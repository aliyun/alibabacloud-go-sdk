// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaslUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateSaslUserResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateSaslUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSaslUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSaslUserResponseBody
	GetSuccess() *bool
}

type CreateSaslUserResponseBody struct {
	// The HTTP status code. The HTTP status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C5CA600C-7D5A-45B5-B6DB-44FAC2C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSaslUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSaslUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSaslUserResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateSaslUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSaslUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSaslUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSaslUserResponseBody) SetCode(v int32) *CreateSaslUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSaslUserResponseBody) SetMessage(v string) *CreateSaslUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSaslUserResponseBody) SetRequestId(v string) *CreateSaslUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSaslUserResponseBody) SetSuccess(v bool) *CreateSaslUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSaslUserResponseBody) Validate() error {
	return dara.Validate(s)
}
