// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHttpEventConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateHttpEventConfigResponseBody
	GetCode() *string
	SetData(v string) *UpdateHttpEventConfigResponseBody
	GetData() *string
	SetMessage(v string) *UpdateHttpEventConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHttpEventConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateHttpEventConfigResponseBody
	GetSuccess() *bool
}

type UpdateHttpEventConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 4
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateHttpEventConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHttpEventConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHttpEventConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateHttpEventConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateHttpEventConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHttpEventConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHttpEventConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateHttpEventConfigResponseBody) SetCode(v string) *UpdateHttpEventConfigResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHttpEventConfigResponseBody) SetData(v string) *UpdateHttpEventConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateHttpEventConfigResponseBody) SetMessage(v string) *UpdateHttpEventConfigResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHttpEventConfigResponseBody) SetRequestId(v string) *UpdateHttpEventConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHttpEventConfigResponseBody) SetSuccess(v bool) *UpdateHttpEventConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateHttpEventConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
