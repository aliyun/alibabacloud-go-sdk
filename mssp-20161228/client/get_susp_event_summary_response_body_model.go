// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspEventSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSuspEventSummaryResponseBody
	GetCode() *string
	SetData(v *GetSuspEventSummaryResponseBodyData) *GetSuspEventSummaryResponseBody
	GetData() *GetSuspEventSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetSuspEventSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSuspEventSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSuspEventSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSuspEventSummaryResponseBody
	GetSuccess() *bool
}

type GetSuspEventSummaryResponseBody struct {
	// API response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the interface.
	Data *GetSuspEventSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Prompt message for the returned result.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9B2DAE9B-B901-5818-AFEF-E5637D938280
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the call was successful.
	//
	// - true: Call succeeded.
	//
	// - false: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSuspEventSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSuspEventSummaryResponseBody) GetData() *GetSuspEventSummaryResponseBodyData {
	return s.Data
}

func (s *GetSuspEventSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSuspEventSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSuspEventSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSuspEventSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSuspEventSummaryResponseBody) SetCode(v string) *GetSuspEventSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetData(v *GetSuspEventSummaryResponseBodyData) *GetSuspEventSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetHttpStatusCode(v int32) *GetSuspEventSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetMessage(v string) *GetSuspEventSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetRequestId(v string) *GetSuspEventSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) SetSuccess(v bool) *GetSuspEventSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetSuspEventSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSuspEventSummaryResponseBodyData struct {
	// Network attack trend.
	NetworkAttackTrendDTO *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO `json:"NetworkAttackTrendDTO,omitempty" xml:"NetworkAttackTrendDTO,omitempty" type:"Struct"`
	// Overview of alert handling.
	SuspEventDealSummaryDTO *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO `json:"SuspEventDealSummaryDTO,omitempty" xml:"SuspEventDealSummaryDTO,omitempty" type:"Struct"`
	// Top 10 alerts before handling.
	SuspEventTopDTO *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO `json:"SuspEventTopDTO,omitempty" xml:"SuspEventTopDTO,omitempty" type:"Struct"`
	// Trend of alert responses.
	SuspEventTrendDTO *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO `json:"SuspEventTrendDTO,omitempty" xml:"SuspEventTrendDTO,omitempty" type:"Struct"`
}

func (s GetSuspEventSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyData) GetNetworkAttackTrendDTO() *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO {
	return s.NetworkAttackTrendDTO
}

func (s *GetSuspEventSummaryResponseBodyData) GetSuspEventDealSummaryDTO() *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	return s.SuspEventDealSummaryDTO
}

func (s *GetSuspEventSummaryResponseBodyData) GetSuspEventTopDTO() *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO {
	return s.SuspEventTopDTO
}

func (s *GetSuspEventSummaryResponseBodyData) GetSuspEventTrendDTO() *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO {
	return s.SuspEventTrendDTO
}

