// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpApiDomainInfo interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *HttpApiDomainInfo
	GetDomainId() *string
	SetName(v string) *HttpApiDomainInfo
	GetName() *string
	SetProtocol(v string) *HttpApiDomainInfo
	GetProtocol() *string
}

type HttpApiDomainInfo struct {
	// The domain name ID.
	//
	// example:
	//
	// d-xxx
	DomainId *string `json:"domainId,omitempty" xml:"domainId,omitempty"`
	// The domain name.
	//
	// example:
	//
	// www.example.com
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The protocol.
	//
	// Valid values:
	//
	// 	- HTTPS
	//
	// 	- HTTP
	//
	// example:
	//
	// HTTP
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
}

func (s HttpApiDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s HttpApiDomainInfo) GoString() string {
	return s.String()
}

func (s *HttpApiDomainInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *HttpApiDomainInfo) GetName() *string {
	return s.Name
}

func (s *HttpApiDomainInfo) GetProtocol() *string {
	return s.Protocol
}

func (s *HttpApiDomainInfo) SetDomainId(v string) *HttpApiDomainInfo {
	s.DomainId = &v
	return s
}

func (s *HttpApiDomainInfo) SetName(v string) *HttpApiDomainInfo {
	s.Name = &v
	return s
}

func (s *HttpApiDomainInfo) SetProtocol(v string) *HttpApiDomainInfo {
	s.Protocol = &v
	return s
}

func (s *HttpApiDomainInfo) Validate() error {
	return dara.Validate(s)
}
