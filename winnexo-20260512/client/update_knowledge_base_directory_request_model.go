// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeBaseDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetAppToken() *string
	SetDescription(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetDescription() *string
	SetDirectoryId(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetDirectoryId() *string
	SetName(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetName() *string
	SetOpsUserCode(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetOpsUserId() *string
	SetParentDirectoryId(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetParentDirectoryId() *string
	SetTenantId(v string) *UpdateKnowledgeBaseDirectoryRequest
	GetTenantId() *string
}

type UpdateKnowledgeBaseDirectoryRequest struct {
	// The platform application token issued by the console. If you specify this parameter, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The code of the proxy user, which is unique within the tenant. This parameter is used only when the application credential is configured as `source_uid`.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the proxy user, which is the WINNEXO platform user ID. This parameter is used only when the application credential is configured as `winnexo_uid`.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The ID of the parent directory.
	//
	// example:
	//
	// exampleParentDirectoryId
	ParentDirectoryId *string `json:"parentDirectoryId,omitempty" xml:"parentDirectoryId,omitempty"`
	// The ID of the tenant.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateKnowledgeBaseDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeBaseDirectoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *UpdateKnowledgeBaseDirectoryRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetAppToken(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.AppToken = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetDescription(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetDirectoryId(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.DirectoryId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetName(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetOpsUserCode(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.OpsUserCode = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetOpsUserId(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.OpsUserId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetParentDirectoryId(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) SetTenantId(v string) *UpdateKnowledgeBaseDirectoryRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateKnowledgeBaseDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
