// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistrantProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteRegistrantProfileRequest
	GetLang() *string
	SetRegistrantProfileId(v int64) *DeleteRegistrantProfileRequest
	GetRegistrantProfileId() *int64
	SetUserClientIp(v string) *DeleteRegistrantProfileRequest
	GetUserClientIp() *string
}

type DeleteRegistrantProfileRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3600000
	RegistrantProfileId *int64 `json:"RegistrantProfileId,omitempty" xml:"RegistrantProfileId,omitempty"`
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DeleteRegistrantProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistrantProfileRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegistrantProfileRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteRegistrantProfileRequest) GetRegistrantProfileId() *int64 {
	return s.RegistrantProfileId
}

func (s *DeleteRegistrantProfileRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DeleteRegistrantProfileRequest) SetLang(v string) *DeleteRegistrantProfileRequest {
	s.Lang = &v
	return s
}

func (s *DeleteRegistrantProfileRequest) SetRegistrantProfileId(v int64) *DeleteRegistrantProfileRequest {
	s.RegistrantProfileId = &v
	return s
}

func (s *DeleteRegistrantProfileRequest) SetUserClientIp(v string) *DeleteRegistrantProfileRequest {
	s.UserClientIp = &v
	return s
}

func (s *DeleteRegistrantProfileRequest) Validate() error {
	return dara.Validate(s)
}
