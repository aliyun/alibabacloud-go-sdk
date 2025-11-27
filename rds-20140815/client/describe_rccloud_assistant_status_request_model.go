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
	// The list of instance IDs.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The maximum number of entries per page. If you specify `InstanceId`, this parameter does not take effect.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that marks the end of the current returned page. If this parameter is empty, the data is queried from the first entry.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The operating system type of the instance. Only **Linux*	- is supported.
	//
	// Valid values:
	//
	// 	- Windows
	//
	// 	- Linux
	//
	// 	- FreeBSD
	//
	// example:
	//
	// Linux
	OSType *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use `NextToken` and `MaxResults` for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use `NextToken` and `MaxResults` for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
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
