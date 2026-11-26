// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateKnowledgeBaseDirectoryRequest
	GetAppToken() *string
	SetDescription(v string) *CreateKnowledgeBaseDirectoryRequest
	GetDescription() *string
	SetName(v string) *CreateKnowledgeBaseDirectoryRequest
	GetName() *string
	SetOpsUserCode(v string) *CreateKnowledgeBaseDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateKnowledgeBaseDirectoryRequest
	GetOpsUserId() *string
	SetParentDirectoryId(v string) *CreateKnowledgeBaseDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *CreateKnowledgeBaseDirectoryRequest
	GetTenantId() *string
}

type CreateKnowledgeBaseDirectoryRequest struct {
	// The platform application token issued by the console. If specified, the operation is executed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the AI assistant.
	//
	// example:
	//
	// recorder function
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The image name.
	//
	// This parameter is required.
	//
	// example:
	//
	// oklabs_tongyici
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
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
	// The directory ID.
	//
	// example:
	//
	// wd-lxykjnnw4lyl9eq
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 692318833855074
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBaseDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreateKnowledgeBaseDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetAppToken(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.AppToken = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetDescription(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetName(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetOpsUserCode(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetOpsUserId(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetParentDirectoryId(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) SetTenantId(v string) *CreateKnowledgeBaseDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBaseDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
