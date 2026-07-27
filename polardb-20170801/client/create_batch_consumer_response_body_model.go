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
	// Indicates whether the key is active.
	//
	// example:
	//
	// true
	Active *bool `json:"Active,omitempty" xml:"Active,omitempty"`
	// The full API key. Returned only in this response.
	//
	// example:
	//
	// xxxxxxxx
	ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	// The MD5 hash of the API key.
	//
	// example:
	//
	// 0769a11c2d474f96fbb527f8e273d3de
	ApiKeyMd5 *string `json:"ApiKeyMd5,omitempty" xml:"ApiKeyMd5,omitempty"`
	// The status of the API key. Default value: Active.
	//
	// example:
	//
	// Active
	ApiKeyStatus *string `json:"ApiKeyStatus,omitempty" xml:"ApiKeyStatus,omitempty"`
	// The API key status. Default value: Active.
	//
	// example:
	//
	// Active
	ApiStatus *string `json:"ApiStatus,omitempty" xml:"ApiStatus,omitempty"`
	// The budget limit, which equals the number of credits per package.
	//
	// example:
	//
	// 3000
	BudgetLimit *int64 `json:"BudgetLimit,omitempty" xml:"BudgetLimit,omitempty"`
	// The budget policy ID. Each key has an independent budget policy.
	//
	// example:
	//
	// 023aacc1effc4b56bb154bfbec6ba9**
	BudgetPolicyId *string `json:"BudgetPolicyId,omitempty" xml:"BudgetPolicyId,omitempty"`
	// The used quota.
	//
	// example:
	//
	// 0
	BudgetUsed *int64 `json:"BudgetUsed,omitempty" xml:"BudgetUsed,omitempty"`
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
	// The application description or remarks.
	//
	// example:
	//
	// myapp
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expiration time.
	//
	// example:
	//
	// 2025-06-25T09:37:10Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Indicates whether the key is expired.
	//
	// example:
	//
	// false
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The last modification time.
	//
	// example:
	//
	// 2026-01-04T16:09:29+08:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
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

func (s *CreateBatchConsumerResponseBodyItems) GetActive() *bool {
	return s.Active
}

func (s *CreateBatchConsumerResponseBodyItems) GetApiKey() *string {
	return s.ApiKey
}

func (s *CreateBatchConsumerResponseBodyItems) GetApiKeyMd5() *string {
	return s.ApiKeyMd5
}

func (s *CreateBatchConsumerResponseBodyItems) GetApiKeyStatus() *string {
	return s.ApiKeyStatus
}

func (s *CreateBatchConsumerResponseBodyItems) GetApiStatus() *string {
	return s.ApiStatus
}

func (s *CreateBatchConsumerResponseBodyItems) GetBudgetLimit() *int64 {
	return s.BudgetLimit
}

func (s *CreateBatchConsumerResponseBodyItems) GetBudgetPolicyId() *string {
	return s.BudgetPolicyId
}

func (s *CreateBatchConsumerResponseBodyItems) GetBudgetUsed() *int64 {
	return s.BudgetUsed
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

func (s *CreateBatchConsumerResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *CreateBatchConsumerResponseBodyItems) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateBatchConsumerResponseBodyItems) GetExpired() *bool {
	return s.Expired
}

func (s *CreateBatchConsumerResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateBatchConsumerResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
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

func (s *CreateBatchConsumerResponseBodyItems) SetActive(v bool) *CreateBatchConsumerResponseBodyItems {
	s.Active = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetApiKey(v string) *CreateBatchConsumerResponseBodyItems {
	s.ApiKey = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetApiKeyMd5(v string) *CreateBatchConsumerResponseBodyItems {
	s.ApiKeyMd5 = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetApiKeyStatus(v string) *CreateBatchConsumerResponseBodyItems {
	s.ApiKeyStatus = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetApiStatus(v string) *CreateBatchConsumerResponseBodyItems {
	s.ApiStatus = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetBudgetLimit(v int64) *CreateBatchConsumerResponseBodyItems {
	s.BudgetLimit = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetBudgetPolicyId(v string) *CreateBatchConsumerResponseBodyItems {
	s.BudgetPolicyId = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetBudgetUsed(v int64) *CreateBatchConsumerResponseBodyItems {
	s.BudgetUsed = &v
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

func (s *CreateBatchConsumerResponseBodyItems) SetDescription(v string) *CreateBatchConsumerResponseBodyItems {
	s.Description = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetExpireTime(v string) *CreateBatchConsumerResponseBodyItems {
	s.ExpireTime = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetExpired(v bool) *CreateBatchConsumerResponseBodyItems {
	s.Expired = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetGmtCreated(v string) *CreateBatchConsumerResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *CreateBatchConsumerResponseBodyItems) SetGmtModified(v string) *CreateBatchConsumerResponseBodyItems {
	s.GmtModified = &v
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
