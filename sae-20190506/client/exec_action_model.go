// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExecAction interface {
  dara.Model
  String() string
  GoString() string
  SetCommand(v []*string) *ExecAction
  GetCommand() []*string 
}

type ExecAction struct {
  Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
}

func (s ExecAction) String() string {
  return dara.Prettify(s)
}

func (s ExecAction) GoString() string {
  return s.String()
}

func (s *ExecAction) GetCommand() []*string  {
  return s.Command
}

func (s *ExecAction) SetCommand(v []*string) *ExecAction {
  s.Command = v
  return s
}

func (s *ExecAction) Validate() error {
  return dara.Validate(s)
}

