// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataStorageRegionId(v string) *UpdateDataStorageRequest
	GetDataStorageRegionId() *string
	SetDeliveryStatus(v string) *UpdateDataStorageRequest
	GetDeliveryStatus() *string
	SetLang(v string) *UpdateDataStorageRequest
	GetLang() *string
	SetRegionId(v string) *UpdateDataStorageRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataStorageRequest
	GetRoleFor() *int64
}

type UpdateDataStorageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	DataStorageRegionId *string `json:"DataStorageRegionId,omitempty" xml:"DataStorageRegionId,omitempty"`
	// example:
	//
	// enable
	DeliveryStatus *string `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty"`
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageRequest) GetDataStorageRegionId() *string {
	return s.DataStorageRegionId
}

func (s *UpdateDataStorageRequest) GetDeliveryStatus() *string {
	return s.DeliveryStatus
}

func (s *UpdateDataStorageRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataStorageRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataStorageRequest) SetDataStorageRegionId(v string) *UpdateDataStorageRequest {
	s.DataStorageRegionId = &v
	return s
}

func (s *UpdateDataStorageRequest) SetDeliveryStatus(v string) *UpdateDataStorageRequest {
	s.DeliveryStatus = &v
	return s
}

func (s *UpdateDataStorageRequest) SetLang(v string) *UpdateDataStorageRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataStorageRequest) SetRegionId(v string) *UpdateDataStorageRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataStorageRequest) SetRoleFor(v int64) *UpdateDataStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataStorageRequest) Validate() error {
	return dara.Validate(s)
}
