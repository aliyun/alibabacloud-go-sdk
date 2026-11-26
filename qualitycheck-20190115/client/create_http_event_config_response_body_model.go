// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHttpEventConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateHttpEventConfigResponseBody
	GetCode() *string
	SetData(v string) *CreateHttpEventConfigResponseBody
	GetData() *string
	SetMessage(v string) *CreateHttpEventConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHttpEventConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateHttpEventConfigResponseBody
	GetSuccess() *bool
}

type CreateHttpEventConfigResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
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

func (s CreateHttpEventConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHttpEventConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHttpEventConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateHttpEventConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateHttpEventConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHttpEventConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHttpEventConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateHttpEventConfigResponseBody) SetCode(v string) *CreateHttpEventConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateHttpEventConfigResponseBody) SetData(v string) *CreateHttpEventConfigResponseBody {
	s.Data = &v
	return s
}

func (s *CreateHttpEventConfigResponseBody) SetMessage(v string) *CreateHttpEventConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHttpEventConfigResponseBody) SetRequestId(v string) *CreateHttpEventConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHttpEventConfigResponseBody) SetSuccess(v bool) *CreateHttpEventConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateHttpEventConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
