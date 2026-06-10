// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSeoTrendsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppSeoTrendsRequest
	GetBizId() *string
	SetDomain(v string) *GetAppSeoTrendsRequest
	GetDomain() *string
	SetSeType(v string) *GetAppSeoTrendsRequest
	GetSeType() *string
}

type GetAppSeoTrendsRequest struct {
	// Business ID
	//
	// example:
	//
	// WD20250814102215000001
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

func (s GetAppSeoTrendsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppSeoTrendsRequest) GoString() string {
	return s.String()
}

func (s *GetAppSeoTrendsRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppSeoTrendsRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetAppSeoTrendsRequest) GetSeType() *string {
	return s.SeType
}

func (s *GetAppSeoTrendsRequest) SetBizId(v string) *GetAppSeoTrendsRequest {
	s.BizId = &v
	return s
}

func (s *GetAppSeoTrendsRequest) SetDomain(v string) *GetAppSeoTrendsRequest {
	s.Domain = &v
	return s
}

func (s *GetAppSeoTrendsRequest) SetSeType(v string) *GetAppSeoTrendsRequest {
	s.SeType = &v
	return s
}

func (s *GetAppSeoTrendsRequest) Validate() error {
	return dara.Validate(s)
}
