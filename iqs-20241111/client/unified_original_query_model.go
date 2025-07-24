// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnifiedOriginalQuery interface {
	dara.Model
	String() string
	GoString() string
	SetQuery(v string) *UnifiedOriginalQuery
	GetQuery() *string
	SetTimeRange(v string) *UnifiedOriginalQuery
	GetTimeRange() *string
}

type UnifiedOriginalQuery struct {
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s UnifiedOriginalQuery) String() string {
	return dara.Prettify(s)
}

func (s UnifiedOriginalQuery) GoString() string {
	return s.String()
}

func (s *UnifiedOriginalQuery) GetQuery() *string {
	return s.Query
}

func (s *UnifiedOriginalQuery) GetTimeRange() *string {
	return s.TimeRange
}

func (s *UnifiedOriginalQuery) SetQuery(v string) *UnifiedOriginalQuery {
	s.Query = &v
	return s
}

func (s *UnifiedOriginalQuery) SetTimeRange(v string) *UnifiedOriginalQuery {
	s.TimeRange = &v
	return s
}

func (s *UnifiedOriginalQuery) Validate() error {
	return dara.Validate(s)
}
