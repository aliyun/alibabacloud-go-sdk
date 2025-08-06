// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnycastEipAddressSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAnycastId(v string) *ModifyAnycastEipAddressSpecRequest
	GetAnycastId() *string
	SetBandwidth(v string) *ModifyAnycastEipAddressSpecRequest
	GetBandwidth() *string
}

type ModifyAnycastEipAddressSpecRequest struct {
	// The ID of the Anycast EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// aeip-bp1ix34fralt4ykf3****
	AnycastId *string `json:"AnycastId,omitempty" xml:"AnycastId,omitempty"`
	// The maximum bandwidth of the Anycast EIP. Unit: Mbit/s.
	//
	// Valid values: **200*	- to **1000**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 200
	Bandwidth *string `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
}

func (s ModifyAnycastEipAddressSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnycastEipAddressSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressSpecRequest) GetAnycastId() *string {
	return s.AnycastId
}

func (s *ModifyAnycastEipAddressSpecRequest) GetBandwidth() *string {
	return s.Bandwidth
}

func (s *ModifyAnycastEipAddressSpecRequest) SetAnycastId(v string) *ModifyAnycastEipAddressSpecRequest {
	s.AnycastId = &v
	return s
}

func (s *ModifyAnycastEipAddressSpecRequest) SetBandwidth(v string) *ModifyAnycastEipAddressSpecRequest {
	s.Bandwidth = &v
	return s
}

func (s *ModifyAnycastEipAddressSpecRequest) Validate() error {
	return dara.Validate(s)
}
