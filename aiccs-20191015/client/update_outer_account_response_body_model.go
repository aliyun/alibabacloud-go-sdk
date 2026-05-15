// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOuterAccountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateOuterAccountResponseBody
	GetCode() *string
	SetData(v string) *UpdateOuterAccountResponseBody
	GetData() *string
	SetMessage(v string) *UpdateOuterAccountResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateOuterAccountResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateOuterAccountResponseBody
	GetSuccess() *bool
}

type UpdateOuterAccountResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateOuterAccountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateOuterAccountResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateOuterAccountResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateOuterAccountResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateOuterAccountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateOuterAccountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateOuterAccountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateOuterAccountResponseBody) SetCode(v string) *UpdateOuterAccountResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) SetData(v string) *UpdateOuterAccountResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) SetMessage(v string) *UpdateOuterAccountResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) SetRequestId(v string) *UpdateOuterAccountResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) SetSuccess(v bool) *UpdateOuterAccountResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateOuterAccountResponseBody) Validate() error {
	return dara.Validate(s)
}
