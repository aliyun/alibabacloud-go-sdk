// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCdnMigrateRegisterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CdnMigrateRegisterRequest
	GetDomainName() *string
}

type CdnMigrateRegisterRequest struct {
	// The accelerated domain name for which you want to register the dynamic routing feature. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s CdnMigrateRegisterRequest) String() string {
	return dara.Prettify(s)
}

func (s CdnMigrateRegisterRequest) GoString() string {
	return s.String()
}

func (s *CdnMigrateRegisterRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *CdnMigrateRegisterRequest) SetDomainName(v string) *CdnMigrateRegisterRequest {
	s.DomainName = &v
	return s
}

func (s *CdnMigrateRegisterRequest) Validate() error {
	return dara.Validate(s)
}
