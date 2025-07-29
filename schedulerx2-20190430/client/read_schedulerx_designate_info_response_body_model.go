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
	AccessDeniedDetail *ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ReadSchedulerxDesignateInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// jobId=xxx is not existed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401108
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	return dara.Validate(s)
}

type ReadSchedulerxDesignateInfoResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// edas:ReadSchedulerxDesignateInfo
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// example:
	//
	// 209312833131416xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// example:
	//
	// 1827811800526xxx
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// example:
	//
	// SubUser
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// example:
	//
	// AQFn/cLPZ/3Cz0YxQkZBMjVGLTY0REUtNTlGNS05NzUwLTgyMUE4M0MwMTFDRQ==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
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
	// -
	DesignateDetailVos []*ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos `json:"DesignateDetailVos,omitempty" xml:"DesignateDetailVos,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	DesignateType *int32 `json:"DesignateType,omitempty" xml:"DesignateType,omitempty"`
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
	return dara.Validate(s)
}

type ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVos struct {
	// example:
	//
	// FREE
	Busy *string `json:"Busy,omitempty" xml:"Busy,omitempty"`
	// example:
	//
	// true
	Checked *bool `json:"Checked,omitempty" xml:"Checked,omitempty"`
	// example:
	//
	// 10.52.169.25
	Key     *string                                                               `json:"Key,omitempty" xml:"Key,omitempty"`
	Metrics *ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
	// example:
	//
	// fasle
	Offline *bool `json:"Offline,omitempty" xml:"Offline,omitempty"`
	// example:
	//
	// 1
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// pod
	Starter *string `json:"Starter,omitempty" xml:"Starter,omitempty"`
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
	return dara.Validate(s)
}

type ReadSchedulerxDesignateInfoResponseBodyDataDesignateDetailVosMetrics struct {
	// example:
	//
	// 0
	CpuLoad1 *float64 `json:"CpuLoad1,omitempty" xml:"CpuLoad1,omitempty"`
	// example:
	//
	// 0
	CpuLoad5 *float64 `json:"CpuLoad5,omitempty" xml:"CpuLoad5,omitempty"`
	// example:
	//
	// 1
	CpuProcessors *int32 `json:"CpuProcessors,omitempty" xml:"CpuProcessors,omitempty"`
	// example:
	//
	// 1024
	DiskMax *int32 `json:"DiskMax,omitempty" xml:"DiskMax,omitempty"`
	// example:
	//
	// 0.19142496008515167
	DiskUsage *float64 `json:"DiskUsage,omitempty" xml:"DiskUsage,omitempty"`
	// example:
	//
	// 148
	DiskUsed *int32 `json:"DiskUsed,omitempty" xml:"DiskUsed,omitempty"`
	// example:
	//
	// 56
	ExecCount *int64 `json:"ExecCount,omitempty" xml:"ExecCount,omitempty"`
	// example:
	//
	// 0.06072874493927125
	Heap1Usage *float64 `json:"Heap1Usage,omitempty" xml:"Heap1Usage,omitempty"`
	// example:
	//
	// 14
	Heap1Used *float64 `json:"Heap1Used,omitempty" xml:"Heap1Used,omitempty"`
	// example:
	//
	// 0.06477732793522267
	Heap5Usage *float64 `json:"Heap5Usage,omitempty" xml:"Heap5Usage,omitempty"`
	// example:
	//
	// 500
	HeapMax *int32 `json:"HeapMax,omitempty" xml:"HeapMax,omitempty"`
	// example:
	//
	// 72
	SharePoolAvailableSize *int32 `json:"SharePoolAvailableSize,omitempty" xml:"SharePoolAvailableSize,omitempty"`
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
