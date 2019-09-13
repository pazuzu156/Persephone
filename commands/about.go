package commands

import (
	"fmt"
	"persephone/lib"
	"strconv"

	"github.com/andersfylling/disgord"
	"github.com/pazuzu156/aurora"
)

// About command.
type About struct {
	Command
	Version string
}

// Version holds the bot's version number
const Version string = "1.0.1"

// InitAbout initialized the about command.
func InitAbout() About {
	return About{
		Command: Init(&CommandItem{
			Name:        "about",
			Description: "Gets information about the bot",
		}),
		Version: Version,
	}
}

// Register registers and runs the about command.
func (c About) Register() *aurora.Command {
	c.CommandInterface.Run = func(ctx aurora.Context) {
		id, _ := strconv.Atoi(config.BotID)
		bot, _ := ctx.Aurora.GetMember(ctx.Message.GuildID, disgord.NewSnowflake(uint64(id)))

		// Gets roles the bot has, so they can be displayed in
		// the embed
		var roles []string
		for _, r := range bot.Roles {
			groles, _ := ctx.Aurora.GetGuildRoles(ctx.Message.GuildID)

			for _, gr := range groles {
				if gr.ID == r {
					roles = append(roles, gr.Name)
				}
			}
		}

		f, t := c.embedFooter(ctx)
		ctx.Aurora.CreateMessage(ctx.Message.ChannelID, &disgord.CreateMessageParams{
			Embed: &disgord.Embed{
				Title:       "About Persephone",
				Description: fmt.Sprintf("Persephone is a bot written in Go. Version v%s", c.Version),
				Color:       0x7FFF00,
				Thumbnail: &disgord.EmbedThumbnail{
					URL: lib.GenAvatarURL(bot.User),
				},
				Fields: []*disgord.EmbedField{
					{
						Name:  "Name",
						Value: fmt.Sprintf("%s#%s", bot.User.Username, bot.User.Discriminator),
					},
					{
						Name:  "ID",
						Value: bot.User.ID.String(),
					},
					{
						Name:  "Roles",
						Value: lib.JoinString(roles, ", "),
					},
				},
				Footer: f, Timestamp: t,
			},
		})
	}

	return c.CommandInterface
}
