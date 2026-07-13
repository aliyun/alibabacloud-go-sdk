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
	// - **zh-CN**: Chinese.
	//
	// - **en-US**: English.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// A client token that is used to ensure the idempotence of the request. The client generates this value. The value must be unique among different requests. It can be up to 64 ASCII characters in length and cannot contain non-ASCII characters.
	//
	// example:
	//
	// 1ae05db4-10e7-11ef-b126-00163e24**22
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the instance configuration. For the same access domain name and GTM instance, you can configure both A and AAAA records. This creates two instance configurations. \\`ConfigId\\` uniquely identifies an instance configuration.
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
