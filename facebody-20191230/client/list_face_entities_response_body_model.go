// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaceEntitiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListFaceEntitiesResponseBodyData) *ListFaceEntitiesResponseBody
	GetData() *ListFaceEntitiesResponseBodyData
	SetRequestId(v string) *ListFaceEntitiesResponseBody
	GetRequestId() *string
}

type ListFaceEntitiesResponseBody struct {
	Data *ListFaceEntitiesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// DA7CAFEB-0A37-4496-9CDF-E3DBB6309CB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFaceEntitiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFaceEntitiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBody) GetData() *ListFaceEntitiesResponseBodyData {
	return s.Data
}

func (s *ListFaceEntitiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFaceEntitiesResponseBody) SetData(v *ListFaceEntitiesResponseBodyData) *ListFaceEntitiesResponseBody {
	s.Data = v
	return s
}

func (s *ListFaceEntitiesResponseBody) SetRequestId(v string) *ListFaceEntitiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaceEntitiesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListFaceEntitiesResponseBodyData struct {
	Entities []*ListFaceEntitiesResponseBodyDataEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// example:
	//
	// 200
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListFaceEntitiesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListFaceEntitiesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBodyData) GetEntities() []*ListFaceEntitiesResponseBodyDataEntities {
	return s.Entities
}

func (s *ListFaceEntitiesResponseBodyData) GetToken() *string {
	return s.Token
}

func (s *ListFaceEntitiesResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListFaceEntitiesResponseBodyData) SetEntities(v []*ListFaceEntitiesResponseBodyDataEntities) *ListFaceEntitiesResponseBodyData {
	s.Entities = v
	return s
}

func (s *ListFaceEntitiesResponseBodyData) SetToken(v string) *ListFaceEntitiesResponseBodyData {
	s.Token = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyData) SetTotalCount(v int32) *ListFaceEntitiesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyData) Validate() error {
	if s.Entities != nil {
		for _, item := range s.Entities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFaceEntitiesResponseBodyDataEntities struct {
	// example:
	//
	// 1589203871832
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// 222
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// example:
	//
	// 3
	FaceCount *int32  `json:"FaceCount,omitempty" xml:"FaceCount,omitempty"`
	Labels    *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 1589203871832
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s ListFaceEntitiesResponseBodyDataEntities) String() string {
	return dara.Prettify(s)
}

func (s ListFaceEntitiesResponseBodyDataEntities) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesResponseBodyDataEntities) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *ListFaceEntitiesResponseBodyDataEntities) GetDbName() *string {
	return s.DbName
}

func (s *ListFaceEntitiesResponseBodyDataEntities) GetEntityId() *string {
	return s.EntityId
}

func (s *ListFaceEntitiesResponseBodyDataEntities) GetFaceCount() *int32 {
	return s.FaceCount
}

func (s *ListFaceEntitiesResponseBodyDataEntities) GetLabels() *string {
	return s.Labels
}

func (s *ListFaceEntitiesResponseBodyDataEntities) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetCreatedAt(v int64) *ListFaceEntitiesResponseBodyDataEntities {
	s.CreatedAt = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetDbName(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.DbName = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetEntityId(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.EntityId = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetFaceCount(v int32) *ListFaceEntitiesResponseBodyDataEntities {
	s.FaceCount = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetLabels(v string) *ListFaceEntitiesResponseBodyDataEntities {
	s.Labels = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) SetUpdatedAt(v int64) *ListFaceEntitiesResponseBodyDataEntities {
	s.UpdatedAt = &v
	return s
}

func (s *ListFaceEntitiesResponseBodyDataEntities) Validate() error {
	return dara.Validate(s)
}
