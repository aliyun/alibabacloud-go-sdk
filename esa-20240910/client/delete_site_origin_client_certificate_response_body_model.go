// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSiteOriginClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteSiteOriginClientCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteSiteOriginClientCertificateResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DeleteSiteOriginClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *DeleteSiteOriginClientCertificateResponseBody
	GetSiteName() *string
}

type DeleteSiteOriginClientCertificateResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId    *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName  *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s DeleteSiteOriginClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSiteOriginClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSiteOriginClientCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteSiteOriginClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSiteOriginClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteSiteOriginClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *DeleteSiteOriginClientCertificateResponseBody) SetId(v string) *DeleteSiteOriginClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteSiteOriginClientCertificateResponseBody) SetRequestId(v string) *DeleteSiteOriginClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSiteOriginClientCertificateResponseBody) SetSiteId(v int64) *DeleteSiteOriginClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *DeleteSiteOriginClientCertificateResponseBody) SetSiteName(v string) *DeleteSiteOriginClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *DeleteSiteOriginClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
