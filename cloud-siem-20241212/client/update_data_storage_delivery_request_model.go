// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageDeliveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateDataStorageDeliveryRequest
	GetLang() *string
	SetLogCode(v string) *UpdateDataStorageDeliveryRequest
	GetLogCode() *string
	SetLogDeliveryStatus(v string) *UpdateDataStorageDeliveryRequest
	GetLogDeliveryStatus() *string
	SetRegionId(v string) *UpdateDataStorageDeliveryRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataStorageDeliveryRequest
	GetRoleFor() *int64
}

type UpdateDataStorageDeliveryRequest struct {
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// aegis-log-login
	LogCode *string `json:"LogCode,omitempty" xml:"LogCode,omitempty"`
	// example:
	//
	// enable
	LogDeliveryStatus *string `json:"LogDeliveryStatus,omitempty" xml:"LogDeliveryStatus,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataStorageDeliveryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageDeliveryRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageDeliveryRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataStorageDeliveryRequest) GetLogCode() *string {
	return s.LogCode
}

func (s *UpdateDataStorageDeliveryRequest) GetLogDeliveryStatus() *string {
	return s.LogDeliveryStatus
}

func (s *UpdateDataStorageDeliveryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataStorageDeliveryRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataStorageDeliveryRequest) SetLang(v string) *UpdateDataStorageDeliveryRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetLogCode(v string) *UpdateDataStorageDeliveryRequest {
	s.LogCode = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetLogDeliveryStatus(v string) *UpdateDataStorageDeliveryRequest {
	s.LogDeliveryStatus = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetRegionId(v string) *UpdateDataStorageDeliveryRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) SetRoleFor(v int64) *UpdateDataStorageDeliveryRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataStorageDeliveryRequest) Validate() error {
	return dara.Validate(s)
}
