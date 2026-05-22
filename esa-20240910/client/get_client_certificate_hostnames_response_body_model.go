// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCertificateHostnamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *GetClientCertificateHostnamesResponseBody
	GetHostnames() []*string
	SetId(v string) *GetClientCertificateHostnamesResponseBody
	GetId() *string
	SetRequestId(v string) *GetClientCertificateHostnamesResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetClientCertificateHostnamesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetClientCertificateHostnamesResponseBody
	GetSiteName() *string
}

type GetClientCertificateHostnamesResponseBody struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	Id        *string   `json:"Id,omitempty" xml:"Id,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SiteId    *int64    `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName  *string   `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s GetClientCertificateHostnamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientCertificateHostnamesResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientCertificateHostnamesResponseBody) GetHostnames() []*string {
	return s.Hostnames
}

func (s *GetClientCertificateHostnamesResponseBody) GetId() *string {
	return s.Id
}

func (s *GetClientCertificateHostnamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientCertificateHostnamesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetClientCertificateHostnamesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetClientCertificateHostnamesResponseBody) SetHostnames(v []*string) *GetClientCertificateHostnamesResponseBody {
	s.Hostnames = v
	return s
}

func (s *GetClientCertificateHostnamesResponseBody) SetId(v string) *GetClientCertificateHostnamesResponseBody {
	s.Id = &v
	return s
}

func (s *GetClientCertificateHostnamesResponseBody) SetRequestId(v string) *GetClientCertificateHostnamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientCertificateHostnamesResponseBody) SetSiteId(v int64) *GetClientCertificateHostnamesResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetClientCertificateHostnamesResponseBody) SetSiteName(v string) *GetClientCertificateHostnamesResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetClientCertificateHostnamesResponseBody) Validate() error {
	return dara.Validate(s)
}
