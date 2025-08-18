// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginCaCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteOriginCaCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *DeleteOriginCaCertificateResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *DeleteOriginCaCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *DeleteOriginCaCertificateResponseBody
	GetSiteName() *string
}

type DeleteOriginCaCertificateResponseBody struct {
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
	// C370DAF1-C838-4288-A1A0-9A87633D248E
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

func (s DeleteOriginCaCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginCaCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOriginCaCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *DeleteOriginCaCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOriginCaCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteOriginCaCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *DeleteOriginCaCertificateResponseBody) SetId(v string) *DeleteOriginCaCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *DeleteOriginCaCertificateResponseBody) SetRequestId(v string) *DeleteOriginCaCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOriginCaCertificateResponseBody) SetSiteId(v int64) *DeleteOriginCaCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *DeleteOriginCaCertificateResponseBody) SetSiteName(v string) *DeleteOriginCaCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *DeleteOriginCaCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
