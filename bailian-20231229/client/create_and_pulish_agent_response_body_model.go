// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndPulishAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateAndPulishAgentResponseBody
	GetCode() *string
	SetData(v string) *CreateAndPulishAgentResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *CreateAndPulishAgentResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateAndPulishAgentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAndPulishAgentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAndPulishAgentResponseBody
	GetSuccess() *bool
}

type CreateAndPulishAgentResponseBody struct {
	Code           *string `json:"code,omitempty" xml:"code,omitempty"`
	Data           *string `json:"data,omitempty" xml:"data,omitempty"`
	HttpStatusCode *int32  `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId      *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success        *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateAndPulishAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAndPulishAgentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAndPulishAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateAndPulishAgentResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateAndPulishAgentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateAndPulishAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAndPulishAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAndPulishAgentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAndPulishAgentResponseBody) SetCode(v string) *CreateAndPulishAgentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetData(v string) *CreateAndPulishAgentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetHttpStatusCode(v int32) *CreateAndPulishAgentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetMessage(v string) *CreateAndPulishAgentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetRequestId(v string) *CreateAndPulishAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) SetSuccess(v bool) *CreateAndPulishAgentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAndPulishAgentResponseBody) Validate() error {
	return dara.Validate(s)
}
