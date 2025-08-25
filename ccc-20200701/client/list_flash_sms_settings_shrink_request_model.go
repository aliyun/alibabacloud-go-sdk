// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsSettingsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFlashSmsSettingsShrinkRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListFlashSmsSettingsShrinkRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFlashSmsSettingsShrinkRequest
	GetPageSize() *int32
	SetSkillGroupIdListShrink(v string) *ListFlashSmsSettingsShrinkRequest
	GetSkillGroupIdListShrink() *string
	SetSkillGroupName(v string) *ListFlashSmsSettingsShrinkRequest
	GetSkillGroupName() *string
}

type ListFlashSmsSettingsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize               *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SkillGroupIdListShrink *string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty"`
	SkillGroupName         *string `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s ListFlashSmsSettingsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsSettingsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListFlashSmsSettingsShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsSettingsShrinkRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsSettingsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsSettingsShrinkRequest) GetSkillGroupIdListShrink() *string {
	return s.SkillGroupIdListShrink
}

func (s *ListFlashSmsSettingsShrinkRequest) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListFlashSmsSettingsShrinkRequest) SetInstanceId(v string) *ListFlashSmsSettingsShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsSettingsShrinkRequest) SetPageNumber(v int32) *ListFlashSmsSettingsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsSettingsShrinkRequest) SetPageSize(v int32) *ListFlashSmsSettingsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsSettingsShrinkRequest) SetSkillGroupIdListShrink(v string) *ListFlashSmsSettingsShrinkRequest {
	s.SkillGroupIdListShrink = &v
	return s
}

func (s *ListFlashSmsSettingsShrinkRequest) SetSkillGroupName(v string) *ListFlashSmsSettingsShrinkRequest {
	s.SkillGroupName = &v
	return s
}

func (s *ListFlashSmsSettingsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
