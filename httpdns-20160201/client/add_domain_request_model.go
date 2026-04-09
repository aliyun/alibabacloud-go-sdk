// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *AddDomainRequest
	GetAccountId() *string
	SetDomainName(v string) *AddDomainRequest
	GetDomainName() *string
}

type AddDomainRequest struct {
	// example:
	//
	// 123456
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s AddDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDomainRequest) GoString() string {
	return s.String()
}

func (s *AddDomainRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *AddDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDomainRequest) SetAccountId(v string) *AddDomainRequest {
	s.AccountId = &v
	return s
}

func (s *AddDomainRequest) SetDomainName(v string) *AddDomainRequest {
	s.DomainName = &v
	return s
}

func (s *AddDomainRequest) Validate() error {
	return dara.Validate(s)
}
