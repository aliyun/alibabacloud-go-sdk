// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmInstanceConfigFullInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmInstanceConfigFullInfoRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DescribeCloudGtmInstanceConfigFullInfoRequest
	GetClientToken() *string
	SetConfigId(v string) *DescribeCloudGtmInstanceConfigFullInfoRequest
	GetConfigId() *string
	SetInstanceId(v string) *DescribeCloudGtmInstanceConfigFullInfoRequest
	GetInstanceId() *string
}

type DescribeCloudGtmInstanceConfigFullInfoRequest struct {
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
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// You can call the [ListCloudGtmInstanceConfigs](~~ListCloudGtmInstanceConfigs~~) operation to query the value of ConfigId for the access domain name.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the GTM 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCloudGtmInstanceConfigFullInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) SetAcceptLanguage(v string) *DescribeCloudGtmInstanceConfigFullInfoRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) SetClientToken(v string) *DescribeCloudGtmInstanceConfigFullInfoRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) SetConfigId(v string) *DescribeCloudGtmInstanceConfigFullInfoRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) SetInstanceId(v string) *DescribeCloudGtmInstanceConfigFullInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoRequest) Validate() error {
	return dara.Validate(s)
}
