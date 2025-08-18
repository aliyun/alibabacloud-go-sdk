// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ActivateClientCertificateResponseBody
	GetId() *string
	SetRequestId(v string) *ActivateClientCertificateResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *ActivateClientCertificateResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *ActivateClientCertificateResponseBody
	GetSiteName() *string
}

type ActivateClientCertificateResponseBody struct {
	// The certificate ID.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0AEDAF20-4DDF-4165-8750-47FF9C1929C9
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

func (s ActivateClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateClientCertificateResponseBody) GetId() *string {
	return s.Id
}

func (s *ActivateClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateClientCertificateResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *ActivateClientCertificateResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *ActivateClientCertificateResponseBody) SetId(v string) *ActivateClientCertificateResponseBody {
	s.Id = &v
	return s
}

func (s *ActivateClientCertificateResponseBody) SetRequestId(v string) *ActivateClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateClientCertificateResponseBody) SetSiteId(v int64) *ActivateClientCertificateResponseBody {
	s.SiteId = &v
	return s
}

func (s *ActivateClientCertificateResponseBody) SetSiteName(v string) *ActivateClientCertificateResponseBody {
	s.SiteName = &v
	return s
}

func (s *ActivateClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
