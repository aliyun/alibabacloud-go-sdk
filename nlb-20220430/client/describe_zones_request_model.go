// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeZonesRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DescribeZonesRequest
	GetClientToken() *string
	SetRegionId(v string) *DescribeZonesRequest
	GetRegionId() *string
	SetServiceCode(v string) *DescribeZonesRequest
	GetServiceCode() *string
}

type DescribeZonesRequest struct {
	// The supported language. Valid values:
	//
	// - **zh-CN**: Chinese.
	//
	// - **en-US**: English. This is the default value.
	//
	// - **ja**: Japanese.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must ensure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The region ID of the zone. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/443657.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The service code of the cloud service. Valid value: **nlb**.
	//
	// example:
	//
	// nlb
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeZonesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeZonesRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeZonesRequest) SetClientToken(v string) *DescribeZonesRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeZonesRequest) SetRegionId(v string) *DescribeZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeZonesRequest) SetServiceCode(v string) *DescribeZonesRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}
