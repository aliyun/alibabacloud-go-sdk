// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListTagValuesRequest
	GetMatchType() *string
	SetMaxResults(v int32) *ListTagValuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTagValuesRequest
	GetNextToken() *string
	SetTagKey(v string) *ListTagValuesRequest
	GetTagKey() *string
	SetTagValue(v string) *ListTagValuesRequest
	GetTagValue() *string
}

type ListTagValuesRequest struct {
	MatchType  *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListTagValuesRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListTagValuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagValuesRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *ListTagValuesRequest) SetMatchType(v string) *ListTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListTagValuesRequest) SetMaxResults(v int32) *ListTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTagValuesRequest) SetNextToken(v string) *ListTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesRequest) SetTagKey(v string) *ListTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListTagValuesRequest) SetTagValue(v string) *ListTagValuesRequest {
	s.TagValue = &v
	return s
}

func (s *ListTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
