// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextTranslateAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TextTranslateAsyncResponseBody
	GetCode() *string
	SetData(v *TextTranslateAsyncResponseBodyData) *TextTranslateAsyncResponseBody
	GetData() *TextTranslateAsyncResponseBodyData
	SetMessage(v string) *TextTranslateAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *TextTranslateAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TextTranslateAsyncResponseBody
	GetSuccess() *bool
}

type TextTranslateAsyncResponseBody struct {
	// The response code of the submission call. Example: success.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submission result of the text translation asynchronous node.
	Data *TextTranslateAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The invocation message. Example of a successful call: Success. If the call fails, check the actual error message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the submission request. It is used for troubleshooting.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the submission call is successful. A successful call does not mean that the translation is complete.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s TextTranslateAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *TextTranslateAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *TextTranslateAsyncResponseBody) GetData() *TextTranslateAsyncResponseBodyData {
	return s.Data
}

func (s *TextTranslateAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TextTranslateAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TextTranslateAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TextTranslateAsyncResponseBody) SetCode(v string) *TextTranslateAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *TextTranslateAsyncResponseBody) SetData(v *TextTranslateAsyncResponseBodyData) *TextTranslateAsyncResponseBody {
	s.Data = v
	return s
}

func (s *TextTranslateAsyncResponseBody) SetMessage(v string) *TextTranslateAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *TextTranslateAsyncResponseBody) SetRequestId(v string) *TextTranslateAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *TextTranslateAsyncResponseBody) SetSuccess(v bool) *TextTranslateAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *TextTranslateAsyncResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TextTranslateAsyncResponseBodyData struct {
	// The ID of the asynchronous task. It is used for result queries and notification association. Do not use the example prefix as a business validation rule.
	//
	// example:
	//
	// TID_550e8400-e29b-41d4-a716-446655440000
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s TextTranslateAsyncResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TextTranslateAsyncResponseBodyData) GoString() string {
	return s.String()
}

func (s *TextTranslateAsyncResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *TextTranslateAsyncResponseBodyData) SetTaskId(v string) *TextTranslateAsyncResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *TextTranslateAsyncResponseBodyData) Validate() error {
	return dara.Validate(s)
}
