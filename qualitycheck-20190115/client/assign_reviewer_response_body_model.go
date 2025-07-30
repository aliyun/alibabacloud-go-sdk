// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssignReviewerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AssignReviewerResponseBody
	GetCode() *string
	SetMessage(v string) *AssignReviewerResponseBody
	GetMessage() *string
	SetRequestId(v string) *AssignReviewerResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AssignReviewerResponseBody
	GetSuccess() *bool
}

type AssignReviewerResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AssignReviewerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssignReviewerResponseBody) GoString() string {
	return s.String()
}

func (s *AssignReviewerResponseBody) GetCode() *string {
	return s.Code
}

func (s *AssignReviewerResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AssignReviewerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssignReviewerResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AssignReviewerResponseBody) SetCode(v string) *AssignReviewerResponseBody {
	s.Code = &v
	return s
}

func (s *AssignReviewerResponseBody) SetMessage(v string) *AssignReviewerResponseBody {
	s.Message = &v
	return s
}

func (s *AssignReviewerResponseBody) SetRequestId(v string) *AssignReviewerResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssignReviewerResponseBody) SetSuccess(v bool) *AssignReviewerResponseBody {
	s.Success = &v
	return s
}

func (s *AssignReviewerResponseBody) Validate() error {
	return dara.Validate(s)
}
