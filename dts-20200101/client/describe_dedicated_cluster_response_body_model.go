// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCpuUtilization(v int64) *DescribeDedicatedClusterResponseBody
	GetCpuUtilization() *int64
	SetDedicatedClusterId(v string) *DescribeDedicatedClusterResponseBody
	GetDedicatedClusterId() *string
	SetDedicatedClusterName(v string) *DescribeDedicatedClusterResponseBody
	GetDedicatedClusterName() *string
	SetDiskUtilization(v int64) *DescribeDedicatedClusterResponseBody
	GetDiskUtilization() *int64
	SetDtsInstanceID(v string) *DescribeDedicatedClusterResponseBody
	GetDtsInstanceID() *string
	SetDu(v int64) *DescribeDedicatedClusterResponseBody
	GetDu() *int64
	SetDuUtilization(v int64) *DescribeDedicatedClusterResponseBody
	GetDuUtilization() *int64
	SetErrCode(v string) *DescribeDedicatedClusterResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDedicatedClusterResponseBody
	GetErrMessage() *string
	SetGmtCreated(v int64) *DescribeDedicatedClusterResponseBody
	GetGmtCreated() *int64
	SetGmtFinished(v int64) *DescribeDedicatedClusterResponseBody
	GetGmtFinished() *int64
	SetHttpStatusCode(v string) *DescribeDedicatedClusterResponseBody
	GetHttpStatusCode() *string
	SetMemUtilization(v int64) *DescribeDedicatedClusterResponseBody
	GetMemUtilization() *int64
	SetNodeCount(v int64) *DescribeDedicatedClusterResponseBody
	GetNodeCount() *int64
	SetOversoldDu(v int64) *DescribeDedicatedClusterResponseBody
	GetOversoldDu() *int64
	SetRegionId(v string) *DescribeDedicatedClusterResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeDedicatedClusterResponseBody
	GetRequestId() *string
	SetState(v string) *DescribeDedicatedClusterResponseBody
	GetState() *string
	SetSuccess(v string) *DescribeDedicatedClusterResponseBody
	GetSuccess() *string
	SetTotalCpuCore(v int64) *DescribeDedicatedClusterResponseBody
	GetTotalCpuCore() *int64
	SetTotalDiskGBSize(v int64) *DescribeDedicatedClusterResponseBody
	GetTotalDiskGBSize() *int64
	SetTotalMemGBSize(v int64) *DescribeDedicatedClusterResponseBody
	GetTotalMemGBSize() *int64
	SetUsedCpuCore(v int64) *DescribeDedicatedClusterResponseBody
	GetUsedCpuCore() *int64
	SetUsedDiskGBSize(v int64) *DescribeDedicatedClusterResponseBody
	GetUsedDiskGBSize() *int64
	SetUsedDu(v int64) *DescribeDedicatedClusterResponseBody
	GetUsedDu() *int64
	SetUsedMemGBSize(v int64) *DescribeDedicatedClusterResponseBody
	GetUsedMemGBSize() *int64
}

type DescribeDedicatedClusterResponseBody struct {
	// The CPU utilization. Unit: percentage.
	//
	// example:
	//
	// 30
	CpuUtilization *int64 `json:"CpuUtilization,omitempty" xml:"CpuUtilization,omitempty"`
	// The ID of the cluster.
	//
	// example:
	//
	// dtsCluster****
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
	// The ID of the instance.
	//
	// example:
	//
	// dtsb8r****
	DtsInstanceID *string `json:"DtsInstanceID,omitempty" xml:"DtsInstanceID,omitempty"`
	// The number of DTS units (DUs).
	//
	// example:
	//
	// 30
	Du *int64 `json:"Du,omitempty" xml:"Du,omitempty"`
	// The DU usage. Unit: percentage.
	//
	// example:
	//
	// 16
	DuUtilization *int64 `json:"DuUtilization,omitempty" xml:"DuUtilization,omitempty"`
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
	// The time when the cluster was created.
	//
	// example:
	//
	// 1642476144000
	GmtCreated *int64 `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the cluster stopped.
	//
	// example:
	//
	// 1645200000000
	GmtFinished *int64 `json:"GmtFinished,omitempty" xml:"GmtFinished,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
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
	// The number of DUs that exceeds the upper limit.
	//
	// example:
	//
	// 60
	OversoldDu *int64 `json:"OversoldDu,omitempty" xml:"OversoldDu,omitempty"`
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// inti
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of CPU cores.
	//
	// example:
	//
	// 100
	TotalCpuCore *int64 `json:"TotalCpuCore,omitempty" xml:"TotalCpuCore,omitempty"`
	// The total disk size. Unit: GB.
	//
	// example:
	//
	// 2048
	TotalDiskGBSize *int64 `json:"TotalDiskGBSize,omitempty" xml:"TotalDiskGBSize,omitempty"`
	// The total amount of memory. Unit: GB.
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
	// The used disk size. Unit: GB.
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
	// The amount of used memory. Unit: GB.
	//
	// example:
	//
	// 128
	UsedMemGBSize *int64 `json:"UsedMemGBSize,omitempty" xml:"UsedMemGBSize,omitempty"`
}

