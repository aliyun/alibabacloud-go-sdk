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
	SetRegionId(v string) *ListHostGroupsRequest
	GetRegionId() *string
}

type ListHostGroupsRequest struct {
	// The name of the host group that you want to query. Only exact match is supported.
	//
	// example:
	//
	// Host group 1
	HostGroupName *string `json:"HostGroupName,omitempty" xml:"HostGroupName,omitempty"`
	// The ID of the bastion host to query.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/462953.html) operation to query the bastion host ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host in which you want to query the host group.
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

func (s *ListHostGroupsRequest) SetRegionId(v string) *ListHostGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *ListHostGroupsRequest) Validate() error {
	return dara.Validate(s)
}
