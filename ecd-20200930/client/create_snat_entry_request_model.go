// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnatEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEipAffinity(v int32) *CreateSnatEntryRequest
	GetEipAffinity() *int32
	SetRegionId(v string) *CreateSnatEntryRequest
	GetRegionId() *string
	SetSnatEntryName(v string) *CreateSnatEntryRequest
	GetSnatEntryName() *string
	SetSnatIp(v string) *CreateSnatEntryRequest
	GetSnatIp() *string
	SetSnatTableId(v string) *CreateSnatEntryRequest
	GetSnatTableId() *string
	SetSourceCIDR(v string) *CreateSnatEntryRequest
	GetSourceCIDR() *string
}

type CreateSnatEntryRequest struct {
	EipAffinity *int32 `json:"EipAffinity,omitempty" xml:"EipAffinity,omitempty"`
	// This parameter is required.
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	SnatEntryName *string `json:"SnatEntryName,omitempty" xml:"SnatEntryName,omitempty"`
	// This parameter is required.
	SnatIp *string `json:"SnatIp,omitempty" xml:"SnatIp,omitempty"`
	// This parameter is required.
	SnatTableId *string `json:"SnatTableId,omitempty" xml:"SnatTableId,omitempty"`
	// This parameter is required.
	SourceCIDR *string `json:"SourceCIDR,omitempty" xml:"SourceCIDR,omitempty"`
}

func (s CreateSnatEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnatEntryRequest) GoString() string {
	return s.String()
}

func (s *CreateSnatEntryRequest) GetEipAffinity() *int32 {
	return s.EipAffinity
}

func (s *CreateSnatEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSnatEntryRequest) GetSnatEntryName() *string {
	return s.SnatEntryName
}

func (s *CreateSnatEntryRequest) GetSnatIp() *string {
	return s.SnatIp
}

func (s *CreateSnatEntryRequest) GetSnatTableId() *string {
	return s.SnatTableId
}

func (s *CreateSnatEntryRequest) GetSourceCIDR() *string {
	return s.SourceCIDR
}

func (s *CreateSnatEntryRequest) SetEipAffinity(v int32) *CreateSnatEntryRequest {
	s.EipAffinity = &v
	return s
}

func (s *CreateSnatEntryRequest) SetRegionId(v string) *CreateSnatEntryRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSnatEntryName(v string) *CreateSnatEntryRequest {
	s.SnatEntryName = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSnatIp(v string) *CreateSnatEntryRequest {
	s.SnatIp = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSnatTableId(v string) *CreateSnatEntryRequest {
	s.SnatTableId = &v
	return s
}

func (s *CreateSnatEntryRequest) SetSourceCIDR(v string) *CreateSnatEntryRequest {
	s.SourceCIDR = &v
	return s
}

func (s *CreateSnatEntryRequest) Validate() error {
	return dara.Validate(s)
}
