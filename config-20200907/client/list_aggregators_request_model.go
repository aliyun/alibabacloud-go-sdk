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
	// The maximum number of entries to return in a request. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// TGlzdFJlc291cmNlU2hhcmVzJjE1MTI2NjY4NzY5MTAzOTEmMiZORnI4NDhVeEtrUT0
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The tags of the resource.
	//
	// You can add up to 20 tags to a resource.
	Tag []*ListAggregatorsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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
	// The tag key of the resource. You can specify up to 20 tag keys.
	//
	// The tag key cannot be an empty string. The tag key must be 1 to 64 characters in length and cannot start with `aliyun` or `acs`:. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// key-1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag values.
	//
	// The tag values can be an empty string or up to 128 characters in length. The tag values cannot start with `aliyun` or `acs:` and cannot contain `http://` or `https://`.
	//
	// Each key-value must be unique. You can specify at most 20 tag values in each call.
	//
	// example:
	//
	// value-1
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
