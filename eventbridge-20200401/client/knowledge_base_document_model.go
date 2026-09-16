// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKnowledgeBaseDocument interface {
	dara.Model
	String() string
	GoString() string
	SetChunkConfiguration(v *KnowledgeBaseDocumentChunkConfiguration) *KnowledgeBaseDocument
	GetChunkConfiguration() *KnowledgeBaseDocumentChunkConfiguration
	SetChunkCount(v int32) *KnowledgeBaseDocument
	GetChunkCount() *int32
	SetCreatedAt(v string) *KnowledgeBaseDocument
	GetCreatedAt() *string
	SetDocumentId(v string) *KnowledgeBaseDocument
	GetDocumentId() *string
	SetErrorCode(v string) *KnowledgeBaseDocument
	GetErrorCode() *string
	SetErrorMessage(v string) *KnowledgeBaseDocument
	GetErrorMessage() *string
	SetFileName(v string) *KnowledgeBaseDocument
	GetFileName() *string
	SetFileSize(v int64) *KnowledgeBaseDocument
	GetFileSize() *int64
	SetMetadata(v []*KnowledgeBaseDocumentMetadata) *KnowledgeBaseDocument
	GetMetadata() []*KnowledgeBaseDocumentMetadata
	SetSourceModifiedTime(v int64) *KnowledgeBaseDocument
	GetSourceModifiedTime() *int64
	SetSourceType(v string) *KnowledgeBaseDocument
	GetSourceType() *string
	SetSourceUri(v string) *KnowledgeBaseDocument
	GetSourceUri() *string
	SetStatus(v string) *KnowledgeBaseDocument
	GetStatus() *string
	SetUpdatedAt(v string) *KnowledgeBaseDocument
	GetUpdatedAt() *string
}

type KnowledgeBaseDocument struct {
	// The snapshot of the document-level chunking policy actually used for this document. This field is returned only if ChunkConfiguration was explicitly specified during upload (BeginUpload) or update (UpdateDocument). If not specified, the document is chunked based on the knowledge base-level default configurations, and this field is not returned. The knowledge base-level configuration is not echoed back to avoid misleading users about the actual chunking basis for this document when the knowledge base-level configuration is subsequently changed.
	ChunkConfiguration *KnowledgeBaseDocumentChunkConfiguration `json:"ChunkConfiguration,omitempty" xml:"ChunkConfiguration,omitempty" type:"Struct"`
	// The number of chunks generated after processing is complete.
	//
	// example:
	//
	// 120
	ChunkCount *int32 `json:"ChunkCount,omitempty" xml:"ChunkCount,omitempty"`
	// The time when the document was created.
	//
	// example:
	//
	// 2026-08-24T10:00:00Z
	CreatedAt *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The unique identifier of the document.
	//
	// example:
	//
	// doc-bp1xxxxxxxxxxxx
	DocumentId *string `json:"DocumentId,omitempty" xml:"DocumentId,omitempty"`
	// The stable error code returned when processing fails.
	//
	// example:
	//
	// FILE_CORRUPTED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The desensitized error message returned when processing fails.
	//
	// example:
	//
	// parse pdf failed
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The file name of the document.
	//
	// example:
	//
	// manual.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file size of the document, in bytes.
	//
	// example:
	//
	// 1048576
	FileSize *int64 `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The document-level metadata key-value pairs, including constant field values and system variable values. This field is not returned if no metadata is specified.
	//
	// example:
	//
	// [{"Key":"department","Value":"R&D"}]
	Metadata []*KnowledgeBaseDocumentMetadata `json:"Metadata,omitempty" xml:"Metadata,omitempty" type:"Repeated"`
	// The last modification time in the upstream source system, in epoch milliseconds. This field is empty if no source information is available.
	//
	// example:
	//
	// 1788000000000
	SourceModifiedTime *int64 `json:"SourceModifiedTime,omitempty" xml:"SourceModifiedTime,omitempty"`
	// The delivery channel through which the document entered the knowledge base. This value is written by the system and cannot be specified by users. Valid values:
	//
	// - UPLOAD: manually uploaded through the console or API.
	//
	// - OSS: imported through an OSS event stream.
	//
	// New values may be added when new channels are supported. The values are not restricted to a fixed enumeration.
	//
	// example:
	//
	// UPLOAD
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The original source address of the document, such as oss://bucket/path/file.md. This field may be empty for manually uploaded documents.
	//
	// example:
	//
	// oss://my-bucket/docs/handbook.pdf
	SourceUri *string `json:"SourceUri,omitempty" xml:"SourceUri,omitempty"`
	// The processing status of the document. Valid values:
	//
	// - UPLOADING: uploading in progress.
	//
	// - PENDING: upload complete and queued for processing. This is typically a transitional state that lasts for seconds.
	//
	// - PROCESSING: parsing and processing in progress.
	//
	// - COMPLETED: processing complete and searchable.
	//
	// - FAILED: processing failed.
	//
	// - DELETING: deletion in progress.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the document was last updated.
	//
	// example:
	//
	// 2026-08-24T10:00:00Z
	UpdatedAt *string `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
}

