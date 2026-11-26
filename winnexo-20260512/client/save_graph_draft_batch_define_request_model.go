// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveGraphDraftBatchDefineRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *SaveGraphDraftBatchDefineRequest
	GetAppToken() *string
	SetDraftChangeIds(v []*int64) *SaveGraphDraftBatchDefineRequest
	GetDraftChangeIds() []*int64
	SetGraphName(v string) *SaveGraphDraftBatchDefineRequest
	GetGraphName() *string
	SetOpsUserCode(v string) *SaveGraphDraftBatchDefineRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *SaveGraphDraftBatchDefineRequest
	GetOpsUserId() *string
	SetSaveMode(v string) *SaveGraphDraftBatchDefineRequest
	GetSaveMode() *string
	SetTenantId(v string) *SaveGraphDraftBatchDefineRequest
	GetTenantId() *string
	SetYamlEdit(v string) *SaveGraphDraftBatchDefineRequest
	GetYamlEdit() *string
}

type SaveGraphDraftBatchDefineRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The list of draft change IDs.
	//
	// example:
	//
	// [401001, 401002]
	DraftChangeIds []*int64 `json:"draftChangeIds,omitempty" xml:"draftChangeIds,omitempty" type:"Repeated"`
	// The graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// crm_graph
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
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
	// The save mode.
	//
	// example:
	//
	// FULL_YAML
	SaveMode *string `json:"saveMode,omitempty" xml:"saveMode,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The raw YAML text of the graph schema trimmed by READ permissions, with $ref references retained within the authorized subgraph.
	//
	// This parameter is required.
	//
	// example:
	//
	// name: crm_graph
	YamlEdit *string `json:"yamlEdit,omitempty" xml:"yamlEdit,omitempty"`
}

func (s SaveGraphDraftBatchDefineRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveGraphDraftBatchDefineRequest) GoString() string {
	return s.String()
}

func (s *SaveGraphDraftBatchDefineRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *SaveGraphDraftBatchDefineRequest) GetDraftChangeIds() []*int64 {
	return s.DraftChangeIds
}

func (s *SaveGraphDraftBatchDefineRequest) GetGraphName() *string {
	return s.GraphName
}

func (s *SaveGraphDraftBatchDefineRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *SaveGraphDraftBatchDefineRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *SaveGraphDraftBatchDefineRequest) GetSaveMode() *string {
	return s.SaveMode
}

func (s *SaveGraphDraftBatchDefineRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *SaveGraphDraftBatchDefineRequest) GetYamlEdit() *string {
	return s.YamlEdit
}

func (s *SaveGraphDraftBatchDefineRequest) SetAppToken(v string) *SaveGraphDraftBatchDefineRequest {
	s.AppToken = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetDraftChangeIds(v []*int64) *SaveGraphDraftBatchDefineRequest {
	s.DraftChangeIds = v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetGraphName(v string) *SaveGraphDraftBatchDefineRequest {
	s.GraphName = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetOpsUserCode(v string) *SaveGraphDraftBatchDefineRequest {
	s.OpsUserCode = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetOpsUserId(v string) *SaveGraphDraftBatchDefineRequest {
	s.OpsUserId = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetSaveMode(v string) *SaveGraphDraftBatchDefineRequest {
	s.SaveMode = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetTenantId(v string) *SaveGraphDraftBatchDefineRequest {
	s.TenantId = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) SetYamlEdit(v string) *SaveGraphDraftBatchDefineRequest {
	s.YamlEdit = &v
	return s
}

func (s *SaveGraphDraftBatchDefineRequest) Validate() error {
	return dara.Validate(s)
}
