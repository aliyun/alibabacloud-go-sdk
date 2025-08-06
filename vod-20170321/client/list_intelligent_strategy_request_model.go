// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntelligentStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNo(v int32) *ListIntelligentStrategyRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListIntelligentStrategyRequest
	GetPageSize() *int32
	SetState(v string) *ListIntelligentStrategyRequest
	GetState() *string
	SetStrategyIds(v string) *ListIntelligentStrategyRequest
	GetStrategyIds() *string
	SetType(v string) *ListIntelligentStrategyRequest
	GetType() *string
}

type ListIntelligentStrategyRequest struct {
	PageNo      *int32  `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	PageSize    *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	StrategyIds *string `json:"StrategyIds,omitempty" xml:"StrategyIds,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListIntelligentStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntelligentStrategyRequest) GoString() string {
	return s.String()
}

func (s *ListIntelligentStrategyRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListIntelligentStrategyRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListIntelligentStrategyRequest) GetState() *string {
	return s.State
}

func (s *ListIntelligentStrategyRequest) GetStrategyIds() *string {
	return s.StrategyIds
}

func (s *ListIntelligentStrategyRequest) GetType() *string {
	return s.Type
}

func (s *ListIntelligentStrategyRequest) SetPageNo(v int32) *ListIntelligentStrategyRequest {
	s.PageNo = &v
	return s
}

func (s *ListIntelligentStrategyRequest) SetPageSize(v int32) *ListIntelligentStrategyRequest {
	s.PageSize = &v
	return s
}

func (s *ListIntelligentStrategyRequest) SetState(v string) *ListIntelligentStrategyRequest {
	s.State = &v
	return s
}

func (s *ListIntelligentStrategyRequest) SetStrategyIds(v string) *ListIntelligentStrategyRequest {
	s.StrategyIds = &v
	return s
}

func (s *ListIntelligentStrategyRequest) SetType(v string) *ListIntelligentStrategyRequest {
	s.Type = &v
	return s
}

func (s *ListIntelligentStrategyRequest) Validate() error {
	return dara.Validate(s)
}
