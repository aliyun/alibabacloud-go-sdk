// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendRenderingInstanceCommandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommands(v string) *SendRenderingInstanceCommandsRequest
	GetCommands() *string
	SetMode(v string) *SendRenderingInstanceCommandsRequest
	GetMode() *string
	SetRenderingInstanceId(v string) *SendRenderingInstanceCommandsRequest
	GetRenderingInstanceId() *string
	SetTimeout(v int32) *SendRenderingInstanceCommandsRequest
	GetTimeout() *int32
}

type SendRenderingInstanceCommandsRequest struct {
	// A shell command string. Enter multiple commands separated by semicolons (;) or line feeds.
	//
	// - Dangerous commands such as rm and reboot are disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// date;ls -l /tmp
	Commands *string `json:"Commands,omitempty" xml:"Commands,omitempty"`
	// The response pattern for the command. Valid values:
	//
	// 1. Sync: The response is returned synchronously. This is the default value.
	//
	// 2. Async: The response is returned asynchronously.
	//
	// example:
	//
	// Async
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the cloud application service instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	// The timeout period for command execution, in seconds. The value range depends on the Mode parameter:
	//
	// 1. If Mode is set to Sync, the value range is 0 to 30. The default value is 30.
	//
	// 2. If Mode is set to Async, the value range is 0 to 3600. The default value is 300.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s SendRenderingInstanceCommandsRequest) String() string {
	return dara.Prettify(s)
}

func (s SendRenderingInstanceCommandsRequest) GoString() string {
	return s.String()
}

func (s *SendRenderingInstanceCommandsRequest) GetCommands() *string {
	return s.Commands
}

func (s *SendRenderingInstanceCommandsRequest) GetMode() *string {
	return s.Mode
}

func (s *SendRenderingInstanceCommandsRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *SendRenderingInstanceCommandsRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *SendRenderingInstanceCommandsRequest) SetCommands(v string) *SendRenderingInstanceCommandsRequest {
	s.Commands = &v
	return s
}

func (s *SendRenderingInstanceCommandsRequest) SetMode(v string) *SendRenderingInstanceCommandsRequest {
	s.Mode = &v
	return s
}

func (s *SendRenderingInstanceCommandsRequest) SetRenderingInstanceId(v string) *SendRenderingInstanceCommandsRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *SendRenderingInstanceCommandsRequest) SetTimeout(v int32) *SendRenderingInstanceCommandsRequest {
	s.Timeout = &v
	return s
}

func (s *SendRenderingInstanceCommandsRequest) Validate() error {
	return dara.Validate(s)
}
