// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchCaseFullTextShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunSearchCaseFullTextShrinkRequest
	GetAppId() *string
	SetFilterConditionShrink(v string) *RunSearchCaseFullTextShrinkRequest
	GetFilterConditionShrink() *string
	SetPageParamShrink(v string) *RunSearchCaseFullTextShrinkRequest
	GetPageParamShrink() *string
	SetQuery(v string) *RunSearchCaseFullTextShrinkRequest
	GetQuery() *string
	SetQueryKeywordsShrink(v string) *RunSearchCaseFullTextShrinkRequest
	GetQueryKeywordsShrink() *string
	SetReferLevel(v string) *RunSearchCaseFullTextShrinkRequest
	GetReferLevel() *string
	SetSortKeyAndDirectionShrink(v string) *RunSearchCaseFullTextShrinkRequest
	GetSortKeyAndDirectionShrink() *string
	SetThreadShrink(v string) *RunSearchCaseFullTextShrinkRequest
	GetThreadShrink() *string
}

type RunSearchCaseFullTextShrinkRequest struct {
	// example:
	//
	// farui
	AppId                 *string `json:"appId,omitempty" xml:"appId,omitempty"`
	FilterConditionShrink *string `json:"filterCondition,omitempty" xml:"filterCondition,omitempty"`
	// This parameter is required.
	PageParamShrink *string `json:"pageParam,omitempty" xml:"pageParam,omitempty"`
	// This parameter is required.
	Query                     *string `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywordsShrink       *string `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty"`
	ReferLevel                *string `json:"referLevel,omitempty" xml:"referLevel,omitempty"`
	SortKeyAndDirectionShrink *string `json:"sortKeyAndDirection,omitempty" xml:"sortKeyAndDirection,omitempty"`
	ThreadShrink              *string `json:"thread,omitempty" xml:"thread,omitempty"`
}

func (s RunSearchCaseFullTextShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunSearchCaseFullTextShrinkRequest) GetFilterConditionShrink() *string {
	return s.FilterConditionShrink
}

func (s *RunSearchCaseFullTextShrinkRequest) GetPageParamShrink() *string {
	return s.PageParamShrink
}

func (s *RunSearchCaseFullTextShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *RunSearchCaseFullTextShrinkRequest) GetQueryKeywordsShrink() *string {
	return s.QueryKeywordsShrink
}

func (s *RunSearchCaseFullTextShrinkRequest) GetReferLevel() *string {
	return s.ReferLevel
}

func (s *RunSearchCaseFullTextShrinkRequest) GetSortKeyAndDirectionShrink() *string {
	return s.SortKeyAndDirectionShrink
}

func (s *RunSearchCaseFullTextShrinkRequest) GetThreadShrink() *string {
	return s.ThreadShrink
}

func (s *RunSearchCaseFullTextShrinkRequest) SetAppId(v string) *RunSearchCaseFullTextShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetFilterConditionShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetPageParamShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.PageParamShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetQuery(v string) *RunSearchCaseFullTextShrinkRequest {
	s.Query = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetQueryKeywordsShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.QueryKeywordsShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetReferLevel(v string) *RunSearchCaseFullTextShrinkRequest {
	s.ReferLevel = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetSortKeyAndDirectionShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.SortKeyAndDirectionShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) SetThreadShrink(v string) *RunSearchCaseFullTextShrinkRequest {
	s.ThreadShrink = &v
	return s
}

func (s *RunSearchCaseFullTextShrinkRequest) Validate() error {
	return dara.Validate(s)
}
