// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateWorkflowTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAction(v string) *UpdateWorkflowTemplateRequest
	GetAction() *string
	SetAppToken(v string) *UpdateWorkflowTemplateRequest
	GetAppToken() *string
	SetCollaborationGroupId(v string) *UpdateWorkflowTemplateRequest
	GetCollaborationGroupId() *string
	SetDefinition(v map[string]interface{}) *UpdateWorkflowTemplateRequest
	GetDefinition() map[string]interface{}
	SetOpsUserCode(v string) *UpdateWorkflowTemplateRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateWorkflowTemplateRequest
	GetOpsUserId() *string
	SetSourceVersion(v string) *UpdateWorkflowTemplateRequest
	GetSourceVersion() *string
	SetTemplateId(v string) *UpdateWorkflowTemplateRequest
	GetTemplateId() *string
	SetTenantId(v string) *UpdateWorkflowTemplateRequest
	GetTenantId() *string
	SetVersion(v string) *UpdateWorkflowTemplateRequest
	GetVersion() *string
	SetVersionDescription(v string) *UpdateWorkflowTemplateRequest
	GetVersionDescription() *string
}

type UpdateWorkflowTemplateRequest struct {
	// The write action. Valid values: SAVE (saves the template as a draft without taking effect, and allows node names and executors to be left empty) and SUBMIT (requires complete nodes, and saves and submits the template as the only effective version). Default value: SUBMIT.
	//
	// example:
	//
	// SUBMIT
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The collaboration workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cg_9401
	CollaborationGroupId *string `json:"collaborationGroupId,omitempty" xml:"collaborationGroupId,omitempty"`
	// The complete JSON definition of the template for full overwriting.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"name":"Customer Complaint Handling Process","schemaVersion":2,"steps":[{"stepId":"collect","sortOrder":1,"name":"Collect Customer Information","specification":[{"type":"text","content":"Collect Customer Information"}],"executorType":"MEMBER","executorUserId":10086}],"edges":[]}
	Definition map[string]interface{} `json:"definition,omitempty" xml:"definition,omitempty"`
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
	// The submitted version number on which the new version is based. This parameter is mutually exclusive with version. If not specified, the current effective version is used as the fallback.
	//
	// example:
	//
	// v1
	SourceVersion *string `json:"sourceVersion,omitempty" xml:"sourceVersion,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl_9401
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The version number to which the returned content belongs.
	//
	// example:
	//
	// v2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// The change description for this version. An empty value indicates that no description is provided.
	//
	// example:
	//
	// Added review step
	VersionDescription *string `json:"versionDescription,omitempty" xml:"versionDescription,omitempty"`
}

func (s UpdateWorkflowTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateWorkflowTemplateRequest) GoString() string {
	return s.String()
}

func (s *UpdateWorkflowTemplateRequest) GetAction() *string {
	return s.Action
}

func (s *UpdateWorkflowTemplateRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateWorkflowTemplateRequest) GetCollaborationGroupId() *string {
	return s.CollaborationGroupId
}

func (s *UpdateWorkflowTemplateRequest) GetDefinition() map[string]interface{} {
	return s.Definition
}

func (s *UpdateWorkflowTemplateRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateWorkflowTemplateRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateWorkflowTemplateRequest) GetSourceVersion() *string {
	return s.SourceVersion
}

func (s *UpdateWorkflowTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UpdateWorkflowTemplateRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateWorkflowTemplateRequest) GetVersion() *string {
	return s.Version
}

func (s *UpdateWorkflowTemplateRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateWorkflowTemplateRequest) SetAction(v string) *UpdateWorkflowTemplateRequest {
	s.Action = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetAppToken(v string) *UpdateWorkflowTemplateRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetCollaborationGroupId(v string) *UpdateWorkflowTemplateRequest {
	s.CollaborationGroupId = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetDefinition(v map[string]interface{}) *UpdateWorkflowTemplateRequest {
	s.Definition = v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetOpsUserCode(v string) *UpdateWorkflowTemplateRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetOpsUserId(v string) *UpdateWorkflowTemplateRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetSourceVersion(v string) *UpdateWorkflowTemplateRequest {
	s.SourceVersion = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetTemplateId(v string) *UpdateWorkflowTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetTenantId(v string) *UpdateWorkflowTemplateRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetVersion(v string) *UpdateWorkflowTemplateRequest {
	s.Version = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) SetVersionDescription(v string) *UpdateWorkflowTemplateRequest {
	s.VersionDescription = &v
	return s
}

func (s *UpdateWorkflowTemplateRequest) Validate() error {
	return dara.Validate(s)
}
