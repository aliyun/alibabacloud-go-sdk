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
	// The exact asset address to search for in the pending approval list.
	//
	// example:
	//
	// 10.167.XX.XX
	AssetAddress *string `json:"AssetAddress,omitempty" xml:"AssetAddress,omitempty"`
	// The instance ID of the bastion host.
	//
	// > You can invoke [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number for a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries per page for a paged query.
	//
	// The maximum value of PageSize is 1000. The default number of entries per page is 20. If PageSize is left empty, 20 entries are returned by default.
	//
	// > Do not leave PageSize empty.
	//
	// example:
	//
	// 20
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the bastion host.
	//
	// > For the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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
