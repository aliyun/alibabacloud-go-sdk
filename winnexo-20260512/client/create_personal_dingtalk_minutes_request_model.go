// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePersonalDingtalkMinutesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppToken(v string) *CreatePersonalDingtalkMinutesRequest
	GetAppToken() *string
	SetDescription(v string) *CreatePersonalDingtalkMinutesRequest
	GetDescription() *string
	SetDirectoryId(v string) *CreatePersonalDingtalkMinutesRequest
	GetDirectoryId() *string
	SetName(v string) *CreatePersonalDingtalkMinutesRequest
	GetName() *string
	SetNotes(v string) *CreatePersonalDingtalkMinutesRequest
	GetNotes() *string
	SetOperatingObjectName(v string) *CreatePersonalDingtalkMinutesRequest
	GetOperatingObjectName() *string
	SetOpsUserCode(v string) *CreatePersonalDingtalkMinutesRequest
	GetOpsUserCode() *string
	SetOpsUserId(v string) *CreatePersonalDingtalkMinutesRequest
	GetOpsUserId() *string
	SetShanjiUrl(v string) *CreatePersonalDingtalkMinutesRequest
	GetShanjiUrl() *string
	SetSourceTags(v string) *CreatePersonalDingtalkMinutesRequest
	GetSourceTags() *string
	SetTenantId(v string) *CreatePersonalDingtalkMinutesRequest
	GetTenantId() *string
}

type CreatePersonalDingtalkMinutesRequest struct {
	// 平台应用 Token（控制台签发）。传入后将以 opsUserId 或 opsUserCode 指定的用户身份执行
	//
	// example:
	//
	// wn_app_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
	AppToken *string `json:"appToken,omitempty" xml:"appToken,omitempty"`
	// The pipeline description.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// exampleDirectoryId
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The name of the worksheet.
	//
	// This parameter is required.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The meeting notes content (optional). The notes are used for auxiliary analysis.
	//
	// example:
	//
	// string_value
	Notes *string `json:"notes,omitempty" xml:"notes,omitempty"`
	// The name of the digital employee (operating object name, optional).
	//
	// example:
	//
	// string_value
	OperatingObjectName *string `json:"operatingObjectName,omitempty" xml:"operatingObjectName,omitempty"`
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
	// The original Shanji note link (required).
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	ShanjiUrl *string `json:"shanjiUrl,omitempty" xml:"shanjiUrl,omitempty"`
	// A list of resource tag JSON strings.
	//
	// example:
	//
	// ["Important","Shanji"]
	SourceTags *string `json:"sourceTags,omitempty" xml:"sourceTags,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 692318833855074
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreatePersonalDingtalkMinutesRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePersonalDingtalkMinutesRequest) GoString() string {
	return s.String()
}

func (s *CreatePersonalDingtalkMinutesRequest) GetAppToken() *string {
	return s.AppToken
}

func (s *CreatePersonalDingtalkMinutesRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePersonalDingtalkMinutesRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreatePersonalDingtalkMinutesRequest) GetName() *string {
	return s.Name
}

func (s *CreatePersonalDingtalkMinutesRequest) GetNotes() *string {
	return s.Notes
}

func (s *CreatePersonalDingtalkMinutesRequest) GetOperatingObjectName() *string {
	return s.OperatingObjectName
}

func (s *CreatePersonalDingtalkMinutesRequest) GetOpsUserCode() *string {
	return s.OpsUserCode
}

func (s *CreatePersonalDingtalkMinutesRequest) GetOpsUserId() *string {
	return s.OpsUserId
}

func (s *CreatePersonalDingtalkMinutesRequest) GetShanjiUrl() *string {
	return s.ShanjiUrl
}

func (s *CreatePersonalDingtalkMinutesRequest) GetSourceTags() *string {
	return s.SourceTags
}

func (s *CreatePersonalDingtalkMinutesRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreatePersonalDingtalkMinutesRequest) SetAppToken(v string) *CreatePersonalDingtalkMinutesRequest {
	s.AppToken = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetDescription(v string) *CreatePersonalDingtalkMinutesRequest {
	s.Description = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetDirectoryId(v string) *CreatePersonalDingtalkMinutesRequest {
	s.DirectoryId = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetName(v string) *CreatePersonalDingtalkMinutesRequest {
	s.Name = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetNotes(v string) *CreatePersonalDingtalkMinutesRequest {
	s.Notes = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetOperatingObjectName(v string) *CreatePersonalDingtalkMinutesRequest {
	s.OperatingObjectName = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetOpsUserCode(v string) *CreatePersonalDingtalkMinutesRequest {
	s.OpsUserCode = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetOpsUserId(v string) *CreatePersonalDingtalkMinutesRequest {
	s.OpsUserId = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetShanjiUrl(v string) *CreatePersonalDingtalkMinutesRequest {
	s.ShanjiUrl = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetSourceTags(v string) *CreatePersonalDingtalkMinutesRequest {
	s.SourceTags = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) SetTenantId(v string) *CreatePersonalDingtalkMinutesRequest {
	s.TenantId = &v
	return s
}

func (s *CreatePersonalDingtalkMinutesRequest) Validate() error {
	return dara.Validate(s)
}
