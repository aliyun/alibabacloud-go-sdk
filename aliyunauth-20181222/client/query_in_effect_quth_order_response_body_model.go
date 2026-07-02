// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInEffectQuthOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *QueryInEffectQuthOrderResponseBody
	GetCode() *int32
	SetData(v *QueryInEffectQuthOrderResponseBodyData) *QueryInEffectQuthOrderResponseBody
	GetData() *QueryInEffectQuthOrderResponseBodyData
	SetMessage(v string) *QueryInEffectQuthOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInEffectQuthOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInEffectQuthOrderResponseBody
	GetSuccess() *bool
}

type QueryInEffectQuthOrderResponseBody struct {
	Code      *int32                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *QueryInEffectQuthOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInEffectQuthOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInEffectQuthOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInEffectQuthOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *QueryInEffectQuthOrderResponseBody) GetData() *QueryInEffectQuthOrderResponseBodyData {
	return s.Data
}

func (s *QueryInEffectQuthOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInEffectQuthOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInEffectQuthOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInEffectQuthOrderResponseBody) SetCode(v int32) *QueryInEffectQuthOrderResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBody) SetData(v *QueryInEffectQuthOrderResponseBodyData) *QueryInEffectQuthOrderResponseBody {
	s.Data = v
	return s
}

func (s *QueryInEffectQuthOrderResponseBody) SetMessage(v string) *QueryInEffectQuthOrderResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBody) SetRequestId(v string) *QueryInEffectQuthOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBody) SetSuccess(v bool) *QueryInEffectQuthOrderResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInEffectQuthOrderResponseBodyData struct {
	Count *int32                                        `json:"Count,omitempty" xml:"Count,omitempty"`
	List  []*QueryInEffectQuthOrderResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s QueryInEffectQuthOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryInEffectQuthOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInEffectQuthOrderResponseBodyData) GetCount() *int32 {
	return s.Count
}

func (s *QueryInEffectQuthOrderResponseBodyData) GetList() []*QueryInEffectQuthOrderResponseBodyDataList {
	return s.List
}

func (s *QueryInEffectQuthOrderResponseBodyData) SetCount(v int32) *QueryInEffectQuthOrderResponseBodyData {
	s.Count = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyData) SetList(v []*QueryInEffectQuthOrderResponseBodyDataList) *QueryInEffectQuthOrderResponseBodyData {
	s.List = v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyData) Validate() error {
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

type QueryInEffectQuthOrderResponseBodyDataList struct {
	AuthItemRecordGroupItemDTOS []*QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS `json:"AuthItemRecordGroupItemDTOS,omitempty" xml:"AuthItemRecordGroupItemDTOS,omitempty" type:"Repeated"`
	CreatedTime                 *string                                                                  `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	LastModifyTime              *string                                                                  `json:"LastModifyTime,omitempty" xml:"LastModifyTime,omitempty"`
	OperateTimes                []*QueryInEffectQuthOrderResponseBodyDataListOperateTimes                `json:"OperateTimes,omitempty" xml:"OperateTimes,omitempty" type:"Repeated"`
	OrderVid                    *string                                                                  `json:"OrderVid,omitempty" xml:"OrderVid,omitempty"`
	Status                      *int32                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s QueryInEffectQuthOrderResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s QueryInEffectQuthOrderResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) GetAuthItemRecordGroupItemDTOS() []*QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS {
	return s.AuthItemRecordGroupItemDTOS
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) GetLastModifyTime() *string {
	return s.LastModifyTime
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) GetOperateTimes() []*QueryInEffectQuthOrderResponseBodyDataListOperateTimes {
	return s.OperateTimes
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) GetOrderVid() *string {
	return s.OrderVid
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) GetStatus() *int32 {
	return s.Status
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) SetAuthItemRecordGroupItemDTOS(v []*QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) *QueryInEffectQuthOrderResponseBodyDataList {
	s.AuthItemRecordGroupItemDTOS = v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) SetCreatedTime(v string) *QueryInEffectQuthOrderResponseBodyDataList {
	s.CreatedTime = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) SetLastModifyTime(v string) *QueryInEffectQuthOrderResponseBodyDataList {
	s.LastModifyTime = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) SetOperateTimes(v []*QueryInEffectQuthOrderResponseBodyDataListOperateTimes) *QueryInEffectQuthOrderResponseBodyDataList {
	s.OperateTimes = v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) SetOrderVid(v string) *QueryInEffectQuthOrderResponseBodyDataList {
	s.OrderVid = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) SetStatus(v int32) *QueryInEffectQuthOrderResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataList) Validate() error {
	if s.AuthItemRecordGroupItemDTOS != nil {
		for _, item := range s.AuthItemRecordGroupItemDTOS {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OperateTimes != nil {
		for _, item := range s.OperateTimes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS struct {
	AuthitemID     *string `json:"AuthitemID,omitempty" xml:"AuthitemID,omitempty"`
	GmtCreated     *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	Msg            *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	Name           *string `json:"Name,omitempty" xml:"Name,omitempty"`
	RemarkDataJson *string `json:"RemarkDataJson,omitempty" xml:"RemarkDataJson,omitempty"`
	Status         *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	Vid            *string `json:"Vid,omitempty" xml:"Vid,omitempty"`
}

func (s QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) String() string {
	return dara.Prettify(s)
}

func (s QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) GoString() string {
	return s.String()
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) GetAuthitemID() *string {
	return s.AuthitemID
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) GetMsg() *string {
	return s.Msg
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) GetName() *string {
	return s.Name
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) GetRemarkDataJson() *string {
	return s.RemarkDataJson
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) GetStatus() *int32 {
	return s.Status
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) GetVid() *string {
	return s.Vid
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) SetAuthitemID(v string) *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS {
	s.AuthitemID = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) SetGmtCreated(v string) *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS {
	s.GmtCreated = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) SetMsg(v string) *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS {
	s.Msg = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) SetName(v string) *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS {
	s.Name = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) SetRemarkDataJson(v string) *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS {
	s.RemarkDataJson = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) SetStatus(v int32) *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS {
	s.Status = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) SetVid(v string) *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS {
	s.Vid = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListAuthItemRecordGroupItemDTOS) Validate() error {
	return dara.Validate(s)
}

type QueryInEffectQuthOrderResponseBodyDataListOperateTimes struct {
	End   *string `json:"End,omitempty" xml:"End,omitempty"`
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s QueryInEffectQuthOrderResponseBodyDataListOperateTimes) String() string {
	return dara.Prettify(s)
}

func (s QueryInEffectQuthOrderResponseBodyDataListOperateTimes) GoString() string {
	return s.String()
}

func (s *QueryInEffectQuthOrderResponseBodyDataListOperateTimes) GetEnd() *string {
	return s.End
}

func (s *QueryInEffectQuthOrderResponseBodyDataListOperateTimes) GetStart() *string {
	return s.Start
}

func (s *QueryInEffectQuthOrderResponseBodyDataListOperateTimes) SetEnd(v string) *QueryInEffectQuthOrderResponseBodyDataListOperateTimes {
	s.End = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListOperateTimes) SetStart(v string) *QueryInEffectQuthOrderResponseBodyDataListOperateTimes {
	s.Start = &v
	return s
}

func (s *QueryInEffectQuthOrderResponseBodyDataListOperateTimes) Validate() error {
	return dara.Validate(s)
}
