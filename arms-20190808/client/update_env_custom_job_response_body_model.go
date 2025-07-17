// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEnvCustomJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateEnvCustomJobResponseBody
	GetCode() *int32
	SetData(v string) *UpdateEnvCustomJobResponseBody
	GetData() *string
	SetMessage(v string) *UpdateEnvCustomJobResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateEnvCustomJobResponseBody
	GetRequestId() *string
}

type UpdateEnvCustomJobResponseBody struct {
	// The status code or error code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the operation.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// message
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C21AB7CF-B7AF-410F-BD61-82D1567F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEnvCustomJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEnvCustomJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEnvCustomJobResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateEnvCustomJobResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateEnvCustomJobResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateEnvCustomJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEnvCustomJobResponseBody) SetCode(v int32) *UpdateEnvCustomJobResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateEnvCustomJobResponseBody) SetData(v string) *UpdateEnvCustomJobResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateEnvCustomJobResponseBody) SetMessage(v string) *UpdateEnvCustomJobResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateEnvCustomJobResponseBody) SetRequestId(v string) *UpdateEnvCustomJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEnvCustomJobResponseBody) Validate() error {
	return dara.Validate(s)
}
