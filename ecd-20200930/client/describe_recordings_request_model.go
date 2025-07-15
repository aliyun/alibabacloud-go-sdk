// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribeRecordingsRequest
	GetDesktopId() *string
	SetEndTime(v string) *DescribeRecordingsRequest
	GetEndTime() *string
	SetMaxResults(v int32) *DescribeRecordingsRequest
	GetMaxResults() *int32
	SetNeedSignedUrl(v bool) *DescribeRecordingsRequest
	GetNeedSignedUrl() *bool
	SetNextToken(v string) *DescribeRecordingsRequest
	GetNextToken() *string
	SetPolicyGroupId(v string) *DescribeRecordingsRequest
	GetPolicyGroupId() *string
	SetRegionId(v string) *DescribeRecordingsRequest
	GetRegionId() *string
	SetSignedUrlExpireMinutes(v int32) *DescribeRecordingsRequest
	GetSignedUrlExpireMinutes() *int32
	SetStandardEndTime(v string) *DescribeRecordingsRequest
	GetStandardEndTime() *string
	SetStandardStartTime(v string) *DescribeRecordingsRequest
	GetStandardStartTime() *string
	SetStartTime(v string) *DescribeRecordingsRequest
	GetStartTime() *string
}

type DescribeRecordingsRequest struct {
	// The cloud computer ID. If this parameter is not specified, the screen recording files on all cloud computers in the designated region will be queried.
	//
	// example:
	//
	// ecd-hlh41mk78dugw****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The end time of the query. Specify the time in the `YYYYMMDDhhmmss` format. The time must be in UTC+8.
	//
	// example:
	//
	// 20230424004441
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The maximum number of entries per page.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Specifies whether to return a URL.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	NeedSignedUrl *bool `json:"NeedSignedUrl,omitempty" xml:"NeedSignedUrl,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// aGN4YzAxQGNuLWhhbmd6aG91LjExNzU5NTMyNjgzMTQ1****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-gx2x1dhsmthe9****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The region ID. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The validity period of the returned URL. Unit: minutes.
	//
	// example:
	//
	// 10
	SignedUrlExpireMinutes *int32 `json:"SignedUrlExpireMinutes,omitempty" xml:"SignedUrlExpireMinutes,omitempty"`
	// The end time of the query. Specify the time in the ISO 8601 standard in the `yyyy-mm-ddthh:mm:ssz` format. The time must be in UTC+0.
	//
	// example:
	//
	// 2025-01-27T02:30:10Z
	StandardEndTime *string `json:"StandardEndTime,omitempty" xml:"StandardEndTime,omitempty"`
	// The start time of the query. Specify the time in the ISO 8601 standard in the `yyyy-mm-ddthh:mm:ssz` format. The time must be in UTC+0.
	//
	// example:
	//
	// 2025-01-27T02:20:10Z
	StandardStartTime *string `json:"StandardStartTime,omitempty" xml:"StandardStartTime,omitempty"`
	// The start time of the query. Specify the time in the `YYYYMMDDhhmmss` format. The time must be in UTC+8.
	//
	// example:
	//
	// 20230424000000
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRecordingsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordingsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecordingsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeRecordingsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRecordingsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeRecordingsRequest) GetNeedSignedUrl() *bool {
	return s.NeedSignedUrl
}

func (s *DescribeRecordingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRecordingsRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeRecordingsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRecordingsRequest) GetSignedUrlExpireMinutes() *int32 {
	return s.SignedUrlExpireMinutes
}

func (s *DescribeRecordingsRequest) GetStandardEndTime() *string {
	return s.StandardEndTime
}

func (s *DescribeRecordingsRequest) GetStandardStartTime() *string {
	return s.StandardStartTime
}

func (s *DescribeRecordingsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRecordingsRequest) SetDesktopId(v string) *DescribeRecordingsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribeRecordingsRequest) SetEndTime(v string) *DescribeRecordingsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordingsRequest) SetMaxResults(v int32) *DescribeRecordingsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeRecordingsRequest) SetNeedSignedUrl(v bool) *DescribeRecordingsRequest {
	s.NeedSignedUrl = &v
	return s
}

func (s *DescribeRecordingsRequest) SetNextToken(v string) *DescribeRecordingsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeRecordingsRequest) SetPolicyGroupId(v string) *DescribeRecordingsRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeRecordingsRequest) SetRegionId(v string) *DescribeRecordingsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRecordingsRequest) SetSignedUrlExpireMinutes(v int32) *DescribeRecordingsRequest {
	s.SignedUrlExpireMinutes = &v
	return s
}

func (s *DescribeRecordingsRequest) SetStandardEndTime(v string) *DescribeRecordingsRequest {
	s.StandardEndTime = &v
	return s
}

func (s *DescribeRecordingsRequest) SetStandardStartTime(v string) *DescribeRecordingsRequest {
	s.StandardStartTime = &v
	return s
}

func (s *DescribeRecordingsRequest) SetStartTime(v string) *DescribeRecordingsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordingsRequest) Validate() error {
	return dara.Validate(s)
}
