// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApproveCommandsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListApproveCommandsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListApproveCommandsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListApproveCommandsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListApproveCommandsRequest
	GetRegionId() *string
}

type ListApproveCommandsRequest struct {
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-zvp2xvysf08
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page. Default value: **1**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.\\
	//
	// Maximum value: 1000. Default value: 20. If you leave this parameter empty, 20 entries are returned on each page.
	//
	// >  We recommend that you do not leave this parameter empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	//
	// >  For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListApproveCommandsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListApproveCommandsRequest) GoString() string {
	return s.String()
}

func (s *ListApproveCommandsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApproveCommandsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListApproveCommandsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListApproveCommandsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListApproveCommandsRequest) SetInstanceId(v string) *ListApproveCommandsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListApproveCommandsRequest) SetPageNumber(v string) *ListApproveCommandsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListApproveCommandsRequest) SetPageSize(v string) *ListApproveCommandsRequest {
	s.PageSize = &v
	return s
}

func (s *ListApproveCommandsRequest) SetRegionId(v string) *ListApproveCommandsRequest {
	s.RegionId = &v
	return s
}

func (s *ListApproveCommandsRequest) Validate() error {
	return dara.Validate(s)
}
