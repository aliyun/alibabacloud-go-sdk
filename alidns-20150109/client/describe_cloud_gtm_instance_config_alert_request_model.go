// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmInstanceConfigAlertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeCloudGtmInstanceConfigAlertRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DescribeCloudGtmInstanceConfigAlertRequest
	GetClientToken() *string
	SetConfigId(v string) *DescribeCloudGtmInstanceConfigAlertRequest
	GetConfigId() *string
	SetInstanceId(v string) *DescribeCloudGtmInstanceConfigAlertRequest
	GetInstanceId() *string
}

type DescribeCloudGtmInstanceConfigAlertRequest struct {
	// The language of the response. Valid values:
	//
	// 	- **zh-CN**: Chinese
	//
	// 	- **en-US**: English
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
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCloudGtmInstanceConfigAlertRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigAlertRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) SetAcceptLanguage(v string) *DescribeCloudGtmInstanceConfigAlertRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) SetClientToken(v string) *DescribeCloudGtmInstanceConfigAlertRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) SetConfigId(v string) *DescribeCloudGtmInstanceConfigAlertRequest {
	s.ConfigId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) SetInstanceId(v string) *DescribeCloudGtmInstanceConfigAlertRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigAlertRequest) Validate() error {
	return dara.Validate(s)
}
