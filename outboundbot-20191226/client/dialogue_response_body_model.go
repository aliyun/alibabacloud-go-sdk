// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DialogueResponseBody
	GetCode() *string
	SetFeedback(v *DialogueResponseBodyFeedback) *DialogueResponseBody
	GetFeedback() *DialogueResponseBodyFeedback
	SetHttpStatusCode(v int32) *DialogueResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DialogueResponseBody
	GetMessage() *string
	SetRequestId(v string) *DialogueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DialogueResponseBody
	GetSuccess() *bool
}

type DialogueResponseBody struct {
	// example:
	//
	// OK
	Code     *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Feedback *DialogueResponseBodyFeedback `json:"Feedback,omitempty" xml:"Feedback,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBody) GoString() string {
	return s.String()
}

func (s *DialogueResponseBody) GetCode() *string {
	return s.Code
}

func (s *DialogueResponseBody) GetFeedback() *DialogueResponseBodyFeedback {
	return s.Feedback
}

func (s *DialogueResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DialogueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DialogueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DialogueResponseBody) SetCode(v string) *DialogueResponseBody {
	s.Code = &v
	return s
}

func (s *DialogueResponseBody) SetFeedback(v *DialogueResponseBodyFeedback) *DialogueResponseBody {
	s.Feedback = v
	return s
}

func (s *DialogueResponseBody) SetHttpStatusCode(v int32) *DialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DialogueResponseBody) SetMessage(v string) *DialogueResponseBody {
	s.Message = &v
	return s
}

func (s *DialogueResponseBody) SetRequestId(v string) *DialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *DialogueResponseBody) SetSuccess(v bool) *DialogueResponseBody {
	s.Success = &v
	return s
}

func (s *DialogueResponseBody) Validate() error {
	if s.Feedback != nil {
		if err := s.Feedback.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DialogueResponseBodyFeedback struct {
	// example:
	//
	// broadcast
	Action       *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ActionParams *string `json:"ActionParams,omitempty" xml:"ActionParams,omitempty"`
	Content      *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 已废弃
	//
	// example:
	//
	// “”
	ContentParams *string `json:"ContentParams,omitempty" xml:"ContentParams,omitempty"`
	// example:
	//
	// true
	Interruptible *bool `json:"Interruptible,omitempty" xml:"Interruptible,omitempty"`
}

func (s DialogueResponseBodyFeedback) String() string {
	return dara.Prettify(s)
}

func (s DialogueResponseBodyFeedback) GoString() string {
	return s.String()
}

func (s *DialogueResponseBodyFeedback) GetAction() *string {
	return s.Action
}

func (s *DialogueResponseBodyFeedback) GetActionParams() *string {
	return s.ActionParams
}

func (s *DialogueResponseBodyFeedback) GetContent() *string {
	return s.Content
}

func (s *DialogueResponseBodyFeedback) GetContentParams() *string {
	return s.ContentParams
}

func (s *DialogueResponseBodyFeedback) GetInterruptible() *bool {
	return s.Interruptible
}

func (s *DialogueResponseBodyFeedback) SetAction(v string) *DialogueResponseBodyFeedback {
	s.Action = &v
	return s
}

func (s *DialogueResponseBodyFeedback) SetActionParams(v string) *DialogueResponseBodyFeedback {
	s.ActionParams = &v
	return s
}

func (s *DialogueResponseBodyFeedback) SetContent(v string) *DialogueResponseBodyFeedback {
	s.Content = &v
	return s
}

func (s *DialogueResponseBodyFeedback) SetContentParams(v string) *DialogueResponseBodyFeedback {
	s.ContentParams = &v
	return s
}

func (s *DialogueResponseBodyFeedback) SetInterruptible(v bool) *DialogueResponseBodyFeedback {
	s.Interruptible = &v
	return s
}

func (s *DialogueResponseBodyFeedback) Validate() error {
	return dara.Validate(s)
}
