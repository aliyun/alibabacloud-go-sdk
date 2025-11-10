// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagKeysResponseBody
	GetRequestId() *string
	SetTags(v []*ListTagKeysResponseBodyTags) *ListTagKeysResponseBody
	GetTags() []*ListTagKeysResponseBodyTags
}

type ListTagKeysResponseBody struct {
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
	// The information about the tag keys.
	Tags []*ListTagKeysResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagKeysResponseBody) GetTags() []*ListTagKeysResponseBodyTags {
	return s.Tags
}

func (s *ListTagKeysResponseBody) SetNextToken(v string) *ListTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagKeysResponseBody) SetRequestId(v string) *ListTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagKeysResponseBody) SetTags(v []*ListTagKeysResponseBodyTags) *ListTagKeysResponseBody {
	s.Tags = v
	return s
}

func (s *ListTagKeysResponseBody) Validate() error {
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

type ListTagKeysResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// team
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ListTagKeysResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s ListTagKeysResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagKeysResponseBodyTags) GetKey() *string {
	return s.Key
}

func (s *ListTagKeysResponseBodyTags) SetKey(v string) *ListTagKeysResponseBodyTags {
	s.Key = &v
	return s
}

func (s *ListTagKeysResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
