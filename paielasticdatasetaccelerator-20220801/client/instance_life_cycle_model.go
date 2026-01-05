// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceLifeCycle interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v string) *InstanceLifeCycle
	GetConfig() *string
	SetType(v string) *InstanceLifeCycle
	GetType() *string
}

type InstanceLifeCycle struct {
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// example:
	//
	// Running
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s InstanceLifeCycle) String() string {
	return dara.Prettify(s)
}

func (s InstanceLifeCycle) GoString() string {
	return s.String()
}

func (s *InstanceLifeCycle) GetConfig() *string {
	return s.Config
}

func (s *InstanceLifeCycle) GetType() *string {
	return s.Type
}

func (s *InstanceLifeCycle) SetConfig(v string) *InstanceLifeCycle {
	s.Config = &v
	return s
}

func (s *InstanceLifeCycle) SetType(v string) *InstanceLifeCycle {
	s.Type = &v
	return s
}

func (s *InstanceLifeCycle) Validate() error {
	return dara.Validate(s)
}
