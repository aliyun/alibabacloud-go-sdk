// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevokeFeedbackResponseBody
	GetCode() *string
	SetMessage(v string) *RevokeFeedbackResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokeFeedbackResponseBody
	GetRequestId() *string
}

type RevokeFeedbackResponseBody struct {
	// Interface status or POP error code. Value explanations are as follows: 2xx: Success. 3xx: Redirect. 4xx: Request error. 5xx: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Public parameter, each request ID is unique and can be used for troubleshooting and problem localization.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevokeFeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeFeedbackResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevokeFeedbackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeFeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeFeedbackResponseBody) SetCode(v string) *RevokeFeedbackResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeFeedbackResponseBody) SetMessage(v string) *RevokeFeedbackResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeFeedbackResponseBody) SetRequestId(v string) *RevokeFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeFeedbackResponseBody) Validate() error {
	return dara.Validate(s)
}
