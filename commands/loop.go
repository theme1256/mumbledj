/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/loop.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
    "errors"
    "fmt"

    "github.com/layeh/gumble/gumble"
    "github.com/matthieugrieger/mumbledj/interfaces"
    "github.com/spf13/viper"
)

// loop is a command that toggles continuous playback if no tracks are in the queue
type loop struct{}

// Aliases returns the current aliases for the command.
func (c *loop) Aliases() []string {
    return viper.GetStringSlice("commands.loop.aliases")
}

// Description returns the description for the command.
func (c *loop) Description() string {
    return viper.GetString("commands.loop.description")
}

// IsAdminCommand returns true if the command is only for admin use, and
// returns false otherwise.
func (c *loop) IsAdminCommand() bool {
    return viper.GetBool("commands.loop.is_admin")
}

// Execute executes the command with the given user and arguments.
// Return value descriptions:
//    string: A message to be returned to the user upon successful execution.
//    bool:   Whether the message should be private or not. true = private,
//            false = public (sent to whole channel).
//    error:  An error message to be returned upon unsuccessful execution.
//            If no error has occurred, pass nil instead.
// Example return statement:
//    return "This is a private message!", true, nil
func (c *loop) Execute(user *gumble.User, args ...string) (string, bool, error) {
	length := DJ.Queue.Length()
}