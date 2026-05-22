// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSiteId(v int64) *GetCertificateQuotaRequest
	GetSiteId() *int64
	SetType(v string) *GetCertificateQuotaRequest
	GetType() *string
}

type GetCertificateQuotaRequest struct {
	// Site ID, which can be obtained by calling the [ListSites](~~ListSites~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// Certificate Quota type.
	//
	// This parameter is required.
	//
	// example:
	//
	// free
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetCertificateQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetCertificateQuotaRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetCertificateQuotaRequest) GetType() *string {
	return s.Type
}

func (s *GetCertificateQuotaRequest) SetSiteId(v int64) *GetCertificateQuotaRequest {
	s.SiteId = &v
	return s
}

func (s *GetCertificateQuotaRequest) SetType(v string) *GetCertificateQuotaRequest {
	s.Type = &v
	return s
}

func (s *GetCertificateQuotaRequest) Validate() error {
	return dara.Validate(s)
}
