// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteCertificateResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DeleteCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *DeleteCertificateResponseBody
	GetSiteName() *string
}

type DeleteCertificateResponseBody struct {
	// The certificate ID.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1d95
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F32C57AA-7BF8-49AE-A2CC-9F42390F5A19
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

func (s DeleteCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *DeleteCertificateResponseBody) SetId(v string) *DeleteCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteCertificateResponseBody) SetRequestId(v string) *DeleteCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCertificateResponseBody) SetSiteId(v int64) *DeleteCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *DeleteCertificateResponseBody) SetSiteName(v string) *DeleteCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *DeleteCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
