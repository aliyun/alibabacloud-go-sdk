// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLifecycleHook interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v []*string) *LifecycleHook
	GetCommand() []*string
	SetHandler(v string) *LifecycleHook
	GetHandler() *string
	SetTimeout(v int32) *LifecycleHook
	GetTimeout() *int32
}

type LifecycleHook struct {
	Command []*string `json:"command" xml:"command" type:"Repeated"`
	// example:
	//
	// index.initializer
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// example:
	//
	// 10
	Timeout *int32 `json:"timeout,omitempty" xml:"timeout,omitempty"`
}

func (s LifecycleHook) String() string {
	return dara.Prettify(s)
}

func (s LifecycleHook) GoString() string {
	return s.String()
}

func (s *LifecycleHook) GetCommand() []*string {
	return s.Command
}

func (s *LifecycleHook) GetHandler() *string {
	return s.Handler
}

func (s *LifecycleHook) GetTimeout() *int32 {
	return s.Timeout
}

func (s *LifecycleHook) SetCommand(v []*string) *LifecycleHook {
	s.Command = v
	return s
}

func (s *LifecycleHook) SetHandler(v string) *LifecycleHook {
	s.Handler = &v
	return s
}

func (s *LifecycleHook) SetTimeout(v int32) *LifecycleHook {
	s.Timeout = &v
	return s
}

func (s *LifecycleHook) Validate() error {
	return dara.Validate(s)
}
