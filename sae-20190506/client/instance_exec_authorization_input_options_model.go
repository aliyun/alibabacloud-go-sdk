// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceExecAuthorizationInputOptions interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v []*string) *InstanceExecAuthorizationInputOptions
	GetCommand() []*string
	SetStderr(v bool) *InstanceExecAuthorizationInputOptions
	GetStderr() *bool
	SetStdin(v bool) *InstanceExecAuthorizationInputOptions
	GetStdin() *bool
	SetStdout(v bool) *InstanceExecAuthorizationInputOptions
	GetStdout() *bool
	SetTty(v bool) *InstanceExecAuthorizationInputOptions
	GetTty() *bool
}

type InstanceExecAuthorizationInputOptions struct {
	Command []*string `json:"command,omitempty" xml:"command,omitempty" type:"Repeated"`
	Stderr  *bool     `json:"stderr,omitempty" xml:"stderr,omitempty"`
	Stdin   *bool     `json:"stdin,omitempty" xml:"stdin,omitempty"`
	Stdout  *bool     `json:"stdout,omitempty" xml:"stdout,omitempty"`
	Tty     *bool     `json:"tty,omitempty" xml:"tty,omitempty"`
}

func (s InstanceExecAuthorizationInputOptions) String() string {
	return dara.Prettify(s)
}

func (s InstanceExecAuthorizationInputOptions) GoString() string {
	return s.String()
}

func (s *InstanceExecAuthorizationInputOptions) GetCommand() []*string {
	return s.Command
}

func (s *InstanceExecAuthorizationInputOptions) GetStderr() *bool {
	return s.Stderr
}

func (s *InstanceExecAuthorizationInputOptions) GetStdin() *bool {
	return s.Stdin
}

func (s *InstanceExecAuthorizationInputOptions) GetStdout() *bool {
	return s.Stdout
}

func (s *InstanceExecAuthorizationInputOptions) GetTty() *bool {
	return s.Tty
}

func (s *InstanceExecAuthorizationInputOptions) SetCommand(v []*string) *InstanceExecAuthorizationInputOptions {
	s.Command = v
	return s
}

func (s *InstanceExecAuthorizationInputOptions) SetStderr(v bool) *InstanceExecAuthorizationInputOptions {
	s.Stderr = &v
	return s
}

func (s *InstanceExecAuthorizationInputOptions) SetStdin(v bool) *InstanceExecAuthorizationInputOptions {
	s.Stdin = &v
	return s
}

func (s *InstanceExecAuthorizationInputOptions) SetStdout(v bool) *InstanceExecAuthorizationInputOptions {
	s.Stdout = &v
	return s
}

func (s *InstanceExecAuthorizationInputOptions) SetTty(v bool) *InstanceExecAuthorizationInputOptions {
	s.Tty = &v
	return s
}

func (s *InstanceExecAuthorizationInputOptions) Validate() error {
	return dara.Validate(s)
}
