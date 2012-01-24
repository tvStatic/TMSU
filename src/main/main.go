/*
Copyright 2011 Paul Ruane.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package main

import (
    "os"
)

var commands map[string]Command

func main() {
	commandArray := []Command{
		DeleteCommand{},
		DupesCommand{},
		ExportCommand{},
		FilesCommand{},
		HelpCommand{},
		MergeCommand{},
		MountCommand{},
		RenameCommand{},
		StatsCommand{},
		StatusCommand{},
		TagCommand{},
		TagsCommand{},
		UnmountCommand{},
		UntagCommand{},
		VersionCommand{},
		VfsCommand{},
	}

	commands = make(map[string]Command, len(commandArray))
	for _, command := range commandArray {
		commands[command.Name()] = command
	}

    args := os.Args[1:] // strip off binary name

	var commandName string
	if len(args) > 0 {
		commandName = args[0]
		args = args[1:]
	} else {
		commandName = "help"
	}

	command := commands[commandName]
	if command == nil {
		fatal("unknown command '%v'.", commandName)
	}

	err := command.Exec(args)
	if err != nil {
		fatal(err.Error())
	}
}
