// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteClientCaCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteClientCaCertificateResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DeleteClientCaCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *DeleteClientCaCertificateResponseBody
	GetSiteName() *string
}

type DeleteClientCaCertificateResponseBody struct {
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId    *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName  *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s DeleteClientCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClientCaCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteClientCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClientCaCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteClientCaCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *DeleteClientCaCertificateResponseBody) SetId(v string) *DeleteClientCaCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteClientCaCertificateResponseBody) SetRequestId(v string) *DeleteClientCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClientCaCertificateResponseBody) SetSiteId(v int64) *DeleteClientCaCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *DeleteClientCaCertificateResponseBody) SetSiteName(v string) *DeleteClientCaCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *DeleteClientCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
