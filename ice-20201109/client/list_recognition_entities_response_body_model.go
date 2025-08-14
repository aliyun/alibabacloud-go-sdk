// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntities(v *ListRecognitionEntitiesResponseBodyEntities) *ListRecognitionEntitiesResponseBody
	GetEntities() *ListRecognitionEntitiesResponseBodyEntities
	SetPageNumber(v int32) *ListRecognitionEntitiesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecognitionEntitiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRecognitionEntitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListRecognitionEntitiesResponseBody
	GetTotalCount() *int64
}

type ListRecognitionEntitiesResponseBody struct {
	Entities *ListRecognitionEntitiesResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Struct"`
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
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecognitionEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecognitionEntitiesResponseBody) GetEntities() *ListRecognitionEntitiesResponseBodyEntities {
	return s.Entities
}

func (s *ListRecognitionEntitiesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecognitionEntitiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecognitionEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecognitionEntitiesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRecognitionEntitiesResponseBody) SetEntities(v *ListRecognitionEntitiesResponseBodyEntities) *ListRecognitionEntitiesResponseBody {
	s.Entities = v
	return s
}

func (s *ListRecognitionEntitiesResponseBody) SetPageNumber(v int32) *ListRecognitionEntitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecognitionEntitiesResponseBody) SetPageSize(v int32) *ListRecognitionEntitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecognitionEntitiesResponseBody) SetRequestId(v string) *ListRecognitionEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecognitionEntitiesResponseBody) SetTotalCount(v int64) *ListRecognitionEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecognitionEntitiesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRecognitionEntitiesResponseBodyEntities struct {
	Entity []*ListRecognitionEntitiesResponseBodyEntitiesEntity `json:"Entity,omitempty" xml:"Entity,omitempty" type:"Repeated"`
}

func (s ListRecognitionEntitiesResponseBodyEntities) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionEntitiesResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *ListRecognitionEntitiesResponseBodyEntities) GetEntity() []*ListRecognitionEntitiesResponseBodyEntitiesEntity {
	return s.Entity
}

func (s *ListRecognitionEntitiesResponseBodyEntities) SetEntity(v []*ListRecognitionEntitiesResponseBodyEntitiesEntity) *ListRecognitionEntitiesResponseBodyEntities {
	s.Entity = v
	return s
}

func (s *ListRecognitionEntitiesResponseBodyEntities) Validate() error {
	return dara.Validate(s)
}

type ListRecognitionEntitiesResponseBodyEntitiesEntity struct {
	// example:
	//
	// **************544cb84754************
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// {}
	EntityInfo *string `json:"EntityInfo,omitempty" xml:"EntityInfo,omitempty"`
	EntityName *string `json:"EntityName,omitempty" xml:"EntityName,omitempty"`
}

func (s ListRecognitionEntitiesResponseBodyEntitiesEntity) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionEntitiesResponseBodyEntitiesEntity) GoString() string {
	return s.String()
}

func (s *ListRecognitionEntitiesResponseBodyEntitiesEntity) GetEntityId() *string {
	return s.EntityId
}

func (s *ListRecognitionEntitiesResponseBodyEntitiesEntity) GetEntityInfo() *string {
	return s.EntityInfo
}

func (s *ListRecognitionEntitiesResponseBodyEntitiesEntity) GetEntityName() *string {
	return s.EntityName
}

func (s *ListRecognitionEntitiesResponseBodyEntitiesEntity) SetEntityId(v string) *ListRecognitionEntitiesResponseBodyEntitiesEntity {
	s.EntityId = &v
	return s
}

func (s *ListRecognitionEntitiesResponseBodyEntitiesEntity) SetEntityInfo(v string) *ListRecognitionEntitiesResponseBodyEntitiesEntity {
	s.EntityInfo = &v
	return s
}

func (s *ListRecognitionEntitiesResponseBodyEntitiesEntity) SetEntityName(v string) *ListRecognitionEntitiesResponseBodyEntitiesEntity {
	s.EntityName = &v
	return s
}

func (s *ListRecognitionEntitiesResponseBodyEntitiesEntity) Validate() error {
	return dara.Validate(s)
}
