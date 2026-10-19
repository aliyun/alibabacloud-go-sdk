// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteWorkItemResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteWorkItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteWorkItemResponseBody
	GetRequestId() *string
}

type DeleteWorkItemResponseBody struct {
	// 业务状态码；成功为200
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteWorkItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWorkItemResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteWorkItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteWorkItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWorkItemResponseBody) SetCode(v string) *DeleteWorkItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteWorkItemResponseBody) SetMessage(v string) *DeleteWorkItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteWorkItemResponseBody) SetRequestId(v string) *DeleteWorkItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWorkItemResponseBody) Validate() error {
	return dara.Validate(s)
}
