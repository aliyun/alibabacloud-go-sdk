// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDSEntityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntities(v []*ListDSEntityResponseBodyEntities) *ListDSEntityResponseBody
	GetEntities() []*ListDSEntityResponseBodyEntities
	SetPageNumber(v int32) *ListDSEntityResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListDSEntityResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListDSEntityResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListDSEntityResponseBody
	GetTotalCount() *int32
}

type ListDSEntityResponseBody struct {
	Entities []*ListDSEntityResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
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
	// ga4h345defgwet2sdf223
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDSEntityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDSEntityResponseBody) GoString() string {
	return s.String()
}

func (s *ListDSEntityResponseBody) GetEntities() []*ListDSEntityResponseBodyEntities {
	return s.Entities
}

func (s *ListDSEntityResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDSEntityResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDSEntityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDSEntityResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDSEntityResponseBody) SetEntities(v []*ListDSEntityResponseBodyEntities) *ListDSEntityResponseBody {
	s.Entities = v
	return s
}

func (s *ListDSEntityResponseBody) SetPageNumber(v int32) *ListDSEntityResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDSEntityResponseBody) SetPageSize(v int32) *ListDSEntityResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListDSEntityResponseBody) SetRequestId(v string) *ListDSEntityResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDSEntityResponseBody) SetTotalCount(v int32) *ListDSEntityResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDSEntityResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDSEntityResponseBodyEntities struct {
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
	// 234564567445
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
	SysEntityCode  *string `json:"SysEntityCode,omitempty" xml:"SysEntityCode,omitempty"`
}

func (s ListDSEntityResponseBodyEntities) String() string {
	return dara.Prettify(s)
}

func (s ListDSEntityResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *ListDSEntityResponseBodyEntities) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDSEntityResponseBodyEntities) GetCreateUserId() *string {
	return s.CreateUserId
}

func (s *ListDSEntityResponseBodyEntities) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *ListDSEntityResponseBodyEntities) GetEntityId() *int64 {
	return s.EntityId
}

func (s *ListDSEntityResponseBodyEntities) GetEntityName() *string {
	return s.EntityName
}

func (s *ListDSEntityResponseBodyEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *ListDSEntityResponseBodyEntities) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListDSEntityResponseBodyEntities) GetModifyUserId() *string {
	return s.ModifyUserId
}

func (s *ListDSEntityResponseBodyEntities) GetModifyUserName() *string {
	return s.ModifyUserName
}

func (s *ListDSEntityResponseBodyEntities) GetSysEntityCode() *string {
	return s.SysEntityCode
}

func (s *ListDSEntityResponseBodyEntities) SetCreateTime(v string) *ListDSEntityResponseBodyEntities {
	s.CreateTime = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetCreateUserId(v string) *ListDSEntityResponseBodyEntities {
	s.CreateUserId = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetCreateUserName(v string) *ListDSEntityResponseBodyEntities {
	s.CreateUserName = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetEntityId(v int64) *ListDSEntityResponseBodyEntities {
	s.EntityId = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetEntityName(v string) *ListDSEntityResponseBodyEntities {
	s.EntityName = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetEntityType(v string) *ListDSEntityResponseBodyEntities {
	s.EntityType = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetModifyTime(v string) *ListDSEntityResponseBodyEntities {
	s.ModifyTime = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetModifyUserId(v string) *ListDSEntityResponseBodyEntities {
	s.ModifyUserId = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetModifyUserName(v string) *ListDSEntityResponseBodyEntities {
	s.ModifyUserName = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) SetSysEntityCode(v string) *ListDSEntityResponseBodyEntities {
	s.SysEntityCode = &v
	return s
}

func (s *ListDSEntityResponseBodyEntities) Validate() error {
	return dara.Validate(s)
}
