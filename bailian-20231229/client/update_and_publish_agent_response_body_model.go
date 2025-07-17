// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndPublishAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAndPublishAgentResponseBody
	GetCode() *string
	SetData(v string) *UpdateAndPublishAgentResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateAndPublishAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateAndPublishAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAndPublishAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAndPublishAgentResponseBody
	GetSuccess() *bool
}

type UpdateAndPublishAgentResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAndPublishAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAndPublishAgentResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateAndPublishAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAndPublishAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAndPublishAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAndPublishAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAndPublishAgentResponseBody) SetCode(v string) *UpdateAndPublishAgentResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetData(v string) *UpdateAndPublishAgentResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetHttpStatusCode(v int32) *UpdateAndPublishAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetMessage(v string) *UpdateAndPublishAgentResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetRequestId(v string) *UpdateAndPublishAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) SetSuccess(v bool) *UpdateAndPublishAgentResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAndPublishAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
