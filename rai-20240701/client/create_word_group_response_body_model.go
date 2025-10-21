// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWordGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWordGroupResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateWordGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateWordGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWordGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWordGroupResponseBody
	GetSuccess() *bool
	SetWordErrorInfoList(v []*CreateWordGroupResponseBodyWordErrorInfoList) *CreateWordGroupResponseBody
	GetWordErrorInfoList() []*CreateWordGroupResponseBodyWordErrorInfoList
}

type CreateWordGroupResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// If there is an error, returns the error message.
	//
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. true: Call succeeded. false: Call failed.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Error information list
	WordErrorInfoList []*CreateWordGroupResponseBodyWordErrorInfoList `json:"WordErrorInfoList,omitempty" xml:"WordErrorInfoList,omitempty" type:"Repeated"`
}

func (s CreateWordGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWordGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWordGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWordGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateWordGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWordGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWordGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWordGroupResponseBody) GetWordErrorInfoList() []*CreateWordGroupResponseBodyWordErrorInfoList {
	return s.WordErrorInfoList
}

func (s *CreateWordGroupResponseBody) SetCode(v string) *CreateWordGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWordGroupResponseBody) SetHttpStatusCode(v int32) *CreateWordGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWordGroupResponseBody) SetMessage(v string) *CreateWordGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWordGroupResponseBody) SetRequestId(v string) *CreateWordGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWordGroupResponseBody) SetSuccess(v bool) *CreateWordGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWordGroupResponseBody) SetWordErrorInfoList(v []*CreateWordGroupResponseBodyWordErrorInfoList) *CreateWordGroupResponseBody {
	s.WordErrorInfoList = v
	return s
}

func (s *CreateWordGroupResponseBody) Validate() error {
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

type CreateWordGroupResponseBodyWordErrorInfoList struct {
	// Error message description
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
	// Label
	//
	// example:
	//
	// Buss.
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// Keyword
	//
	// example:
	//
	// Inv.
	Word *string `json:"Word,omitempty" xml:"Word,omitempty"`
}

func (s CreateWordGroupResponseBodyWordErrorInfoList) String() string {
	return dara.Prettify(s)
}

func (s CreateWordGroupResponseBodyWordErrorInfoList) GoString() string {
	return s.String()
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) GetErrorStatus() *int32 {
	return s.ErrorStatus
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) GetIndex() *int64 {
	return s.Index
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) GetLabel() *string {
	return s.Label
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) GetWord() *string {
	return s.Word
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) SetErrorMessage(v string) *CreateWordGroupResponseBodyWordErrorInfoList {
	s.ErrorMessage = &v
	return s
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) SetErrorStatus(v int32) *CreateWordGroupResponseBodyWordErrorInfoList {
	s.ErrorStatus = &v
	return s
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) SetIndex(v int64) *CreateWordGroupResponseBodyWordErrorInfoList {
	s.Index = &v
	return s
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) SetLabel(v string) *CreateWordGroupResponseBodyWordErrorInfoList {
	s.Label = &v
	return s
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) SetWord(v string) *CreateWordGroupResponseBodyWordErrorInfoList {
	s.Word = &v
	return s
}

func (s *CreateWordGroupResponseBodyWordErrorInfoList) Validate() error {
	return dara.Validate(s)
}
