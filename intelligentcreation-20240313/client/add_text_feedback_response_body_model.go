// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTextFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddTextFeedbackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTextFeedbackResponseBody
	GetSuccess() *bool
}

type AddTextFeedbackResponseBody struct {
	// example:
	//
	// 3239281273464326823
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s AddTextFeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTextFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *AddTextFeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTextFeedbackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTextFeedbackResponseBody) SetRequestId(v string) *AddTextFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTextFeedbackResponseBody) SetSuccess(v bool) *AddTextFeedbackResponseBody {
	s.Success = &v
	return s
}

func (s *AddTextFeedbackResponseBody) Validate() error {
	return dara.Validate(s)
}
