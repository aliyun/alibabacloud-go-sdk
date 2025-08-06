// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserSolutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *ListUserSolutionsResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*ListUserSolutionsResponseBodyData) *ListUserSolutionsResponseBody
	GetData() []*ListUserSolutionsResponseBodyData
	SetPageSize(v int32) *ListUserSolutionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserSolutionsResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *ListUserSolutionsResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListUserSolutionsResponseBody
	GetTotalPageNum() *int32
}

type ListUserSolutionsResponseBody struct {
	// example:
	//
	// 8
	CurrentPageNum *int32                               `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListUserSolutionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 344
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserSolutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserSolutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListUserSolutionsResponseBody) GetData() []*ListUserSolutionsResponseBodyData {
	return s.Data
}

func (s *ListUserSolutionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserSolutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserSolutionsResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListUserSolutionsResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListUserSolutionsResponseBody) SetCurrentPageNum(v int32) *ListUserSolutionsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListUserSolutionsResponseBody) SetData(v []*ListUserSolutionsResponseBodyData) *ListUserSolutionsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserSolutionsResponseBody) SetPageSize(v int32) *ListUserSolutionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserSolutionsResponseBody) SetRequestId(v string) *ListUserSolutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserSolutionsResponseBody) SetTotalItemNum(v int32) *ListUserSolutionsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserSolutionsResponseBody) SetTotalPageNum(v int32) *ListUserSolutionsResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListUserSolutionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserSolutionsResponseBodyData struct {
	// example:
	//
	// S20210924151843000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.logo
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 164454443222
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// P20210924151843000001
	DeliveryOrderBizId *string `json:"DeliveryOrderBizId,omitempty" xml:"DeliveryOrderBizId,omitempty"`
	// example:
	//
	// A20210924151843000001
	IntentionAssignBizId *string `json:"IntentionAssignBizId,omitempty" xml:"IntentionAssignBizId,omitempty"`
	// example:
	//
	// I20210924151843000001
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
	// 164454443222
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1219541161213057
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserSolutionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserSolutionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserSolutionsResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *ListUserSolutionsResponseBodyData) GetBizType() *string {
	return s.BizType
}

func (s *ListUserSolutionsResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUserSolutionsResponseBodyData) GetDeliveryOrderBizId() *string {
	return s.DeliveryOrderBizId
}

func (s *ListUserSolutionsResponseBodyData) GetIntentionAssignBizId() *string {
	return s.IntentionAssignBizId
}

func (s *ListUserSolutionsResponseBodyData) GetIntentionBizId() *string {
	return s.IntentionBizId
}

func (s *ListUserSolutionsResponseBodyData) GetPartnerCode() *string {
	return s.PartnerCode
}

func (s *ListUserSolutionsResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *ListUserSolutionsResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListUserSolutionsResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListUserSolutionsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListUserSolutionsResponseBodyData) SetBizId(v string) *ListUserSolutionsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetBizType(v string) *ListUserSolutionsResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetCreateTime(v int64) *ListUserSolutionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetDeliveryOrderBizId(v string) *ListUserSolutionsResponseBodyData {
	s.DeliveryOrderBizId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetIntentionAssignBizId(v string) *ListUserSolutionsResponseBodyData {
	s.IntentionAssignBizId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetIntentionBizId(v string) *ListUserSolutionsResponseBodyData {
	s.IntentionBizId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetPartnerCode(v string) *ListUserSolutionsResponseBodyData {
	s.PartnerCode = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetReason(v string) *ListUserSolutionsResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetStatus(v int32) *ListUserSolutionsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetUpdateTime(v int64) *ListUserSolutionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) SetUserId(v string) *ListUserSolutionsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUserSolutionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
