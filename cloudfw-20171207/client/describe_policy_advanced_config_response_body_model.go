// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyAdvancedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEips(v []*string) *DescribePolicyAdvancedConfigResponseBody
	GetEips() []*string
	SetInternetSwitch(v string) *DescribePolicyAdvancedConfigResponseBody
	GetInternetSwitch() *string
	SetRequestId(v string) *DescribePolicyAdvancedConfigResponseBody
	GetRequestId() *string
}

type DescribePolicyAdvancedConfigResponseBody struct {
	// The list of elastic IP addresses (EIPs) that take effect when strict mode is enabled. A maximum of 100 IP addresses can be returned.
	//
	// example:
	//
	// ["47.100.XX.XX","47.101.XX.XX"]
	Eips []*string `json:"Eips,omitempty" xml:"Eips,omitempty" type:"Repeated"`
	// The enabling status of the strict mode for access control policies. Valid values:
	//
	// - **on**: The strict mode is enabled.
	//
	// - **off**: The strict mode is disabled.
	//
	// example:
	//
	// off
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 850A84D6-0DE4-4797-A1E8-00090125EEB1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolicyAdvancedConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolicyAdvancedConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolicyAdvancedConfigResponseBody) GetEips() []*string {
	return s.Eips
}

func (s *DescribePolicyAdvancedConfigResponseBody) GetInternetSwitch() *string {
	return s.InternetSwitch
}

func (s *DescribePolicyAdvancedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolicyAdvancedConfigResponseBody) SetEips(v []*string) *DescribePolicyAdvancedConfigResponseBody {
	s.Eips = v
	return s
}

func (s *DescribePolicyAdvancedConfigResponseBody) SetInternetSwitch(v string) *DescribePolicyAdvancedConfigResponseBody {
	s.InternetSwitch = &v
	return s
}

func (s *DescribePolicyAdvancedConfigResponseBody) SetRequestId(v string) *DescribePolicyAdvancedConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolicyAdvancedConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
