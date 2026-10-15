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
	// example:
	//
	// success
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TextTranslateAsyncResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
