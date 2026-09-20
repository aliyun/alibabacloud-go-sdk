// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableThemeLevelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEntity(v *GetMetaTableThemeLevelResponseBodyEntity) *GetMetaTableThemeLevelResponseBody
	GetEntity() *GetMetaTableThemeLevelResponseBodyEntity
	SetErrorCode(v string) *GetMetaTableThemeLevelResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetMetaTableThemeLevelResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetMetaTableThemeLevelResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetMetaTableThemeLevelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetaTableThemeLevelResponseBody
	GetSuccess() *bool
}

type GetMetaTableThemeLevelResponseBody struct {
	// The returned data.
	Entity *GetMetaTableThemeLevelResponseBodyEntity `json:"Entity,omitempty" xml:"Entity,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// 0
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// success
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 1324afdsfde
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetaTableThemeLevelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableThemeLevelResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaTableThemeLevelResponseBody) GetEntity() *GetMetaTableThemeLevelResponseBodyEntity {
	return s.Entity
}

func (s *GetMetaTableThemeLevelResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetMetaTableThemeLevelResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetMetaTableThemeLevelResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMetaTableThemeLevelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetaTableThemeLevelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaTableThemeLevelResponseBody) SetEntity(v *GetMetaTableThemeLevelResponseBodyEntity) *GetMetaTableThemeLevelResponseBody {
	s.Entity = v
	return s
}

func (s *GetMetaTableThemeLevelResponseBody) SetErrorCode(v string) *GetMetaTableThemeLevelResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBody) SetErrorMessage(v string) *GetMetaTableThemeLevelResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBody) SetHttpStatusCode(v int32) *GetMetaTableThemeLevelResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBody) SetRequestId(v string) *GetMetaTableThemeLevelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBody) SetSuccess(v bool) *GetMetaTableThemeLevelResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBody) Validate() error {
	if s.Entity != nil {
		if err := s.Entity.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMetaTableThemeLevelResponseBodyEntity struct {
	// The level information.
	Level []*GetMetaTableThemeLevelResponseBodyEntityLevel `json:"Level,omitempty" xml:"Level,omitempty" type:"Repeated"`
	// The theme information.
	Theme []*GetMetaTableThemeLevelResponseBodyEntityTheme `json:"Theme,omitempty" xml:"Theme,omitempty" type:"Repeated"`
}

func (s GetMetaTableThemeLevelResponseBodyEntity) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableThemeLevelResponseBodyEntity) GoString() string {
	return s.String()
}

func (s *GetMetaTableThemeLevelResponseBodyEntity) GetLevel() []*GetMetaTableThemeLevelResponseBodyEntityLevel {
	return s.Level
}

func (s *GetMetaTableThemeLevelResponseBodyEntity) GetTheme() []*GetMetaTableThemeLevelResponseBodyEntityTheme {
	return s.Theme
}

func (s *GetMetaTableThemeLevelResponseBodyEntity) SetLevel(v []*GetMetaTableThemeLevelResponseBodyEntityLevel) *GetMetaTableThemeLevelResponseBodyEntity {
	s.Level = v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntity) SetTheme(v []*GetMetaTableThemeLevelResponseBodyEntityTheme) *GetMetaTableThemeLevelResponseBodyEntity {
	s.Theme = v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntity) Validate() error {
	if s.Level != nil {
		for _, item := range s.Level {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Theme != nil {
		for _, item := range s.Theme {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMetaTableThemeLevelResponseBodyEntityLevel struct {
	// The description of the level.
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The level ID.
	//
	// example:
	//
	// 1
	LevelId *int64 `json:"LevelId,omitempty" xml:"LevelId,omitempty"`
	// The level name.
	//
	// example:
	//
	// level1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The level type. Valid values:
	//
	// - 1: logical level.
	//
	// - 2: physical level.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetMetaTableThemeLevelResponseBodyEntityLevel) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableThemeLevelResponseBodyEntityLevel) GoString() string {
	return s.String()
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) GetDescription() *string {
	return s.Description
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) GetLevelId() *int64 {
	return s.LevelId
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) GetName() *string {
	return s.Name
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) GetType() *int32 {
	return s.Type
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) SetDescription(v string) *GetMetaTableThemeLevelResponseBodyEntityLevel {
	s.Description = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) SetLevelId(v int64) *GetMetaTableThemeLevelResponseBodyEntityLevel {
	s.LevelId = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) SetName(v string) *GetMetaTableThemeLevelResponseBodyEntityLevel {
	s.Name = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) SetType(v int32) *GetMetaTableThemeLevelResponseBodyEntityLevel {
	s.Type = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntityLevel) Validate() error {
	return dara.Validate(s)
}

type GetMetaTableThemeLevelResponseBodyEntityTheme struct {
	// The theme type. Valid values:
	//
	// - 1: primary theme.
	//
	// - 2: secondary theme.
	//
	// example:
	//
	// 1
	Level *int32 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The theme name.
	//
	// example:
	//
	// theme1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the parent theme.
	//
	// example:
	//
	// 0
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The theme ID.
	//
	// example:
	//
	// 123
	ThemeId *int64 `json:"ThemeId,omitempty" xml:"ThemeId,omitempty"`
}

func (s GetMetaTableThemeLevelResponseBodyEntityTheme) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableThemeLevelResponseBodyEntityTheme) GoString() string {
	return s.String()
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) GetLevel() *int32 {
	return s.Level
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) GetName() *string {
	return s.Name
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) GetThemeId() *int64 {
	return s.ThemeId
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) SetLevel(v int32) *GetMetaTableThemeLevelResponseBodyEntityTheme {
	s.Level = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) SetName(v string) *GetMetaTableThemeLevelResponseBodyEntityTheme {
	s.Name = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) SetParentId(v int64) *GetMetaTableThemeLevelResponseBodyEntityTheme {
	s.ParentId = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) SetThemeId(v int64) *GetMetaTableThemeLevelResponseBodyEntityTheme {
	s.ThemeId = &v
	return s
}

func (s *GetMetaTableThemeLevelResponseBodyEntityTheme) Validate() error {
	return dara.Validate(s)
}
