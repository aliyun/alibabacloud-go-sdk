// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginClientCertificateHostnamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *GetOriginClientCertificateHostnamesResponseBody
	GetHostnames() []*string
	SetId(v string) *GetOriginClientCertificateHostnamesResponseBody
	GetId() *string
	SetRequestId(v string) *GetOriginClientCertificateHostnamesResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetOriginClientCertificateHostnamesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetOriginClientCertificateHostnamesResponseBody
	GetSiteName() *string
}

type GetOriginClientCertificateHostnamesResponseBody struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	Id        *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId    *int64    `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName  *string   `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s GetOriginClientCertificateHostnamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateHostnamesResponseBody) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateHostnamesResponseBody) GetHostnames() []*string {
	return s.Hostnames
}

func (s *GetOriginClientCertificateHostnamesResponseBody) GetId() *string {
	return s.Id
}

func (s *GetOriginClientCertificateHostnamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOriginClientCertificateHostnamesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginClientCertificateHostnamesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetOriginClientCertificateHostnamesResponseBody) SetHostnames(v []*string) *GetOriginClientCertificateHostnamesResponseBody {
	s.Hostnames = v
	return s
}

func (s *GetOriginClientCertificateHostnamesResponseBody) SetId(v string) *GetOriginClientCertificateHostnamesResponseBody {
	s.Id = &v
	return s
}

func (s *GetOriginClientCertificateHostnamesResponseBody) SetRequestId(v string) *GetOriginClientCertificateHostnamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOriginClientCertificateHostnamesResponseBody) SetSiteId(v int64) *GetOriginClientCertificateHostnamesResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetOriginClientCertificateHostnamesResponseBody) SetSiteName(v string) *GetOriginClientCertificateHostnamesResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetOriginClientCertificateHostnamesResponseBody) Validate() error {
	return dara.Validate(s)
}
