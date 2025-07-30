// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDedicatedClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedClusterStatusList(v *ListDedicatedClusterResponseBodyDedicatedClusterStatusList) *ListDedicatedClusterResponseBody
	GetDedicatedClusterStatusList() *ListDedicatedClusterResponseBodyDedicatedClusterStatusList
	SetErrCode(v string) *ListDedicatedClusterResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ListDedicatedClusterResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *ListDedicatedClusterResponseBody
	GetHttpStatusCode() *string
	SetPageNumber(v int32) *ListDedicatedClusterResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListDedicatedClusterResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListDedicatedClusterResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ListDedicatedClusterResponseBody
	GetSuccess() *string
	SetTotalRecordCount(v int32) *ListDedicatedClusterResponseBody
	GetTotalRecordCount() *int32
}

type ListDedicatedClusterResponseBody struct {
	// The statuses of all clusters.
	DedicatedClusterStatusList *ListDedicatedClusterResponseBodyDedicatedClusterStatusList `json:"DedicatedClusterStatusList,omitempty" xml:"DedicatedClusterStatusList,omitempty" type:"Struct"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The page number of the returned page. The value must be an integer that is greater than 0. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of clusters that meet the query condition.
	//
	// example:
	//
	// 15
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListDedicatedClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDedicatedClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ListDedicatedClusterResponseBody) GetDedicatedClusterStatusList() *ListDedicatedClusterResponseBodyDedicatedClusterStatusList {
	return s.DedicatedClusterStatusList
}

func (s *ListDedicatedClusterResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ListDedicatedClusterResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ListDedicatedClusterResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ListDedicatedClusterResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDedicatedClusterResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListDedicatedClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDedicatedClusterResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ListDedicatedClusterResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListDedicatedClusterResponseBody) SetDedicatedClusterStatusList(v *ListDedicatedClusterResponseBodyDedicatedClusterStatusList) *ListDedicatedClusterResponseBody {
	s.DedicatedClusterStatusList = v
	return s
}

func (s *ListDedicatedClusterResponseBody) SetErrCode(v string) *ListDedicatedClusterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ListDedicatedClusterResponseBody) SetErrMessage(v string) *ListDedicatedClusterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ListDedicatedClusterResponseBody) SetHttpStatusCode(v string) *ListDedicatedClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDedicatedClusterResponseBody) SetPageNumber(v int32) *ListDedicatedClusterResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListDedicatedClusterResponseBody) SetPageRecordCount(v int32) *ListDedicatedClusterResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListDedicatedClusterResponseBody) SetRequestId(v string) *ListDedicatedClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDedicatedClusterResponseBody) SetSuccess(v string) *ListDedicatedClusterResponseBody {
	s.Success = &v
	return s
}

func (s *ListDedicatedClusterResponseBody) SetTotalRecordCount(v int32) *ListDedicatedClusterResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListDedicatedClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDedicatedClusterResponseBodyDedicatedClusterStatusList struct {
	DedicatedClusterStatus []*ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus `json:"DedicatedClusterStatus,omitempty" xml:"DedicatedClusterStatus,omitempty" type:"Repeated"`
}

func (s ListDedicatedClusterResponseBodyDedicatedClusterStatusList) String() string {
	return dara.Prettify(s)
}

func (s ListDedicatedClusterResponseBodyDedicatedClusterStatusList) GoString() string {
	return s.String()
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusList) GetDedicatedClusterStatus() []*ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	return s.DedicatedClusterStatus
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusList) SetDedicatedClusterStatus(v []*ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) *ListDedicatedClusterResponseBodyDedicatedClusterStatusList {
	s.DedicatedClusterStatus = v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusList) Validate() error {
	return dara.Validate(s)
}

type ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus struct {
	// The CPU utilization, in percentage.
	//
	// example:
	//
	// 30
	CpuUtilization *int64 `json:"CpuUtilization,omitempty" xml:"CpuUtilization,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// dtscluster*******
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// daily_test
	DedicatedClusterName *string `json:"DedicatedClusterName,omitempty" xml:"DedicatedClusterName,omitempty"`
	// The disk usage.
	//
	// example:
	//
	// 50
	DiskUtilization *int64 `json:"DiskUtilization,omitempty" xml:"DiskUtilization,omitempty"`
	// The ID of the DTS instance.
	//
	// example:
	//
	// dtsi03e3zty16i****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The number of DTS units (DUs).
	//
	// example:
	//
	// 30
	Du *int64 `json:"Du,omitempty" xml:"Du,omitempty"`
	// The DU usage, in percentage.
	//
	// example:
	//
	// 16.6667
	DuUtilization *int64 `json:"DuUtilization,omitempty" xml:"DuUtilization,omitempty"`
	// The time when the cluster was created.
	//
	// example:
	//
	// 1647424384606
	GmtCreated *int64 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The memory usage.
	//
	// example:
	//
	// 20
	MemUtilization *int64 `json:"MemUtilization,omitempty" xml:"MemUtilization,omitempty"`
	// The number of nodes in the cluster.
	//
	// example:
	//
	// 5
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The number of over-provisioned DUs.
	//
	// example:
	//
	// 60
	OversoldDu *int64 `json:"OversoldDu,omitempty" xml:"OversoldDu,omitempty"`
	// The ID of the region in which the DTS instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the cluster. Valid values:
	//
	// 	- **init**: The cluster is being initialized.
	//
	// 	- **schedule**: The cluster is pending scheduling.
	//
	// 	- **running**: The cluster is running.
	//
	// 	- **upgrade**: The cluster is being upgraded.
	//
	// 	- **downgrade**: The cluster is being downgraded.
	//
	// 	- **locked**: The cluster is locked.
	//
	// 	- **releasing**: The cluster is being released.
	//
	// 	- **released**: The cluster is released.
	//
	// example:
	//
	// init
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The total number of CPU cores.
	//
	// example:
	//
	// 100
	TotalCpuCore *int64 `json:"TotalCpuCore,omitempty" xml:"TotalCpuCore,omitempty"`
	// The total disk capacity. Unit: GB.
	//
	// example:
	//
	// 2048
	TotalDiskGBSize *int64 `json:"TotalDiskGBSize,omitempty" xml:"TotalDiskGBSize,omitempty"`
	// The total memory capacity. Unit: GB.
	//
	// example:
	//
	// 256
	TotalMemGBSize *int64 `json:"TotalMemGBSize,omitempty" xml:"TotalMemGBSize,omitempty"`
	// The number of used CPU cores.
	//
	// example:
	//
	// 30
	UsedCpuCore *int64 `json:"UsedCpuCore,omitempty" xml:"UsedCpuCore,omitempty"`
	// The used disk capacity. Unit: GB.
	//
	// example:
	//
	// 1024
	UsedDiskGBSize *int64 `json:"UsedDiskGBSize,omitempty" xml:"UsedDiskGBSize,omitempty"`
	// The number of used DUs.
	//
	// example:
	//
	// 5
	UsedDu *int64 `json:"UsedDu,omitempty" xml:"UsedDu,omitempty"`
	// The used memory capacity. Unit: GB.
	//
	// example:
	//
	// 128
	UsedMemGBSize *int64 `json:"UsedMemGBSize,omitempty" xml:"UsedMemGBSize,omitempty"`
}

