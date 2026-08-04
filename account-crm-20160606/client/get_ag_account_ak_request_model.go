// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgAccountAkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *GetAgAccountAkRequest
	GetAgAccountType() *string
	SetMpk(v string) *GetAgAccountAkRequest
	GetMpk() *string
	SetPk(v string) *GetAgAccountAkRequest
	GetPk() *string
}

type GetAgAccountAkRequest struct {
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s GetAgAccountAkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAgAccountAkRequest) GoString() string {
	return s.String()
}

func (s *GetAgAccountAkRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *GetAgAccountAkRequest) GetMpk() *string {
	return s.Mpk
}

func (s *GetAgAccountAkRequest) GetPk() *string {
	return s.Pk
}

func (s *GetAgAccountAkRequest) SetAgAccountType(v string) *GetAgAccountAkRequest {
	s.AgAccountType = &v
	return s
}

func (s *GetAgAccountAkRequest) SetMpk(v string) *GetAgAccountAkRequest {
	s.Mpk = &v
	return s
}

func (s *GetAgAccountAkRequest) SetPk(v string) *GetAgAccountAkRequest {
	s.Pk = &v
	return s
}

func (s *GetAgAccountAkRequest) Validate() error {
	return dara.Validate(s)
}
