// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryUserByMobileAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMobileType(v string) *QueryUserByMobileAccountRequest
	GetMobileType() *string
	SetMobileUserId(v string) *QueryUserByMobileAccountRequest
	GetMobileUserId() *string
}

type QueryUserByMobileAccountRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ding
	MobileType *string `json:"MobileType,omitempty" xml:"MobileType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sasda
	MobileUserId *string `json:"MobileUserId,omitempty" xml:"MobileUserId,omitempty"`
}

func (s QueryUserByMobileAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryUserByMobileAccountRequest) GoString() string {
	return s.String()
}

func (s *QueryUserByMobileAccountRequest) GetMobileType() *string {
	return s.MobileType
}

func (s *QueryUserByMobileAccountRequest) GetMobileUserId() *string {
	return s.MobileUserId
}

func (s *QueryUserByMobileAccountRequest) SetMobileType(v string) *QueryUserByMobileAccountRequest {
	s.MobileType = &v
	return s
}

func (s *QueryUserByMobileAccountRequest) SetMobileUserId(v string) *QueryUserByMobileAccountRequest {
	s.MobileUserId = &v
	return s
}

func (s *QueryUserByMobileAccountRequest) Validate() error {
	return dara.Validate(s)
}
