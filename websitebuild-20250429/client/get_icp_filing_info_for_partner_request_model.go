// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIcpFilingInfoForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetIcpFilingInfoForPartnerRequest
	GetBizId() *string
	SetDomain(v string) *GetIcpFilingInfoForPartnerRequest
	GetDomain() *string
}

type GetIcpFilingInfoForPartnerRequest struct {
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// yjdw.bpu.edu.cn-waf
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetIcpFilingInfoForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIcpFilingInfoForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GetIcpFilingInfoForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetIcpFilingInfoForPartnerRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetIcpFilingInfoForPartnerRequest) SetBizId(v string) *GetIcpFilingInfoForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *GetIcpFilingInfoForPartnerRequest) SetDomain(v string) *GetIcpFilingInfoForPartnerRequest {
	s.Domain = &v
	return s
}

func (s *GetIcpFilingInfoForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
