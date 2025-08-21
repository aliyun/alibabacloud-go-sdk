// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDSEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTime(v string) *DescribeDSEntityResponseBody
	GetCreateTime() *string
	SetCreateUserId(v string) *DescribeDSEntityResponseBody
	GetCreateUserId() *string
	SetCreateUserName(v string) *DescribeDSEntityResponseBody
	GetCreateUserName() *string
	SetEntityId(v int64) *DescribeDSEntityResponseBody
	GetEntityId() *int64
	SetEntityName(v string) *DescribeDSEntityResponseBody
	GetEntityName() *string
	SetEntityType(v string) *DescribeDSEntityResponseBody
	GetEntityType() *string
	SetModifyTime(v string) *DescribeDSEntityResponseBody
	GetModifyTime() *string
	SetModifyUserId(v string) *DescribeDSEntityResponseBody
	GetModifyUserId() *string
	SetModifyUserName(v string) *DescribeDSEntityResponseBody
	GetModifyUserName() *string
	SetRequestId(v string) *DescribeDSEntityResponseBody
	GetRequestId() *string
	SetSysEntityCode(v string) *DescribeDSEntityResponseBody
	GetSysEntityCode() *string
}

type DescribeDSEntityResponseBody struct {
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
	// 123
	EntityId   *int64  `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
	// example:
	//
	// synonyms
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
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
	// ad23234dsf234fga
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SysEntityCode *string `json:"SysEntityCode,omitempty" xml:"SysEntityCode,omitempty"`
}

func (s DescribeDSEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDSEntityResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDSEntityResponseBody) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *DescribeDSEntityResponseBody) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *DescribeDSEntityResponseBody) GetEntityId() *int64 {
	return s.EntityId
}

func (s *DescribeDSEntityResponseBody) GetEntityName() *string {
	return s.EntityName
}

func (s *DescribeDSEntityResponseBody) GetEntityType() *string {
	return s.EntityType
}

func (s *DescribeDSEntityResponseBody) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeDSEntityResponseBody) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *DescribeDSEntityResponseBody) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *DescribeDSEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDSEntityResponseBody) GetSysEntityCode() *string {
	return s.SysEntityCode
}

func (s *DescribeDSEntityResponseBody) SetCreateTime(v string) *DescribeDSEntityResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetCreateUserId(v string) *DescribeDSEntityResponseBody {
	s.CreateUserId = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetCreateUserName(v string) *DescribeDSEntityResponseBody {
	s.CreateUserName = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetEntityId(v int64) *DescribeDSEntityResponseBody {
	s.EntityId = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetEntityName(v string) *DescribeDSEntityResponseBody {
	s.EntityName = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetEntityType(v string) *DescribeDSEntityResponseBody {
	s.EntityType = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetModifyTime(v string) *DescribeDSEntityResponseBody {
	s.ModifyTime = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetModifyUserId(v string) *DescribeDSEntityResponseBody {
	s.ModifyUserId = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetModifyUserName(v string) *DescribeDSEntityResponseBody {
	s.ModifyUserName = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetRequestId(v string) *DescribeDSEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDSEntityResponseBody) SetSysEntityCode(v string) *DescribeDSEntityResponseBody {
	s.SysEntityCode = &v
	return s
}

func (s *DescribeDSEntityResponseBody) Validate() error {
	return dara.Validate(s)
}
