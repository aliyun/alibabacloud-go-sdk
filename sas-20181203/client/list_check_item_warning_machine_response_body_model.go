// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemWarningMachineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetList(v []*ListCheckItemWarningMachineResponseBodyList) *ListCheckItemWarningMachineResponseBody
	GetList() []*ListCheckItemWarningMachineResponseBodyList
	SetPageInfo(v *ListCheckItemWarningMachineResponseBodyPageInfo) *ListCheckItemWarningMachineResponseBody
	GetPageInfo() *ListCheckItemWarningMachineResponseBodyPageInfo
	SetRequestId(v string) *ListCheckItemWarningMachineResponseBody
	GetRequestId() *string
}

type ListCheckItemWarningMachineResponseBody struct {
	// The servers on which the alerts are generated.
	List []*ListCheckItemWarningMachineResponseBodyList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListCheckItemWarningMachineResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 22B5615F-700E-575A-A6D5-DC8D7741****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCheckItemWarningMachineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningMachineResponseBody) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningMachineResponseBody) GetList() []*ListCheckItemWarningMachineResponseBodyList {
	return s.List
}

func (s *ListCheckItemWarningMachineResponseBody) GetPageInfo() *ListCheckItemWarningMachineResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCheckItemWarningMachineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCheckItemWarningMachineResponseBody) SetList(v []*ListCheckItemWarningMachineResponseBodyList) *ListCheckItemWarningMachineResponseBody {
	s.List = v
	return s
}

