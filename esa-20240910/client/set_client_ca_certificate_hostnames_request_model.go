// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClientCaCertificateHostnamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnames(v []*string) *SetClientCaCertificateHostnamesRequest
	GetHostnames() []*string
	SetId(v string) *SetClientCaCertificateHostnamesRequest
	GetId() *string
	SetSiteId(v int64) *SetClientCaCertificateHostnamesRequest
	GetSiteId() *int64
}

type SetClientCaCertificateHostnamesRequest struct {
	// This parameter is required.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetClientCaCertificateHostnamesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetClientCaCertificateHostnamesRequest) GoString() string {
	return s.String()
}

func (s *SetClientCaCertificateHostnamesRequest) GetHostnames() []*string {
	return s.Hostnames
}

func (s *SetClientCaCertificateHostnamesRequest) GetId() *string {
	return s.Id
}

func (s *SetClientCaCertificateHostnamesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetClientCaCertificateHostnamesRequest) SetHostnames(v []*string) *SetClientCaCertificateHostnamesRequest {
	s.Hostnames = v
	return s
}

func (s *SetClientCaCertificateHostnamesRequest) SetId(v string) *SetClientCaCertificateHostnamesRequest {
	s.Id = &v
	return s
}

func (s *SetClientCaCertificateHostnamesRequest) SetSiteId(v int64) *SetClientCaCertificateHostnamesRequest {
	s.SiteId = &v
	return s
}

func (s *SetClientCaCertificateHostnamesRequest) Validate() error {
	return dara.Validate(s)
}
