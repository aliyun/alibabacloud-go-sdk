// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxDesignateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) *ReadSchedulerxDesignateInfoResponseBody
	GetAccessDeniedDetail() *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail
	SetCode(v int32) *ReadSchedulerxDesignateInfoResponseBody
	GetCode() *int32
	SetData(v *ReadSchedulerxDesignateInfoResponseBodyData) *ReadSchedulerxDesignateInfoResponseBody
	GetData() *ReadSchedulerxDesignateInfoResponseBodyData
	SetMessage(v string) *ReadSchedulerxDesignateInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadSchedulerxDesignateInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadSchedulerxDesignateInfoResponseBody
	GetSuccess() *bool
}

type ReadSchedulerxDesignateInfoResponseBody struct {
	// The access denial details.
	AccessDeniedDetail *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// *
	Data *ReadSchedulerxDesignateInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned only if an error occurs.
	//
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadSchedulerxDesignateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateInfoResponseBody) GetAccessDeniedDetail() *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ReadSchedulerxDesignateInfoResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReadSchedulerxDesignateInfoResponseBody) GetData() *ReadSchedulerxDesignateInfoResponseBodyData {
	return s.Data
}

func (s *ReadSchedulerxDesignateInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadSchedulerxDesignateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadSchedulerxDesignateInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadSchedulerxDesignateInfoResponseBody) SetAccessDeniedDetail(v *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) *ReadSchedulerxDesignateInfoResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBody) SetCode(v int32) *ReadSchedulerxDesignateInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBody) SetData(v *ReadSchedulerxDesignateInfoResponseBodyData) *ReadSchedulerxDesignateInfoResponseBody {
	s.Data = v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBody) SetMessage(v string) *ReadSchedulerxDesignateInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBody) SetRequestId(v string) *ReadSchedulerxDesignateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBody) SetSuccess(v bool) *ReadSchedulerxDesignateInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail struct {
	// The authentication operation.
	//
	// example:
	//
	// edas:ReadSchedulerxDesignateInfo
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The principal name.
	//
	// example:
	//
	// 209312833131416xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The principal account.
	//
	// example:
	//
	// 1827811800526xxx
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The principal type.
	//
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQFn/cLPZ/3Cz0YxQkZBMjVGLTY0REUtNTlGNS05NzUwLTgyMUE4M0MwMTFDRQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The permission denial type.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// AccountLevelIdentityBasedPolicy
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ReadSchedulerxDesignateInfoResponseBodyData struct {
	// *
	DesignateDetailVos []*ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos `json:"DesignateDetailVos,omitempty" xml:"DesignateDetailVos,omitempty" type:"Repeated"`
	// The information type of the specified workers.
	//
	// 	- 1: the IP address of the specified workers.
	//
	// 	- 2: the tags of the specified workers.
	//
	// >  The default value of the DesignateType parameter is 1.
	//
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
	// Indicates whether to enable failover for the workers. If you set this parameter to true, the job is scheduled to other workers when the specified workers go offline.
	//
	// 	- true: enables failover for the workers.
	//
	// 	- false: disables failover for the workers.
	//
	// >  The default value of the Transferable parameter is false.
	//
	// example:
	//
	// true
	Transferable *bool `json:"Transferable,omitempty" xml:"Transferable,omitempty"`
}

func (s ReadSchedulerxDesignateInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateInfoResponseBodyData) GetDesignateDetailVos() []*ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	return s.DesignateDetailVos
}

func (s *ReadSchedulerxDesignateInfoResponseBodyData) GetDesignateType() *int32 {
	return s.DesignateType
}

func (s *ReadSchedulerxDesignateInfoResponseBodyData) GetTransferable() *bool {
	return s.Transferable
}

