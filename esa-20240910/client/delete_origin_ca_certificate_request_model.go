// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOriginCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteOriginCaCertificateRequest
	GetId() *string
	SetSiteId(v int64) *DeleteOriginCaCertificateRequest
	GetSiteId() *int64
}

type DeleteOriginCaCertificateRequest struct {
	// The certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaabcd****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteOriginCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteOriginCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteOriginCaCertificateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteOriginCaCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteOriginCaCertificateRequest) SetId(v string) *DeleteOriginCaCertificateRequest {
	s.Id = &v
	return s
}

func (s *DeleteOriginCaCertificateRequest) SetSiteId(v int64) *DeleteOriginCaCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteOriginCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
