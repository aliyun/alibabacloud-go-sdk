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
	// The language in which the returned results are displayed. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US*	- (default): English
	//
	// example:
	//
	// en-US
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
	// The ID of the GTM 3.0 instance for which you want to modify the description.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
