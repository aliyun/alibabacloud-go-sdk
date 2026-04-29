// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUnknownThreatDetectMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListUnknownThreatDetectMachineResponseBodyData) *ListUnknownThreatDetectMachineResponseBody
	GetData() []*ListUnknownThreatDetectMachineResponseBodyData
	SetPageInfo(v *ListUnknownThreatDetectMachineResponseBodyPageInfo) *ListUnknownThreatDetectMachineResponseBody
	GetPageInfo() *ListUnknownThreatDetectMachineResponseBodyPageInfo
	SetRequestId(v string) *ListUnknownThreatDetectMachineResponseBody
	GetRequestId() *string
}

type ListUnknownThreatDetectMachineResponseBody struct {
	Data     []*ListUnknownThreatDetectMachineResponseBodyData   `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageInfo *ListUnknownThreatDetectMachineResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A4EB8B1C-1DEC-5E18-BCD0-XXXXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUnknownThreatDetectMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectMachineResponseBody) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectMachineResponseBody) GetData() []*ListUnknownThreatDetectMachineResponseBodyData {
	return s.Data
}

func (s *ListUnknownThreatDetectMachineResponseBody) GetPageInfo() *ListUnknownThreatDetectMachineResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListUnknownThreatDetectMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUnknownThreatDetectMachineResponseBody) SetData(v []*ListUnknownThreatDetectMachineResponseBodyData) *ListUnknownThreatDetectMachineResponseBody {
	s.Data = v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBody) SetPageInfo(v *ListUnknownThreatDetectMachineResponseBodyPageInfo) *ListUnknownThreatDetectMachineResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBody) SetRequestId(v string) *ListUnknownThreatDetectMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBody) Validate() error {
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

type ListUnknownThreatDetectMachineResponseBodyData struct {
	// example:
	//
	// 12
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// example:
	//
	// 172.16.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// example:
	//
	// 10.42.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// example:
	//
	// 1
	ProcessCount *int32 `json:"ProcessCount,omitempty" xml:"ProcessCount,omitempty"`
	// example:
	//
	// studying
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// hash
	StudyMode *string `json:"StudyMode,omitempty" xml:"StudyMode,omitempty"`
	// example:
	//
	// 1766988192
	StudyStartTime *int64 `json:"StudyStartTime,omitempty" xml:"StudyStartTime,omitempty"`
	// example:
	//
	// 6690a46c-0edb-4663-a641-3629d1a9****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ListUnknownThreatDetectMachineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectMachineResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) GetProcessCount() *int32 {
	return s.ProcessCount
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) GetStudyMode() *string {
	return s.StudyMode
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) GetStudyStartTime() *int64 {
	return s.StudyStartTime
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) SetInstanceName(v string) *ListUnknownThreatDetectMachineResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) SetInternetIp(v string) *ListUnknownThreatDetectMachineResponseBodyData {
	s.InternetIp = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) SetIntranetIp(v string) *ListUnknownThreatDetectMachineResponseBodyData {
	s.IntranetIp = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) SetProcessCount(v int32) *ListUnknownThreatDetectMachineResponseBodyData {
	s.ProcessCount = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) SetStatus(v string) *ListUnknownThreatDetectMachineResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) SetStudyMode(v string) *ListUnknownThreatDetectMachineResponseBodyData {
	s.StudyMode = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) SetStudyStartTime(v int64) *ListUnknownThreatDetectMachineResponseBodyData {
	s.StudyStartTime = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) SetUuid(v string) *ListUnknownThreatDetectMachineResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListUnknownThreatDetectMachineResponseBodyPageInfo struct {
	// example:
	//
	// 10
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
	// 149
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUnknownThreatDetectMachineResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListUnknownThreatDetectMachineResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) SetCount(v int32) *ListUnknownThreatDetectMachineResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) SetCurrentPage(v int32) *ListUnknownThreatDetectMachineResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) SetPageSize(v int32) *ListUnknownThreatDetectMachineResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) SetTotalCount(v int32) *ListUnknownThreatDetectMachineResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListUnknownThreatDetectMachineResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
