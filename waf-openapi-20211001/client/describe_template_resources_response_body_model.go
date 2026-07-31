// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplateResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeTemplateResourcesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTemplateResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTemplateResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*string) *DescribeTemplateResourcesResponseBody
	GetResources() []*string
	SetTemplateId(v int64) *DescribeTemplateResourcesResponseBody
	GetTemplateId() *int64
	SetTotalCount(v int32) *DescribeTemplateResourcesResponseBody
	GetTotalCount() *int32
}

type DescribeTemplateResourcesResponseBody struct {
	// The number of entries per page in a paged query. Valid values: 1 to 500. Default value: 500.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token (Token) for the next page. If a next page exists, this field contains a value.
	//
	// > If this parameter has a return value, a next page exists. Use the returned **NextToken*	- as a request parameter to retrieve the next page of data. Repeat until no value is returned, which indicates that all data has been retrieved.
	//
	// example:
	//
	// AAAAABLQv******37sHZaHk4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C7BC9373-3960-53B0-8968-2B13454AE18F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of bound protected object names, protected object group names, or protected asset IDs.
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// The ID of the protection template.
	//
	// example:
	//
	// 168465
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 25
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeTemplateResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplateResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplateResourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTemplateResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTemplateResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplateResourcesResponseBody) GetResources() []*string {
	return s.Resources
}

func (s *DescribeTemplateResourcesResponseBody) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *DescribeTemplateResourcesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeTemplateResourcesResponseBody) SetMaxResults(v int32) *DescribeTemplateResourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeTemplateResourcesResponseBody) SetNextToken(v string) *DescribeTemplateResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTemplateResourcesResponseBody) SetRequestId(v string) *DescribeTemplateResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplateResourcesResponseBody) SetResources(v []*string) *DescribeTemplateResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeTemplateResourcesResponseBody) SetTemplateId(v int64) *DescribeTemplateResourcesResponseBody {
	s.TemplateId = &v
	return s
}

func (s *DescribeTemplateResourcesResponseBody) SetTotalCount(v int32) *DescribeTemplateResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeTemplateResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
