// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceTimelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetResourceComplianceTimelineRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *GetResourceComplianceTimelineRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetResourceComplianceTimelineRequest
	GetNextToken() *string
	SetRegion(v string) *GetResourceComplianceTimelineRequest
	GetRegion() *string
	SetResourceId(v string) *GetResourceComplianceTimelineRequest
	GetResourceId() *string
	SetResourceType(v string) *GetResourceComplianceTimelineRequest
	GetResourceType() *string
	SetStartTime(v int64) *GetResourceComplianceTimelineRequest
	GetStartTime() *int64
}

type GetResourceComplianceTimelineRequest struct {
	// The timestamp that specifies the end of the time range to query. The default value is the time when the GetResourceComplianceTimeline operation is called. Unit: milliseconds.
	//
	// example:
	//
	// 1625821156000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries to return for a single request. Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token that is used to initiate the next request. If the response of the current request is truncated, this token is used to initiate another request and obtain the remaining entries.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the resource resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The ID of the resource.
	//
	// For more information about how to obtain the ID of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// new-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource.
	//
	// For more information about how to obtain the type of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The timestamp that specifies the beginning of the time range to query. By default, Cloud Config retrieves the compliance evaluations in the last 30 days for the specified resource. Unit: milliseconds.
	//
	// example:
	//
	// 1623211156000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetResourceComplianceTimelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetResourceComplianceTimelineRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourceComplianceTimelineRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourceComplianceTimelineRequest) GetRegion() *string {
	return s.Region
}

func (s *GetResourceComplianceTimelineRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourceComplianceTimelineRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceComplianceTimelineRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetResourceComplianceTimelineRequest) SetEndTime(v int64) *GetResourceComplianceTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetMaxResults(v int32) *GetResourceComplianceTimelineRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetNextToken(v string) *GetResourceComplianceTimelineRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetRegion(v string) *GetResourceComplianceTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetResourceId(v string) *GetResourceComplianceTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetResourceType(v string) *GetResourceComplianceTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) SetStartTime(v int64) *GetResourceComplianceTimelineRequest {
	s.StartTime = &v
	return s
}

func (s *GetResourceComplianceTimelineRequest) Validate() error {
	return dara.Validate(s)
}
