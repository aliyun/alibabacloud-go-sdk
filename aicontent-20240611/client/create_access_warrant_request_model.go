// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessWarrantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateAccessWarrantRequest
	GetAppId() *string
	SetRequestSign(v string) *CreateAccessWarrantRequest
	GetRequestSign() *string
	SetTimestamp(v string) *CreateAccessWarrantRequest
	GetTimestamp() *string
	SetUserClientIp(v string) *CreateAccessWarrantRequest
	GetUserClientIp() *string
	SetUserId(v string) *CreateAccessWarrantRequest
	GetUserId() *string
	SetWarrantAvailable(v int32) *CreateAccessWarrantRequest
	GetWarrantAvailable() *int32
}

type CreateAccessWarrantRequest struct {
	// example:
	//
	// a123
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// e32fac43df0b0b0be32fac43df0b0b0b
	RequestSign *string `json:"requestSign,omitempty" xml:"requestSign,omitempty"`
	// example:
	//
	// 1701000000
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	// example:
	//
	// 110.25.23.12
	UserClientIp *string `json:"userClientIp,omitempty" xml:"userClientIp,omitempty"`
	// example:
	//
	// sn123
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
	// example:
	//
	// 7200
	WarrantAvailable *int32 `json:"warrantAvailable,omitempty" xml:"warrantAvailable,omitempty"`
}

func (s CreateAccessWarrantRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessWarrantRequest) GoString() string {
	return s.String()
}

func (s *CreateAccessWarrantRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateAccessWarrantRequest) GetRequestSign() *string {
	return s.RequestSign
}

func (s *CreateAccessWarrantRequest) GetTimestamp() *string {
	return s.Timestamp
}

func (s *CreateAccessWarrantRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *CreateAccessWarrantRequest) GetUserId() *string {
	return s.UserId
}

func (s *CreateAccessWarrantRequest) GetWarrantAvailable() *int32 {
	return s.WarrantAvailable
}

func (s *CreateAccessWarrantRequest) SetAppId(v string) *CreateAccessWarrantRequest {
	s.AppId = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetRequestSign(v string) *CreateAccessWarrantRequest {
	s.RequestSign = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetTimestamp(v string) *CreateAccessWarrantRequest {
	s.Timestamp = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetUserClientIp(v string) *CreateAccessWarrantRequest {
	s.UserClientIp = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetUserId(v string) *CreateAccessWarrantRequest {
	s.UserId = &v
	return s
}

func (s *CreateAccessWarrantRequest) SetWarrantAvailable(v int32) *CreateAccessWarrantRequest {
	s.WarrantAvailable = &v
	return s
}

func (s *CreateAccessWarrantRequest) Validate() error {
	return dara.Validate(s)
}
