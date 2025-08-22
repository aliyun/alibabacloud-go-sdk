// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDCdnDomainSchdmByPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyDCdnDomainSchdmByPropertyRequest
	GetDomainName() *string
	SetProperty(v string) *ModifyDCdnDomainSchdmByPropertyRequest
	GetProperty() *string
}

type ModifyDCdnDomainSchdmByPropertyRequest struct {
	// The name of the accelerated domain for which you want to change the acceleration region. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The region where the acceleration service is deployed. Valid values:
	//
	// 	- **domestic**: Chinese mainland
	//
	// 	- **overseas**: global (excluding mainland China)
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

func (s ModifyDCdnDomainSchdmByPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDCdnDomainSchdmByPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyDCdnDomainSchdmByPropertyRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyDCdnDomainSchdmByPropertyRequest) GetProperty() *string {
	return s.Property
}

func (s *ModifyDCdnDomainSchdmByPropertyRequest) SetDomainName(v string) *ModifyDCdnDomainSchdmByPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyRequest) SetProperty(v string) *ModifyDCdnDomainSchdmByPropertyRequest {
	s.Property = &v
	return s
}

func (s *ModifyDCdnDomainSchdmByPropertyRequest) Validate() error {
	return dara.Validate(s)
}
