// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserDetailSolutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *ListUserDetailSolutionsResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*ListUserDetailSolutionsResponseBodyData) *ListUserDetailSolutionsResponseBody
	GetData() []*ListUserDetailSolutionsResponseBodyData
	SetPageSize(v int32) *ListUserDetailSolutionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserDetailSolutionsResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *ListUserDetailSolutionsResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListUserDetailSolutionsResponseBody
	GetTotalPageNum() *int32
}

type ListUserDetailSolutionsResponseBody struct {
	// example:
	//
	// 5
	CurrentPageNum *int32                                     `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListUserDetailSolutionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2174AA97-56FB-50FA-B243-0460B9E4CE0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserDetailSolutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserDetailSolutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserDetailSolutionsResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListUserDetailSolutionsResponseBody) GetData() []*ListUserDetailSolutionsResponseBodyData {
	return s.Data
}

func (s *ListUserDetailSolutionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserDetailSolutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserDetailSolutionsResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListUserDetailSolutionsResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListUserDetailSolutionsResponseBody) SetCurrentPageNum(v int32) *ListUserDetailSolutionsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetData(v []*ListUserDetailSolutionsResponseBodyData) *ListUserDetailSolutionsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetPageSize(v int32) *ListUserDetailSolutionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetRequestId(v string) *ListUserDetailSolutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetTotalItemNum(v int32) *ListUserDetailSolutionsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) SetTotalPageNum(v int32) *ListUserDetailSolutionsResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBody) Validate() error {
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

type ListUserDetailSolutionsResponseBodyData struct {
	// example:
	//
	// S20211222161651000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.wangwen
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 15556223433
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// P20211222161651000001
	DeliveryOrderBizId *string `json:"DeliveryOrderBizId,omitempty" xml:"DeliveryOrderBizId,omitempty"`
	// example:
	//
	// {}
	ExtInfo *string `json:"ExtInfo,omitempty" xml:"ExtInfo,omitempty"`
	// example:
	//
	// A20211222161651000001
	IntentionAssignBizId *string `json:"IntentionAssignBizId,omitempty" xml:"IntentionAssignBizId,omitempty"`
	// example:
	//
	// I20211222161651000001
	IntentionBizId *string `json:"IntentionBizId,omitempty" xml:"IntentionBizId,omitempty"`
	// example:
	//
	// jinsan
	PartnerCode *string `json:"PartnerCode,omitempty" xml:"PartnerCode,omitempty"`
	Reason      *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 15556223433
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1219541161213057
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserDetailSolutionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserDetailSolutionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserDetailSolutionsResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *ListUserDetailSolutionsResponseBodyData) GetBizType() *string {
	return s.BizType
}

func (s *ListUserDetailSolutionsResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUserDetailSolutionsResponseBodyData) GetDeliveryOrderBizId() *string {
	return s.DeliveryOrderBizId
}

func (s *ListUserDetailSolutionsResponseBodyData) GetExtInfo() *string {
	return s.ExtInfo
}

func (s *ListUserDetailSolutionsResponseBodyData) GetIntentionAssignBizId() *string {
	return s.IntentionAssignBizId
}

func (s *ListUserDetailSolutionsResponseBodyData) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListUserDetailSolutionsResponseBodyData) GetPartnerCode() *string {
	return s.PartnerCode
}

func (s *ListUserDetailSolutionsResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *ListUserDetailSolutionsResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListUserDetailSolutionsResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListUserDetailSolutionsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListUserDetailSolutionsResponseBodyData) SetBizId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetBizType(v string) *ListUserDetailSolutionsResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetCreateTime(v int64) *ListUserDetailSolutionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetDeliveryOrderBizId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.DeliveryOrderBizId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetExtInfo(v string) *ListUserDetailSolutionsResponseBodyData {
	s.ExtInfo = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetIntentionAssignBizId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.IntentionAssignBizId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetIntentionBizId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetPartnerCode(v string) *ListUserDetailSolutionsResponseBodyData {
	s.PartnerCode = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetReason(v string) *ListUserDetailSolutionsResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetStatus(v int32) *ListUserDetailSolutionsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetUpdateTime(v int64) *ListUserDetailSolutionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) SetUserId(v string) *ListUserDetailSolutionsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUserDetailSolutionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