func (s DescribeDedicatedClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterResponseBody) GetCpuUtilization() *int64 {
	return s.CpuUtilization
}

func (s *DescribeDedicatedClusterResponseBody) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *DescribeDedicatedClusterResponseBody) GetDedicatedClusterName() *string {
	return s.DedicatedClusterName
}

func (s *DescribeDedicatedClusterResponseBody) GetDiskUtilization() *int64 {
	return s.DiskUtilization
}

func (s *DescribeDedicatedClusterResponseBody) GetDtsInstanceID() *string {
	return s.DtsInstanceID
}

func (s *DescribeDedicatedClusterResponseBody) GetDu() *int64 {
	return s.Du
}

func (s *DescribeDedicatedClusterResponseBody) GetDuUtilization() *int64 {
	return s.DuUtilization
}

func (s *DescribeDedicatedClusterResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDedicatedClusterResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDedicatedClusterResponseBody) GetGmtCreated() *int64 {
	return s.GmtCreated
}

func (s *DescribeDedicatedClusterResponseBody) GetGmtFinished() *int64 {
	return s.GmtFinished
}

func (s *DescribeDedicatedClusterResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeDedicatedClusterResponseBody) GetMemUtilization() *int64 {
	return s.MemUtilization
}

func (s *DescribeDedicatedClusterResponseBody) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *DescribeDedicatedClusterResponseBody) GetOversoldDu() *int64 {
	return s.OversoldDu
}

func (s *DescribeDedicatedClusterResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedClusterResponseBody) GetState() *string {
	return s.State
}

func (s *DescribeDedicatedClusterResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDedicatedClusterResponseBody) GetTotalCpuCore() *int64 {
	return s.TotalCpuCore
}

func (s *DescribeDedicatedClusterResponseBody) GetTotalDiskGBSize() *int64 {
	return s.TotalDiskGBSize
}

func (s *DescribeDedicatedClusterResponseBody) GetTotalMemGBSize() *int64 {
	return s.TotalMemGBSize
}

func (s *DescribeDedicatedClusterResponseBody) GetUsedCpuCore() *int64 {
	return s.UsedCpuCore
}

func (s *DescribeDedicatedClusterResponseBody) GetUsedDiskGBSize() *int64 {
	return s.UsedDiskGBSize
}

func (s *DescribeDedicatedClusterResponseBody) GetUsedDu() *int64 {
	return s.UsedDu
}

func (s *DescribeDedicatedClusterResponseBody) GetUsedMemGBSize() *int64 {
	return s.UsedMemGBSize
}

func (s *DescribeDedicatedClusterResponseBody) SetCpuUtilization(v int64) *DescribeDedicatedClusterResponseBody {
	s.CpuUtilization = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetDedicatedClusterId(v string) *DescribeDedicatedClusterResponseBody {
	s.DedicatedClusterId = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetDedicatedClusterName(v string) *DescribeDedicatedClusterResponseBody {
	s.DedicatedClusterName = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetDiskUtilization(v int64) *DescribeDedicatedClusterResponseBody {
	s.DiskUtilization = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetDtsInstanceID(v string) *DescribeDedicatedClusterResponseBody {
	s.DtsInstanceID = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetDu(v int64) *DescribeDedicatedClusterResponseBody {
	s.Du = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetDuUtilization(v int64) *DescribeDedicatedClusterResponseBody {
	s.DuUtilization = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetErrCode(v string) *DescribeDedicatedClusterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetErrMessage(v string) *DescribeDedicatedClusterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetGmtCreated(v int64) *DescribeDedicatedClusterResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetGmtFinished(v int64) *DescribeDedicatedClusterResponseBody {
	s.GmtFinished = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetHttpStatusCode(v string) *DescribeDedicatedClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetMemUtilization(v int64) *DescribeDedicatedClusterResponseBody {
	s.MemUtilization = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetNodeCount(v int64) *DescribeDedicatedClusterResponseBody {
	s.NodeCount = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetOversoldDu(v int64) *DescribeDedicatedClusterResponseBody {
	s.OversoldDu = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetRegionId(v string) *DescribeDedicatedClusterResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetRequestId(v string) *DescribeDedicatedClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetState(v string) *DescribeDedicatedClusterResponseBody {
	s.State = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetSuccess(v string) *DescribeDedicatedClusterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetTotalCpuCore(v int64) *DescribeDedicatedClusterResponseBody {
	s.TotalCpuCore = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetTotalDiskGBSize(v int64) *DescribeDedicatedClusterResponseBody {
	s.TotalDiskGBSize = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetTotalMemGBSize(v int64) *DescribeDedicatedClusterResponseBody {
	s.TotalMemGBSize = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetUsedCpuCore(v int64) *DescribeDedicatedClusterResponseBody {
	s.UsedCpuCore = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetUsedDiskGBSize(v int64) *DescribeDedicatedClusterResponseBody {
	s.UsedDiskGBSize = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetUsedDu(v int64) *DescribeDedicatedClusterResponseBody {
	s.UsedDu = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) SetUsedMemGBSize(v int64) *DescribeDedicatedClusterResponseBody {
	s.UsedMemGBSize = &v
	return s
}

func (s *DescribeDedicatedClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
