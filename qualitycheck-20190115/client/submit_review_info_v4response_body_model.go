// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitReviewInfoV4ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitReviewInfoV4ResponseBody
	GetCode() *string
	SetMessage(v string) *SubmitReviewInfoV4ResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitReviewInfoV4ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitReviewInfoV4ResponseBody
	GetSuccess() *bool
}

type SubmitReviewInfoV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitReviewInfoV4ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitReviewInfoV4ResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitReviewInfoV4ResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitReviewInfoV4ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitReviewInfoV4ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitReviewInfoV4ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitReviewInfoV4ResponseBody) SetCode(v string) *SubmitReviewInfoV4ResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitReviewInfoV4ResponseBody) SetMessage(v string) *SubmitReviewInfoV4ResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitReviewInfoV4ResponseBody) SetRequestId(v string) *SubmitReviewInfoV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitReviewInfoV4ResponseBody) SetSuccess(v bool) *SubmitReviewInfoV4ResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitReviewInfoV4ResponseBody) Validate() error {
	return dara.Validate(s)
}
