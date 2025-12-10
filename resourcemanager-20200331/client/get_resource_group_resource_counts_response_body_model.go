// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceGroupResourceCountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceGroupResourceCountsResponseBody
	GetRequestId() *string
	SetResourceCounts(v []*GetResourceGroupResourceCountsResponseBodyResourceCounts) *GetResourceGroupResourceCountsResponseBody
	GetResourceCounts() []*GetResourceGroupResourceCountsResponseBodyResourceCounts
}

type GetResourceGroupResourceCountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The numbers of the resources.
	ResourceCounts []*GetResourceGroupResourceCountsResponseBodyResourceCounts `json:"ResourceCounts,omitempty" xml:"ResourceCounts,omitempty" type:"Repeated"`
}

func (s GetResourceGroupResourceCountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResourceCountsResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResourceCountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceGroupResourceCountsResponseBody) GetResourceCounts() []*GetResourceGroupResourceCountsResponseBodyResourceCounts {
	return s.ResourceCounts
}

func (s *GetResourceGroupResourceCountsResponseBody) SetRequestId(v string) *GetResourceGroupResourceCountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceGroupResourceCountsResponseBody) SetResourceCounts(v []*GetResourceGroupResourceCountsResponseBodyResourceCounts) *GetResourceGroupResourceCountsResponseBody {
	s.ResourceCounts = v
	return s
}

func (s *GetResourceGroupResourceCountsResponseBody) Validate() error {
	if s.ResourceCounts != nil {
		for _, item := range s.ResourceCounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceGroupResourceCountsResponseBodyResourceCounts struct {
	// The number of the resources.
	//
	// example:
	//
	// 1
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The dimension by which resources are queried.
	//
	// example:
	//
	// ResourceGroupId
	GroupByKey *string `json:"GroupByKey,omitempty" xml:"GroupByKey,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-9gLOoK****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetResourceGroupResourceCountsResponseBodyResourceCounts) String() string {
	return dara.Prettify(s)
}

func (s GetResourceGroupResourceCountsResponseBodyResourceCounts) GoString() string {
	return s.String()
}

func (s *GetResourceGroupResourceCountsResponseBodyResourceCounts) GetCount() *int64 {
	return s.Count
}

func (s *GetResourceGroupResourceCountsResponseBodyResourceCounts) GetGroupByKey() *string {
	return s.GroupByKey
}

func (s *GetResourceGroupResourceCountsResponseBodyResourceCounts) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetResourceGroupResourceCountsResponseBodyResourceCounts) SetCount(v int64) *GetResourceGroupResourceCountsResponseBodyResourceCounts {
	s.Count = &v
	return s
}

func (s *GetResourceGroupResourceCountsResponseBodyResourceCounts) SetGroupByKey(v string) *GetResourceGroupResourceCountsResponseBodyResourceCounts {
	s.GroupByKey = &v
	return s
}

func (s *GetResourceGroupResourceCountsResponseBodyResourceCounts) SetResourceGroupId(v string) *GetResourceGroupResourceCountsResponseBodyResourceCounts {
	s.ResourceGroupId = &v
	return s
}

func (s *GetResourceGroupResourceCountsResponseBodyResourceCounts) Validate() error {
	return dara.Validate(s)
}
