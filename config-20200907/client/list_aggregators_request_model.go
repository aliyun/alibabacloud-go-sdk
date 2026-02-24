// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAggregatorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListAggregatorsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAggregatorsRequest
	GetNextToken() *string
	SetTag(v []*ListAggregatorsRequestTag) *ListAggregatorsRequest
	GetTag() []*ListAggregatorsRequestTag
}

type ListAggregatorsRequest struct {
	// This parameter is required.
	MaxResults *int32                       `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                      `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Tag        []*ListAggregatorsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAggregatorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAggregatorsRequest) GoString() string {
	return s.String()
}

func (s *ListAggregatorsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAggregatorsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAggregatorsRequest) GetTag() []*ListAggregatorsRequestTag {
	return s.Tag
}

func (s *ListAggregatorsRequest) SetMaxResults(v int32) *ListAggregatorsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAggregatorsRequest) SetNextToken(v string) *ListAggregatorsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAggregatorsRequest) SetTag(v []*ListAggregatorsRequestTag) *ListAggregatorsRequest {
	s.Tag = v
	return s
}

func (s *ListAggregatorsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAggregatorsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAggregatorsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAggregatorsRequestTag) GoString() string {
	return s.String()
}

func (s *ListAggregatorsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAggregatorsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAggregatorsRequestTag) SetKey(v string) *ListAggregatorsRequestTag {
	s.Key = &v
	return s
}

func (s *ListAggregatorsRequestTag) SetValue(v string) *ListAggregatorsRequestTag {
	s.Value = &v
	return s
}

func (s *ListAggregatorsRequestTag) Validate() error {
	return dara.Validate(s)
}