func (s *ReadSchedulerxDesignateInfoResponseBodyData) SetDesignateDetailVos(v []*ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) *ReadSchedulerxDesignateInfoResponseBodyData {
	s.DesignateDetailVos = v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyData) SetDesignateType(v int32) *ReadSchedulerxDesignateInfoResponseBodyData {
	s.DesignateType = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyData) SetTransferable(v bool) *ReadSchedulerxDesignateInfoResponseBodyData {
	s.Transferable = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyData) Validate() error {
	if s.DesignateDetailVos != nil {
		for _, item := range s.DesignateDetailVos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos struct {
	// The status of the workers. Valid values:
	//
	// 	- FREE: idle.
	//
	// 	- LOAD5_BUSY: The average of the latest five values of CPU workload is too high.
	//
	// 	- HEAP5_BUSY: The average of the latest five values of heap memory usage is too high.
	//
	// 	- DISK_BUSY: The disk usage is too high.
	//
	// example:
	//
	// FREE
	Busy *string `json:"Busy,omitempty" xml:"Busy,omitempty"`
	// Indicates whether the workers are specified.
	//
	// 	- true: The workers are specified.
	//
	// 	- false: The workers are not specified.
	//
	// example:
	//
	// true
	Checked *bool `json:"Checked,omitempty" xml:"Checked,omitempty"`
	// The information returned based on the value of the DesignateType parameter.
	//
	// 	- If you set the DesignateType parameter to 2, the tags of the workers are returned.
	//
	// 	- If you set the DesignateType parameter to 1, the IP addresses of the workers are returned.
	//
	// example:
	//
	// 10.52.169.25
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The metric values.
	Metrics *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// Indicates whether the workers are offline.
	//
	// example:
	//
	// fasle
	Offline *bool `json:"Offline,omitempty" xml:"Offline,omitempty"`
	// The number of workers.
	//
	// example:
	//
	// 1
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The startup method of the workers.
	//
	// example:
	//
	// pod
	Starter *string `json:"Starter,omitempty" xml:"Starter,omitempty"`
	// The version of the workers.
	//
	// example:
	//
	// 1.12.5
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GetBusy() *string {
	return s.Busy
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GetChecked() *bool {
	return s.Checked
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GetKey() *string {
	return s.Key
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GetMetrics() *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	return s.Metrics
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GetOffline() *bool {
	return s.Offline
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GetSize() *int32 {
	return s.Size
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GetStarter() *string {
	return s.Starter
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) GetVersion() *string {
	return s.Version
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) SetBusy(v string) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	s.Busy = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) SetChecked(v bool) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	s.Checked = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) SetKey(v string) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	s.Key = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) SetMetrics(v *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	s.Metrics = v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) SetOffline(v bool) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	s.Offline = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) SetSize(v int32) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	s.Size = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) SetStarter(v string) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	s.Starter = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) SetVersion(v string) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos {
	s.Version = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos) Validate() error {
	if s.Metrics != nil {
		if err := s.Metrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics struct {
	// The most recent value of CPU workload.
	//
	// example:
	//
	// 0
	CpuLoad1 *float64 `json:"CpuLoad1,omitempty" xml:"CpuLoad1,omitempty"`
	// The average of the latest five values of CPU workload.
	//
	// example:
	//
	// 0
	CpuLoad5 *float64 `json:"CpuLoad5,omitempty" xml:"CpuLoad5,omitempty"`
	// The number of available CPU processors.
	//
	// example:
	//
	// 1
	CpuProcessors *int32 `json:"CpuProcessors,omitempty" xml:"CpuProcessors,omitempty"`
	// The total disk capacity in MB.
	//
	// example:
	//
	// 1024
	DiskMax *int32 `json:"DiskMax,omitempty" xml:"DiskMax,omitempty"`
	// The disk usage.
	//
	// example:
	//
	// 0.19142496008515167
	DiskUsage *float64 `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// The used disk space in MB.
	//
	// example:
	//
	// 148
	DiskUsed *int32 `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	// The number of job executions.
	//
	// example:
	//
	// 56
	ExecCount *int64 `json:"ExecCount,omitempty" xml:"ExecCount,omitempty"`
	// The most recent value of heap memory usage.
	//
	// example:
	//
	// 0.06072874493927125
	Heap1Usage *float64 `json:"Heap1Usage,omitempty" xml:"Heap1Usage,omitempty"`
	// The most recent value of used heap memory in MB.
	//
	// example:
	//
	// 14
	Heap1Used *float64 `json:"Heap1Used,omitempty" xml:"Heap1Used,omitempty"`
	// The average of the latest five values of heap memory usage.
	//
	// example:
	//
	// 0.06477732793522267
	Heap5Usage *float64 `json:"Heap5Usage,omitempty" xml:"Heap5Usage,omitempty"`
	// The maximum heap memory in MB.
	//
	// example:
	//
	// 500
	HeapMax *int32 `json:"HeapMax,omitempty" xml:"HeapMax,omitempty"`
	// The number of available resources in the shared pool.
	//
	// example:
	//
	// 72
	SharePoolAvailableSize *int32 `json:"SharePoolAvailableSize,omitempty" xml:"SharePoolAvailableSize,omitempty"`
	// The queue size in the shared pool.
	//
	// example:
	//
	// 1
	SharePoolQueueSize *int32 `json:"SharePoolQueueSize,omitempty" xml:"SharePoolQueueSize,omitempty"`
}

func (s ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetCpuLoad1() *float64 {
	return s.CpuLoad1
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetCpuLoad5() *float64 {
	return s.CpuLoad5
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetCpuProcessors() *int32 {
	return s.CpuProcessors
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetDiskMax() *int32 {
	return s.DiskMax
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetDiskUsage() *float64 {
	return s.DiskUsage
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetDiskUsed() *int32 {
	return s.DiskUsed
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetExecCount() *int64 {
	return s.ExecCount
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetHeap1Usage() *float64 {
	return s.Heap1Usage
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetHeap1Used() *float64 {
	return s.Heap1Used
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetHeap5Usage() *float64 {
	return s.Heap5Usage
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetHeapMax() *int32 {
	return s.HeapMax
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetSharePoolAvailableSize() *int32 {
	return s.SharePoolAvailableSize
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) GetSharePoolQueueSize() *int32 {
	return s.SharePoolQueueSize
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetCpuLoad1(v float64) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.CpuLoad1 = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetCpuLoad5(v float64) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.CpuLoad5 = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetCpuProcessors(v int32) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.CpuProcessors = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetDiskMax(v int32) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.DiskMax = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetDiskUsage(v float64) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.DiskUsage = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetDiskUsed(v int32) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.DiskUsed = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetExecCount(v int64) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.ExecCount = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetHeap1Usage(v float64) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.Heap1Usage = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetHeap1Used(v float64) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.Heap1Used = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetHeap5Usage(v float64) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.Heap5Usage = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetHeapMax(v int32) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.HeapMax = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetSharePoolAvailableSize(v int32) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.SharePoolAvailableSize = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) SetSharePoolQueueSize(v int32) *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics {
	s.SharePoolQueueSize = &v
	return s
}

func (s *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics) Validate() error {
	return dara.Validate(s)
}
