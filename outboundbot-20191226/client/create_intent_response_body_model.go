// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateIntentResponseBody
	GetHttpStatusCode() *int32
	SetIntentId(v string) *CreateIntentResponseBody
	GetIntentId() *string
	SetMessage(v string) *CreateIntentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateIntentResponseBody
	GetSuccess() *bool
}

type CreateIntentResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 390515b5-6115-4ccf-83e2-52d5bfaf2ddf
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Response message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation succeeded
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIntentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateIntentResponseBody) GetIntentId() *string {
	return s.IntentId
}

func (s *CreateIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateIntentResponseBody) SetCode(v string) *CreateIntentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateIntentResponseBody) SetHttpStatusCode(v int32) *CreateIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateIntentResponseBody) SetIntentId(v string) *CreateIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *CreateIntentResponseBody) SetMessage(v string) *CreateIntentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateIntentResponseBody) SetRequestId(v string) *CreateIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIntentResponseBody) SetSuccess(v bool) *CreateIntentResponseBody {
	s.Success = &v
	return s
}

func (s *CreateIntentResponseBody) Validate() error {
	return dara.Validate(s)
}
