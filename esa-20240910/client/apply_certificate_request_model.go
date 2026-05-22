// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomains(v string) *ApplyCertificateRequest
	GetDomains() *string
	SetSiteId(v int64) *ApplyCertificateRequest
	GetSiteId() *int64
	SetType(v string) *ApplyCertificateRequest
	GetType() *string
}

type ApplyCertificateRequest struct {
	// This parameter is required.
	Domains *string `json:"Domains,omitempty" xml:"Domains,omitempty"`
	// This parameter is required.
	SiteId *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ApplyCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyCertificateRequest) GoString() string {
	return s.String()
}

func (s *ApplyCertificateRequest) GetDomains() *string {
	return s.Domains
}

func (s *ApplyCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ApplyCertificateRequest) GetType() *string {
	return s.Type
}

func (s *ApplyCertificateRequest) SetDomains(v string) *ApplyCertificateRequest {
	s.Domains = &v
	return s
}

func (s *ApplyCertificateRequest) SetSiteId(v int64) *ApplyCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *ApplyCertificateRequest) SetType(v string) *ApplyCertificateRequest {
	s.Type = &v
	return s
}

func (s *ApplyCertificateRequest) Validate() error {
	return dara.Validate(s)
}
