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
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// aliyun-cloudsiem-data-173326*******-cn-hangzhou
	LogProjectName *string `json:"LogProjectName,omitempty" xml:"LogProjectName,omitempty"`
	// The ID of the log storage region.
	//
	// example:
	//
	// cn-hangzhou
	LogRegionId *string `json:"LogRegionId,omitempty" xml:"LogRegionId,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// ssglauncher-log
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The user ID for data access.
	//
	// example:
	//
	// 173326*******
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// The region of the Data Management Center for threat analysis. Select a region based on the location of your assets. Valid values:
	//
	// - cn-hangzhou: The Chinese mainland.
	//
	// - ap-southeast-1: Regions outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. This lets an administrator switch to the member\\"s perspective.
	//
	// example:
	//
	// 173326*******
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
