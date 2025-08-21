// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDocRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateDocRequest
	GetAgentKey() *string
	SetCategoryId(v int64) *CreateDocRequest
	GetCategoryId() *int64
	SetConfig(v string) *CreateDocRequest
	GetConfig() *string
	SetContent(v string) *CreateDocRequest
	GetContent() *string
	SetDocMetadata(v []*CreateDocRequestDocMetadata) *CreateDocRequest
	GetDocMetadata() []*CreateDocRequestDocMetadata
	SetEndDate(v string) *CreateDocRequest
	GetEndDate() *string
	SetMeta(v string) *CreateDocRequest
	GetMeta() *string
	SetStartDate(v string) *CreateDocRequest
	GetStartDate() *string
	SetTagIds(v []*int64) *CreateDocRequest
	GetTagIds() []*int64
	SetTitle(v string) *CreateDocRequest
	GetTitle() *string
	SetUrl(v string) *CreateDocRequest
	GetUrl() *string
}

type CreateDocRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30000049006
	CategoryId *int64 `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// example:
	//
	// {"Splitter":"treeSplitter","ChunkSize":500,"TreePatterns":["^# .*","^## .*","^### .*","^#### .*"],"TitleSource":""}
	Config      *string                        `json:"Config,omitempty" xml:"Config,omitempty"`
	Content     *string                        `json:"Content,omitempty" xml:"Content,omitempty"`
	DocMetadata []*CreateDocRequestDocMetadata `json:"DocMetadata,omitempty" xml:"DocMetadata,omitempty" type:"Repeated"`
	// example:
	//
	// 2032-05-25T16:28:36Z
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// {"code":"xxx"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// example:
	//
	// 2022-05-25T16:28:36Z
	StartDate *string  `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	TagIds    []*int64 `json:"TagIds,omitempty" xml:"TagIds,omitempty" type:"Repeated"`
	// This parameter is required.
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// example:
	//
	// https://example.com/example.pdf
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s CreateDocRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDocRequest) GoString() string {
	return s.String()
}

func (s *CreateDocRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateDocRequest) GetCategoryId() *int64 {
	return s.CategoryId
}

func (s *CreateDocRequest) GetConfig() *string {
	return s.Config
}

func (s *CreateDocRequest) GetContent() *string {
	return s.Content
}

func (s *CreateDocRequest) GetDocMetadata() []*CreateDocRequestDocMetadata {
	return s.DocMetadata
}

func (s *CreateDocRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *CreateDocRequest) GetMeta() *string {
	return s.Meta
}

func (s *CreateDocRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateDocRequest) GetTagIds() []*int64 {
	return s.TagIds
}

func (s *CreateDocRequest) GetTitle() *string {
	return s.Title
}

func (s *CreateDocRequest) GetUrl() *string {
	return s.Url
}

func (s *CreateDocRequest) SetAgentKey(v string) *CreateDocRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateDocRequest) SetCategoryId(v int64) *CreateDocRequest {
	s.CategoryId = &v
	return s
}

func (s *CreateDocRequest) SetConfig(v string) *CreateDocRequest {
	s.Config = &v
	return s
}

func (s *CreateDocRequest) SetContent(v string) *CreateDocRequest {
	s.Content = &v
	return s
}

func (s *CreateDocRequest) SetDocMetadata(v []*CreateDocRequestDocMetadata) *CreateDocRequest {
	s.DocMetadata = v
	return s
}

func (s *CreateDocRequest) SetEndDate(v string) *CreateDocRequest {
	s.EndDate = &v
	return s
}

func (s *CreateDocRequest) SetMeta(v string) *CreateDocRequest {
	s.Meta = &v
	return s
}

func (s *CreateDocRequest) SetStartDate(v string) *CreateDocRequest {
	s.StartDate = &v
	return s
}

func (s *CreateDocRequest) SetTagIds(v []*int64) *CreateDocRequest {
	s.TagIds = v
	return s
}

func (s *CreateDocRequest) SetTitle(v string) *CreateDocRequest {
	s.Title = &v
	return s
}

func (s *CreateDocRequest) SetUrl(v string) *CreateDocRequest {
	s.Url = &v
	return s
}

func (s *CreateDocRequest) Validate() error {
	return dara.Validate(s)
}

type CreateDocRequestDocMetadata struct {
	BusinessViewId      *string                                           `json:"BusinessViewId,omitempty" xml:"BusinessViewId,omitempty"`
	BusinessViewName    *string                                           `json:"BusinessViewName,omitempty" xml:"BusinessViewName,omitempty"`
	MetaCellInfoDTOList []*CreateDocRequestDocMetadataMetaCellInfoDTOList `json:"MetaCellInfoDTOList,omitempty" xml:"MetaCellInfoDTOList,omitempty" type:"Repeated"`
}

func (s CreateDocRequestDocMetadata) String() string {
	return dara.Prettify(s)
}

func (s CreateDocRequestDocMetadata) GoString() string {
	return s.String()
}

func (s *CreateDocRequestDocMetadata) GetBusinessViewId() *string {
	return s.BusinessViewId
}

func (s *CreateDocRequestDocMetadata) GetBusinessViewName() *string {
	return s.BusinessViewName
}

func (s *CreateDocRequestDocMetadata) GetMetaCellInfoDTOList() []*CreateDocRequestDocMetadataMetaCellInfoDTOList {
	return s.MetaCellInfoDTOList
}

func (s *CreateDocRequestDocMetadata) SetBusinessViewId(v string) *CreateDocRequestDocMetadata {
	s.BusinessViewId = &v
	return s
}

func (s *CreateDocRequestDocMetadata) SetBusinessViewName(v string) *CreateDocRequestDocMetadata {
	s.BusinessViewName = &v
	return s
}

func (s *CreateDocRequestDocMetadata) SetMetaCellInfoDTOList(v []*CreateDocRequestDocMetadataMetaCellInfoDTOList) *CreateDocRequestDocMetadata {
	s.MetaCellInfoDTOList = v
	return s
}

func (s *CreateDocRequestDocMetadata) Validate() error {
	return dara.Validate(s)
}

type CreateDocRequestDocMetadataMetaCellInfoDTOList struct {
	FieldCode *string `json:"FieldCode,omitempty" xml:"FieldCode,omitempty"`
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	Value     *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateDocRequestDocMetadataMetaCellInfoDTOList) String() string {
	return dara.Prettify(s)
}

func (s CreateDocRequestDocMetadataMetaCellInfoDTOList) GoString() string {
	return s.String()
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) GetFieldCode() *string {
	return s.FieldCode
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) GetValue() *string {
	return s.Value
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) SetFieldCode(v string) *CreateDocRequestDocMetadataMetaCellInfoDTOList {
	s.FieldCode = &v
	return s
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) SetFieldName(v string) *CreateDocRequestDocMetadataMetaCellInfoDTOList {
	s.FieldName = &v
	return s
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) SetValue(v string) *CreateDocRequestDocMetadataMetaCellInfoDTOList {
	s.Value = &v
	return s
}

func (s *CreateDocRequestDocMetadataMetaCellInfoDTOList) Validate() error {
	return dara.Validate(s)
}