func (s KnowledgeBaseDocument) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseDocument) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseDocument) GetChunkConfiguration() *KnowledgeBaseDocumentChunkConfiguration {
	return s.ChunkConfiguration
}

func (s *KnowledgeBaseDocument) GetChunkCount() *int32 {
	return s.ChunkCount
}

func (s *KnowledgeBaseDocument) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *KnowledgeBaseDocument) GetDocumentId() *string {
	return s.DocumentId
}

func (s *KnowledgeBaseDocument) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *KnowledgeBaseDocument) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *KnowledgeBaseDocument) GetFileName() *string {
	return s.FileName
}

func (s *KnowledgeBaseDocument) GetFileSize() *int64 {
	return s.FileSize
}

func (s *KnowledgeBaseDocument) GetMetadata() []*KnowledgeBaseDocumentMetadata {
	return s.Metadata
}

func (s *KnowledgeBaseDocument) GetSourceModifiedTime() *int64 {
	return s.SourceModifiedTime
}

func (s *KnowledgeBaseDocument) GetSourceType() *string {
	return s.SourceType
}

func (s *KnowledgeBaseDocument) GetSourceUri() *string {
	return s.SourceUri
}

func (s *KnowledgeBaseDocument) GetStatus() *string {
	return s.Status
}

func (s *KnowledgeBaseDocument) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *KnowledgeBaseDocument) SetChunkConfiguration(v *KnowledgeBaseDocumentChunkConfiguration) *KnowledgeBaseDocument {
	s.ChunkConfiguration = v
	return s
}

func (s *KnowledgeBaseDocument) SetChunkCount(v int32) *KnowledgeBaseDocument {
	s.ChunkCount = &v
	return s
}

func (s *KnowledgeBaseDocument) SetCreatedAt(v string) *KnowledgeBaseDocument {
	s.CreatedAt = &v
	return s
}

func (s *KnowledgeBaseDocument) SetDocumentId(v string) *KnowledgeBaseDocument {
	s.DocumentId = &v
	return s
}

func (s *KnowledgeBaseDocument) SetErrorCode(v string) *KnowledgeBaseDocument {
	s.ErrorCode = &v
	return s
}

func (s *KnowledgeBaseDocument) SetErrorMessage(v string) *KnowledgeBaseDocument {
	s.ErrorMessage = &v
	return s
}

func (s *KnowledgeBaseDocument) SetFileName(v string) *KnowledgeBaseDocument {
	s.FileName = &v
	return s
}

func (s *KnowledgeBaseDocument) SetFileSize(v int64) *KnowledgeBaseDocument {
	s.FileSize = &v
	return s
}

func (s *KnowledgeBaseDocument) SetMetadata(v []*KnowledgeBaseDocumentMetadata) *KnowledgeBaseDocument {
	s.Metadata = v
	return s
}

func (s *KnowledgeBaseDocument) SetSourceModifiedTime(v int64) *KnowledgeBaseDocument {
	s.SourceModifiedTime = &v
	return s
}

