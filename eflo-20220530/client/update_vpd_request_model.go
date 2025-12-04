// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVpdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *UpdateVpdRequest
	GetRegionId() *string
	SetVpdId(v string) *UpdateVpdRequest
	GetVpdId() *string
	SetVpdName(v string) *UpdateVpdRequest
	GetVpdName() *string
}

type UpdateVpdRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the VPD instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpd-omqutbff
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The name of the VPD instance.
	//
	// example:
	//
	// vpd-lingjun
	VpdName *string `json:"VpdName,omitempty" xml:"VpdName,omitempty"`
}

func (s UpdateVpdRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateVpdRequest) GoString() string {
	return s.String()
}

func (s *UpdateVpdRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateVpdRequest) GetVpdId() *string {
	return s.VpdId
}

func (s *UpdateVpdRequest) GetVpdName() *string {
	return s.VpdName
}

func (s *UpdateVpdRequest) SetRegionId(v string) *UpdateVpdRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateVpdRequest) SetVpdId(v string) *UpdateVpdRequest {
	s.VpdId = &v
	return s
}

func (s *UpdateVpdRequest) SetVpdName(v string) *UpdateVpdRequest {
	s.VpdName = &v
	return s
}

func (s *UpdateVpdRequest) Validate() error {
	return dara.Validate(s)
}
