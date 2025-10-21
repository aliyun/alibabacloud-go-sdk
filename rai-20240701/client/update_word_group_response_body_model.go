// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWordGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateWordGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateWordGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateWordGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateWordGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateWordGroupResponseBody
	GetSuccess() *bool
	SetWordErrorInfoList(v []*UpdateWordGroupResponseBodyWordErrorInfoList) *UpdateWordGroupResponseBody
	GetWordErrorInfoList() []*UpdateWordGroupResponseBodyWordErrorInfoList
}

type UpdateWordGroupResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// If there is an error, return the error message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether it was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// List of error information.
	WordErrorInfoList []*UpdateWordGroupResponseBodyWordErrorInfoList `json:"WordErrorInfoList,omitempty" xml:"WordErrorInfoList,omitempty" type:"Repeated"`
}

func (s UpdateWordGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateWordGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateWordGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateWordGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateWordGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateWordGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateWordGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateWordGroupResponseBody) GetWordErrorInfoList() []*UpdateWordGroupResponseBodyWordErrorInfoList {
	return s.WordErrorInfoList
}

func (s *UpdateWordGroupResponseBody) SetCode(v string) *UpdateWordGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateWordGroupResponseBody) SetHttpStatusCode(v int32) *UpdateWordGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateWordGroupResponseBody) SetMessage(v string) *UpdateWordGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateWordGroupResponseBody) SetRequestId(v string) *UpdateWordGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateWordGroupResponseBody) SetSuccess(v bool) *UpdateWordGroupResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateWordGroupResponseBody) SetWordErrorInfoList(v []*UpdateWordGroupResponseBodyWordErrorInfoList) *UpdateWordGroupResponseBody {
	s.WordErrorInfoList = v
	return s
}

func (s *UpdateWordGroupResponseBody) Validate() error {
	if s.WordErrorInfoList != nil {
		for _, item := range s.WordErrorInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateWordGroupResponseBodyWordErrorInfoList struct {
	// Error message description.
	//
	// example:
	//
	// Keyword can not be empty
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Error status information.
	//
	// example:
	//
	// 1
	ErrorStatus *int32 `json:"ErrorStatus,omitempty" xml:"ErrorStatus,omitempty"`
	// Position of the error information in the array.
	//
	// example:
	//
	// 1
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
	// Label.
	//
	// example:
	//
	// Buss.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Keyword.
	//
	// example:
	//
	// Inv.
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s UpdateWordGroupResponseBodyWordErrorInfoList) String() string {
	return dara.Prettify(s)
}

func (s UpdateWordGroupResponseBodyWordErrorInfoList) GoString() string {
	return s.String()
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) GetErrorStatus() *int32 {
	return s.ErrorStatus
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) GetIndex() *int64 {
	return s.Index
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) GetLabel() *string {
	return s.Label
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) GetWord() *string {
	return s.Word
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) SetErrorMessage(v string) *UpdateWordGroupResponseBodyWordErrorInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) SetErrorStatus(v int32) *UpdateWordGroupResponseBodyWordErrorInfoList {
	s.ErrorStatus = &v
	return s
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) SetIndex(v int64) *UpdateWordGroupResponseBodyWordErrorInfoList {
	s.Index = &v
	return s
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) SetLabel(v string) *UpdateWordGroupResponseBodyWordErrorInfoList {
	s.Label = &v
	return s
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) SetWord(v string) *UpdateWordGroupResponseBodyWordErrorInfoList {
	s.Word = &v
	return s
}

func (s *UpdateWordGroupResponseBodyWordErrorInfoList) Validate() error {
	return dara.Validate(s)
}
