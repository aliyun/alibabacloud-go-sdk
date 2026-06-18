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
	// The ID of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the site. Call the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API to get this value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
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
