// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *RevokeClientCertificateRequest
	GetId() *string
	SetSiteId(v int64) *RevokeClientCertificateRequest
	GetSiteId() *int64
}

type RevokeClientCertificateRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s RevokeClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *RevokeClientCertificateRequest) GetId() *string {
	return s.Id
}

func (s *RevokeClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *RevokeClientCertificateRequest) SetId(v string) *RevokeClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *RevokeClientCertificateRequest) SetSiteId(v int64) *RevokeClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *RevokeClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
