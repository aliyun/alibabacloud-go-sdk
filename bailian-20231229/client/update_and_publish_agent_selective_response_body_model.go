// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAndPublishAgentSelectiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateAndPublishAgentSelectiveResponseBody
	GetCode() *string
	SetData(v string) *UpdateAndPublishAgentSelectiveResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *UpdateAndPublishAgentSelectiveResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateAndPublishAgentSelectiveResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAndPublishAgentSelectiveResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAndPublishAgentSelectiveResponseBody
	GetSuccess() *bool
}

type UpdateAndPublishAgentSelectiveResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateAndPublishAgentSelectiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAndPublishAgentSelectiveResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) SetCode(v string) *UpdateAndPublishAgentSelectiveResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) SetData(v string) *UpdateAndPublishAgentSelectiveResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) SetHttpStatusCode(v int32) *UpdateAndPublishAgentSelectiveResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) SetMessage(v string) *UpdateAndPublishAgentSelectiveResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) SetRequestId(v string) *UpdateAndPublishAgentSelectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) SetSuccess(v bool) *UpdateAndPublishAgentSelectiveResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAndPublishAgentSelectiveResponseBody) Validate() error {
	return dara.Validate(s)
}
