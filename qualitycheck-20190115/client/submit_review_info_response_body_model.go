// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitReviewInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitReviewInfoResponseBody
	GetCode() *string
	SetData(v string) *SubmitReviewInfoResponseBody
	GetData() *string
	SetMessage(v string) *SubmitReviewInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitReviewInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitReviewInfoResponseBody
	GetSuccess() *bool
}

type SubmitReviewInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 95
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

func (s SubmitReviewInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitReviewInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitReviewInfoResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitReviewInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitReviewInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitReviewInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitReviewInfoResponseBody) SetCode(v string) *SubmitReviewInfoResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetData(v string) *SubmitReviewInfoResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetMessage(v string) *SubmitReviewInfoResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetRequestId(v string) *SubmitReviewInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) SetSuccess(v bool) *SubmitReviewInfoResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitReviewInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
