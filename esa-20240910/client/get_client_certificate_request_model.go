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
	// The certificate ID.
	//
	// This parameter is required.
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
