// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliasName(v string) *UnbindAliasRequest
	GetAliasName() *string
	SetAppKey(v int64) *UnbindAliasRequest
	GetAppKey() *int64
	SetDeviceId(v string) *UnbindAliasRequest
	GetDeviceId() *string
	SetUnbindAll(v bool) *UnbindAliasRequest
	GetUnbindAll() *bool
}

type UnbindAliasRequest struct {
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
	// example:
	//
	// true
	UnbindAll *bool `json:"UnbindAll,omitempty" xml:"UnbindAll,omitempty"`
}

func (s UnbindAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindAliasRequest) GoString() string {
	return s.String()
}

func (s *UnbindAliasRequest) GetAliasName() *string {
	return s.AliasName
}

func (s *UnbindAliasRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *UnbindAliasRequest) GetDeviceId() *string {
	return s.DeviceId
}

func (s *UnbindAliasRequest) GetUnbindAll() *bool {
	return s.UnbindAll
}

func (s *UnbindAliasRequest) SetAliasName(v string) *UnbindAliasRequest {
	s.AliasName = &v
	return s
}

func (s *UnbindAliasRequest) SetAppKey(v int64) *UnbindAliasRequest {
	s.AppKey = &v
	return s
}

func (s *UnbindAliasRequest) SetDeviceId(v string) *UnbindAliasRequest {
	s.DeviceId = &v
	return s
}

func (s *UnbindAliasRequest) SetUnbindAll(v bool) *UnbindAliasRequest {
	s.UnbindAll = &v
	return s
}

func (s *UnbindAliasRequest) Validate() error {
	return dara.Validate(s)
}
