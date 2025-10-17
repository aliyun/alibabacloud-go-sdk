// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *DescribeIntentResponseBody
	GetAliasName() *string
	SetCreateTime(v string) *DescribeIntentResponseBody
	GetCreateTime() *string
	SetCreateUserId(v string) *DescribeIntentResponseBody
	GetCreateUserId() *string
	SetCreateUserName(v string) *DescribeIntentResponseBody
	GetCreateUserName() *string
	SetIntentId(v int64) *DescribeIntentResponseBody
	GetIntentId() *int64
	SetIntentName(v string) *DescribeIntentResponseBody
	GetIntentName() *string
	SetModifyTime(v string) *DescribeIntentResponseBody
	GetModifyTime() *string
	SetModifyUserId(v string) *DescribeIntentResponseBody
	GetModifyUserId() *string
	SetModifyUserName(v string) *DescribeIntentResponseBody
	GetModifyUserName() *string
	SetRequestId(v string) *DescribeIntentResponseBody
	GetRequestId() *string
	SetSlotInfos(v []*DescribeIntentResponseBodySlotInfos) *DescribeIntentResponseBody
	GetSlotInfos() []*DescribeIntentResponseBodySlotInfos
}

type DescribeIntentResponseBody struct {
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
	// 84243341
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
	ModifyUserName *string `json:"ModifyUserName,omitempty" xml:"ModifyUserName,omitempty"`
	// example:
	//
	// a22afaf2adfasf2gr345fga45ada
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SlotInfos []*DescribeIntentResponseBodySlotInfos `json:"SlotInfos,omitempty" xml:"SlotInfos,omitempty" type:"Repeated"`
}

func (s DescribeIntentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBody) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeIntentResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeIntentResponseBody) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *DescribeIntentResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeIntentResponseBody) GetIntentId() *int64 {
	return s.IntentId
}

func (s *DescribeIntentResponseBody) GetIntentName() *string {
	return s.IntentName
}

func (s *DescribeIntentResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeIntentResponseBody) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *DescribeIntentResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeIntentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeIntentResponseBody) GetSlotInfos() []*DescribeIntentResponseBodySlotInfos {
	return s.SlotInfos
}

func (s *DescribeIntentResponseBody) SetAliasName(v string) *DescribeIntentResponseBody {
	s.AliasName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetCreateTime(v string) *DescribeIntentResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeIntentResponseBody) SetCreateUserId(v string) *DescribeIntentResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetCreateUserName(v string) *DescribeIntentResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetIntentId(v int64) *DescribeIntentResponseBody {
	s.IntentId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetIntentName(v string) *DescribeIntentResponseBody {
	s.IntentName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyTime(v string) *DescribeIntentResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyUserId(v string) *DescribeIntentResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetModifyUserName(v string) *DescribeIntentResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeIntentResponseBody) SetRequestId(v string) *DescribeIntentResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeIntentResponseBody) SetSlotInfos(v []*DescribeIntentResponseBodySlotInfos) *DescribeIntentResponseBody {
	s.SlotInfos = v
	return s
}

func (s *DescribeIntentResponseBody) Validate() error {
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

type DescribeIntentResponseBodySlotInfos struct {
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
	// aa4d2a343a3ad4afad
	SlotId *string `json:"SlotId,omitempty" xml:"SlotId,omitempty"`
	Value  *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeIntentResponseBodySlotInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentResponseBodySlotInfos) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponseBodySlotInfos) GetArray() *bool {
	return s.Array
}

func (s *DescribeIntentResponseBodySlotInfos) GetEncrypt() *bool {
	return s.Encrypt
}

func (s *DescribeIntentResponseBodySlotInfos) GetInteractive() *bool {
	return s.Interactive
}

func (s *DescribeIntentResponseBodySlotInfos) GetName() *string {
	return s.Name
}

func (s *DescribeIntentResponseBodySlotInfos) GetSlotId() *string {
	return s.SlotId
}

func (s *DescribeIntentResponseBodySlotInfos) GetValue() *string {
	return s.Value
}

func (s *DescribeIntentResponseBodySlotInfos) SetArray(v bool) *DescribeIntentResponseBodySlotInfos {
	s.Array = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetEncrypt(v bool) *DescribeIntentResponseBodySlotInfos {
	s.Encrypt = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetInteractive(v bool) *DescribeIntentResponseBodySlotInfos {
	s.Interactive = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetName(v string) *DescribeIntentResponseBodySlotInfos {
	s.Name = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetSlotId(v string) *DescribeIntentResponseBodySlotInfos {
	s.SlotId = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) SetValue(v string) *DescribeIntentResponseBodySlotInfos {
	s.Value = &v
	return s
}

func (s *DescribeIntentResponseBodySlotInfos) Validate() error {
	return dara.Validate(s)
}
