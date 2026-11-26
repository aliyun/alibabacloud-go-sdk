// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeBaseVoiceMeetingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetAppToken() *string
	SetDescription(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetDirectoryId() *string
	SetFileUrl(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetFileUrl() *string
	SetName(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetName() *string
	SetOperatingObjectName(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetOpsUserId() *string
	SetSourceTags(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreateKnowledgeBaseVoiceMeetingRequest
	GetTenantId() *string
}

type CreateKnowledgeBaseVoiceMeetingRequest struct {
	// token
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The description of the AI assistant.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The attachment URL in Yida.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// The name of the AI assistant.
	//
	// This parameter is required.
	//
	// example:
	//
	// Project Review Recording
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The name of the digital employee. This is the name of the operating object and is optional.
	//
	// example:
	//
	// exampleAgent
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
	// The code of the operator.
	//
	// example:
	//
	// alice
	OpsUserCode *string `json:"opsUserCode,omitempty" xml:"opsUserCode,omitempty"`
	// The ID of the operator.
	//
	// example:
	//
	// 10001
	OpsUserId *string `json:"opsUserId,omitempty" xml:"opsUserId,omitempty"`
	// The resource tags. This parameter is optional and accepts a JSON string list, such as ["tagA","tagB"].
	//
	// example:
	//
	// ["Important","Meeting"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID. This is a common parameter. If not specified, the default tenant of the caller is used.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateKnowledgeBaseVoiceMeetingRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeBaseVoiceMeetingRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetAppToken(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.AppToken = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetDescription(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetDirectoryId(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetFileUrl(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetName(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetOperatingObjectName(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetOpsUserCode(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetOpsUserId(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetSourceTags(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.SourceTags = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) SetTenantId(v string) *CreateKnowledgeBaseVoiceMeetingRequest {
	s.TenantId = &v
	return s
}

func (s *CreateKnowledgeBaseVoiceMeetingRequest) Validate() error {
	return dara.Validate(s)
}
