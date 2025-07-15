// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMiguLivePullToPushStartRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *MiguLivePullToPushStartRequest
	GetDomainName() *string
	SetMiguData(v string) *MiguLivePullToPushStartRequest
	GetMiguData() *string
	SetOwnerId(v int64) *MiguLivePullToPushStartRequest
	GetOwnerId() *int64
	SetRegionId(v string) *MiguLivePullToPushStartRequest
	GetRegionId() *string
}

type MiguLivePullToPushStartRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	MiguData *string `json:"MiguData,omitempty" xml:"MiguData,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MiguLivePullToPushStartRequest) String() string {
	return dara.Prettify(s)
}

func (s MiguLivePullToPushStartRequest) GoString() string {
	return s.String()
}

func (s *MiguLivePullToPushStartRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *MiguLivePullToPushStartRequest) GetMiguData() *string {
	return s.MiguData
}

func (s *MiguLivePullToPushStartRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MiguLivePullToPushStartRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MiguLivePullToPushStartRequest) SetDomainName(v string) *MiguLivePullToPushStartRequest {
	s.DomainName = &v
	return s
}

func (s *MiguLivePullToPushStartRequest) SetMiguData(v string) *MiguLivePullToPushStartRequest {
	s.MiguData = &v
	return s
}

func (s *MiguLivePullToPushStartRequest) SetOwnerId(v int64) *MiguLivePullToPushStartRequest {
	s.OwnerId = &v
	return s
}

func (s *MiguLivePullToPushStartRequest) SetRegionId(v string) *MiguLivePullToPushStartRequest {
	s.RegionId = &v
	return s
}

func (s *MiguLivePullToPushStartRequest) Validate() error {
	return dara.Validate(s)
}
