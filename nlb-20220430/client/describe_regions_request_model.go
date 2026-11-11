// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
	SetClientToken(v string) *DescribeRegionsRequest
	GetClientToken() *string
	SetServiceCode(v string) *DescribeRegionsRequest
	GetServiceCode() *string
}

type DescribeRegionsRequest struct {
	// The supported language. Valid values:
	//
	// - **zh-CN**: Chinese.
	//
	// - **en-US*	- (default): English.
	//
	// - **ja**: Japanese.
	//
	// example:
	//
	// en-US
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can generate the value from your client to ensure that the value is unique across different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- of each API request may be different.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The service code of the cloud service. Valid value: **nlb**.
	//
	// example:
	//
	// nlb
	ServiceCode *string `json:"ServiceCode,omitempty" xml:"ServiceCode,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeRegionsRequest) GetServiceCode() *string {
	return s.ServiceCode
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) SetClientToken(v string) *DescribeRegionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeRegionsRequest) SetServiceCode(v string) *DescribeRegionsRequest {
	s.ServiceCode = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}
