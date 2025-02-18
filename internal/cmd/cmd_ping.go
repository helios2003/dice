// Copyright (c) 2022-present, DiceDB contributors
// All rights reserved. Licensed under the BSD 3-Clause License. See LICENSE file in the project root for full license information.

package cmd

import (
	dstore "github.com/dicedb/dice/internal/store"
	"github.com/dicedb/dicedb-go/wire"
)

var cPING = &DiceDBCommand{
	Name:      "PING",
	HelpShort: "PING returns with an encoded \"PONG\" if no message is added with the ping command, the message will be returned.",
	Eval:      evalPING,
}

func init() {
	commandRegistry.AddCommand(cPING)
}

func evalPING(c *Cmd, s *dstore.Store) (*CmdRes, error) {
	if len(c.C.Args) >= 2 {
		return cmdResNil, errWrongArgumentCount("PING")
	}
	if len(c.C.Args) == 0 {
		return &CmdRes{R: &wire.Response{
			Value: &wire.Response_VStr{VStr: "PONG"},
		}}, nil
	}
	return &CmdRes{R: &wire.Response{
		Value: &wire.Response_VStr{VStr: c.C.Args[0]},
	}}, nil
}
