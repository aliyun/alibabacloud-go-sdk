// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteOriginClientCertificateRequest
	GetId() *string
	SetSiteId(v int64) *DeleteOriginClientCertificateRequest
	GetSiteId() *int64
}

type DeleteOriginClientCertificateRequest struct {
	// The certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteOriginClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteOriginClientCertificateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteOriginClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteOriginClientCertificateRequest) SetId(v string) *DeleteOriginClientCertificateRequest {
	s.Id = &v
	return s
}

func (s *DeleteOriginClientCertificateRequest) SetSiteId(v int64) *DeleteOriginClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteOriginClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
