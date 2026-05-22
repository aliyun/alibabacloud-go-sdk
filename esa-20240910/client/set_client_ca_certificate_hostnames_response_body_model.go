// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClientCaCertificateHostnamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *SetClientCaCertificateHostnamesResponseBody
	GetHostnames() []*string
	SetId(v string) *SetClientCaCertificateHostnamesResponseBody
	GetId() *string
	SetRequestId(v string) *SetClientCaCertificateHostnamesResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *SetClientCaCertificateHostnamesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *SetClientCaCertificateHostnamesResponseBody
	GetSiteName() *string
}

type SetClientCaCertificateHostnamesResponseBody struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 34DCBC8A-****-****-****-6DAA11D7DDBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s SetClientCaCertificateHostnamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetClientCaCertificateHostnamesResponseBody) GoString() string {
	return s.String()
}

func (s *SetClientCaCertificateHostnamesResponseBody) GetHostnames() []*string {
	return s.Hostnames
}

func (s *SetClientCaCertificateHostnamesResponseBody) GetId() *string {
	return s.Id
}

func (s *SetClientCaCertificateHostnamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetClientCaCertificateHostnamesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetClientCaCertificateHostnamesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *SetClientCaCertificateHostnamesResponseBody) SetHostnames(v []*string) *SetClientCaCertificateHostnamesResponseBody {
	s.Hostnames = v
	return s
}

func (s *SetClientCaCertificateHostnamesResponseBody) SetId(v string) *SetClientCaCertificateHostnamesResponseBody {
	s.Id = &v
	return s
}

func (s *SetClientCaCertificateHostnamesResponseBody) SetRequestId(v string) *SetClientCaCertificateHostnamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetClientCaCertificateHostnamesResponseBody) SetSiteId(v int64) *SetClientCaCertificateHostnamesResponseBody {
	s.SiteId = &v
	return s
}

func (s *SetClientCaCertificateHostnamesResponseBody) SetSiteName(v string) *SetClientCaCertificateHostnamesResponseBody {
	s.SiteName = &v
	return s
}

func (s *SetClientCaCertificateHostnamesResponseBody) Validate() error {
	return dara.Validate(s)
}
