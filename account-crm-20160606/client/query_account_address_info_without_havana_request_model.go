// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAccountAddressInfoWithoutHavanaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressVersion(v string) *QueryAccountAddressInfoWithoutHavanaRequest
	GetAddressVersion() *string
	SetHavanaId(v string) *QueryAccountAddressInfoWithoutHavanaRequest
	GetHavanaId() *string
	SetPK(v string) *QueryAccountAddressInfoWithoutHavanaRequest
	GetPK() *string
}

type QueryAccountAddressInfoWithoutHavanaRequest struct {
	AddressVersion *string `json:"AddressVersion,omitempty" xml:"AddressVersion,omitempty"`
	HavanaId       *string `json:"HavanaId,omitempty" xml:"HavanaId,omitempty"`
	PK             *string `json:"PK,omitempty" xml:"PK,omitempty"`
}

func (s QueryAccountAddressInfoWithoutHavanaRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAccountAddressInfoWithoutHavanaRequest) GoString() string {
	return s.String()
}

func (s *QueryAccountAddressInfoWithoutHavanaRequest) GetAddressVersion() *string {
	return s.AddressVersion
}

func (s *QueryAccountAddressInfoWithoutHavanaRequest) GetHavanaId() *string {
	return s.HavanaId
}

func (s *QueryAccountAddressInfoWithoutHavanaRequest) GetPK() *string {
	return s.PK
}

func (s *QueryAccountAddressInfoWithoutHavanaRequest) SetAddressVersion(v string) *QueryAccountAddressInfoWithoutHavanaRequest {
	s.AddressVersion = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaRequest) SetHavanaId(v string) *QueryAccountAddressInfoWithoutHavanaRequest {
	s.HavanaId = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaRequest) SetPK(v string) *QueryAccountAddressInfoWithoutHavanaRequest {
	s.PK = &v
	return s
}

func (s *QueryAccountAddressInfoWithoutHavanaRequest) Validate() error {
	return dara.Validate(s)
}
