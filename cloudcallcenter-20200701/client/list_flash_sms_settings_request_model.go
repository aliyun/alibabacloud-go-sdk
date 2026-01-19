// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsSettingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListFlashSmsSettingsRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListFlashSmsSettingsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListFlashSmsSettingsRequest
	GetPageSize() *int32
	SetSkillGroupIdList(v []*string) *ListFlashSmsSettingsRequest
	GetSkillGroupIdList() []*string
	SetSkillGroupName(v string) *ListFlashSmsSettingsRequest
	GetSkillGroupName() *string
}

type ListFlashSmsSettingsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize         *int32    `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	SkillGroupIdList []*string `json:"SkillGroupIdList,omitempty" xml:"SkillGroupIdList,omitempty" type:"Repeated"`
	SkillGroupName   *string   `json:"SkillGroupName,omitempty" xml:"SkillGroupName,omitempty"`
}

func (s ListFlashSmsSettingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsSettingsRequest) GoString() string {
	return s.String()
}

func (s *ListFlashSmsSettingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListFlashSmsSettingsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListFlashSmsSettingsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListFlashSmsSettingsRequest) GetSkillGroupIdList() []*string {
	return s.SkillGroupIdList
}

func (s *ListFlashSmsSettingsRequest) GetSkillGroupName() *string {
	return s.SkillGroupName
}

func (s *ListFlashSmsSettingsRequest) SetInstanceId(v string) *ListFlashSmsSettingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListFlashSmsSettingsRequest) SetPageNumber(v int32) *ListFlashSmsSettingsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListFlashSmsSettingsRequest) SetPageSize(v int32) *ListFlashSmsSettingsRequest {
	s.PageSize = &v
	return s
}

func (s *ListFlashSmsSettingsRequest) SetSkillGroupIdList(v []*string) *ListFlashSmsSettingsRequest {
	s.SkillGroupIdList = v
	return s
}

func (s *ListFlashSmsSettingsRequest) SetSkillGroupName(v string) *ListFlashSmsSettingsRequest {
	s.SkillGroupName = &v
	return s
}

func (s *ListFlashSmsSettingsRequest) Validate() error {
	return dara.Validate(s)
}
