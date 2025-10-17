// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultDiagnosisUserDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCallInfo(v *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) *DescribeFaultDiagnosisUserDetailResponseBody
	GetCallInfo() *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo
	SetFactorList(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList) *DescribeFaultDiagnosisUserDetailResponseBody
	GetFactorList() []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList
	SetFaultMetricData(v *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) *DescribeFaultDiagnosisUserDetailResponseBody
	GetFaultMetricData() *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData
	SetNetworkOperators(v []*string) *DescribeFaultDiagnosisUserDetailResponseBody
	GetNetworkOperators() []*string
	SetRequestId(v string) *DescribeFaultDiagnosisUserDetailResponseBody
	GetRequestId() *string
	SetUserDetail(v *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) *DescribeFaultDiagnosisUserDetailResponseBody
	GetUserDetail() *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail
}

type DescribeFaultDiagnosisUserDetailResponseBody struct {
	CallInfo         *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo        `json:"CallInfo,omitempty" xml:"CallInfo,omitempty" type:"Struct"`
	FactorList       []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList    `json:"FactorList,omitempty" xml:"FactorList,omitempty" type:"Repeated"`
	FaultMetricData  *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData `json:"FaultMetricData,omitempty" xml:"FaultMetricData,omitempty" type:"Struct"`
	NetworkOperators []*string                                                    `json:"NetworkOperators,omitempty" xml:"NetworkOperators,omitempty" type:"Repeated"`
	// example:
	//
	// 231470C1-ACFB-4C9F-844F-4CFE1E3804C5
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserDetail *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail `json:"UserDetail,omitempty" xml:"UserDetail,omitempty" type:"Struct"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) GetCallInfo() *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	return s.CallInfo
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) GetFactorList() []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	return s.FactorList
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) GetFaultMetricData() *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData {
	return s.FaultMetricData
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) GetNetworkOperators() []*string {
	return s.NetworkOperators
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) GetUserDetail() *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	return s.UserDetail
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetCallInfo(v *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.CallInfo = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetFactorList(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorList) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.FactorList = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetFaultMetricData(v *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.FaultMetricData = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetNetworkOperators(v []*string) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.NetworkOperators = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetRequestId(v string) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) SetUserDetail(v *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) *DescribeFaultDiagnosisUserDetailResponseBody {
	s.UserDetail = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBody) Validate() error {
	if s.CallInfo != nil {
		if err := s.CallInfo.Validate(); err != nil {
			return err
		}
	}
	if s.FactorList != nil {
		for _, item := range s.FactorList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FaultMetricData != nil {
		if err := s.FaultMetricData.Validate(); err != nil {
			return err
		}
	}
	if s.UserDetail != nil {
		if err := s.UserDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFaultDiagnosisUserDetailResponseBodyCallInfo struct {
	// App ID。
	//
	// example:
	//
	// 0rbd****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// IN
	CallStatus *string `json:"CallStatus,omitempty" xml:"CallStatus,omitempty"`
	// example:
	//
	// 311
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// example:
	//
	// 1620957905
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1620958150
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 100
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GetAppId() *string {
	return s.AppId
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GetCallStatus() *string {
	return s.CallStatus
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GetChannelId() *string {
	return s.ChannelId
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetAppId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.AppId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetCallStatus(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.CallStatus = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetChannelId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.ChannelId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) SetDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo {
	s.Duration = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyCallInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorList struct {
	// example:
	//
	// 1
	FactorId *string `json:"FactorId,omitempty" xml:"FactorId,omitempty"`
	// example:
	//
	// LOCAL
	FaultSource        *string                                                                     `json:"FaultSource,omitempty" xml:"FaultSource,omitempty"`
	RelatedEventDatas  []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas  `json:"RelatedEventDatas,omitempty" xml:"RelatedEventDatas,omitempty" type:"Repeated"`
	RelatedMetricDatas []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas `json:"RelatedMetricDatas,omitempty" xml:"RelatedMetricDatas,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) GetFactorId() *string {
	return s.FactorId
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) GetFaultSource() *string {
	return s.FaultSource
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) GetRelatedEventDatas() []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	return s.RelatedEventDatas
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) GetRelatedMetricDatas() []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	return s.RelatedMetricDatas
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetFactorId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.FactorId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetFaultSource(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.FaultSource = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetRelatedEventDatas(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.RelatedEventDatas = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) SetRelatedMetricDatas(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) *DescribeFaultDiagnosisUserDetailResponseBodyFactorList {
	s.RelatedMetricDatas = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorList) Validate() error {
	if s.RelatedEventDatas != nil {
		for _, item := range s.RelatedEventDatas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RelatedMetricDatas != nil {
		for _, item := range s.RelatedMetricDatas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas struct {
	EventDataItems []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems `json:"EventDataItems,omitempty" xml:"EventDataItems,omitempty" type:"Repeated"`
	// example:
	//
	// SENDER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// 0a497933****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) GetEventDataItems() []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems {
	return s.EventDataItems
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) GetRole() *string {
	return s.Role
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) GetUserId() *string {
	return s.UserId
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetEventDataItems(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.EventDataItems = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetRole(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.Role = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas {
	s.UserId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatas) Validate() error {
	if s.EventDataItems != nil {
		for _, item := range s.EventDataItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems struct {
	EventList []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList `json:"EventList,omitempty" xml:"EventList,omitempty" type:"Repeated"`
	// example:
	//
	// 1614936817
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) GetEventList() []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	return s.EventList
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) GetTs() *int64 {
	return s.Ts
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) SetEventList(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems {
	s.EventList = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) SetTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems {
	s.Ts = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItems) Validate() error {
	if s.EventList != nil {
		for _, item := range s.EventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList struct {
	// example:
	//
	// 开始发布
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// example:
	//
	// USER
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// example:
	//
	// 1614936817
	Ts *int64 `json:"Ts,omitempty" xml:"Ts,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) GetEventName() *string {
	return s.EventName
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) GetEventType() *string {
	return s.EventType
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) GetTs() *int64 {
	return s.Ts
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetEventName(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.EventName = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetEventType(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.EventType = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) SetTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList {
	s.Ts = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedEventDatasEventDataItemsEventList) Validate() error {
	return dara.Validate(s)
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas struct {
	Nodes []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// SENDER
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
	// example:
	//
	// AUDIO_STUCK
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 0a497933****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) GetNodes() []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	return s.Nodes
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) GetRole() *string {
	return s.Role
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) GetType() *string {
	return s.Type
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) GetUserId() *string {
	return s.UserId
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetNodes(v []*DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Nodes = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetRole(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Role = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetType(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.Type = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas {
	s.UserId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatas) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes struct {
	Ext map[string]interface{} `json:"Ext,omitempty" xml:"Ext,omitempty"`
	// example:
	//
	// 1615892596
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 20
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) GetExt() map[string]interface{} {
	return s.Ext
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) GetX() *string {
	return s.X
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) GetY() *string {
	return s.Y
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetExt(v map[string]interface{}) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.Ext = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetX(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) SetY(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes {
	s.Y = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFactorListRelatedMetricDatasNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData struct {
	Nodes []*DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) GetNodes() []*DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes {
	return s.Nodes
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) SetNodes(v []*DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData {
	s.Nodes = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricData) Validate() error {
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes struct {
	// example:
	//
	// 1620957900
	X *string `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 0.4540
	Y *string `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) GetX() *string {
	return s.X
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) GetY() *string {
	return s.Y
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) SetX(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes {
	s.X = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) SetY(v string) *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes {
	s.Y = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyFaultMetricDataNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeFaultDiagnosisUserDetailResponseBodyUserDetail struct {
	// example:
	//
	// 1620957919
	CreatedTs *int64 `json:"CreatedTs,omitempty" xml:"CreatedTs,omitempty"`
	// example:
	//
	// 1620958150
	DestroyedTs *int64 `json:"DestroyedTs,omitempty" xml:"DestroyedTs,omitempty"`
	// example:
	//
	// 231
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// example:
	//
	// 浙江省-杭州市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// example:
	//
	// 4G
	Network *string `json:"Network,omitempty" xml:"Network,omitempty"`
	// example:
	//
	// 231
	OnlineDuration *int64                                                                 `json:"OnlineDuration,omitempty" xml:"OnlineDuration,omitempty"`
	OnlinePeriods  []*DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods `json:"OnlinePeriods,omitempty" xml:"OnlinePeriods,omitempty" type:"Repeated"`
	// example:
	//
	// iOS
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// 1.0.0
	SdkVersion *string `json:"SdkVersion,omitempty" xml:"SdkVersion,omitempty"`
	// example:
	//
	// 0a497933****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetCreatedTs() *int64 {
	return s.CreatedTs
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetDestroyedTs() *int64 {
	return s.DestroyedTs
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetLocation() *string {
	return s.Location
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetNetwork() *string {
	return s.Network
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetOnlineDuration() *int64 {
	return s.OnlineDuration
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetOnlinePeriods() []*DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods {
	return s.OnlinePeriods
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetOs() *string {
	return s.Os
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetSdkVersion() *string {
	return s.SdkVersion
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) GetUserId() *string {
	return s.UserId
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetCreatedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.CreatedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetDestroyedTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.DestroyedTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Duration = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetLocation(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Location = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetNetwork(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Network = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOnlineDuration(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.OnlineDuration = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOnlinePeriods(v []*DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.OnlinePeriods = v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetOs(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.Os = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetSdkVersion(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.SdkVersion = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) SetUserId(v string) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail {
	s.UserId = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetail) Validate() error {
	if s.OnlinePeriods != nil {
		for _, item := range s.OnlinePeriods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods struct {
	// example:
	//
	// 1620957919
	JoinTs *int64 `json:"JoinTs,omitempty" xml:"JoinTs,omitempty"`
	// example:
	//
	// 1620958150
	LeaveTs *int64 `json:"LeaveTs,omitempty" xml:"LeaveTs,omitempty"`
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) GoString() string {
	return s.String()
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) GetJoinTs() *int64 {
	return s.JoinTs
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) GetLeaveTs() *int64 {
	return s.LeaveTs
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) SetJoinTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods {
	s.JoinTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) SetLeaveTs(v int64) *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods {
	s.LeaveTs = &v
	return s
}

func (s *DescribeFaultDiagnosisUserDetailResponseBodyUserDetailOnlinePeriods) Validate() error {
	return dara.Validate(s)
}
