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
	// - **zh-CN**: Chinese.
	//
	// - **en-US*	- (default): English.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A client-generated token that is used to ensure the idempotence of the request. The token must be unique among different requests and can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the domain name instance configuration. For the same access domain name and GTM instance, you can configure both A and AAAA records. This results in two domain name instance configurations for the GTM instance. The ConfigId uniquely identifies a specific configuration.
	//
	// Call the [ListCloudGtmInstanceConfigs](https://help.aliyun.com/document_detail/2797349.html) operation to query the ConfigId of a domain name instance.
	//
	// example:
	//
	// Config-000****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The enablement status of the domain name instance. Valid values:
	//
	// - enable: Enables the domain name instance.
	//
	// - disable: Disables the domain name instance.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The ID of the GTM 3.0 instance that you want to modify.
	//
	// example:
	//
	// gtm-cn-wwo3a3h****
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
