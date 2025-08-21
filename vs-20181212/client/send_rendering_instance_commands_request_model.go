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
	// This parameter is required.
	//
	// example:
	//
	// date;ls -l /tmp
	Commands *string `json:"Commands,omitempty" xml:"Commands,omitempty"`
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
	Timeout             *int32  `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
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
