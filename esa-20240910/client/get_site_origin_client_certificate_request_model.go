// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteOriginClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetSiteOriginClientCertificateRequest
	GetId() *string
	SetSiteId(v int64) *GetSiteOriginClientCertificateRequest
	GetSiteId() *int64
}

type GetSiteOriginClientCertificateRequest struct {
	// The ID of the client certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the site.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetSiteOriginClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSiteOriginClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetSiteOriginClientCertificateRequest) GetId() *string {
	return s.Id
}

func (s *GetSiteOriginClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetSiteOriginClientCertificateRequest) SetId(v string) *GetSiteOriginClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *GetSiteOriginClientCertificateRequest) SetSiteId(v int64) *GetSiteOriginClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *GetSiteOriginClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
