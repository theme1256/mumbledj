/*
 * MumbleDJ
 * By Matthieu Grieger
 * commands/yournewcommand_test.go
 * Copyright (c) 2016 Matthieu Grieger (MIT License)
 */

package commands

import (
    "testing"

    "github.com/layeh/gumble/gumbleffmpeg"
    "github.com/matthieugrieger/mumbledj/bot"
    "github.com/spf13/viper"
    "github.com/stretchr/testify/suite"
)

type loopTestSuite struct {
    Command loop
    suite.Suite
}

func (suite *loopTestSuite) SetupSuite() {
    DJ = bot.NewMumbleDJ()
    bot.DJ = DJ

    viper.Set("commands.loop.aliases", []string{"loop", "c"})
    viper.Set("commands.loop.description", "loop")
    viper.Set("commands.loop.is_admin", false)
}

func (suite *loopTestSuite) SetupTest() {
    DJ.Queue = bot.NewQueue()
}

func (suite *loopTestSuite) TestAliases() {
    suite.Equal([]string{"loop", "c"}, suite.Command.Aliases())
}

func (suite *loopTestSuite) TestDescription() {
    suite.Equal("loop", suite.Command.Description())
}

func (suite *loopTestSuite) TestIsAdminCommand() {
    suite.False(suite.Command.IsAdminCommand())
}

// Implement more tests here as necessary! It may be helpful to take a look
// at the stretchr/testify documentation:
// https://github.com/stretchr/testify
// Remove this comment before sending a pull request.

func TestloopTestSuite(t *testing.T) {
    suite.Run(t, new(loopTestSuite))
}