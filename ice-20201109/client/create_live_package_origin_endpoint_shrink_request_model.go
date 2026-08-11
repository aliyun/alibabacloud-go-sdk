// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageOriginEndpointShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationCode(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetAuthorizationCode() *string
	SetChannelName(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetChannelName() *string
	SetClientToken(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetClientToken() *string
	SetDescription(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetDescription() *string
	SetEndpointName(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetEndpointName() *string
	SetGroupName(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetGroupName() *string
	SetIpBlacklist(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetIpBlacklist() *string
	SetIpWhitelist(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetIpWhitelist() *string
	SetLivePackagingConfigShrink(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetLivePackagingConfigShrink() *string
	SetManifestName(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetManifestName() *string
	SetProtocol(v string) *CreateLivePackageOriginEndpointShrinkRequest
	GetProtocol() *string
	SetTimeshiftVision(v int32) *CreateLivePackageOriginEndpointShrinkRequest
	GetTimeshiftVision() *int32
}

type CreateLivePackageOriginEndpointShrinkRequest struct {
	// The authorization code. You must specify at least one of AuthorizationCode and IpWhitelist. Maximum length: 200 characters. Format: [A-Za-z0-9-_.]+
	//
	// example:
	//
	// AbcDef123
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// The name of the channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The idempotency token.
	//
	// example:
	//
	// ****0311a423d11a5f7dee713535****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the endpoint.
	//
	// example:
	//
	// This is an origin endpoint.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the origin endpoint. The name must contain only letters, digits, hyphens (-), and underscores (_). The name must be 1 to 200 characters in length. Format: [A-Za-z0-9_-]+
	//
	// This parameter is required.
	//
	// example:
	//
	// endpoint-1
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	// The name of the channel group.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-group-1
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP blacklist. Subnet masks are supported. The value 0.0.0.0/0 is not allowed. Separate multiple IP addresses with commas (,). Maximum length: 1000 characters.
	//
	// example:
	//
	// 103.21.222.1/32,192.168.100.0/24
	IpBlacklist *string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The IP whitelist. Subnet masks are supported. The value 0.0.0.0/0 is not allowed. Separate multiple IP addresses with commas (,). You must specify at least one of IpWhitelist and AuthorizationCode. Maximum length: 1000 characters.
	//
	// example:
	//
	// 192.168.1.0/24,10.0.0.1/24
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	// The live packaging configuration.
	LivePackagingConfigShrink *string `json:"LivePackagingConfig,omitempty" xml:"LivePackagingConfig,omitempty"`
	// The manifest name. Default value: manifest.
	//
	// example:
	//
	// manifest
	ManifestName *string `json:"ManifestName,omitempty" xml:"ManifestName,omitempty"`
	// The distribution protocol.
	//
	// This parameter is required.
	//
	// example:
	//
	// HLS
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The number of days for time-shifting. Maximum value: 30. Default value: 0, which indicates that time-shifting is not supported.
	//
	// example:
	//
	// 1
	TimeshiftVision *int32 `json:"TimeshiftVision,omitempty" xml:"TimeshiftVision,omitempty"`
}

func (s CreateLivePackageOriginEndpointShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageOriginEndpointShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetLivePackagingConfigShrink() *string {
	return s.LivePackagingConfigShrink
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetManifestName() *string {
	return s.ManifestName
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) GetTimeshiftVision() *int32 {
	return s.TimeshiftVision
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetAuthorizationCode(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.AuthorizationCode = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetChannelName(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetClientToken(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetDescription(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetEndpointName(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.EndpointName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetGroupName(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.GroupName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetIpBlacklist(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.IpBlacklist = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetIpWhitelist(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.IpWhitelist = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetLivePackagingConfigShrink(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.LivePackagingConfigShrink = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetManifestName(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.ManifestName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetProtocol(v string) *CreateLivePackageOriginEndpointShrinkRequest {
	s.Protocol = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) SetTimeshiftVision(v int32) *CreateLivePackageOriginEndpointShrinkRequest {
	s.TimeshiftVision = &v
	return s
}

func (s *CreateLivePackageOriginEndpointShrinkRequest) Validate() error {
	return dara.Validate(s)
}
