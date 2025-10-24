// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomEntities(v *ListCustomEntitiesResponseBodyCustomEntities) *ListCustomEntitiesResponseBody
	GetCustomEntities() *ListCustomEntitiesResponseBodyCustomEntities
	SetPageNumber(v int32) *ListCustomEntitiesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCustomEntitiesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCustomEntitiesResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListCustomEntitiesResponseBody
	GetTotalCount() *int64
}

type ListCustomEntitiesResponseBody struct {
	CustomEntities *ListCustomEntitiesResponseBodyCustomEntities `json:"CustomEntities,omitempty" xml:"CustomEntities,omitempty" type:"Struct"`
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
	// 580e8ce3-3b80-44c5-9f3f-36ac3cc5bdd5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCustomEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCustomEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCustomEntitiesResponseBody) GetCustomEntities() *ListCustomEntitiesResponseBodyCustomEntities {
	return s.CustomEntities
}

func (s *ListCustomEntitiesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCustomEntitiesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCustomEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCustomEntitiesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCustomEntitiesResponseBody) SetCustomEntities(v *ListCustomEntitiesResponseBodyCustomEntities) *ListCustomEntitiesResponseBody {
	s.CustomEntities = v
	return s
}

func (s *ListCustomEntitiesResponseBody) SetPageNumber(v int32) *ListCustomEntitiesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCustomEntitiesResponseBody) SetPageSize(v int32) *ListCustomEntitiesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCustomEntitiesResponseBody) SetRequestId(v string) *ListCustomEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCustomEntitiesResponseBody) SetTotalCount(v int64) *ListCustomEntitiesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCustomEntitiesResponseBody) Validate() error {
	if s.CustomEntities != nil {
		if err := s.CustomEntities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCustomEntitiesResponseBodyCustomEntities struct {
	CustomEntity []*ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity `json:"CustomEntity,omitempty" xml:"CustomEntity,omitempty" type:"Repeated"`
}

func (s ListCustomEntitiesResponseBodyCustomEntities) String() string {
	return dara.Prettify(s)
}

func (s ListCustomEntitiesResponseBodyCustomEntities) GoString() string {
	return s.String()
}

func (s *ListCustomEntitiesResponseBodyCustomEntities) GetCustomEntity() []*ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity {
	return s.CustomEntity
}

func (s *ListCustomEntitiesResponseBodyCustomEntities) SetCustomEntity(v []*ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) *ListCustomEntitiesResponseBodyCustomEntities {
	s.CustomEntity = v
	return s
}

func (s *ListCustomEntitiesResponseBodyCustomEntities) Validate() error {
	if s.CustomEntity != nil {
		for _, item := range s.CustomEntity {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity struct {
	// example:
	//
	// 1
	CustomEntityId *string `json:"CustomEntityId,omitempty" xml:"CustomEntityId,omitempty"`
	// example:
	//
	// { "finegrainName":"example" }
	CustomEntityInfo *string `json:"CustomEntityInfo,omitempty" xml:"CustomEntityInfo,omitempty"`
	// example:
	//
	// exampleName
	CustomEntityName *string `json:"CustomEntityName,omitempty" xml:"CustomEntityName,omitempty"`
}

func (s ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) String() string {
	return dara.Prettify(s)
}

func (s ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) GoString() string {
	return s.String()
}

func (s *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) GetCustomEntityId() *string {
	return s.CustomEntityId
}

func (s *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) GetCustomEntityInfo() *string {
	return s.CustomEntityInfo
}

func (s *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) GetCustomEntityName() *string {
	return s.CustomEntityName
}

func (s *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) SetCustomEntityId(v string) *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity {
	s.CustomEntityId = &v
	return s
}

func (s *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) SetCustomEntityInfo(v string) *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity {
	s.CustomEntityInfo = &v
	return s
}

func (s *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) SetCustomEntityName(v string) *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity {
	s.CustomEntityName = &v
	return s
}

func (s *ListCustomEntitiesResponseBodyCustomEntitiesCustomEntity) Validate() error {
	return dara.Validate(s)
}
