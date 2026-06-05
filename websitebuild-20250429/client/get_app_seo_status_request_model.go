// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppSeoStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppSeoStatusRequest
	GetBizId() *string
	SetDomain(v string) *GetAppSeoStatusRequest
	GetDomain() *string
	SetSeType(v string) *GetAppSeoStatusRequest
	GetSeType() *string
}

type GetAppSeoStatusRequest struct {
	// example:
	//
	// WD20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// yjdw.bpu.edu.cn-waf
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// type
	SeType *string `json:"SeType,omitempty" xml:"SeType,omitempty"`
}

func (s GetAppSeoStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppSeoStatusRequest) GoString() string {
	return s.String()
}

func (s *GetAppSeoStatusRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppSeoStatusRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetAppSeoStatusRequest) GetSeType() *string {
	return s.SeType
}

func (s *GetAppSeoStatusRequest) SetBizId(v string) *GetAppSeoStatusRequest {
	s.BizId = &v
	return s
}

func (s *GetAppSeoStatusRequest) SetDomain(v string) *GetAppSeoStatusRequest {
	s.Domain = &v
	return s
}

func (s *GetAppSeoStatusRequest) SetSeType(v string) *GetAppSeoStatusRequest {
	s.SeType = &v
	return s
}

func (s *GetAppSeoStatusRequest) Validate() error {
	return dara.Validate(s)
}
