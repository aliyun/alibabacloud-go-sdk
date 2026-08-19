// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDDoSSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDDoSBillingMode(v string) *UpdateDDoSSpecRequest
	GetDDoSBillingMode() *string
	SetDDoSBurstableDomesticProtection(v string) *UpdateDDoSSpecRequest
	GetDDoSBurstableDomesticProtection() *string
	SetDDoSBurstableOverseasProtection(v string) *UpdateDDoSSpecRequest
	GetDDoSBurstableOverseasProtection() *string
	SetInstanceId(v string) *UpdateDDoSSpecRequest
	GetInstanceId() *string
}

type UpdateDDoSSpecRequest struct {
	// The billing method.
	//
	// This parameter is required.
	//
	// example:
	//
	// CleanTraffic
	DDoSBillingMode *string `json:"DDoSBillingMode,omitempty" xml:"DDoSBillingMode,omitempty"`
	// The instance specification for the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn_300
	DDoSBurstableDomesticProtection *string `json:"DDoSBurstableDomesticProtection,omitempty" xml:"DDoSBurstableDomesticProtection,omitempty"`
	// The instance specification for regions outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// overseas_300
	DDoSBurstableOverseasProtection *string `json:"DDoSBurstableOverseasProtection,omitempty" xml:"DDoSBurstableOverseasProtection,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-ddos-9tuv*********
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateDDoSSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDDoSSpecRequest) GoString() string {
	return s.String()
}

func (s *UpdateDDoSSpecRequest) GetDDoSBillingMode() *string {
	return s.DDoSBillingMode
}

func (s *UpdateDDoSSpecRequest) GetDDoSBurstableDomesticProtection() *string {
	return s.DDoSBurstableDomesticProtection
}

func (s *UpdateDDoSSpecRequest) GetDDoSBurstableOverseasProtection() *string {
	return s.DDoSBurstableOverseasProtection
}

func (s *UpdateDDoSSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateDDoSSpecRequest) SetDDoSBillingMode(v string) *UpdateDDoSSpecRequest {
	s.DDoSBillingMode = &v
	return s
}

func (s *UpdateDDoSSpecRequest) SetDDoSBurstableDomesticProtection(v string) *UpdateDDoSSpecRequest {
	s.DDoSBurstableDomesticProtection = &v
	return s
}

func (s *UpdateDDoSSpecRequest) SetDDoSBurstableOverseasProtection(v string) *UpdateDDoSSpecRequest {
	s.DDoSBurstableOverseasProtection = &v
	return s
}

func (s *UpdateDDoSSpecRequest) SetInstanceId(v string) *UpdateDDoSSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateDDoSSpecRequest) Validate() error {
	return dara.Validate(s)
}
