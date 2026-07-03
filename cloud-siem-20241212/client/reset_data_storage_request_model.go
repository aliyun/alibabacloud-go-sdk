// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDataStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ResetDataStorageRequest
	GetLang() *string
	SetRegionId(v string) *ResetDataStorageRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ResetDataStorageRequest
	GetRoleFor() *int64
}

type ResetDataStorageRequest struct {
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
	// The region of the Data Management center. Select a region for the Data Management center based on the region where your assets are located. Valid values:
	//
	// - cn-hangzhou: The assets are in the Chinese mainland.
	//
	// - ap-southeast-1: The assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator uses this ID to switch to the member\\"s perspective.
	//
	// example:
	//
	// 113091674488****
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s ResetDataStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetDataStorageRequest) GoString() string {
	return s.String()
}

func (s *ResetDataStorageRequest) GetLang() *string {
	return s.Lang
}

func (s *ResetDataStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ResetDataStorageRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ResetDataStorageRequest) SetLang(v string) *ResetDataStorageRequest {
	s.Lang = &v
	return s
}

func (s *ResetDataStorageRequest) SetRegionId(v string) *ResetDataStorageRequest {
	s.RegionId = &v
	return s
}

func (s *ResetDataStorageRequest) SetRoleFor(v int64) *ResetDataStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *ResetDataStorageRequest) Validate() error {
	return dara.Validate(s)
}
