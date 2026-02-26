// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySupabaseInstanceInfoForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *QuerySupabaseInstanceInfoForAdminRequest
	GetBizId() *string
	SetEnv(v string) *QuerySupabaseInstanceInfoForAdminRequest
	GetEnv() *string
	SetUserId(v string) *QuerySupabaseInstanceInfoForAdminRequest
	GetUserId() *string
}

type QuerySupabaseInstanceInfoForAdminRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// staging
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 123
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QuerySupabaseInstanceInfoForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySupabaseInstanceInfoForAdminRequest) GoString() string {
	return s.String()
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetEnv() *string {
	return s.Env
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) GetUserId() *string {
	return s.UserId
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetBizId(v string) *QuerySupabaseInstanceInfoForAdminRequest {
	s.BizId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetEnv(v string) *QuerySupabaseInstanceInfoForAdminRequest {
	s.Env = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) SetUserId(v string) *QuerySupabaseInstanceInfoForAdminRequest {
	s.UserId = &v
	return s
}

func (s *QuerySupabaseInstanceInfoForAdminRequest) Validate() error {
	return dara.Validate(s)
}
