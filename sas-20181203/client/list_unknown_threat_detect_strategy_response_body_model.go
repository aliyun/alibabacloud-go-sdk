// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUnknownThreatDetectStrategyResponseBodyData) *ListUnknownThreatDetectStrategyResponseBody
	GetData() []*ListUnknownThreatDetectStrategyResponseBodyData
	SetPageInfo(v *ListUnknownThreatDetectStrategyResponseBodyPageInfo) *ListUnknownThreatDetectStrategyResponseBody
	GetPageInfo() *ListUnknownThreatDetectStrategyResponseBodyPageInfo
	SetRequestId(v string) *ListUnknownThreatDetectStrategyResponseBody
	GetRequestId() *string
}

type ListUnknownThreatDetectStrategyResponseBody struct {
	Data     []*ListUnknownThreatDetectStrategyResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageInfo *ListUnknownThreatDetectStrategyResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 898F7AA7-CECD-5EC7-AF4D-664C601B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUnknownThreatDetectStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectStrategyResponseBody) GetData() []*ListUnknownThreatDetectStrategyResponseBodyData {
	return s.Data
}

func (s *ListUnknownThreatDetectStrategyResponseBody) GetPageInfo() *ListUnknownThreatDetectStrategyResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListUnknownThreatDetectStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUnknownThreatDetectStrategyResponseBody) SetData(v []*ListUnknownThreatDetectStrategyResponseBodyData) *ListUnknownThreatDetectStrategyResponseBody {
	s.Data = v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBody) SetPageInfo(v *ListUnknownThreatDetectStrategyResponseBodyPageInfo) *ListUnknownThreatDetectStrategyResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBody) SetRequestId(v string) *ListUnknownThreatDetectStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUnknownThreatDetectStrategyResponseBodyData struct {
	// example:
	//
	// UNKNOWN_THREAT_DETECT_CONFIG_****
	AssetSelectionType *string `json:"AssetSelectionType,omitempty" xml:"AssetSelectionType,omitempty"`
	// example:
	//
	// 1
	DurationDaysAfterInit *int32 `json:"DurationDaysAfterInit,omitempty" xml:"DurationDaysAfterInit,omitempty"`
	// example:
	//
	// 1
	DurationDaysAfterStop *int32 `json:"DurationDaysAfterStop,omitempty" xml:"DurationDaysAfterStop,omitempty"`
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1
	MachineCount *int32 `json:"MachineCount,omitempty" xml:"MachineCount,omitempty"`
	// example:
	//
	// test****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// hash
	StudyMode *string `json:"StudyMode,omitempty" xml:"StudyMode,omitempty"`
}

func (s ListUnknownThreatDetectStrategyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectStrategyResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) GetAssetSelectionType() *string {
	return s.AssetSelectionType
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) GetDurationDaysAfterInit() *int32 {
	return s.DurationDaysAfterInit
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) GetDurationDaysAfterStop() *int32 {
	return s.DurationDaysAfterStop
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) GetMachineCount() *int32 {
	return s.MachineCount
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) GetStudyMode() *string {
	return s.StudyMode
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) SetAssetSelectionType(v string) *ListUnknownThreatDetectStrategyResponseBodyData {
	s.AssetSelectionType = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) SetDurationDaysAfterInit(v int32) *ListUnknownThreatDetectStrategyResponseBodyData {
	s.DurationDaysAfterInit = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) SetDurationDaysAfterStop(v int32) *ListUnknownThreatDetectStrategyResponseBodyData {
	s.DurationDaysAfterStop = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) SetId(v int64) *ListUnknownThreatDetectStrategyResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) SetMachineCount(v int32) *ListUnknownThreatDetectStrategyResponseBodyData {
	s.MachineCount = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) SetName(v string) *ListUnknownThreatDetectStrategyResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) SetStudyMode(v string) *ListUnknownThreatDetectStrategyResponseBodyData {
	s.StudyMode = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUnknownThreatDetectStrategyResponseBodyPageInfo struct {
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 69
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUnknownThreatDetectStrategyResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectStrategyResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) SetCount(v int32) *ListUnknownThreatDetectStrategyResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) SetCurrentPage(v int32) *ListUnknownThreatDetectStrategyResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) SetPageSize(v int32) *ListUnknownThreatDetectStrategyResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) SetTotalCount(v int32) *ListUnknownThreatDetectStrategyResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListUnknownThreatDetectStrategyResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
