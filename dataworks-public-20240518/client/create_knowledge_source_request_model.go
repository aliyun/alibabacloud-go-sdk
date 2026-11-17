// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKnowledgeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateKnowledgeSourceRequest
	GetClientToken() *string
	SetDataSource(v *CreateKnowledgeSourceRequestDataSource) *CreateKnowledgeSourceRequest
	GetDataSource() *CreateKnowledgeSourceRequestDataSource
	SetDescription(v string) *CreateKnowledgeSourceRequest
	GetDescription() *string
	SetDocument(v *CreateKnowledgeSourceRequestDocument) *CreateKnowledgeSourceRequest
	GetDocument() *CreateKnowledgeSourceRequestDocument
	SetKnowledgeBaseId(v string) *CreateKnowledgeSourceRequest
	GetKnowledgeBaseId() *string
	SetName(v string) *CreateKnowledgeSourceRequest
	GetName() *string
	SetProject(v *CreateKnowledgeSourceRequestProject) *CreateKnowledgeSourceRequest
	GetProject() *CreateKnowledgeSourceRequestProject
	SetType(v string) *CreateKnowledgeSourceRequest
	GetType() *string
}

type CreateKnowledgeSourceRequest struct {
	// The client token used to ensure idempotence.
	//
	// example:
	//
	// source-create-001
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The data source configuration. This parameter is required when Type is set to DATA_SOURCE.
	DataSource *CreateKnowledgeSourceRequestDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The description.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document configuration. This parameter is required when Type is set to DOCUMENT.
	Document *CreateKnowledgeSourceRequestDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The ID of the knowledge base.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The name of the source.
	//
	// This parameter is required.
	//
	// example:
	//
	// Sales topic project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The project configuration. This parameter is required when Type is set to PROJECT.
	Project *CreateKnowledgeSourceRequestProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The type of the knowledge source. Valid values: PROJECT, DATA_SOURCE, and DOCUMENT.
	//
	// This parameter is required.
	//
	// example:
	//
	// PROJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateKnowledgeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateKnowledgeSourceRequest) GetDataSource() *CreateKnowledgeSourceRequestDataSource {
	return s.DataSource
}

func (s *CreateKnowledgeSourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateKnowledgeSourceRequest) GetDocument() *CreateKnowledgeSourceRequestDocument {
	return s.Document
}

func (s *CreateKnowledgeSourceRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *CreateKnowledgeSourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateKnowledgeSourceRequest) GetProject() *CreateKnowledgeSourceRequestProject {
	return s.Project
}

func (s *CreateKnowledgeSourceRequest) GetType() *string {
	return s.Type
}

func (s *CreateKnowledgeSourceRequest) SetClientToken(v string) *CreateKnowledgeSourceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateKnowledgeSourceRequest) SetDataSource(v *CreateKnowledgeSourceRequestDataSource) *CreateKnowledgeSourceRequest {
	s.DataSource = v
	return s
}

func (s *CreateKnowledgeSourceRequest) SetDescription(v string) *CreateKnowledgeSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateKnowledgeSourceRequest) SetDocument(v *CreateKnowledgeSourceRequestDocument) *CreateKnowledgeSourceRequest {
	s.Document = v
	return s
}

func (s *CreateKnowledgeSourceRequest) SetKnowledgeBaseId(v string) *CreateKnowledgeSourceRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *CreateKnowledgeSourceRequest) SetName(v string) *CreateKnowledgeSourceRequest {
	s.Name = &v
	return s
}

func (s *CreateKnowledgeSourceRequest) SetProject(v *CreateKnowledgeSourceRequestProject) *CreateKnowledgeSourceRequest {
	s.Project = v
	return s
}

func (s *CreateKnowledgeSourceRequest) SetType(v string) *CreateKnowledgeSourceRequest {
	s.Type = &v
	return s
}

func (s *CreateKnowledgeSourceRequest) Validate() error {
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

type CreateKnowledgeSourceRequestDataSource struct {
	// The ID of the DataWorks data source. The server uses this ID to resolve the associated workspace, environment, and data source type.
	//
	// example:
	//
	// 12345
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s CreateKnowledgeSourceRequestDataSource) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestDataSource) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestDataSource) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *CreateKnowledgeSourceRequestDataSource) SetDataSourceId(v int64) *CreateKnowledgeSourceRequestDataSource {
	s.DataSourceId = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDataSource) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeSourceRequestDocument struct {
	// The type of the connector. Valid values: DINGTALK, FEISHU, LARK, WECOM, and YUQUE. The connector type cannot be changed after creation.
	//
	// example:
	//
	// WECOM
	ConnectorType *string `json:"ConnectorType,omitempty" xml:"ConnectorType,omitempty"`
	// The DingTalk configuration. This parameter is available only when ConnectorType is set to DINGTALK.
	DingTalk *CreateKnowledgeSourceRequestDocumentDingTalk `json:"DingTalk,omitempty" xml:"DingTalk,omitempty" type:"Struct"`
	// The Feishu configuration. This parameter is available only when ConnectorType is set to FEISHU.
	Feishu *CreateKnowledgeSourceRequestDocumentFeishu `json:"Feishu,omitempty" xml:"Feishu,omitempty" type:"Struct"`
	// The Lark configuration. This parameter is available only when ConnectorType is set to LARK.
	Lark *CreateKnowledgeSourceRequestDocumentLark `json:"Lark,omitempty" xml:"Lark,omitempty" type:"Struct"`
	// The WeCom configuration. This parameter is available only when ConnectorType is set to WECOM.
	WeCom *CreateKnowledgeSourceRequestDocumentWeCom `json:"WeCom,omitempty" xml:"WeCom,omitempty" type:"Struct"`
	// The Yuque configuration. This parameter is available only when ConnectorType is set to YUQUE.
	Yuque *CreateKnowledgeSourceRequestDocumentYuque `json:"Yuque,omitempty" xml:"Yuque,omitempty" type:"Struct"`
}

