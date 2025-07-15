// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMiguLivePullToPushStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *MiguLivePullToPushStatusRequest
	GetDomainName() *string
	SetMiguData(v string) *MiguLivePullToPushStatusRequest
	GetMiguData() *string
	SetOwnerId(v int64) *MiguLivePullToPushStatusRequest
	GetOwnerId() *int64
	SetRegionId(v string) *MiguLivePullToPushStatusRequest
	GetRegionId() *string
}

type MiguLivePullToPushStatusRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	MiguData *string `json:"MiguData,omitempty" xml:"MiguData,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s MiguLivePullToPushStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s MiguLivePullToPushStatusRequest) GoString() string {
	return s.String()
}

func (s *MiguLivePullToPushStatusRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *MiguLivePullToPushStatusRequest) GetMiguData() *string {
	return s.MiguData
}

func (s *MiguLivePullToPushStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *MiguLivePullToPushStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *MiguLivePullToPushStatusRequest) SetDomainName(v string) *MiguLivePullToPushStatusRequest {
	s.DomainName = &v
	return s
}

func (s *MiguLivePullToPushStatusRequest) SetMiguData(v string) *MiguLivePullToPushStatusRequest {
	s.MiguData = &v
	return s
}

func (s *MiguLivePullToPushStatusRequest) SetOwnerId(v int64) *MiguLivePullToPushStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *MiguLivePullToPushStatusRequest) SetRegionId(v string) *MiguLivePullToPushStatusRequest {
	s.RegionId = &v
	return s
}

func (s *MiguLivePullToPushStatusRequest) Validate() error {
	return dara.Validate(s)
}
