// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteClientCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteClientCertificateResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DeleteClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *DeleteClientCertificateResponseBody
	GetSiteName() *string
}

type DeleteClientCertificateResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId    *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName  *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s DeleteClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *DeleteClientCertificateResponseBody) SetId(v string) *DeleteClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteClientCertificateResponseBody) SetRequestId(v string) *DeleteClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientCertificateResponseBody) SetSiteId(v int64) *DeleteClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *DeleteClientCertificateResponseBody) SetSiteName(v string) *DeleteClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *DeleteClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
