// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseDDoSInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDDoSBillingMode(v string) *PurchaseDDoSInstanceRequest
	GetDDoSBillingMode() *string
	SetDDoSBurstableDomesticProtection(v string) *PurchaseDDoSInstanceRequest
	GetDDoSBurstableDomesticProtection() *string
	SetDDoSBurstableOverseasProtection(v string) *PurchaseDDoSInstanceRequest
	GetDDoSBurstableOverseasProtection() *string
	SetSiteInstanceId(v string) *PurchaseDDoSInstanceRequest
	GetSiteInstanceId() *string
}

type PurchaseDDoSInstanceRequest struct {
	// The billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// CleanTraffic
	DDoSBillingMode *string `json:"DDoSBillingMode,omitempty" xml:"DDoSBillingMode,omitempty"`
	// The instance specifications for the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn_300
	DDoSBurstableDomesticProtection *string `json:"DDoSBurstableDomesticProtection,omitempty" xml:"DDoSBurstableDomesticProtection,omitempty"`
	// The instance specifications for outside China.
	//
	// This parameter is required.
	//
	// example:
	//
	// overseas_300
	DDoSBurstableOverseasProtection *string `json:"DDoSBurstableOverseasProtection,omitempty" xml:"DDoSBurstableOverseasProtection,omitempty"`
	// The site instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-site-23kde*****
	SiteInstanceId *string `json:"SiteInstanceId,omitempty" xml:"SiteInstanceId,omitempty"`
}

func (s PurchaseDDoSInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s PurchaseDDoSInstanceRequest) GoString() string {
	return s.String()
}

func (s *PurchaseDDoSInstanceRequest) GetDDoSBillingMode() *string {
	return s.DDoSBillingMode
}

func (s *PurchaseDDoSInstanceRequest) GetDDoSBurstableDomesticProtection() *string {
	return s.DDoSBurstableDomesticProtection
}

func (s *PurchaseDDoSInstanceRequest) GetDDoSBurstableOverseasProtection() *string {
	return s.DDoSBurstableOverseasProtection
}

func (s *PurchaseDDoSInstanceRequest) GetSiteInstanceId() *string {
	return s.SiteInstanceId
}

func (s *PurchaseDDoSInstanceRequest) SetDDoSBillingMode(v string) *PurchaseDDoSInstanceRequest {
	s.DDoSBillingMode = &v
	return s
}

func (s *PurchaseDDoSInstanceRequest) SetDDoSBurstableDomesticProtection(v string) *PurchaseDDoSInstanceRequest {
	s.DDoSBurstableDomesticProtection = &v
	return s
}

func (s *PurchaseDDoSInstanceRequest) SetDDoSBurstableOverseasProtection(v string) *PurchaseDDoSInstanceRequest {
	s.DDoSBurstableOverseasProtection = &v
	return s
}

func (s *PurchaseDDoSInstanceRequest) SetSiteInstanceId(v string) *PurchaseDDoSInstanceRequest {
	s.SiteInstanceId = &v
	return s
}

func (s *PurchaseDDoSInstanceRequest) Validate() error {
	return dara.Validate(s)
}
