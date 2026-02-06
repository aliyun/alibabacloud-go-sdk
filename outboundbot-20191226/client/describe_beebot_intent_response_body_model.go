// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBeebotIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *DescribeBeebotIntentResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *DescribeBeebotIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeBeebotIntentResponseBody
	GetHttpStatusCode() *int32
	SetIntent(v *DescribeBeebotIntentResponseBodyIntent) *DescribeBeebotIntentResponseBody
	GetIntent() *DescribeBeebotIntentResponseBodyIntent
	SetIntentId(v int64) *DescribeBeebotIntentResponseBody
	GetIntentId() *int64
	SetMessage(v string) *DescribeBeebotIntentResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeBeebotIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeBeebotIntentResponseBody
	GetSuccess() *bool
}

type DescribeBeebotIntentResponseBody struct {
	// example:
	//
	// 0B219FCB-EC71-1F08-BB1B-0E87C20158C8
	BeebotRequestId *string `json:"BeebotRequestId,omitempty" xml:"BeebotRequestId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Intent         *DescribeBeebotIntentResponseBodyIntent `json:"Intent,omitempty" xml:"Intent,omitempty" type:"Struct"`
	// example:
	//
	// 10717802
	IntentId *int64 `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
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

func (s DescribeBeebotIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBeebotIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBeebotIntentResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *DescribeBeebotIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeBeebotIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeBeebotIntentResponseBody) GetIntent() *DescribeBeebotIntentResponseBodyIntent {
	return s.Intent
}

func (s *DescribeBeebotIntentResponseBody) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DescribeBeebotIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeBeebotIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBeebotIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeBeebotIntentResponseBody) SetBeebotRequestId(v string) *DescribeBeebotIntentResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *DescribeBeebotIntentResponseBody) SetCode(v string) *DescribeBeebotIntentResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeBeebotIntentResponseBody) SetHttpStatusCode(v int32) *DescribeBeebotIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeBeebotIntentResponseBody) SetIntent(v *DescribeBeebotIntentResponseBodyIntent) *DescribeBeebotIntentResponseBody {
	s.Intent = v
	return s
}

func (s *DescribeBeebotIntentResponseBody) SetIntentId(v int64) *DescribeBeebotIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *DescribeBeebotIntentResponseBody) SetMessage(v string) *DescribeBeebotIntentResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeBeebotIntentResponseBody) SetRequestId(v string) *DescribeBeebotIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBeebotIntentResponseBody) SetSuccess(v bool) *DescribeBeebotIntentResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeBeebotIntentResponseBody) Validate() error {
	if s.Intent != nil {
		if err := s.Intent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeBeebotIntentResponseBodyIntent struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 2025-04-21 10:29:59.+0800
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1252504
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// xxx@voice-navigator-testonaliyun.com
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 10717802
	IntentId   *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// 2025-04-21 15:52:57.+0800
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 1252504
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// xxx@voice-navigator-testonaliyun.com
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
}

func (s DescribeBeebotIntentResponseBodyIntent) String() string {
	return dara.Prettify(s)
}

func (s DescribeBeebotIntentResponseBodyIntent) GoString() string {
	return s.String()
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetIntentName() *string {
	return s.IntentName
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *DescribeBeebotIntentResponseBodyIntent) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetAliasName(v string) *DescribeBeebotIntentResponseBodyIntent {
	s.AliasName = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetCreateTime(v string) *DescribeBeebotIntentResponseBodyIntent {
	s.CreateTime = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetCreateUserId(v string) *DescribeBeebotIntentResponseBodyIntent {
	s.CreateUserId = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetCreateUserName(v string) *DescribeBeebotIntentResponseBodyIntent {
	s.CreateUserName = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetIntentId(v int64) *DescribeBeebotIntentResponseBodyIntent {
	s.IntentId = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetIntentName(v string) *DescribeBeebotIntentResponseBodyIntent {
	s.IntentName = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetModifyTime(v string) *DescribeBeebotIntentResponseBodyIntent {
	s.ModifyTime = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetModifyUserId(v string) *DescribeBeebotIntentResponseBodyIntent {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) SetModifyUserName(v string) *DescribeBeebotIntentResponseBodyIntent {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeBeebotIntentResponseBodyIntent) Validate() error {
	return dara.Validate(s)
}
