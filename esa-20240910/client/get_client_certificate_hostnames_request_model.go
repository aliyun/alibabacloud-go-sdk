// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCertificateHostnamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetClientCertificateHostnamesRequest
	GetId() *string
	SetSiteId(v int64) *GetClientCertificateHostnamesRequest
	GetSiteId() *int64
}

type GetClientCertificateHostnamesRequest struct {
	// The certificate ID.
	//
	// example:
	//
	// baba39055622c008b90285a8838ed09a
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

func (s GetClientCertificateHostnamesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientCertificateHostnamesRequest) GoString() string {
	return s.String()
}

func (s *GetClientCertificateHostnamesRequest) GetId() *string {
	return s.Id
}

func (s *GetClientCertificateHostnamesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetClientCertificateHostnamesRequest) SetId(v string) *GetClientCertificateHostnamesRequest {
	s.Id = &v
	return s
}

func (s *GetClientCertificateHostnamesRequest) SetSiteId(v int64) *GetClientCertificateHostnamesRequest {
	s.SiteId = &v
	return s
}

func (s *GetClientCertificateHostnamesRequest) Validate() error {
	return dara.Validate(s)
}
