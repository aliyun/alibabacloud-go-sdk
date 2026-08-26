// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLiveDomainSchdmByPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyLiveDomainSchdmByPropertyRequest
	GetDomainName() *string
	SetOwnerId(v int64) *ModifyLiveDomainSchdmByPropertyRequest
	GetOwnerId() *int64
	SetProperty(v string) *ModifyLiveDomainSchdmByPropertyRequest
	GetProperty() *string
	SetRegionId(v string) *ModifyLiveDomainSchdmByPropertyRequest
	GetRegionId() *string
}

type ModifyLiveDomainSchdmByPropertyRequest struct {
	// The live streaming domain for which you want to modify the acceleration region.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The acceleration region. A value of {"coverage":"overseas"} specifies that the configuration is for regions outside mainland China. The following list describes the valid values for coverage:
	//
	// - domestic: mainland China.
	//
	// - overseas: regions outside mainland China.
	//
	// - global: regions in and outside mainland China.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"coverage":"global"}
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyLiveDomainSchdmByPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLiveDomainSchdmByPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) GetProperty() *string {
	return s.Property
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) SetDomainName(v string) *ModifyLiveDomainSchdmByPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) SetOwnerId(v int64) *ModifyLiveDomainSchdmByPropertyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) SetProperty(v string) *ModifyLiveDomainSchdmByPropertyRequest {
	s.Property = &v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) SetRegionId(v string) *ModifyLiveDomainSchdmByPropertyRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyLiveDomainSchdmByPropertyRequest) Validate() error {
	return dara.Validate(s)
}
