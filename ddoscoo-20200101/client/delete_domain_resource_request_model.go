// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDomainResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomain(v string) *DeleteDomainResourceRequest
	GetDomain() *string
}

type DeleteDomainResourceRequest struct {
	// The domain name corresponding to the website business forwarding rule to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s DeleteDomainResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDomainResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainResourceRequest) GetDomain() *string {
	return s.Domain
}

func (s *DeleteDomainResourceRequest) SetDomain(v string) *DeleteDomainResourceRequest {
	s.Domain = &v
	return s
}

func (s *DeleteDomainResourceRequest) Validate() error {
	return dara.Validate(s)
}