func (s *GetSuspEventSummaryResponseBodyData) SetNetworkAttackTrendDTO(v *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) *GetSuspEventSummaryResponseBodyData {
	s.NetworkAttackTrendDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventDealSummaryDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventDealSummaryDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventTopDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventTopDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) SetSuspEventTrendDTO(v *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) *GetSuspEventSummaryResponseBodyData {
	s.SuspEventTrendDTO = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyData) Validate() error {
	if s.NetworkAttackTrendDTO != nil {
		if err := s.NetworkAttackTrendDTO.Validate(); err != nil {
			return err
		}
	}
	if s.SuspEventDealSummaryDTO != nil {
		if err := s.SuspEventDealSummaryDTO.Validate(); err != nil {
			return err
		}
	}
	if s.SuspEventTopDTO != nil {
		if err := s.SuspEventTopDTO.Validate(); err != nil {
			return err
		}
	}
	if s.SuspEventTrendDTO != nil {
		if err := s.SuspEventTrendDTO.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO struct {
	// Collection of trend nodes for each attack item.
	TrendList []*GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) GetTrendList() []*GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	return s.TrendList
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) SetTrendList(v []*GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO {
	s.TrendList = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTO) Validate() error {
	if s.TrendList != nil {
		for _, item := range s.TrendList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList struct {
	// Date.
	//
	// example:
	//
	// 202409或20240901
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// DDoS count.
	//
	// example:
	//
	// 10
	DdosCount *int64 `json:"DdosCount,omitempty" xml:"DdosCount,omitempty"`
	// EIP count.
	//
	// example:
	//
	// 10
	EipCount *int64 `json:"EipCount,omitempty" xml:"EipCount,omitempty"`
	// WAF count.
	//
	// example:
	//
	// 10
	WafCount *int64 `json:"WafCount,omitempty" xml:"WafCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) GetDate() *string {
	return s.Date
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) GetDdosCount() *int64 {
	return s.DdosCount
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) GetEipCount() *int64 {
	return s.EipCount
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) GetWafCount() *int64 {
	return s.WafCount
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetDate(v string) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.Date = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetDdosCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.DdosCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetEipCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.EipCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) SetWafCount(v int64) *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList {
	s.WafCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataNetworkAttackTrendDTOTrendList) Validate() error {
	return dara.Validate(s)
}

type GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO struct {
	// Completed.
	//
	// example:
	//
	// 20
	CompletedCount *int64 `json:"CompletedCount,omitempty" xml:"CompletedCount,omitempty"`
	// In progress.
	//
	// example:
	//
	// 5
	HandingCount *int64 `json:"HandingCount,omitempty" xml:"HandingCount,omitempty"`
	// Alert handling rate.
	//
	// example:
	//
	// 90
	HandingRate *string `json:"HandingRate,omitempty" xml:"HandingRate,omitempty"`
	// Total number of alerts.
	//
	// example:
	//
	// 35
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Year-over-year comparison of alerts.
	//
	// example:
	//
	// 10
	TotalGrowthRate *string `json:"TotalGrowthRate,omitempty" xml:"TotalGrowthRate,omitempty"`
	// Number of unhandled alerts.
	//
	// example:
	//
	// 10
	WaitHandleCount *int64 `json:"WaitHandleCount,omitempty" xml:"WaitHandleCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GetCompletedCount() *int64 {
	return s.CompletedCount
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GetHandingCount() *int64 {
	return s.HandingCount
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GetHandingRate() *string {
	return s.HandingRate
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GetTotalGrowthRate() *string {
	return s.TotalGrowthRate
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) GetWaitHandleCount() *int64 {
	return s.WaitHandleCount
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetCompletedCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.CompletedCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetHandingCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.HandingCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetHandingRate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.HandingRate = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetTotalCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.TotalCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetTotalGrowthRate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.TotalGrowthRate = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) SetWaitHandleCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO {
	s.WaitHandleCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventDealSummaryDTO) Validate() error {
	return dara.Validate(s)
}

type GetSuspEventSummaryResponseBodyDataSuspEventTopDTO struct {
	// Top 10 before handling alarms
	SuspEventList []*GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList `json:"SuspEventList,omitempty" xml:"SuspEventList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) GetSuspEventList() []*GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList {
	return s.SuspEventList
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) SetSuspEventList(v []*GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO {
	s.SuspEventList = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTO) Validate() error {
	if s.SuspEventList != nil {
		for _, item := range s.SuspEventList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList struct {
	// Alert name.
	//
	// example:
	//
	// 主动外连风险 IP
	EventName *string `json:"EventName,omitempty" xml:"EventName,omitempty"`
	// Count.
	//
	// example:
	//
	// 7
	TaskCount *int64 `json:"TaskCount,omitempty" xml:"TaskCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) GetEventName() *string {
	return s.EventName
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) GetTaskCount() *int64 {
	return s.TaskCount
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) SetEventName(v string) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList {
	s.EventName = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) SetTaskCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList {
	s.TaskCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTopDTOSuspEventList) Validate() error {
	return dara.Validate(s)
}

type GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO struct {
	// Trend of alerts.
	TrendList []*GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList `json:"TrendList,omitempty" xml:"TrendList,omitempty" type:"Repeated"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) GetTrendList() []*GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	return s.TrendList
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) SetTrendList(v []*GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO {
	s.TrendList = v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTO) Validate() error {
	if s.TrendList != nil {
		for _, item := range s.TrendList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList struct {
	// Time point.
	//
	// example:
	//
	// 202405或者20240501
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// Number of handled alerts.
	//
	// example:
	//
	// 10
	DealCount *int64 `json:"DealCount,omitempty" xml:"DealCount,omitempty"`
	// Number of discovered alerts.
	//
	// example:
	//
	// 15
	FindCount *int64 `json:"FindCount,omitempty" xml:"FindCount,omitempty"`
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) String() string {
	return dara.Prettify(s)
}

func (s GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) GoString() string {
	return s.String()
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) GetDate() *string {
	return s.Date
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) GetDealCount() *int64 {
	return s.DealCount
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) GetFindCount() *int64 {
	return s.FindCount
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetDate(v string) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.Date = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetDealCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.DealCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) SetFindCount(v int64) *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList {
	s.FindCount = &v
	return s
}

func (s *GetSuspEventSummaryResponseBodyDataSuspEventTrendDTOTrendList) Validate() error {
	return dara.Validate(s)
}
