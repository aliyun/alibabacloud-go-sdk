// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteOriginClientCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteOriginClientCertificateResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DeleteOriginClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *DeleteOriginClientCertificateResponseBody
	GetSiteName() *string
}

type DeleteOriginClientCertificateResponseBody struct {
	// The certificate ID.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s DeleteOriginClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOriginClientCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteOriginClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOriginClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteOriginClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *DeleteOriginClientCertificateResponseBody) SetId(v string) *DeleteOriginClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteOriginClientCertificateResponseBody) SetRequestId(v string) *DeleteOriginClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOriginClientCertificateResponseBody) SetSiteId(v int64) *DeleteOriginClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *DeleteOriginClientCertificateResponseBody) SetSiteName(v string) *DeleteOriginClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *DeleteOriginClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
