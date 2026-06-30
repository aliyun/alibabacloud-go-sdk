// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPrecisionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitPrecisionTaskResponseBody
	GetCode() *string
	SetData(v string) *SubmitPrecisionTaskResponseBody
	GetData() *string
	SetMessage(v string) *SubmitPrecisionTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitPrecisionTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitPrecisionTaskResponseBody
	GetSuccess() *bool
}

type SubmitPrecisionTaskResponseBody struct {
	// Result code. A value of **200*	- indicates success. Any other value indicates failure. Use this field to identify the cause of failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ID of the created task.
	//
	// example:
	//
	// EA701F66-8CA2-4A79-8E3C-A6F2FA****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Error details if the call fails. Returns successful if the call succeeds.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. true means success. false or null means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitPrecisionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitPrecisionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPrecisionTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitPrecisionTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitPrecisionTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitPrecisionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitPrecisionTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitPrecisionTaskResponseBody) SetCode(v string) *SubmitPrecisionTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetData(v string) *SubmitPrecisionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetMessage(v string) *SubmitPrecisionTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetRequestId(v string) *SubmitPrecisionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) SetSuccess(v bool) *SubmitPrecisionTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitPrecisionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
