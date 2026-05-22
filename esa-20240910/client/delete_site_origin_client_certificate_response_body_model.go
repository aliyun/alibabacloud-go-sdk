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
	// The certificate ID on ESA.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
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
