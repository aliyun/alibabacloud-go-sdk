// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelApiDomainInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDefault(v bool) *ModelApiDomainInfo
	GetDefault() *bool
	SetDomainId(v string) *ModelApiDomainInfo
	GetDomainId() *string
	SetDomainScope(v string) *ModelApiDomainInfo
	GetDomainScope() *string
	SetName(v string) *ModelApiDomainInfo
	GetName() *string
	SetProtocol(v string) *ModelApiDomainInfo
	GetProtocol() *string
}

type ModelApiDomainInfo struct {
	// Indicates whether it is the default domain name of the instance.
	//
	// example:
	//
	// false
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
	// The domain name ID.
	//
	// example:
	//
	// d-d00000000000000000000
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name scope.
	//
	// example:
	//
	// Dedicated
	DomainScope *string `json:"domainScope,omitempty" xml:"domainScope,omitempty"`
	// The domain name.
	//
	// example:
	//
	// api.example.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The access protocol.
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s ModelApiDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s ModelApiDomainInfo) GoString() string {
	return s.String()
}

func (s *ModelApiDomainInfo) GetDefault() *bool {
	return s.Default
}

func (s *ModelApiDomainInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *ModelApiDomainInfo) GetDomainScope() *string {
	return s.DomainScope
}

func (s *ModelApiDomainInfo) GetName() *string {
	return s.Name
}

func (s *ModelApiDomainInfo) GetProtocol() *string {
	return s.Protocol
}

func (s *ModelApiDomainInfo) SetDefault(v bool) *ModelApiDomainInfo {
	s.Default = &v
	return s
}

func (s *ModelApiDomainInfo) SetDomainId(v string) *ModelApiDomainInfo {
	s.DomainId = &v
	return s
}

func (s *ModelApiDomainInfo) SetDomainScope(v string) *ModelApiDomainInfo {
	s.DomainScope = &v
	return s
}

func (s *ModelApiDomainInfo) SetName(v string) *ModelApiDomainInfo {
	s.Name = &v
	return s
}

func (s *ModelApiDomainInfo) SetProtocol(v string) *ModelApiDomainInfo {
	s.Protocol = &v
	return s
}

func (s *ModelApiDomainInfo) Validate() error {
	return dara.Validate(s)
}
