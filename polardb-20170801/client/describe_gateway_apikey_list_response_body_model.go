// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGatewayApikeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeGatewayApikeyListResponseBodyItems) *DescribeGatewayApikeyListResponseBody
	GetItems() []*DescribeGatewayApikeyListResponseBodyItems
	SetPageNumber(v int32) *DescribeGatewayApikeyListResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeGatewayApikeyListResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeGatewayApikeyListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGatewayApikeyListResponseBody
	GetRequestId() *string
	SetTotalPages(v int32) *DescribeGatewayApikeyListResponseBody
	GetTotalPages() *int32
	SetTotalRecordCount(v int32) *DescribeGatewayApikeyListResponseBody
	GetTotalRecordCount() *int32
}

type DescribeGatewayApikeyListResponseBody struct {
	// The list of consumer objects.
	Items []*DescribeGatewayApikeyListResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The number of entries per page. Valid values:
	//
	// 	- **30**
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// Default value: 30.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Id of the request
	//
	// example:
	//
	// A0450B18-BBD4-5DF9-8E71-610F1A921CDE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeGatewayApikeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayApikeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGatewayApikeyListResponseBody) GetItems() []*DescribeGatewayApikeyListResponseBodyItems {
	return s.Items
}

func (s *DescribeGatewayApikeyListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGatewayApikeyListResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeGatewayApikeyListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGatewayApikeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGatewayApikeyListResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeGatewayApikeyListResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeGatewayApikeyListResponseBody) SetItems(v []*DescribeGatewayApikeyListResponseBodyItems) *DescribeGatewayApikeyListResponseBody {
	s.Items = v
	return s
}

func (s *DescribeGatewayApikeyListResponseBody) SetPageNumber(v int32) *DescribeGatewayApikeyListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBody) SetPageRecordCount(v int32) *DescribeGatewayApikeyListResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBody) SetPageSize(v int32) *DescribeGatewayApikeyListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBody) SetRequestId(v string) *DescribeGatewayApikeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBody) SetTotalPages(v int32) *DescribeGatewayApikeyListResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBody) SetTotalRecordCount(v int32) *DescribeGatewayApikeyListResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBody) Validate() error {
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

type DescribeGatewayApikeyListResponseBodyItems struct {
	// The API key in use.
	//
	// example:
	//
	// xxxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The user group ID.
	//
	// example:
	//
	// cg-xxxxxx
	ConsumerGroupId *string `json:"ConsumerGroupId,omitempty" xml:"ConsumerGroupId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// c-mqveroemc***
	ConsumerId *string `json:"ConsumerId,omitempty" xml:"ConsumerId,omitempty"`
	// The consumer tag.
	//
	// example:
	//
	// test
	ConsumerTag *string `json:"ConsumerTag,omitempty" xml:"ConsumerTag,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2020-08-14T05:58:42Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The gateway instance ID.
	//
	// example:
	//
	// pg-xxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2024-10-29T09:31:37Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The username.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeGatewayApikeyListResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeGatewayApikeyListResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeGatewayApikeyListResponseBodyItems) GetApiKey() *string {
	return s.ApiKey
}

func (s *DescribeGatewayApikeyListResponseBodyItems) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *DescribeGatewayApikeyListResponseBodyItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *DescribeGatewayApikeyListResponseBodyItems) GetConsumerTag() *string {
	return s.ConsumerTag
}

func (s *DescribeGatewayApikeyListResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGatewayApikeyListResponseBodyItems) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *DescribeGatewayApikeyListResponseBodyItems) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeGatewayApikeyListResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *DescribeGatewayApikeyListResponseBodyItems) SetApiKey(v string) *DescribeGatewayApikeyListResponseBodyItems {
	s.ApiKey = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBodyItems) SetConsumerGroupId(v string) *DescribeGatewayApikeyListResponseBodyItems {
	s.ConsumerGroupId = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBodyItems) SetConsumerId(v string) *DescribeGatewayApikeyListResponseBodyItems {
	s.ConsumerId = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBodyItems) SetConsumerTag(v string) *DescribeGatewayApikeyListResponseBodyItems {
	s.ConsumerTag = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBodyItems) SetCreateTime(v string) *DescribeGatewayApikeyListResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBodyItems) SetGwClusterId(v string) *DescribeGatewayApikeyListResponseBodyItems {
	s.GwClusterId = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBodyItems) SetModifyTime(v string) *DescribeGatewayApikeyListResponseBodyItems {
	s.ModifyTime = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBodyItems) SetName(v string) *DescribeGatewayApikeyListResponseBodyItems {
	s.Name = &v
	return s
}

func (s *DescribeGatewayApikeyListResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
