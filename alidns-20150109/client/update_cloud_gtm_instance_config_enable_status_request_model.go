// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigEnableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest
	GetClientToken() *string
	SetConfigId(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest
	GetConfigId() *string
	SetEnableStatus(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest
	GetEnableStatus() *string
	SetInstanceId(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest
	GetInstanceId() *string
}

type UpdateCloudGtmInstanceConfigEnableStatusRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// You can call the [ListCloudGtmInstanceConfigs](~~ListCloudGtmInstanceConfigs~~) operation to query the configuration ID of the access domain name.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The enabling state of the access domain name. Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigEnableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) SetClientToken(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) SetConfigId(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) SetEnableStatus(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest {
	s.EnableStatus = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) SetInstanceId(v string) *UpdateCloudGtmInstanceConfigEnableStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigEnableStatusRequest) Validate() error {
	return dara.Validate(s)
}
