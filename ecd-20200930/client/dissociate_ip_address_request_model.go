// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipId(v string) *DissociateIpAddressRequest
	GetEipId() *string
	SetRegionId(v string) *DissociateIpAddressRequest
	GetRegionId() *string
}

type DissociateIpAddressRequest struct {
	// This parameter is required.
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DissociateIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateIpAddressRequest) GoString() string {
	return s.String()
}

func (s *DissociateIpAddressRequest) GetEipId() *string {
	return s.EipId
}

func (s *DissociateIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DissociateIpAddressRequest) SetEipId(v string) *DissociateIpAddressRequest {
	s.EipId = &v
	return s
}

func (s *DissociateIpAddressRequest) SetRegionId(v string) *DissociateIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
