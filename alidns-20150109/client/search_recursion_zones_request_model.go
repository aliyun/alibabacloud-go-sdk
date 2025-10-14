// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRecursionZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *SearchRecursionZonesRequest
	GetDirection() *string
	SetEffectiveScopes(v []*SearchRecursionZonesRequestEffectiveScopes) *SearchRecursionZonesRequest
	GetEffectiveScopes() []*SearchRecursionZonesRequestEffectiveScopes
	SetMaxResults(v int32) *SearchRecursionZonesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchRecursionZonesRequest
	GetNextToken() *string
	SetOrderBy(v string) *SearchRecursionZonesRequest
	GetOrderBy() *string
	SetPageNumber(v int32) *SearchRecursionZonesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchRecursionZonesRequest
	GetPageSize() *int32
	SetRemark(v string) *SearchRecursionZonesRequest
	GetRemark() *string
	SetZoneName(v string) *SearchRecursionZonesRequest
	GetZoneName() *string
}

type SearchRecursionZonesRequest struct {
	// example:
	//
	// asc
	Direction       *string                                       `json:"Direction,omitempty" xml:"Direction,omitempty"`
	EffectiveScopes []*SearchRecursionZonesRequestEffectiveScopes `json:"EffectiveScopes,omitempty" xml:"EffectiveScopes,omitempty" type:"Repeated"`
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 4698691
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// default
	OrderBy *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
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
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// example:
	//
	// cheng.suow.cc
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s SearchRecursionZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesRequest) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesRequest) GetDirection() *string {
	return s.Direction
}

func (s *SearchRecursionZonesRequest) GetEffectiveScopes() []*SearchRecursionZonesRequestEffectiveScopes {
	return s.EffectiveScopes
}

func (s *SearchRecursionZonesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchRecursionZonesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchRecursionZonesRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchRecursionZonesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchRecursionZonesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchRecursionZonesRequest) GetRemark() *string {
	return s.Remark
}

func (s *SearchRecursionZonesRequest) GetZoneName() *string {
	return s.ZoneName
}

func (s *SearchRecursionZonesRequest) SetDirection(v string) *SearchRecursionZonesRequest {
	s.Direction = &v
	return s
}

func (s *SearchRecursionZonesRequest) SetEffectiveScopes(v []*SearchRecursionZonesRequestEffectiveScopes) *SearchRecursionZonesRequest {
	s.EffectiveScopes = v
	return s
}

func (s *SearchRecursionZonesRequest) SetMaxResults(v int32) *SearchRecursionZonesRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchRecursionZonesRequest) SetNextToken(v string) *SearchRecursionZonesRequest {
	s.NextToken = &v
	return s
}

func (s *SearchRecursionZonesRequest) SetOrderBy(v string) *SearchRecursionZonesRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchRecursionZonesRequest) SetPageNumber(v int32) *SearchRecursionZonesRequest {
	s.PageNumber = &v
	return s
}

func (s *SearchRecursionZonesRequest) SetPageSize(v int32) *SearchRecursionZonesRequest {
	s.PageSize = &v
	return s
}

func (s *SearchRecursionZonesRequest) SetRemark(v string) *SearchRecursionZonesRequest {
	s.Remark = &v
	return s
}

func (s *SearchRecursionZonesRequest) SetZoneName(v string) *SearchRecursionZonesRequest {
	s.ZoneName = &v
	return s
}

func (s *SearchRecursionZonesRequest) Validate() error {
	if s.EffectiveScopes != nil {
		for _, item := range s.EffectiveScopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchRecursionZonesRequestEffectiveScopes struct {
	// example:
	//
	// account
	EffectiveType *string `json:"EffectiveType,omitempty" xml:"EffectiveType,omitempty"`
	// example:
	//
	// [20003]
	Scope []*string `json:"Scope,omitempty" xml:"Scope,omitempty" type:"Repeated"`
}

func (s SearchRecursionZonesRequestEffectiveScopes) String() string {
	return dara.Prettify(s)
}

func (s SearchRecursionZonesRequestEffectiveScopes) GoString() string {
	return s.String()
}

func (s *SearchRecursionZonesRequestEffectiveScopes) GetEffectiveType() *string {
	return s.EffectiveType
}

func (s *SearchRecursionZonesRequestEffectiveScopes) GetScope() []*string {
	return s.Scope
}

func (s *SearchRecursionZonesRequestEffectiveScopes) SetEffectiveType(v string) *SearchRecursionZonesRequestEffectiveScopes {
	s.EffectiveType = &v
	return s
}

func (s *SearchRecursionZonesRequestEffectiveScopes) SetScope(v []*string) *SearchRecursionZonesRequestEffectiveScopes {
	s.Scope = v
	return s
}

func (s *SearchRecursionZonesRequestEffectiveScopes) Validate() error {
	return dara.Validate(s)
}
