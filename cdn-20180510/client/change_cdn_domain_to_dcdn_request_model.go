// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeCdnDomainToDcdnRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ChangeCdnDomainToDcdnRequest
	GetDomainName() *string
	SetOperation(v string) *ChangeCdnDomainToDcdnRequest
	GetOperation() *string
	SetOwnerAccount(v string) *ChangeCdnDomainToDcdnRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ChangeCdnDomainToDcdnRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *ChangeCdnDomainToDcdnRequest
	GetSecurityToken() *string
}

type ChangeCdnDomainToDcdnRequest struct {
	// The accelerated domain name. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The operation to perform. Set the value to preCheck. Precheck is performed, and the result is returned. If the precheck passes, set the value to enforce to perform the transfer.
	//
	// example:
	//
	// preCheck
	Operation     *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ChangeCdnDomainToDcdnRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeCdnDomainToDcdnRequest) GoString() string {
	return s.String()
}

func (s *ChangeCdnDomainToDcdnRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ChangeCdnDomainToDcdnRequest) GetOperation() *string {
	return s.Operation
}

func (s *ChangeCdnDomainToDcdnRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ChangeCdnDomainToDcdnRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ChangeCdnDomainToDcdnRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ChangeCdnDomainToDcdnRequest) SetDomainName(v string) *ChangeCdnDomainToDcdnRequest {
	s.DomainName = &v
	return s
}

func (s *ChangeCdnDomainToDcdnRequest) SetOperation(v string) *ChangeCdnDomainToDcdnRequest {
	s.Operation = &v
	return s
}

func (s *ChangeCdnDomainToDcdnRequest) SetOwnerAccount(v string) *ChangeCdnDomainToDcdnRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ChangeCdnDomainToDcdnRequest) SetOwnerId(v int64) *ChangeCdnDomainToDcdnRequest {
	s.OwnerId = &v
	return s
}

func (s *ChangeCdnDomainToDcdnRequest) SetSecurityToken(v string) *ChangeCdnDomainToDcdnRequest {
	s.SecurityToken = &v
	return s
}

func (s *ChangeCdnDomainToDcdnRequest) Validate() error {
	return dara.Validate(s)
}
