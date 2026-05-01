// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelTemplateResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListModelTemplateResourceGroupResponseBodyData) *ListModelTemplateResourceGroupResponseBody
	GetData() []*ListModelTemplateResourceGroupResponseBodyData
	SetPageNumber(v int32) *ListModelTemplateResourceGroupResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListModelTemplateResourceGroupResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListModelTemplateResourceGroupResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListModelTemplateResourceGroupResponseBody
	GetTotalCount() *int32
}

type ListModelTemplateResourceGroupResponseBody struct {
	Data []*ListModelTemplateResourceGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListModelTemplateResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplateResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelTemplateResourceGroupResponseBody) GetData() []*ListModelTemplateResourceGroupResponseBodyData {
	return s.Data
}

func (s *ListModelTemplateResourceGroupResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelTemplateResourceGroupResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListModelTemplateResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelTemplateResourceGroupResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListModelTemplateResourceGroupResponseBody) SetData(v []*ListModelTemplateResourceGroupResponseBodyData) *ListModelTemplateResourceGroupResponseBody {
	s.Data = v
	return s
}

func (s *ListModelTemplateResourceGroupResponseBody) SetPageNumber(v int32) *ListModelTemplateResourceGroupResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelTemplateResourceGroupResponseBody) SetPageSize(v int32) *ListModelTemplateResourceGroupResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListModelTemplateResourceGroupResponseBody) SetRequestId(v string) *ListModelTemplateResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelTemplateResourceGroupResponseBody) SetTotalCount(v int32) *ListModelTemplateResourceGroupResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelTemplateResourceGroupResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelTemplateResourceGroupResponseBodyData struct {
	// example:
	//
	// rg-xxxxx
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListModelTemplateResourceGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListModelTemplateResourceGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelTemplateResourceGroupResponseBodyData) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListModelTemplateResourceGroupResponseBodyData) SetResourceGroupId(v string) *ListModelTemplateResourceGroupResponseBodyData {
	s.ResourceGroupId = &v
	return s
}

func (s *ListModelTemplateResourceGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
