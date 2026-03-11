// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInstanceForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *GetAppInstanceForAdminRequest
	GetBizId() *string
	SetDomain(v string) *GetAppInstanceForAdminRequest
	GetDomain() *string
}

type GetAppInstanceForAdminRequest struct {
	// example:
	//
	// WS20250801154628000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// yjdw.bpu.edu.cn-waf
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
}

func (s GetAppInstanceForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppInstanceForAdminRequest) GoString() string {
	return s.String()
}

func (s *GetAppInstanceForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *GetAppInstanceForAdminRequest) GetDomain() *string {
	return s.Domain
}

func (s *GetAppInstanceForAdminRequest) SetBizId(v string) *GetAppInstanceForAdminRequest {
	s.BizId = &v
	return s
}

func (s *GetAppInstanceForAdminRequest) SetDomain(v string) *GetAppInstanceForAdminRequest {
	s.Domain = &v
	return s
}

func (s *GetAppInstanceForAdminRequest) Validate() error {
	return dara.Validate(s)
}
