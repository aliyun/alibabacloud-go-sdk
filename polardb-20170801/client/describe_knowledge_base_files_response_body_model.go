// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKnowledgeBaseFilesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeKnowledgeBaseFilesResponseBodyItems) *DescribeKnowledgeBaseFilesResponseBody
	GetItems() []*DescribeKnowledgeBaseFilesResponseBodyItems
	SetPageNumber(v int32) *DescribeKnowledgeBaseFilesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeKnowledgeBaseFilesResponseBody
	GetPageRecordCount() *int32
	SetPageSize(v int32) *DescribeKnowledgeBaseFilesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeKnowledgeBaseFilesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeKnowledgeBaseFilesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeKnowledgeBaseFilesResponseBody struct {
	Items []*DescribeKnowledgeBaseFilesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 9
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// CED079B7-A408-41A1-BFF1-EC608E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 9
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeKnowledgeBaseFilesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBody) GetItems() []*DescribeKnowledgeBaseFilesResponseBodyItems {
	return s.Items
}

func (s *DescribeKnowledgeBaseFilesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKnowledgeBaseFilesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeKnowledgeBaseFilesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeKnowledgeBaseFilesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKnowledgeBaseFilesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeKnowledgeBaseFilesResponseBody) SetItems(v []*DescribeKnowledgeBaseFilesResponseBodyItems) *DescribeKnowledgeBaseFilesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBody) SetPageNumber(v int32) *DescribeKnowledgeBaseFilesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBody) SetPageRecordCount(v int32) *DescribeKnowledgeBaseFilesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBody) SetPageSize(v int32) *DescribeKnowledgeBaseFilesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBody) SetRequestId(v string) *DescribeKnowledgeBaseFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBody) SetTotalRecordCount(v int32) *DescribeKnowledgeBaseFilesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKnowledgeBaseFilesResponseBodyItems struct {
	// example:
	//
	// Not Support.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 5b2dbb13-xxxx-xxxx-xxxx-a55fe8daec8f
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// 财报.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// 318881
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// example:
	//
	// pdf
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// true
	InheritSpaceStrategy *bool `json:"InheritSpaceStrategy,omitempty" xml:"InheritSpaceStrategy,omitempty"`
	// example:
	//
	// pkb-xxxxxx
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// example:
	//
	// pks-xxxxxx
	KnowledgeSpaceId *string `json:"KnowledgeSpaceId,omitempty" xml:"KnowledgeSpaceId,omitempty"`
	// example:
	//
	// {}
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// oss://test-bucket-example/pks-xxxx/pkb-xxxx/财报.pdf
	OSSPath *string `json:"OSSPath,omitempty" xml:"OSSPath,omitempty"`
	// example:
	//
	// 10
	ShardCount             *int32                                                             `json:"ShardCount,omitempty" xml:"ShardCount,omitempty"`
	ShardingStrategyConfig *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig `json:"ShardingStrategyConfig,omitempty" xml:"ShardingStrategyConfig,omitempty" type:"Struct"`
	// example:
	//
	// OSS
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2026-06-15T22:28:53Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// example:
	//
	// 2026-06-15T22:28:53Z
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s DescribeKnowledgeBaseFilesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetFileId() *string {
	return s.FileId
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetFileName() *string {
	return s.FileName
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetFileSize() *int64 {
	return s.FileSize
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetFileType() *string {
	return s.FileType
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetInheritSpaceStrategy() *bool {
	return s.InheritSpaceStrategy
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetKnowledgeSpaceId() *string {
	return s.KnowledgeSpaceId
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetOSSPath() *string {
	return s.OSSPath
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetShardCount() *int32 {
	return s.ShardCount
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetShardingStrategyConfig() *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig {
	return s.ShardingStrategyConfig
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) GetUploadTime() *string {
	return s.UploadTime
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetErrorMessage(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetFileId(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.FileId = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetFileName(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.FileName = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetFileSize(v int64) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.FileSize = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetFileType(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.FileType = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetInheritSpaceStrategy(v bool) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.InheritSpaceStrategy = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetKnowledgeBaseId(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.KnowledgeBaseId = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetKnowledgeSpaceId(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.KnowledgeSpaceId = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetMetadata(v map[string]interface{}) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.Metadata = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetOSSPath(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.OSSPath = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetShardCount(v int32) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.ShardCount = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetShardingStrategyConfig(v *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.ShardingStrategyConfig = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetSourceType(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetStatus(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.Status = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetUpdatedAt(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.UpdatedAt = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) SetUploadTime(v string) *DescribeKnowledgeBaseFilesResponseBodyItems {
	s.UploadTime = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItems) Validate() error {
	if s.ShardingStrategyConfig != nil {
		if err := s.ShardingStrategyConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig struct {
	DefaultStrategy *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy `json:"DefaultStrategy,omitempty" xml:"DefaultStrategy,omitempty" type:"Struct"`
	Rules           []*DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules         `json:"Rules,omitempty" xml:"Rules,omitempty" type:"Repeated"`
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig) GetDefaultStrategy() *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy {
	return s.DefaultStrategy
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig) GetRules() []*DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules {
	return s.Rules
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig) SetDefaultStrategy(v *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig {
	s.DefaultStrategy = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig) SetRules(v []*DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig {
	s.Rules = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfig) Validate() error {
	if s.DefaultStrategy != nil {
		if err := s.DefaultStrategy.Validate(); err != nil {
			return err
		}
	}
	if s.Rules != nil {
		for _, item := range s.Rules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy struct {
	Parameters *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// example:
	//
	// hybrid
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy) GetParameters() *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters {
	return s.Parameters
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy) GetType() *string {
	return s.Type
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy) SetParameters(v *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy {
	s.Parameters = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy) SetType(v string) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy {
	s.Type = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategy) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters struct {
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
	// example:
	//
	// true
	MergePeers *bool `json:"MergePeers,omitempty" xml:"MergePeers,omitempty"`
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters) GetMergePeers() *bool {
	return s.MergePeers
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters) SetMaxTokens(v int32) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters {
	s.MaxTokens = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters) SetMergePeers(v bool) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters {
	s.MergePeers = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigDefaultStrategyParameters) Validate() error {
	return dara.Validate(s)
}

type DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules struct {
	Match    *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch    `json:"Match,omitempty" xml:"Match,omitempty" type:"Struct"`
	Strategy *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy `json:"Strategy,omitempty" xml:"Strategy,omitempty" type:"Struct"`
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules) GetMatch() *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch {
	return s.Match
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules) GetStrategy() *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy {
	return s.Strategy
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules) SetMatch(v *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules {
	s.Match = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules) SetStrategy(v *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules {
	s.Strategy = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRules) Validate() error {
	if s.Match != nil {
		if err := s.Match.Validate(); err != nil {
			return err
		}
	}
	if s.Strategy != nil {
		if err := s.Strategy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch struct {
	// example:
	//
	// table
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch) GetContentType() *string {
	return s.ContentType
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch) SetContentType(v string) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch {
	s.ContentType = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesMatch) Validate() error {
	return dara.Validate(s)
}

type DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy struct {
	Parameters *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// example:
	//
	// hierarchical
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy) GetParameters() *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters {
	return s.Parameters
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy) GetType() *string {
	return s.Type
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy) SetParameters(v *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy {
	s.Parameters = v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy) SetType(v string) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy {
	s.Type = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategy) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters struct {
	// example:
	//
	// auto
	MarkdownTables *string `json:"MarkdownTables,omitempty" xml:"MarkdownTables,omitempty"`
	// example:
	//
	// 512
	MaxTokens *int32 `json:"MaxTokens,omitempty" xml:"MaxTokens,omitempty"`
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters) GoString() string {
	return s.String()
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters) GetMarkdownTables() *string {
	return s.MarkdownTables
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters) GetMaxTokens() *int32 {
	return s.MaxTokens
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters) SetMarkdownTables(v string) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters {
	s.MarkdownTables = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters) SetMaxTokens(v int32) *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters {
	s.MaxTokens = &v
	return s
}

func (s *DescribeKnowledgeBaseFilesResponseBodyItemsShardingStrategyConfigRulesStrategyParameters) Validate() error {
	return dara.Validate(s)
}
