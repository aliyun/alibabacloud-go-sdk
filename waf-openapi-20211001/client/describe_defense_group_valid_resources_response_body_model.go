// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDefenseGroupValidResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDefenseGroupValidResourcesResponseBody
	GetRequestId() *string
	SetResources(v []*string) *DescribeDefenseGroupValidResourcesResponseBody
	GetResources() []*string
	SetTotalCount(v int64) *DescribeDefenseGroupValidResourcesResponseBody
	GetTotalCount() *int64
}

type DescribeDefenseGroupValidResourcesResponseBody struct {
	// example:
	//
	// A57BA089-3B28-5C82-8331-7B94****4978
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Resources []*string `json:"Resources,omitempty" xml:"Resources,omitempty" type:"Repeated"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDefenseGroupValidResourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDefenseGroupValidResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDefenseGroupValidResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDefenseGroupValidResourcesResponseBody) GetResources() []*string {
	return s.Resources
}

func (s *DescribeDefenseGroupValidResourcesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDefenseGroupValidResourcesResponseBody) SetRequestId(v string) *DescribeDefenseGroupValidResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesResponseBody) SetResources(v []*string) *DescribeDefenseGroupValidResourcesResponseBody {
	s.Resources = v
	return s
}

func (s *DescribeDefenseGroupValidResourcesResponseBody) SetTotalCount(v int64) *DescribeDefenseGroupValidResourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDefenseGroupValidResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}
