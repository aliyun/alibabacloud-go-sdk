// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveKnowledgeBaseResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *MoveKnowledgeBaseResourceRequest
	GetAppToken() *string
	SetKnowledgeId(v string) *MoveKnowledgeBaseResourceRequest
	GetKnowledgeId() *string
	SetOpsUserCode(v string) *MoveKnowledgeBaseResourceRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *MoveKnowledgeBaseResourceRequest
	GetOpsUserId() *string
	SetSourceDirectoryId(v string) *MoveKnowledgeBaseResourceRequest
	GetSourceDirectoryId() *string
	SetSourceId(v string) *MoveKnowledgeBaseResourceRequest
	GetSourceId() *string
	SetTargetDirectoryId(v string) *MoveKnowledgeBaseResourceRequest
	GetTargetDirectoryId() *string
	SetTenantId(v string) *MoveKnowledgeBaseResourceRequest
	GetTenantId() *string
}

type MoveKnowledgeBaseResourceRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// Not supported. This parameter is ignored.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleKnowledgeId
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
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
	// The source directory ID. This is the enterprise knowledge base directory where the resource currently resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleSourceDirectoryId
	SourceDirectoryId *string `json:"sourceDirectoryId,omitempty" xml:"sourceDirectoryId,omitempty"`
	// The data source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2000627
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The target directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleTargetDirectoryId
	TargetDirectoryId *string `json:"targetDirectoryId,omitempty" xml:"targetDirectoryId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 1729094555111072
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s MoveKnowledgeBaseResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveKnowledgeBaseResourceRequest) GoString() string {
	return s.String()
}

func (s *MoveKnowledgeBaseResourceRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *MoveKnowledgeBaseResourceRequest) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *MoveKnowledgeBaseResourceRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *MoveKnowledgeBaseResourceRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *MoveKnowledgeBaseResourceRequest) GetSourceDirectoryId() *string {
	return s.SourceDirectoryId
}

func (s *MoveKnowledgeBaseResourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *MoveKnowledgeBaseResourceRequest) GetTargetDirectoryId() *string {
	return s.TargetDirectoryId
}

func (s *MoveKnowledgeBaseResourceRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *MoveKnowledgeBaseResourceRequest) SetAppToken(v string) *MoveKnowledgeBaseResourceRequest {
	s.AppToken = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetKnowledgeId(v string) *MoveKnowledgeBaseResourceRequest {
	s.KnowledgeId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetOpsUserCode(v string) *MoveKnowledgeBaseResourceRequest {
	s.OpsUserCode = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetOpsUserId(v string) *MoveKnowledgeBaseResourceRequest {
	s.OpsUserId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetSourceDirectoryId(v string) *MoveKnowledgeBaseResourceRequest {
	s.SourceDirectoryId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetSourceId(v string) *MoveKnowledgeBaseResourceRequest {
	s.SourceId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetTargetDirectoryId(v string) *MoveKnowledgeBaseResourceRequest {
	s.TargetDirectoryId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) SetTenantId(v string) *MoveKnowledgeBaseResourceRequest {
	s.TenantId = &v
	return s
}

func (s *MoveKnowledgeBaseResourceRequest) Validate() error {
	return dara.Validate(s)
}
