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
	// Result code. **200*	- indicates success. Other values indicate failure. Callers can determine the failure reason using this field.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Current detection accuracy: Number of incorrect characters in verified files / Total number of characters in verified files.
	//
	// example:
	//
	// 0.9485294
	Data *float32 `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error details when an error occurs. Successful when successful.
	//
	// example:
	//
	// s
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Callers can determine if the request was successful using this field: true indicates success; false/null indicates failure.
	//
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
