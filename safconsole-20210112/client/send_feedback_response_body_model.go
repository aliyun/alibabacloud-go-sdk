// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SendFeedbackResponseBody
	GetCode() *int32
	SetMessage(v string) *SendFeedbackResponseBody
	GetMessage() *string
	SetRequestId(v string) *SendFeedbackResponseBody
	GetRequestId() *string
}

type SendFeedbackResponseBody struct {
	// Interface status or POP error code. The values are as follows: 2xx: Success. 3xx: Redirect. 4xx: Request error. 5xx: Server error.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
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

func (s SendFeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SendFeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *SendFeedbackResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SendFeedbackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SendFeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendFeedbackResponseBody) SetCode(v int32) *SendFeedbackResponseBody {
	s.Code = &v
	return s
}

func (s *SendFeedbackResponseBody) SetMessage(v string) *SendFeedbackResponseBody {
	s.Message = &v
	return s
}

func (s *SendFeedbackResponseBody) SetRequestId(v string) *SendFeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendFeedbackResponseBody) Validate() error {
	return dara.Validate(s)
}
