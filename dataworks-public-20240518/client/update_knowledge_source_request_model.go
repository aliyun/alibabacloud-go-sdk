// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateKnowledgeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSource(v *UpdateKnowledgeSourceRequestDataSource) *UpdateKnowledgeSourceRequest
	GetDataSource() *UpdateKnowledgeSourceRequestDataSource
	SetDescription(v string) *UpdateKnowledgeSourceRequest
	GetDescription() *string
	SetDocument(v *UpdateKnowledgeSourceRequestDocument) *UpdateKnowledgeSourceRequest
	GetDocument() *UpdateKnowledgeSourceRequestDocument
	SetKnowledgeBaseId(v string) *UpdateKnowledgeSourceRequest
	GetKnowledgeBaseId() *string
	SetName(v string) *UpdateKnowledgeSourceRequest
	GetName() *string
	SetProject(v *UpdateKnowledgeSourceRequestProject) *UpdateKnowledgeSourceRequest
	GetProject() *UpdateKnowledgeSourceRequestProject
	SetSourceId(v string) *UpdateKnowledgeSourceRequest
	GetSourceId() *string
}

type UpdateKnowledgeSourceRequest struct {
	// The complete configuration for Type=DATA_SOURCE.
	DataSource *UpdateKnowledgeSourceRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The new description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The complete configuration for Type=DOCUMENT. Credentials are not returned after they are written.
	Document *UpdateKnowledgeSourceRequestDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The ID of the knowledge base to which the source belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The new name.
	//
	// example:
	//
	// Sales Project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The complete configuration for Type=PROJECT.
	Project *UpdateKnowledgeSourceRequestProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The source ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
}

func (s UpdateKnowledgeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequest) GetDataSource() *UpdateKnowledgeSourceRequestDataSource {
	return s.DataSource
}

func (s *UpdateKnowledgeSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateKnowledgeSourceRequest) GetDocument() *UpdateKnowledgeSourceRequestDocument {
	return s.Document
}

func (s *UpdateKnowledgeSourceRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *UpdateKnowledgeSourceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateKnowledgeSourceRequest) GetProject() *UpdateKnowledgeSourceRequestProject {
	return s.Project
}

func (s *UpdateKnowledgeSourceRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *UpdateKnowledgeSourceRequest) SetDataSource(v *UpdateKnowledgeSourceRequestDataSource) *UpdateKnowledgeSourceRequest {
	s.DataSource = v
	return s
}

func (s *UpdateKnowledgeSourceRequest) SetDescription(v string) *UpdateKnowledgeSourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateKnowledgeSourceRequest) SetDocument(v *UpdateKnowledgeSourceRequestDocument) *UpdateKnowledgeSourceRequest {
	s.Document = v
	return s
}

func (s *UpdateKnowledgeSourceRequest) SetKnowledgeBaseId(v string) *UpdateKnowledgeSourceRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *UpdateKnowledgeSourceRequest) SetName(v string) *UpdateKnowledgeSourceRequest {
	s.Name = &v
	return s
}

func (s *UpdateKnowledgeSourceRequest) SetProject(v *UpdateKnowledgeSourceRequestProject) *UpdateKnowledgeSourceRequest {
	s.Project = v
	return s
}

func (s *UpdateKnowledgeSourceRequest) SetSourceId(v string) *UpdateKnowledgeSourceRequest {
	s.SourceId = &v
	return s
}

