// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiAccountTagValuesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMatchType(v string) *ListMultiAccountTagValuesRequest
	GetMatchType() *string
	SetMaxResults(v int32) *ListMultiAccountTagValuesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListMultiAccountTagValuesRequest
	GetNextToken() *string
	SetScope(v string) *ListMultiAccountTagValuesRequest
	GetScope() *string
	SetTagKey(v string) *ListMultiAccountTagValuesRequest
	GetTagKey() *string
	SetTagValue(v string) *ListMultiAccountTagValuesRequest
	GetTagValue() *string
}

type ListMultiAccountTagValuesRequest struct {
	MatchType  *string `json:"MatchType,omitempty" xml:"MatchType,omitempty"`
	MaxResults *int32  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Scope      *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// This parameter is required.
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListMultiAccountTagValuesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMultiAccountTagValuesRequest) GoString() string {
	return s.String()
}

func (s *ListMultiAccountTagValuesRequest) GetMatchType() *string {
	return s.MatchType
}

func (s *ListMultiAccountTagValuesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMultiAccountTagValuesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMultiAccountTagValuesRequest) GetScope() *string {
	return s.Scope
}

func (s *ListMultiAccountTagValuesRequest) GetTagKey() *string {
	return s.TagKey
}

func (s *ListMultiAccountTagValuesRequest) GetTagValue() *string {
	return s.TagValue
}

func (s *ListMultiAccountTagValuesRequest) SetMatchType(v string) *ListMultiAccountTagValuesRequest {
	s.MatchType = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetMaxResults(v int32) *ListMultiAccountTagValuesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetNextToken(v string) *ListMultiAccountTagValuesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetScope(v string) *ListMultiAccountTagValuesRequest {
	s.Scope = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagKey(v string) *ListMultiAccountTagValuesRequest {
	s.TagKey = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) SetTagValue(v string) *ListMultiAccountTagValuesRequest {
	s.TagValue = &v
	return s
}

func (s *ListMultiAccountTagValuesRequest) Validate() error {
	return dara.Validate(s)
}
