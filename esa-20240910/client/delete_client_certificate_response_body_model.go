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
	// 15C66C7B-671A-4297-9187-2C4477247A74
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
