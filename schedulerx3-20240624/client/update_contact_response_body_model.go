// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateContactResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateContactResponseBody
	GetSuccess() *bool
}

type UpdateContactResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// not support query script history, please upgrade engine version to 2.2.2+
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// F131C3E0-3FAA-5FA4-A6F3-E974D69EF3C5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateContactResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateContactResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateContactResponseBody) SetCode(v int32) *UpdateContactResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateContactResponseBody) SetMessage(v string) *UpdateContactResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateContactResponseBody) SetRequestId(v string) *UpdateContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateContactResponseBody) SetSuccess(v bool) *UpdateContactResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateContactResponseBody) Validate() error {
	return dara.Validate(s)
}
