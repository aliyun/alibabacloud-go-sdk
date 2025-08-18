// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetCertificateRequest
	GetId() *string
	SetSiteId(v int64) *GetCertificateRequest
	GetSiteId() *int64
}

type GetCertificateRequest struct {
	// Certificate ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// babaded901474b9693acf530e0fb1d95
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetCertificateRequest) GetId() *string {
	return s.Id
}

func (s *GetCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCertificateRequest) SetId(v string) *GetCertificateRequest {
	s.Id = &v
	return s
}

func (s *GetCertificateRequest) SetSiteId(v int64) *GetCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *GetCertificateRequest) Validate() error {
	return dara.Validate(s)
}
