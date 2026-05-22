// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClientCertificateHostnamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *SetClientCertificateHostnamesResponseBody
	GetHostnames() []*string
	SetId(v string) *SetClientCertificateHostnamesResponseBody
	GetId() *string
	SetRequestId(v string) *SetClientCertificateHostnamesResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *SetClientCertificateHostnamesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *SetClientCertificateHostnamesResponseBody
	GetSiteName() *string
}

type SetClientCertificateHostnamesResponseBody struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The ID of the client CA certificate.
	//
	// example:
	//
	// babab9db65ee5efcca9f3d41d4b50d66
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A528000
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

func (s SetClientCertificateHostnamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetClientCertificateHostnamesResponseBody) GoString() string {
	return s.String()
}

func (s *SetClientCertificateHostnamesResponseBody) GetHostnames() []*string {
	return s.Hostnames
}

func (s *SetClientCertificateHostnamesResponseBody) GetId() *string {
	return s.Id
}

func (s *SetClientCertificateHostnamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetClientCertificateHostnamesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetClientCertificateHostnamesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *SetClientCertificateHostnamesResponseBody) SetHostnames(v []*string) *SetClientCertificateHostnamesResponseBody {
	s.Hostnames = v
	return s
}

func (s *SetClientCertificateHostnamesResponseBody) SetId(v string) *SetClientCertificateHostnamesResponseBody {
	s.Id = &v
	return s
}

func (s *SetClientCertificateHostnamesResponseBody) SetRequestId(v string) *SetClientCertificateHostnamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetClientCertificateHostnamesResponseBody) SetSiteId(v int64) *SetClientCertificateHostnamesResponseBody {
	s.SiteId = &v
	return s
}

func (s *SetClientCertificateHostnamesResponseBody) SetSiteName(v string) *SetClientCertificateHostnamesResponseBody {
	s.SiteName = &v
	return s
}

func (s *SetClientCertificateHostnamesResponseBody) Validate() error {
	return dara.Validate(s)
}
