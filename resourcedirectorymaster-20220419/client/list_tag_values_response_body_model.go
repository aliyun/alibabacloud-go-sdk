// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagValuesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagValuesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagValuesResponseBody
	GetRequestId() *string
	SetTags(v []*ListTagValuesResponseBodyTags) *ListTagValuesResponseBody
	GetTags() []*ListTagValuesResponseBodyTags
}

type ListTagValuesResponseBody struct {
	// Indicates whether the next query is required.
	//
	// 	- If the value of this parameter is empty (`"NextToken": ""`), all results are returned, and the next query is not required.
	//
	// 	- If the value of this parameter is not empty, the next query is required, and the value is the token used to start the next query.
	//
	// example:
	//
	// TGlzdFJlc291cm****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DC09A6AA-2713-4E10-A2E9-E6C5C43A8842
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the tag values.
	Tags []*ListTagValuesResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagValuesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagValuesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagValuesResponseBody) GetTags() []*ListTagValuesResponseBodyTags {
	return s.Tags
}

func (s *ListTagValuesResponseBody) SetNextToken(v string) *ListTagValuesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagValuesResponseBody) SetRequestId(v string) *ListTagValuesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagValuesResponseBody) SetTags(v []*ListTagValuesResponseBodyTags) *ListTagValuesResponseBody {
	s.Tags = v
	return s
}

func (s *ListTagValuesResponseBody) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListTagValuesResponseBodyTags struct {
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagValuesResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagValuesResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagValuesResponseBodyTags) GetValue() *string {
	return s.Value
}

func (s *ListTagValuesResponseBodyTags) SetValue(v string) *ListTagValuesResponseBodyTags {
	s.Value = &v
	return s
}

func (s *ListTagValuesResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
