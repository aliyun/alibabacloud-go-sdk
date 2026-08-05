// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLakebaseS3AccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListLakebaseS3AccountsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLakebaseS3AccountsRequest
	GetPageSize() *int32
	SetPfsInstanceId(v string) *ListLakebaseS3AccountsRequest
	GetPfsInstanceId() *string
	SetRegionId(v string) *ListLakebaseS3AccountsRequest
	GetRegionId() *string
}

type ListLakebaseS3AccountsRequest struct {
	// The page number.
	//
	// example:
	//
	// 5
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The PolarFS instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pfs-xxx
	PfsInstanceId *string `json:"PfsInstanceId,omitempty" xml:"PfsInstanceId,omitempty"`
	// The region ID.
	//
	// >You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query region IDs.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListLakebaseS3AccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLakebaseS3AccountsRequest) GoString() string {
	return s.String()
}

func (s *ListLakebaseS3AccountsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLakebaseS3AccountsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLakebaseS3AccountsRequest) GetPfsInstanceId() *string {
	return s.PfsInstanceId
}

func (s *ListLakebaseS3AccountsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLakebaseS3AccountsRequest) SetPageNumber(v int32) *ListLakebaseS3AccountsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListLakebaseS3AccountsRequest) SetPageSize(v int32) *ListLakebaseS3AccountsRequest {
	s.PageSize = &v
	return s
}

func (s *ListLakebaseS3AccountsRequest) SetPfsInstanceId(v string) *ListLakebaseS3AccountsRequest {
	s.PfsInstanceId = &v
	return s
}

func (s *ListLakebaseS3AccountsRequest) SetRegionId(v string) *ListLakebaseS3AccountsRequest {
	s.RegionId = &v
	return s
}

func (s *ListLakebaseS3AccountsRequest) Validate() error {
	return dara.Validate(s)
}
