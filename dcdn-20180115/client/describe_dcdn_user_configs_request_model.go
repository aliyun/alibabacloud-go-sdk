// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnUserConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFunctionName(v string) *DescribeDcdnUserConfigsRequest
	GetFunctionName() *string
}

type DescribeDcdnUserConfigsRequest struct {
	// The configuration that you want to query. Valid values:
	//
	// 	- domain_business_control: user configurations
	//
	// 	- bot_basic: the basic edition of bot traffic management, which supports authorized crawlers and provides threat intelligence
	//
	// 	- bot_Advance: the advanced edition of bot traffic management, which supports authorized crawlers and AI intelligent protection and provides threat intelligence
	//
	// This parameter is required.
	//
	// example:
	//
	// domain_business_control
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
}

func (s DescribeDcdnUserConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnUserConfigsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnUserConfigsRequest) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeDcdnUserConfigsRequest) SetFunctionName(v string) *DescribeDcdnUserConfigsRequest {
	s.FunctionName = &v
	return s
}

func (s *DescribeDcdnUserConfigsRequest) Validate() error {
	return dara.Validate(s)
}
