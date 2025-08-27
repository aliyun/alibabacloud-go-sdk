// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTripTaskQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TripTaskQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TripTaskQueryResponseBody
	GetMessage() *string
	SetModule(v *TripTaskQueryResponseBodyModule) *TripTaskQueryResponseBody
	GetModule() *TripTaskQueryResponseBodyModule
	SetRequestId(v string) *TripTaskQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TripTaskQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TripTaskQueryResponseBody
	GetTraceId() *string
}

type TripTaskQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module。
	Module *TripTaskQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 2FB0D7A8-BA41-5D04-BEFC-CADA5481AC53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s TripTaskQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TripTaskQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TripTaskQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TripTaskQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TripTaskQueryResponseBody) GetModule() *TripTaskQueryResponseBodyModule {
	return s.Module
}

func (s *TripTaskQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TripTaskQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TripTaskQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TripTaskQueryResponseBody) SetCode(v string) *TripTaskQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TripTaskQueryResponseBody) SetMessage(v string) *TripTaskQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TripTaskQueryResponseBody) SetModule(v *TripTaskQueryResponseBodyModule) *TripTaskQueryResponseBody {
	s.Module = v
	return s
}

func (s *TripTaskQueryResponseBody) SetRequestId(v string) *TripTaskQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TripTaskQueryResponseBody) SetSuccess(v bool) *TripTaskQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TripTaskQueryResponseBody) SetTraceId(v string) *TripTaskQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TripTaskQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TripTaskQueryResponseBodyModule struct {
	// example:
	//
	// true
	NeedRefresh  *bool                                          `json:"needRefresh,omitempty" xml:"needRefresh,omitempty"`
	RecordTasks  []*TripTaskQueryResponseBodyModuleRecordTasks  `json:"record_tasks,omitempty" xml:"record_tasks,omitempty" type:"Repeated"`
	RunningTasks []*TripTaskQueryResponseBodyModuleRunningTasks `json:"running_tasks,omitempty" xml:"running_tasks,omitempty" type:"Repeated"`
}

func (s TripTaskQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TripTaskQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TripTaskQueryResponseBodyModule) GetNeedRefresh() *bool {
	return s.NeedRefresh
}

func (s *TripTaskQueryResponseBodyModule) GetRecordTasks() []*TripTaskQueryResponseBodyModuleRecordTasks {
	return s.RecordTasks
}

func (s *TripTaskQueryResponseBodyModule) GetRunningTasks() []*TripTaskQueryResponseBodyModuleRunningTasks {
	return s.RunningTasks
}

func (s *TripTaskQueryResponseBodyModule) SetNeedRefresh(v bool) *TripTaskQueryResponseBodyModule {
	s.NeedRefresh = &v
	return s
}

func (s *TripTaskQueryResponseBodyModule) SetRecordTasks(v []*TripTaskQueryResponseBodyModuleRecordTasks) *TripTaskQueryResponseBodyModule {
	s.RecordTasks = v
	return s
}

func (s *TripTaskQueryResponseBodyModule) SetRunningTasks(v []*TripTaskQueryResponseBodyModuleRunningTasks) *TripTaskQueryResponseBodyModule {
	s.RunningTasks = v
	return s
}

func (s *TripTaskQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TripTaskQueryResponseBodyModuleRecordTasks struct {
	// example:
	//
	// 123
	Actioner   *string `json:"actioner,omitempty" xml:"actioner,omitempty"`
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// example:
	//
	// 1525104000
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1525104000
	GmtFinished *int64 `json:"gmt_finished,omitempty" xml:"gmt_finished,omitempty"`
	// example:
	//
	// 34360
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 575003
	NodeId *string `json:"node_id,omitempty" xml:"node_id,omitempty"`
	// example:
	//
	// agree
	OutResult *string `json:"out_result,omitempty" xml:"out_result,omitempty"`
	// example:
	//
	// 123
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TripTaskQueryResponseBodyModuleRecordTasks) String() string {
	return dara.Prettify(s)
}

func (s TripTaskQueryResponseBodyModuleRecordTasks) GoString() string {
	return s.String()
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetActioner() *string {
	return s.Actioner
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetAttributes() *string {
	return s.Attributes
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetGmtFinished() *int64 {
	return s.GmtFinished
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetId() *int64 {
	return s.Id
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetNodeId() *string {
	return s.NodeId
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetOutResult() *string {
	return s.OutResult
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetOwner() *string {
	return s.Owner
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) GetStatus() *string {
	return s.Status
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetActioner(v string) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.Actioner = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetAttributes(v string) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.Attributes = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetGmtCreate(v int64) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.GmtCreate = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetGmtFinished(v int64) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.GmtFinished = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetId(v int64) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.Id = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetNodeId(v string) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.NodeId = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetOutResult(v string) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.OutResult = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetOwner(v string) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.Owner = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) SetStatus(v string) *TripTaskQueryResponseBodyModuleRecordTasks {
	s.Status = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRecordTasks) Validate() error {
	return dara.Validate(s)
}

type TripTaskQueryResponseBodyModuleRunningTasks struct {
	// example:
	//
	// 123
	Actioner   *string `json:"actioner,omitempty" xml:"actioner,omitempty"`
	Attributes *string `json:"attributes,omitempty" xml:"attributes,omitempty"`
	// example:
	//
	// 1525104000
	GmtCreate *int64 `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 1525104000
	GmtFinished *int64 `json:"gmt_finished,omitempty" xml:"gmt_finished,omitempty"`
	// example:
	//
	// 34360
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 575003
	NodeId *string `json:"node_id,omitempty" xml:"node_id,omitempty"`
	// example:
	//
	// agree
	OutResult *string `json:"out_result,omitempty" xml:"out_result,omitempty"`
	// example:
	//
	// 123
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// RUNNING
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s TripTaskQueryResponseBodyModuleRunningTasks) String() string {
	return dara.Prettify(s)
}

func (s TripTaskQueryResponseBodyModuleRunningTasks) GoString() string {
	return s.String()
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetActioner() *string {
	return s.Actioner
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetAttributes() *string {
	return s.Attributes
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetGmtFinished() *int64 {
	return s.GmtFinished
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetId() *int64 {
	return s.Id
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetNodeId() *string {
	return s.NodeId
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetOutResult() *string {
	return s.OutResult
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetOwner() *string {
	return s.Owner
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) GetStatus() *string {
	return s.Status
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetActioner(v string) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.Actioner = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetAttributes(v string) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.Attributes = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetGmtCreate(v int64) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.GmtCreate = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetGmtFinished(v int64) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.GmtFinished = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetId(v int64) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.Id = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetNodeId(v string) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.NodeId = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetOutResult(v string) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.OutResult = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetOwner(v string) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.Owner = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) SetStatus(v string) *TripTaskQueryResponseBodyModuleRunningTasks {
	s.Status = &v
	return s
}

func (s *TripTaskQueryResponseBodyModuleRunningTasks) Validate() error {
	return dara.Validate(s)
}
