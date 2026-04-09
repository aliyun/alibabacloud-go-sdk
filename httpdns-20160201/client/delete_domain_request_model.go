// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountId(v string) *DeleteDomainRequest
	GetAccountId() *string
	SetDomainName(v string) *DeleteDomainRequest
	GetDomainName() *string
}

type DeleteDomainRequest struct {
	// example:
	//
	// 12****
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) GetAccountId() *string {
	return s.AccountId
}

func (s *DeleteDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteDomainRequest) SetAccountId(v string) *DeleteDomainRequest {
	s.AccountId = &v
	return s
}

func (s *DeleteDomainRequest) SetDomainName(v string) *DeleteDomainRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteDomainRequest) Validate() error {
	return dara.Validate(s)
}
