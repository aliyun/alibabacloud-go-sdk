// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationTicketsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetAddress(v string) *ListOperationTicketsRequest
	GetAssetAddress() *string
	SetInstanceId(v string) *ListOperationTicketsRequest
	GetInstanceId() *string
	SetPageNumber(v string) *ListOperationTicketsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListOperationTicketsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListOperationTicketsRequest
	GetRegionId() *string
}

type ListOperationTicketsRequest struct {
	// The IP address of the asset that is contained in the O\\&M application to be reviewed.
	//
	// example:
	//
	// 10.167.XX.XX
	AssetAddress *string `json:"AssetAddress,omitempty" xml:"AssetAddress,omitempty"`
	// The ID of the bastion host.
	//
	// >  You can call the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to query the ID of the bastion host.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page. Default value: **1**.
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
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOperationTicketsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOperationTicketsRequest) GoString() string {
	return s.String()
}

func (s *ListOperationTicketsRequest) GetAssetAddress() *string {
	return s.AssetAddress
}

func (s *ListOperationTicketsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListOperationTicketsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListOperationTicketsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListOperationTicketsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOperationTicketsRequest) SetAssetAddress(v string) *ListOperationTicketsRequest {
	s.AssetAddress = &v
	return s
}

func (s *ListOperationTicketsRequest) SetInstanceId(v string) *ListOperationTicketsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListOperationTicketsRequest) SetPageNumber(v string) *ListOperationTicketsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListOperationTicketsRequest) SetPageSize(v string) *ListOperationTicketsRequest {
	s.PageSize = &v
	return s
}

func (s *ListOperationTicketsRequest) SetRegionId(v string) *ListOperationTicketsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOperationTicketsRequest) Validate() error {
	return dara.Validate(s)
}
