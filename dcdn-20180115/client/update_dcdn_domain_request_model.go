// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateDcdnDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *UpdateDcdnDomainRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *UpdateDcdnDomainRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *UpdateDcdnDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *UpdateDcdnDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *UpdateDcdnDomainRequest
	GetTopLevelDomain() *string
}

type UpdateDcdnDomainRequest struct {
	// The accelerated domain name. You can specify only one domain name in each call.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The information about the addresses of origin servers.
	//
	// example:
	//
	// [{"content":"10.10.10.10","type":"ipaddr","priority":"20","port":80}]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The top-level domain.
	//
	// example:
	//
	// yourTopLevelDomain
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s UpdateDcdnDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDcdnDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateDcdnDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateDcdnDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateDcdnDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *UpdateDcdnDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *UpdateDcdnDomainRequest) SetDomainName(v string) *UpdateDcdnDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetOwnerId(v int64) *UpdateDcdnDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetResourceGroupId(v string) *UpdateDcdnDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetSecurityToken(v string) *UpdateDcdnDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetSources(v string) *UpdateDcdnDomainRequest {
	s.Sources = &v
	return s
}

func (s *UpdateDcdnDomainRequest) SetTopLevelDomain(v string) *UpdateDcdnDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *UpdateDcdnDomainRequest) Validate() error {
	return dara.Validate(s)
}
