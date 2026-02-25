// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListParameterSetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListParameterSetsResponseBody
	GetPageSize() *int32
	SetParameterSets(v []*ListParameterSetsResponseBodyParameterSets) *ListParameterSetsResponseBody
	GetParameterSets() []*ListParameterSetsResponseBodyParameterSets
	SetRequestId(v string) *ListParameterSetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListParameterSetsResponseBody
	GetTotalCount() *int32
}

type ListParameterSetsResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize      *int32                                        `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ParameterSets []*ListParameterSetsResponseBodyParameterSets `json:"parameterSets,omitempty" xml:"parameterSets,omitempty" type:"Repeated"`
	// example:
	//
	// 4E188A8C-D77A-53F2-9578-E9AD8ABF2FA9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 50
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListParameterSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListParameterSetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListParameterSetsResponseBody) GetParameterSets() []*ListParameterSetsResponseBodyParameterSets {
	return s.ParameterSets
}

func (s *ListParameterSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListParameterSetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListParameterSetsResponseBody) SetPageNumber(v int32) *ListParameterSetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListParameterSetsResponseBody) SetPageSize(v int32) *ListParameterSetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListParameterSetsResponseBody) SetParameterSets(v []*ListParameterSetsResponseBodyParameterSets) *ListParameterSetsResponseBody {
	s.ParameterSets = v
	return s
}

func (s *ListParameterSetsResponseBody) SetRequestId(v string) *ListParameterSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParameterSetsResponseBody) SetTotalCount(v int32) *ListParameterSetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListParameterSetsResponseBody) Validate() error {
	if s.ParameterSets != nil {
		for _, item := range s.ParameterSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListParameterSetsResponseBodyParameterSets struct {
	// example:
	//
	// 2022-05-14T10:05:19Z
	CreateTime         *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	DeletionProtection *bool   `json:"deletionProtection,omitempty" xml:"deletionProtection,omitempty"`
	// example:
	//
	// OK
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 12
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// pts-433aead756057ea135b21e89c
	ParameterSetId *string                                                   `json:"parameterSetId,omitempty" xml:"parameterSetId,omitempty"`
	Parameters     []*ListParameterSetsResponseBodyParameterSetsParameters   `json:"parameters,omitempty" xml:"parameters,omitempty" type:"Repeated"`
	RelationList   []*ListParameterSetsResponseBodyParameterSetsRelationList `json:"relationList,omitempty" xml:"relationList,omitempty" type:"Repeated"`
}

func (s ListParameterSetsResponseBodyParameterSets) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetsResponseBodyParameterSets) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponseBodyParameterSets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListParameterSetsResponseBodyParameterSets) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ListParameterSetsResponseBodyParameterSets) GetDescription() *string {
	return s.Description
}

func (s *ListParameterSetsResponseBodyParameterSets) GetName() *string {
	return s.Name
}

func (s *ListParameterSetsResponseBodyParameterSets) GetParameterSetId() *string {
	return s.ParameterSetId
}

func (s *ListParameterSetsResponseBodyParameterSets) GetParameters() []*ListParameterSetsResponseBodyParameterSetsParameters {
	return s.Parameters
}

func (s *ListParameterSetsResponseBodyParameterSets) GetRelationList() []*ListParameterSetsResponseBodyParameterSetsRelationList {
	return s.RelationList
}

func (s *ListParameterSetsResponseBodyParameterSets) SetCreateTime(v string) *ListParameterSetsResponseBodyParameterSets {
	s.CreateTime = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetDeletionProtection(v bool) *ListParameterSetsResponseBodyParameterSets {
	s.DeletionProtection = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetDescription(v string) *ListParameterSetsResponseBodyParameterSets {
	s.Description = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetName(v string) *ListParameterSetsResponseBodyParameterSets {
	s.Name = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetParameterSetId(v string) *ListParameterSetsResponseBodyParameterSets {
	s.ParameterSetId = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetParameters(v []*ListParameterSetsResponseBodyParameterSetsParameters) *ListParameterSetsResponseBodyParameterSets {
	s.Parameters = v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) SetRelationList(v []*ListParameterSetsResponseBodyParameterSetsRelationList) *ListParameterSetsResponseBodyParameterSets {
	s.RelationList = v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSets) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelationList != nil {
		for _, item := range s.RelationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListParameterSetsResponseBodyParameterSetsParameters struct {
	Name   *string `json:"name,omitempty" xml:"name,omitempty"`
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// string
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// 111
	Value interface{} `json:"value,omitempty" xml:"value,omitempty"`
}

func (s ListParameterSetsResponseBodyParameterSetsParameters) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetsResponseBodyParameterSetsParameters) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) GetName() *string {
	return s.Name
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) GetStatus() *string {
	return s.Status
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) GetType() *string {
	return s.Type
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) GetValue() interface{} {
	return s.Value
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) SetName(v string) *ListParameterSetsResponseBodyParameterSetsParameters {
	s.Name = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) SetStatus(v string) *ListParameterSetsResponseBodyParameterSetsParameters {
	s.Status = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) SetType(v string) *ListParameterSetsResponseBodyParameterSetsParameters {
	s.Type = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) SetValue(v interface{}) *ListParameterSetsResponseBodyParameterSetsParameters {
	s.Value = v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsParameters) Validate() error {
	return dara.Validate(s)
}

type ListParameterSetsResponseBodyParameterSetsRelationList struct {
	// example:
	//
	// 2022-06-09T03:46:18Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// task-433aead756057ffdf5326bf1e12ed
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// Module
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListParameterSetsResponseBodyParameterSetsRelationList) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetsResponseBodyParameterSetsRelationList) GoString() string {
	return s.String()
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) SetCreateTime(v string) *ListParameterSetsResponseBodyParameterSetsRelationList {
	s.CreateTime = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) SetResourceId(v string) *ListParameterSetsResponseBodyParameterSetsRelationList {
	s.ResourceId = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) SetResourceType(v string) *ListParameterSetsResponseBodyParameterSetsRelationList {
	s.ResourceType = &v
	return s
}

func (s *ListParameterSetsResponseBodyParameterSetsRelationList) Validate() error {
	return dara.Validate(s)
}
