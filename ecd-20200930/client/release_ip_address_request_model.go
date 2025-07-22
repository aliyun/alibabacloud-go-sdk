// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseIpAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipId(v string) *ReleaseIpAddressRequest
	GetEipId() *string
	SetRegionId(v string) *ReleaseIpAddressRequest
	GetRegionId() *string
}

type ReleaseIpAddressRequest struct {
	// This parameter is required.
	EipId *string `json:"EipId,omitempty" xml:"EipId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ReleaseIpAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseIpAddressRequest) GoString() string {
	return s.String()
}

func (s *ReleaseIpAddressRequest) GetEipId() *string {
	return s.EipId
}

func (s *ReleaseIpAddressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReleaseIpAddressRequest) SetEipId(v string) *ReleaseIpAddressRequest {
	s.EipId = &v
	return s
}

func (s *ReleaseIpAddressRequest) SetRegionId(v string) *ReleaseIpAddressRequest {
	s.RegionId = &v
	return s
}

func (s *ReleaseIpAddressRequest) Validate() error {
	return dara.Validate(s)
}
