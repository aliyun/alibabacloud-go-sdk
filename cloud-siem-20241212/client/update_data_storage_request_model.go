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
	// The log storage region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	DataStorageRegionId *string `json:"DataStorageRegionId,omitempty" xml:"DataStorageRegionId,omitempty"`
	// The global switch for log delivery in Log Management. This parameter is not yet available. Valid values:
	//
	// - enable: Enables global delivery.
	//
	// - disable: Disables global delivery.
	//
	// example:
	//
	// enable
	DeliveryStatus *string `json:"DeliveryStatus,omitempty" xml:"DeliveryStatus,omitempty"`
	// The language of the response message. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The region where the Data Management center for threat analysis is located. This region must be the same as the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: The assets are in the Chinese mainland.
	//
	// - ap-southeast-1: The assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can specify this parameter to switch to the perspective of the member.
	//
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
