// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSetDcdnIpaDomainConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainNames(v string) *BatchSetDcdnIpaDomainConfigsRequest
	GetDomainNames() *string
	SetFunctions(v string) *BatchSetDcdnIpaDomainConfigsRequest
	GetFunctions() *string
	SetOwnerAccount(v string) *BatchSetDcdnIpaDomainConfigsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *BatchSetDcdnIpaDomainConfigsRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *BatchSetDcdnIpaDomainConfigsRequest
	GetSecurityToken() *string
}

type BatchSetDcdnIpaDomainConfigsRequest struct {
	// The domain names accelerated by IPA. Separate multiple domain names with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The list of features. `[{"functionArgs":[{"argName":"parameter name","argValue":"parameter value"}],"functionName":"feature name"}]`
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"functionArgs":[{"argName":"domain_name","argValue":"api.*com"}],"functionName":"protogw"}]
	Functions     *string `json:"Functions,omitempty" xml:"Functions,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s BatchSetDcdnIpaDomainConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSetDcdnIpaDomainConfigsRequest) GoString() string {
	return s.String()
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) GetDomainNames() *string {
	return s.DomainNames
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) GetFunctions() *string {
	return s.Functions
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetDomainNames(v string) *BatchSetDcdnIpaDomainConfigsRequest {
	s.DomainNames = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetFunctions(v string) *BatchSetDcdnIpaDomainConfigsRequest {
	s.Functions = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetOwnerAccount(v string) *BatchSetDcdnIpaDomainConfigsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetOwnerId(v int64) *BatchSetDcdnIpaDomainConfigsRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) SetSecurityToken(v string) *BatchSetDcdnIpaDomainConfigsRequest {
	s.SecurityToken = &v
	return s
}

func (s *BatchSetDcdnIpaDomainConfigsRequest) Validate() error {
	return dara.Validate(s)
}
