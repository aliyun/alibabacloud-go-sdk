// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCCloudAssistantStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *DescribeRCCloudAssistantStatusRequest
	GetInstanceIds() []*string
	SetMaxResults(v int32) *DescribeRCCloudAssistantStatusRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeRCCloudAssistantStatusRequest
	GetNextToken() *string
	SetOSType(v string) *DescribeRCCloudAssistantStatusRequest
	GetOSType() *string
	SetPageNumber(v int32) *DescribeRCCloudAssistantStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRCCloudAssistantStatusRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRCCloudAssistantStatusRequest
	GetRegionId() *string
}

type DescribeRCCloudAssistantStatusRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
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

func (s DescribeRCCloudAssistantStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCCloudAssistantStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeRCCloudAssistantStatusRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeRCCloudAssistantStatusRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRCCloudAssistantStatusRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRCCloudAssistantStatusRequest) GetOSType() *string {
	return s.OSType
}

func (s *DescribeRCCloudAssistantStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRCCloudAssistantStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRCCloudAssistantStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRCCloudAssistantStatusRequest) SetInstanceIds(v []*string) *DescribeRCCloudAssistantStatusRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeRCCloudAssistantStatusRequest) SetMaxResults(v int32) *DescribeRCCloudAssistantStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusRequest) SetNextToken(v string) *DescribeRCCloudAssistantStatusRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusRequest) SetOSType(v string) *DescribeRCCloudAssistantStatusRequest {
	s.OSType = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusRequest) SetPageNumber(v int32) *DescribeRCCloudAssistantStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusRequest) SetPageSize(v int32) *DescribeRCCloudAssistantStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusRequest) SetRegionId(v string) *DescribeRCCloudAssistantStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRCCloudAssistantStatusRequest) Validate() error {
	return dara.Validate(s)
}
