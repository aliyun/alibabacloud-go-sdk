// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageOriginEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageOriginEndpoint(v *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) *CreateLivePackageOriginEndpointResponseBody
	GetLivePackageOriginEndpoint() *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint
	SetRequestId(v string) *CreateLivePackageOriginEndpointResponseBody
	GetRequestId() *string
}

type CreateLivePackageOriginEndpointResponseBody struct {
	// The information about the origin endpoint.
	LivePackageOriginEndpoint *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint `json:"LivePackageOriginEndpoint,omitempty" xml:"LivePackageOriginEndpoint,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateLivePackageOriginEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageOriginEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *CreateLivePackageOriginEndpointResponseBody) GetLivePackageOriginEndpoint() *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	return s.LivePackageOriginEndpoint
}

func (s *CreateLivePackageOriginEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateLivePackageOriginEndpointResponseBody) SetLivePackageOriginEndpoint(v *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) *CreateLivePackageOriginEndpointResponseBody {
	s.LivePackageOriginEndpoint = v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBody) SetRequestId(v string) *CreateLivePackageOriginEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint struct {
	// The authorization code.
	//
	// example:
	//
	// Abcded123
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
	// https://xxx.packagepull-abcxxx.ap-southeast-1.aliyuncsiceintl.com/v1/group01/1/ch01/manifest
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
	// 103.21.222.1/32,192.168.100.0/24
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
	LastModified        *string              `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	LivePackagingConfig *LivePackagingConfig `json:"LivePackagingConfig,omitempty" xml:"LivePackagingConfig,omitempty"`
	// The playlist name.
	//
	// example:
	//
	// manifest
	ManifestName *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	// The protocol. Only HLS is supported.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of days that time-shifted content is available. Maximum value: 30.
	//
	// example:
	//
	// 1
	TimeshiftVision *int32 `json:"TimeshiftVision,omitempty" xml:"TimeshiftVision,omitempty"`
}

func (s CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GoString() string {
	return s.String()
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetDescription() *string {
	return s.Description
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetEndpointName() *string {
	return s.EndpointName
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetLastModified() *string {
	return s.LastModified
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetLivePackagingConfig() *LivePackagingConfig {
	return s.LivePackagingConfig
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetManifestName() *string {
	return s.ManifestName
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetTimeshiftVision() *int32 {
	return s.TimeshiftVision
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetAuthorizationCode(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.AuthorizationCode = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetChannelName(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.ChannelName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetCreateTime(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.CreateTime = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetDescription(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.Description = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetEndpointName(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.EndpointName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetEndpointUrl(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.EndpointUrl = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetGroupName(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.GroupName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetIpBlacklist(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.IpBlacklist = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetIpWhitelist(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.IpWhitelist = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetLastModified(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.LastModified = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetLivePackagingConfig(v *LivePackagingConfig) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.LivePackagingConfig = v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetManifestName(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.ManifestName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetProtocol(v string) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.Protocol = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetTimeshiftVision(v int32) *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.TimeshiftVision = &v
	return s
}

func (s *CreateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) Validate() error {
	return dara.Validate(s)
}
