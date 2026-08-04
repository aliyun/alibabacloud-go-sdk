// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAsyncCreateAgAccountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoginEmail(v string) *AsyncCreateAgAccountRequest
	GetLoginEmail() *string
	SetMaserAccountInfo(v string) *AsyncCreateAgAccountRequest
	GetMaserAccountInfo() *string
	SetMpk(v string) *AsyncCreateAgAccountRequest
	GetMpk() *string
}

type AsyncCreateAgAccountRequest struct {
	// This parameter is required.
	LoginEmail *string `json:"LoginEmail,omitempty" xml:"LoginEmail,omitempty"`
	// This parameter is required.
	MaserAccountInfo *string `json:"MaserAccountInfo,omitempty" xml:"MaserAccountInfo,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
}

func (s AsyncCreateAgAccountRequest) String() string {
	return dara.Prettify(s)
}

func (s AsyncCreateAgAccountRequest) GoString() string {
	return s.String()
}

func (s *AsyncCreateAgAccountRequest) GetLoginEmail() *string {
	return s.LoginEmail
}

func (s *AsyncCreateAgAccountRequest) GetMaserAccountInfo() *string {
	return s.MaserAccountInfo
}

func (s *AsyncCreateAgAccountRequest) GetMpk() *string {
	return s.Mpk
}

func (s *AsyncCreateAgAccountRequest) SetLoginEmail(v string) *AsyncCreateAgAccountRequest {
	s.LoginEmail = &v
	return s
}

func (s *AsyncCreateAgAccountRequest) SetMaserAccountInfo(v string) *AsyncCreateAgAccountRequest {
	s.MaserAccountInfo = &v
	return s
}

func (s *AsyncCreateAgAccountRequest) SetMpk(v string) *AsyncCreateAgAccountRequest {
	s.Mpk = &v
	return s
}

func (s *AsyncCreateAgAccountRequest) Validate() error {
	return dara.Validate(s)
}
