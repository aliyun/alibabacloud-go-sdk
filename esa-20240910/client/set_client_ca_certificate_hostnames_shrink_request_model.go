// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetClientCaCertificateHostnamesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHostnamesShrink(v string) *SetClientCaCertificateHostnamesShrinkRequest
	GetHostnamesShrink() *string
	SetId(v string) *SetClientCaCertificateHostnamesShrinkRequest
	GetId() *string
	SetSiteId(v int64) *SetClientCaCertificateHostnamesShrinkRequest
	GetSiteId() *int64
}

type SetClientCaCertificateHostnamesShrinkRequest struct {
	// This parameter is required.
	HostnamesShrink *string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty"`
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

func (s SetClientCaCertificateHostnamesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SetClientCaCertificateHostnamesShrinkRequest) GoString() string {
	return s.String()
}

func (s *SetClientCaCertificateHostnamesShrinkRequest) GetHostnamesShrink() *string {
	return s.HostnamesShrink
}

func (s *SetClientCaCertificateHostnamesShrinkRequest) GetId() *string {
	return s.Id
}

func (s *SetClientCaCertificateHostnamesShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *SetClientCaCertificateHostnamesShrinkRequest) SetHostnamesShrink(v string) *SetClientCaCertificateHostnamesShrinkRequest {
	s.HostnamesShrink = &v
	return s
}

func (s *SetClientCaCertificateHostnamesShrinkRequest) SetId(v string) *SetClientCaCertificateHostnamesShrinkRequest {
	s.Id = &v
	return s
}

func (s *SetClientCaCertificateHostnamesShrinkRequest) SetSiteId(v int64) *SetClientCaCertificateHostnamesShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *SetClientCaCertificateHostnamesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
