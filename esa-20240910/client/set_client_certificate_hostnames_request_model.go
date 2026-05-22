// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClientCertificateHostnamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *SetClientCertificateHostnamesRequest
	GetHostnames() []*string
	SetId(v string) *SetClientCertificateHostnamesRequest
	GetId() *string
	SetSiteId(v int64) *SetClientCertificateHostnamesRequest
	GetSiteId() *int64
}

type SetClientCertificateHostnamesRequest struct {
	// The domain names to associate.
	//
	// This parameter is required.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The ID of the client CA certificate.
	//
	// example:
	//
	// babab9db65ee5efcca9f3d41d4b50d66
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The website ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetClientCertificateHostnamesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetClientCertificateHostnamesRequest) GoString() string {
	return s.String()
}

func (s *SetClientCertificateHostnamesRequest) GetHostnames() []*string {
	return s.Hostnames
}

func (s *SetClientCertificateHostnamesRequest) GetId() *string {
	return s.Id
}

func (s *SetClientCertificateHostnamesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetClientCertificateHostnamesRequest) SetHostnames(v []*string) *SetClientCertificateHostnamesRequest {
	s.Hostnames = v
	return s
}

func (s *SetClientCertificateHostnamesRequest) SetId(v string) *SetClientCertificateHostnamesRequest {
	s.Id = &v
	return s
}

func (s *SetClientCertificateHostnamesRequest) SetSiteId(v int64) *SetClientCertificateHostnamesRequest {
	s.SiteId = &v
	return s
}

func (s *SetClientCertificateHostnamesRequest) Validate() error {
	return dara.Validate(s)
}
