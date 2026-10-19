// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *SetAccountAttributesResponseBody
	GetCode() *int64
	SetMessage(v string) *SetAccountAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetAccountAttributesResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetAccountAttributesResponseBody
	GetStatus() *string
	SetSuccess(v bool) *SetAccountAttributesResponseBody
	GetSuccess() *bool
}

type SetAccountAttributesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 06273500-249F-5863-121D-74D51123E62C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response status.
	//
	// example:
	//
	// Success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetAccountAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAccountAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *SetAccountAttributesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *SetAccountAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetAccountAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAccountAttributesResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetAccountAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetAccountAttributesResponseBody) SetCode(v int64) *SetAccountAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *SetAccountAttributesResponseBody) SetMessage(v string) *SetAccountAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *SetAccountAttributesResponseBody) SetRequestId(v string) *SetAccountAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAccountAttributesResponseBody) SetStatus(v string) *SetAccountAttributesResponseBody {
	s.Status = &v
	return s
}

func (s *SetAccountAttributesResponseBody) SetSuccess(v bool) *SetAccountAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *SetAccountAttributesResponseBody) Validate() error {
	return dara.Validate(s)
}
