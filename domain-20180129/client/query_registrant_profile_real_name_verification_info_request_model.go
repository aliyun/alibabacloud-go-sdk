// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRegistrantProfileRealNameVerificationInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFetchImage(v bool) *QueryRegistrantProfileRealNameVerificationInfoRequest
	GetFetchImage() *bool
	SetLang(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest
	GetLang() *string
	SetRegistrantProfileId(v int64) *QueryRegistrantProfileRealNameVerificationInfoRequest
	GetRegistrantProfileId() *int64
	SetUserClientIp(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest
	GetUserClientIp() *string
}

type QueryRegistrantProfileRealNameVerificationInfoRequest struct {
	// example:
	//
	// false
	FetchImage *bool `json:"FetchImage,omitempty" xml:"FetchImage,omitempty"`
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryRegistrantProfileRealNameVerificationInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryRegistrantProfileRealNameVerificationInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) GetFetchImage() *bool {
	return s.FetchImage
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetFetchImage(v bool) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.FetchImage = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetLang(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.Lang = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetRegistrantProfileId(v int64) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) SetUserClientIp(v string) *QueryRegistrantProfileRealNameVerificationInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryRegistrantProfileRealNameVerificationInfoRequest) Validate() error {
	return dara.Validate(s)
}
