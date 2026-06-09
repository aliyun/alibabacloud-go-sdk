// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAppSeoStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UpdateAppSeoStatusRequest
	GetBizId() *string
	SetDomain(v string) *UpdateAppSeoStatusRequest
	GetDomain() *string
	SetSeAuthInfo(v string) *UpdateAppSeoStatusRequest
	GetSeAuthInfo() *string
	SetSeType(v string) *UpdateAppSeoStatusRequest
	GetSeType() *string
}

type UpdateAppSeoStatusRequest struct {
	// Business ID
	//
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// Domain Name
	//
	// example:
	//
	// yjdw.bpu.edu.cn-waf
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 123123
	SeAuthInfo *string `json:"SeAuthInfo,omitempty" xml:"SeAuthInfo,omitempty"`
	// Search Engine Type
	//
	// example:
	//
	// type
	SeType *string `json:"SeType,omitempty" xml:"SeType,omitempty"`
}

func (s UpdateAppSeoStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAppSeoStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateAppSeoStatusRequest) GetBizId() *string {
	return s.BizId
}

func (s *UpdateAppSeoStatusRequest) GetDomain() *string {
	return s.Domain
}

func (s *UpdateAppSeoStatusRequest) GetSeAuthInfo() *string {
	return s.SeAuthInfo
}

func (s *UpdateAppSeoStatusRequest) GetSeType() *string {
	return s.SeType
}

func (s *UpdateAppSeoStatusRequest) SetBizId(v string) *UpdateAppSeoStatusRequest {
	s.BizId = &v
	return s
}

func (s *UpdateAppSeoStatusRequest) SetDomain(v string) *UpdateAppSeoStatusRequest {
	s.Domain = &v
	return s
}

func (s *UpdateAppSeoStatusRequest) SetSeAuthInfo(v string) *UpdateAppSeoStatusRequest {
	s.SeAuthInfo = &v
	return s
}

func (s *UpdateAppSeoStatusRequest) SetSeType(v string) *UpdateAppSeoStatusRequest {
	s.SeType = &v
	return s
}

func (s *UpdateAppSeoStatusRequest) Validate() error {
	return dara.Validate(s)
}
