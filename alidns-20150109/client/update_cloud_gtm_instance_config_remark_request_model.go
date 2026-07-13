// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmInstanceConfigRemarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigRemarkRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *UpdateCloudGtmInstanceConfigRemarkRequest
	GetClientToken() *string
	SetConfigId(v string) *UpdateCloudGtmInstanceConfigRemarkRequest
	GetConfigId() *string
	SetInstanceId(v string) *UpdateCloudGtmInstanceConfigRemarkRequest
	GetInstanceId() *string
	SetRemark(v string) *UpdateCloudGtmInstanceConfigRemarkRequest
	GetRemark() *string
}

type UpdateCloudGtmInstanceConfigRemarkRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh-CN**: Chinese.
	//
	// - **en-US*	- (default): English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request. Generate a unique token for each request. The token can contain a maximum of 64 ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance configuration. A GTM instance can have multiple configurations if you add both an A record and an AAAA record for the same domain name. The ConfigId uniquely identifies the configuration that you want to update.
	//
	// For more information, see [ListCloudGtmInstanceConfigs](https://help.aliyun.com/document_detail/2797349.html).
	//
	// example:
	//
	// Config-000****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the Global Traffic Manager (GTM) instance for which you want to update the remarks.
	//
	// example:
	//
	// gtm-cn-wwo3a3h****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new remarks for the configuration.
	//
	// example:
	//
	// API
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s UpdateCloudGtmInstanceConfigRemarkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmInstanceConfigRemarkRequest) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) GetRemark() *string {
	return s.Remark
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) SetAcceptLanguage(v string) *UpdateCloudGtmInstanceConfigRemarkRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) SetClientToken(v string) *UpdateCloudGtmInstanceConfigRemarkRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) SetConfigId(v string) *UpdateCloudGtmInstanceConfigRemarkRequest {
	s.ConfigId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) SetInstanceId(v string) *UpdateCloudGtmInstanceConfigRemarkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) SetRemark(v string) *UpdateCloudGtmInstanceConfigRemarkRequest {
	s.Remark = &v
	return s
}

func (s *UpdateCloudGtmInstanceConfigRemarkRequest) Validate() error {
	return dara.Validate(s)
}
