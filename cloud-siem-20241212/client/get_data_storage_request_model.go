// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataStorageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *GetDataStorageRequest
	GetLang() *string
	SetRegionId(v string) *GetDataStorageRequest
	GetRegionId() *string
	SetRoleFor(v int64) *GetDataStorageRequest
	GetRoleFor() *int64
}

type GetDataStorageRequest struct {
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

func (s GetDataStorageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataStorageRequest) GoString() string {
	return s.String()
}

func (s *GetDataStorageRequest) GetLang() *string {
	return s.Lang
}

func (s *GetDataStorageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDataStorageRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *GetDataStorageRequest) SetLang(v string) *GetDataStorageRequest {
	s.Lang = &v
	return s
}

func (s *GetDataStorageRequest) SetRegionId(v string) *GetDataStorageRequest {
	s.RegionId = &v
	return s
}

func (s *GetDataStorageRequest) SetRoleFor(v int64) *GetDataStorageRequest {
	s.RoleFor = &v
	return s
}

func (s *GetDataStorageRequest) Validate() error {
	return dara.Validate(s)
}
