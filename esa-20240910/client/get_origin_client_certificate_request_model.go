// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetOriginClientCertificateRequest
	GetId() *string
	SetSiteId(v int64) *GetOriginClientCertificateRequest
	GetSiteId() *int64
}

type GetOriginClientCertificateRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetOriginClientCertificateRequest) GetId() *string {
	return s.Id
}

func (s *GetOriginClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginClientCertificateRequest) SetId(v string) *GetOriginClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *GetOriginClientCertificateRequest) SetSiteId(v int64) *GetOriginClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *GetOriginClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
