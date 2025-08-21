// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *BindAliasRequest
	GetAliasName() *string
	SetAppKey(v int64) *BindAliasRequest
	GetAppKey() *int64
	SetDeviceId(v string) *BindAliasRequest
	GetDeviceId() *string
}

type BindAliasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test_alias
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 23267207
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e2ba19de97604f55b16557673****
	DeviceId *string `json:"DeviceId,omitempty" xml:"DeviceId,omitempty"`
}

func (s BindAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s BindAliasRequest) GoString() string {
	return s.String()
}

func (s *BindAliasRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *BindAliasRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *BindAliasRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *BindAliasRequest) SetAliasName(v string) *BindAliasRequest {
	s.AliasName = &v
	return s
}

func (s *BindAliasRequest) SetAppKey(v int64) *BindAliasRequest {
	s.AppKey = &v
	return s
}

func (s *BindAliasRequest) SetDeviceId(v string) *BindAliasRequest {
	s.DeviceId = &v
	return s
}

func (s *BindAliasRequest) Validate() error {
	return dara.Validate(s)
}
