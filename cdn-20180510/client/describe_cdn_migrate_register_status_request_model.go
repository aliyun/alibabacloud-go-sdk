// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnMigrateRegisterStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnMigrateRegisterStatusRequest
	GetDomainName() *string
}

type DescribeCdnMigrateRegisterStatusRequest struct {
	// The accelerated domain name. You can specify only one domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DescribeCdnMigrateRegisterStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnMigrateRegisterStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCdnMigrateRegisterStatusRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnMigrateRegisterStatusRequest) SetDomainName(v string) *DescribeCdnMigrateRegisterStatusRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnMigrateRegisterStatusRequest) Validate() error {
	return dara.Validate(s)
}
