// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkItemArtifactAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWorkItemArtifactAccessResponseBody
	GetCode() *string
	SetMessage(v string) *GetWorkItemArtifactAccessResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWorkItemArtifactAccessResponseBody
	GetRequestId() *string
}

type GetWorkItemArtifactAccessResponseBody struct {
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

func (s GetWorkItemArtifactAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWorkItemArtifactAccessResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkItemArtifactAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWorkItemArtifactAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWorkItemArtifactAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWorkItemArtifactAccessResponseBody) SetCode(v string) *GetWorkItemArtifactAccessResponseBody {
	s.Code = &v
	return s
}

func (s *GetWorkItemArtifactAccessResponseBody) SetMessage(v string) *GetWorkItemArtifactAccessResponseBody {
	s.Message = &v
	return s
}

func (s *GetWorkItemArtifactAccessResponseBody) SetRequestId(v string) *GetWorkItemArtifactAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWorkItemArtifactAccessResponseBody) Validate() error {
	return dara.Validate(s)
}
