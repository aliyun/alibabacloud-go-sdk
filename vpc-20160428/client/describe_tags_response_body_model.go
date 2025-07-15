// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeTagsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
	SetTagResources(v *DescribeTagsResponseBodyTagResources) *DescribeTagsResponseBody
	GetTagResources() *DescribeTagsResponseBodyTagResources
}

type DescribeTagsResponseBody struct {
	// The token that is used for the next query. Valid values:
	//
	// 	- If the value of **NextToken*	- is not returned, it indicates that no next query is to be sent.
	//
	// 	- If a value of **NextToken*	- is returned, the value is the token that is used for the subsequent query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C46FF5A8-C5F0-4024-8262-B16B639225A0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tags that meet the filter conditions.
	TagResources *DescribeTagsResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) GetTagResources() *DescribeTagsResponseBodyTagResources {
	return s.TagResources
}

func (s *DescribeTagsResponseBody) SetNextToken(v string) *DescribeTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTagResources(v *DescribeTagsResponseBodyTagResources) *DescribeTagsResponseBody {
	s.TagResources = v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsResponseBodyTagResources struct {
	TagResource []*DescribeTagsResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagResources) GetTagResource() []*DescribeTagsResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *DescribeTagsResponseBodyTagResources) SetTagResource(v []*DescribeTagsResponseBodyTagResourcesTagResource) *DescribeTagsResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *DescribeTagsResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}

type DescribeTagsResponseBodyTagResourcesTagResource struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// FinanceJoshua
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeTagsResponseBodyTagResourcesTagResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeTagsResponseBodyTagResourcesTagResource) SetTagKey(v string) *DescribeTagsResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTagResourcesTagResource) SetTagValue(v string) *DescribeTagsResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

func (s *DescribeTagsResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}
