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
	// Pagination information.
	PagingInfo *ListComputeResourcesResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. Used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 7C352CB7-CD88-50CF-9D0D-E81BDF02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The list of computing resources. Each element is a computing resource group that contains information about the development environment (if any) and the production environment.
	ComputeResources []*ListComputeResourcesResponseBodyPagingInfoComputeResources `json:"ComputeResources,omitempty" xml:"ComputeResources,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	// Details of a single computing resource.
	ComputeResource []*ListComputeResourcesResponseBodyPagingInfoComputeResourcesComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Repeated"`
	// The name of the computing resource.
	//
	// example:
	//
	// cal_all_vemg_workflow_parallel
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the computing resource.
	//
	// example:
	//
	// hologres
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
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
	// The category of the added compute resource. Different types have different subtypes with corresponding parameter constraints. Examples: InstanceMode: The instance mode. UrlMode: The connection string mode.
	//
	// example:
	//
	// {\\n    \\"clusterIdentifier\\": \\"c-da123456\\",\\n    \\"database\\": \\"testdb\\",\\n    \\"loginMode\\":\\"Anonymous\\",\\n    \\"defaultFS\\":\\"127.0.0.1\\",\\n    \\"envType\\": \\"Prod\\"\\n}
	ConnectionProperties interface{} `json:"ConnectionProperties,omitempty" xml:"ConnectionProperties,omitempty"`
	// The specific connection configuration details for the computing resource, including the connection address, access identity, and environment information. envType, which specifies the computing resource environment, is a property of this object. Valid values:
	//
	// 	- Dev
	//
	// 	- Prod Different types of computing resources have different attribute specifications under different configuration modes (ConnectionPropertiesMode).
	//
	// example:
	//
	// UrlMode
	ConnectionPropertiesMode *string `json:"ConnectionPropertiesMode,omitempty" xml:"ConnectionPropertiesMode,omitempty"`
	// The creation time (timestamp).
	//
	// example:
	//
	// 1624387842781448
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator ID.
	//
	// example:
	//
	// 1648711113000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The unique identifier of the computing resource.
	//
	// example:
	//
	// home_feed
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The computing resource ID, which is the unique identifier for the resource.
	//
	// example:
	//
	// 8649832502405358603
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modified time of the computing resource (timestamp).
	//
	// example:
	//
	// 1624387842781448
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The modifier ID.
	//
	// example:
	//
	// 1648711113000
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// Specifies whether it is the default compute resource.
	//
	// example:
	//
	// true
	WhetherDefault *bool `json:"WhetherDefault,omitempty" xml:"WhetherDefault,omitempty"`
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
