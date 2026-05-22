// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCaCertificateHostnamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *GetClientCaCertificateHostnamesResponseBody
	GetHostnames() []*string
	SetId(v string) *GetClientCaCertificateHostnamesResponseBody
	GetId() *string
	SetRequestId(v string) *GetClientCaCertificateHostnamesResponseBody
	GetRequestId() *string
	SetSiteId(v int64) *GetClientCaCertificateHostnamesResponseBody
	GetSiteId() *int64
	SetSiteName(v string) *GetClientCaCertificateHostnamesResponseBody
	GetSiteName() *string
}

type GetClientCaCertificateHostnamesResponseBody struct {
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// example:
	//
	// babaded901474b9693acf530e0fb****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 3558df77-8a7a-4060-a900-2d7949403836
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// example:
	//
	// example.com
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s GetClientCaCertificateHostnamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetClientCaCertificateHostnamesResponseBody) GoString() string {
	return s.String()
}

func (s *GetClientCaCertificateHostnamesResponseBody) GetHostnames() []*string {
	return s.Hostnames
}

func (s *GetClientCaCertificateHostnamesResponseBody) GetId() *string {
	return s.Id
}

func (s *GetClientCaCertificateHostnamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetClientCaCertificateHostnamesResponseBody) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetClientCaCertificateHostnamesResponseBody) GetSiteName() *string {
	return s.SiteName
}

func (s *GetClientCaCertificateHostnamesResponseBody) SetHostnames(v []*string) *GetClientCaCertificateHostnamesResponseBody {
	s.Hostnames = v
	return s
}

func (s *GetClientCaCertificateHostnamesResponseBody) SetId(v string) *GetClientCaCertificateHostnamesResponseBody {
	s.Id = &v
	return s
}

func (s *GetClientCaCertificateHostnamesResponseBody) SetRequestId(v string) *GetClientCaCertificateHostnamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetClientCaCertificateHostnamesResponseBody) SetSiteId(v int64) *GetClientCaCertificateHostnamesResponseBody {
	s.SiteId = &v
	return s
}

func (s *GetClientCaCertificateHostnamesResponseBody) SetSiteName(v string) *GetClientCaCertificateHostnamesResponseBody {
	s.SiteName = &v
	return s
}

func (s *GetClientCaCertificateHostnamesResponseBody) Validate() error {
	return dara.Validate(s)
}
