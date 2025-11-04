// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageOriginEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLivePackageOriginEndpoint(v *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) *UpdateLivePackageOriginEndpointResponseBody
	GetLivePackageOriginEndpoint() *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint
	SetRequestId(v string) *UpdateLivePackageOriginEndpointResponseBody
	GetRequestId() *string
}

type UpdateLivePackageOriginEndpointResponseBody struct {
	// The information about the origin endpoint.
	LivePackageOriginEndpoint *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint `json:"LivePackageOriginEndpoint,omitempty" xml:"LivePackageOriginEndpoint,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// b1f8d6c4-a123-4cd5-9e88-d0819e3bfa70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLivePackageOriginEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageOriginEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageOriginEndpointResponseBody) GetLivePackageOriginEndpoint() *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	return s.LivePackageOriginEndpoint
}

func (s *UpdateLivePackageOriginEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLivePackageOriginEndpointResponseBody) SetLivePackageOriginEndpoint(v *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) *UpdateLivePackageOriginEndpointResponseBody {
	s.LivePackageOriginEndpoint = v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBody) SetRequestId(v string) *UpdateLivePackageOriginEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBody) Validate() error {
	if s.LivePackageOriginEndpoint != nil {
		if err := s.LivePackageOriginEndpoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint struct {
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
	// The IP address blacklist. It supports subnet masks. Multiple IP addresses are separated by commas (,).
	//
	// example:
	//
	// 10.21.222.1/32,192.168.100.0/24
	IpBlacklist *string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The IP address whitelist. It supports subnet masks. Multiple IP addresses are separated by commas (,).
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
	// The playlist name. Default value: manifest.
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
	// 5
	TimeshiftVision *int32 `json:"TimeshiftVision,omitempty" xml:"TimeshiftVision,omitempty"`
}

func (s UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetDescription() *string {
	return s.Description
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetEndpointName() *string {
	return s.EndpointName
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetEndpointUrl() *string {
	return s.EndpointUrl
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetLastModified() *string {
	return s.LastModified
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetLivePackagingConfig() *LivePackagingConfig {
	return s.LivePackagingConfig
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetManifestName() *string {
	return s.ManifestName
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) GetTimeshiftVision() *int32 {
	return s.TimeshiftVision
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetAuthorizationCode(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.AuthorizationCode = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetChannelName(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.ChannelName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetCreateTime(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.CreateTime = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetDescription(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.Description = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetEndpointName(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.EndpointName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetEndpointUrl(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.EndpointUrl = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetGroupName(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.GroupName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetIpBlacklist(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.IpBlacklist = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetIpWhitelist(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.IpWhitelist = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetLastModified(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.LastModified = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetLivePackagingConfig(v *LivePackagingConfig) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.LivePackagingConfig = v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetManifestName(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.ManifestName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetProtocol(v string) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.Protocol = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) SetTimeshiftVision(v int32) *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint {
	s.TimeshiftVision = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointResponseBodyLivePackageOriginEndpoint) Validate() error {
	if s.LivePackagingConfig != nil {
		if err := s.LivePackagingConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
