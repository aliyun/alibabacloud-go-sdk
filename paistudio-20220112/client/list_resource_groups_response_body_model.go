// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListResourceGroupsResponseBody
	GetRequestId() *string
	SetResourceGroups(v []*ResourceGroup) *ListResourceGroupsResponseBody
	GetResourceGroups() []*ResourceGroup
	SetTotalCount(v int64) *ListResourceGroupsResponseBody
	GetTotalCount() *int64
}

type ListResourceGroupsResponseBody struct {
	// example:
	//
	// 9CFA2665-1FFE-5929-8468-C14C25890486
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// RG1
	ResourceGroups []*ResourceGroup `json:"ResourceGroups,omitempty" xml:"ResourceGroups,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceGroupsResponseBody) GetResourceGroups() []*ResourceGroup {
	return s.ResourceGroups
}

func (s *ListResourceGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListResourceGroupsResponseBody) SetRequestId(v string) *ListResourceGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceGroupsResponseBody) SetResourceGroups(v []*ResourceGroup) *ListResourceGroupsResponseBody {
	s.ResourceGroups = v
	return s
}

func (s *ListResourceGroupsResponseBody) SetTotalCount(v int64) *ListResourceGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceGroupsResponseBody) Validate() error {
	if s.ResourceGroups != nil {
		for _, item := range s.ResourceGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
