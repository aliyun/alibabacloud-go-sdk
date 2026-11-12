// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAutoPassPolicy interface {
	dara.Model
	String() string
	GoString() string
	SetCommandPrefix(v []*string) *AutoPassPolicy
	GetCommandPrefix() []*string
	SetFullCommand(v string) *AutoPassPolicy
	GetFullCommand() *string
	SetVersion(v string) *AutoPassPolicy
	GetVersion() *string
}

type AutoPassPolicy struct {
	// The selected command prefix. The array must contain exactly one prefix. You must specify either this parameter or fullCommand.
	//
	// example:
	//
	// ["starops sls"]
	CommandPrefix []*string `json:"commandPrefix,omitempty" xml:"commandPrefix,omitempty" type:"Repeated"`
	// The full command for exact matching. The command must not be empty. You must specify either this parameter or commandPrefix.
	//
	// example:
	//
	// starops tool list
	FullCommand *string `json:"fullCommand,omitempty" xml:"fullCommand,omitempty"`
	// The version required for a valid rule. The current version is v1.0.0.
	//
	// example:
	//
	// v1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s AutoPassPolicy) String() string {
	return dara.Prettify(s)
}

func (s AutoPassPolicy) GoString() string {
	return s.String()
}

func (s *AutoPassPolicy) GetCommandPrefix() []*string {
	return s.CommandPrefix
}

func (s *AutoPassPolicy) GetFullCommand() *string {
	return s.FullCommand
}

func (s *AutoPassPolicy) GetVersion() *string {
	return s.Version
}

func (s *AutoPassPolicy) SetCommandPrefix(v []*string) *AutoPassPolicy {
	s.CommandPrefix = v
	return s
}

func (s *AutoPassPolicy) SetFullCommand(v string) *AutoPassPolicy {
	s.FullCommand = &v
	return s
}

func (s *AutoPassPolicy) SetVersion(v string) *AutoPassPolicy {
	s.Version = &v
	return s
}

func (s *AutoPassPolicy) Validate() error {
	return dara.Validate(s)
}
