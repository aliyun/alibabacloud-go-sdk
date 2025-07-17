// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvCustomJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateEnvCustomJobResponseBody
	GetCode() *int32
	SetData(v string) *CreateEnvCustomJobResponseBody
	GetData() *string
	SetMessage(v string) *CreateEnvCustomJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEnvCustomJobResponseBody
	GetRequestId() *string
}

type CreateEnvCustomJobResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The name of the custom job that was created, or the exception information.
	//
	// example:
	//
	// cutomJob1
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4C518054-852F-4023-ABC1-4AF95FF7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnvCustomJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvCustomJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvCustomJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateEnvCustomJobResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateEnvCustomJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEnvCustomJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnvCustomJobResponseBody) SetCode(v int32) *CreateEnvCustomJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnvCustomJobResponseBody) SetData(v string) *CreateEnvCustomJobResponseBody {
	s.Data = &v
	return s
}

func (s *CreateEnvCustomJobResponseBody) SetMessage(v string) *CreateEnvCustomJobResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEnvCustomJobResponseBody) SetRequestId(v string) *CreateEnvCustomJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnvCustomJobResponseBody) Validate() error {
	return dara.Validate(s)
}
