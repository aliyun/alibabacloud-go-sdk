// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMainDomainNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteName(v string) *GetMainDomainNameRequest
	GetSiteName() *string
}

type GetMainDomainNameRequest struct {
	// The website name.
	//
	// This parameter is required.
	//
	// example:
	//
	// sub.example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s GetMainDomainNameRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMainDomainNameRequest) GoString() string {
	return s.String()
}

func (s *GetMainDomainNameRequest) GetSiteName() *string {
	return s.SiteName
}

func (s *GetMainDomainNameRequest) SetSiteName(v string) *GetMainDomainNameRequest {
	s.SiteName = &v
	return s
}

func (s *GetMainDomainNameRequest) Validate() error {
	return dara.Validate(s)
}
