// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHoneypotEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHoneypotEvents(v []*ListHoneypotEventsResponseBodyHoneypotEvents) *ListHoneypotEventsResponseBody
	GetHoneypotEvents() []*ListHoneypotEventsResponseBodyHoneypotEvents
	SetPageInfo(v *ListHoneypotEventsResponseBodyPageInfo) *ListHoneypotEventsResponseBody
	GetPageInfo() *ListHoneypotEventsResponseBodyPageInfo
	SetRequestId(v string) *ListHoneypotEventsResponseBody
	GetRequestId() *string
}

type ListHoneypotEventsResponseBody struct {
	// The list of honeypot attack events.
	HoneypotEvents []*ListHoneypotEventsResponseBodyHoneypotEvents `json:"HoneypotEvents,omitempty" xml:"HoneypotEvents,omitempty" type:"Repeated"`
	// The paging information.
	PageInfo *ListHoneypotEventsResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// FBD28009-6096-5E90-BFE6-62CCD67*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListHoneypotEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventsResponseBody) GetHoneypotEvents() []*ListHoneypotEventsResponseBodyHoneypotEvents {
	return s.HoneypotEvents
}

func (s *ListHoneypotEventsResponseBody) GetPageInfo() *ListHoneypotEventsResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListHoneypotEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHoneypotEventsResponseBody) SetHoneypotEvents(v []*ListHoneypotEventsResponseBodyHoneypotEvents) *ListHoneypotEventsResponseBody {
	s.HoneypotEvents = v
	return s
}

