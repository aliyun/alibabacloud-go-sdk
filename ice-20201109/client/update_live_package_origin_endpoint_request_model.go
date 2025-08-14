// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageOriginEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationCode(v string) *UpdateLivePackageOriginEndpointRequest
	GetAuthorizationCode() *string
	SetChannelName(v string) *UpdateLivePackageOriginEndpointRequest
	GetChannelName() *string
	SetDescription(v string) *UpdateLivePackageOriginEndpointRequest
	GetDescription() *string
	SetEndpointName(v string) *UpdateLivePackageOriginEndpointRequest
	GetEndpointName() *string
	SetGroupName(v string) *UpdateLivePackageOriginEndpointRequest
	GetGroupName() *string
	SetIpBlacklist(v string) *UpdateLivePackageOriginEndpointRequest
	GetIpBlacklist() *string
	SetIpWhitelist(v string) *UpdateLivePackageOriginEndpointRequest
	GetIpWhitelist() *string
	SetLivePackagingConfig(v *LivePackagingConfig) *UpdateLivePackageOriginEndpointRequest
	GetLivePackagingConfig() *LivePackagingConfig
	SetManifestName(v string) *UpdateLivePackageOriginEndpointRequest
	GetManifestName() *string
	SetProtocol(v string) *UpdateLivePackageOriginEndpointRequest
	GetProtocol() *string
	SetTimeshiftVision(v int32) *UpdateLivePackageOriginEndpointRequest
	GetTimeshiftVision() *int32
}

type UpdateLivePackageOriginEndpointRequest struct {
	// The authorization code. It can be up to 200 characters in length. You must configure AuthorizationCode, IpWhitelist, or both. Format: [A-Za-z0-9-_.]+
	//
	// example:
	//
	// Abc123Def456
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// The channel name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The endpoint description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The origin endpoint name. It can contain letters, digits, hyphens (-), and underscores (_). The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// endpoint-1
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The channel group name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP address blacklist. It supports subnet masks. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 103.0.0.0/8
	IpBlacklist *string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The IP address whitelist. It supports subnet masks. 0.0.0.0/0 is not allowed. It can be up to 1,000 characters in length. Separate multiple IP addresses with commas (,). You must configure AuthorizationCode, IpWhitelist, or both.
	//
	// example:
	//
	// 192.168.1.0/24,10.0.0.1
	IpWhitelist         *string              `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	LivePackagingConfig *LivePackagingConfig `json:"LivePackagingConfig,omitempty" xml:"LivePackagingConfig,omitempty"`
	// The playlist name. Default value: manifest.
	//
	// example:
	//
	// manifest
	ManifestName *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	// The protocol. Only HLS is supported.
	//
	// This parameter is required.
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

func (s UpdateLivePackageOriginEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageOriginEndpointRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageOriginEndpointRequest) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *UpdateLivePackageOriginEndpointRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateLivePackageOriginEndpointRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLivePackageOriginEndpointRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *UpdateLivePackageOriginEndpointRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLivePackageOriginEndpointRequest) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *UpdateLivePackageOriginEndpointRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *UpdateLivePackageOriginEndpointRequest) GetLivePackagingConfig() *LivePackagingConfig {
	return s.LivePackagingConfig
}

func (s *UpdateLivePackageOriginEndpointRequest) GetManifestName() *string {
	return s.ManifestName
}

func (s *UpdateLivePackageOriginEndpointRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateLivePackageOriginEndpointRequest) GetTimeshiftVision() *int32 {
	return s.TimeshiftVision
}

func (s *UpdateLivePackageOriginEndpointRequest) SetAuthorizationCode(v string) *UpdateLivePackageOriginEndpointRequest {
	s.AuthorizationCode = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetChannelName(v string) *UpdateLivePackageOriginEndpointRequest {
	s.ChannelName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetDescription(v string) *UpdateLivePackageOriginEndpointRequest {
	s.Description = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetEndpointName(v string) *UpdateLivePackageOriginEndpointRequest {
	s.EndpointName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetGroupName(v string) *UpdateLivePackageOriginEndpointRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetIpBlacklist(v string) *UpdateLivePackageOriginEndpointRequest {
	s.IpBlacklist = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetIpWhitelist(v string) *UpdateLivePackageOriginEndpointRequest {
	s.IpWhitelist = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetLivePackagingConfig(v *LivePackagingConfig) *UpdateLivePackageOriginEndpointRequest {
	s.LivePackagingConfig = v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetManifestName(v string) *UpdateLivePackageOriginEndpointRequest {
	s.ManifestName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetProtocol(v string) *UpdateLivePackageOriginEndpointRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) SetTimeshiftVision(v int32) *UpdateLivePackageOriginEndpointRequest {
	s.TimeshiftVision = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointRequest) Validate() error {
	return dara.Validate(s)
}
