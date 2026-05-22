// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *RevokeClientCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *RevokeClientCertificateResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *RevokeClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *RevokeClientCertificateResponseBody
	GetSiteName() *string
}

type RevokeClientCertificateResponseBody struct {
	// The certificate ID.
	//
	// example:
	//
	// baba39055622c008b90285a8838ed09a
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A123425345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s RevokeClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeClientCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *RevokeClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *RevokeClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *RevokeClientCertificateResponseBody) SetId(v string) *RevokeClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *RevokeClientCertificateResponseBody) SetRequestId(v string) *RevokeClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeClientCertificateResponseBody) SetSiteId(v int64) *RevokeClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *RevokeClientCertificateResponseBody) SetSiteName(v string) *RevokeClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *RevokeClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
