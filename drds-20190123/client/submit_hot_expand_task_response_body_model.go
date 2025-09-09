// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotExpandTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SubmitHotExpandTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitHotExpandTaskResponseBody
	GetSuccess() *bool
}

type SubmitHotExpandTaskResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0B6B7BDC-575D-4A77-A4F8-24B7EFERV45Y
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitHotExpandTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotExpandTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotExpandTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitHotExpandTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitHotExpandTaskResponseBody) SetRequestId(v string) *SubmitHotExpandTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotExpandTaskResponseBody) SetSuccess(v bool) *SubmitHotExpandTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitHotExpandTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
