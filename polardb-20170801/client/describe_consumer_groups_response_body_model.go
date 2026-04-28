// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsumerGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeConsumerGroupsResponseBodyItems) *DescribeConsumerGroupsResponseBody
	GetItems() []*DescribeConsumerGroupsResponseBodyItems
	SetPageNumber(v int32) *DescribeConsumerGroupsResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeConsumerGroupsResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeConsumerGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeConsumerGroupsResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeConsumerGroupsResponseBody
	GetTotalRecordCount() *int32
}

type DescribeConsumerGroupsResponseBody struct {
	Items []*DescribeConsumerGroupsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 30
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 24A1990B-4F6E-482B-B8CB-75C612******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeConsumerGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupsResponseBody) GetItems() []*DescribeConsumerGroupsResponseBodyItems {
	return s.Items
}

func (s *DescribeConsumerGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeConsumerGroupsResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeConsumerGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeConsumerGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeConsumerGroupsResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeConsumerGroupsResponseBody) SetItems(v []*DescribeConsumerGroupsResponseBodyItems) *DescribeConsumerGroupsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeConsumerGroupsResponseBody) SetPageNumber(v int32) *DescribeConsumerGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBody) SetPageRecordCount(v int32) *DescribeConsumerGroupsResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBody) SetPageSize(v int32) *DescribeConsumerGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBody) SetRequestId(v string) *DescribeConsumerGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBody) SetTotalRecordCount(v int32) *DescribeConsumerGroupsResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeConsumerGroupsResponseBodyItems struct {
	// example:
	//
	// "[]"
	AllowedModels *string `json:"AllowedModels,omitempty" xml:"AllowedModels,omitempty"`
	// example:
	//
	// cg-xxxxxxxx
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// example:
	//
	// test
	ConsumerGroupName *string `json:"ConsumerGroupName,omitempty" xml:"ConsumerGroupName,omitempty"`
	// example:
	//
	// 2026-01-28T09:56:03+08:00
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2026-01-04T16:09:29+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 0
	IsDefault *string `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	// example:
	//
	// test
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
}

func (s DescribeConsumerGroupsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsumerGroupsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeConsumerGroupsResponseBodyItems) GetAllowedModels() *string {
	return s.AllowedModels
}

func (s *DescribeConsumerGroupsResponseBodyItems) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *DescribeConsumerGroupsResponseBodyItems) GetConsumerGroupName() *string {
	return s.ConsumerGroupName
}

func (s *DescribeConsumerGroupsResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeConsumerGroupsResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeConsumerGroupsResponseBodyItems) GetIsDefault() *string {
	return s.IsDefault
}

func (s *DescribeConsumerGroupsResponseBodyItems) GetNickName() *string {
	return s.NickName
}

func (s *DescribeConsumerGroupsResponseBodyItems) SetAllowedModels(v string) *DescribeConsumerGroupsResponseBodyItems {
	s.AllowedModels = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBodyItems) SetConsumerGroupId(v string) *DescribeConsumerGroupsResponseBodyItems {
	s.ConsumerGroupId = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBodyItems) SetConsumerGroupName(v string) *DescribeConsumerGroupsResponseBodyItems {
	s.ConsumerGroupName = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBodyItems) SetGmtCreated(v string) *DescribeConsumerGroupsResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBodyItems) SetGmtModified(v string) *DescribeConsumerGroupsResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBodyItems) SetIsDefault(v string) *DescribeConsumerGroupsResponseBodyItems {
	s.IsDefault = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBodyItems) SetNickName(v string) *DescribeConsumerGroupsResponseBodyItems {
	s.NickName = &v
	return s
}

func (s *DescribeConsumerGroupsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
