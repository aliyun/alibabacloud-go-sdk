// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledTaskConversationRecordsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *GetScheduledTaskConversationRecordsRequest
	GetAppToken() *string
	SetInitiatorUserId(v string) *GetScheduledTaskConversationRecordsRequest
	GetInitiatorUserId() *string
	SetOpsUserCode(v string) *GetScheduledTaskConversationRecordsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *GetScheduledTaskConversationRecordsRequest
	GetOpsUserId() *string
	SetPage(v int32) *GetScheduledTaskConversationRecordsRequest
	GetPage() *int32
	SetPageSize(v int32) *GetScheduledTaskConversationRecordsRequest
	GetPageSize() *int32
	SetStatus(v string) *GetScheduledTaskConversationRecordsRequest
	GetStatus() *string
	SetTaskId(v string) *GetScheduledTaskConversationRecordsRequest
	GetTaskId() *string
	SetTenantId(v string) *GetScheduledTaskConversationRecordsRequest
	GetTenantId() *string
}

type GetScheduledTaskConversationRecordsRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The user ID of the task initiator.
	//
	// example:
	//
	// 5
	InitiatorUserId *string `json:"initiatorUserId,omitempty" xml:"initiatorUserId,omitempty"`
	// 代理用户编码（租户内唯一），仅当应用凭证配置为 source_uid 时使用
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// 代理用户ID（WINNEXO 平台用户ID），仅当应用凭证配置为 winnexo_uid 时使用
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The maximum number of data records to return in this request.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The running status of the task.
	//
	// example:
	//
	// failed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The execution status filter in lowercase. Valid values: pending (queued), running (running), success (successful), failed (failed), timeout (timed out), and cancelled (cancelled). If you do not specify this parameter, no filtering is applied. If you specify this parameter, future scheduled items are not generated.
	//
	// example:
	//
	// exampleTaskId
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The tenant ID. This is a common parameter. If you do not specify this parameter, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetScheduledTaskConversationRecordsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledTaskConversationRecordsRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledTaskConversationRecordsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *GetScheduledTaskConversationRecordsRequest) GetInitiatorUserId() *string {
	return s.InitiatorUserId
}

func (s *GetScheduledTaskConversationRecordsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *GetScheduledTaskConversationRecordsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *GetScheduledTaskConversationRecordsRequest) GetPage() *int32 {
	return s.Page
}

func (s *GetScheduledTaskConversationRecordsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetScheduledTaskConversationRecordsRequest) GetStatus() *string {
	return s.Status
}

func (s *GetScheduledTaskConversationRecordsRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *GetScheduledTaskConversationRecordsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *GetScheduledTaskConversationRecordsRequest) SetAppToken(v string) *GetScheduledTaskConversationRecordsRequest {
	s.AppToken = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) SetInitiatorUserId(v string) *GetScheduledTaskConversationRecordsRequest {
	s.InitiatorUserId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) SetOpsUserCode(v string) *GetScheduledTaskConversationRecordsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) SetOpsUserId(v string) *GetScheduledTaskConversationRecordsRequest {
	s.OpsUserId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) SetPage(v int32) *GetScheduledTaskConversationRecordsRequest {
	s.Page = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) SetPageSize(v int32) *GetScheduledTaskConversationRecordsRequest {
	s.PageSize = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) SetStatus(v string) *GetScheduledTaskConversationRecordsRequest {
	s.Status = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) SetTaskId(v string) *GetScheduledTaskConversationRecordsRequest {
	s.TaskId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) SetTenantId(v string) *GetScheduledTaskConversationRecordsRequest {
	s.TenantId = &v
	return s
}

func (s *GetScheduledTaskConversationRecordsRequest) Validate() error {
	return dara.Validate(s)
}
