// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteOriginClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteSiteOriginClientCertificateRequest
	GetId() *string
	SetSiteId(v int64) *DeleteSiteOriginClientCertificateRequest
	GetSiteId() *int64
}

type DeleteSiteOriginClientCertificateRequest struct {
	// The certificate ID on ESA.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteSiteOriginClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteOriginClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteSiteOriginClientCertificateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteSiteOriginClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteSiteOriginClientCertificateRequest) SetId(v string) *DeleteSiteOriginClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *DeleteSiteOriginClientCertificateRequest) SetSiteId(v int64) *DeleteSiteOriginClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteSiteOriginClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
