// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskAttachmentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DescribeDiskAttachmentsRequest
	GetDiskId() *string
	SetMaxResults(v int32) *DescribeDiskAttachmentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDiskAttachmentsRequest
	GetNextToken() *string
	SetOwnerId(v int64) *DescribeDiskAttachmentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeDiskAttachmentsRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DescribeDiskAttachmentsRequest
	GetResourceOwnerId() *int64
}

type DescribeDiskAttachmentsRequest struct {
	// The disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp131n0q38u3a4zi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The maximum number of attachments to return in a single request. Valid values: 10 to 500.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Obtain the value from the NextToken response parameter of the previous call. Do not specify this parameter in the first call.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerId   *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDiskAttachmentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskAttachmentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskAttachmentsRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskAttachmentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDiskAttachmentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiskAttachmentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDiskAttachmentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiskAttachmentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDiskAttachmentsRequest) SetDiskId(v string) *DescribeDiskAttachmentsRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskAttachmentsRequest) SetMaxResults(v int32) *DescribeDiskAttachmentsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiskAttachmentsRequest) SetNextToken(v string) *DescribeDiskAttachmentsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiskAttachmentsRequest) SetOwnerId(v int64) *DescribeDiskAttachmentsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDiskAttachmentsRequest) SetRegionId(v string) *DescribeDiskAttachmentsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiskAttachmentsRequest) SetResourceOwnerId(v int64) *DescribeDiskAttachmentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDiskAttachmentsRequest) Validate() error {
	return dara.Validate(s)
}
