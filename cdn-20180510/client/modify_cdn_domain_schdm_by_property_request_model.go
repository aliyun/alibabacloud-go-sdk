// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdnDomainSchdmByPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyCdnDomainSchdmByPropertyRequest
	GetDomainName() *string
	SetProperty(v string) *ModifyCdnDomainSchdmByPropertyRequest
	GetProperty() *string
}

type ModifyCdnDomainSchdmByPropertyRequest struct {
	// The accelerated domain name for which you want to change the acceleration region. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The information about the acceleration region. {"coverage":"overseas"}
	//
	// Valid values for coverage:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: global (excluding the Chinese mainland)
	//
	// 	- **global**: global
	//
	// This parameter is required.
	//
	// example:
	//
	// {"coverage":"overseas"}
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
}

func (s ModifyCdnDomainSchdmByPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdnDomainSchdmByPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyCdnDomainSchdmByPropertyRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyCdnDomainSchdmByPropertyRequest) GetProperty() *string {
	return s.Property
}

func (s *ModifyCdnDomainSchdmByPropertyRequest) SetDomainName(v string) *ModifyCdnDomainSchdmByPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyRequest) SetProperty(v string) *ModifyCdnDomainSchdmByPropertyRequest {
	s.Property = &v
	return s
}

func (s *ModifyCdnDomainSchdmByPropertyRequest) Validate() error {
	return dara.Validate(s)
}
