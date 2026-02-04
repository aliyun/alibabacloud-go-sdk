// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppDomainCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DeleteAppDomainCertificateRequest
	GetBizId() *string
	SetDomainName(v string) *DeleteAppDomainCertificateRequest
	GetDomainName() *string
}

type DeleteAppDomainCertificateRequest struct {
	// Business ID of the application instance
	//
	// example:
	//
	// 202506170003
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Domain name
	//
	// example:
	//
	// aliwmzs.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s DeleteAppDomainCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppDomainCertificateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppDomainCertificateRequest) GetBizId() *string {
	return s.BizId
}

func (s *DeleteAppDomainCertificateRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DeleteAppDomainCertificateRequest) SetBizId(v string) *DeleteAppDomainCertificateRequest {
	s.BizId = &v
	return s
}

func (s *DeleteAppDomainCertificateRequest) SetDomainName(v string) *DeleteAppDomainCertificateRequest {
	s.DomainName = &v
	return s
}

func (s *DeleteAppDomainCertificateRequest) Validate() error {
	return dara.Validate(s)
}
