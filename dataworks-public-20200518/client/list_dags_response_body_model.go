// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDagsResponseBodyData) *ListDagsResponseBody
	GetData() *ListDagsResponseBodyData
	SetErrorCode(v string) *ListDagsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListDagsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int32) *ListDagsResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ListDagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDagsResponseBody
	GetSuccess() *bool
}

type ListDagsResponseBody struct {
	// The details of DAGs.
	Data *ListDagsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The request ID. You can use the ID to locate logs and troubleshoot issues.
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

func (s ListDagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDagsResponseBody) GetData() *ListDagsResponseBodyData {
	return s.Data
}

func (s *ListDagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDagsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListDagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDagsResponseBody) SetData(v *ListDagsResponseBodyData) *ListDagsResponseBody {
	s.Data = v
	return s
}

func (s *ListDagsResponseBody) SetErrorCode(v string) *ListDagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDagsResponseBody) SetErrorMessage(v string) *ListDagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListDagsResponseBody) SetHttpStatusCode(v int32) *ListDagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDagsResponseBody) SetRequestId(v string) *ListDagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDagsResponseBody) SetSuccess(v bool) *ListDagsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDagsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDagsResponseBodyData struct {
	// The entities returned.
	Dags []*ListDagsResponseBodyDataDags `json:"Dags,omitempty" xml:"Dags,omitempty" type:"Repeated"`
}

func (s ListDagsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDagsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDagsResponseBodyData) GetDags() []*ListDagsResponseBodyDataDags {
	return s.Dags
}

func (s *ListDagsResponseBodyData) SetDags(v []*ListDagsResponseBodyDataDags) *ListDagsResponseBodyData {
	s.Dags = v
	return s
}

func (s *ListDagsResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ListDagsResponseBodyDataDags struct {
	// The data timestamp.
	//
	// example:
	//
	// 1605052800000
	Bizdate *int64 `json:"Bizdate,omitempty" xml:"Bizdate,omitempty"`
	// The creation time.
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
	// The end time.
	//
	// example:
	//
	// 1605052800000
	FinishTime *int64 `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The actual running time.
	//
	// example:
	//
	// 1605052800000
	Gmtdate *int64 `json:"Gmtdate,omitempty" xml:"Gmtdate,omitempty"`
	// The modification time.
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
	// The start time.
	//
	// example:
	//
	// 1605052800000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the DAG. Valid values:
	//
	// 	- CREATED
	//
	// 	- RUNNING
	//
	// 	- FAILURE
	//
	// 	- SUCCESS
	//
	// example:
	//
	// FAILURE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the DAG. Valid values:
	//
	// 	- MANUAL: DAG for a manually triggered workflow
	//
	// 	- SMOKE_TEST: DAG for a smoke testing workflow
	//
	// 	- SUPPLY_DATA: DAG for a data backfill instance
	//
	// 	- BUSINESS_PROCESS_DAG: DAG for a one-time workflow
	//
	// example:
	//
	// MANUAL_FLOW
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListDagsResponseBodyDataDags) String() string {
	return dara.Prettify(s)
}

func (s ListDagsResponseBodyDataDags) GoString() string {
	return s.String()
}

func (s *ListDagsResponseBodyDataDags) GetBizdate() *int64 {
	return s.Bizdate
}

func (s *ListDagsResponseBodyDataDags) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListDagsResponseBodyDataDags) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListDagsResponseBodyDataDags) GetDagId() *int64 {
	return s.DagId
}

func (s *ListDagsResponseBodyDataDags) GetFinishTime() *int64 {
	return s.FinishTime
}

func (s *ListDagsResponseBodyDataDags) GetGmtdate() *int64 {
	return s.Gmtdate
}

func (s *ListDagsResponseBodyDataDags) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListDagsResponseBodyDataDags) GetName() *string {
	return s.Name
}

func (s *ListDagsResponseBodyDataDags) GetOpSeq() *int64 {
	return s.OpSeq
}

func (s *ListDagsResponseBodyDataDags) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDagsResponseBodyDataDags) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListDagsResponseBodyDataDags) GetStatus() *string {
	return s.Status
}

func (s *ListDagsResponseBodyDataDags) GetType() *string {
	return s.Type
}

func (s *ListDagsResponseBodyDataDags) SetBizdate(v int64) *ListDagsResponseBodyDataDags {
	s.Bizdate = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetCreateTime(v int64) *ListDagsResponseBodyDataDags {
	s.CreateTime = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetCreateUser(v string) *ListDagsResponseBodyDataDags {
	s.CreateUser = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetDagId(v int64) *ListDagsResponseBodyDataDags {
	s.DagId = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetFinishTime(v int64) *ListDagsResponseBodyDataDags {
	s.FinishTime = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetGmtdate(v int64) *ListDagsResponseBodyDataDags {
	s.Gmtdate = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetModifyTime(v int64) *ListDagsResponseBodyDataDags {
	s.ModifyTime = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetName(v string) *ListDagsResponseBodyDataDags {
	s.Name = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetOpSeq(v int64) *ListDagsResponseBodyDataDags {
	s.OpSeq = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetProjectId(v int64) *ListDagsResponseBodyDataDags {
	s.ProjectId = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetStartTime(v int64) *ListDagsResponseBodyDataDags {
	s.StartTime = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetStatus(v string) *ListDagsResponseBodyDataDags {
	s.Status = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) SetType(v string) *ListDagsResponseBodyDataDags {
	s.Type = &v
	return s
}

func (s *ListDagsResponseBodyDataDags) Validate() error {
	return dara.Validate(s)
}
