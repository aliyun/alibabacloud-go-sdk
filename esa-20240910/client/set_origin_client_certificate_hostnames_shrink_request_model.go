// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetOriginClientCertificateHostnamesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnamesShrink(v string) *SetOriginClientCertificateHostnamesShrinkRequest
	GetHostnamesShrink() *string
	SetId(v string) *SetOriginClientCertificateHostnamesShrinkRequest
	GetId() *string
	SetSiteId(v int64) *SetOriginClientCertificateHostnamesShrinkRequest
	GetSiteId() *int64
}

type SetOriginClientCertificateHostnamesShrinkRequest struct {
	// The domain names to associate.
	//
	// This parameter is required.
	HostnamesShrink *string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty"`
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

func (s SetOriginClientCertificateHostnamesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetOriginClientCertificateHostnamesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetOriginClientCertificateHostnamesShrinkRequest) GetHostnamesShrink() *string {
	return s.HostnamesShrink
}

func (s *SetOriginClientCertificateHostnamesShrinkRequest) GetId() *string {
	return s.Id
}

func (s *SetOriginClientCertificateHostnamesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetOriginClientCertificateHostnamesShrinkRequest) SetHostnamesShrink(v string) *SetOriginClientCertificateHostnamesShrinkRequest {
	s.HostnamesShrink = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesShrinkRequest) SetId(v string) *SetOriginClientCertificateHostnamesShrinkRequest {
	s.Id = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesShrinkRequest) SetSiteId(v int64) *SetOriginClientCertificateHostnamesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *SetOriginClientCertificateHostnamesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
