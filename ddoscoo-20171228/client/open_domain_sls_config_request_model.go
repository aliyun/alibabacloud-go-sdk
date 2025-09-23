// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenDomainSlsConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *OpenDomainSlsConfigRequest
	GetDomain() *string
	SetLang(v string) *OpenDomainSlsConfigRequest
	GetLang() *string
	SetResourceGroupId(v string) *OpenDomainSlsConfigRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *OpenDomainSlsConfigRequest
	GetSourceIp() *string
}

type OpenDomainSlsConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// cn
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// xx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// 1.1.1.1
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s OpenDomainSlsConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenDomainSlsConfigRequest) GoString() string {
	return s.String()
}

func (s *OpenDomainSlsConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *OpenDomainSlsConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *OpenDomainSlsConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *OpenDomainSlsConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *OpenDomainSlsConfigRequest) SetDomain(v string) *OpenDomainSlsConfigRequest {
	s.Domain = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetLang(v string) *OpenDomainSlsConfigRequest {
	s.Lang = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetResourceGroupId(v string) *OpenDomainSlsConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) SetSourceIp(v string) *OpenDomainSlsConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *OpenDomainSlsConfigRequest) Validate() error {
	return dara.Validate(s)
}
