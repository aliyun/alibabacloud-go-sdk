// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnIpaDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateDcdnIpaDomainRequest
	GetDomainName() *string
	SetOwnerId(v int64) *UpdateDcdnIpaDomainRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *UpdateDcdnIpaDomainRequest
	GetResourceGroupId() *string
	SetSecurityToken(v string) *UpdateDcdnIpaDomainRequest
	GetSecurityToken() *string
	SetSources(v string) *UpdateDcdnIpaDomainRequest
	GetSources() *string
	SetTopLevelDomain(v string) *UpdateDcdnIpaDomainRequest
	GetTopLevelDomain() *string
}

type UpdateDcdnIpaDomainRequest struct {
	// The accelerated domain name that you want to modify. You can specify only one domain name in each request.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmyuji4b6r4**
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SecurityToken   *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The information about the addresses of the origin server.
	//
	// example:
	//
	// [{"content":"10.10.10.10","type":"ipaddr","priority":"20","port":80,"weight":"15"}]
	Sources *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The top-level domain name.
	//
	// example:
	//
	// example.edu
	TopLevelDomain *string `json:"TopLevelDomain,omitempty" xml:"TopLevelDomain,omitempty"`
}

func (s UpdateDcdnIpaDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnIpaDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnIpaDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDcdnIpaDomainRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateDcdnIpaDomainRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *UpdateDcdnIpaDomainRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *UpdateDcdnIpaDomainRequest) GetSources() *string {
	return s.Sources
}

func (s *UpdateDcdnIpaDomainRequest) GetTopLevelDomain() *string {
	return s.TopLevelDomain
}

func (s *UpdateDcdnIpaDomainRequest) SetDomainName(v string) *UpdateDcdnIpaDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetOwnerId(v int64) *UpdateDcdnIpaDomainRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetResourceGroupId(v string) *UpdateDcdnIpaDomainRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetSecurityToken(v string) *UpdateDcdnIpaDomainRequest {
	s.SecurityToken = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetSources(v string) *UpdateDcdnIpaDomainRequest {
	s.Sources = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) SetTopLevelDomain(v string) *UpdateDcdnIpaDomainRequest {
	s.TopLevelDomain = &v
	return s
}

func (s *UpdateDcdnIpaDomainRequest) Validate() error {
	return dara.Validate(s)
}
