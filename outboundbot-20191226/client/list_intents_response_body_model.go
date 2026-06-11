// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListIntentsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListIntentsResponseBody
	GetHttpStatusCode() *int32
	SetIntents(v *ListIntentsResponseBodyIntents) *ListIntentsResponseBody
	GetIntents() *ListIntentsResponseBodyIntents
	SetMessage(v string) *ListIntentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListIntentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListIntentsResponseBody
	GetSuccess() *bool
}

type ListIntentsResponseBody struct {
	// The status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The paginated list of intents.
	Intents *ListIntentsResponseBodyIntents `json:"Intents,omitempty" xml:"Intents,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListIntentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListIntentsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListIntentsResponseBody) GetIntents() *ListIntentsResponseBodyIntents {
	return s.Intents
}

func (s *ListIntentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListIntentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListIntentsResponseBody) SetCode(v string) *ListIntentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListIntentsResponseBody) SetHttpStatusCode(v int32) *ListIntentsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListIntentsResponseBody) SetIntents(v *ListIntentsResponseBodyIntents) *ListIntentsResponseBody {
	s.Intents = v
	return s
}

func (s *ListIntentsResponseBody) SetMessage(v string) *ListIntentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListIntentsResponseBody) SetRequestId(v string) *ListIntentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntentsResponseBody) SetSuccess(v bool) *ListIntentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListIntentsResponseBody) Validate() error {
	if s.Intents != nil {
		if err := s.Intents.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListIntentsResponseBodyIntents struct {
	// The list of intents.
	List []*ListIntentsResponseBodyIntentsList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of intents.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIntentsResponseBodyIntents) String() string {
	return dara.Prettify(s)
}

func (s ListIntentsResponseBodyIntents) GoString() string {
	return s.String()
}

func (s *ListIntentsResponseBodyIntents) GetList() []*ListIntentsResponseBodyIntentsList {
	return s.List
}

func (s *ListIntentsResponseBodyIntents) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIntentsResponseBodyIntents) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntentsResponseBodyIntents) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIntentsResponseBodyIntents) SetList(v []*ListIntentsResponseBodyIntentsList) *ListIntentsResponseBodyIntents {
	s.List = v
	return s
}

func (s *ListIntentsResponseBodyIntents) SetPageNumber(v int32) *ListIntentsResponseBodyIntents {
	s.PageNumber = &v
	return s
}

func (s *ListIntentsResponseBodyIntents) SetPageSize(v int32) *ListIntentsResponseBodyIntents {
	s.PageSize = &v
	return s
}

func (s *ListIntentsResponseBodyIntents) SetTotalCount(v int32) *ListIntentsResponseBodyIntents {
	s.TotalCount = &v
	return s
}

func (s *ListIntentsResponseBodyIntents) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntentsResponseBodyIntentsList struct {
	// The time when the intent was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1578469042851
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the intent.
	//
	// example:
	//
	// 同意还款意图
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// The intent ID.
	//
	// example:
	//
	// a8494b35-eefb-4c8a-887b-b60d2f0fa57a
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// The intent name.
	//
	// example:
	//
	// 统一还款
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// The keywords of the intent.
	//
	// example:
	//
	// ["还款"]
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// The script ID.
	//
	// example:
	//
	// 6ef95fd5-558f-4ee8-af34-b2ede087a87c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// The time when the intent was last updated. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1578469042851
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The user utterances that trigger the intent.
	//
	// example:
	//
	// ["ok","好的","好吧","好嘞","可以","行啊","行吧","那行","知道了","我看一下","能的","等会吧","等一下","马上还","等一会","过两天","我会想办法处理"]
	Utterances *string `json:"Utterances,omitempty" xml:"Utterances,omitempty"`
}

func (s ListIntentsResponseBodyIntentsList) String() string {
	return dara.Prettify(s)
}

func (s ListIntentsResponseBodyIntentsList) GoString() string {
	return s.String()
}

func (s *ListIntentsResponseBodyIntentsList) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListIntentsResponseBodyIntentsList) GetIntentDescription() *string {
	return s.IntentDescription
}

func (s *ListIntentsResponseBodyIntentsList) GetIntentId() *string {
	return s.IntentId
}

func (s *ListIntentsResponseBodyIntentsList) GetIntentName() *string {
	return s.IntentName
}

func (s *ListIntentsResponseBodyIntentsList) GetKeywords() *string {
	return s.Keywords
}

func (s *ListIntentsResponseBodyIntentsList) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListIntentsResponseBodyIntentsList) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListIntentsResponseBodyIntentsList) GetUtterances() *string {
	return s.Utterances
}

func (s *ListIntentsResponseBodyIntentsList) SetCreateTime(v int64) *ListIntentsResponseBodyIntentsList {
	s.CreateTime = &v
	return s
}

func (s *ListIntentsResponseBodyIntentsList) SetIntentDescription(v string) *ListIntentsResponseBodyIntentsList {
	s.IntentDescription = &v
	return s
}

func (s *ListIntentsResponseBodyIntentsList) SetIntentId(v string) *ListIntentsResponseBodyIntentsList {
	s.IntentId = &v
	return s
}

func (s *ListIntentsResponseBodyIntentsList) SetIntentName(v string) *ListIntentsResponseBodyIntentsList {
	s.IntentName = &v
	return s
}

func (s *ListIntentsResponseBodyIntentsList) SetKeywords(v string) *ListIntentsResponseBodyIntentsList {
	s.Keywords = &v
	return s
}

func (s *ListIntentsResponseBodyIntentsList) SetScriptId(v string) *ListIntentsResponseBodyIntentsList {
	s.ScriptId = &v
	return s
}

func (s *ListIntentsResponseBodyIntentsList) SetUpdateTime(v int64) *ListIntentsResponseBodyIntentsList {
	s.UpdateTime = &v
	return s
}

func (s *ListIntentsResponseBodyIntentsList) SetUtterances(v string) *ListIntentsResponseBodyIntentsList {
	s.Utterances = &v
	return s
}

func (s *ListIntentsResponseBodyIntentsList) Validate() error {
	return dara.Validate(s)
}
