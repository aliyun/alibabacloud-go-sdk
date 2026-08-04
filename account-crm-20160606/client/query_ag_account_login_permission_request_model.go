// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgAccountLoginPermissionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgAccountType(v string) *QueryAgAccountLoginPermissionRequest
	GetAgAccountType() *string
	SetAppName(v string) *QueryAgAccountLoginPermissionRequest
	GetAppName() *string
	SetMpk(v string) *QueryAgAccountLoginPermissionRequest
	GetMpk() *string
	SetPk(v string) *QueryAgAccountLoginPermissionRequest
	GetPk() *string
}

type QueryAgAccountLoginPermissionRequest struct {
	// This parameter is required.
	AgAccountType *string `json:"AgAccountType,omitempty" xml:"AgAccountType,omitempty"`
	// This parameter is required.
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	Mpk *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	// This parameter is required.
	Pk *string `json:"Pk,omitempty" xml:"Pk,omitempty"`
}

func (s QueryAgAccountLoginPermissionRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAgAccountLoginPermissionRequest) GoString() string {
	return s.String()
}

func (s *QueryAgAccountLoginPermissionRequest) GetAgAccountType() *string {
	return s.AgAccountType
}

func (s *QueryAgAccountLoginPermissionRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryAgAccountLoginPermissionRequest) GetMpk() *string {
	return s.Mpk
}

func (s *QueryAgAccountLoginPermissionRequest) GetPk() *string {
	return s.Pk
}

func (s *QueryAgAccountLoginPermissionRequest) SetAgAccountType(v string) *QueryAgAccountLoginPermissionRequest {
	s.AgAccountType = &v
	return s
}

func (s *QueryAgAccountLoginPermissionRequest) SetAppName(v string) *QueryAgAccountLoginPermissionRequest {
	s.AppName = &v
	return s
}

func (s *QueryAgAccountLoginPermissionRequest) SetMpk(v string) *QueryAgAccountLoginPermissionRequest {
	s.Mpk = &v
	return s
}

func (s *QueryAgAccountLoginPermissionRequest) SetPk(v string) *QueryAgAccountLoginPermissionRequest {
	s.Pk = &v
	return s
}

func (s *QueryAgAccountLoginPermissionRequest) Validate() error {
	return dara.Validate(s)
}
