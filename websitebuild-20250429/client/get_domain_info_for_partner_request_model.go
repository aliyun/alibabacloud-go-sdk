// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainInfoForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetDomainInfoForPartnerRequest
	GetBizId() *string
	SetDomainName(v string) *GetDomainInfoForPartnerRequest
	GetDomainName() *string
	SetUserId(v string) *GetDomainInfoForPartnerRequest
	GetUserId() *string
}

type GetDomainInfoForPartnerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yywq.qimengwenhua.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDomainInfoForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDomainInfoForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GetDomainInfoForPartnerRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetDomainInfoForPartnerRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *GetDomainInfoForPartnerRequest) GetUserId() *string {
	return s.UserId
}

func (s *GetDomainInfoForPartnerRequest) SetBizId(v string) *GetDomainInfoForPartnerRequest {
	s.BizId = &v
	return s
}

func (s *GetDomainInfoForPartnerRequest) SetDomainName(v string) *GetDomainInfoForPartnerRequest {
	s.DomainName = &v
	return s
}

func (s *GetDomainInfoForPartnerRequest) SetUserId(v string) *GetDomainInfoForPartnerRequest {
	s.UserId = &v
	return s
}

func (s *GetDomainInfoForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
