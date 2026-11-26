// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestHttpEventConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TestHttpEventConfigResponseBody
	GetCode() *string
	SetData(v string) *TestHttpEventConfigResponseBody
	GetData() *string
	SetMessage(v string) *TestHttpEventConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *TestHttpEventConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TestHttpEventConfigResponseBody
	GetSuccess() *bool
}

type TestHttpEventConfigResponseBody struct {
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

func (s TestHttpEventConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TestHttpEventConfigResponseBody) GoString() string {
	return s.String()
}

func (s *TestHttpEventConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *TestHttpEventConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *TestHttpEventConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TestHttpEventConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TestHttpEventConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TestHttpEventConfigResponseBody) SetCode(v string) *TestHttpEventConfigResponseBody {
	s.Code = &v
	return s
}

func (s *TestHttpEventConfigResponseBody) SetData(v string) *TestHttpEventConfigResponseBody {
	s.Data = &v
	return s
}

func (s *TestHttpEventConfigResponseBody) SetMessage(v string) *TestHttpEventConfigResponseBody {
	s.Message = &v
	return s
}

func (s *TestHttpEventConfigResponseBody) SetRequestId(v string) *TestHttpEventConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *TestHttpEventConfigResponseBody) SetSuccess(v bool) *TestHttpEventConfigResponseBody {
	s.Success = &v
	return s
}

func (s *TestHttpEventConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
