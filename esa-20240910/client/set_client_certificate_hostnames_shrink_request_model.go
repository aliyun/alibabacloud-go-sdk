// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClientCertificateHostnamesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnamesShrink(v string) *SetClientCertificateHostnamesShrinkRequest
	GetHostnamesShrink() *string
	SetId(v string) *SetClientCertificateHostnamesShrinkRequest
	GetId() *string
	SetSiteId(v int64) *SetClientCertificateHostnamesShrinkRequest
	GetSiteId() *int64
}

type SetClientCertificateHostnamesShrinkRequest struct {
	// This parameter is required.
	HostnamesShrink *string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty"`
	Id              *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s SetClientCertificateHostnamesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetClientCertificateHostnamesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetClientCertificateHostnamesShrinkRequest) GetHostnamesShrink() *string {
	return s.HostnamesShrink
}

func (s *SetClientCertificateHostnamesShrinkRequest) GetId() *string {
	return s.Id
}

func (s *SetClientCertificateHostnamesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetClientCertificateHostnamesShrinkRequest) SetHostnamesShrink(v string) *SetClientCertificateHostnamesShrinkRequest {
	s.HostnamesShrink = &v
	return s
}

func (s *SetClientCertificateHostnamesShrinkRequest) SetId(v string) *SetClientCertificateHostnamesShrinkRequest {
	s.Id = &v
	return s
}

func (s *SetClientCertificateHostnamesShrinkRequest) SetSiteId(v int64) *SetClientCertificateHostnamesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *SetClientCertificateHostnamesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
