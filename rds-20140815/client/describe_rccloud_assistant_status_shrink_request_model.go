// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCCloudAssistantStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdsShrink(v string) *DescribeRCCloudAssistantStatusShrinkRequest
	GetInstanceIdsShrink() *string
	SetMaxResults(v int32) *DescribeRCCloudAssistantStatusShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRCCloudAssistantStatusShrinkRequest
	GetNextToken() *string
	SetOSType(v string) *DescribeRCCloudAssistantStatusShrinkRequest
	GetOSType() *string
	SetPageNumber(v int32) *DescribeRCCloudAssistantStatusShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCCloudAssistantStatusShrinkRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCCloudAssistantStatusShrinkRequest
	GetRegionId() *string
}

type DescribeRCCloudAssistantStatusShrinkRequest struct {
	InstanceIdsShrink *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRCCloudAssistantStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCCloudAssistantStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) GetInstanceIdsShrink() *string {
	return s.InstanceIdsShrink
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) GetOSType() *string {
	return s.OSType
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) SetInstanceIdsShrink(v string) *DescribeRCCloudAssistantStatusShrinkRequest {
	s.InstanceIdsShrink = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) SetMaxResults(v int32) *DescribeRCCloudAssistantStatusShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) SetNextToken(v string) *DescribeRCCloudAssistantStatusShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) SetOSType(v string) *DescribeRCCloudAssistantStatusShrinkRequest {
	s.OSType = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) SetPageNumber(v int32) *DescribeRCCloudAssistantStatusShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) SetPageSize(v int32) *DescribeRCCloudAssistantStatusShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) SetRegionId(v string) *DescribeRCCloudAssistantStatusShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}
