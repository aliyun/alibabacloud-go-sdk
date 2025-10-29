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
	OriginalQuery *QueryContextOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty" type:"Struct"`
	Rewrite       *QueryContextRewrite       `json:"rewrite,omitempty" xml:"rewrite,omitempty" type:"Struct"`
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
	Industry  *string `json:"industry,omitempty" xml:"industry,omitempty"`
	Page      *string `json:"page,omitempty" xml:"page,omitempty"`
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
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
	Enabled   *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
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
