// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMpsSchedulerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryMpsSchedulerListResponseBody
	GetRequestId() *string
	SetResultCode(v string) *QueryMpsSchedulerListResponseBody
	GetResultCode() *string
	SetResultContent(v *QueryMpsSchedulerListResponseBodyResultContent) *QueryMpsSchedulerListResponseBody
	GetResultContent() *QueryMpsSchedulerListResponseBodyResultContent
	SetResultMessage(v string) *QueryMpsSchedulerListResponseBody
	GetResultMessage() *string
}

type QueryMpsSchedulerListResponseBody struct {
	RequestId     *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultCode    *string                                         `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	ResultContent *QueryMpsSchedulerListResponseBodyResultContent `json:"ResultContent,omitempty" xml:"ResultContent,omitempty" type:"Struct"`
	ResultMessage *string                                         `json:"ResultMessage,omitempty" xml:"ResultMessage,omitempty"`
}

func (s QueryMpsSchedulerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryMpsSchedulerListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryMpsSchedulerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryMpsSchedulerListResponseBody) GetResultCode() *string {
	return s.ResultCode
}

func (s *QueryMpsSchedulerListResponseBody) GetResultContent() *QueryMpsSchedulerListResponseBodyResultContent {
	return s.ResultContent
}

func (s *QueryMpsSchedulerListResponseBody) GetResultMessage() *string {
	return s.ResultMessage
}

func (s *QueryMpsSchedulerListResponseBody) SetRequestId(v string) *QueryMpsSchedulerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBody) SetResultCode(v string) *QueryMpsSchedulerListResponseBody {
	s.ResultCode = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBody) SetResultContent(v *QueryMpsSchedulerListResponseBodyResultContent) *QueryMpsSchedulerListResponseBody {
	s.ResultContent = v
	return s
}

func (s *QueryMpsSchedulerListResponseBody) SetResultMessage(v string) *QueryMpsSchedulerListResponseBody {
	s.ResultMessage = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBody) Validate() error {
	if s.ResultContent != nil {
		if err := s.ResultContent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMpsSchedulerListResponseBodyResultContent struct {
	Data *QueryMpsSchedulerListResponseBodyResultContentData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
}

func (s QueryMpsSchedulerListResponseBodyResultContent) String() string {
	return dara.Prettify(s)
}

func (s QueryMpsSchedulerListResponseBodyResultContent) GoString() string {
	return s.String()
}

func (s *QueryMpsSchedulerListResponseBodyResultContent) GetData() *QueryMpsSchedulerListResponseBodyResultContentData {
	return s.Data
}

func (s *QueryMpsSchedulerListResponseBodyResultContent) SetData(v *QueryMpsSchedulerListResponseBodyResultContentData) *QueryMpsSchedulerListResponseBodyResultContent {
	s.Data = v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContent) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryMpsSchedulerListResponseBodyResultContentData struct {
	List       []*QueryMpsSchedulerListResponseBodyResultContentDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	TotalCount *int32                                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s QueryMpsSchedulerListResponseBodyResultContentData) String() string {
	return dara.Prettify(s)
}

func (s QueryMpsSchedulerListResponseBodyResultContentData) GoString() string {
	return s.String()
}

func (s *QueryMpsSchedulerListResponseBodyResultContentData) GetList() []*QueryMpsSchedulerListResponseBodyResultContentDataList {
	return s.List
}

func (s *QueryMpsSchedulerListResponseBodyResultContentData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *QueryMpsSchedulerListResponseBodyResultContentData) SetList(v []*QueryMpsSchedulerListResponseBodyResultContentDataList) *QueryMpsSchedulerListResponseBodyResultContentData {
	s.List = v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentData) SetTotalCount(v int32) *QueryMpsSchedulerListResponseBodyResultContentData {
	s.TotalCount = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryMpsSchedulerListResponseBodyResultContentDataList struct {
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

func (s QueryMpsSchedulerListResponseBodyResultContentDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryMpsSchedulerListResponseBodyResultContentDataList) GoString() string {
	return s.String()
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetCreateType() *int32 {
	return s.CreateType
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetDeliveryType() *int32 {
	return s.DeliveryType
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetExecutedStatus() *string {
	return s.ExecutedStatus
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetParentId() *string {
	return s.ParentId
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetPushContent() *string {
	return s.PushContent
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetPushTime() *int64 {
	return s.PushTime
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetPushTitle() *string {
	return s.PushTitle
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetStrategyType() *int32 {
	return s.StrategyType
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetType() *int32 {
	return s.Type
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) GetUniqueId() *string {
	return s.UniqueId
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetCreateType(v int32) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.CreateType = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetDeliveryType(v int32) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.DeliveryType = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetExecutedStatus(v string) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.ExecutedStatus = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetGmtCreate(v int64) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.GmtCreate = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetParentId(v string) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.ParentId = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetPushContent(v string) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.PushContent = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetPushTime(v int64) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.PushTime = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetPushTitle(v string) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.PushTitle = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetStrategyType(v int32) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.StrategyType = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetType(v int32) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.Type = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) SetUniqueId(v string) *QueryMpsSchedulerListResponseBodyResultContentDataList {
	s.UniqueId = &v
	return s
}

func (s *QueryMpsSchedulerListResponseBodyResultContentDataList) Validate() error {
	return dara.Validate(s)
}
