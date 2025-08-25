// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushSchedulerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPushSchedulerListResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryPushSchedulerListResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryPushSchedulerListResponseBodyResultContent) *QueryPushSchedulerListResponseBody
	GetResultContent() *QueryPushSchedulerListResponseBodyResultContent
	SetResultMessage(v string) *QueryPushSchedulerListResponseBody
	GetResultMessage() *string
}

type QueryPushSchedulerListResponseBody struct {
	RequestId     *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                          `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryPushSchedulerListResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                          `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryPushSchedulerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPushSchedulerListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPushSchedulerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPushSchedulerListResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryPushSchedulerListResponseBody) GetResultContent() *QueryPushSchedulerListResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryPushSchedulerListResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryPushSchedulerListResponseBody) SetRequestId(v string) *QueryPushSchedulerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPushSchedulerListResponseBody) SetResultCode(v string) *QueryPushSchedulerListResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryPushSchedulerListResponseBody) SetResultContent(v *QueryPushSchedulerListResponseBodyResultContent) *QueryPushSchedulerListResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryPushSchedulerListResponseBody) SetResultMessage(v string) *QueryPushSchedulerListResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryPushSchedulerListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryPushSchedulerListResponseBodyResultContent struct {
	Data *QueryPushSchedulerListResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryPushSchedulerListResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryPushSchedulerListResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryPushSchedulerListResponseBodyResultContent) GetData() *QueryPushSchedulerListResponseBodyResultContentData {
	return s.Data
}

func (s *QueryPushSchedulerListResponseBodyResultContent) SetData(v *QueryPushSchedulerListResponseBodyResultContentData) *QueryPushSchedulerListResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContent) Validate() error {
	return dara.Validate(s)
}

type QueryPushSchedulerListResponseBodyResultContentData struct {
	List       []*QueryPushSchedulerListResponseBodyResultContentDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int32                                                     `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryPushSchedulerListResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s QueryPushSchedulerListResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *QueryPushSchedulerListResponseBodyResultContentData) GetList() []*QueryPushSchedulerListResponseBodyResultContentDataList {
	return s.List
}

func (s *QueryPushSchedulerListResponseBodyResultContentData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryPushSchedulerListResponseBodyResultContentData) SetList(v []*QueryPushSchedulerListResponseBodyResultContentDataList) *QueryPushSchedulerListResponseBodyResultContentData {
	s.List = v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentData) SetTotalCount(v int32) *QueryPushSchedulerListResponseBodyResultContentData {
	s.TotalCount = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentData) Validate() error {
	return dara.Validate(s)
}

type QueryPushSchedulerListResponseBodyResultContentDataList struct {
	CreateType     *int32  `json:"CreateType,omitempty" xml:"CreateType,omitempty"`
	DeliveryType   *int32  `json:"DeliveryType,omitempty" xml:"DeliveryType,omitempty"`
	ExecutedStatus *string `json:"ExecutedStatus,omitempty" xml:"ExecutedStatus,omitempty"`
	GmtCreate      *int64  `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	ParentId       *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	PushContent    *string `json:"PushContent,omitempty" xml:"PushContent,omitempty"`
	PushTime       *int64  `json:"PushTime,omitempty" xml:"PushTime,omitempty"`
	PushTitle      *string `json:"PushTitle,omitempty" xml:"PushTitle,omitempty"`
	StrategyType   *int32  `json:"StrategyType,omitempty" xml:"StrategyType,omitempty"`
	Type           *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
	UniqueId       *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s QueryPushSchedulerListResponseBodyResultContentDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryPushSchedulerListResponseBodyResultContentDataList) GoString() string {
	return s.String()
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetCreateType() *int32 {
	return s.CreateType
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetDeliveryType() *int32 {
	return s.DeliveryType
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetExecutedStatus() *string {
	return s.ExecutedStatus
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetParentId() *string {
	return s.ParentId
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetPushContent() *string {
	return s.PushContent
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetPushTime() *int64 {
	return s.PushTime
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetPushTitle() *string {
	return s.PushTitle
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetType() *int32 {
	return s.Type
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetCreateType(v int32) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.CreateType = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetDeliveryType(v int32) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.DeliveryType = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetExecutedStatus(v string) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.ExecutedStatus = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetGmtCreate(v int64) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.GmtCreate = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetParentId(v string) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.ParentId = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetPushContent(v string) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.PushContent = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetPushTime(v int64) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.PushTime = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetPushTitle(v string) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.PushTitle = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetStrategyType(v int32) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.StrategyType = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetType(v int32) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.Type = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) SetUniqueId(v string) *QueryPushSchedulerListResponseBodyResultContentDataList {
	s.UniqueId = &v
	return s
}

func (s *QueryPushSchedulerListResponseBodyResultContentDataList) Validate() error {
	return dara.Validate(s)
}
