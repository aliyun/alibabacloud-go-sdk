// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainVerifyDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainVerifyDataRequest
	GetDomainName() *string
	SetGlobalResourcePlan(v string) *DescribeDomainVerifyDataRequest
	GetGlobalResourcePlan() *string
}

type DescribeDomainVerifyDataRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.yourdomain.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Specifies whether to enable the global resource plan.
	//
	// Valid values:
	//
	// 	- off
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- on
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// off
	GlobalResourcePlan *string `json:"GlobalResourcePlan,omitempty" xml:"GlobalResourcePlan,omitempty"`
}

func (s DescribeDomainVerifyDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainVerifyDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainVerifyDataRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainVerifyDataRequest) GetGlobalResourcePlan() *string {
	return s.GlobalResourcePlan
}

func (s *DescribeDomainVerifyDataRequest) SetDomainName(v string) *DescribeDomainVerifyDataRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainVerifyDataRequest) SetGlobalResourcePlan(v string) *DescribeDomainVerifyDataRequest {
	s.GlobalResourcePlan = &v
	return s
}

func (s *DescribeDomainVerifyDataRequest) Validate() error {
	return dara.Validate(s)
}
