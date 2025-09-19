// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAppDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *BindAppDomainRequest
	GetBizId() *string
	SetDomainName(v string) *BindAppDomainRequest
	GetDomainName() *string
	SetExtend(v string) *BindAppDomainRequest
	GetExtend() *string
	SetOperateType(v string) *BindAppDomainRequest
	GetOperateType() *string
}

type BindAppDomainRequest struct {
	// example:
	//
	// WD20250820143531000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// xiaoniu.link
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	Extend     *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// example:
	//
	// vul_fix
	OperateType *string `json:"OperateType,omitempty" xml:"OperateType,omitempty"`
}

func (s BindAppDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAppDomainRequest) GoString() string {
	return s.String()
}

func (s *BindAppDomainRequest) GetBizId() *string {
	return s.BizId
}

func (s *BindAppDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *BindAppDomainRequest) GetExtend() *string {
	return s.Extend
}

func (s *BindAppDomainRequest) GetOperateType() *string {
	return s.OperateType
}

func (s *BindAppDomainRequest) SetBizId(v string) *BindAppDomainRequest {
	s.BizId = &v
	return s
}

func (s *BindAppDomainRequest) SetDomainName(v string) *BindAppDomainRequest {
	s.DomainName = &v
	return s
}

func (s *BindAppDomainRequest) SetExtend(v string) *BindAppDomainRequest {
	s.Extend = &v
	return s
}

func (s *BindAppDomainRequest) SetOperateType(v string) *BindAppDomainRequest {
	s.OperateType = &v
	return s
}

func (s *BindAppDomainRequest) Validate() error {
	return dara.Validate(s)
}
