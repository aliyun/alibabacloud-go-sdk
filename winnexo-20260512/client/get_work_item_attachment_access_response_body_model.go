// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemAttachmentAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkItemAttachmentAccessResponseBody
	GetCode() *string
	SetMessage(v string) *GetWorkItemAttachmentAccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkItemAttachmentAccessResponseBody
	GetRequestId() *string
}

type GetWorkItemAttachmentAccessResponseBody struct {
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

func (s GetWorkItemAttachmentAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemAttachmentAccessResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemAttachmentAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkItemAttachmentAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkItemAttachmentAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemAttachmentAccessResponseBody) SetCode(v string) *GetWorkItemAttachmentAccessResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkItemAttachmentAccessResponseBody) SetMessage(v string) *GetWorkItemAttachmentAccessResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkItemAttachmentAccessResponseBody) SetRequestId(v string) *GetWorkItemAttachmentAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemAttachmentAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
