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
	// The domain name for which you want to modify the acceleration region.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The acceleration region that you want to set. {"coverage":"overseas"} specifies regions outside the Chinese mainland. Valid values of coverage:
	//
	// 	- domestic: regions in the Chinese mainland.
	//
	// 	- overseas: regions outside the Chinese mainland.
	//
	// 	- global: regions in and outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"coverage":"global"}
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
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
