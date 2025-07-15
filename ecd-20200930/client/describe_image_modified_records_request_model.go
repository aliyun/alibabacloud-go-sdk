// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageModifiedRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribeImageModifiedRecordsRequest
	GetDesktopId() *string
	SetMaxResults(v int32) *DescribeImageModifiedRecordsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeImageModifiedRecordsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeImageModifiedRecordsRequest
	GetRegionId() *string
}

type DescribeImageModifiedRecordsRequest struct {
	// The ID of the cloud computer.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecd-bd53sfmysz8ir****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// 	- Maximum value: 100.
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that determines the start point of the next query. If this parameter is left empty, all results are returned.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6l3d+SWeOobbIlDLjwhjkTk
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeImageModifiedRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageModifiedRecordsRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageModifiedRecordsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeImageModifiedRecordsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImageModifiedRecordsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImageModifiedRecordsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageModifiedRecordsRequest) SetDesktopId(v string) *DescribeImageModifiedRecordsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeImageModifiedRecordsRequest) SetMaxResults(v int32) *DescribeImageModifiedRecordsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImageModifiedRecordsRequest) SetNextToken(v string) *DescribeImageModifiedRecordsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImageModifiedRecordsRequest) SetRegionId(v string) *DescribeImageModifiedRecordsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageModifiedRecordsRequest) Validate() error {
	return dara.Validate(s)
}
