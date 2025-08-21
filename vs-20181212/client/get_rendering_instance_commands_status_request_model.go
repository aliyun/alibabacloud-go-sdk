// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingInstanceCommandsStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCmdId(v string) *GetRenderingInstanceCommandsStatusRequest
	GetCmdId() *string
	SetRenderingInstanceId(v string) *GetRenderingInstanceCommandsStatusRequest
	GetRenderingInstanceId() *string
}

type GetRenderingInstanceCommandsStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cmd-81de027b66e442e99c1e0e09a16a0be5
	CmdId *string `json:"CmdId,omitempty" xml:"CmdId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// render-9f8c57355d224ad7beaf95e145f22111
	RenderingInstanceId *string `json:"RenderingInstanceId,omitempty" xml:"RenderingInstanceId,omitempty"`
}

func (s GetRenderingInstanceCommandsStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingInstanceCommandsStatusRequest) GoString() string {
	return s.String()
}

func (s *GetRenderingInstanceCommandsStatusRequest) GetCmdId() *string {
	return s.CmdId
}

func (s *GetRenderingInstanceCommandsStatusRequest) GetRenderingInstanceId() *string {
	return s.RenderingInstanceId
}

func (s *GetRenderingInstanceCommandsStatusRequest) SetCmdId(v string) *GetRenderingInstanceCommandsStatusRequest {
	s.CmdId = &v
	return s
}

func (s *GetRenderingInstanceCommandsStatusRequest) SetRenderingInstanceId(v string) *GetRenderingInstanceCommandsStatusRequest {
	s.RenderingInstanceId = &v
	return s
}

func (s *GetRenderingInstanceCommandsStatusRequest) Validate() error {
	return dara.Validate(s)
}