func (s *UpdateKnowledgeSourceRequest) Validate() error {
	if s.DataSource != nil {
		if err := s.DataSource.Validate(); err != nil {
			return err
		}
	}
	if s.Document != nil {
		if err := s.Document.Validate(); err != nil {
			return err
		}
	}
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeSourceRequestDataSource struct {
	// The DataWorks data source ID. The server uses this ID to resolve the associated workspace, environment, and data source type.
	//
	// example:
	//
	// 12345
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s UpdateKnowledgeSourceRequestDataSource) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestDataSource) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestDataSource) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *UpdateKnowledgeSourceRequestDataSource) SetDataSourceId(v int64) *UpdateKnowledgeSourceRequestDataSource {
	s.DataSourceId = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDataSource) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeSourceRequestDocument struct {
	// The connector type. Valid values: DINGTALK, FEISHU, LARK, WECOM, and YUQUE. This value cannot be changed after creation.
	//
	// example:
	//
	// WECOM
	ConnectorType *string `json:"ConnectorType,omitempty" xml:"ConnectorType,omitempty"`
	// The DingTalk configuration. This parameter is present only when the connector type is DINGTALK.
	DingTalk *UpdateKnowledgeSourceRequestDocumentDingTalk `json:"DingTalk,omitempty" xml:"DingTalk,omitempty" type:"Struct"`
	// The Feishu configuration. This parameter is present only when the connector type is FEISHU.
	Feishu *UpdateKnowledgeSourceRequestDocumentFeishu `json:"Feishu,omitempty" xml:"Feishu,omitempty" type:"Struct"`
	// The Lark configuration. This parameter is present only when the connector type is LARK.
	Lark *UpdateKnowledgeSourceRequestDocumentLark `json:"Lark,omitempty" xml:"Lark,omitempty" type:"Struct"`
	// The WeCom configuration. This parameter is present only when the connector type is WECOM.
	WeCom *UpdateKnowledgeSourceRequestDocumentWeCom `json:"WeCom,omitempty" xml:"WeCom,omitempty" type:"Struct"`
	// The Yuque configuration. This parameter is present only when the connector type is YUQUE.
	Yuque *UpdateKnowledgeSourceRequestDocumentYuque `json:"Yuque,omitempty" xml:"Yuque,omitempty" type:"Struct"`
}

func (s UpdateKnowledgeSourceRequestDocument) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestDocument) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestDocument) GetConnectorType() *string {
	return s.ConnectorType
}

func (s *UpdateKnowledgeSourceRequestDocument) GetDingTalk() *UpdateKnowledgeSourceRequestDocumentDingTalk {
	return s.DingTalk
}

func (s *UpdateKnowledgeSourceRequestDocument) GetFeishu() *UpdateKnowledgeSourceRequestDocumentFeishu {
	return s.Feishu
}

func (s *UpdateKnowledgeSourceRequestDocument) GetLark() *UpdateKnowledgeSourceRequestDocumentLark {
	return s.Lark
}

func (s *UpdateKnowledgeSourceRequestDocument) GetWeCom() *UpdateKnowledgeSourceRequestDocumentWeCom {
	return s.WeCom
}

func (s *UpdateKnowledgeSourceRequestDocument) GetYuque() *UpdateKnowledgeSourceRequestDocumentYuque {
	return s.Yuque
}

