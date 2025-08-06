// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserIntentionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *ListUserIntentionsResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*ListUserIntentionsResponseBodyData) *ListUserIntentionsResponseBody
	GetData() []*ListUserIntentionsResponseBodyData
	SetPageSize(v int32) *ListUserIntentionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListUserIntentionsResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int32) *ListUserIntentionsResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListUserIntentionsResponseBody
	GetTotalPageNum() *int32
}

type ListUserIntentionsResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListUserIntentionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
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
	// 2
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListUserIntentionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserIntentionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserIntentionsResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListUserIntentionsResponseBody) GetData() []*ListUserIntentionsResponseBodyData {
	return s.Data
}

func (s *ListUserIntentionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUserIntentionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserIntentionsResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListUserIntentionsResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListUserIntentionsResponseBody) SetCurrentPageNum(v int32) *ListUserIntentionsResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListUserIntentionsResponseBody) SetData(v []*ListUserIntentionsResponseBodyData) *ListUserIntentionsResponseBody {
	s.Data = v
	return s
}

func (s *ListUserIntentionsResponseBody) SetPageSize(v int32) *ListUserIntentionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListUserIntentionsResponseBody) SetRequestId(v string) *ListUserIntentionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserIntentionsResponseBody) SetTotalItemNum(v int32) *ListUserIntentionsResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListUserIntentionsResponseBody) SetTotalPageNum(v int32) *ListUserIntentionsResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListUserIntentionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListUserIntentionsResponseBodyData struct {
	Area *string `json:"Area,omitempty" xml:"Area,omitempty"`
	// example:
	//
	// I100000033443
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.lgo
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// example:
	//
	// 2022-01-24 15:43:58
	CreateTime  *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Ext         *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 18000000000
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2022-01-24 15:43:58
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// example:
	//
	// 1219541161213057
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s ListUserIntentionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUserIntentionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUserIntentionsResponseBodyData) GetArea() *string {
	return s.Area
}

func (s *ListUserIntentionsResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *ListUserIntentionsResponseBodyData) GetBizType() *string {
	return s.BizType
}

func (s *ListUserIntentionsResponseBodyData) GetContactName() *string {
	return s.ContactName
}

func (s *ListUserIntentionsResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListUserIntentionsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListUserIntentionsResponseBodyData) GetExt() *string {
	return s.Ext
}

func (s *ListUserIntentionsResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *ListUserIntentionsResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *ListUserIntentionsResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListUserIntentionsResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListUserIntentionsResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *ListUserIntentionsResponseBodyData) SetArea(v string) *ListUserIntentionsResponseBodyData {
	s.Area = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetBizId(v string) *ListUserIntentionsResponseBodyData {
	s.BizId = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetBizType(v string) *ListUserIntentionsResponseBodyData {
	s.BizType = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetContactName(v string) *ListUserIntentionsResponseBodyData {
	s.ContactName = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetCreateTime(v int64) *ListUserIntentionsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetDescription(v string) *ListUserIntentionsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetExt(v string) *ListUserIntentionsResponseBodyData {
	s.Ext = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetMobile(v string) *ListUserIntentionsResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetReason(v string) *ListUserIntentionsResponseBodyData {
	s.Reason = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetStatus(v int32) *ListUserIntentionsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetUpdateTime(v int64) *ListUserIntentionsResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) SetUserId(v string) *ListUserIntentionsResponseBodyData {
	s.UserId = &v
	return s
}

func (s *ListUserIntentionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
