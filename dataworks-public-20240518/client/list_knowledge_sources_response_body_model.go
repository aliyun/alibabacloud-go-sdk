// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeSourcesResponseBody
	GetCode() *string
	SetData(v *ListKnowledgeSourcesResponseBodyData) *ListKnowledgeSourcesResponseBody
	GetData() *ListKnowledgeSourcesResponseBodyData
	SetMaxResults(v int32) *ListKnowledgeSourcesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListKnowledgeSourcesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListKnowledgeSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListKnowledgeSourcesResponseBody
	GetSuccess() *bool
}

type ListKnowledgeSourcesResponseBody struct {
	// The business error code. This code is used to determine the error type when the request fails.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The business data returned when the request is successful.
	Data *ListKnowledgeSourcesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// The request parameters are invalid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListKnowledgeSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeSourcesResponseBody) GetData() *ListKnowledgeSourcesResponseBodyData {
	return s.Data
}

func (s *ListKnowledgeSourcesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKnowledgeSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListKnowledgeSourcesResponseBody) SetCode(v string) *ListKnowledgeSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBody) SetData(v *ListKnowledgeSourcesResponseBodyData) *ListKnowledgeSourcesResponseBody {
	s.Data = v
	return s
}

func (s *ListKnowledgeSourcesResponseBody) SetMaxResults(v int32) *ListKnowledgeSourcesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBody) SetMessage(v string) *ListKnowledgeSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBody) SetRequestId(v string) *ListKnowledgeSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBody) SetSuccess(v bool) *ListKnowledgeSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKnowledgeSourcesResponseBodyData struct {
	// The pagination token for the next page. If no next page exists, the value is null.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The knowledge sources that you have permission to read on the current page.
	Sources []*ListKnowledgeSourcesResponseBodyDataSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
}

func (s ListKnowledgeSourcesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKnowledgeSourcesResponseBodyData) GetSources() []*ListKnowledgeSourcesResponseBodyDataSources {
	return s.Sources
}

