// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForbiddenAgAccountLoginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *ForbiddenAgAccountLoginRequest
	GetAgAccountType() *string
	SetAppName(v string) *ForbiddenAgAccountLoginRequest
	GetAppName() *string
	SetMpk(v string) *ForbiddenAgAccountLoginRequest
	GetMpk() *string
	SetPk(v string) *ForbiddenAgAccountLoginRequest
	GetPk() *string
}

type ForbiddenAgAccountLoginRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s ForbiddenAgAccountLoginRequest) String() string {
	return dara.Prettify(s)
}

func (s ForbiddenAgAccountLoginRequest) GoString() string {
	return s.String()
}

func (s *ForbiddenAgAccountLoginRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *ForbiddenAgAccountLoginRequest) GetAppName() *string {
	return s.AppName
}

func (s *ForbiddenAgAccountLoginRequest) GetMpk() *string {
	return s.Mpk
}

func (s *ForbiddenAgAccountLoginRequest) GetPk() *string {
	return s.Pk
}

func (s *ForbiddenAgAccountLoginRequest) SetAgAccountType(v string) *ForbiddenAgAccountLoginRequest {
	s.AgAccountType = &v
	return s
}

func (s *ForbiddenAgAccountLoginRequest) SetAppName(v string) *ForbiddenAgAccountLoginRequest {
	s.AppName = &v
	return s
}

func (s *ForbiddenAgAccountLoginRequest) SetMpk(v string) *ForbiddenAgAccountLoginRequest {
	s.Mpk = &v
	return s
}

func (s *ForbiddenAgAccountLoginRequest) SetPk(v string) *ForbiddenAgAccountLoginRequest {
	s.Pk = &v
	return s
}

func (s *ForbiddenAgAccountLoginRequest) Validate() error {
	return dara.Validate(s)
}
