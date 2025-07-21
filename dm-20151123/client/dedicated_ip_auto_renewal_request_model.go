// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpAutoRenewalRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRenewal(v string) *DedicatedIpAutoRenewalRequest
	GetAutoRenewal() *string
	SetBuyResourceIds(v string) *DedicatedIpAutoRenewalRequest
	GetBuyResourceIds() *string
}

type DedicatedIpAutoRenewalRequest struct {
	// Whether to enable auto-renewal
	//
	// This parameter is required.
	//
	// example:
	//
	// false
	AutoRenewal *string `json:"AutoRenewal,omitempty" xml:"AutoRenewal,omitempty"`
	// Purchase instance ID, separated by English commas if multiple.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx,xxx
	BuyResourceIds *string `json:"BuyResourceIds,omitempty" xml:"BuyResourceIds,omitempty"`
}

func (s DedicatedIpAutoRenewalRequest) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpAutoRenewalRequest) GoString() string {
	return s.String()
}

func (s *DedicatedIpAutoRenewalRequest) GetAutoRenewal() *string {
	return s.AutoRenewal
}

func (s *DedicatedIpAutoRenewalRequest) GetBuyResourceIds() *string {
	return s.BuyResourceIds
}

func (s *DedicatedIpAutoRenewalRequest) SetAutoRenewal(v string) *DedicatedIpAutoRenewalRequest {
	s.AutoRenewal = &v
	return s
}

func (s *DedicatedIpAutoRenewalRequest) SetBuyResourceIds(v string) *DedicatedIpAutoRenewalRequest {
	s.BuyResourceIds = &v
	return s
}

func (s *DedicatedIpAutoRenewalRequest) Validate() error {
	return dara.Validate(s)
}