func (s *ListKnowledgeSourcesResponseBodyData) SetNextToken(v string) *ListKnowledgeSourcesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyData) SetSources(v []*ListKnowledgeSourcesResponseBodyDataSources) *ListKnowledgeSourcesResponseBodyData {
	s.Sources = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyData) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeSourcesResponseBodyDataSources struct {
	// The creation time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1790200400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator. This field is written by the server.
	//
	// example:
	//
	// 1000000000000000
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The data source information. This field is available only when Type is set to DATA_SOURCE.
	DataSource *ListKnowledgeSourcesResponseBodyDataSourcesDataSource `json:"DataSource,omitempty" xml:"DataSource,omitempty" type:"Struct"`
	// The optional description of the purpose.
	//
	// example:
	//
	// Sales metrics and policy documents
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The document information. This field is available only when Type is set to DOCUMENT.
	Document *ListKnowledgeSourcesResponseBodyDataSourcesDocument `json:"Document,omitempty" xml:"Document,omitempty" type:"Struct"`
	// The ID of the knowledge base to which the source belongs. The source cannot be migrated to another knowledge base after it is created.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The last modification time. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1790200300000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The name used to identify the source.
	//
	// example:
	//
	// Sales Topic Project
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The project information. This field is available only when Type is set to PROJECT.
	Project *ListKnowledgeSourcesResponseBodyDataSourcesProject `json:"Project,omitempty" xml:"Project,omitempty" type:"Struct"`
	// The stable source ID generated by the service.
	//
	// example:
	//
	// src-001
	SourceId *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	// The source type. Valid values: PROJECT (workspace), DATA_SOURCE (DataWorks data source), or DOCUMENT (external document source).
	//
	// example:
	//
	// PROJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListKnowledgeSourcesResponseBodyDataSources) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetDataSource() *ListKnowledgeSourcesResponseBodyDataSourcesDataSource {
	return s.DataSource
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetDescription() *string {
	return s.Description
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetDocument() *ListKnowledgeSourcesResponseBodyDataSourcesDocument {
	return s.Document
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetProject() *ListKnowledgeSourcesResponseBodyDataSourcesProject {
	return s.Project
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetSourceId() *string {
	return s.SourceId
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) GetType() *string {
	return s.Type
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetCreateTime(v int64) *ListKnowledgeSourcesResponseBodyDataSources {
	s.CreateTime = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetCreateUser(v string) *ListKnowledgeSourcesResponseBodyDataSources {
	s.CreateUser = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetDataSource(v *ListKnowledgeSourcesResponseBodyDataSourcesDataSource) *ListKnowledgeSourcesResponseBodyDataSources {
	s.DataSource = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetDescription(v string) *ListKnowledgeSourcesResponseBodyDataSources {
	s.Description = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetDocument(v *ListKnowledgeSourcesResponseBodyDataSourcesDocument) *ListKnowledgeSourcesResponseBodyDataSources {
	s.Document = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetKnowledgeBaseId(v string) *ListKnowledgeSourcesResponseBodyDataSources {
	s.KnowledgeBaseId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetModifyTime(v int64) *ListKnowledgeSourcesResponseBodyDataSources {
	s.ModifyTime = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetName(v string) *ListKnowledgeSourcesResponseBodyDataSources {
	s.Name = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetProject(v *ListKnowledgeSourcesResponseBodyDataSourcesProject) *ListKnowledgeSourcesResponseBodyDataSources {
	s.Project = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetSourceId(v string) *ListKnowledgeSourcesResponseBodyDataSources {
	s.SourceId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) SetType(v string) *ListKnowledgeSourcesResponseBodyDataSources {
	s.Type = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSources) Validate() error {
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

type ListKnowledgeSourcesResponseBodyDataSourcesDataSource struct {
	// The ID of the DataWorks data source. The server uses this ID to resolve the workspace, environment, and data source type.
	//
	// example:
	//
	// 12345
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDataSource) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDataSource) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDataSource) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDataSource) SetDataSourceId(v int64) *ListKnowledgeSourcesResponseBodyDataSourcesDataSource {
	s.DataSourceId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDataSource) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeSourcesResponseBodyDataSourcesDocument struct {
	// The connector type. Valid values: DINGTALK, FEISHU, LARK, WECOM, or YUQUE. The connector type cannot be changed after the source is created.
	//
	// example:
	//
	// WECOM
	ConnectorType *string `json:"ConnectorType,omitempty" xml:"ConnectorType,omitempty"`
	// Indicates whether the credentials required by the connector are configured. This field is read-only, and the credential content or ciphertext is not returned.
	CredentialConfigured *bool `json:"CredentialConfigured,omitempty" xml:"CredentialConfigured,omitempty"`
	// The DingTalk configuration. This field is available only when the connector type is DingTalk.
	DingTalk *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk `json:"DingTalk,omitempty" xml:"DingTalk,omitempty" type:"Struct"`
	// The Feishu configuration. This field is available only when the connector type is Feishu.
	Feishu *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu `json:"Feishu,omitempty" xml:"Feishu,omitempty" type:"Struct"`
	// The Lark configuration. This field is available only when the connector type is Lark.
	Lark *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark `json:"Lark,omitempty" xml:"Lark,omitempty" type:"Struct"`
	// The WeCom configuration. This field is available only when the connector type is WeCom.
	WeCom *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom `json:"WeCom,omitempty" xml:"WeCom,omitempty" type:"Struct"`
	// The Yuque configuration. This field is available only when the connector type is Yuque.
	Yuque *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque `json:"Yuque,omitempty" xml:"Yuque,omitempty" type:"Struct"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocument) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocument) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) GetConnectorType() *string {
	return s.ConnectorType
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) GetCredentialConfigured() *bool {
	return s.CredentialConfigured
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) GetDingTalk() *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk {
	return s.DingTalk
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) GetFeishu() *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu {
	return s.Feishu
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) GetLark() *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark {
	return s.Lark
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) GetWeCom() *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom {
	return s.WeCom
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) GetYuque() *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque {
	return s.Yuque
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) SetConnectorType(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocument {
	s.ConnectorType = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) SetCredentialConfigured(v bool) *ListKnowledgeSourcesResponseBodyDataSourcesDocument {
	s.CredentialConfigured = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) SetDingTalk(v *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk) *ListKnowledgeSourcesResponseBodyDataSourcesDocument {
	s.DingTalk = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) SetFeishu(v *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu) *ListKnowledgeSourcesResponseBodyDataSourcesDocument {
	s.Feishu = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) SetLark(v *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark) *ListKnowledgeSourcesResponseBodyDataSourcesDocument {
	s.Lark = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) SetWeCom(v *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom) *ListKnowledgeSourcesResponseBodyDataSourcesDocument {
	s.WeCom = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) SetYuque(v *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque) *ListKnowledgeSourcesResponseBodyDataSourcesDocument {
	s.Yuque = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocument) Validate() error {
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

type ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk struct {
	// The Feishu or Lark application ID. This is not the AppSecret.
	//
	// example:
	//
	// app-001
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk) GetAppId() *string {
	return s.AppId
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk) SetAppId(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk {
	s.AppId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk) SetSourceUrls(v []*string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk {
	s.SourceUrls = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentDingTalk) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu struct {
	// The Feishu or Lark application ID. This is not the AppSecret.
	//
	// example:
	//
	// app-001
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu) GetAppId() *string {
	return s.AppId
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu) SetAppId(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu {
	s.AppId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu) SetSourceUrls(v []*string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu {
	s.SourceUrls = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentFeishu) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark struct {
	// The Feishu or Lark application ID. This is not the AppSecret.
	//
	// example:
	//
	// app-001
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark) GetAppId() *string {
	return s.AppId
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark) SetAppId(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark {
	s.AppId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark) SetSourceUrls(v []*string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark {
	s.SourceUrls = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentLark) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom struct {
	// The ID of the WeCom bot.
	//
	// example:
	//
	// bot-001
	BotId *string `json:"BotId,omitempty" xml:"BotId,omitempty"`
	// The selected shared spaces, including their identities and display names.
	Spaces []*ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces `json:"Spaces,omitempty" xml:"Spaces,omitempty" type:"Repeated"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom) GetBotId() *string {
	return s.BotId
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom) GetSpaces() []*ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces {
	return s.Spaces
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom) SetBotId(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom {
	s.BotId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom) SetSpaces(v []*ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom {
	s.Spaces = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeCom) Validate() error {
	if s.Spaces != nil {
		for _, item := range s.Spaces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces struct {
	// The identity of the shared space.
	//
	// example:
	//
	// space-001
	SpaceId *string `json:"SpaceId,omitempty" xml:"SpaceId,omitempty"`
	// The display name verified by the server. This field is not used for identity verification.
	//
	// example:
	//
	// R&D Knowledge
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces) GetSpaceId() *string {
	return s.SpaceId
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces) GetSpaceName() *string {
	return s.SpaceName
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces) SetSpaceId(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces {
	s.SpaceId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces) SetSpaceName(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces {
	s.SpaceName = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentWeComSpaces) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque struct {
	// The Yuque team configurations. Each team independently selects the synchronization scope.
	Teams []*ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams `json:"Teams,omitempty" xml:"Teams,omitempty" type:"Repeated"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque) GetTeams() []*ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams {
	return s.Teams
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque) SetTeams(v []*ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque {
	s.Teams = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuque) Validate() error {
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

type ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams struct {
	// The URL of the team site.
	//
	// example:
	//
	// https://www.yuque.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The URLs of the knowledge bases, spaces, or documents selected for this connector. The actual allowed scope is determined by the connector and authorization.
	SourceUrls []*string `json:"SourceUrls,omitempty" xml:"SourceUrls,omitempty" type:"Repeated"`
	// The team ID verified by the server.
	//
	// example:
	//
	// team-001
	TeamId *string `json:"TeamId,omitempty" xml:"TeamId,omitempty"`
	// The team logon name. This is read-only display information.
	//
	// example:
	//
	// example-team
	TeamLogin *string `json:"TeamLogin,omitempty" xml:"TeamLogin,omitempty"`
	// The team name. This is read-only display information.
	//
	// example:
	//
	// Sample Team
	TeamName *string `json:"TeamName,omitempty" xml:"TeamName,omitempty"`
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) GetHost() *string {
	return s.Host
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) GetSourceUrls() []*string {
	return s.SourceUrls
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) GetTeamId() *string {
	return s.TeamId
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) GetTeamLogin() *string {
	return s.TeamLogin
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) GetTeamName() *string {
	return s.TeamName
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) SetHost(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams {
	s.Host = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) SetSourceUrls(v []*string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams {
	s.SourceUrls = v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) SetTeamId(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams {
	s.TeamId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) SetTeamLogin(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams {
	s.TeamLogin = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) SetTeamName(v string) *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams {
	s.TeamName = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesDocumentYuqueTeams) Validate() error {
	return dara.Validate(s)
}

type ListKnowledgeSourcesResponseBodyDataSourcesProject struct {
	// The environment type. This optional field restricts the read environment to Dev or Prod. If this field is not specified, it does not mean that you have the permissions to read both environments.
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

func (s ListKnowledgeSourcesResponseBodyDataSourcesProject) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeSourcesResponseBodyDataSourcesProject) GoString() string {
	return s.String()
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesProject) GetEnvType() *string {
	return s.EnvType
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesProject) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesProject) SetEnvType(v string) *ListKnowledgeSourcesResponseBodyDataSourcesProject {
	s.EnvType = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesProject) SetProjectId(v int64) *ListKnowledgeSourcesResponseBodyDataSourcesProject {
	s.ProjectId = &v
	return s
}

func (s *ListKnowledgeSourcesResponseBodyDataSourcesProject) Validate() error {
	return dara.Validate(s)
}
