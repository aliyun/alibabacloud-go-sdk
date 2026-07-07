// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillAuthedIdentitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSkillAuthedIdentitiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSkillAuthedIdentitiesRequest
	GetPageSize() *int32
	SetSkillChannel(v string) *ListSkillAuthedIdentitiesRequest
	GetSkillChannel() *string
	SetSkillId(v string) *ListSkillAuthedIdentitiesRequest
	GetSkillId() *string
}

type ListSkillAuthedIdentitiesRequest struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The skill channel. Valid values:
	//
	// - ENTERPRISE: enterprise edition
	//
	// - BUSINESS: business edition
	//
	// This parameter is required.
	//
	// example:
	//
	// ENTERPRISE
	SkillChannel *string `json:"SkillChannel,omitempty" xml:"SkillChannel,omitempty"`
	// The unique identifier of the skill.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-04rj8mzqj1fu****
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
}

func (s ListSkillAuthedIdentitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillAuthedIdentitiesRequest) GoString() string {
	return s.String()
}

func (s *ListSkillAuthedIdentitiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillAuthedIdentitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillAuthedIdentitiesRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *ListSkillAuthedIdentitiesRequest) GetSkillId() *string {
	return s.SkillId
}

func (s *ListSkillAuthedIdentitiesRequest) SetPageNumber(v int32) *ListSkillAuthedIdentitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillAuthedIdentitiesRequest) SetPageSize(v int32) *ListSkillAuthedIdentitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillAuthedIdentitiesRequest) SetSkillChannel(v string) *ListSkillAuthedIdentitiesRequest {
	s.SkillChannel = &v
	return s
}

func (s *ListSkillAuthedIdentitiesRequest) SetSkillId(v string) *ListSkillAuthedIdentitiesRequest {
	s.SkillId = &v
	return s
}

func (s *ListSkillAuthedIdentitiesRequest) Validate() error {
	return dara.Validate(s)
}
