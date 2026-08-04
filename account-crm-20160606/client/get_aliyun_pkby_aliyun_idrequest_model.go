// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAliyunPKByAliyunIDRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunId(v string) *GetAliyunPKByAliyunIDRequest
	GetAliyunId() *string
	SetEmail(v string) *GetAliyunPKByAliyunIDRequest
	GetEmail() *string
	SetHavanaId(v string) *GetAliyunPKByAliyunIDRequest
	GetHavanaId() *string
	SetMobile(v string) *GetAliyunPKByAliyunIDRequest
	GetMobile() *string
	SetPK(v string) *GetAliyunPKByAliyunIDRequest
	GetPK() *string
}

type GetAliyunPKByAliyunIDRequest struct {
	// This parameter is required.
	AliyunId *string `json:"AliyunId,omitempty" xml:"AliyunId,omitempty"`
	Email    *string `json:"Email,omitempty" xml:"Email,omitempty"`
	HavanaId *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	Mobile   *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	PK       *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s GetAliyunPKByAliyunIDRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAliyunPKByAliyunIDRequest) GoString() string {
	return s.String()
}

func (s *GetAliyunPKByAliyunIDRequest) GetAliyunId() *string {
	return s.AliyunId
}

func (s *GetAliyunPKByAliyunIDRequest) GetEmail() *string {
	return s.Email
}

func (s *GetAliyunPKByAliyunIDRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *GetAliyunPKByAliyunIDRequest) GetMobile() *string {
	return s.Mobile
}

func (s *GetAliyunPKByAliyunIDRequest) GetPK() *string {
	return s.PK
}

func (s *GetAliyunPKByAliyunIDRequest) SetAliyunId(v string) *GetAliyunPKByAliyunIDRequest {
	s.AliyunId = &v
	return s
}

func (s *GetAliyunPKByAliyunIDRequest) SetEmail(v string) *GetAliyunPKByAliyunIDRequest {
	s.Email = &v
	return s
}

func (s *GetAliyunPKByAliyunIDRequest) SetHavanaId(v string) *GetAliyunPKByAliyunIDRequest {
	s.HavanaId = &v
	return s
}

func (s *GetAliyunPKByAliyunIDRequest) SetMobile(v string) *GetAliyunPKByAliyunIDRequest {
	s.Mobile = &v
	return s
}

func (s *GetAliyunPKByAliyunIDRequest) SetPK(v string) *GetAliyunPKByAliyunIDRequest {
	s.PK = &v
	return s
}

func (s *GetAliyunPKByAliyunIDRequest) Validate() error {
	return dara.Validate(s)
}
