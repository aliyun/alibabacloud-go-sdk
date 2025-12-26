// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKnowledgeBaseFilters(v []*SearchInfoRequestKnowledgeBaseFilters) *SearchInfoRequest
	GetKnowledgeBaseFilters() []*SearchInfoRequestKnowledgeBaseFilters
	SetWebFilters(v []*SearchInfoRequestWebFilters) *SearchInfoRequest
	GetWebFilters() []*SearchInfoRequestWebFilters
}

type SearchInfoRequest struct {
	KnowledgeBaseFilters []*SearchInfoRequestKnowledgeBaseFilters `json:"KnowledgeBaseFilters,omitempty" xml:"KnowledgeBaseFilters,omitempty" type:"Repeated"`
	WebFilters           []*SearchInfoRequestWebFilters           `json:"WebFilters,omitempty" xml:"WebFilters,omitempty" type:"Repeated"`
}

func (s SearchInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchInfoRequest) GoString() string {
	return s.String()
}

func (s *SearchInfoRequest) GetKnowledgeBaseFilters() []*SearchInfoRequestKnowledgeBaseFilters {
	return s.KnowledgeBaseFilters
}

func (s *SearchInfoRequest) GetWebFilters() []*SearchInfoRequestWebFilters {
	return s.WebFilters
}

func (s *SearchInfoRequest) SetKnowledgeBaseFilters(v []*SearchInfoRequestKnowledgeBaseFilters) *SearchInfoRequest {
	s.KnowledgeBaseFilters = v
	return s
}

func (s *SearchInfoRequest) SetWebFilters(v []*SearchInfoRequestWebFilters) *SearchInfoRequest {
	s.WebFilters = v
	return s
}

func (s *SearchInfoRequest) Validate() error {
	if s.KnowledgeBaseFilters != nil {
		for _, item := range s.KnowledgeBaseFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.WebFilters != nil {
		for _, item := range s.WebFilters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchInfoRequestKnowledgeBaseFilters struct {
	// This parameter is required.
	//
	// example:
	//
	// pai_qa
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// what is pai
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 5
	ResultLimit *int32 `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// 0.7827883223
	ScoreThreshold *float64 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
}

func (s SearchInfoRequestKnowledgeBaseFilters) String() string {
	return dara.Prettify(s)
}

func (s SearchInfoRequestKnowledgeBaseFilters) GoString() string {
	return s.String()
}

func (s *SearchInfoRequestKnowledgeBaseFilters) GetCollectionName() *string {
	return s.CollectionName
}

func (s *SearchInfoRequestKnowledgeBaseFilters) GetQuery() *string {
	return s.Query
}

func (s *SearchInfoRequestKnowledgeBaseFilters) GetResultLimit() *int32 {
	return s.ResultLimit
}

func (s *SearchInfoRequestKnowledgeBaseFilters) GetScoreThreshold() *float64 {
	return s.ScoreThreshold
}

func (s *SearchInfoRequestKnowledgeBaseFilters) SetCollectionName(v string) *SearchInfoRequestKnowledgeBaseFilters {
	s.CollectionName = &v
	return s
}

func (s *SearchInfoRequestKnowledgeBaseFilters) SetQuery(v string) *SearchInfoRequestKnowledgeBaseFilters {
	s.Query = &v
	return s
}

func (s *SearchInfoRequestKnowledgeBaseFilters) SetResultLimit(v int32) *SearchInfoRequestKnowledgeBaseFilters {
	s.ResultLimit = &v
	return s
}

func (s *SearchInfoRequestKnowledgeBaseFilters) SetScoreThreshold(v float64) *SearchInfoRequestKnowledgeBaseFilters {
	s.ScoreThreshold = &v
	return s
}

func (s *SearchInfoRequestKnowledgeBaseFilters) Validate() error {
	return dara.Validate(s)
}

type SearchInfoRequestWebFilters struct {
	Category     *string   `json:"Category,omitempty" xml:"Category,omitempty"`
	IncludeSites []*string `json:"IncludeSites,omitempty" xml:"IncludeSites,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// what is dsw
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// 5
	ResultLimit *int32 `json:"ResultLimit,omitempty" xml:"ResultLimit,omitempty"`
	// example:
	//
	// 0.832883223
	ScoreThreshold *float64 `json:"ScoreThreshold,omitempty" xml:"ScoreThreshold,omitempty"`
}

func (s SearchInfoRequestWebFilters) String() string {
	return dara.Prettify(s)
}

func (s SearchInfoRequestWebFilters) GoString() string {
	return s.String()
}

func (s *SearchInfoRequestWebFilters) GetCategory() *string {
	return s.Category
}

func (s *SearchInfoRequestWebFilters) GetIncludeSites() []*string {
	return s.IncludeSites
}

func (s *SearchInfoRequestWebFilters) GetQuery() *string {
	return s.Query
}

func (s *SearchInfoRequestWebFilters) GetResultLimit() *int32 {
	return s.ResultLimit
}

func (s *SearchInfoRequestWebFilters) GetScoreThreshold() *float64 {
	return s.ScoreThreshold
}

func (s *SearchInfoRequestWebFilters) SetCategory(v string) *SearchInfoRequestWebFilters {
	s.Category = &v
	return s
}

func (s *SearchInfoRequestWebFilters) SetIncludeSites(v []*string) *SearchInfoRequestWebFilters {
	s.IncludeSites = v
	return s
}

func (s *SearchInfoRequestWebFilters) SetQuery(v string) *SearchInfoRequestWebFilters {
	s.Query = &v
	return s
}

func (s *SearchInfoRequestWebFilters) SetResultLimit(v int32) *SearchInfoRequestWebFilters {
	s.ResultLimit = &v
	return s
}

func (s *SearchInfoRequestWebFilters) SetScoreThreshold(v float64) *SearchInfoRequestWebFilters {
	s.ScoreThreshold = &v
	return s
}

func (s *SearchInfoRequestWebFilters) Validate() error {
	return dara.Validate(s)
}
