// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStsTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthCode(v string) *GetStsTokenRequest
	GetAuthCode() *string
	SetClientId(v string) *GetStsTokenRequest
	GetClientId() *string
	SetClientIp(v string) *GetStsTokenRequest
	GetClientIp() *string
	SetClientOS(v string) *GetStsTokenRequest
	GetClientOS() *string
	SetClientVersion(v string) *GetStsTokenRequest
	GetClientVersion() *string
	SetUuid(v string) *GetStsTokenRequest
	GetUuid() *string
}

type GetStsTokenRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// e4e169bea1cc48e8afac53**********
	AuthCode *string `json:"AuthCode,omitempty" xml:"AuthCode,omitempty"`
	// example:
	//
	// eac19bef-1e45-4190-a03a-4ea74b69****
	ClientId *string `json:"ClientId,omitempty" xml:"ClientId,omitempty"`
	// example:
	//
	// 22.21.2.**
	ClientIp *string `json:"ClientIp,omitempty" xml:"ClientIp,omitempty"`
	// example:
	//
	// windows_\\"Windows 10 Enterprise\\" 10.0 (Build 14393)
	ClientOS *string `json:"ClientOS,omitempty" xml:"ClientOS,omitempty"`
	// example:
	//
	// 6.3.0-R-20231106.210000
	ClientVersion *string `json:"ClientVersion,omitempty" xml:"ClientVersion,omitempty"`
	// example:
	//
	// 2943802884B27030B6759F9132B2****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetStsTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStsTokenRequest) GoString() string {
	return s.String()
}

func (s *GetStsTokenRequest) GetAuthCode() *string {
	return s.AuthCode
}

func (s *GetStsTokenRequest) GetClientId() *string {
	return s.ClientId
}

func (s *GetStsTokenRequest) GetClientIp() *string {
	return s.ClientIp
}

func (s *GetStsTokenRequest) GetClientOS() *string {
	return s.ClientOS
}

func (s *GetStsTokenRequest) GetClientVersion() *string {
	return s.ClientVersion
}

func (s *GetStsTokenRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetStsTokenRequest) SetAuthCode(v string) *GetStsTokenRequest {
	s.AuthCode = &v
	return s
}

func (s *GetStsTokenRequest) SetClientId(v string) *GetStsTokenRequest {
	s.ClientId = &v
	return s
}

func (s *GetStsTokenRequest) SetClientIp(v string) *GetStsTokenRequest {
	s.ClientIp = &v
	return s
}

func (s *GetStsTokenRequest) SetClientOS(v string) *GetStsTokenRequest {
	s.ClientOS = &v
	return s
}

func (s *GetStsTokenRequest) SetClientVersion(v string) *GetStsTokenRequest {
	s.ClientVersion = &v
	return s
}

func (s *GetStsTokenRequest) SetUuid(v string) *GetStsTokenRequest {
	s.Uuid = &v
	return s
}

func (s *GetStsTokenRequest) Validate() error {
	return dara.Validate(s)
}