func (s *UpdateKnowledgeSourceRequestDocument) SetConnectorType(v string) *UpdateKnowledgeSourceRequestDocument {
	s.ConnectorType = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocument) SetDingTalk(v *UpdateKnowledgeSourceRequestDocumentDingTalk) *UpdateKnowledgeSourceRequestDocument {
	s.DingTalk = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocument) SetFeishu(v *UpdateKnowledgeSourceRequestDocumentFeishu) *UpdateKnowledgeSourceRequestDocument {
	s.Feishu = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocument) SetLark(v *UpdateKnowledgeSourceRequestDocumentLark) *UpdateKnowledgeSourceRequestDocument {
	s.Lark = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocument) SetWeCom(v *UpdateKnowledgeSourceRequestDocumentWeCom) *UpdateKnowledgeSourceRequestDocument {
	s.WeCom = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocument) SetYuque(v *UpdateKnowledgeSourceRequestDocumentYuque) *UpdateKnowledgeSourceRequestDocument {
	s.Yuque = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocument) Validate() error {
	if s.DingTalk != nil {
		if err := s.DingTalk.Validate(); err != nil {
			return err
		}
	}
	if s.Feishu != nil {
		if err := s.Feishu.Validate(); err != nil {
			return err
		}
	}
	if s.Lark != nil {
		if err := s.Lark.Validate(); err != nil {
			return err
		}
	}
	if s.WeCom != nil {
		if err := s.WeCom.Validate(); err != nil {
			return err
		}
	}
	if s.Yuque != nil {
		if err := s.Yuque.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateKnowledgeSourceRequestDocumentDingTalk struct {
	// The MCP authorization credential for the DingTalk document connector. This credential is used only for creation or updates and is not returned in the response.
	//
	// example:
	//
	// 123456****
	McpKey *string `json:"McpKey,omitempty" xml:"McpKey,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeSourceRequestDocumentDingTalk) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestDocumentDingTalk) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestDocumentDingTalk) GetMcpKey() *string {
	return s.McpKey
}

func (s *UpdateKnowledgeSourceRequestDocumentDingTalk) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *UpdateKnowledgeSourceRequestDocumentDingTalk) SetMcpKey(v string) *UpdateKnowledgeSourceRequestDocumentDingTalk {
	s.McpKey = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentDingTalk) SetSourceUrls(v []*string) *UpdateKnowledgeSourceRequestDocumentDingTalk {
	s.SourceUrls = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentDingTalk) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeSourceRequestDocumentFeishu struct {
	// The application ID of Feishu or Lark. This is not the AppSecret.
	//
	// example:
	//
	// app-001
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The secret of the Feishu or Lark application. This secret is used only for creation or updates and is not returned in the response.
	//
	// example:
	//
	// 123456****
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeSourceRequestDocumentFeishu) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestDocumentFeishu) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestDocumentFeishu) GetAppId() *string {
	return s.AppId
}

func (s *UpdateKnowledgeSourceRequestDocumentFeishu) GetAppSecret() *string {
	return s.AppSecret
}

func (s *UpdateKnowledgeSourceRequestDocumentFeishu) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *UpdateKnowledgeSourceRequestDocumentFeishu) SetAppId(v string) *UpdateKnowledgeSourceRequestDocumentFeishu {
	s.AppId = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentFeishu) SetAppSecret(v string) *UpdateKnowledgeSourceRequestDocumentFeishu {
	s.AppSecret = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentFeishu) SetSourceUrls(v []*string) *UpdateKnowledgeSourceRequestDocumentFeishu {
	s.SourceUrls = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentFeishu) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeSourceRequestDocumentLark struct {
	// The application ID of Feishu or Lark. This is not the AppSecret.
	//
	// example:
	//
	// app-001
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The secret of the Feishu or Lark application. This secret is used only for creation or updates and is not returned in the response.
	//
	// example:
	//
	// 123456****
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeSourceRequestDocumentLark) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestDocumentLark) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestDocumentLark) GetAppId() *string {
	return s.AppId
}

func (s *UpdateKnowledgeSourceRequestDocumentLark) GetAppSecret() *string {
	return s.AppSecret
}

func (s *UpdateKnowledgeSourceRequestDocumentLark) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *UpdateKnowledgeSourceRequestDocumentLark) SetAppId(v string) *UpdateKnowledgeSourceRequestDocumentLark {
	s.AppId = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentLark) SetAppSecret(v string) *UpdateKnowledgeSourceRequestDocumentLark {
	s.AppSecret = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentLark) SetSourceUrls(v []*string) *UpdateKnowledgeSourceRequestDocumentLark {
	s.SourceUrls = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentLark) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeSourceRequestDocumentWeCom struct {
	// The WeCom bot ID.
	//
	// example:
	//
	// bot-001
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// The secret of the WeCom bot. This secret is used only for creation or updates and is not returned in the response.
	//
	// example:
	//
	// 123456****
	BotSecret *string `json:"BotSecret,omitempty" xml:"BotSecret,omitempty"`
	// The IDs of the shared spaces selected for the WeCom connector.
	SpaceIds []*string `json:"SpaceIds,omitempty" xml:"SpaceIds,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeSourceRequestDocumentWeCom) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestDocumentWeCom) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestDocumentWeCom) GetBotId() *string {
	return s.BotId
}

func (s *UpdateKnowledgeSourceRequestDocumentWeCom) GetBotSecret() *string {
	return s.BotSecret
}

func (s *UpdateKnowledgeSourceRequestDocumentWeCom) GetSpaceIds() []*string {
	return s.SpaceIds
}

func (s *UpdateKnowledgeSourceRequestDocumentWeCom) SetBotId(v string) *UpdateKnowledgeSourceRequestDocumentWeCom {
	s.BotId = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentWeCom) SetBotSecret(v string) *UpdateKnowledgeSourceRequestDocumentWeCom {
	s.BotSecret = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentWeCom) SetSpaceIds(v []*string) *UpdateKnowledgeSourceRequestDocumentWeCom {
	s.SpaceIds = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentWeCom) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeSourceRequestDocumentYuque struct {
	// The Yuque team configuration. Each team independently selects the synchronization scope.
	Teams []*UpdateKnowledgeSourceRequestDocumentYuqueTeams `json:"Teams,omitempty" xml:"Teams,omitempty" type:"Repeated"`
}

func (s UpdateKnowledgeSourceRequestDocumentYuque) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestDocumentYuque) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestDocumentYuque) GetTeams() []*UpdateKnowledgeSourceRequestDocumentYuqueTeams {
	return s.Teams
}

