// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecureSkillIdentitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSecureSkillIdentitiesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSecureSkillIdentitiesRequest
	GetPageSize() *int32
	SetSkillChannel(v string) *ListSecureSkillIdentitiesRequest
	GetSkillChannel() *string
}

type ListSecureSkillIdentitiesRequest struct {
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The skill channel. Valid values:
	//
	// - ENTERPRISE: Enterprise edition.
	//
	// - BUSINESS: Business edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// ENTERPRISE
	SkillChannel *string `json:"SkillChannel,omitempty" xml:"SkillChannel,omitempty"`
}

func (s ListSecureSkillIdentitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecureSkillIdentitiesRequest) GoString() string {
	return s.String()
}

func (s *ListSecureSkillIdentitiesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecureSkillIdentitiesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecureSkillIdentitiesRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *ListSecureSkillIdentitiesRequest) SetPageNumber(v int32) *ListSecureSkillIdentitiesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecureSkillIdentitiesRequest) SetPageSize(v int32) *ListSecureSkillIdentitiesRequest {
	s.PageSize = &v
	return s
}

func (s *ListSecureSkillIdentitiesRequest) SetSkillChannel(v string) *ListSecureSkillIdentitiesRequest {
	s.SkillChannel = &v
	return s
}

func (s *ListSecureSkillIdentitiesRequest) Validate() error {
	return dara.Validate(s)
}
