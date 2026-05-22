// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetClientCertificateRequest
	GetId() *string
	SetSiteId(v int64) *GetClientCertificateRequest
	GetSiteId() *int64
}

type GetClientCertificateRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetClientCertificateRequest) GetId() *string {
	return s.Id
}

func (s *GetClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetClientCertificateRequest) SetId(v string) *GetClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *GetClientCertificateRequest) SetSiteId(v int64) *GetClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *GetClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