func (s *ListCheckItemWarningMachineResponseBody) SetPageInfo(v *ListCheckItemWarningMachineResponseBodyPageInfo) *ListCheckItemWarningMachineResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCheckItemWarningMachineResponseBody) SetRequestId(v string) *ListCheckItemWarningMachineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemWarningMachineResponseBodyList struct {
	// The edition of Security Center that is authorized to protect the asset. Valid values:
	//
	// 	- **1**: Basic edition
	//
	// 	- **6**: Anti-virus edition
	//
	// 	- **5**: Advanced edition
	//
	// 	- **3**: Enterprise edition
	//
	// 	- **7**: Ultimate edition
	//
	// 	- **10**: Value-added Plan edition
	//
	// example:
	//
	// 3
	AuthVersion *int32 `json:"AuthVersion,omitempty" xml:"AuthVersion,omitempty"`
	// Indicates whether Security Center is authorized to protect the asset. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Bind *bool `json:"Bind,omitempty" xml:"Bind,omitempty"`
	// The ID of the container.
	//
	// example:
	//
	// 48a6d9a92435a13ad573372c3f3c63b7e04d106458141df9f9215570********
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The name of the container.
	//
	// example:
	//
	// step-build-ui-build
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The details of the baselines for which the risk item can be fixed.
	FixList []*ListCheckItemWarningMachineResponseBodyListFixList `json:"FixList,omitempty" xml:"FixList,omitempty" type:"Repeated"`
	// Whether the repair is supported. Valid values:
	//
	// 	- **0**: Supported
	//
	// 	- **1**: Not Supported
	//
	// example:
	//
	// 1
	FixStatus *int32 `json:"FixStatus,omitempty" xml:"FixStatus,omitempty"`
	// The instance ID of the server.
	//
	// example:
	//
	// i-bp1a69mvjujbakxu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// sql-test-0****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The public IP address of the affected asset.
	//
	// example:
	//
	// 8.210.XX.XX
	InternetIp *string `json:"InternetIp,omitempty" xml:"InternetIp,omitempty"`
	// The private IP address of the affected asset.
	//
	// example:
	//
	// 172.25.XX.XX
	IntranetIp *string `json:"IntranetIp,omitempty" xml:"IntranetIp,omitempty"`
	// The timestamp of the latest processing of the check item risk of the machine. Unit: milliseconds.
	//
	// example:
	//
	// 1694692471000
	LastHandleTime *int64 `json:"LastHandleTime,omitempty" xml:"LastHandleTime,omitempty"`
	// The timestamp generated when the last scan was performed. Unit: milliseconds.
	//
	// example:
	//
	// 1694692471000
	LastScanTime *int64 `json:"LastScanTime,omitempty" xml:"LastScanTime,omitempty"`
	// Indicates whether a port on the server is accessible over the Internet. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	PortOpen *bool `json:"PortOpen,omitempty" xml:"PortOpen,omitempty"`
	// The prompt for the risk item.
	//
	// example:
	//
	// There is a weak password (username/password): root/he*****34
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// The region ID of the asset.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- 1: failed
	//
	// 	- 2: verifying
	//
	// 	- 3: passed
	//
	// 	- 6: ignored
	//
	// 	- 7: fixing
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the asset that is scanned.
	//
	// example:
	//
	// 30****
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// The name of the asset on which the malicious image sample is detected.
	//
	// example:
	//
	// jenkins****
	TargetName *string `json:"TargetName,omitempty" xml:"TargetName,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- **ECS_SNAPSHOT**
	//
	// 	- **ECS_IMAGE**
	//
	// example:
	//
	// ECS_IMAGE
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The UUID of the server.
	//
	// example:
	//
	// 49e25e0f-bb51-4a5a-a1b3-13a4ddaa****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// Deprecated
	//
	// The information about the baselines on which the risk item is detected.
	WarningRiskList []*ListCheckItemWarningMachineResponseBodyListWarningRiskList `json:"WarningRiskList,omitempty" xml:"WarningRiskList,omitempty" type:"Repeated"`
}

func (s ListCheckItemWarningMachineResponseBodyList) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningMachineResponseBodyList) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetAuthVersion() *int32 {
	return s.AuthVersion
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetBind() *bool {
	return s.Bind
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetContainerId() *string {
	return s.ContainerId
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetContainerName() *string {
	return s.ContainerName
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetFixList() []*ListCheckItemWarningMachineResponseBodyListFixList {
	return s.FixList
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetFixStatus() *int32 {
	return s.FixStatus
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetInternetIp() *string {
	return s.InternetIp
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetIntranetIp() *string {
	return s.IntranetIp
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetLastHandleTime() *int64 {
	return s.LastHandleTime
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetLastScanTime() *int64 {
	return s.LastScanTime
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetPortOpen() *bool {
	return s.PortOpen
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetPrompt() *string {
	return s.Prompt
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetStatus() *int32 {
	return s.Status
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetTargetId() *string {
	return s.TargetId
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetTargetName() *string {
	return s.TargetName
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetTargetType() *string {
	return s.TargetType
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetUuid() *string {
	return s.Uuid
}

func (s *ListCheckItemWarningMachineResponseBodyList) GetWarningRiskList() []*ListCheckItemWarningMachineResponseBodyListWarningRiskList {
	return s.WarningRiskList
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetAuthVersion(v int32) *ListCheckItemWarningMachineResponseBodyList {
	s.AuthVersion = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetBind(v bool) *ListCheckItemWarningMachineResponseBodyList {
	s.Bind = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetContainerId(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.ContainerId = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetContainerName(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.ContainerName = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetFixList(v []*ListCheckItemWarningMachineResponseBodyListFixList) *ListCheckItemWarningMachineResponseBodyList {
	s.FixList = v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetFixStatus(v int32) *ListCheckItemWarningMachineResponseBodyList {
	s.FixStatus = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetInstanceId(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.InstanceId = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetInstanceName(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.InstanceName = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetInternetIp(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.InternetIp = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetIntranetIp(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.IntranetIp = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetLastHandleTime(v int64) *ListCheckItemWarningMachineResponseBodyList {
	s.LastHandleTime = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetLastScanTime(v int64) *ListCheckItemWarningMachineResponseBodyList {
	s.LastScanTime = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetPortOpen(v bool) *ListCheckItemWarningMachineResponseBodyList {
	s.PortOpen = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetPrompt(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.Prompt = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetRegionId(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.RegionId = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetStatus(v int32) *ListCheckItemWarningMachineResponseBodyList {
	s.Status = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetTargetId(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.TargetId = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetTargetName(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.TargetName = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetTargetType(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.TargetType = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetUuid(v string) *ListCheckItemWarningMachineResponseBodyList {
	s.Uuid = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) SetWarningRiskList(v []*ListCheckItemWarningMachineResponseBodyListWarningRiskList) *ListCheckItemWarningMachineResponseBodyList {
	s.WarningRiskList = v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyList) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemWarningMachineResponseBodyListFixList struct {
	// The ID of the baseline.
	//
	// example:
	//
	// 72
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The name of the baseline.
	//
	// example:
	//
	// Alibaba Cloud Linux/Aliyun Linux 2 Baseline for China classified protection of cybersecurity-Level II
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
}

func (s ListCheckItemWarningMachineResponseBodyListFixList) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningMachineResponseBodyListFixList) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningMachineResponseBodyListFixList) GetRiskId() *int64 {
	return s.RiskId
}

func (s *ListCheckItemWarningMachineResponseBodyListFixList) GetRiskName() *string {
	return s.RiskName
}

func (s *ListCheckItemWarningMachineResponseBodyListFixList) SetRiskId(v int64) *ListCheckItemWarningMachineResponseBodyListFixList {
	s.RiskId = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyListFixList) SetRiskName(v string) *ListCheckItemWarningMachineResponseBodyListFixList {
	s.RiskName = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyListFixList) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemWarningMachineResponseBodyListWarningRiskList struct {
	// Deprecated
	//
	// The ID of the baseline.
	//
	// example:
	//
	// 72
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// Deprecated
	//
	// The name of the baseline.
	//
	// example:
	//
	// Alibaba Cloud Linux/Aliyun Linux 2 Baseline for China classified protection of cybersecurity-Level II
	RiskName *string `json:"RiskName,omitempty" xml:"RiskName,omitempty"`
}

func (s ListCheckItemWarningMachineResponseBodyListWarningRiskList) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningMachineResponseBodyListWarningRiskList) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningMachineResponseBodyListWarningRiskList) GetRiskId() *int64 {
	return s.RiskId
}

func (s *ListCheckItemWarningMachineResponseBodyListWarningRiskList) GetRiskName() *string {
	return s.RiskName
}

func (s *ListCheckItemWarningMachineResponseBodyListWarningRiskList) SetRiskId(v int64) *ListCheckItemWarningMachineResponseBodyListWarningRiskList {
	s.RiskId = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyListWarningRiskList) SetRiskName(v string) *ListCheckItemWarningMachineResponseBodyListWarningRiskList {
	s.RiskName = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyListWarningRiskList) Validate() error {
	return dara.Validate(s)
}

type ListCheckItemWarningMachineResponseBodyPageInfo struct {
	// The number of affected assets returned on the current page.
	//
	// example:
	//
	// 4
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of affected assets.
	//
	// example:
	//
	// 107
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCheckItemWarningMachineResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningMachineResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) SetCount(v int32) *ListCheckItemWarningMachineResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) SetCurrentPage(v int32) *ListCheckItemWarningMachineResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) SetPageSize(v int32) *ListCheckItemWarningMachineResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) SetTotalCount(v int32) *ListCheckItemWarningMachineResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCheckItemWarningMachineResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
