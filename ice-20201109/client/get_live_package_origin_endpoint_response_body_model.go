// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLivePackageOriginEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageOriginEndpoint(v *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) *GetLivePackageOriginEndpointResponseBody
	GetLivePackageOriginEndpoint() *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint
	SetRequestId(v string) *GetLivePackageOriginEndpointResponseBody
	GetRequestId() *string
}

type GetLivePackageOriginEndpointResponseBody struct {
	// The information about the origin endpoints.
	LivePackageOriginEndpoint *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint `json:"LivePackageOriginEndpoint,omitempty" xml:"LivePackageOriginEndpoint,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// requestIdExample123
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLivePackageOriginEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageOriginEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetLivePackageOriginEndpointResponseBody) GetLivePackageOriginEndpoint() *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	return s.LivePackageOriginEndpoint
}

func (s *GetLivePackageOriginEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLivePackageOriginEndpointResponseBody) SetLivePackageOriginEndpoint(v *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) *GetLivePackageOriginEndpointResponseBody {
	s.LivePackageOriginEndpoint = v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBody) SetRequestId(v string) *GetLivePackageOriginEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBody) Validate() error {
	if s.LivePackageOriginEndpoint != nil {
		if err := s.LivePackageOriginEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint struct {
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
	// 10.21.222.1/32
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
	// 5
	TimeshiftVision *int32 `json:"TimeshiftVision,omitempty" xml:"TimeshiftVision,omitempty"`
}

func (s GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) String() string {
	return dara.Prettify(s)
}

func (s GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GoString() string {
	return s.String()
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetChannelName() *string {
	return s.ChannelName
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetDescription() *string {
	return s.Description
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetEndpointName() *string {
	return s.EndpointName
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetGroupName() *string {
	return s.GroupName
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetLastModified() *string {
	return s.LastModified
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetLivePackagingConfig() *LivePackagingConfig {
	return s.LivePackagingConfig
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetManifestName() *string {
	return s.ManifestName
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetProtocol() *string {
	return s.Protocol
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetTimeshiftVision() *int32 {
	return s.TimeshiftVision
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetAuthorizationCode(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.AuthorizationCode = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetChannelName(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.ChannelName = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetCreateTime(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.CreateTime = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetDescription(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.Description = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetEndpointName(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.EndpointName = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetEndpointUrl(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.EndpointUrl = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetGroupName(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.GroupName = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetIpBlacklist(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.IpBlacklist = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetIpWhitelist(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.IpWhitelist = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetLastModified(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.LastModified = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetLivePackagingConfig(v *LivePackagingConfig) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.LivePackagingConfig = v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetManifestName(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.ManifestName = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetProtocol(v string) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.Protocol = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetTimeshiftVision(v int32) *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.TimeshiftVision = &v
	return s
}

func (s *GetLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) Validate() error {
	if s.LivePackagingConfig != nil {
		if err := s.LivePackagingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
