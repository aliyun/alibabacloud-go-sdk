// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountAddressInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressVersion(v string) *QueryAccountAddressInfoRequest
	GetAddressVersion() *string
	SetHavanaId(v string) *QueryAccountAddressInfoRequest
	GetHavanaId() *string
	SetPK(v string) *QueryAccountAddressInfoRequest
	GetPK() *string
}

type QueryAccountAddressInfoRequest struct {
	AddressVersion *string `json:"AddressVersion,omitempty" xml:"AddressVersion,omitempty"`
	HavanaId       *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	PK             *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryAccountAddressInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoRequest) GetAddressVersion() *string {
	return s.AddressVersion
}

func (s *QueryAccountAddressInfoRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryAccountAddressInfoRequest) GetPK() *string {
	return s.PK
}

func (s *QueryAccountAddressInfoRequest) SetAddressVersion(v string) *QueryAccountAddressInfoRequest {
	s.AddressVersion = &v
	return s
}

func (s *QueryAccountAddressInfoRequest) SetHavanaId(v string) *QueryAccountAddressInfoRequest {
	s.HavanaId = &v
	return s
}

func (s *QueryAccountAddressInfoRequest) SetPK(v string) *QueryAccountAddressInfoRequest {
	s.PK = &v
	return s
}

func (s *QueryAccountAddressInfoRequest) Validate() error {
	return dara.Validate(s)
}
