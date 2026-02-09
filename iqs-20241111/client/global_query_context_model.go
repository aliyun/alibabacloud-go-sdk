// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGlobalQueryContext interface {
	dara.Model
	String() string
	GoString() string
	SetOriginalQuery(v *GlobalQueryContextOriginalQuery) *GlobalQueryContext
	GetOriginalQuery() *GlobalQueryContextOriginalQuery
}

type GlobalQueryContext struct {
	OriginalQuery *GlobalQueryContextOriginalQuery `json:"originalQuery,omitempty" xml:"originalQuery,omitempty" type:"Struct"`
}

func (s GlobalQueryContext) String() string {
	return dara.Prettify(s)
}

func (s GlobalQueryContext) GoString() string {
	return s.String()
}

func (s *GlobalQueryContext) GetOriginalQuery() *GlobalQueryContextOriginalQuery {
	return s.OriginalQuery
}

func (s *GlobalQueryContext) SetOriginalQuery(v *GlobalQueryContextOriginalQuery) *GlobalQueryContext {
	s.OriginalQuery = v
	return s
}

func (s *GlobalQueryContext) Validate() error {
	if s.OriginalQuery != nil {
		if err := s.OriginalQuery.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GlobalQueryContextOriginalQuery struct {
	// example:
	//
	// 1
	Page  *string `json:"page,omitempty" xml:"page,omitempty"`
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// example:
	//
	// NoLimit
	TimeRange *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s GlobalQueryContextOriginalQuery) String() string {
	return dara.Prettify(s)
}

func (s GlobalQueryContextOriginalQuery) GoString() string {
	return s.String()
}

func (s *GlobalQueryContextOriginalQuery) GetPage() *string {
	return s.Page
}

func (s *GlobalQueryContextOriginalQuery) GetQuery() *string {
	return s.Query
}

func (s *GlobalQueryContextOriginalQuery) GetTimeRange() *string {
	return s.TimeRange
}

func (s *GlobalQueryContextOriginalQuery) SetPage(v string) *GlobalQueryContextOriginalQuery {
	s.Page = &v
	return s
}

func (s *GlobalQueryContextOriginalQuery) SetQuery(v string) *GlobalQueryContextOriginalQuery {
	s.Query = &v
	return s
}

func (s *GlobalQueryContextOriginalQuery) SetTimeRange(v string) *GlobalQueryContextOriginalQuery {
	s.TimeRange = &v
	return s
}

func (s *GlobalQueryContextOriginalQuery) Validate() error {
	return dara.Validate(s)
}
