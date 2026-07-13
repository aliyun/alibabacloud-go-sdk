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
	// - zh-CN: Chinese
	//
	// - en-US (default): English
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
	// The ID of the instance configuration. You can configure both A and AAAA records for the same access domain name and Global Traffic Manager (GTM) instance. In this case, the GTM instance has two configurations. The ConfigId uniquely identifies an instance configuration.
	//
	// For more information, see [ListCloudGtmInstanceConfigs](https://help.aliyun.com/document_detail/2797349.html).
	//
	// example:
	//
	// Config-000****
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The ID of the Global Traffic Manager 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3h****
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
