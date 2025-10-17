// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIntents(v []*ListIntentResponseBodyIntents) *ListIntentResponseBody
	GetIntents() []*ListIntentResponseBodyIntents
	SetPageNumber(v int32) *ListIntentResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListIntentResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListIntentResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListIntentResponseBody
	GetTotalCount() *int32
}

type ListIntentResponseBody struct {
	Intents []*ListIntentResponseBodyIntents `json:"Intents,omitempty" xml:"Intents,omitempty" type:"Repeated"`
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
	// 23dsfa34r2s2s2sd12
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIntentResponseBody) GoString() string {
	return s.String()
}

func (s *ListIntentResponseBody) GetIntents() []*ListIntentResponseBodyIntents {
	return s.Intents
}

func (s *ListIntentResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListIntentResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIntentResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListIntentResponseBody) SetIntents(v []*ListIntentResponseBodyIntents) *ListIntentResponseBody {
	s.Intents = v
	return s
}

func (s *ListIntentResponseBody) SetPageNumber(v int32) *ListIntentResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListIntentResponseBody) SetPageSize(v int32) *ListIntentResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListIntentResponseBody) SetRequestId(v string) *ListIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIntentResponseBody) SetTotalCount(v int32) *ListIntentResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIntentResponseBody) Validate() error {
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

type ListIntentResponseBodyIntents struct {
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123231
	CreateUserId *string `json:"CreateUserId,omitempty" xml:"CreateUserId,omitempty"`
	// example:
	//
	// test
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	// example:
	//
	// 234234234234
	IntentId   *int64  `json:"IntentId,omitempty" xml:"IntentId,omitempty"`
	IntentName *string `json:"IntentName,omitempty" xml:"IntentName,omitempty"`
	// example:
	//
	// 2021-08-12T16:00:01Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 123231
	ModifyUserId *string `json:"ModifyUserId,omitempty" xml:"ModifyUserId,omitempty"`
	// example:
	//
	// test
	ModifyUserName *string                                   `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	SlotInfos      []*ListIntentResponseBodyIntentsSlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s ListIntentResponseBodyIntents) String() string {
	return dara.Prettify(s)
}

func (s ListIntentResponseBodyIntents) GoString() string {
	return s.String()
}

func (s *ListIntentResponseBodyIntents) GetAliasName() *string {
	return s.AliasName
}

func (s *ListIntentResponseBodyIntents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListIntentResponseBodyIntents) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListIntentResponseBodyIntents) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListIntentResponseBodyIntents) GetIntentId() *int64 {
	return s.IntentId
}

func (s *ListIntentResponseBodyIntents) GetIntentName() *string {
	return s.IntentName
}

func (s *ListIntentResponseBodyIntents) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListIntentResponseBodyIntents) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *ListIntentResponseBodyIntents) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *ListIntentResponseBodyIntents) GetSlotInfos() []*ListIntentResponseBodyIntentsSlotInfos {
	return s.SlotInfos
}

func (s *ListIntentResponseBodyIntents) SetAliasName(v string) *ListIntentResponseBodyIntents {
	s.AliasName = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetCreateTime(v string) *ListIntentResponseBodyIntents {
	s.CreateTime = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetCreateUserId(v string) *ListIntentResponseBodyIntents {
	s.CreateUserId = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetCreateUserName(v string) *ListIntentResponseBodyIntents {
	s.CreateUserName = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetIntentId(v int64) *ListIntentResponseBodyIntents {
	s.IntentId = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetIntentName(v string) *ListIntentResponseBodyIntents {
	s.IntentName = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetModifyTime(v string) *ListIntentResponseBodyIntents {
	s.ModifyTime = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetModifyUserId(v string) *ListIntentResponseBodyIntents {
	s.ModifyUserId = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetModifyUserName(v string) *ListIntentResponseBodyIntents {
	s.ModifyUserName = &v
	return s
}

func (s *ListIntentResponseBodyIntents) SetSlotInfos(v []*ListIntentResponseBodyIntentsSlotInfos) *ListIntentResponseBodyIntents {
	s.SlotInfos = v
	return s
}

func (s *ListIntentResponseBodyIntents) Validate() error {
	if s.SlotInfos != nil {
		for _, item := range s.SlotInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIntentResponseBodyIntentsSlotInfos struct {
	// example:
	//
	// false
	Array *bool `json:"Array,omitempty" xml:"Array,omitempty"`
	// example:
	//
	// false
	Encrypt *bool `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"`
	// example:
	//
	// false
	Interactive *bool   `json:"Interactive,omitempty" xml:"Interactive,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 12134223
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListIntentResponseBodyIntentsSlotInfos) String() string {
	return dara.Prettify(s)
}

func (s ListIntentResponseBodyIntentsSlotInfos) GoString() string {
	return s.String()
}

func (s *ListIntentResponseBodyIntentsSlotInfos) GetArray() *bool {
	return s.Array
}

func (s *ListIntentResponseBodyIntentsSlotInfos) GetEncrypt() *bool {
	return s.Encrypt
}

func (s *ListIntentResponseBodyIntentsSlotInfos) GetInteractive() *bool {
	return s.Interactive
}

func (s *ListIntentResponseBodyIntentsSlotInfos) GetName() *string {
	return s.Name
}

func (s *ListIntentResponseBodyIntentsSlotInfos) GetSlotId() *string {
	return s.SlotId
}

func (s *ListIntentResponseBodyIntentsSlotInfos) GetValue() *string {
	return s.Value
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetArray(v bool) *ListIntentResponseBodyIntentsSlotInfos {
	s.Array = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetEncrypt(v bool) *ListIntentResponseBodyIntentsSlotInfos {
	s.Encrypt = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetInteractive(v bool) *ListIntentResponseBodyIntentsSlotInfos {
	s.Interactive = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetName(v string) *ListIntentResponseBodyIntentsSlotInfos {
	s.Name = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetSlotId(v string) *ListIntentResponseBodyIntentsSlotInfos {
	s.SlotId = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) SetValue(v string) *ListIntentResponseBodyIntentsSlotInfos {
	s.Value = &v
	return s
}

func (s *ListIntentResponseBodyIntentsSlotInfos) Validate() error {
	return dara.Validate(s)
}
