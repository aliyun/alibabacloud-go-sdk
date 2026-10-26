// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHostGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostGroupName(v string) *ListHostGroupsRequest
	GetHostGroupName() *string
	SetInstanceId(v string) *ListHostGroupsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListHostGroupsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListHostGroupsRequest
	GetPageSize() *string
	SetProjectId(v int64) *ListHostGroupsRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListHostGroupsRequest
	GetRegionId() *string
}

type ListHostGroupsRequest struct {
	// The name of the asset group. Only exact match is supported.
	//
	// example:
	//
	// HostGroup1
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/462953.html) operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the current page for paging in a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page for paging in a paged query. Maximum value: 100. Default value: 20. If this parameter is left empty, 20 entries are returned by default.
	//
	// > Specify a value for this parameter.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 3
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The region ID of the bastion host.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListHostGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHostGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListHostGroupsRequest) GetHostGroupName() *string {
	return s.HostGroupName
}

func (s *ListHostGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListHostGroupsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListHostGroupsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListHostGroupsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListHostGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListHostGroupsRequest) SetHostGroupName(v string) *ListHostGroupsRequest {
	s.HostGroupName = &v
	return s
}

func (s *ListHostGroupsRequest) SetInstanceId(v string) *ListHostGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListHostGroupsRequest) SetPageNumber(v string) *ListHostGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHostGroupsRequest) SetPageSize(v string) *ListHostGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHostGroupsRequest) SetProjectId(v int64) *ListHostGroupsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListHostGroupsRequest) SetRegionId(v string) *ListHostGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsRequest) Validate() error {
	return dara.Validate(s)
}
