// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOriginClientCertificateHostnamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *SetOriginClientCertificateHostnamesResponseBody
	GetHostnames() []*string
	SetId(v string) *SetOriginClientCertificateHostnamesResponseBody
	GetId() *string
	SetRequestId(v string) *SetOriginClientCertificateHostnamesResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *SetOriginClientCertificateHostnamesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *SetOriginClientCertificateHostnamesResponseBody
	GetSiteName() *string
}

type SetOriginClientCertificateHostnamesResponseBody struct {
	// The domain name.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The ID of the client certificate.
	//
	// example:
	//
	// babaabcd****
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
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The website name.
	//
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s SetOriginClientCertificateHostnamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetOriginClientCertificateHostnamesResponseBody) GoString() string {
	return s.String()
}

func (s *SetOriginClientCertificateHostnamesResponseBody) GetHostnames() []*string {
	return s.Hostnames
}

func (s *SetOriginClientCertificateHostnamesResponseBody) GetId() *string {
	return s.Id
}

func (s *SetOriginClientCertificateHostnamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetOriginClientCertificateHostnamesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetOriginClientCertificateHostnamesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *SetOriginClientCertificateHostnamesResponseBody) SetHostnames(v []*string) *SetOriginClientCertificateHostnamesResponseBody {
	s.Hostnames = v
	return s
}

func (s *SetOriginClientCertificateHostnamesResponseBody) SetId(v string) *SetOriginClientCertificateHostnamesResponseBody {
	s.Id = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesResponseBody) SetRequestId(v string) *SetOriginClientCertificateHostnamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesResponseBody) SetSiteId(v int64) *SetOriginClientCertificateHostnamesResponseBody {
	s.SiteId = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesResponseBody) SetSiteName(v string) *SetOriginClientCertificateHostnamesResponseBody {
	s.SiteName = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesResponseBody) Validate() error {
	return dara.Validate(s)
}
