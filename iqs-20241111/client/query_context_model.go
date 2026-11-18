// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContext interface {
	dara.Model
	String() string
	GoString() string
	SetOriginalQuery(v *QueryContextOriginalQuery) *QueryContext
	GetOriginalQuery() *QueryContextOriginalQuery
	SetRewrite(v *QueryContextRewrite) *QueryContext
	GetRewrite() *QueryContextRewrite
}

type QueryContext struct {
	// The initial query condition.
	OriginalQuery *QueryContextOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty" type:"Struct"`
	// The query rewrite information.
	Rewrite *QueryContextRewrite `json:"rewrite,omitempty" xml:"rewrite,omitempty" type:"Struct"`
}

func (s QueryContext) String() string {
	return dara.Prettify(s)
}

func (s QueryContext) GoString() string {
	return s.String()
}

func (s *QueryContext) GetOriginalQuery() *QueryContextOriginalQuery {
	return s.OriginalQuery
}

func (s *QueryContext) GetRewrite() *QueryContextRewrite {
	return s.Rewrite
}

func (s *QueryContext) SetOriginalQuery(v *QueryContextOriginalQuery) *QueryContext {
	s.OriginalQuery = v
	return s
}

func (s *QueryContext) SetRewrite(v *QueryContextRewrite) *QueryContext {
	s.Rewrite = v
	return s
}

func (s *QueryContext) Validate() error {
	if s.OriginalQuery != nil {
		if err := s.OriginalQuery.Validate(); err != nil {
			return err
		}
	}
	if s.Rewrite != nil {
		if err := s.Rewrite.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryContextOriginalQuery struct {
	// The industry classification of the original request.
	//
	// example:
	//
	// medical
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
	// The page number of the original request.
	//
	// example:
	//
	// 1
	Page *string `json:"page,omitempty" xml:"page,omitempty"`
	// The query content of the original request.
	//
	// example:
	//
	// Apsara Conference
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// The time range of the original request. Valid values: OneDay, OneWeek, OneMonth, OneYear, and NoLimit.
	//
	// example:
	//
	// NoLimit
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s QueryContextOriginalQuery) String() string {
	return dara.Prettify(s)
}

func (s QueryContextOriginalQuery) GoString() string {
	return s.String()
}

func (s *QueryContextOriginalQuery) GetIndustry() *string {
	return s.Industry
}

func (s *QueryContextOriginalQuery) GetPage() *string {
	return s.Page
}

func (s *QueryContextOriginalQuery) GetQuery() *string {
	return s.Query
}

func (s *QueryContextOriginalQuery) GetTimeRange() *string {
	return s.TimeRange
}

func (s *QueryContextOriginalQuery) SetIndustry(v string) *QueryContextOriginalQuery {
	s.Industry = &v
	return s
}

func (s *QueryContextOriginalQuery) SetPage(v string) *QueryContextOriginalQuery {
	s.Page = &v
	return s
}

func (s *QueryContextOriginalQuery) SetQuery(v string) *QueryContextOriginalQuery {
	s.Query = &v
	return s
}

func (s *QueryContextOriginalQuery) SetTimeRange(v string) *QueryContextOriginalQuery {
	s.TimeRange = &v
	return s
}

func (s *QueryContextOriginalQuery) Validate() error {
	return dara.Validate(s)
}

type QueryContextRewrite struct {
	// Specifies whether query rewrite is enabled.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The rewritten time range. An empty value indicates that no rewrite is performed or the rewrite result is the same as the original result.
	//
	// example:
	//
	// OneYear
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s QueryContextRewrite) String() string {
	return dara.Prettify(s)
}

func (s QueryContextRewrite) GoString() string {
	return s.String()
}

func (s *QueryContextRewrite) GetEnabled() *bool {
	return s.Enabled
}

func (s *QueryContextRewrite) GetTimeRange() *string {
	return s.TimeRange
}

func (s *QueryContextRewrite) SetEnabled(v bool) *QueryContextRewrite {
	s.Enabled = &v
	return s
}

func (s *QueryContextRewrite) SetTimeRange(v string) *QueryContextRewrite {
	s.TimeRange = &v
	return s
}

func (s *QueryContextRewrite) Validate() error {
	return dara.Validate(s)
}
