// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTableModelInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirstLevelThemeId(v int64) *UpdateTableModelInfoRequest
	GetFirstLevelThemeId() *int64
	SetLevelId(v int64) *UpdateTableModelInfoRequest
	GetLevelId() *int64
	SetLevelType(v int32) *UpdateTableModelInfoRequest
	GetLevelType() *int32
	SetSecondLevelThemeId(v int64) *UpdateTableModelInfoRequest
	GetSecondLevelThemeId() *int64
	SetTableGuid(v string) *UpdateTableModelInfoRequest
	GetTableGuid() *string
}

type UpdateTableModelInfoRequest struct {
	// The ID of the first-level table folder.
	//
	// example:
	//
	// 101
	FirstLevelThemeId *int64 `json:"FirstLevelThemeId,omitempty" xml:"FirstLevelThemeId,omitempty"`
	// The table level ID.
	//
	// example:
	//
	// 101
	LevelId *int64 `json:"LevelId,omitempty" xml:"LevelId,omitempty"`
	// The type of the table level. Valid values: 1 and 2. The value 1 indicates the logical level. The value 2 indicates the physical level.
	//
	// example:
	//
	// 1
	LevelType *int32 `json:"LevelType,omitempty" xml:"LevelType,omitempty"`
	// The ID of the second-level table folder.
	//
	// example:
	//
	// 101
	SecondLevelThemeId *int64 `json:"SecondLevelThemeId,omitempty" xml:"SecondLevelThemeId,omitempty"`
	// The GUID of the table. Specify the GUID in the odps.{projectName}.{tableName} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.test.table1
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s UpdateTableModelInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTableModelInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateTableModelInfoRequest) GetFirstLevelThemeId() *int64 {
	return s.FirstLevelThemeId
}

func (s *UpdateTableModelInfoRequest) GetLevelId() *int64 {
	return s.LevelId
}

func (s *UpdateTableModelInfoRequest) GetLevelType() *int32 {
	return s.LevelType
}

func (s *UpdateTableModelInfoRequest) GetSecondLevelThemeId() *int64 {
	return s.SecondLevelThemeId
}

func (s *UpdateTableModelInfoRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *UpdateTableModelInfoRequest) SetFirstLevelThemeId(v int64) *UpdateTableModelInfoRequest {
	s.FirstLevelThemeId = &v
	return s
}

func (s *UpdateTableModelInfoRequest) SetLevelId(v int64) *UpdateTableModelInfoRequest {
	s.LevelId = &v
	return s
}

func (s *UpdateTableModelInfoRequest) SetLevelType(v int32) *UpdateTableModelInfoRequest {
	s.LevelType = &v
	return s
}

func (s *UpdateTableModelInfoRequest) SetSecondLevelThemeId(v int64) *UpdateTableModelInfoRequest {
	s.SecondLevelThemeId = &v
	return s
}

func (s *UpdateTableModelInfoRequest) SetTableGuid(v string) *UpdateTableModelInfoRequest {
	s.TableGuid = &v
	return s
}

func (s *UpdateTableModelInfoRequest) Validate() error {
	return dara.Validate(s)
}
