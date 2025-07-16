// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnUserConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *DescribeCdnUserConfigsRequest
	GetFunctionName() *string
}

type DescribeCdnUserConfigsRequest struct {
	// The configuration that you want to query. Valid values:
	//
	// 	- **domain_business_control**: user configurations
	//
	// 	- **waf**: Web Application Firewall (WAF) configurations
	//
	// This parameter is required.
	//
	// example:
	//
	// domain_business_control
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeCdnUserConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnUserConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnUserConfigsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeCdnUserConfigsRequest) SetFunctionName(v string) *DescribeCdnUserConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *DescribeCdnUserConfigsRequest) Validate() error {
	return dara.Validate(s)
}
