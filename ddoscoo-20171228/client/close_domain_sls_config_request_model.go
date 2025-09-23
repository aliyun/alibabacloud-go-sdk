// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDomainSlsConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *CloseDomainSlsConfigRequest
	GetDomain() *string
	SetLang(v string) *CloseDomainSlsConfigRequest
	GetLang() *string
	SetResourceGroupId(v string) *CloseDomainSlsConfigRequest
	GetResourceGroupId() *string
	SetSourceIp(v string) *CloseDomainSlsConfigRequest
	GetSourceIp() *string
}

type CloseDomainSlsConfigRequest struct {
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

func (s CloseDomainSlsConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s CloseDomainSlsConfigRequest) GoString() string {
	return s.String()
}

func (s *CloseDomainSlsConfigRequest) GetDomain() *string {
	return s.Domain
}

func (s *CloseDomainSlsConfigRequest) GetLang() *string {
	return s.Lang
}

func (s *CloseDomainSlsConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CloseDomainSlsConfigRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CloseDomainSlsConfigRequest) SetDomain(v string) *CloseDomainSlsConfigRequest {
	s.Domain = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetLang(v string) *CloseDomainSlsConfigRequest {
	s.Lang = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetResourceGroupId(v string) *CloseDomainSlsConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) SetSourceIp(v string) *CloseDomainSlsConfigRequest {
	s.SourceIp = &v
	return s
}

func (s *CloseDomainSlsConfigRequest) Validate() error {
	return dara.Validate(s)
}
