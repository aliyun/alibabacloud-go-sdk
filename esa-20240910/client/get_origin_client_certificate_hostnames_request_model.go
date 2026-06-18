// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginClientCertificateHostnamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetOriginClientCertificateHostnamesRequest
	GetId() *string
	SetSiteId(v int64) *GetOriginClientCertificateHostnamesRequest
	GetSiteId() *int64
}

type GetOriginClientCertificateHostnamesRequest struct {
	// The certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginClientCertificateHostnamesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateHostnamesRequest) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateHostnamesRequest) GetId() *string {
	return s.Id
}

func (s *GetOriginClientCertificateHostnamesRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginClientCertificateHostnamesRequest) SetId(v string) *GetOriginClientCertificateHostnamesRequest {
	s.Id = &v
	return s
}

func (s *GetOriginClientCertificateHostnamesRequest) SetSiteId(v int64) *GetOriginClientCertificateHostnamesRequest {
	s.SiteId = &v
	return s
}

func (s *GetOriginClientCertificateHostnamesRequest) Validate() error {
	return dara.Validate(s)
}
