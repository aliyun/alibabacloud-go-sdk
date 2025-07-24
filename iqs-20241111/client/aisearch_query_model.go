// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchQuery interface {
	dara.Model
	String() string
	GoString() string
	SetCard(v string) *AISearchQuery
	GetCard() *string
	SetCardQuery(v string) *AISearchQuery
	GetCardQuery() *string
	SetPage(v int32) *AISearchQuery
	GetPage() *int32
	SetQuery(v string) *AISearchQuery
	GetQuery() *string
	SetSearchEngine(v string) *AISearchQuery
	GetSearchEngine() *string
	SetSearchPlan(v string) *AISearchQuery
	GetSearchPlan() *string
	SetSessionId(v string) *AISearchQuery
	GetSessionId() *string
	SetTimeRange(v string) *AISearchQuery
	GetTimeRange() *string
}

type AISearchQuery struct {
	Card         *string `json:"card,omitempty" xml:"card,omitempty"`
	CardQuery    *string `json:"cardQuery,omitempty" xml:"cardQuery,omitempty"`
	Page         *int32  `json:"page,omitempty" xml:"page,omitempty"`
	Query        *string `json:"query,omitempty" xml:"query,omitempty"`
	SearchEngine *string `json:"searchEngine,omitempty" xml:"searchEngine,omitempty"`
	SearchPlan   *string `json:"searchPlan,omitempty" xml:"searchPlan,omitempty"`
	SessionId    *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	TimeRange    *string `json:"timeRange,omitempty" xml:"timeRange,omitempty"`
}

func (s AISearchQuery) String() string {
	return dara.Prettify(s)
}

func (s AISearchQuery) GoString() string {
	return s.String()
}

func (s *AISearchQuery) GetCard() *string {
	return s.Card
}

func (s *AISearchQuery) GetCardQuery() *string {
	return s.CardQuery
}

func (s *AISearchQuery) GetPage() *int32 {
	return s.Page
}

func (s *AISearchQuery) GetQuery() *string {
	return s.Query
}

func (s *AISearchQuery) GetSearchEngine() *string {
	return s.SearchEngine
}

func (s *AISearchQuery) GetSearchPlan() *string {
	return s.SearchPlan
}

func (s *AISearchQuery) GetSessionId() *string {
	return s.SessionId
}

func (s *AISearchQuery) GetTimeRange() *string {
	return s.TimeRange
}

func (s *AISearchQuery) SetCard(v string) *AISearchQuery {
	s.Card = &v
	return s
}

func (s *AISearchQuery) SetCardQuery(v string) *AISearchQuery {
	s.CardQuery = &v
	return s
}

func (s *AISearchQuery) SetPage(v int32) *AISearchQuery {
	s.Page = &v
	return s
}

func (s *AISearchQuery) SetQuery(v string) *AISearchQuery {
	s.Query = &v
	return s
}

func (s *AISearchQuery) SetSearchEngine(v string) *AISearchQuery {
	s.SearchEngine = &v
	return s
}

func (s *AISearchQuery) SetSearchPlan(v string) *AISearchQuery {
	s.SearchPlan = &v
	return s
}

func (s *AISearchQuery) SetSessionId(v string) *AISearchQuery {
	s.SessionId = &v
	return s
}

func (s *AISearchQuery) SetTimeRange(v string) *AISearchQuery {
	s.TimeRange = &v
	return s
}

func (s *AISearchQuery) Validate() error {
	return dara.Validate(s)
}
