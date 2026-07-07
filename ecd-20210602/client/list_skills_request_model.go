// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListSkillsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSkillsRequest
	GetPageSize() *int32
	SetSkillChannel(v string) *ListSkillsRequest
	GetSkillChannel() *string
	SetSkillIds(v []*string) *ListSkillsRequest
	GetSkillIds() []*string
	SetSupplierType(v string) *ListSkillsRequest
	GetSupplierType() *string
}

type ListSkillsRequest struct {
	// The page number of the current page in a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The maximum number of rows per page in a paged query. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The skill channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUSINESS
	SkillChannel *string `json:"SkillChannel,omitempty" xml:"SkillChannel,omitempty"`
	// The list of skill IDs.
	SkillIds []*string `json:"SkillIds,omitempty" xml:"SkillIds,omitempty" type:"Repeated"`
	// The supply type.
	//
	// example:
	//
	// WUYING
	SupplierType *string `json:"SupplierType,omitempty" xml:"SupplierType,omitempty"`
}

func (s ListSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillsRequest) GetSkillChannel() *string {
	return s.SkillChannel
}

func (s *ListSkillsRequest) GetSkillIds() []*string {
	return s.SkillIds
}

func (s *ListSkillsRequest) GetSupplierType() *string {
	return s.SupplierType
}

func (s *ListSkillsRequest) SetPageNumber(v int32) *ListSkillsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillsRequest) SetPageSize(v int32) *ListSkillsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillsRequest) SetSkillChannel(v string) *ListSkillsRequest {
	s.SkillChannel = &v
	return s
}

func (s *ListSkillsRequest) SetSkillIds(v []*string) *ListSkillsRequest {
	s.SkillIds = v
	return s
}

func (s *ListSkillsRequest) SetSupplierType(v string) *ListSkillsRequest {
	s.SupplierType = &v
	return s
}

func (s *ListSkillsRequest) Validate() error {
	return dara.Validate(s)
}
