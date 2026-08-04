// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgAccountAddressInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *UpdateAgAccountAddressInfoRequest
	GetAddress() *string
	SetAddress2(v string) *UpdateAgAccountAddressInfoRequest
	GetAddress2() *string
	SetAppName(v string) *UpdateAgAccountAddressInfoRequest
	GetAppName() *string
	SetCity(v string) *UpdateAgAccountAddressInfoRequest
	GetCity() *string
	SetMpk(v string) *UpdateAgAccountAddressInfoRequest
	GetMpk() *string
	SetPK(v string) *UpdateAgAccountAddressInfoRequest
	GetPK() *string
	SetPostCode(v string) *UpdateAgAccountAddressInfoRequest
	GetPostCode() *string
	SetProvince(v string) *UpdateAgAccountAddressInfoRequest
	GetProvince() *string
}

type UpdateAgAccountAddressInfoRequest struct {
	Address  *string `json:"Address,omitempty" xml:"Address,omitempty"`
	Address2 *string `json:"Address2,omitempty" xml:"Address2,omitempty"`
	AppName  *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	City     *string `json:"City,omitempty" xml:"City,omitempty"`
	Mpk      *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	PK       *string `json:"PK,omitempty" xml:"PK,omitempty"`
	PostCode *string `json:"PostCode,omitempty" xml:"PostCode,omitempty"`
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
}

func (s UpdateAgAccountAddressInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgAccountAddressInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgAccountAddressInfoRequest) GetAddress() *string {
	return s.Address
}

func (s *UpdateAgAccountAddressInfoRequest) GetAddress2() *string {
	return s.Address2
}

func (s *UpdateAgAccountAddressInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateAgAccountAddressInfoRequest) GetCity() *string {
	return s.City
}

func (s *UpdateAgAccountAddressInfoRequest) GetMpk() *string {
	return s.Mpk
}

func (s *UpdateAgAccountAddressInfoRequest) GetPK() *string {
	return s.PK
}

func (s *UpdateAgAccountAddressInfoRequest) GetPostCode() *string {
	return s.PostCode
}

func (s *UpdateAgAccountAddressInfoRequest) GetProvince() *string {
	return s.Province
}

func (s *UpdateAgAccountAddressInfoRequest) SetAddress(v string) *UpdateAgAccountAddressInfoRequest {
	s.Address = &v
	return s
}

func (s *UpdateAgAccountAddressInfoRequest) SetAddress2(v string) *UpdateAgAccountAddressInfoRequest {
	s.Address2 = &v
	return s
}

func (s *UpdateAgAccountAddressInfoRequest) SetAppName(v string) *UpdateAgAccountAddressInfoRequest {
	s.AppName = &v
	return s
}

func (s *UpdateAgAccountAddressInfoRequest) SetCity(v string) *UpdateAgAccountAddressInfoRequest {
	s.City = &v
	return s
}

func (s *UpdateAgAccountAddressInfoRequest) SetMpk(v string) *UpdateAgAccountAddressInfoRequest {
	s.Mpk = &v
	return s
}

func (s *UpdateAgAccountAddressInfoRequest) SetPK(v string) *UpdateAgAccountAddressInfoRequest {
	s.PK = &v
	return s
}

func (s *UpdateAgAccountAddressInfoRequest) SetPostCode(v string) *UpdateAgAccountAddressInfoRequest {
	s.PostCode = &v
	return s
}

func (s *UpdateAgAccountAddressInfoRequest) SetProvince(v string) *UpdateAgAccountAddressInfoRequest {
	s.Province = &v
	return s
}

func (s *UpdateAgAccountAddressInfoRequest) Validate() error {
	return dara.Validate(s)
}