func (s *UpdateKnowledgeSourceRequestDocumentYuque) SetTeams(v []*UpdateKnowledgeSourceRequestDocumentYuqueTeams) *UpdateKnowledgeSourceRequestDocumentYuque {
	s.Teams = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentYuque) Validate() error {
	if s.Teams != nil {
		for _, item := range s.Teams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateKnowledgeSourceRequestDocumentYuqueTeams struct {
	// The URL of the team site.
	//
	// example:
	//
	// https://www.yuque.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
	// The Yuque team access credential. This credential is used only for creation or updates and is not returned in the response.
	//
	// example:
	//
	// 123456****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s UpdateKnowledgeSourceRequestDocumentYuqueTeams) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestDocumentYuqueTeams) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestDocumentYuqueTeams) GetHost() *string {
	return s.Host
}

func (s *UpdateKnowledgeSourceRequestDocumentYuqueTeams) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *UpdateKnowledgeSourceRequestDocumentYuqueTeams) GetToken() *string {
	return s.Token
}

func (s *UpdateKnowledgeSourceRequestDocumentYuqueTeams) SetHost(v string) *UpdateKnowledgeSourceRequestDocumentYuqueTeams {
	s.Host = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentYuqueTeams) SetSourceUrls(v []*string) *UpdateKnowledgeSourceRequestDocumentYuqueTeams {
	s.SourceUrls = v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentYuqueTeams) SetToken(v string) *UpdateKnowledgeSourceRequestDocumentYuqueTeams {
	s.Token = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestDocumentYuqueTeams) Validate() error {
	return dara.Validate(s)
}

type UpdateKnowledgeSourceRequestProject struct {
	// Optional. Limits the read environment to Dev or Prod. If this parameter is not specified, cross-environment read access is not granted.
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The DataWorks workspace ID, which is used to locate the source.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s UpdateKnowledgeSourceRequestProject) String() string {
	return dara.Prettify(s)
}

func (s UpdateKnowledgeSourceRequestProject) GoString() string {
	return s.String()
}

func (s *UpdateKnowledgeSourceRequestProject) GetEnvType() *string {
	return s.EnvType
}

func (s *UpdateKnowledgeSourceRequestProject) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateKnowledgeSourceRequestProject) SetEnvType(v string) *UpdateKnowledgeSourceRequestProject {
	s.EnvType = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestProject) SetProjectId(v int64) *UpdateKnowledgeSourceRequestProject {
	s.ProjectId = &v
	return s
}

func (s *UpdateKnowledgeSourceRequestProject) Validate() error {
	return dara.Validate(s)
}
