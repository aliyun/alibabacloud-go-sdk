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
