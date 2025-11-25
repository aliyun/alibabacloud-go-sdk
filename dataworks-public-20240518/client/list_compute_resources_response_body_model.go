// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListComputeResourcesResponseBodyPagingInfo) *ListComputeResourcesResponseBody
	GetPagingInfo() *ListComputeResourcesResponseBodyPagingInfo
	SetRequestId(v string) *ListComputeResourcesResponseBody
	GetRequestId() *string
}

type ListComputeResourcesResponseBody struct {
	PagingInfo *ListComputeResourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	RequestId  *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListComputeResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesResponseBody) GetPagingInfo() *ListComputeResourcesResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListComputeResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeResourcesResponseBody) SetPagingInfo(v *ListComputeResourcesResponseBodyPagingInfo) *ListComputeResourcesResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListComputeResourcesResponseBody) SetRequestId(v string) *ListComputeResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeResourcesResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeResourcesResponseBodyPagingInfo struct {
	ComputeResources []*ListComputeResourcesResponseBodyPagingInfoComputeResources `json:"ComputeResources,omitempty" xml:"ComputeResources,omitempty" type:"Repeated"`
	PageNumber       *int64                                                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize         *int64                                                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalCount       *int64                                                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComputeResourcesResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesResponseBodyPagingInfo) GetComputeResources() []*ListComputeResourcesResponseBodyPagingInfoComputeResources {
	return s.ComputeResources
}

func (s *ListComputeResourcesResponseBodyPagingInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListComputeResourcesResponseBodyPagingInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListComputeResourcesResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListComputeResourcesResponseBodyPagingInfo) SetComputeResources(v []*ListComputeResourcesResponseBodyPagingInfoComputeResources) *ListComputeResourcesResponseBodyPagingInfo {
	s.ComputeResources = v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfo) SetPageNumber(v int64) *ListComputeResourcesResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfo) SetPageSize(v int64) *ListComputeResourcesResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfo) SetTotalCount(v int64) *ListComputeResourcesResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfo) Validate() error {
	if s.ComputeResources != nil {
		for _, item := range s.ComputeResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeResourcesResponseBodyPagingInfoComputeResources struct {
	ComputeResource []*ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Repeated"`
	Name            *string                                                                      `json:"Name,omitempty" xml:"Name,omitempty"`
	Type            *string                                                                      `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListComputeResourcesResponseBodyPagingInfoComputeResources) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesResponseBodyPagingInfoComputeResources) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResources) GetComputeResource() []*ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	return s.ComputeResource
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResources) GetName() *string {
	return s.Name
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResources) GetType() *string {
	return s.Type
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResources) SetComputeResource(v []*ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) *ListComputeResourcesResponseBodyPagingInfoComputeResources {
	s.ComputeResource = v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResources) SetName(v string) *ListComputeResourcesResponseBodyPagingInfoComputeResources {
	s.Name = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResources) SetType(v string) *ListComputeResourcesResponseBodyPagingInfoComputeResources {
	s.Type = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResources) Validate() error {
	if s.ComputeResource != nil {
		for _, item := range s.ComputeResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource struct {
	ConnectionProperties     interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	ConnectionPropertiesMode *string     `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	CreateTime               *int64      `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateUser               *string     `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	Description              *string     `json:"Description,omitempty" xml:"Description,omitempty"`
	Id                       *int64      `json:"Id,omitempty" xml:"Id,omitempty"`
	ModifyTime               *int64      `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	ModifyUser               *string     `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	WhetherDefault           *bool       `json:"WhetherDefault,omitempty" xml:"WhetherDefault,omitempty"`
}

func (s ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) String() string {
	return dara.Prettify(s)
}

func (s ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GoString() string {
	return s.String()
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetConnectionProperties() interface{} {
	return s.ConnectionProperties
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetConnectionPropertiesMode() *string {
	return s.ConnectionPropertiesMode
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetDescription() *string {
	return s.Description
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetId() *int64 {
	return s.Id
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) GetWhetherDefault() *bool {
	return s.WhetherDefault
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetConnectionProperties(v interface{}) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.ConnectionProperties = v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetConnectionPropertiesMode(v string) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.ConnectionPropertiesMode = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetCreateTime(v int64) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.CreateTime = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetCreateUser(v string) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.CreateUser = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetDescription(v string) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.Description = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetId(v int64) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.Id = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetModifyTime(v int64) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.ModifyTime = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetModifyUser(v string) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.ModifyUser = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) SetWhetherDefault(v bool) *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource {
	s.WhetherDefault = &v
	return s
}

func (s *ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource) Validate() error {
	return dara.Validate(s)
}
