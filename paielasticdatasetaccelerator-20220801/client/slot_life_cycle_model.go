// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSlotLifeCycle interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *SlotLifeCycle
	GetConfig() *string
	SetType(v string) *SlotLifeCycle
	GetType() *string
}

type SlotLifeCycle struct {
	// example:
	//
	// {}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// Running
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SlotLifeCycle) String() string {
	return dara.Prettify(s)
}

func (s SlotLifeCycle) GoString() string {
	return s.String()
}

func (s *SlotLifeCycle) GetConfig() *string {
	return s.Config
}

func (s *SlotLifeCycle) GetType() *string {
	return s.Type
}

func (s *SlotLifeCycle) SetConfig(v string) *SlotLifeCycle {
	s.Config = &v
	return s
}

func (s *SlotLifeCycle) SetType(v string) *SlotLifeCycle {
	s.Type = &v
	return s
}

func (s *SlotLifeCycle) Validate() error {
	return dara.Validate(s)
}
