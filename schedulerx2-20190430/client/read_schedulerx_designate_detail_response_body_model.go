// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxDesignateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) *ReadSchedulerxDesignateDetailResponseBody
	GetAccessDeniedDetail() *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail
	SetCode(v int32) *ReadSchedulerxDesignateDetailResponseBody
	GetCode() *int32
	SetData(v *ReadSchedulerxDesignateDetailResponseBodyData) *ReadSchedulerxDesignateDetailResponseBody
	GetData() *ReadSchedulerxDesignateDetailResponseBodyData
	SetMessage(v string) *ReadSchedulerxDesignateDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReadSchedulerxDesignateDetailResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReadSchedulerxDesignateDetailResponseBody
	GetSuccess() *bool
}

type ReadSchedulerxDesignateDetailResponseBody struct {
	AccessDeniedDetail *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ReadSchedulerxDesignateDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Invalid parameter: jobId=368 invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 883AFE93-FB03-4FA9-A958-E750C6DE120C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReadSchedulerxDesignateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateDetailResponseBody) GetAccessDeniedDetail() *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *ReadSchedulerxDesignateDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReadSchedulerxDesignateDetailResponseBody) GetData() *ReadSchedulerxDesignateDetailResponseBodyData {
	return s.Data
}

func (s *ReadSchedulerxDesignateDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReadSchedulerxDesignateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReadSchedulerxDesignateDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReadSchedulerxDesignateDetailResponseBody) SetAccessDeniedDetail(v *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) *ReadSchedulerxDesignateDetailResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBody) SetCode(v int32) *ReadSchedulerxDesignateDetailResponseBody {
	s.Code = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBody) SetData(v *ReadSchedulerxDesignateDetailResponseBodyData) *ReadSchedulerxDesignateDetailResponseBody {
	s.Data = v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBody) SetMessage(v string) *ReadSchedulerxDesignateDetailResponseBody {
	s.Message = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBody) SetRequestId(v string) *ReadSchedulerxDesignateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBody) SetSuccess(v bool) *ReadSchedulerxDesignateDetailResponseBody {
	s.Success = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBody) Validate() error {
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

type ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail struct {
	// example:
	//
	// edas:ReadSchedulerxDesignateDetail
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

func (s ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) SetAuthAction(v string) *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) SetPolicyType(v string) *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type ReadSchedulerxDesignateDetailResponseBodyData struct {
	// -
	DesignateDetailVos []*ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos `json:"DesignateDetailVos,omitempty" xml:"DesignateDetailVos,omitempty" type:"Repeated"`
}

func (s ReadSchedulerxDesignateDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateDetailResponseBodyData) GetDesignateDetailVos() []*ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	return s.DesignateDetailVos
}

func (s *ReadSchedulerxDesignateDetailResponseBodyData) SetDesignateDetailVos(v []*ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) *ReadSchedulerxDesignateDetailResponseBodyData {
	s.DesignateDetailVos = v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyData) Validate() error {
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

type ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos struct {
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
	Key     *string                                                                 `json:"Key,omitempty" xml:"Key,omitempty"`
	Metrics *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Struct"`
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

func (s ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GetBusy() *string {
	return s.Busy
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GetChecked() *bool {
	return s.Checked
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GetKey() *string {
	return s.Key
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GetMetrics() *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	return s.Metrics
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GetOffline() *bool {
	return s.Offline
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GetSize() *int32 {
	return s.Size
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GetStarter() *string {
	return s.Starter
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) GetVersion() *string {
	return s.Version
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) SetBusy(v string) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	s.Busy = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) SetChecked(v bool) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	s.Checked = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) SetKey(v string) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	s.Key = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) SetMetrics(v *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	s.Metrics = v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) SetOffline(v bool) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	s.Offline = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) SetSize(v int32) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	s.Size = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) SetStarter(v string) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	s.Starter = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) SetVersion(v string) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos {
	s.Version = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVos) Validate() error {
	if s.Metrics != nil {
		if err := s.Metrics.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics struct {
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
	// 0.14865875
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

func (s ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetCpuLoad1() *float64 {
	return s.CpuLoad1
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetCpuLoad5() *float64 {
	return s.CpuLoad5
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetCpuProcessors() *int32 {
	return s.CpuProcessors
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetDiskMax() *int32 {
	return s.DiskMax
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetDiskUsage() *float64 {
	return s.DiskUsage
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetDiskUsed() *int32 {
	return s.DiskUsed
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetExecCount() *int64 {
	return s.ExecCount
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetHeap1Usage() *float64 {
	return s.Heap1Usage
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetHeap1Used() *float64 {
	return s.Heap1Used
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetHeap5Usage() *float64 {
	return s.Heap5Usage
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetHeapMax() *int32 {
	return s.HeapMax
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetSharePoolAvailableSize() *int32 {
	return s.SharePoolAvailableSize
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) GetSharePoolQueueSize() *int32 {
	return s.SharePoolQueueSize
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetCpuLoad1(v float64) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.CpuLoad1 = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetCpuLoad5(v float64) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.CpuLoad5 = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetCpuProcessors(v int32) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.CpuProcessors = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetDiskMax(v int32) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.DiskMax = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetDiskUsage(v float64) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.DiskUsage = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetDiskUsed(v int32) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.DiskUsed = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetExecCount(v int64) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.ExecCount = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetHeap1Usage(v float64) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.Heap1Usage = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetHeap1Used(v float64) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.Heap1Used = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetHeap5Usage(v float64) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.Heap5Usage = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetHeapMax(v int32) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.HeapMax = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetSharePoolAvailableSize(v int32) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.SharePoolAvailableSize = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) SetSharePoolQueueSize(v int32) *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics {
	s.SharePoolQueueSize = &v
	return s
}

func (s *ReadSchedulerxDesignateDetailResponseBodyDataDesignateDetailVosMetrics) Validate() error {
	return dara.Validate(s)
}
