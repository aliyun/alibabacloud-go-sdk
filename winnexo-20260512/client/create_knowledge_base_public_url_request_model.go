// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBasePublicUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateKnowledgeBasePublicUrlRequest
	GetAppToken() *string
	SetDescription(v string) *CreateKnowledgeBasePublicUrlRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBasePublicUrlRequest
	GetDirectoryId() *string
	SetName(v string) *CreateKnowledgeBasePublicUrlRequest
	GetName() *string
	SetNotes(v string) *CreateKnowledgeBasePublicUrlRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreateKnowledgeBasePublicUrlRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *CreateKnowledgeBasePublicUrlRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateKnowledgeBasePublicUrlRequest
	GetOpsUserId() *string
	SetOriginalUrl(v string) *CreateKnowledgeBasePublicUrlRequest
	GetOriginalUrl() *string
	SetSourceTags(v string) *CreateKnowledgeBasePublicUrlRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateKnowledgeBasePublicUrlRequest
	GetTenantId() *string
}

type CreateKnowledgeBasePublicUrlRequest struct {
	// The platform application token issued by the console. If specified, the operation is performed as the user specified by opsUserId or opsUserCode.
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the resource.
	//
	// example:
	//
	// Project design document
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the target folder in the enterprise knowledge base. This parameter is required, and you must have knowledge base management permissions on the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// dir_tenant_child
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name of the resource. If not specified, the URL is used.
	//
	// example:
	//
	// Project plan
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The analysis instructions.
	//
	// example:
	//
	// Extract decisions and action items
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the operating object.
	//
	// example:
	//
	// R&D assistant
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The URL of the public HTTP or HTTPS web page.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com
	OriginalUrl *string `json:"originalUrl,omitempty" xml:"originalUrl,omitempty"`
	// The list of resource tags in JSON string format.
	//
	// example:
	//
	// ["R&D"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBasePublicUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBasePublicUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetOriginalUrl() *string {
	return s.OriginalUrl
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBasePublicUrlRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetAppToken(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.AppToken = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetDescription(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetDirectoryId(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetName(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetNotes(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.Notes = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetOperatingObjectName(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetOpsUserCode(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetOpsUserId(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetOriginalUrl(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.OriginalUrl = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetSourceTags(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) SetTenantId(v string) *CreateKnowledgeBasePublicUrlRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBasePublicUrlRequest) Validate() error {
	return dara.Validate(s)
}