func (s CreateKnowledgeSourceRequestDocument) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestDocument) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestDocument) GetConnectorType() *string {
	return s.ConnectorType
}

func (s *CreateKnowledgeSourceRequestDocument) GetDingTalk() *CreateKnowledgeSourceRequestDocumentDingTalk {
	return s.DingTalk
}

func (s *CreateKnowledgeSourceRequestDocument) GetFeishu() *CreateKnowledgeSourceRequestDocumentFeishu {
	return s.Feishu
}

func (s *CreateKnowledgeSourceRequestDocument) GetLark() *CreateKnowledgeSourceRequestDocumentLark {
	return s.Lark
}

func (s *CreateKnowledgeSourceRequestDocument) GetWeCom() *CreateKnowledgeSourceRequestDocumentWeCom {
	return s.WeCom
}

func (s *CreateKnowledgeSourceRequestDocument) GetYuque() *CreateKnowledgeSourceRequestDocumentYuque {
	return s.Yuque
}

func (s *CreateKnowledgeSourceRequestDocument) SetConnectorType(v string) *CreateKnowledgeSourceRequestDocument {
	s.ConnectorType = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocument) SetDingTalk(v *CreateKnowledgeSourceRequestDocumentDingTalk) *CreateKnowledgeSourceRequestDocument {
	s.DingTalk = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocument) SetFeishu(v *CreateKnowledgeSourceRequestDocumentFeishu) *CreateKnowledgeSourceRequestDocument {
	s.Feishu = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocument) SetLark(v *CreateKnowledgeSourceRequestDocumentLark) *CreateKnowledgeSourceRequestDocument {
	s.Lark = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocument) SetWeCom(v *CreateKnowledgeSourceRequestDocumentWeCom) *CreateKnowledgeSourceRequestDocument {
	s.WeCom = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocument) SetYuque(v *CreateKnowledgeSourceRequestDocumentYuque) *CreateKnowledgeSourceRequestDocument {
	s.Yuque = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocument) Validate() error {
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

type CreateKnowledgeSourceRequestDocumentDingTalk struct {
	// The MCP authorization credential for the DingTalk document connector. This credential is used only for creation or updates and is not returned in the response.
	//
	// example:
	//
	// 123456****
	McpKey *string `json:"McpKey,omitempty" xml:"McpKey,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization settings.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeSourceRequestDocumentDingTalk) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestDocumentDingTalk) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestDocumentDingTalk) GetMcpKey() *string {
	return s.McpKey
}

func (s *CreateKnowledgeSourceRequestDocumentDingTalk) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *CreateKnowledgeSourceRequestDocumentDingTalk) SetMcpKey(v string) *CreateKnowledgeSourceRequestDocumentDingTalk {
	s.McpKey = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentDingTalk) SetSourceUrls(v []*string) *CreateKnowledgeSourceRequestDocumentDingTalk {
	s.SourceUrls = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentDingTalk) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeSourceRequestDocumentFeishu struct {
	// The ID of the Feishu or Lark application. This is not the AppSecret.
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
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization settings.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeSourceRequestDocumentFeishu) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestDocumentFeishu) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestDocumentFeishu) GetAppId() *string {
	return s.AppId
}

func (s *CreateKnowledgeSourceRequestDocumentFeishu) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateKnowledgeSourceRequestDocumentFeishu) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *CreateKnowledgeSourceRequestDocumentFeishu) SetAppId(v string) *CreateKnowledgeSourceRequestDocumentFeishu {
	s.AppId = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentFeishu) SetAppSecret(v string) *CreateKnowledgeSourceRequestDocumentFeishu {
	s.AppSecret = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentFeishu) SetSourceUrls(v []*string) *CreateKnowledgeSourceRequestDocumentFeishu {
	s.SourceUrls = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentFeishu) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeSourceRequestDocumentLark struct {
	// The ID of the Feishu or Lark application. This is not the AppSecret.
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
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization settings.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeSourceRequestDocumentLark) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestDocumentLark) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestDocumentLark) GetAppId() *string {
	return s.AppId
}

func (s *CreateKnowledgeSourceRequestDocumentLark) GetAppSecret() *string {
	return s.AppSecret
}

func (s *CreateKnowledgeSourceRequestDocumentLark) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *CreateKnowledgeSourceRequestDocumentLark) SetAppId(v string) *CreateKnowledgeSourceRequestDocumentLark {
	s.AppId = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentLark) SetAppSecret(v string) *CreateKnowledgeSourceRequestDocumentLark {
	s.AppSecret = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentLark) SetSourceUrls(v []*string) *CreateKnowledgeSourceRequestDocumentLark {
	s.SourceUrls = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentLark) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeSourceRequestDocumentWeCom struct {
	// The ID of the WeCom bot.
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

func (s CreateKnowledgeSourceRequestDocumentWeCom) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestDocumentWeCom) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestDocumentWeCom) GetBotId() *string {
	return s.BotId
}

func (s *CreateKnowledgeSourceRequestDocumentWeCom) GetBotSecret() *string {
	return s.BotSecret
}

func (s *CreateKnowledgeSourceRequestDocumentWeCom) GetSpaceIds() []*string {
	return s.SpaceIds
}

func (s *CreateKnowledgeSourceRequestDocumentWeCom) SetBotId(v string) *CreateKnowledgeSourceRequestDocumentWeCom {
	s.BotId = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentWeCom) SetBotSecret(v string) *CreateKnowledgeSourceRequestDocumentWeCom {
	s.BotSecret = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentWeCom) SetSpaceIds(v []*string) *CreateKnowledgeSourceRequestDocumentWeCom {
	s.SpaceIds = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentWeCom) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeSourceRequestDocumentYuque struct {
	// The Yuque team configurations. Each team independently selects the synchronization scope.
	Teams []*CreateKnowledgeSourceRequestDocumentYuqueTeams `json:"Teams,omitempty" xml:"Teams,omitempty" type:"Repeated"`
}

func (s CreateKnowledgeSourceRequestDocumentYuque) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestDocumentYuque) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestDocumentYuque) GetTeams() []*CreateKnowledgeSourceRequestDocumentYuqueTeams {
	return s.Teams
}

func (s *CreateKnowledgeSourceRequestDocumentYuque) SetTeams(v []*CreateKnowledgeSourceRequestDocumentYuqueTeams) *CreateKnowledgeSourceRequestDocumentYuque {
	s.Teams = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentYuque) Validate() error {
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

type CreateKnowledgeSourceRequestDocumentYuqueTeams struct {
	// The URL of the team site.
	//
	// example:
	//
	// https://www.yuque.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization settings.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
	// The access token for the Yuque team. This token is used only for creation or updates and is not returned in the response.
	//
	// example:
	//
	// 123456****
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s CreateKnowledgeSourceRequestDocumentYuqueTeams) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestDocumentYuqueTeams) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestDocumentYuqueTeams) GetHost() *string {
	return s.Host
}

func (s *CreateKnowledgeSourceRequestDocumentYuqueTeams) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *CreateKnowledgeSourceRequestDocumentYuqueTeams) GetToken() *string {
	return s.Token
}

func (s *CreateKnowledgeSourceRequestDocumentYuqueTeams) SetHost(v string) *CreateKnowledgeSourceRequestDocumentYuqueTeams {
	s.Host = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentYuqueTeams) SetSourceUrls(v []*string) *CreateKnowledgeSourceRequestDocumentYuqueTeams {
	s.SourceUrls = v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentYuqueTeams) SetToken(v string) *CreateKnowledgeSourceRequestDocumentYuqueTeams {
	s.Token = &v
	return s
}

func (s *CreateKnowledgeSourceRequestDocumentYuqueTeams) Validate() error {
	return dara.Validate(s)
}

type CreateKnowledgeSourceRequestProject struct {
	// The environment type. This parameter is optional and restricts the read environment to Dev or Prod. If this parameter is not specified, it does not mean that you have the permissions to read data from both environments.
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The ID of the DataWorks workspace. This field is used to locate the source.
	//
	// example:
	//
	// 12345
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s CreateKnowledgeSourceRequestProject) String() string {
	return dara.Prettify(s)
}

func (s CreateKnowledgeSourceRequestProject) GoString() string {
	return s.String()
}

func (s *CreateKnowledgeSourceRequestProject) GetEnvType() *string {
	return s.EnvType
}

func (s *CreateKnowledgeSourceRequestProject) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateKnowledgeSourceRequestProject) SetEnvType(v string) *CreateKnowledgeSourceRequestProject {
	s.EnvType = &v
	return s
}

func (s *CreateKnowledgeSourceRequestProject) SetProjectId(v int64) *CreateKnowledgeSourceRequestProject {
	s.ProjectId = &v
	return s
}

func (s *CreateKnowledgeSourceRequestProject) Validate() error {
	return dara.Validate(s)
}