func (s ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) String() string {
	return dara.Prettify(s)
}

func (s ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GoString() string {
	return s.String()
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetCpuUtilization() *int64 {
	return s.CpuUtilization
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetDedicatedClusterName() *string {
	return s.DedicatedClusterName
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetDiskUtilization() *int64 {
	return s.DiskUtilization
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetDu() *int64 {
	return s.Du
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetDuUtilization() *int64 {
	return s.DuUtilization
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetMemUtilization() *int64 {
	return s.MemUtilization
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetOversoldDu() *int64 {
	return s.OversoldDu
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetState() *string {
	return s.State
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetTotalCpuCore() *int64 {
	return s.TotalCpuCore
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetTotalDiskGBSize() *int64 {
	return s.TotalDiskGBSize
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetTotalMemGBSize() *int64 {
	return s.TotalMemGBSize
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetUsedCpuCore() *int64 {
	return s.UsedCpuCore
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetUsedDiskGBSize() *int64 {
	return s.UsedDiskGBSize
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetUsedDu() *int64 {
	return s.UsedDu
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) GetUsedMemGBSize() *int64 {
	return s.UsedMemGBSize
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetCpuUtilization(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.CpuUtilization = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetDedicatedClusterId(v string) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.DedicatedClusterId = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetDedicatedClusterName(v string) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.DedicatedClusterName = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetDiskUtilization(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.DiskUtilization = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetDtsInstanceID(v string) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.DtsInstanceID = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetDu(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.Du = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetDuUtilization(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.DuUtilization = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetGmtCreated(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.GmtCreated = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetMemUtilization(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.MemUtilization = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetNodeCount(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.NodeCount = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetOversoldDu(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.OversoldDu = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetRegionId(v string) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.RegionId = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetState(v string) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.State = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetTotalCpuCore(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.TotalCpuCore = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetTotalDiskGBSize(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.TotalDiskGBSize = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetTotalMemGBSize(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.TotalMemGBSize = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetUsedCpuCore(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.UsedCpuCore = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetUsedDiskGBSize(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.UsedDiskGBSize = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetUsedDu(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.UsedDu = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) SetUsedMemGBSize(v int64) *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus {
	s.UsedMemGBSize = &v
	return s
}

func (s *ListDedicatedClusterResponseBodyDedicatedClusterStatusListDedicatedClusterStatus) Validate() error {
	return dara.Validate(s)
}
