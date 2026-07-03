// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataStorageTtlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateDataStorageTtlRequest
	GetLang() *string
	SetLogStoreColdTtl(v string) *UpdateDataStorageTtlRequest
	GetLogStoreColdTtl() *string
	SetLogStoreHotTtl(v string) *UpdateDataStorageTtlRequest
	GetLogStoreHotTtl() *string
	SetLogStoreName(v string) *UpdateDataStorageTtlRequest
	GetLogStoreName() *string
	SetLogStoreTtl(v string) *UpdateDataStorageTtlRequest
	GetLogStoreTtl() *string
	SetRegionId(v string) *UpdateDataStorageTtlRequest
	GetRegionId() *string
	SetRoleFor(v int64) *UpdateDataStorageTtlRequest
	GetRoleFor() *int64
}

type UpdateDataStorageTtlRequest struct {
	// The language of the response. Valid values:
	//
	// - **zh*	- (default): Chinese.
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The storage duration for cold storage in Simple Log Service. This setting is not yet available.
	//
	// example:
	//
	// 0
	LogStoreColdTtl *string `json:"LogStoreColdTtl,omitempty" xml:"LogStoreColdTtl,omitempty"`
	// The storage duration for hot storage in Simple Log Service.
	//
	// example:
	//
	// 180
	LogStoreHotTtl *string `json:"LogStoreHotTtl,omitempty" xml:"LogStoreHotTtl,omitempty"`
	// The name of the Logstore for threat analysis.
	//
	// example:
	//
	// network-activity
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The storage duration of the Logstore.
	//
	// example:
	//
	// 180
	LogStoreTtl *string `json:"LogStoreTtl,omitempty" xml:"LogStoreTtl,omitempty"`
	// The region of the Data Management Center for threat analysis. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: The Chinese mainland.
	//
	// - ap-southeast-1: Regions outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of the member whose perspective the administrator assumes.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s UpdateDataStorageTtlRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataStorageTtlRequest) GoString() string {
	return s.String()
}

func (s *UpdateDataStorageTtlRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDataStorageTtlRequest) GetLogStoreColdTtl() *string {
	return s.LogStoreColdTtl
}

func (s *UpdateDataStorageTtlRequest) GetLogStoreHotTtl() *string {
	return s.LogStoreHotTtl
}

func (s *UpdateDataStorageTtlRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *UpdateDataStorageTtlRequest) GetLogStoreTtl() *string {
	return s.LogStoreTtl
}

func (s *UpdateDataStorageTtlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateDataStorageTtlRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *UpdateDataStorageTtlRequest) SetLang(v string) *UpdateDataStorageTtlRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDataStorageTtlRequest) SetLogStoreColdTtl(v string) *UpdateDataStorageTtlRequest {
	s.LogStoreColdTtl = &v
	return s
}

func (s *UpdateDataStorageTtlRequest) SetLogStoreHotTtl(v string) *UpdateDataStorageTtlRequest {
	s.LogStoreHotTtl = &v
	return s
}

func (s *UpdateDataStorageTtlRequest) SetLogStoreName(v string) *UpdateDataStorageTtlRequest {
	s.LogStoreName = &v
	return s
}

func (s *UpdateDataStorageTtlRequest) SetLogStoreTtl(v string) *UpdateDataStorageTtlRequest {
	s.LogStoreTtl = &v
	return s
}

func (s *UpdateDataStorageTtlRequest) SetRegionId(v string) *UpdateDataStorageTtlRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateDataStorageTtlRequest) SetRoleFor(v int64) *UpdateDataStorageTtlRequest {
	s.RoleFor = &v
	return s
}

func (s *UpdateDataStorageTtlRequest) Validate() error {
	return dara.Validate(s)
}
