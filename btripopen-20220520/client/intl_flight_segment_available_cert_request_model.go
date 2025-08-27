// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIntlFlightSegmentAvailableCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsvName(v string) *IntlFlightSegmentAvailableCertRequest
	GetIsvName() *string
	SetLanguage(v string) *IntlFlightSegmentAvailableCertRequest
	GetLanguage() *string
	SetUserId(v string) *IntlFlightSegmentAvailableCertRequest
	GetUserId() *string
	SetUserName(v string) *IntlFlightSegmentAvailableCertRequest
	GetUserName() *string
}

type IntlFlightSegmentAvailableCertRequest struct {
	// example:
	//
	// ZJTD
	IsvName *string `json:"isv_name,omitempty" xml:"isv_name,omitempty"`
	// example:
	//
	// chinese
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// example:
	//
	// 21341234
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s IntlFlightSegmentAvailableCertRequest) String() string {
	return dara.Prettify(s)
}

func (s IntlFlightSegmentAvailableCertRequest) GoString() string {
	return s.String()
}

func (s *IntlFlightSegmentAvailableCertRequest) GetIsvName() *string {
	return s.IsvName
}

func (s *IntlFlightSegmentAvailableCertRequest) GetLanguage() *string {
	return s.Language
}

func (s *IntlFlightSegmentAvailableCertRequest) GetUserId() *string {
	return s.UserId
}

func (s *IntlFlightSegmentAvailableCertRequest) GetUserName() *string {
	return s.UserName
}

func (s *IntlFlightSegmentAvailableCertRequest) SetIsvName(v string) *IntlFlightSegmentAvailableCertRequest {
	s.IsvName = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertRequest) SetLanguage(v string) *IntlFlightSegmentAvailableCertRequest {
	s.Language = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertRequest) SetUserId(v string) *IntlFlightSegmentAvailableCertRequest {
	s.UserId = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertRequest) SetUserName(v string) *IntlFlightSegmentAvailableCertRequest {
	s.UserName = &v
	return s
}

func (s *IntlFlightSegmentAvailableCertRequest) Validate() error {
	return dara.Validate(s)
}