func (s *KnowledgeBaseDocument) SetSourceType(v string) *KnowledgeBaseDocument {
	s.SourceType = &v
	return s
}

func (s *KnowledgeBaseDocument) SetSourceUri(v string) *KnowledgeBaseDocument {
	s.SourceUri = &v
	return s
}

func (s *KnowledgeBaseDocument) SetStatus(v string) *KnowledgeBaseDocument {
	s.Status = &v
	return s
}

func (s *KnowledgeBaseDocument) SetUpdatedAt(v string) *KnowledgeBaseDocument {
	s.UpdatedAt = &v
	return s
}

func (s *KnowledgeBaseDocument) Validate() error {
	if s.ChunkConfiguration != nil {
		if err := s.ChunkConfiguration.Validate(); err != nil {
			return err
		}
	}
	if s.Metadata != nil {
		for _, item := range s.Metadata {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type KnowledgeBaseDocumentChunkConfiguration struct {
	// The heading level (1 to 6) used for splitting in the BY_HEADING strategy. Headings at or above this level serve as split boundaries. Deeper-level headings are retained in the chunk body.
	//
	// example:
	//
	// 2
	HeadingLevel *int32 `json:"HeadingLevel,omitempty" xml:"HeadingLevel,omitempty"`
	// The maximum character length of a single chunk. Starting from revision 22, this value is character-based. Valid values: 1 to 6000.
	//
	// example:
	//
	// 600
	MaxChunkSize *int32 `json:"MaxChunkSize,omitempty" xml:"MaxChunkSize,omitempty"`
	// The overlap character length between adjacent chunks. This parameter takes effect only for the BY_LENGTH strategy. If the value is greater than 0, the beginning of the next chunk repeats the content from the end of the previous chunk within this window. The overlap does not cause a chunk to exceed MaxChunkSize. A value of 0 indicates no overlap.
	//
	// example:
	//
	// 40
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// The snapshot of preprocessing rules.
	PreprocessRules *KnowledgeBaseDocumentChunkConfigurationPreprocessRules `json:"PreprocessRules,omitempty" xml:"PreprocessRules,omitempty" type:"Struct"`
	// The separator used in the BY_SEPARATOR strategy. The separator is matched as a literal string (not a regular expression). The maximum length is 32 characters.
	//
	// example:
	//
	// \\\\n\\\\n
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The chunking strategy. Valid values:
	//
	// - AUTO: intelligent splitting (heading-aware + paragraph packing).
	//
	// - BY_LENGTH: sliding window splitting by length. You can specify OverlapSize.
	//
	// - BY_SEPARATOR: splitting by separator. You must specify Separator.
	//
	// - BY_HEADING: splitting by heading level. You must specify HeadingLevel.
	//
	// example:
	//
	// BY_SEPARATOR
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s KnowledgeBaseDocumentChunkConfiguration) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseDocumentChunkConfiguration) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseDocumentChunkConfiguration) GetHeadingLevel() *int32 {
	return s.HeadingLevel
}

func (s *KnowledgeBaseDocumentChunkConfiguration) GetMaxChunkSize() *int32 {
	return s.MaxChunkSize
}

func (s *KnowledgeBaseDocumentChunkConfiguration) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *KnowledgeBaseDocumentChunkConfiguration) GetPreprocessRules() *KnowledgeBaseDocumentChunkConfigurationPreprocessRules {
	return s.PreprocessRules
}

func (s *KnowledgeBaseDocumentChunkConfiguration) GetSeparator() *string {
	return s.Separator
}

func (s *KnowledgeBaseDocumentChunkConfiguration) GetStrategy() *string {
	return s.Strategy
}

func (s *KnowledgeBaseDocumentChunkConfiguration) SetHeadingLevel(v int32) *KnowledgeBaseDocumentChunkConfiguration {
	s.HeadingLevel = &v
	return s
}

