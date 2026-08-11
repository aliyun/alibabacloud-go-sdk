// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageOriginEndpointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationCode(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetAuthorizationCode() *string
	SetChannelName(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetChannelName() *string
	SetDescription(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetDescription() *string
	SetEndpointName(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetEndpointName() *string
	SetGroupName(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetGroupName() *string
	SetIpBlacklist(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetIpBlacklist() *string
	SetIpWhitelist(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetIpWhitelist() *string
	SetLivePackagingConfigShrink(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetLivePackagingConfigShrink() *string
	SetManifestName(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetManifestName() *string
	SetProtocol(v string) *UpdateLivePackageOriginEndpointShrinkRequest
	GetProtocol() *string
	SetTimeshiftVision(v int32) *UpdateLivePackageOriginEndpointShrinkRequest
	GetTimeshiftVision() *int32
}

type UpdateLivePackageOriginEndpointShrinkRequest struct {
	// The authorization code. You must specify at least one of AuthorizationCode and IpWhitelist. Maximum length: 200 characters. Format: [A-Za-z0-9-_.]+
	//
	// example:
	//
	// Abc123Def456
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// The name of an existing channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The description of the endpoint.
	//
	// example:
	//
	// This is an origin endpoint.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the origin endpoint. The name can contain uppercase and lowercase letters, digits, hyphens (-), and underscores (_). The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// endpoint-1
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The name of an existing channel group.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP blacklist. Subnet masks are supported. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 103.0.0.0/8
	IpBlacklist *string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The IP whitelist. Subnet masks are supported. The value 0.0.0.0/0 is not allowed. Separate multiple IP addresses with commas (,). You must specify at least one of IpWhitelist and AuthorizationCode. Maximum length: 1000 characters.
	//
	// example:
	//
	// 192.168.1.0/24,10.0.0.1
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The live packaging configuration.
	LivePackagingConfigShrink *string `json:"LivePackagingConfig,omitempty" xml:"LivePackagingConfig,omitempty"`
	// The playlist name. Default value: manifest.
	//
	// example:
	//
	// manifest
	ManifestName *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	// The protocol. Currently, only HLS is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of time-shifting days. Maximum value: 30.
	//
	// example:
	//
	// 5
	TimeshiftVision *int32 `json:"TimeshiftVision,omitempty" xml:"TimeshiftVision,omitempty"`
}

func (s UpdateLivePackageOriginEndpointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageOriginEndpointShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetLivePackagingConfigShrink() *string {
	return s.LivePackagingConfigShrink
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetManifestName() *string {
	return s.ManifestName
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) GetTimeshiftVision() *int32 {
	return s.TimeshiftVision
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetAuthorizationCode(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.AuthorizationCode = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetChannelName(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.ChannelName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetDescription(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetEndpointName(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.EndpointName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetGroupName(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetIpBlacklist(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.IpBlacklist = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetIpWhitelist(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.IpWhitelist = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetLivePackagingConfigShrink(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.LivePackagingConfigShrink = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetManifestName(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.ManifestName = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetProtocol(v string) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) SetTimeshiftVision(v int32) *UpdateLivePackageOriginEndpointShrinkRequest {
	s.TimeshiftVision = &v
	return s
}

func (s *UpdateLivePackageOriginEndpointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
