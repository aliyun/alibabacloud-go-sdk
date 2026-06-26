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
	// 函数生命周期初始化阶段回调指令，生命周期回调方法的执行入口 handler 和 command 不允许同时配置，只能有一个生效，同时配置会产生错误提示
	Command []*string `json:"command" xml:"command" type:"Repeated"`
	// The handler of the hook. The definition is similar to that of a request handler.
	//
	// example:
	//
	// index.initializer
	Handler *string `json:"handler,omitempty" xml:"handler,omitempty"`
	// The timeout period of the hook. Unit: seconds.
	//
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
