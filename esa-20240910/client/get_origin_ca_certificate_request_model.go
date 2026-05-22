// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOriginCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetOriginCaCertificateRequest
	GetId() *string
	SetSiteId(v int64) *GetOriginCaCertificateRequest
	GetSiteId() *int64
}

type GetOriginCaCertificateRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetOriginCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOriginCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetOriginCaCertificateRequest) GetId() *string {
	return s.Id
}

func (s *GetOriginCaCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetOriginCaCertificateRequest) SetId(v string) *GetOriginCaCertificateRequest {
	s.Id = &v
	return s
}

func (s *GetOriginCaCertificateRequest) SetSiteId(v int64) *GetOriginCaCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *GetOriginCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
