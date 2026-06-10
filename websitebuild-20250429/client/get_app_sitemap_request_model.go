// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSitemapRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppSitemapRequest
	GetBizId() *string
	SetDomain(v string) *GetAppSitemapRequest
	GetDomain() *string
	SetSeType(v string) *GetAppSitemapRequest
	GetSeType() *string
}

type GetAppSitemapRequest struct {
	// Business ID
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Domain name
	//
	// example:
	//
	// yjdw.bpu.edu.cn-waf
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Search engine type
	//
	// example:
	//
	// type
	SeType *string `json:"SeType,omitempty" xml:"SeType,omitempty"`
}

func (s GetAppSitemapRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppSitemapRequest) GoString() string {
	return s.String()
}

func (s *GetAppSitemapRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppSitemapRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetAppSitemapRequest) GetSeType() *string {
	return s.SeType
}

func (s *GetAppSitemapRequest) SetBizId(v string) *GetAppSitemapRequest {
	s.BizId = &v
	return s
}

func (s *GetAppSitemapRequest) SetDomain(v string) *GetAppSitemapRequest {
	s.Domain = &v
	return s
}

func (s *GetAppSitemapRequest) SetSeType(v string) *GetAppSitemapRequest {
	s.SeType = &v
	return s
}

func (s *GetAppSitemapRequest) Validate() error {
	return dara.Validate(s)
}
