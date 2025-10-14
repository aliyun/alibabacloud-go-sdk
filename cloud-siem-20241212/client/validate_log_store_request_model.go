// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateLogStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *ValidateLogStoreRequest
	GetLang() *string
	SetLogProjectName(v string) *ValidateLogStoreRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *ValidateLogStoreRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *ValidateLogStoreRequest
	GetLogStoreName() *string
	SetLogUserId(v int64) *ValidateLogStoreRequest
	GetLogUserId() *int64
	SetRegionId(v string) *ValidateLogStoreRequest
	GetRegionId() *string
	SetRoleFor(v int64) *ValidateLogStoreRequest
	GetRoleFor() *int64
}

type ValidateLogStoreRequest struct {
	// example:
	//
	// zh。
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou。
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// example:
	//
	// cn-hangzhou。
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// example:
	//
	// ssglauncher-log。
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

func (s ValidateLogStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateLogStoreRequest) GoString() string {
	return s.String()
}

func (s *ValidateLogStoreRequest) GetLang() *string {
	return s.Lang
}

func (s *ValidateLogStoreRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *ValidateLogStoreRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *ValidateLogStoreRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *ValidateLogStoreRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *ValidateLogStoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ValidateLogStoreRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *ValidateLogStoreRequest) SetLang(v string) *ValidateLogStoreRequest {
	s.Lang = &v
	return s
}

func (s *ValidateLogStoreRequest) SetLogProjectName(v string) *ValidateLogStoreRequest {
	s.LogProjectName = &v
	return s
}

func (s *ValidateLogStoreRequest) SetLogRegionId(v string) *ValidateLogStoreRequest {
	s.LogRegionId = &v
	return s
}

func (s *ValidateLogStoreRequest) SetLogStoreName(v string) *ValidateLogStoreRequest {
	s.LogStoreName = &v
	return s
}

func (s *ValidateLogStoreRequest) SetLogUserId(v int64) *ValidateLogStoreRequest {
	s.LogUserId = &v
	return s
}

func (s *ValidateLogStoreRequest) SetRegionId(v string) *ValidateLogStoreRequest {
	s.RegionId = &v
	return s
}

func (s *ValidateLogStoreRequest) SetRoleFor(v int64) *ValidateLogStoreRequest {
	s.RoleFor = &v
	return s
}

func (s *ValidateLogStoreRequest) Validate() error {
	return dara.Validate(s)
}
