// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteCertificateRequest
	GetId() *string
	SetSiteId(v int64) *DeleteCertificateRequest
	GetSiteId() *int64
}

type DeleteCertificateRequest struct {
	// The certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1d95
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](~~ListSites~~) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCertificateRequest) GetId() *string {
	return s.Id
}

func (s *DeleteCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteCertificateRequest) SetId(v string) *DeleteCertificateRequest {
	s.Id = &v
	return s
}

func (s *DeleteCertificateRequest) SetSiteId(v int64) *DeleteCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteCertificateRequest) Validate() error {
	return dara.Validate(s)
}
