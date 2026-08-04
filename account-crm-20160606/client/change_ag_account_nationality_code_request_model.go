// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeAgAccountNationalityCodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ChangeAgAccountNationalityCodeRequest
	GetAppName() *string
	SetMpk(v string) *ChangeAgAccountNationalityCodeRequest
	GetMpk() *string
	SetNationalityCode(v string) *ChangeAgAccountNationalityCodeRequest
	GetNationalityCode() *string
	SetPK(v string) *ChangeAgAccountNationalityCodeRequest
	GetPK() *string
}

type ChangeAgAccountNationalityCodeRequest struct {
	AppName         *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	Mpk             *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	NationalityCode *string `json:"NationalityCode,omitempty" xml:"NationalityCode,omitempty"`
	PK              *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s ChangeAgAccountNationalityCodeRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeAgAccountNationalityCodeRequest) GoString() string {
	return s.String()
}

func (s *ChangeAgAccountNationalityCodeRequest) GetAppName() *string {
	return s.AppName
}

func (s *ChangeAgAccountNationalityCodeRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ChangeAgAccountNationalityCodeRequest) GetNationalityCode() *string {
	return s.NationalityCode
}

func (s *ChangeAgAccountNationalityCodeRequest) GetPK() *string {
	return s.PK
}

func (s *ChangeAgAccountNationalityCodeRequest) SetAppName(v string) *ChangeAgAccountNationalityCodeRequest {
	s.AppName = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeRequest) SetMpk(v string) *ChangeAgAccountNationalityCodeRequest {
	s.Mpk = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeRequest) SetNationalityCode(v string) *ChangeAgAccountNationalityCodeRequest {
	s.NationalityCode = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeRequest) SetPK(v string) *ChangeAgAccountNationalityCodeRequest {
	s.PK = &v
	return s
}

func (s *ChangeAgAccountNationalityCodeRequest) Validate() error {
	return dara.Validate(s)
}
