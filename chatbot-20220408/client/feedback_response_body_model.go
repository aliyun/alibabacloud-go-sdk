// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeedback(v string) *FeedbackResponseBody
	GetFeedback() *string
	SetMessageId(v string) *FeedbackResponseBody
	GetMessageId() *string
	SetRequestId(v string) *FeedbackResponseBody
	GetRequestId() *string
}

type FeedbackResponseBody struct {
	// example:
	//
	// good
	Feedback *string `json:"Feedback,omitempty" xml:"Feedback,omitempty"`
	// example:
	//
	// 5ca40988-4f99-47ad-ac96-9060d0f81db9
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
	// example:
	//
	// 4e5eea71-f326-450c-8849-49515473ef64
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s FeedbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FeedbackResponseBody) GoString() string {
	return s.String()
}

func (s *FeedbackResponseBody) GetFeedback() *string {
	return s.Feedback
}

func (s *FeedbackResponseBody) GetMessageId() *string {
	return s.MessageId
}

func (s *FeedbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FeedbackResponseBody) SetFeedback(v string) *FeedbackResponseBody {
	s.Feedback = &v
	return s
}

func (s *FeedbackResponseBody) SetMessageId(v string) *FeedbackResponseBody {
	s.MessageId = &v
	return s
}

func (s *FeedbackResponseBody) SetRequestId(v string) *FeedbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *FeedbackResponseBody) Validate() error {
	return dara.Validate(s)
}
