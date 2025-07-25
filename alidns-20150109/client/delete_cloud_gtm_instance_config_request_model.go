// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DeleteCloudGtmInstanceConfigRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DeleteCloudGtmInstanceConfigRequest
	GetClientToken() *string
	SetConfigId(v string) *DeleteCloudGtmInstanceConfigRequest
	GetConfigId() *string
	SetInstanceId(v string) *DeleteCloudGtmInstanceConfigRequest
	GetInstanceId() *string
}

type DeleteCloudGtmInstanceConfigRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US (default): English
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
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration. You can call the [ListCloudGtmInstanceConfigs](~~ListCloudGtmInstanceConfigs~~) operation to query the configuration ID of the access domain name.
	//
	// example:
	//
	// config-000**1
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The access domain name that is configured for the desired GTM 3.0 instance. You can delete only one access domain name.
	//
	// example:
	//
	// gtm-cn-jmp3qnw**03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteCloudGtmInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmInstanceConfigRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DeleteCloudGtmInstanceConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteCloudGtmInstanceConfigRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DeleteCloudGtmInstanceConfigRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCloudGtmInstanceConfigRequest) SetAcceptLanguage(v string) *DeleteCloudGtmInstanceConfigRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DeleteCloudGtmInstanceConfigRequest) SetClientToken(v string) *DeleteCloudGtmInstanceConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteCloudGtmInstanceConfigRequest) SetConfigId(v string) *DeleteCloudGtmInstanceConfigRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteCloudGtmInstanceConfigRequest) SetInstanceId(v string) *DeleteCloudGtmInstanceConfigRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCloudGtmInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
