// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseTemplateValidResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseTemplateValidResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*string) *DescribeDefenseTemplateValidResourcesResponseBody
	GetResources() []*string
	SetTotalCount(v int64) *DescribeDefenseTemplateValidResourcesResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseTemplateValidResourcesResponseBody struct {
	// example:
	//
	// C54DD36B-6380-57E5-89BA-2642757C4DB8
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 34
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseTemplateValidResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseTemplateValidResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseTemplateValidResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseTemplateValidResourcesResponseBody) GetResources() []*string {
	return s.Resources
}

func (s *DescribeDefenseTemplateValidResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseTemplateValidResourcesResponseBody) SetRequestId(v string) *DescribeDefenseTemplateValidResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesResponseBody) SetResources(v []*string) *DescribeDefenseTemplateValidResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesResponseBody) SetTotalCount(v int64) *DescribeDefenseTemplateValidResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseTemplateValidResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
