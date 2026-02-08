// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeIntentResponseBody
	GetHttpStatusCode() *int32
	SetIntent(v *DescribeIntentResponseBodyIntent) *DescribeIntentResponseBody
	GetIntent() *DescribeIntentResponseBodyIntent
	SetMessage(v string) *DescribeIntentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeIntentResponseBody
	GetSuccess() *bool
}

type DescribeIntentResponseBody struct {
	// Response code
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Intent information
	Intent *DescribeIntentResponseBodyIntent `json:"Intent,omitempty" xml:"Intent,omitempty" type:"Struct"`
	// API response message
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 327FEE69-F173-5B2F-9F3B-DCC6182D7BA0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeIntentResponseBody) GetIntent() *DescribeIntentResponseBodyIntent {
	return s.Intent
}

func (s *DescribeIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeIntentResponseBody) SetCode(v string) *DescribeIntentResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeIntentResponseBody) SetHttpStatusCode(v int32) *DescribeIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeIntentResponseBody) SetIntent(v *DescribeIntentResponseBodyIntent) *DescribeIntentResponseBody {
	s.Intent = v
	return s
}

func (s *DescribeIntentResponseBody) SetMessage(v string) *DescribeIntentResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeIntentResponseBody) SetRequestId(v string) *DescribeIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetSuccess(v bool) *DescribeIntentResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeIntentResponseBody) Validate() error {
	if s.Intent != nil {
		if err := s.Intent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeIntentResponseBodyIntent struct {
	// Creation Time
	//
	// example:
	//
	// 1578469042851
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Instance Description
	//
	// example:
	//
	// 同意还款意图
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// Intent ID
	//
	// example:
	//
	// 10722701
	IntentId *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	// Intent Name
	//
	// example:
	//
	// 统一还款
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// Keywords of the intent, which can be used to filter and find intents in the List API
	//
	// example:
	//
	// ["还款"]
	Keywords *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// Scenario ID
	//
	// example:
	//
	// 6ef95fd5-558f-4ee8-af34-b2ede087a87c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// Update Time
	//
	// example:
	//
	// 1578469042851
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// List of user utterances that trigger the intent
	//
	// example:
	//
	// ["ok","好的","好吧","好嘞","可以","行啊","行吧","那行","知道了","我看一下","能的","等会吧","等一下","马上还","等一会","过两天","我会想办法处理"]
	Utterances *string `json:"Utterances,omitempty" xml:"Utterances,omitempty"`
}

func (s DescribeIntentResponseBodyIntent) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentResponseBodyIntent) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBodyIntent) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeIntentResponseBodyIntent) GetIntentDescription() *string {
	return s.IntentDescription
}

func (s *DescribeIntentResponseBodyIntent) GetIntentId() *string {
	return s.IntentId
}

func (s *DescribeIntentResponseBodyIntent) GetIntentName() *string {
	return s.IntentName
}

func (s *DescribeIntentResponseBodyIntent) GetKeywords() *string {
	return s.Keywords
}

func (s *DescribeIntentResponseBodyIntent) GetScriptId() *string {
	return s.ScriptId
}

func (s *DescribeIntentResponseBodyIntent) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeIntentResponseBodyIntent) GetUtterances() *string {
	return s.Utterances
}

func (s *DescribeIntentResponseBodyIntent) SetCreateTime(v int64) *DescribeIntentResponseBodyIntent {
	s.CreateTime = &v
	return s
}

func (s *DescribeIntentResponseBodyIntent) SetIntentDescription(v string) *DescribeIntentResponseBodyIntent {
	s.IntentDescription = &v
	return s
}

func (s *DescribeIntentResponseBodyIntent) SetIntentId(v string) *DescribeIntentResponseBodyIntent {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentResponseBodyIntent) SetIntentName(v string) *DescribeIntentResponseBodyIntent {
	s.IntentName = &v
	return s
}

func (s *DescribeIntentResponseBodyIntent) SetKeywords(v string) *DescribeIntentResponseBodyIntent {
	s.Keywords = &v
	return s
}

func (s *DescribeIntentResponseBodyIntent) SetScriptId(v string) *DescribeIntentResponseBodyIntent {
	s.ScriptId = &v
	return s
}

func (s *DescribeIntentResponseBodyIntent) SetUpdateTime(v int64) *DescribeIntentResponseBodyIntent {
	s.UpdateTime = &v
	return s
}

func (s *DescribeIntentResponseBodyIntent) SetUtterances(v string) *DescribeIntentResponseBodyIntent {
	s.Utterances = &v
	return s
}

func (s *DescribeIntentResponseBodyIntent) Validate() error {
	return dara.Validate(s)
}
