// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWordGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetWordGroupResponseBody
	GetCode() *string
	SetGroupName(v string) *GetWordGroupResponseBody
	GetGroupName() *string
	SetHttpStatusCode(v int32) *GetWordGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetWordGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetWordGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetWordGroupResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *GetWordGroupResponseBody
	GetTotalCount() *int32
	SetWordInfoList(v []*GetWordGroupResponseBodyWordInfoList) *GetWordGroupResponseBody
	GetWordInfoList() []*GetWordGroupResponseBodyWordInfoList
}

type GetWordGroupResponseBody struct {
	// Status code, 00000 indicates success; others indicate failure.
	//
	// example:
	//
	// 00000
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Keyword group name.
	//
	// example:
	//
	// testGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// HTTP status code.
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
	// Request ID.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total count.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Keyword group list.
	WordInfoList []*GetWordGroupResponseBodyWordInfoList `json:"WordInfoList,omitempty" xml:"WordInfoList,omitempty" type:"Repeated"`
}

func (s GetWordGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetWordGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetWordGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetWordGroupResponseBody) GetGroupName() *string {
	return s.GroupName
}

func (s *GetWordGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetWordGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetWordGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetWordGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetWordGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetWordGroupResponseBody) GetWordInfoList() []*GetWordGroupResponseBodyWordInfoList {
	return s.WordInfoList
}

func (s *GetWordGroupResponseBody) SetCode(v string) *GetWordGroupResponseBody {
	s.Code = &v
	return s
}

func (s *GetWordGroupResponseBody) SetGroupName(v string) *GetWordGroupResponseBody {
	s.GroupName = &v
	return s
}

func (s *GetWordGroupResponseBody) SetHttpStatusCode(v int32) *GetWordGroupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetWordGroupResponseBody) SetMessage(v string) *GetWordGroupResponseBody {
	s.Message = &v
	return s
}

func (s *GetWordGroupResponseBody) SetRequestId(v string) *GetWordGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetWordGroupResponseBody) SetSuccess(v bool) *GetWordGroupResponseBody {
	s.Success = &v
	return s
}

func (s *GetWordGroupResponseBody) SetTotalCount(v int32) *GetWordGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetWordGroupResponseBody) SetWordInfoList(v []*GetWordGroupResponseBodyWordInfoList) *GetWordGroupResponseBody {
	s.WordInfoList = v
	return s
}

func (s *GetWordGroupResponseBody) Validate() error {
	if s.WordInfoList != nil {
		for _, item := range s.WordInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetWordGroupResponseBodyWordInfoList struct {
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
	// ID of the successfully added word.
	//
	// example:
	//
	// 1
	WordId *int64 `json:"WordId,omitempty" xml:"WordId,omitempty"`
}

func (s GetWordGroupResponseBodyWordInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetWordGroupResponseBodyWordInfoList) GoString() string {
	return s.String()
}

func (s *GetWordGroupResponseBodyWordInfoList) GetLabel() *string {
	return s.Label
}

func (s *GetWordGroupResponseBodyWordInfoList) GetWord() *string {
	return s.Word
}

func (s *GetWordGroupResponseBodyWordInfoList) GetWordId() *int64 {
	return s.WordId
}

func (s *GetWordGroupResponseBodyWordInfoList) SetLabel(v string) *GetWordGroupResponseBodyWordInfoList {
	s.Label = &v
	return s
}

func (s *GetWordGroupResponseBodyWordInfoList) SetWord(v string) *GetWordGroupResponseBodyWordInfoList {
	s.Word = &v
	return s
}

func (s *GetWordGroupResponseBodyWordInfoList) SetWordId(v int64) *GetWordGroupResponseBodyWordInfoList {
	s.WordId = &v
	return s
}

func (s *GetWordGroupResponseBodyWordInfoList) Validate() error {
	return dara.Validate(s)
}
