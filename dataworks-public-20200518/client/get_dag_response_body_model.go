// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDagResponseBodyData) *GetDagResponseBody
	GetData() *GetDagResponseBodyData
	SetErrorCode(v string) *GetDagResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDagResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *GetDagResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *GetDagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDagResponseBody
	GetSuccess() *bool
}

type GetDagResponseBody struct {
	// The details of the DAG.
	Data *GetDagResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// Invalid.Tenant.ConnectionNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The connection does not exist.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7782a299-b291-4fee-8424-cf8058efa8e8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDagResponseBody) GoString() string {
	return s.String()
}

func (s *GetDagResponseBody) GetData() *GetDagResponseBodyData {
	return s.Data
}

func (s *GetDagResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDagResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDagResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDagResponseBody) SetData(v *GetDagResponseBodyData) *GetDagResponseBody {
	s.Data = v
	return s
}

func (s *GetDagResponseBody) SetErrorCode(v string) *GetDagResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDagResponseBody) SetErrorMessage(v string) *GetDagResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDagResponseBody) SetHttpStatusCode(v int32) *GetDagResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDagResponseBody) SetRequestId(v string) *GetDagResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDagResponseBody) SetSuccess(v bool) *GetDagResponseBody {
	s.Success = &v
	return s
}

func (s *GetDagResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDagResponseBodyData struct {
	// The data timestamp.
	//
	// example:
	//
	// 1605052800000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The time when the DAG was created.
	//
	// example:
	//
	// 1605052800000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator.
	//
	// example:
	//
	// 1736629400048545
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The DAG ID.
	//
	// example:
	//
	// 351249682
	DagId *int64 `json:"DagId,omitempty" xml:"DagId,omitempty"`
	// The time when the DAG finished running.
	//
	// example:
	//
	// 1605052800000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The time when the DAG was scheduled to run.
	//
	// example:
	//
	// 1605052800000
	Gmtdate *int64 `json:"Gmtdate,omitempty" xml:"Gmtdate,omitempty"`
	// The time when the DAG was last modified.
	//
	// example:
	//
	// 1605052800000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name of the DAG.
	//
	// example:
	//
	// test_dag
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The sequence number of the operation.
	//
	// example:
	//
	// 123
	OpSeq *int64 `json:"OpSeq,omitempty" xml:"OpSeq,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 112345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The time when the DAG started to run.
	//
	// example:
	//
	// 1605052800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the DAG. Valid values: CREATED, RUNNING, FAILURE, and SUCCESS.
	//
	// example:
	//
	// FAILURE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the DAG. Valid values: MANUAL, SMOKE_TEST, SUPPLY_DATA, and BUSINESS_PROCESS_DAG. The value MANUAL indicates the DAG for a manually triggered workflow. The value SMOKE_TEST indicates the DAG for a smoke testing workflow. The value SUPPLY_DATA indicates the DAG for a data backfill instance. The value BUSINESS_PROCESS_DAG indicates the DAG for a one-time workflow.
	//
	// example:
	//
	// MANUAL_FLOW
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetDagResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDagResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDagResponseBodyData) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *GetDagResponseBodyData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetDagResponseBodyData) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetDagResponseBodyData) GetDagId() *int64 {
	return s.DagId
}

func (s *GetDagResponseBodyData) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *GetDagResponseBodyData) GetGmtdate() *int64 {
	return s.Gmtdate
}

func (s *GetDagResponseBodyData) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *GetDagResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetDagResponseBodyData) GetOpSeq() *int64 {
	return s.OpSeq
}

func (s *GetDagResponseBodyData) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetDagResponseBodyData) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetDagResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetDagResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetDagResponseBodyData) SetBizdate(v int64) *GetDagResponseBodyData {
	s.Bizdate = &v
	return s
}

func (s *GetDagResponseBodyData) SetCreateTime(v int64) *GetDagResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetDagResponseBodyData) SetCreateUser(v string) *GetDagResponseBodyData {
	s.CreateUser = &v
	return s
}

func (s *GetDagResponseBodyData) SetDagId(v int64) *GetDagResponseBodyData {
	s.DagId = &v
	return s
}

func (s *GetDagResponseBodyData) SetFinishTime(v int64) *GetDagResponseBodyData {
	s.FinishTime = &v
	return s
}

func (s *GetDagResponseBodyData) SetGmtdate(v int64) *GetDagResponseBodyData {
	s.Gmtdate = &v
	return s
}

func (s *GetDagResponseBodyData) SetModifyTime(v int64) *GetDagResponseBodyData {
	s.ModifyTime = &v
	return s
}

func (s *GetDagResponseBodyData) SetName(v string) *GetDagResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetDagResponseBodyData) SetOpSeq(v int64) *GetDagResponseBodyData {
	s.OpSeq = &v
	return s
}

func (s *GetDagResponseBodyData) SetProjectId(v int64) *GetDagResponseBodyData {
	s.ProjectId = &v
	return s
}

func (s *GetDagResponseBodyData) SetStartTime(v int64) *GetDagResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *GetDagResponseBodyData) SetStatus(v string) *GetDagResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetDagResponseBodyData) SetType(v string) *GetDagResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetDagResponseBodyData) Validate() error {
	return dara.Validate(s)
}
