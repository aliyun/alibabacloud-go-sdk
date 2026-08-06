// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedup(v *AddDocumentsRequestDedup) *AddDocumentsRequest
	GetDedup() *AddDocumentsRequestDedup
	SetDocuments(v []*AddDocumentsRequestDocuments) *AddDocumentsRequest
	GetDocuments() []*AddDocumentsRequestDocuments
	SetImportType(v string) *AddDocumentsRequest
	GetImportType() *string
	SetKnowledgeBaseId(v string) *AddDocumentsRequest
	GetKnowledgeBaseId() *string
	SetMetaFields(v interface{}) *AddDocumentsRequest
	GetMetaFields() interface{}
	SetStrategyId(v string) *AddDocumentsRequest
	GetStrategyId() *string
	SetDingTalkConfiguration(v *AddDocumentsRequestDingTalkConfiguration) *AddDocumentsRequest
	GetDingTalkConfiguration() *AddDocumentsRequestDingTalkConfiguration
}

type AddDocumentsRequest struct {
	Dedup     *AddDocumentsRequestDedup       `json:"Dedup,omitempty" xml:"Dedup,omitempty" type:"Struct"`
	Documents []*AddDocumentsRequestDocuments `json:"Documents,omitempty" xml:"Documents,omitempty" type:"Repeated"`
	// 当前支持 LOCAL_UPLOAD；OSS_IMPORT 和 PUBLIC_URL 为后续导入方式预留。
	//
	// example:
	//
	// LOCAL_UPLOAD
	ImportType *string `json:"ImportType,omitempty" xml:"ImportType,omitempty"`
	// example:
	//
	// kb-3bd02617e9be335f
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// 导入时批量设置到本批次所有知识数据的标签键值。Key 必须为知识库已定义标签字段；Value 支持 string、int64、float32、bool、list。
	//
	// example:
	//
	// {"department":"legal","topics":["policy","contract"],"reviewed":true}
	MetaFields interface{} `json:"MetaFields,omitempty" xml:"MetaFields,omitempty"`
	// example:
	//
	// kb-strategy-7043984ca395eabd
	StrategyId            *string                                   `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	DingTalkConfiguration *AddDocumentsRequestDingTalkConfiguration `json:"dingTalkConfiguration,omitempty" xml:"dingTalkConfiguration,omitempty" type:"Struct"`
}

func (s AddDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentsRequest) GoString() string {
	return s.String()
}

func (s *AddDocumentsRequest) GetDedup() *AddDocumentsRequestDedup {
	return s.Dedup
}

func (s *AddDocumentsRequest) GetDocuments() []*AddDocumentsRequestDocuments {
	return s.Documents
}

func (s *AddDocumentsRequest) GetImportType() *string {
	return s.ImportType
}

func (s *AddDocumentsRequest) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *AddDocumentsRequest) GetMetaFields() interface{} {
	return s.MetaFields
}

func (s *AddDocumentsRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *AddDocumentsRequest) GetDingTalkConfiguration() *AddDocumentsRequestDingTalkConfiguration {
	return s.DingTalkConfiguration
}

func (s *AddDocumentsRequest) SetDedup(v *AddDocumentsRequestDedup) *AddDocumentsRequest {
	s.Dedup = v
	return s
}

func (s *AddDocumentsRequest) SetDocuments(v []*AddDocumentsRequestDocuments) *AddDocumentsRequest {
	s.Documents = v
	return s
}

func (s *AddDocumentsRequest) SetImportType(v string) *AddDocumentsRequest {
	s.ImportType = &v
	return s
}

func (s *AddDocumentsRequest) SetKnowledgeBaseId(v string) *AddDocumentsRequest {
	s.KnowledgeBaseId = &v
	return s
}

func (s *AddDocumentsRequest) SetMetaFields(v interface{}) *AddDocumentsRequest {
	s.MetaFields = v
	return s
}

func (s *AddDocumentsRequest) SetStrategyId(v string) *AddDocumentsRequest {
	s.StrategyId = &v
	return s
}

func (s *AddDocumentsRequest) SetDingTalkConfiguration(v *AddDocumentsRequestDingTalkConfiguration) *AddDocumentsRequest {
	s.DingTalkConfiguration = v
	return s
}

func (s *AddDocumentsRequest) Validate() error {
	if s.Dedup != nil {
		if err := s.Dedup.Validate(); err != nil {
			return err
		}
	}
	if s.Documents != nil {
		for _, item := range s.Documents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DingTalkConfiguration != nil {
		if err := s.DingTalkConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddDocumentsRequestDedup struct {
	// example:
	//
	// true
	ContentDedup *bool `json:"ContentDedup,omitempty" xml:"ContentDedup,omitempty"`
	// example:
	//
	// true
	DocNameDedup *bool `json:"DocNameDedup,omitempty" xml:"DocNameDedup,omitempty"`
}

func (s AddDocumentsRequestDedup) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentsRequestDedup) GoString() string {
	return s.String()
}

func (s *AddDocumentsRequestDedup) GetContentDedup() *bool {
	return s.ContentDedup
}

func (s *AddDocumentsRequestDedup) GetDocNameDedup() *bool {
	return s.DocNameDedup
}

func (s *AddDocumentsRequestDedup) SetContentDedup(v bool) *AddDocumentsRequestDedup {
	s.ContentDedup = &v
	return s
}

func (s *AddDocumentsRequestDedup) SetDocNameDedup(v bool) *AddDocumentsRequestDedup {
	s.DocNameDedup = &v
	return s
}

func (s *AddDocumentsRequestDedup) Validate() error {
	return dara.Validate(s)
}

type AddDocumentsRequestDocuments struct {
	// example:
	//
	// CHANGELOG.md
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 本地上传时为预签名上传使用的批次相对路径；不同 ImportType 下含义由导入类型定义。
	//
	// example:
	//
	// 2026_06_23_17_49_52WwGSUezpG2u2iHWxyYGzkf9KtormhkxN/CHANGELOG.md
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// 1024
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s AddDocumentsRequestDocuments) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentsRequestDocuments) GoString() string {
	return s.String()
}

func (s *AddDocumentsRequestDocuments) GetName() *string {
	return s.Name
}

func (s *AddDocumentsRequestDocuments) GetPath() *string {
	return s.Path
}

func (s *AddDocumentsRequestDocuments) GetSize() *int64 {
	return s.Size
}

func (s *AddDocumentsRequestDocuments) SetName(v string) *AddDocumentsRequestDocuments {
	s.Name = &v
	return s
}

func (s *AddDocumentsRequestDocuments) SetPath(v string) *AddDocumentsRequestDocuments {
	s.Path = &v
	return s
}

func (s *AddDocumentsRequestDocuments) SetSize(v int64) *AddDocumentsRequestDocuments {
	s.Size = &v
	return s
}

func (s *AddDocumentsRequestDocuments) Validate() error {
	return dara.Validate(s)
}

type AddDocumentsRequestDingTalkConfiguration struct {
	// example:
	//
	// ignore
	AppId *string `json:"appId,omitempty" xml:"appId,omitempty"`
	// example:
	//
	// ignore
	AppPassword *string `json:"appPassword,omitempty" xml:"appPassword,omitempty"`
	// example:
	//
	// ignore
	DingDocMcpLink *string `json:"dingDocMcpLink,omitempty" xml:"dingDocMcpLink,omitempty"`
	// example:
	//
	// ignore
	DingTableMcpLink *string `json:"dingTableMcpLink,omitempty" xml:"dingTableMcpLink,omitempty"`
	// example:
	//
	// ignore
	KnowledgeId *string `json:"knowledgeId,omitempty" xml:"knowledgeId,omitempty"`
	// example:
	//
	// ignore
	KnowledgeType *string `json:"knowledgeType,omitempty" xml:"knowledgeType,omitempty"`
	// example:
	//
	// ignore
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s AddDocumentsRequestDingTalkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s AddDocumentsRequestDingTalkConfiguration) GoString() string {
	return s.String()
}

func (s *AddDocumentsRequestDingTalkConfiguration) GetAppId() *string {
	return s.AppId
}

func (s *AddDocumentsRequestDingTalkConfiguration) GetAppPassword() *string {
	return s.AppPassword
}

func (s *AddDocumentsRequestDingTalkConfiguration) GetDingDocMcpLink() *string {
	return s.DingDocMcpLink
}

func (s *AddDocumentsRequestDingTalkConfiguration) GetDingTableMcpLink() *string {
	return s.DingTableMcpLink
}

func (s *AddDocumentsRequestDingTalkConfiguration) GetKnowledgeId() *string {
	return s.KnowledgeId
}

func (s *AddDocumentsRequestDingTalkConfiguration) GetKnowledgeType() *string {
	return s.KnowledgeType
}

func (s *AddDocumentsRequestDingTalkConfiguration) GetUserId() *string {
	return s.UserId
}

func (s *AddDocumentsRequestDingTalkConfiguration) SetAppId(v string) *AddDocumentsRequestDingTalkConfiguration {
	s.AppId = &v
	return s
}

func (s *AddDocumentsRequestDingTalkConfiguration) SetAppPassword(v string) *AddDocumentsRequestDingTalkConfiguration {
	s.AppPassword = &v
	return s
}

func (s *AddDocumentsRequestDingTalkConfiguration) SetDingDocMcpLink(v string) *AddDocumentsRequestDingTalkConfiguration {
	s.DingDocMcpLink = &v
	return s
}

func (s *AddDocumentsRequestDingTalkConfiguration) SetDingTableMcpLink(v string) *AddDocumentsRequestDingTalkConfiguration {
	s.DingTableMcpLink = &v
	return s
}

func (s *AddDocumentsRequestDingTalkConfiguration) SetKnowledgeId(v string) *AddDocumentsRequestDingTalkConfiguration {
	s.KnowledgeId = &v
	return s
}

func (s *AddDocumentsRequestDingTalkConfiguration) SetKnowledgeType(v string) *AddDocumentsRequestDingTalkConfiguration {
	s.KnowledgeType = &v
	return s
}

func (s *AddDocumentsRequestDingTalkConfiguration) SetUserId(v string) *AddDocumentsRequestDingTalkConfiguration {
	s.UserId = &v
	return s
}

func (s *AddDocumentsRequestDingTalkConfiguration) Validate() error {
	return dara.Validate(s)
}
