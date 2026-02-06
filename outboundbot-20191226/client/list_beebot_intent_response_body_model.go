// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBeebotIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBeebotRequestId(v string) *ListBeebotIntentResponseBody
	GetBeebotRequestId() *string
	SetCode(v string) *ListBeebotIntentResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListBeebotIntentResponseBody
	GetHttpStatusCode() *int32
	SetIntents(v []*ListBeebotIntentResponseBodyIntents) *ListBeebotIntentResponseBody
	GetIntents() []*ListBeebotIntentResponseBodyIntents
	SetMessage(v string) *ListBeebotIntentResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListBeebotIntentResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBeebotIntentResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListBeebotIntentResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBeebotIntentResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListBeebotIntentResponseBody
	GetTotalCount() *int32
}

type ListBeebotIntentResponseBody struct {
	// example:
	//
	// 497CFAFF-48CC-161A-AD2C-252DED569037
	BeebotRequestId *string `json:"BeebotRequestId,omitempty" xml:"BeebotRequestId,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32                                 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Intents        []*ListBeebotIntentResponseBodyIntents `json:"Intents,omitempty" xml:"Intents,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBeebotIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentResponseBody) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentResponseBody) GetBeebotRequestId() *string {
	return s.BeebotRequestId
}

func (s *ListBeebotIntentResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBeebotIntentResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBeebotIntentResponseBody) GetIntents() []*ListBeebotIntentResponseBodyIntents {
	return s.Intents
}

func (s *ListBeebotIntentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBeebotIntentResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBeebotIntentResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBeebotIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBeebotIntentResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBeebotIntentResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBeebotIntentResponseBody) SetBeebotRequestId(v string) *ListBeebotIntentResponseBody {
	s.BeebotRequestId = &v
	return s
}

func (s *ListBeebotIntentResponseBody) SetCode(v string) *ListBeebotIntentResponseBody {
	s.Code = &v
	return s
}

func (s *ListBeebotIntentResponseBody) SetHttpStatusCode(v int32) *ListBeebotIntentResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBeebotIntentResponseBody) SetIntents(v []*ListBeebotIntentResponseBodyIntents) *ListBeebotIntentResponseBody {
	s.Intents = v
	return s
}

func (s *ListBeebotIntentResponseBody) SetMessage(v string) *ListBeebotIntentResponseBody {
	s.Message = &v
	return s
}

func (s *ListBeebotIntentResponseBody) SetPageNumber(v int32) *ListBeebotIntentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListBeebotIntentResponseBody) SetPageSize(v int32) *ListBeebotIntentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListBeebotIntentResponseBody) SetRequestId(v string) *ListBeebotIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBeebotIntentResponseBody) SetSuccess(v bool) *ListBeebotIntentResponseBody {
	s.Success = &v
	return s
}

func (s *ListBeebotIntentResponseBody) SetTotalCount(v int32) *ListBeebotIntentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListBeebotIntentResponseBody) Validate() error {
	if s.Intents != nil {
		for _, item := range s.Intents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBeebotIntentResponseBodyIntents struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 2025-04-21 16:03:15.+0800
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
	// 2025-04-21 16:03:15.+0800
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

func (s ListBeebotIntentResponseBodyIntents) String() string {
	return dara.Prettify(s)
}

func (s ListBeebotIntentResponseBodyIntents) GoString() string {
	return s.String()
}

func (s *ListBeebotIntentResponseBodyIntents) GetAliasName() *string {
	return s.AliasName
}

func (s *ListBeebotIntentResponseBodyIntents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListBeebotIntentResponseBodyIntents) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListBeebotIntentResponseBodyIntents) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListBeebotIntentResponseBodyIntents) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListBeebotIntentResponseBodyIntents) GetIntentName() *string {
	return s.IntentName
}

func (s *ListBeebotIntentResponseBodyIntents) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListBeebotIntentResponseBodyIntents) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *ListBeebotIntentResponseBodyIntents) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *ListBeebotIntentResponseBodyIntents) SetAliasName(v string) *ListBeebotIntentResponseBodyIntents {
	s.AliasName = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) SetCreateTime(v string) *ListBeebotIntentResponseBodyIntents {
	s.CreateTime = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) SetCreateUserId(v string) *ListBeebotIntentResponseBodyIntents {
	s.CreateUserId = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) SetCreateUserName(v string) *ListBeebotIntentResponseBodyIntents {
	s.CreateUserName = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) SetIntentId(v int64) *ListBeebotIntentResponseBodyIntents {
	s.IntentId = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) SetIntentName(v string) *ListBeebotIntentResponseBodyIntents {
	s.IntentName = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) SetModifyTime(v string) *ListBeebotIntentResponseBodyIntents {
	s.ModifyTime = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) SetModifyUserId(v string) *ListBeebotIntentResponseBodyIntents {
	s.ModifyUserId = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) SetModifyUserName(v string) *ListBeebotIntentResponseBodyIntents {
	s.ModifyUserName = &v
	return s
}

func (s *ListBeebotIntentResponseBodyIntents) Validate() error {
	return dara.Validate(s)
}
