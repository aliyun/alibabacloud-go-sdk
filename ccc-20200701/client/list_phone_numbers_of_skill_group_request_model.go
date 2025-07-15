// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPhoneNumbersOfSkillGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActive(v bool) *ListPhoneNumbersOfSkillGroupRequest
	GetActive() *bool
	SetInstanceId(v string) *ListPhoneNumbersOfSkillGroupRequest
	GetInstanceId() *string
	SetIsMember(v bool) *ListPhoneNumbersOfSkillGroupRequest
	GetIsMember() *bool
	SetPageNumber(v int32) *ListPhoneNumbersOfSkillGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListPhoneNumbersOfSkillGroupRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListPhoneNumbersOfSkillGroupRequest
	GetSearchPattern() *string
	SetSkillGroupId(v string) *ListPhoneNumbersOfSkillGroupRequest
	GetSkillGroupId() *string
}

type ListPhoneNumbersOfSkillGroupRequest struct {
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
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
	// true
	IsMember *bool `json:"IsMember,omitempty" xml:"IsMember,omitempty"`
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
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0833
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// skillgroup@ccc-test
	SkillGroupId *string `json:"SkillGroupId,omitempty" xml:"SkillGroupId,omitempty"`
}

func (s ListPhoneNumbersOfSkillGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPhoneNumbersOfSkillGroupRequest) GoString() string {
	return s.String()
}

func (s *ListPhoneNumbersOfSkillGroupRequest) GetActive() *bool {
	return s.Active
}

func (s *ListPhoneNumbersOfSkillGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListPhoneNumbersOfSkillGroupRequest) GetIsMember() *bool {
	return s.IsMember
}

func (s *ListPhoneNumbersOfSkillGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPhoneNumbersOfSkillGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPhoneNumbersOfSkillGroupRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListPhoneNumbersOfSkillGroupRequest) GetSkillGroupId() *string {
	return s.SkillGroupId
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetActive(v bool) *ListPhoneNumbersOfSkillGroupRequest {
	s.Active = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetInstanceId(v string) *ListPhoneNumbersOfSkillGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetIsMember(v bool) *ListPhoneNumbersOfSkillGroupRequest {
	s.IsMember = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetPageNumber(v int32) *ListPhoneNumbersOfSkillGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetPageSize(v int32) *ListPhoneNumbersOfSkillGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetSearchPattern(v string) *ListPhoneNumbersOfSkillGroupRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) SetSkillGroupId(v string) *ListPhoneNumbersOfSkillGroupRequest {
	s.SkillGroupId = &v
	return s
}

func (s *ListPhoneNumbersOfSkillGroupRequest) Validate() error {
	return dara.Validate(s)
}
