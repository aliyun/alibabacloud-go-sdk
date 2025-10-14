// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DeleteLogStoreRequest
	GetLang() *string
	SetLogProjectName(v string) *DeleteLogStoreRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *DeleteLogStoreRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *DeleteLogStoreRequest
	GetLogStoreName() *string
	SetLogUserId(v int64) *DeleteLogStoreRequest
	GetLogUserId() *int64
	SetRegionId(v string) *DeleteLogStoreRequest
	GetRegionId() *string
	SetRoleFor(v int64) *DeleteLogStoreRequest
	GetRoleFor() *int64
}

type DeleteLogStoreRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// example:
	//
	// rds-logstore。
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// 173326*******。
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// example:
	//
	// cn-hangzhou。
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 173326*******。
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s DeleteLogStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogStoreRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogStoreRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteLogStoreRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *DeleteLogStoreRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *DeleteLogStoreRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DeleteLogStoreRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *DeleteLogStoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLogStoreRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *DeleteLogStoreRequest) SetLang(v string) *DeleteLogStoreRequest {
	s.Lang = &v
	return s
}

func (s *DeleteLogStoreRequest) SetLogProjectName(v string) *DeleteLogStoreRequest {
	s.LogProjectName = &v
	return s
}

func (s *DeleteLogStoreRequest) SetLogRegionId(v string) *DeleteLogStoreRequest {
	s.LogRegionId = &v
	return s
}

func (s *DeleteLogStoreRequest) SetLogStoreName(v string) *DeleteLogStoreRequest {
	s.LogStoreName = &v
	return s
}

func (s *DeleteLogStoreRequest) SetLogUserId(v int64) *DeleteLogStoreRequest {
	s.LogUserId = &v
	return s
}

func (s *DeleteLogStoreRequest) SetRegionId(v string) *DeleteLogStoreRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLogStoreRequest) SetRoleFor(v int64) *DeleteLogStoreRequest {
	s.RoleFor = &v
	return s
}

func (s *DeleteLogStoreRequest) Validate() error {
	return dara.Validate(s)
}
