// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFeedbackDialogueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FeedbackDialogueResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *FeedbackDialogueResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *FeedbackDialogueResponseBody
	GetMessage() *string
	SetRequestId(v string) *FeedbackDialogueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FeedbackDialogueResponseBody
	GetSuccess() *bool
}

type FeedbackDialogueResponseBody struct {
	// The status code.
	//
	// example:
	//
	// successful
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 3f7045e099474ba28ceca1b4eb6d6e21
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of \\`true\\` indicates success. A value of \\`false\\` indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FeedbackDialogueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FeedbackDialogueResponseBody) GoString() string {
	return s.String()
}

func (s *FeedbackDialogueResponseBody) GetCode() *string {
	return s.Code
}

func (s *FeedbackDialogueResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *FeedbackDialogueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FeedbackDialogueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FeedbackDialogueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FeedbackDialogueResponseBody) SetCode(v string) *FeedbackDialogueResponseBody {
	s.Code = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetHttpStatusCode(v int32) *FeedbackDialogueResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetMessage(v string) *FeedbackDialogueResponseBody {
	s.Message = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetRequestId(v string) *FeedbackDialogueResponseBody {
	s.RequestId = &v
	return s
}

func (s *FeedbackDialogueResponseBody) SetSuccess(v bool) *FeedbackDialogueResponseBody {
	s.Success = &v
	return s
}

func (s *FeedbackDialogueResponseBody) Validate() error {
	return dara.Validate(s)
}
