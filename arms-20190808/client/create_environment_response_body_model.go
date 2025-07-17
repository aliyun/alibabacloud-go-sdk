// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateEnvironmentResponseBody
	GetCode() *int32
	SetData(v string) *CreateEnvironmentResponseBody
	GetData() *string
	SetMessage(v string) *CreateEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEnvironmentResponseBody
	GetRequestId() *string
}

type CreateEnvironmentResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the created environment.
	//
	// example:
	//
	// env-xxxxx
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
	// 16AF921B-8187-489F-9913-43C808B4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateEnvironmentResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnvironmentResponseBody) SetCode(v int32) *CreateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetData(v string) *CreateEnvironmentResponseBody {
	s.Data = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetMessage(v string) *CreateEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetRequestId(v string) *CreateEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}
