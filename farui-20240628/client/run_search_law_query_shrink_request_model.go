// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchLawQueryShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *RunSearchLawQueryShrinkRequest
	GetAppId() *string
	SetFilterConditionShrink(v string) *RunSearchLawQueryShrinkRequest
	GetFilterConditionShrink() *string
	SetPageParamShrink(v string) *RunSearchLawQueryShrinkRequest
	GetPageParamShrink() *string
	SetQuery(v string) *RunSearchLawQueryShrinkRequest
	GetQuery() *string
	SetQueryKeywordsShrink(v string) *RunSearchLawQueryShrinkRequest
	GetQueryKeywordsShrink() *string
	SetThreadShrink(v string) *RunSearchLawQueryShrinkRequest
	GetThreadShrink() *string
}

type RunSearchLawQueryShrinkRequest struct {
	// example:
	//
	// farui
	AppId                 *string `json:"appId,omitempty" xml:"appId,omitempty"`
	FilterConditionShrink *string `json:"filterCondition,omitempty" xml:"filterCondition,omitempty"`
	PageParamShrink       *string `json:"pageParam,omitempty" xml:"pageParam,omitempty"`
	// This parameter is required.
	Query               *string `json:"query,omitempty" xml:"query,omitempty"`
	QueryKeywordsShrink *string `json:"queryKeywords,omitempty" xml:"queryKeywords,omitempty"`
	ThreadShrink        *string `json:"thread,omitempty" xml:"thread,omitempty"`
}

func (s RunSearchLawQueryShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSearchLawQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunSearchLawQueryShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *RunSearchLawQueryShrinkRequest) GetFilterConditionShrink() *string {
	return s.FilterConditionShrink
}

func (s *RunSearchLawQueryShrinkRequest) GetPageParamShrink() *string {
	return s.PageParamShrink
}

func (s *RunSearchLawQueryShrinkRequest) GetQuery() *string {
	return s.Query
}

func (s *RunSearchLawQueryShrinkRequest) GetQueryKeywordsShrink() *string {
	return s.QueryKeywordsShrink
}

func (s *RunSearchLawQueryShrinkRequest) GetThreadShrink() *string {
	return s.ThreadShrink
}

func (s *RunSearchLawQueryShrinkRequest) SetAppId(v string) *RunSearchLawQueryShrinkRequest {
	s.AppId = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetFilterConditionShrink(v string) *RunSearchLawQueryShrinkRequest {
	s.FilterConditionShrink = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetPageParamShrink(v string) *RunSearchLawQueryShrinkRequest {
	s.PageParamShrink = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetQuery(v string) *RunSearchLawQueryShrinkRequest {
	s.Query = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetQueryKeywordsShrink(v string) *RunSearchLawQueryShrinkRequest {
	s.QueryKeywordsShrink = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) SetThreadShrink(v string) *RunSearchLawQueryShrinkRequest {
	s.ThreadShrink = &v
	return s
}

func (s *RunSearchLawQueryShrinkRequest) Validate() error {
	return dara.Validate(s)
}
