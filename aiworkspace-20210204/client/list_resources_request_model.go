// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *ListResourcesRequest
	GetGroupName() *string
	SetLabels(v string) *ListResourcesRequest
	GetLabels() *string
	SetOption(v string) *ListResourcesRequest
	GetOption() *string
	SetPageNumber(v int64) *ListResourcesRequest
	GetPageNumber() *int64
	SetPageSize(v int32) *ListResourcesRequest
	GetPageSize() *int32
	SetProductTypes(v string) *ListResourcesRequest
	GetProductTypes() *string
	SetQuotaIds(v string) *ListResourcesRequest
	GetQuotaIds() *string
	SetResourceName(v string) *ListResourcesRequest
	GetResourceName() *string
	SetResourceTypes(v string) *ListResourcesRequest
	GetResourceTypes() *string
	SetVerbose(v bool) *ListResourcesRequest
	GetVerbose() *bool
	SetVerboseFields(v string) *ListResourcesRequest
	GetVerboseFields() *string
	SetWorkspaceId(v string) *ListResourcesRequest
	GetWorkspaceId() *string
}

type ListResourcesRequest struct {
	// The name of the resource group. You can call [ListResources](https://help.aliyun.com/document_detail/449143.html) to obtain the name of the resource group.
	//
	// example:
	//
	// group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Tag-based filter conditions. Multiple conditions are separated by commas (,). Only resources that meet all the specified tag-based filter conditions are returned.
	//
	// This parameter is available only for resources whose ProductType is ACS.
	//
	// example:
	//
	// system.supported.dsw=true,system.supported.dlc=true
	Labels *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// The operation to perform. Valid values:
	//
	// 	- ListResourceByWorkspace: obtains the resources in the workspace. This is the default value.
	//
	// 	- ListResource: obtains the resources of the user.
	//
	// example:
	//
	// ListResourceByWorkspace
	Option *string `json:"Option,omitempty" xml:"Option,omitempty"`
	// The page number. The pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// **This field is no longer used and will be removed. Use the ResourceType field instead.
	//
	// example:
	//
	// MaxCompute
	ProductTypes *string `json:"ProductTypes,omitempty" xml:"ProductTypes,omitempty"`
	// The quota IDs, which are separated by commas (,). Only resources that contain all the specified quotas are returned.
	//
	// >  This parameter is available only for resources whose ResourceTypes is ACS.
	//
	// example:
	//
	// quota-k******da,quota-cd******w
	QuotaIds *string `json:"QuotaIds,omitempty" xml:"QuotaIds,omitempty"`
	// The resource name. The value must meet the following requirements:
	//
	// 	- The name must be 3 to 28 characters in length.
	//
	// 	- The name is unique in the region.
	//
	// example:
	//
	// resource
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// The resource types. Valid values:
	//
	// 	- MaxCompute
	//
	// 	- ECS
	//
	// 	- Lingjun
	//
	// 	- ACS
	//
	// 	- FLINK
	//
	// example:
	//
	// MaxCompute
	ResourceTypes *string `json:"ResourceTypes,omitempty" xml:"ResourceTypes,omitempty"`
	// Specifies whether to show detailed information, which includes the Quotas field. Valid values:
	//
	// 	- true (default)
	//
	// 	- false
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// The fields to return. Multiple fields are separated by commas (,). Valid values:
	//
	// 	- Quota
	//
	// 	- Label
	//
	// 	- IsDefault
	//
	// example:
	//
	// Quota,IsDefault
	VerboseFields *string `json:"VerboseFields,omitempty" xml:"VerboseFields,omitempty"`
	// The workspace ID. You can call [ListWorkspaces](https://help.aliyun.com/document_detail/449124.html) to obtain the workspace ID.
	//
	// 	- This parameter is required when the Option parameter is set to ListResourceByWorkspace.
	//
	// 	- You do not need to configure this parameter when the Option parameter is set to ListResource.
	//
	// example:
	//
	// 123
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListResourcesRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListResourcesRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListResourcesRequest) GetOption() *string {
	return s.Option
}

func (s *ListResourcesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListResourcesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourcesRequest) GetProductTypes() *string {
	return s.ProductTypes
}

func (s *ListResourcesRequest) GetQuotaIds() *string {
	return s.QuotaIds
}

func (s *ListResourcesRequest) GetResourceName() *string {
	return s.ResourceName
}

func (s *ListResourcesRequest) GetResourceTypes() *string {
	return s.ResourceTypes
}

func (s *ListResourcesRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *ListResourcesRequest) GetVerboseFields() *string {
	return s.VerboseFields
}

func (s *ListResourcesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListResourcesRequest) SetGroupName(v string) *ListResourcesRequest {
	s.GroupName = &v
	return s
}

func (s *ListResourcesRequest) SetLabels(v string) *ListResourcesRequest {
	s.Labels = &v
	return s
}

func (s *ListResourcesRequest) SetOption(v string) *ListResourcesRequest {
	s.Option = &v
	return s
}

func (s *ListResourcesRequest) SetPageNumber(v int64) *ListResourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListResourcesRequest) SetPageSize(v int32) *ListResourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListResourcesRequest) SetProductTypes(v string) *ListResourcesRequest {
	s.ProductTypes = &v
	return s
}

func (s *ListResourcesRequest) SetQuotaIds(v string) *ListResourcesRequest {
	s.QuotaIds = &v
	return s
}

func (s *ListResourcesRequest) SetResourceName(v string) *ListResourcesRequest {
	s.ResourceName = &v
	return s
}

func (s *ListResourcesRequest) SetResourceTypes(v string) *ListResourcesRequest {
	s.ResourceTypes = &v
	return s
}

func (s *ListResourcesRequest) SetVerbose(v bool) *ListResourcesRequest {
	s.Verbose = &v
	return s
}

func (s *ListResourcesRequest) SetVerboseFields(v string) *ListResourcesRequest {
	s.VerboseFields = &v
	return s
}

func (s *ListResourcesRequest) SetWorkspaceId(v string) *ListResourcesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListResourcesRequest) Validate() error {
	return dara.Validate(s)
}