func (s *KnowledgeBaseDocumentChunkConfiguration) SetMaxChunkSize(v int32) *KnowledgeBaseDocumentChunkConfiguration {
	s.MaxChunkSize = &v
	return s
}

func (s *KnowledgeBaseDocumentChunkConfiguration) SetOverlapSize(v int32) *KnowledgeBaseDocumentChunkConfiguration {
	s.OverlapSize = &v
	return s
}

func (s *KnowledgeBaseDocumentChunkConfiguration) SetPreprocessRules(v *KnowledgeBaseDocumentChunkConfigurationPreprocessRules) *KnowledgeBaseDocumentChunkConfiguration {
	s.PreprocessRules = v
	return s
}

func (s *KnowledgeBaseDocumentChunkConfiguration) SetSeparator(v string) *KnowledgeBaseDocumentChunkConfiguration {
	s.Separator = &v
	return s
}

func (s *KnowledgeBaseDocumentChunkConfiguration) SetStrategy(v string) *KnowledgeBaseDocumentChunkConfiguration {
	s.Strategy = &v
	return s
}

func (s *KnowledgeBaseDocumentChunkConfiguration) Validate() error {
	if s.PreprocessRules != nil {
		if err := s.PreprocessRules.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type KnowledgeBaseDocumentChunkConfigurationPreprocessRules struct {
	// Specifies whether to remove URLs and email addresses during parsing.
	//
	// example:
	//
	// false
	RemoveUrlsAndEmails *bool `json:"RemoveUrlsAndEmails,omitempty" xml:"RemoveUrlsAndEmails,omitempty"`
	// Specifies whether to replace consecutive whitespace characters (spaces, line breaks, and tab characters) with a single space.
	//
	// example:
	//
	// true
	ReplaceConsecutiveWhitespace *bool `json:"ReplaceConsecutiveWhitespace,omitempty" xml:"ReplaceConsecutiveWhitespace,omitempty"`
}

func (s KnowledgeBaseDocumentChunkConfigurationPreprocessRules) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseDocumentChunkConfigurationPreprocessRules) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseDocumentChunkConfigurationPreprocessRules) GetRemoveUrlsAndEmails() *bool {
	return s.RemoveUrlsAndEmails
}

func (s *KnowledgeBaseDocumentChunkConfigurationPreprocessRules) GetReplaceConsecutiveWhitespace() *bool {
	return s.ReplaceConsecutiveWhitespace
}

func (s *KnowledgeBaseDocumentChunkConfigurationPreprocessRules) SetRemoveUrlsAndEmails(v bool) *KnowledgeBaseDocumentChunkConfigurationPreprocessRules {
	s.RemoveUrlsAndEmails = &v
	return s
}

func (s *KnowledgeBaseDocumentChunkConfigurationPreprocessRules) SetReplaceConsecutiveWhitespace(v bool) *KnowledgeBaseDocumentChunkConfigurationPreprocessRules {
	s.ReplaceConsecutiveWhitespace = &v
	return s
}

func (s *KnowledgeBaseDocumentChunkConfigurationPreprocessRules) Validate() error {
	return dara.Validate(s)
}

type KnowledgeBaseDocumentMetadata struct {
	// The metadata field name.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The metadata field value.
	//
	// example:
	//
	// R&D
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s KnowledgeBaseDocumentMetadata) String() string {
	return dara.Prettify(s)
}

func (s KnowledgeBaseDocumentMetadata) GoString() string {
	return s.String()
}

func (s *KnowledgeBaseDocumentMetadata) GetKey() *string {
	return s.Key
}

func (s *KnowledgeBaseDocumentMetadata) GetValue() *string {
	return s.Value
}

func (s *KnowledgeBaseDocumentMetadata) SetKey(v string) *KnowledgeBaseDocumentMetadata {
	s.Key = &v
	return s
}

func (s *KnowledgeBaseDocumentMetadata) SetValue(v string) *KnowledgeBaseDocumentMetadata {
	s.Value = &v
	return s
}

func (s *KnowledgeBaseDocumentMetadata) Validate() error {
	return dara.Validate(s)
}
