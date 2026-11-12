// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitImageResponseBody
	GetCode() *string
	SetData(v bool) *SubmitImageResponseBody
	GetData() *bool
	SetMessage(v string) *SubmitImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitImageResponseBody
	GetSuccess() *bool
}

type SubmitImageResponseBody struct {
	// The error code.
	//
	// example:
	//
	// InvalidParameter.ProductType
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The processing result.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// NULL
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitImageResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *SubmitImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitImageResponseBody) SetCode(v string) *SubmitImageResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitImageResponseBody) SetData(v bool) *SubmitImageResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitImageResponseBody) SetMessage(v string) *SubmitImageResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitImageResponseBody) SetRequestId(v string) *SubmitImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitImageResponseBody) SetSuccess(v bool) *SubmitImageResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitImageResponseBody) Validate() error {
	return dara.Validate(s)
}
