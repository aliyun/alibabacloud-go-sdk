// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolicyAdvancedConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInternetSwitch(v string) *DescribePolicyAdvancedConfigResponseBody
	GetInternetSwitch() *string
	SetRequestId(v string) *DescribePolicyAdvancedConfigResponseBody
	GetRequestId() *string
}

type DescribePolicyAdvancedConfigResponseBody struct {
	// Indicates whether the strict mode is enabled for the access control policy. Valid values:
	//
	// 	- **on**: The strict mode is enabled.
	//
	// 	- **off**: The strict mode is disabled.
	//
	// example:
	//
	// off
	InternetSwitch *string `json:"InternetSwitch,omitempty" xml:"InternetSwitch,omitempty"`
	// The ID of the request.
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

func (s *DescribePolicyAdvancedConfigResponseBody) GetInternetSwitch() *string {
	return s.InternetSwitch
}

func (s *DescribePolicyAdvancedConfigResponseBody) GetRequestId() *string {
	return s.RequestId
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
