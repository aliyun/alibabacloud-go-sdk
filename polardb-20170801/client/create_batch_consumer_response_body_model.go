// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchConsumerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*CreateBatchConsumerResponseBodyItems) *CreateBatchConsumerResponseBody
	GetItems() []*CreateBatchConsumerResponseBodyItems
	SetPageNumber(v int32) *CreateBatchConsumerResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *CreateBatchConsumerResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *CreateBatchConsumerResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *CreateBatchConsumerResponseBody
	GetRequestId() *string
	SetTotalPages(v int32) *CreateBatchConsumerResponseBody
	GetTotalPages() *int32
	SetTotalRecordCount(v int32) *CreateBatchConsumerResponseBody
	GetTotalRecordCount() *int32
}

type CreateBatchConsumerResponseBody struct {
	// The list of consumer objects.
	Items []*CreateBatchConsumerResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
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
	// The number of records per page. Valid values:
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
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
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
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s CreateBatchConsumerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchConsumerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBatchConsumerResponseBody) GetItems() []*CreateBatchConsumerResponseBodyItems {
	return s.Items
}

func (s *CreateBatchConsumerResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CreateBatchConsumerResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *CreateBatchConsumerResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CreateBatchConsumerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBatchConsumerResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *CreateBatchConsumerResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *CreateBatchConsumerResponseBody) SetItems(v []*CreateBatchConsumerResponseBodyItems) *CreateBatchConsumerResponseBody {
	s.Items = v
	return s
}

func (s *CreateBatchConsumerResponseBody) SetPageNumber(v int32) *CreateBatchConsumerResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CreateBatchConsumerResponseBody) SetPageRecordCount(v int32) *CreateBatchConsumerResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *CreateBatchConsumerResponseBody) SetPageSize(v int32) *CreateBatchConsumerResponseBody {
	s.PageSize = &v
	return s
}

func (s *CreateBatchConsumerResponseBody) SetRequestId(v string) *CreateBatchConsumerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBatchConsumerResponseBody) SetTotalPages(v int32) *CreateBatchConsumerResponseBody {
	s.TotalPages = &v
	return s
}

func (s *CreateBatchConsumerResponseBody) SetTotalRecordCount(v int32) *CreateBatchConsumerResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *CreateBatchConsumerResponseBody) Validate() error {
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

type CreateBatchConsumerResponseBodyItems struct {
	// The complete API key. This value is returned only in the current response.
	//
	// example:
	//
	// xxxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The API key status. Default value: Active.
	//
	// example:
	//
	// Active
	ApiStatus *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// The ID of the user group.
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
	// The gateway instance ID.
	//
	// example:
	//
	// pg-xxxxxxxxxx
	GwClusterId *string `json:"GwClusterId,omitempty" xml:"GwClusterId,omitempty"`
	// The name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The consumer status. Default value: Enabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateBatchConsumerResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchConsumerResponseBodyItems) GoString() string {
	return s.String()
}

func (s *CreateBatchConsumerResponseBodyItems) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateBatchConsumerResponseBodyItems) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *CreateBatchConsumerResponseBodyItems) GetConsumerGroupId() *string {
	return s.ConsumerGroupId
}

func (s *CreateBatchConsumerResponseBodyItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *CreateBatchConsumerResponseBodyItems) GetConsumerTag() *string {
	return s.ConsumerTag
}

func (s *CreateBatchConsumerResponseBodyItems) GetGwClusterId() *string {
	return s.GwClusterId
}

func (s *CreateBatchConsumerResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *CreateBatchConsumerResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *CreateBatchConsumerResponseBodyItems) SetApiKey(v string) *CreateBatchConsumerResponseBodyItems {
	s.ApiKey = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetApiStatus(v string) *CreateBatchConsumerResponseBodyItems {
	s.ApiStatus = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetConsumerGroupId(v string) *CreateBatchConsumerResponseBodyItems {
	s.ConsumerGroupId = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetConsumerId(v string) *CreateBatchConsumerResponseBodyItems {
	s.ConsumerId = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetConsumerTag(v string) *CreateBatchConsumerResponseBodyItems {
	s.ConsumerTag = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetGwClusterId(v string) *CreateBatchConsumerResponseBodyItems {
	s.GwClusterId = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetName(v string) *CreateBatchConsumerResponseBodyItems {
	s.Name = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetStatus(v string) *CreateBatchConsumerResponseBodyItems {
	s.Status = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
