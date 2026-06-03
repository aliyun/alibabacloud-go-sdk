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
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                            `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Intent         *DescribeIntentResponseBodyIntent `json:"Intent,omitempty" xml:"Intent,omitempty" type:"Struct"`
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
	// example:
	//
	// 1578469042851
	CreateTime        *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	IntentDescription *string `json:"IntentDescription,omitempty" xml:"IntentDescription,omitempty"`
	// example:
	//
	// a8494b35-eefb-4c8a-887b-b60d2f0fa57a
	IntentId   *string `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	Keywords   *string `json:"Keywords,omitempty" xml:"Keywords,omitempty"`
	// example:
	//
	// 6ef95fd5-558f-4ee8-af34-b2ede087a87c
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
	// example:
	//
	// 1578469042851
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
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
