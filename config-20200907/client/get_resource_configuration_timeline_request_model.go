// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationTimelineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *GetResourceConfigurationTimelineRequest
	GetEndTime() *int64
	SetMaxResults(v int32) *GetResourceConfigurationTimelineRequest
	GetMaxResults() *int32
	SetNextToken(v string) *GetResourceConfigurationTimelineRequest
	GetNextToken() *string
	SetRegion(v string) *GetResourceConfigurationTimelineRequest
	GetRegion() *string
	SetResourceId(v string) *GetResourceConfigurationTimelineRequest
	GetResourceId() *string
	SetResourceType(v string) *GetResourceConfigurationTimelineRequest
	GetResourceType() *string
	SetStartTime(v int64) *GetResourceConfigurationTimelineRequest
	GetStartTime() *int64
}

type GetResourceConfigurationTimelineRequest struct {
	// The end of the time range to query. The default value indicates the time when the GetResourceConfigurationTimeline operation is called. Unit: milliseconds.
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
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// IWBjqMYSy0is7zSMGu16****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The resource IDs.
	//
	// For more information about how to query the ID of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// new-bucket
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource type.
	//
	// For more information about how to obtain the type of a resource, see [ListDiscoveredResources](https://help.aliyun.com/document_detail/169620.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// ACS::OSS::Bucket
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The beginning of the time range to query. By default, Cloud Config retrieves the configuration changes in the last 30 days for the specified resource. Unit: milliseconds.
	//
	// example:
	//
	// 1623211156000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetResourceConfigurationTimelineRequest) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationTimelineRequest) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetResourceConfigurationTimelineRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *GetResourceConfigurationTimelineRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetResourceConfigurationTimelineRequest) GetRegion() *string {
	return s.Region
}

func (s *GetResourceConfigurationTimelineRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetResourceConfigurationTimelineRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceConfigurationTimelineRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetResourceConfigurationTimelineRequest) SetEndTime(v int64) *GetResourceConfigurationTimelineRequest {
	s.EndTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetMaxResults(v int32) *GetResourceConfigurationTimelineRequest {
	s.MaxResults = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetNextToken(v string) *GetResourceConfigurationTimelineRequest {
	s.NextToken = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetRegion(v string) *GetResourceConfigurationTimelineRequest {
	s.Region = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetResourceId(v string) *GetResourceConfigurationTimelineRequest {
	s.ResourceId = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetResourceType(v string) *GetResourceConfigurationTimelineRequest {
	s.ResourceType = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) SetStartTime(v int64) *GetResourceConfigurationTimelineRequest {
	s.StartTime = &v
	return s
}

func (s *GetResourceConfigurationTimelineRequest) Validate() error {
	return dara.Validate(s)
}
