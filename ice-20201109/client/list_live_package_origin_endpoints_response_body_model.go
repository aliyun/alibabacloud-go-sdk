// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLivePackageOriginEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageOriginEndpoints(v []*ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) *ListLivePackageOriginEndpointsResponseBody
	GetLivePackageOriginEndpoints() []*ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints
	SetPageNo(v int64) *ListLivePackageOriginEndpointsResponseBody
	GetPageNo() *int64
	SetPageSize(v int64) *ListLivePackageOriginEndpointsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListLivePackageOriginEndpointsResponseBody
	GetRequestId() *string
	SetSortBy(v string) *ListLivePackageOriginEndpointsResponseBody
	GetSortBy() *string
	SetTotalCount(v int64) *ListLivePackageOriginEndpointsResponseBody
	GetTotalCount() *int64
}

type ListLivePackageOriginEndpointsResponseBody struct {
	// The origin endpoints returned.
	LivePackageOriginEndpoints []*ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints `json:"LivePackageOriginEndpoints,omitempty" xml:"LivePackageOriginEndpoints,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// b9f90a7ac8904db28dc18e0c2a72c75d
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The sort order. Valid values: `asc` and `desc` (default).
	//
	// example:
	//
	// desc
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLivePackageOriginEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageOriginEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListLivePackageOriginEndpointsResponseBody) GetLivePackageOriginEndpoints() []*ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	return s.LivePackageOriginEndpoints
}

func (s *ListLivePackageOriginEndpointsResponseBody) GetPageNo() *int64 {
	return s.PageNo
}

func (s *ListLivePackageOriginEndpointsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListLivePackageOriginEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLivePackageOriginEndpointsResponseBody) GetSortBy() *string {
	return s.SortBy
}

func (s *ListLivePackageOriginEndpointsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListLivePackageOriginEndpointsResponseBody) SetLivePackageOriginEndpoints(v []*ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) *ListLivePackageOriginEndpointsResponseBody {
	s.LivePackageOriginEndpoints = v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBody) SetPageNo(v int64) *ListLivePackageOriginEndpointsResponseBody {
	s.PageNo = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBody) SetPageSize(v int64) *ListLivePackageOriginEndpointsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBody) SetRequestId(v string) *ListLivePackageOriginEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBody) SetSortBy(v string) *ListLivePackageOriginEndpointsResponseBody {
	s.SortBy = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBody) SetTotalCount(v int64) *ListLivePackageOriginEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints struct {
	// The authorization code.
	//
	// example:
	//
	// Abc123Def456
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// The channel name.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The time when the endpoint was created.
	//
	// example:
	//
	// 2023-04-01T12:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The endpoint description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The endpoint name.
	//
	// example:
	//
	// endpoint-1
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The endpoint URL.
	//
	// example:
	//
	// https://xxx.packagepull-abcxxx.ap-southeast-1.aliyuncsiceintl.com/v1/group01/1/ch01/manifest.m3u8
	EndpointUrl *string `json:"EndpointUrl,omitempty" xml:"EndpointUrl,omitempty"`
	// The channel group name.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP address blacklist.
	//
	// example:
	//
	// 10.21.222.1/32,192.168.100.0/24
	IpBlacklist *string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The IP address whitelist.
	//
	// example:
	//
	// 192.168.1.0/24,10.0.0.1/24
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The time when the endpoint was last modified.
	//
	// example:
	//
	// 2023-04-01T12:00:00Z
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	// The playlist name.
	//
	// example:
	//
	// manifest
	ManifestName *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	// The distribution protocol.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of days that time-shifted content is available.
	//
	// example:
	//
	// 1
	TimeshiftVision *int32 `json:"TimeshiftVision,omitempty" xml:"TimeshiftVision,omitempty"`
}

func (s ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) String() string {
	return dara.Prettify(s)
}

func (s ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GoString() string {
	return s.String()
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetChannelName() *string {
	return s.ChannelName
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetDescription() *string {
	return s.Description
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetGroupName() *string {
	return s.GroupName
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetLastModified() *string {
	return s.LastModified
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetManifestName() *string {
	return s.ManifestName
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetProtocol() *string {
	return s.Protocol
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) GetTimeshiftVision() *int32 {
	return s.TimeshiftVision
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetAuthorizationCode(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.AuthorizationCode = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetChannelName(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.ChannelName = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetCreateTime(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.CreateTime = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetDescription(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.Description = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetEndpointName(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.EndpointName = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetEndpointUrl(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.EndpointUrl = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetGroupName(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.GroupName = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetIpBlacklist(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.IpBlacklist = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetIpWhitelist(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.IpWhitelist = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetLastModified(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.LastModified = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetManifestName(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.ManifestName = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetProtocol(v string) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.Protocol = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) SetTimeshiftVision(v int32) *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints {
	s.TimeshiftVision = &v
	return s
}

func (s *ListLivePackageOriginEndpointsResponseBodyLivePackageOriginEndpoints) Validate() error {
	return dara.Validate(s)
}