func (s *ListHoneypotEventsResponseBody) SetPageInfo(v *ListHoneypotEventsResponseBodyPageInfo) *ListHoneypotEventsResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListHoneypotEventsResponseBody) SetRequestId(v string) *ListHoneypotEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHoneypotEventsResponseBody) Validate() error {
	if s.HoneypotEvents != nil {
		for _, item := range s.HoneypotEvents {
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

type ListHoneypotEventsResponseBodyHoneypotEvents struct {
	// The probe ID.
	//
	// example:
	//
	// 27d44bd5815d401992ea672874d9****
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// The name of the probe.
	//
	// example:
	//
	// 1193474_test_****
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// The ID of the alert event.
	//
	// example:
	//
	// 1900752
	AlarmEventId *int64 `json:"AlarmEventId,omitempty" xml:"AlarmEventId,omitempty"`
	// The destination IP address of the attack.
	//
	// example:
	//
	// 112.126.205.***
	DstIp *string `json:"DstIp,omitempty" xml:"DstIp,omitempty"`
	// The timestamp of the first occurrence.
	//
	// example:
	//
	// 1692670297
	FirstTime *int64 `json:"FirstTime,omitempty" xml:"FirstTime,omitempty"`
	// The name of the honeypot.
	//
	// example:
	//
	// 主机 tcp监听****
	HoneypotName *string `json:"HoneypotName,omitempty" xml:"HoneypotName,omitempty"`
	// The timestamp of the most recent occurrence.
	//
	// example:
	//
	// 1676558664
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The region.
	//
	// example:
	//
	// 北京市
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The list of extended field key-value pairs.
	MergeFieldList []*ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList `json:"MergeFieldList,omitempty" xml:"MergeFieldList,omitempty" type:"Repeated"`
	// The protocol. Valid values:
	//
	// - **tcp**.
	//
	// - **udp**.
	//
	// example:
	//
	// tcp
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The risk level. Valid values:
	//
	// - **2**: Low risk.
	//
	// - **3**: Medium risk.
	//
	// - **4**: High risk.
	//
	// example:
	//
	// 4
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
	// The unique ID of the attack event.
	//
	// example:
	//
	// 70427821
	SecurityEventId *int64 `json:"SecurityEventId,omitempty" xml:"SecurityEventId,omitempty"`
	// The source IP address of the attack.
	//
	// example:
	//
	// 192.168.62.***
	SrcIp *string `json:"SrcIp,omitempty" xml:"SrcIp,omitempty"`
}

func (s ListHoneypotEventsResponseBodyHoneypotEvents) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventsResponseBodyHoneypotEvents) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetAgentId() *string {
	return s.AgentId
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetAgentName() *string {
	return s.AgentName
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetAlarmEventId() *int64 {
	return s.AlarmEventId
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetDstIp() *string {
	return s.DstIp
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetFirstTime() *int64 {
	return s.FirstTime
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetHoneypotName() *string {
	return s.HoneypotName
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetLastTime() *int64 {
	return s.LastTime
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetLocation() *string {
	return s.Location
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetMergeFieldList() []*ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList {
	return s.MergeFieldList
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetProtocol() *string {
	return s.Protocol
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetSecurityEventId() *int64 {
	return s.SecurityEventId
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) GetSrcIp() *string {
	return s.SrcIp
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetAgentId(v string) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.AgentId = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetAgentName(v string) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.AgentName = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetAlarmEventId(v int64) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.AlarmEventId = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetDstIp(v string) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.DstIp = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetFirstTime(v int64) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.FirstTime = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetHoneypotName(v string) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.HoneypotName = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetLastTime(v int64) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.LastTime = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetLocation(v string) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.Location = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetMergeFieldList(v []*ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.MergeFieldList = v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetProtocol(v string) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.Protocol = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetRiskLevel(v string) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.RiskLevel = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetSecurityEventId(v int64) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.SecurityEventId = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) SetSrcIp(v string) *ListHoneypotEventsResponseBodyHoneypotEvents {
	s.SrcIp = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEvents) Validate() error {
	if s.MergeFieldList != nil {
		for _, item := range s.MergeFieldList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList struct {
	// The supplementary information of the field.
	//
	// example:
	//
	// data
	FieldExtInfo *string `json:"FieldExtInfo,omitempty" xml:"FieldExtInfo,omitempty"`
	// The key of the field.
	//
	// example:
	//
	// type
	FieldKey *string `json:"FieldKey,omitempty" xml:"FieldKey,omitempty"`
	// The field type.
	//
	// example:
	//
	// level2_item1
	FieldType *string `json:"FieldType,omitempty" xml:"FieldType,omitempty"`
	// The value of the field key.
	//
	// example:
	//
	// web_access
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
}

func (s ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) GetFieldExtInfo() *string {
	return s.FieldExtInfo
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) GetFieldKey() *string {
	return s.FieldKey
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) GetFieldType() *string {
	return s.FieldType
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) GetFieldValue() *string {
	return s.FieldValue
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) SetFieldExtInfo(v string) *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList {
	s.FieldExtInfo = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) SetFieldKey(v string) *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList {
	s.FieldKey = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) SetFieldType(v string) *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList {
	s.FieldType = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) SetFieldValue(v string) *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList {
	s.FieldValue = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyHoneypotEventsMergeFieldList) Validate() error {
	return dara.Validate(s)
}

type ListHoneypotEventsResponseBodyPageInfo struct {
	// The number of entries on the current page in a paging query.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The key of the last entry.
	//
	// example:
	//
	// CAESGgoSChAKDGNvbXBsZXRlVGltZRABCgQiAggAGAAiQAoJAGYXFWIAAAAACjMDLgAAADFTNzMyZDMwMzAzMDM1Mzc3Njc4MzA2ODY5NmI2YTY*********
	LastRowKey *string `json:"LastRowKey,omitempty" xml:"LastRowKey,omitempty"`
	// The NextToken value returned when the NextToken-based pagination method is used.
	//
	// example:
	//
	// B604532DEF982B875E8360A6EFA3B***
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The maximum number of entries per page in a paging query.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListHoneypotEventsResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListHoneypotEventsResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListHoneypotEventsResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListHoneypotEventsResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListHoneypotEventsResponseBodyPageInfo) GetLastRowKey() *string {
	return s.LastRowKey
}

func (s *ListHoneypotEventsResponseBodyPageInfo) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHoneypotEventsResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHoneypotEventsResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListHoneypotEventsResponseBodyPageInfo) SetCount(v int32) *ListHoneypotEventsResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyPageInfo) SetCurrentPage(v int32) *ListHoneypotEventsResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyPageInfo) SetLastRowKey(v string) *ListHoneypotEventsResponseBodyPageInfo {
	s.LastRowKey = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyPageInfo) SetNextToken(v string) *ListHoneypotEventsResponseBodyPageInfo {
	s.NextToken = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyPageInfo) SetPageSize(v int32) *ListHoneypotEventsResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyPageInfo) SetTotalCount(v int32) *ListHoneypotEventsResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListHoneypotEventsResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
