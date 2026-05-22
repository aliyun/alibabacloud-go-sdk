// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteClientCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteClientCaCertificateRequest
	GetId() *string
	SetSiteId(v int64) *DeleteClientCaCertificateRequest
	GetSiteId() *int64
}

type DeleteClientCaCertificateRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteClientCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteClientCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteClientCaCertificateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteClientCaCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteClientCaCertificateRequest) SetId(v string) *DeleteClientCaCertificateRequest {
	s.Id = &v
	return s
}

func (s *DeleteClientCaCertificateRequest) SetSiteId(v int64) *DeleteClientCaCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteClientCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
