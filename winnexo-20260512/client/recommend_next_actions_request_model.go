// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecommendNextActionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *RecommendNextActionsRequest
	GetAppToken() *string
	SetCustomPrompt(v string) *RecommendNextActionsRequest
	GetCustomPrompt() *string
	SetOpsUserCode(v string) *RecommendNextActionsRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *RecommendNextActionsRequest
	GetOpsUserId() *string
	SetOutputType(v string) *RecommendNextActionsRequest
	GetOutputType() *string
	SetRecentMessageCount(v int64) *RecommendNextActionsRequest
	GetRecentMessageCount() *int64
	SetSessionId(v string) *RecommendNextActionsRequest
	GetSessionId() *string
	SetTenantId(v string) *RecommendNextActionsRequest
	GetTenantId() *string
}

type RecommendNextActionsRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The extraction instruction.
	//
	// example:
	//
	// Only recommend next steps related to data analysis
	CustomPrompt *string `json:"customPrompt,omitempty" xml:"customPrompt,omitempty"`
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
	// The output type: `conversation/skill/task`.
	//
	// example:
	//
	// followUpOnly
	OutputType *string `json:"outputType,omitempty" xml:"outputType,omitempty"`
	// The number of recent messages used to assemble the contextual information.
	//
	// example:
	//
	// 10
	RecentMessageCount *int64 `json:"recentMessageCount,omitempty" xml:"recentMessageCount,omitempty"`
	// The session ID to filter by. If specified, returns all Active/Expired status information associated with this session.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSessionId
	SessionId *string `json:"sessionId,omitempty" xml:"sessionId,omitempty"`
	// The tenant ID. This is a common parameter. In winnexo-cli, pass this value explicitly by using --tenant-id.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s RecommendNextActionsRequest) String() string {
	return dara.Prettify(s)
}

func (s RecommendNextActionsRequest) GoString() string {
	return s.String()
}

func (s *RecommendNextActionsRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *RecommendNextActionsRequest) GetCustomPrompt() *string {
	return s.CustomPrompt
}

func (s *RecommendNextActionsRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *RecommendNextActionsRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *RecommendNextActionsRequest) GetOutputType() *string {
	return s.OutputType
}

func (s *RecommendNextActionsRequest) GetRecentMessageCount() *int64 {
	return s.RecentMessageCount
}

func (s *RecommendNextActionsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RecommendNextActionsRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *RecommendNextActionsRequest) SetAppToken(v string) *RecommendNextActionsRequest {
	s.AppToken = &v
	return s
}

func (s *RecommendNextActionsRequest) SetCustomPrompt(v string) *RecommendNextActionsRequest {
	s.CustomPrompt = &v
	return s
}

func (s *RecommendNextActionsRequest) SetOpsUserCode(v string) *RecommendNextActionsRequest {
	s.OpsUserCode = &v
	return s
}

func (s *RecommendNextActionsRequest) SetOpsUserId(v string) *RecommendNextActionsRequest {
	s.OpsUserId = &v
	return s
}

func (s *RecommendNextActionsRequest) SetOutputType(v string) *RecommendNextActionsRequest {
	s.OutputType = &v
	return s
}

func (s *RecommendNextActionsRequest) SetRecentMessageCount(v int64) *RecommendNextActionsRequest {
	s.RecentMessageCount = &v
	return s
}

func (s *RecommendNextActionsRequest) SetSessionId(v string) *RecommendNextActionsRequest {
	s.SessionId = &v
	return s
}

func (s *RecommendNextActionsRequest) SetTenantId(v string) *RecommendNextActionsRequest {
	s.TenantId = &v
	return s
}

func (s *RecommendNextActionsRequest) Validate() error {
	return dara.Validate(s)
}
