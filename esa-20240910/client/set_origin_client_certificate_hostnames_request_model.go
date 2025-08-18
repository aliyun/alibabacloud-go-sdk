// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOriginClientCertificateHostnamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *SetOriginClientCertificateHostnamesRequest
	GetHostnames() []*string
	SetId(v string) *SetOriginClientCertificateHostnamesRequest
	GetId() *string
	SetSiteId(v int64) *SetOriginClientCertificateHostnamesRequest
	GetSiteId() *int64
}

type SetOriginClientCertificateHostnamesRequest struct {
	// The domain names to associate.
	//
	// This parameter is required.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetOriginClientCertificateHostnamesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetOriginClientCertificateHostnamesRequest) GoString() string {
	return s.String()
}

func (s *SetOriginClientCertificateHostnamesRequest) GetHostnames() []*string {
	return s.Hostnames
}

func (s *SetOriginClientCertificateHostnamesRequest) GetId() *string {
	return s.Id
}

func (s *SetOriginClientCertificateHostnamesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetOriginClientCertificateHostnamesRequest) SetHostnames(v []*string) *SetOriginClientCertificateHostnamesRequest {
	s.Hostnames = v
	return s
}

func (s *SetOriginClientCertificateHostnamesRequest) SetId(v string) *SetOriginClientCertificateHostnamesRequest {
	s.Id = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesRequest) SetSiteId(v int64) *SetOriginClientCertificateHostnamesRequest {
	s.SiteId = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesRequest) Validate() error {
	return dara.Validate(s)
}
