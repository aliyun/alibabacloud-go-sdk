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
	SetProjectId(v int64) *ListOperationTicketsRequest
	GetProjectId() *int64
	SetRegionId(v string) *ListOperationTicketsRequest
	GetRegionId() *string
}

type ListOperationTicketsRequest struct {
	// The address of the O&M asset for exact search.
	//
	// example:
	//
	// 10.167.XX.XX
	AssetAddress *string `json:"AssetAddress,omitempty" xml:"AssetAddress,omitempty"`
	// The instance ID of the Bastionhost instance.
	//
	// > You can invoke the [DescribeInstances](https://help.aliyun.com/document_detail/153281.html) operation to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// bastionhost-cn-st220aw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the page to return in a paged query. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of entries to return on each page in a paged query.
	//
	// The maximum value of the PageSize parameter is 1000. The default value is 20. If this parameter is left empty, 20 entries are returned by default.
	//
	// > We recommend that you specify a value for the PageSize parameter.
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
	// The region ID of the Bastionhost instance.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
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

func (s *ListOperationTicketsRequest) GetProjectId() *int64 {
	return s.ProjectId
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

func (s *ListOperationTicketsRequest) SetProjectId(v int64) *ListOperationTicketsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListOperationTicketsRequest) SetRegionId(v string) *ListOperationTicketsRequest {
	s.RegionId = &v
	return s
}

func (s *ListOperationTicketsRequest) Validate() error {
	return dara.Validate(s)
}
