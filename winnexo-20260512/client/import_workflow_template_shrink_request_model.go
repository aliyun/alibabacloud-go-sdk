// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowTemplateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ImportWorkflowTemplateShrinkRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ImportWorkflowTemplateShrinkRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *ImportWorkflowTemplateShrinkRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ImportWorkflowTemplateShrinkRequest
	GetOpsUserId() *string
	SetTemplateShrink(v string) *ImportWorkflowTemplateShrinkRequest
	GetTemplateShrink() *string
	SetTenantId(v string) *ImportWorkflowTemplateShrinkRequest
	GetTenantId() *string
}

type ImportWorkflowTemplateShrinkRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user identified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The ID of the collaboration workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The proxy user code, which is unique within the tenant. This parameter is used only when the application credential is configured as source_uid.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The proxy user ID, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as winnexo_uid.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The portable template in JSON format, which includes schemaVersion, steps, and edges. Each step must specify stepId and sortOrder.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"name":"Customer complaint handling process","schemaVersion":2,"steps":[{"stepId":"collect","sortOrder":1,"name":"Collect customer information","specification":[{"type":"text","content":"Collect customer information"}],"executorType":"MEMBER","executorUserId":10086}],"edges":[]}
	TemplateShrink *string `json:"template,omitempty" xml:"template,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ImportWorkflowTemplateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowTemplateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ImportWorkflowTemplateShrinkRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ImportWorkflowTemplateShrinkRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ImportWorkflowTemplateShrinkRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ImportWorkflowTemplateShrinkRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ImportWorkflowTemplateShrinkRequest) GetTemplateShrink() *string {
	return s.TemplateShrink
}

func (s *ImportWorkflowTemplateShrinkRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ImportWorkflowTemplateShrinkRequest) SetAppToken(v string) *ImportWorkflowTemplateShrinkRequest {
	s.AppToken = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) SetCollaborationGroupId(v string) *ImportWorkflowTemplateShrinkRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) SetOpsUserCode(v string) *ImportWorkflowTemplateShrinkRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) SetOpsUserId(v string) *ImportWorkflowTemplateShrinkRequest {
	s.OpsUserId = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) SetTemplateShrink(v string) *ImportWorkflowTemplateShrinkRequest {
	s.TemplateShrink = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) SetTenantId(v string) *ImportWorkflowTemplateShrinkRequest {
	s.TenantId = &v
	return s
}

func (s *ImportWorkflowTemplateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
