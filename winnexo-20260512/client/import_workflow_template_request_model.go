// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *ImportWorkflowTemplateRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *ImportWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetOpsUserCode(v string) *ImportWorkflowTemplateRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *ImportWorkflowTemplateRequest
	GetOpsUserId() *string
	SetTemplate(v map[string]interface{}) *ImportWorkflowTemplateRequest
	GetTemplate() map[string]interface{}
	SetTenantId(v string) *ImportWorkflowTemplateRequest
	GetTenantId() *string
}

type ImportWorkflowTemplateRequest struct {
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
	Template map[string]interface{} `json:"template,omitempty" xml:"template,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s ImportWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *ImportWorkflowTemplateRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *ImportWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *ImportWorkflowTemplateRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *ImportWorkflowTemplateRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *ImportWorkflowTemplateRequest) GetTemplate() map[string]interface{} {
	return s.Template
}

func (s *ImportWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *ImportWorkflowTemplateRequest) SetAppToken(v string) *ImportWorkflowTemplateRequest {
	s.AppToken = &v
	return s
}

func (s *ImportWorkflowTemplateRequest) SetCollaborationGroupId(v string) *ImportWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *ImportWorkflowTemplateRequest) SetOpsUserCode(v string) *ImportWorkflowTemplateRequest {
	s.OpsUserCode = &v
	return s
}

func (s *ImportWorkflowTemplateRequest) SetOpsUserId(v string) *ImportWorkflowTemplateRequest {
	s.OpsUserId = &v
	return s
}

func (s *ImportWorkflowTemplateRequest) SetTemplate(v map[string]interface{}) *ImportWorkflowTemplateRequest {
	s.Template = v
	return s
}

func (s *ImportWorkflowTemplateRequest) SetTenantId(v string) *ImportWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *ImportWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
