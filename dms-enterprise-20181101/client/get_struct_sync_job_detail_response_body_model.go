// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStructSyncJobDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetStructSyncJobDetailResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetStructSyncJobDetailResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetStructSyncJobDetailResponseBody
	GetRequestId() *string
	SetStructSyncJobDetail(v *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) *GetStructSyncJobDetailResponseBody
	GetStructSyncJobDetail() *GetStructSyncJobDetailResponseBodyStructSyncJobDetail
	SetSuccess(v bool) *GetStructSyncJobDetailResponseBody
	GetSuccess() *bool
}

type GetStructSyncJobDetailResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 48602B78-0DDF-414C-8688-70CAB6070115
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The details of the schema synchronization task.
	StructSyncJobDetail *GetStructSyncJobDetailResponseBodyStructSyncJobDetail `json:"StructSyncJobDetail,omitempty" xml:"StructSyncJobDetail,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetStructSyncJobDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobDetailResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetStructSyncJobDetailResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetStructSyncJobDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStructSyncJobDetailResponseBody) GetStructSyncJobDetail() *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	return s.StructSyncJobDetail
}

func (s *GetStructSyncJobDetailResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetStructSyncJobDetailResponseBody) SetErrorCode(v string) *GetStructSyncJobDetailResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) SetErrorMessage(v string) *GetStructSyncJobDetailResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) SetRequestId(v string) *GetStructSyncJobDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) SetStructSyncJobDetail(v *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) *GetStructSyncJobDetailResponseBody {
	s.StructSyncJobDetail = v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) SetSuccess(v bool) *GetStructSyncJobDetailResponseBody {
	s.Success = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBody) Validate() error {
	if s.StructSyncJobDetail != nil {
		if err := s.StructSyncJobDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetStructSyncJobDetailResponseBodyStructSyncJobDetail struct {
	// The ID of the SQL task group.
	//
	// example:
	//
	// 12345
	DBTaskGroupId *int64 `json:"DBTaskGroupId,omitempty" xml:"DBTaskGroupId,omitempty"`
	// The number of SQL statements that have been executed.
	//
	// example:
	//
	// 1
	ExecuteCount *int64 `json:"ExecuteCount,omitempty" xml:"ExecuteCount,omitempty"`
	// The status of the task. Valid values:
	//
	// 	- **NEW**: The task was created.
	//
	// 	- **COMPARING**: The schemas of tables were being compared.
	//
	// 	- **COMPARE_BREAK**: The schema comparison was interrupted.
	//
	// 	- **COMPARE_FINISH**: The comparison was finished.
	//
	// 	- **NOT_SCRIPTS**: The comparison was finished but no executable script was available.
	//
	// 	- **SUBMITED_DBTASK**: The task was submitted.
	//
	// 	- **DBTASK_SUCCESS**: The task was complete.
	//
	// 	- **SUBMITED_WORKFLOW**: The ticket was submitted.
	//
	// 	- **WORKFLOW_SUCCESS**: The ticket was approved.
	//
	// example:
	//
	// DBTASK_SUCCESS
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The description of the task.
	//
	// example:
	//
	// test
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The type of security rule. Valid values:
	//
	// 	- **CANNOT_SYNC**: Synchronization cannot be performed.
	//
	// 	- **WITH_APPROVE**: The schema synchronization can be performed after the ticket is approved. You can call the [SubmitStructSyncOrderApproval](https://help.aliyun.com/document_detail/206166.html) operation to submit the ticket for approval.
	//
	// 	- **WITHOUT_APPROVE**: The schema synchronization can be performed without approval.
	//
	// example:
	//
	// WITHOUT_APPROVE
	SecurityRule *string `json:"SecurityRule,omitempty" xml:"SecurityRule,omitempty"`
	// The total number of SQL statements.
	//
	// example:
	//
	// 1
	SqlCount *int64 `json:"SqlCount,omitempty" xml:"SqlCount,omitempty"`
	// The number of tables that have been analyzed.
	//
	// example:
	//
	// 2
	TableAnalyzed *int64 `json:"TableAnalyzed,omitempty" xml:"TableAnalyzed,omitempty"`
	// The total number of tables.
	//
	// example:
	//
	// 2
	TableCount *int64 `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
}

func (s GetStructSyncJobDetailResponseBodyStructSyncJobDetail) String() string {
	return dara.Prettify(s)
}

func (s GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GoString() string {
	return s.String()
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GetDBTaskGroupId() *int64 {
	return s.DBTaskGroupId
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GetExecuteCount() *int64 {
	return s.ExecuteCount
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GetJobStatus() *string {
	return s.JobStatus
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GetMessage() *string {
	return s.Message
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GetSecurityRule() *string {
	return s.SecurityRule
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GetSqlCount() *int64 {
	return s.SqlCount
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GetTableAnalyzed() *int64 {
	return s.TableAnalyzed
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) GetTableCount() *int64 {
	return s.TableCount
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetDBTaskGroupId(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.DBTaskGroupId = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetExecuteCount(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.ExecuteCount = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetJobStatus(v string) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.JobStatus = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetMessage(v string) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.Message = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetSecurityRule(v string) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.SecurityRule = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetSqlCount(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.SqlCount = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetTableAnalyzed(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.TableAnalyzed = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) SetTableCount(v int64) *GetStructSyncJobDetailResponseBodyStructSyncJobDetail {
	s.TableCount = &v
	return s
}

func (s *GetStructSyncJobDetailResponseBodyStructSyncJobDetail) Validate() error {
	return dara.Validate(s)
}
