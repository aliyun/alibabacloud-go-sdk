// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyFileResponseBody
	GetCode() *string
	SetData(v float32) *VerifyFileResponseBody
	GetData() *float32
	SetMessage(v string) *VerifyFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyFileResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VerifyFileResponseBody
	GetSuccess() *bool
}

type VerifyFileResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 0.9485294
	Data *float32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// s
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VerifyFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyFileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyFileResponseBody) GetData() *float32 {
	return s.Data
}

func (s *VerifyFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyFileResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VerifyFileResponseBody) SetCode(v string) *VerifyFileResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyFileResponseBody) SetData(v float32) *VerifyFileResponseBody {
	s.Data = &v
	return s
}

func (s *VerifyFileResponseBody) SetMessage(v string) *VerifyFileResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyFileResponseBody) SetRequestId(v string) *VerifyFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyFileResponseBody) SetSuccess(v bool) *VerifyFileResponseBody {
	s.Success = &v
	return s
}

func (s *VerifyFileResponseBody) Validate() error {
	return dara.Validate(s)
}
