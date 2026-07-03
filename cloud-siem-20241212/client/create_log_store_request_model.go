// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogStoreRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateLogStoreRequest
	GetLang() *string
	SetLogProjectName(v string) *CreateLogStoreRequest
	GetLogProjectName() *string
	SetLogRegionId(v string) *CreateLogStoreRequest
	GetLogRegionId() *string
	SetLogStoreName(v string) *CreateLogStoreRequest
	GetLogStoreName() *string
	SetLogUserId(v int64) *CreateLogStoreRequest
	GetLogUserId() *int64
	SetRegionId(v string) *CreateLogStoreRequest
	GetRegionId() *string
	SetRoleFor(v int64) *CreateLogStoreRequest
	GetRoleFor() *int64
}

type CreateLogStoreRequest struct {
	// The language of the response messages. Valid values:
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
	// aliyun-cloudsiem-channel-173326*******-cn-hangzhou
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
	// logstoreqykug
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The user ID for data access.
	//
	// example:
	//
	// 173326*******
	LogUserId *int64 `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// The region of the Data Management Center for threat analysis. Select a region for the Management Center based on the region of your assets. Valid values:
	//
	// - cn-hangzhou: Your assets are in the Chinese mainland.
	//
	// - ap-southeast-1: Your assets are in a region outside China.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The user ID of a member. An administrator can specify this parameter to switch to the perspective of the member.
	//
	// example:
	//
	// 173326*******
	RoleFor *int64 `json:"RoleFor,omitempty" xml:"RoleFor,omitempty"`
}

func (s CreateLogStoreRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLogStoreRequest) GoString() string {
	return s.String()
}

func (s *CreateLogStoreRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateLogStoreRequest) GetLogProjectName() *string {
	return s.LogProjectName
}

func (s *CreateLogStoreRequest) GetLogRegionId() *string {
	return s.LogRegionId
}

func (s *CreateLogStoreRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *CreateLogStoreRequest) GetLogUserId() *int64 {
	return s.LogUserId
}

func (s *CreateLogStoreRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLogStoreRequest) GetRoleFor() *int64 {
	return s.RoleFor
}

func (s *CreateLogStoreRequest) SetLang(v string) *CreateLogStoreRequest {
	s.Lang = &v
	return s
}

func (s *CreateLogStoreRequest) SetLogProjectName(v string) *CreateLogStoreRequest {
	s.LogProjectName = &v
	return s
}

func (s *CreateLogStoreRequest) SetLogRegionId(v string) *CreateLogStoreRequest {
	s.LogRegionId = &v
	return s
}

func (s *CreateLogStoreRequest) SetLogStoreName(v string) *CreateLogStoreRequest {
	s.LogStoreName = &v
	return s
}

func (s *CreateLogStoreRequest) SetLogUserId(v int64) *CreateLogStoreRequest {
	s.LogUserId = &v
	return s
}

func (s *CreateLogStoreRequest) SetRegionId(v string) *CreateLogStoreRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLogStoreRequest) SetRoleFor(v int64) *CreateLogStoreRequest {
	s.RoleFor = &v
	return s
}

func (s *CreateLogStoreRequest) Validate() error {
	return dara.Validate(s)
}
