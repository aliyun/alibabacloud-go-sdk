// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageCampaignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListMessageCampaignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ListMessageCampaignResponseBody
	GetCode() *string
	SetData(v []*ListMessageCampaignResponseBodyData) *ListMessageCampaignResponseBody
	GetData() []*ListMessageCampaignResponseBodyData
	SetMessage(v string) *ListMessageCampaignResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListMessageCampaignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListMessageCampaignResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListMessageCampaignResponseBody
	GetTotal() *int64
}

type ListMessageCampaignResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - A value of OK indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// Example
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data object.
	Data []*ListMessageCampaignResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 23**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: successful.
	//
	// - false: failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 70
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListMessageCampaignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMessageCampaignResponseBody) GoString() string {
	return s.String()
}

func (s *ListMessageCampaignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListMessageCampaignResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMessageCampaignResponseBody) GetData() []*ListMessageCampaignResponseBodyData {
	return s.Data
}

func (s *ListMessageCampaignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListMessageCampaignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMessageCampaignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListMessageCampaignResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListMessageCampaignResponseBody) SetAccessDeniedDetail(v string) *ListMessageCampaignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListMessageCampaignResponseBody) SetCode(v string) *ListMessageCampaignResponseBody {
	s.Code = &v
	return s
}

func (s *ListMessageCampaignResponseBody) SetData(v []*ListMessageCampaignResponseBodyData) *ListMessageCampaignResponseBody {
	s.Data = v
	return s
}

func (s *ListMessageCampaignResponseBody) SetMessage(v string) *ListMessageCampaignResponseBody {
	s.Message = &v
	return s
}

func (s *ListMessageCampaignResponseBody) SetRequestId(v string) *ListMessageCampaignResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMessageCampaignResponseBody) SetSuccess(v bool) *ListMessageCampaignResponseBody {
	s.Success = &v
	return s
}

func (s *ListMessageCampaignResponseBody) SetTotal(v int64) *ListMessageCampaignResponseBody {
	s.Total = &v
	return s
}

func (s *ListMessageCampaignResponseBody) Validate() error {
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

type ListMessageCampaignResponseBodyData struct {
	// The ID of the Meta ad account.
	//
	// example:
	//
	// 2339**
	AdAccountId *string `json:"AdAccountId,omitempty" xml:"AdAccountId,omitempty"`
	// The budget.
	//
	// example:
	//
	// 62
	Budget *int64 `json:"Budget,omitempty" xml:"Budget,omitempty"`
	// The budget type.
	//
	// example:
	//
	// daily
	BudgetType *string `json:"BudgetType,omitempty" xml:"BudgetType,omitempty"`
	// The campaign ID.
	//
	// example:
	//
	// 233**
	CampaignId *string `json:"CampaignId,omitempty" xml:"CampaignId,omitempty"`
	// The name of the campaign.
	//
	// example:
	//
	// campaign-name
	CampaignName *string `json:"CampaignName,omitempty" xml:"CampaignName,omitempty"`
	// The time when the campaign was created.
	//
	// example:
	//
	// 173029392838
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Page ID for Messenger.
	//
	// example:
	//
	// 238***
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The status of the campaign.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListMessageCampaignResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMessageCampaignResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMessageCampaignResponseBodyData) GetAdAccountId() *string {
	return s.AdAccountId
}

func (s *ListMessageCampaignResponseBodyData) GetBudget() *int64 {
	return s.Budget
}

func (s *ListMessageCampaignResponseBodyData) GetBudgetType() *string {
	return s.BudgetType
}

func (s *ListMessageCampaignResponseBodyData) GetCampaignId() *string {
	return s.CampaignId
}

func (s *ListMessageCampaignResponseBodyData) GetCampaignName() *string {
	return s.CampaignName
}

func (s *ListMessageCampaignResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListMessageCampaignResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *ListMessageCampaignResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListMessageCampaignResponseBodyData) SetAdAccountId(v string) *ListMessageCampaignResponseBodyData {
	s.AdAccountId = &v
	return s
}

func (s *ListMessageCampaignResponseBodyData) SetBudget(v int64) *ListMessageCampaignResponseBodyData {
	s.Budget = &v
	return s
}

func (s *ListMessageCampaignResponseBodyData) SetBudgetType(v string) *ListMessageCampaignResponseBodyData {
	s.BudgetType = &v
	return s
}

func (s *ListMessageCampaignResponseBodyData) SetCampaignId(v string) *ListMessageCampaignResponseBodyData {
	s.CampaignId = &v
	return s
}

func (s *ListMessageCampaignResponseBodyData) SetCampaignName(v string) *ListMessageCampaignResponseBodyData {
	s.CampaignName = &v
	return s
}

func (s *ListMessageCampaignResponseBodyData) SetCreateTime(v int64) *ListMessageCampaignResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListMessageCampaignResponseBodyData) SetPageId(v string) *ListMessageCampaignResponseBodyData {
	s.PageId = &v
	return s
}

func (s *ListMessageCampaignResponseBodyData) SetStatus(v string) *ListMessageCampaignResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListMessageCampaignResponseBodyData) Validate() error {
	return dara.Validate(s)
}
