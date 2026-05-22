// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteClientCertificateRequest
	GetId() *string
	SetSiteId(v int64) *DeleteClientCertificateRequest
	GetSiteId() *int64
}

type DeleteClientCertificateRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientCertificateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteClientCertificateRequest) SetId(v string) *DeleteClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *DeleteClientCertificateRequest) SetSiteId(v int64) *DeleteClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
