// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainSecureAlarmListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlarmList(v []*DescribeDomainSecureAlarmListResponseBodyAlarmList) *DescribeDomainSecureAlarmListResponseBody
	GetAlarmList() []*DescribeDomainSecureAlarmListResponseBodyAlarmList
	SetRequestId(v string) *DescribeDomainSecureAlarmListResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDomainSecureAlarmListResponseBody
	GetTotalCount() *int32
}

type DescribeDomainSecureAlarmListResponseBody struct {
	// The security alerts in your website assets.
	AlarmList []*DescribeDomainSecureAlarmListResponseBodyAlarmList `json:"AlarmList,omitempty" xml:"AlarmList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D03DD0FD-6041-5107-AC00-383E28F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 42
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDomainSecureAlarmListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureAlarmListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureAlarmListResponseBody) GetAlarmList() []*DescribeDomainSecureAlarmListResponseBodyAlarmList {
	return s.AlarmList
}

func (s *DescribeDomainSecureAlarmListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainSecureAlarmListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDomainSecureAlarmListResponseBody) SetAlarmList(v []*DescribeDomainSecureAlarmListResponseBodyAlarmList) *DescribeDomainSecureAlarmListResponseBody {
	s.AlarmList = v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBody) SetRequestId(v string) *DescribeDomainSecureAlarmListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBody) SetTotalCount(v int32) *DescribeDomainSecureAlarmListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBody) Validate() error {
	if s.AlarmList != nil {
		for _, item := range s.AlarmList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDomainSecureAlarmListResponseBodyAlarmList struct {
	// The name of the alert event.
	//
	// example:
	//
	// Trojan
	AlarmEventName *string `json:"AlarmEventName,omitempty" xml:"AlarmEventName,omitempty"`
	// The original parent name of the alert event.
	//
	// example:
	//
	// login_common_location
	AlarmEventNameOriginal *string `json:"AlarmEventNameOriginal,omitempty" xml:"AlarmEventNameOriginal,omitempty"`
	// The type of the alert event.
	//
	// example:
	//
	// Malicious Software
	AlarmEventType *string `json:"AlarmEventType,omitempty" xml:"AlarmEventType,omitempty"`
	// The unique ID of the alert event.
	//
	// example:
	//
	// 8df914418f4211fbf756efe7a6f4****
	AlarmUniqueInfo *string `json:"AlarmUniqueInfo,omitempty" xml:"AlarmUniqueInfo,omitempty"`
	// Indicates whether automatic defense is enabled.
	//
	// example:
	//
	// true
	AutoBreaking *bool `json:"AutoBreaking,omitempty" xml:"AutoBreaking,omitempty"`
	// Indicates whether the alert event can be handled online, such as quarantining the source file of the malicious process, adding the alert event to the whitelist, and ignoring the alert event. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CanBeDealOnLine *bool `json:"CanBeDealOnLine,omitempty" xml:"CanBeDealOnLine,omitempty"`
	// Indicates whether you can cancel marking the alert event as a false positive. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// false
	CanCancelFault *bool `json:"CanCancelFault,omitempty" xml:"CanCancelFault,omitempty"`
	// Indicates whether the safeguard mode for major activities is supported.
	//
	// example:
	//
	// true
	ContainHwMode *bool `json:"ContainHwMode,omitempty" xml:"ContainHwMode,omitempty"`
	// The data source of the alert event.
	//
	// example:
	//
	// aegis_****
	DataSource *string `json:"DataSource,omitempty" xml:"DataSource,omitempty"`
	// Indicates whether the alert event is handled. Valid values:
	//
	// 	- **N**: unhandled
	//
	// 	- **Y**: handled
	//
	// example:
	//
	// y
	Dealed *bool `json:"Dealed,omitempty" xml:"Dealed,omitempty"`
	// The description of the alert event.
	//
	// example:
	//
	// The detection model finds that there is a Trojan horse program on your server. The Trojan horse program is a program specially used to invade the user\\"s host. Generally, it will download and release another malicious program after being implanted into the system through disguise.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The timestamp generated when the alert event was last detected. Unit: milliseconds.
	//
	// example:
	//
	// 1543740301000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The time of the last modification.
	//
	// example:
	//
	// 1656901794000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the alert event has tracing information. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	HasTraceInfo *bool `json:"HasTraceInfo,omitempty" xml:"HasTraceInfo,omitempty"`
	// The instance ID of the affected asset.
	//
	// example:
	//
	// i-e****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name of the affected asset.
	//
	// example:
	//
	// TestInstance
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the server.
	//
	// example:
	//
	// 95.214.*.*
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the affected instance.
	//
	// example:
	//
	// 192.168.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The risk level of the alert event. Valid values:
	//
	// 	- **serious**
	//
	// 	- **suspicious**
	//
	// 	- **remind**
	//
	// example:
	//
	// serious
	Level *string `json:"Level,omitempty" xml:"Level,omitempty"`
	// The handling result code of the alert event.
	//
	// example:
	//
	// kill_and_quara.Success
	OperateErrorCode *string `json:"OperateErrorCode,omitempty" xml:"OperateErrorCode,omitempty"`
	// The timestamp generated when the alert event was handled. Unit: milliseconds.
	//
	// example:
	//
	// 1631699497000
	OperateTime *int64 `json:"OperateTime,omitempty" xml:"OperateTime,omitempty"`
	// The edition of Security Center in which the alert event can be detected. Valid values:
	//
	// 	- **0**: Basic edition.
	//
	// 	- **1**: Advanced edition.
	//
	// 	- **2**: Enterprise edition.
	//
	// example:
	//
	// 1
	SaleVersion *string `json:"SaleVersion,omitempty" xml:"SaleVersion,omitempty"`
	// The ID of the associated alert event.
	//
	// example:
	//
	// 270789
	SecurityEventIds *string `json:"SecurityEventIds,omitempty" xml:"SecurityEventIds,omitempty"`
	// The solution to the alert event.
	//
	// example:
	//
	// A malicious program implanted by hacker after intrusion will occupy your bandwidth and attack other servers, and may affect you own service. The malicious process may also have self-deleting behavior or disguise as a system service to evade detection.
	Solution *string `json:"Solution,omitempty" xml:"Solution,omitempty"`
	// The stage at which the attack or intrusion is detected.
	//
	// example:
	//
	// [\\"authority_maintenance\\"]
	Stages *string `json:"Stages,omitempty" xml:"Stages,omitempty"`
	// The timestamp generated when the alert event was first detected. Unit: milliseconds.
	//
	// example:
	//
	// 1543740301000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The total number of security alerts in your website assets.
	//
	// example:
	//
	// 1
	SuspiciousEventCount *int32 `json:"SuspiciousEventCount,omitempty" xml:"SuspiciousEventCount,omitempty"`
	// The unique ID of the associated instance.
	//
	// example:
	//
	// 47900178-885d-4fa4-9d77-****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s DescribeDomainSecureAlarmListResponseBodyAlarmList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainSecureAlarmListResponseBodyAlarmList) GoString() string {
	return s.String()
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetAlarmEventName() *string {
	return s.AlarmEventName
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetAlarmEventNameOriginal() *string {
	return s.AlarmEventNameOriginal
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetAlarmEventType() *string {
	return s.AlarmEventType
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetAlarmUniqueInfo() *string {
	return s.AlarmUniqueInfo
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetAutoBreaking() *bool {
	return s.AutoBreaking
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetCanBeDealOnLine() *bool {
	return s.CanBeDealOnLine
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetCanCancelFault() *bool {
	return s.CanCancelFault
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetContainHwMode() *bool {
	return s.ContainHwMode
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetDataSource() *string {
	return s.DataSource
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetDealed() *bool {
	return s.Dealed
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetDescription() *string {
	return s.Description
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetHasTraceInfo() *bool {
	return s.HasTraceInfo
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetLevel() *string {
	return s.Level
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetOperateErrorCode() *string {
	return s.OperateErrorCode
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetOperateTime() *int64 {
	return s.OperateTime
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetSaleVersion() *string {
	return s.SaleVersion
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetSecurityEventIds() *string {
	return s.SecurityEventIds
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetSolution() *string {
	return s.Solution
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetStages() *string {
	return s.Stages
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetSuspiciousEventCount() *int32 {
	return s.SuspiciousEventCount
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) GetUuid() *string {
	return s.Uuid
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetAlarmEventName(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.AlarmEventName = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetAlarmEventNameOriginal(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.AlarmEventNameOriginal = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetAlarmEventType(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.AlarmEventType = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetAlarmUniqueInfo(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.AlarmUniqueInfo = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetAutoBreaking(v bool) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.AutoBreaking = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetCanBeDealOnLine(v bool) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.CanBeDealOnLine = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetCanCancelFault(v bool) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.CanCancelFault = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetContainHwMode(v bool) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.ContainHwMode = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetDataSource(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.DataSource = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetDealed(v bool) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.Dealed = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetDescription(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.Description = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetEndTime(v int64) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.EndTime = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetGmtModified(v int64) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.GmtModified = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetHasTraceInfo(v bool) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.HasTraceInfo = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetInstanceId(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.InstanceId = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetInstanceName(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.InstanceName = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetInternetIp(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.InternetIp = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetIntranetIp(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.IntranetIp = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetLevel(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.Level = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetOperateErrorCode(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.OperateErrorCode = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetOperateTime(v int64) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.OperateTime = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetSaleVersion(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.SaleVersion = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetSecurityEventIds(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.SecurityEventIds = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetSolution(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.Solution = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetStages(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.Stages = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetStartTime(v int64) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.StartTime = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetSuspiciousEventCount(v int32) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.SuspiciousEventCount = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) SetUuid(v string) *DescribeDomainSecureAlarmListResponseBodyAlarmList {
	s.Uuid = &v
	return s
}

func (s *DescribeDomainSecureAlarmListResponseBodyAlarmList) Validate() error {
	return dara.Validate(s)
}
