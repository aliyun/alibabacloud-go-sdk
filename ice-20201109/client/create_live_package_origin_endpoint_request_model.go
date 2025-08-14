// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLivePackageOriginEndpointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationCode(v string) *CreateLivePackageOriginEndpointRequest
	GetAuthorizationCode() *string
	SetChannelName(v string) *CreateLivePackageOriginEndpointRequest
	GetChannelName() *string
	SetClientToken(v string) *CreateLivePackageOriginEndpointRequest
	GetClientToken() *string
	SetDescription(v string) *CreateLivePackageOriginEndpointRequest
	GetDescription() *string
	SetEndpointName(v string) *CreateLivePackageOriginEndpointRequest
	GetEndpointName() *string
	SetGroupName(v string) *CreateLivePackageOriginEndpointRequest
	GetGroupName() *string
	SetIpBlacklist(v string) *CreateLivePackageOriginEndpointRequest
	GetIpBlacklist() *string
	SetIpWhitelist(v string) *CreateLivePackageOriginEndpointRequest
	GetIpWhitelist() *string
	SetLivePackagingConfig(v *LivePackagingConfig) *CreateLivePackageOriginEndpointRequest
	GetLivePackagingConfig() *LivePackagingConfig
	SetManifestName(v string) *CreateLivePackageOriginEndpointRequest
	GetManifestName() *string
	SetProtocol(v string) *CreateLivePackageOriginEndpointRequest
	GetProtocol() *string
	SetTimeshiftVision(v int32) *CreateLivePackageOriginEndpointRequest
	GetTimeshiftVision() *int32
}

type CreateLivePackageOriginEndpointRequest struct {
	// The authorization code. It can be up to 200 characters in length. You must configure AuthorizationCode, IpWhitelist, or both. Format: [A-Za-z0-9-_.]+
	//
	// example:
	//
	// AbcDef123
	AuthorizationCode *string `json:"AuthorizationCode,omitempty" xml:"AuthorizationCode,omitempty"`
	// The channel name.
	//
	// This parameter is required.
	//
	// example:
	//
	// channel-1
	ChannelName *string `json:"ChannelName,omitempty" xml:"ChannelName,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// ****0311a423d11a5f7dee713535****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	// The IP address blacklist. It supports subnet masks. 0.0.0.0/0 is not allowed. It can be up to 1,000 characters in length. Separate multiple IP addresses with commas (,).
	//
	// example:
	//
	// 103.21.222.1/32,192.168.100.0/24
	IpBlacklist *string `json:"IpBlacklist,omitempty" xml:"IpBlacklist,omitempty"`
	// The IP address whitelist. It supports subnet masks. 0.0.0.0/0 is not allowed. It can be up to 1,000 characters in length. Separate multiple IP addresses with commas (,). You must configure AuthorizationCode, IpWhitelist, or both.
	//
	// example:
	//
	// 192.168.1.0/24,10.0.0.1/24
	IpWhitelist         *string              `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
	LivePackagingConfig *LivePackagingConfig `json:"LivePackagingConfig,omitempty" xml:"LivePackagingConfig,omitempty"`
	// The playlist name. Default value: manifest.
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
	// The number of days that time-shifted content is available. Maximum value: 30. Default value: 0, which indicates that time shifting is not supported.
	//
	// example:
	//
	// 1
	TimeshiftVision *int32 `json:"TimeshiftVision,omitempty" xml:"TimeshiftVision,omitempty"`
}

func (s CreateLivePackageOriginEndpointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLivePackageOriginEndpointRequest) GoString() string {
	return s.String()
}

func (s *CreateLivePackageOriginEndpointRequest) GetAuthorizationCode() *string {
	return s.AuthorizationCode
}

func (s *CreateLivePackageOriginEndpointRequest) GetChannelName() *string {
	return s.ChannelName
}

func (s *CreateLivePackageOriginEndpointRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateLivePackageOriginEndpointRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLivePackageOriginEndpointRequest) GetEndpointName() *string {
	return s.EndpointName
}

func (s *CreateLivePackageOriginEndpointRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateLivePackageOriginEndpointRequest) GetIpBlacklist() *string {
	return s.IpBlacklist
}

func (s *CreateLivePackageOriginEndpointRequest) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *CreateLivePackageOriginEndpointRequest) GetLivePackagingConfig() *LivePackagingConfig {
	return s.LivePackagingConfig
}

func (s *CreateLivePackageOriginEndpointRequest) GetManifestName() *string {
	return s.ManifestName
}

func (s *CreateLivePackageOriginEndpointRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateLivePackageOriginEndpointRequest) GetTimeshiftVision() *int32 {
	return s.TimeshiftVision
}

func (s *CreateLivePackageOriginEndpointRequest) SetAuthorizationCode(v string) *CreateLivePackageOriginEndpointRequest {
	s.AuthorizationCode = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetChannelName(v string) *CreateLivePackageOriginEndpointRequest {
	s.ChannelName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetClientToken(v string) *CreateLivePackageOriginEndpointRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetDescription(v string) *CreateLivePackageOriginEndpointRequest {
	s.Description = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetEndpointName(v string) *CreateLivePackageOriginEndpointRequest {
	s.EndpointName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetGroupName(v string) *CreateLivePackageOriginEndpointRequest {
	s.GroupName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetIpBlacklist(v string) *CreateLivePackageOriginEndpointRequest {
	s.IpBlacklist = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetIpWhitelist(v string) *CreateLivePackageOriginEndpointRequest {
	s.IpWhitelist = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetLivePackagingConfig(v *LivePackagingConfig) *CreateLivePackageOriginEndpointRequest {
	s.LivePackagingConfig = v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetManifestName(v string) *CreateLivePackageOriginEndpointRequest {
	s.ManifestName = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetProtocol(v string) *CreateLivePackageOriginEndpointRequest {
	s.Protocol = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) SetTimeshiftVision(v int32) *CreateLivePackageOriginEndpointRequest {
	s.TimeshiftVision = &v
	return s
}

func (s *CreateLivePackageOriginEndpointRequest) Validate() error {
	return dara.Validate(s)
}
