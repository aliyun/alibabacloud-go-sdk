// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-oss/client"
	
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"io"
)

type AccessControlList struct {
	Grant *string `json:"Grant,omitempty" xml:"Grant,omitempty"`
}

func (s AccessControlList) String() string {
	return tea.Prettify(s)
}

func (s AccessControlList) GoString() string {
	return s.String()
}

func (s *AccessControlList) SetGrant(v string) *AccessControlList {
	s.Grant = &v
	return s
}

type AccessControlPolicy struct {
	AccessControlList *AccessControlList `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty"`
	Owner             *Owner             `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s AccessControlPolicy) String() string {
	return tea.Prettify(s)
}

func (s AccessControlPolicy) GoString() string {
	return s.String()
}

func (s *AccessControlPolicy) SetAccessControlList(v *AccessControlList) *AccessControlPolicy {
	s.AccessControlList = v
	return s
}

func (s *AccessControlPolicy) SetOwner(v *Owner) *AccessControlPolicy {
	s.Owner = v
	return s
}

type ApplyServerSideEncryptionByDefault struct {
	KMSDataEncryption *string `json:"KMSDataEncryption,omitempty" xml:"KMSDataEncryption,omitempty"`
	KMSMasterKeyID    *string `json:"KMSMasterKeyID,omitempty" xml:"KMSMasterKeyID,omitempty"`
	SSEAlgorithm      *string `json:"SSEAlgorithm,omitempty" xml:"SSEAlgorithm,omitempty"`
}

func (s ApplyServerSideEncryptionByDefault) String() string {
	return tea.Prettify(s)
}

func (s ApplyServerSideEncryptionByDefault) GoString() string {
	return s.String()
}

func (s *ApplyServerSideEncryptionByDefault) SetKMSDataEncryption(v string) *ApplyServerSideEncryptionByDefault {
	s.KMSDataEncryption = &v
	return s
}

func (s *ApplyServerSideEncryptionByDefault) SetKMSMasterKeyID(v string) *ApplyServerSideEncryptionByDefault {
	s.KMSMasterKeyID = &v
	return s
}

func (s *ApplyServerSideEncryptionByDefault) SetSSEAlgorithm(v string) *ApplyServerSideEncryptionByDefault {
	s.SSEAlgorithm = &v
	return s
}

type Bucket struct {
	CreationDate     *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	ExtranetEndpoint *string `json:"ExtranetEndpoint,omitempty" xml:"ExtranetEndpoint,omitempty"`
	IntranetEndpoint *string `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Location         *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Name             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Region           *string `json:"Region,omitempty" xml:"Region,omitempty"`
	StorageClass     *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s Bucket) String() string {
	return tea.Prettify(s)
}

func (s Bucket) GoString() string {
	return s.String()
}

func (s *Bucket) SetCreationDate(v string) *Bucket {
	s.CreationDate = &v
	return s
}

func (s *Bucket) SetExtranetEndpoint(v string) *Bucket {
	s.ExtranetEndpoint = &v
	return s
}

func (s *Bucket) SetIntranetEndpoint(v string) *Bucket {
	s.IntranetEndpoint = &v
	return s
}

func (s *Bucket) SetLocation(v string) *Bucket {
	s.Location = &v
	return s
}

func (s *Bucket) SetName(v string) *Bucket {
	s.Name = &v
	return s
}

func (s *Bucket) SetRegion(v string) *Bucket {
	s.Region = &v
	return s
}

func (s *Bucket) SetStorageClass(v string) *Bucket {
	s.StorageClass = &v
	return s
}

type BucketLoggingStatus struct {
	LoggingEnabled *LoggingEnabled `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
}

func (s BucketLoggingStatus) String() string {
	return tea.Prettify(s)
}

func (s BucketLoggingStatus) GoString() string {
	return s.String()
}

func (s *BucketLoggingStatus) SetLoggingEnabled(v *LoggingEnabled) *BucketLoggingStatus {
	s.LoggingEnabled = v
	return s
}

type CORSConfiguration struct {
	CORSRule     []*CORSRule `json:"CORSRule,omitempty" xml:"CORSRule,omitempty" type:"Repeated"`
	ResponseVary *bool       `json:"ResponseVary,omitempty" xml:"ResponseVary,omitempty"`
}

func (s CORSConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CORSConfiguration) GoString() string {
	return s.String()
}

func (s *CORSConfiguration) SetCORSRule(v []*CORSRule) *CORSConfiguration {
	s.CORSRule = v
	return s
}

func (s *CORSConfiguration) SetResponseVary(v bool) *CORSConfiguration {
	s.ResponseVary = &v
	return s
}

type CORSRule struct {
	AllowedHeader *string   `json:"AllowedHeader,omitempty" xml:"AllowedHeader,omitempty"`
	AllowedMethod []*string `json:"AllowedMethod,omitempty" xml:"AllowedMethod,omitempty" type:"Repeated"`
	AllowedOrigin []*string `json:"AllowedOrigin,omitempty" xml:"AllowedOrigin,omitempty" type:"Repeated"`
	ExposeHeader  []*string `json:"ExposeHeader,omitempty" xml:"ExposeHeader,omitempty" type:"Repeated"`
	MaxAgeSeconds *int64    `json:"MaxAgeSeconds,omitempty" xml:"MaxAgeSeconds,omitempty"`
}

func (s CORSRule) String() string {
	return tea.Prettify(s)
}

func (s CORSRule) GoString() string {
	return s.String()
}

func (s *CORSRule) SetAllowedHeader(v string) *CORSRule {
	s.AllowedHeader = &v
	return s
}

func (s *CORSRule) SetAllowedMethod(v []*string) *CORSRule {
	s.AllowedMethod = v
	return s
}

func (s *CORSRule) SetAllowedOrigin(v []*string) *CORSRule {
	s.AllowedOrigin = v
	return s
}

func (s *CORSRule) SetExposeHeader(v []*string) *CORSRule {
	s.ExposeHeader = v
	return s
}

func (s *CORSRule) SetMaxAgeSeconds(v int64) *CORSRule {
	s.MaxAgeSeconds = &v
	return s
}

type CSVInput struct {
	AllowQuotedRecordDelimiter *bool   `json:"AllowQuotedRecordDelimiter,omitempty" xml:"AllowQuotedRecordDelimiter,omitempty"`
	CommentCharacter           *string `json:"CommentCharacter,omitempty" xml:"CommentCharacter,omitempty"`
	FieldDelimiter             *string `json:"FieldDelimiter,omitempty" xml:"FieldDelimiter,omitempty"`
	FileHeaderInfo             *string `json:"FileHeaderInfo,omitempty" xml:"FileHeaderInfo,omitempty"`
	QuoteCharacter             *string `json:"QuoteCharacter,omitempty" xml:"QuoteCharacter,omitempty"`
	Range                      *string `json:"Range,omitempty" xml:"Range,omitempty"`
	RecordDelimiter            *string `json:"RecordDelimiter,omitempty" xml:"RecordDelimiter,omitempty"`
}

func (s CSVInput) String() string {
	return tea.Prettify(s)
}

func (s CSVInput) GoString() string {
	return s.String()
}

func (s *CSVInput) SetAllowQuotedRecordDelimiter(v bool) *CSVInput {
	s.AllowQuotedRecordDelimiter = &v
	return s
}

func (s *CSVInput) SetCommentCharacter(v string) *CSVInput {
	s.CommentCharacter = &v
	return s
}

func (s *CSVInput) SetFieldDelimiter(v string) *CSVInput {
	s.FieldDelimiter = &v
	return s
}

func (s *CSVInput) SetFileHeaderInfo(v string) *CSVInput {
	s.FileHeaderInfo = &v
	return s
}

func (s *CSVInput) SetQuoteCharacter(v string) *CSVInput {
	s.QuoteCharacter = &v
	return s
}

func (s *CSVInput) SetRange(v string) *CSVInput {
	s.Range = &v
	return s
}

func (s *CSVInput) SetRecordDelimiter(v string) *CSVInput {
	s.RecordDelimiter = &v
	return s
}

type CSVOutput struct {
	FieldDelimiter  *string `json:"FieldDelimiter,omitempty" xml:"FieldDelimiter,omitempty"`
	RecordDelimiter *string `json:"RecordDelimiter,omitempty" xml:"RecordDelimiter,omitempty"`
}

func (s CSVOutput) String() string {
	return tea.Prettify(s)
}

func (s CSVOutput) GoString() string {
	return s.String()
}

func (s *CSVOutput) SetFieldDelimiter(v string) *CSVOutput {
	s.FieldDelimiter = &v
	return s
}

func (s *CSVOutput) SetRecordDelimiter(v string) *CSVOutput {
	s.RecordDelimiter = &v
	return s
}

type CommonPrefix struct {
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s CommonPrefix) String() string {
	return tea.Prettify(s)
}

func (s CommonPrefix) GoString() string {
	return s.String()
}

func (s *CommonPrefix) SetPrefix(v string) *CommonPrefix {
	s.Prefix = &v
	return s
}

type CompleteMultipartUpload struct {
	Part []*Part `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
}

func (s CompleteMultipartUpload) String() string {
	return tea.Prettify(s)
}

func (s CompleteMultipartUpload) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUpload) SetPart(v []*Part) *CompleteMultipartUpload {
	s.Part = v
	return s
}

type CopyObjectResult struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s CopyObjectResult) String() string {
	return tea.Prettify(s)
}

func (s CopyObjectResult) GoString() string {
	return s.String()
}

func (s *CopyObjectResult) SetETag(v string) *CopyObjectResult {
	s.ETag = &v
	return s
}

func (s *CopyObjectResult) SetLastModified(v string) *CopyObjectResult {
	s.LastModified = &v
	return s
}

type CopyPartResult struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s CopyPartResult) String() string {
	return tea.Prettify(s)
}

func (s CopyPartResult) GoString() string {
	return s.String()
}

func (s *CopyPartResult) SetETag(v string) *CopyPartResult {
	s.ETag = &v
	return s
}

func (s *CopyPartResult) SetLastModified(v string) *CopyPartResult {
	s.LastModified = &v
	return s
}

type CreateBucketConfiguration struct {
	DataRedundancyType *string `json:"DataRedundancyType,omitempty" xml:"DataRedundancyType,omitempty"`
	StorageClass       *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s CreateBucketConfiguration) String() string {
	return tea.Prettify(s)
}

func (s CreateBucketConfiguration) GoString() string {
	return s.String()
}

func (s *CreateBucketConfiguration) SetDataRedundancyType(v string) *CreateBucketConfiguration {
	s.DataRedundancyType = &v
	return s
}

func (s *CreateBucketConfiguration) SetStorageClass(v string) *CreateBucketConfiguration {
	s.StorageClass = &v
	return s
}

type Delete struct {
	Objects []*ObjectIdentifier `json:"Object,omitempty" xml:"Object,omitempty" type:"Repeated"`
	Quiet   *bool               `json:"Quiet,omitempty" xml:"Quiet,omitempty"`
}

func (s Delete) String() string {
	return tea.Prettify(s)
}

func (s Delete) GoString() string {
	return s.String()
}

func (s *Delete) SetObjects(v []*ObjectIdentifier) *Delete {
	s.Objects = v
	return s
}

func (s *Delete) SetQuiet(v bool) *Delete {
	s.Quiet = &v
	return s
}

type DeleteMarkerEntry struct {
	IsLatest     *bool   `json:"IsLatest,omitempty" xml:"IsLatest,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Owner        *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeleteMarkerEntry) String() string {
	return tea.Prettify(s)
}

func (s DeleteMarkerEntry) GoString() string {
	return s.String()
}

func (s *DeleteMarkerEntry) SetIsLatest(v bool) *DeleteMarkerEntry {
	s.IsLatest = &v
	return s
}

func (s *DeleteMarkerEntry) SetKey(v string) *DeleteMarkerEntry {
	s.Key = &v
	return s
}

func (s *DeleteMarkerEntry) SetLastModified(v string) *DeleteMarkerEntry {
	s.LastModified = &v
	return s
}

func (s *DeleteMarkerEntry) SetOwner(v *Owner) *DeleteMarkerEntry {
	s.Owner = v
	return s
}

func (s *DeleteMarkerEntry) SetVersionId(v string) *DeleteMarkerEntry {
	s.VersionId = &v
	return s
}

type DeletedObject struct {
	DeleteMarker          *bool   `json:"DeleteMarker,omitempty" xml:"DeleteMarker,omitempty"`
	DeleteMarkerVersionId *string `json:"DeleteMarkerVersionId,omitempty" xml:"DeleteMarkerVersionId,omitempty"`
	Key                   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	VersionId             *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s DeletedObject) String() string {
	return tea.Prettify(s)
}

func (s DeletedObject) GoString() string {
	return s.String()
}

func (s *DeletedObject) SetDeleteMarker(v bool) *DeletedObject {
	s.DeleteMarker = &v
	return s
}

func (s *DeletedObject) SetDeleteMarkerVersionId(v string) *DeletedObject {
	s.DeleteMarkerVersionId = &v
	return s
}

func (s *DeletedObject) SetKey(v string) *DeletedObject {
	s.Key = &v
	return s
}

func (s *DeletedObject) SetVersionId(v string) *DeletedObject {
	s.VersionId = &v
	return s
}

type Error struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HostId    *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s Error) String() string {
	return tea.Prettify(s)
}

func (s Error) GoString() string {
	return s.String()
}

func (s *Error) SetCode(v string) *Error {
	s.Code = &v
	return s
}

func (s *Error) SetHostId(v string) *Error {
	s.HostId = &v
	return s
}

func (s *Error) SetMessage(v string) *Error {
	s.Message = &v
	return s
}

func (s *Error) SetRequestId(v string) *Error {
	s.RequestId = &v
	return s
}

type ErrorDocument struct {
	HttpStatus *string `json:"HttpStatus,omitempty" xml:"HttpStatus,omitempty"`
	Key        *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s ErrorDocument) String() string {
	return tea.Prettify(s)
}

func (s ErrorDocument) GoString() string {
	return s.String()
}

func (s *ErrorDocument) SetHttpStatus(v string) *ErrorDocument {
	s.HttpStatus = &v
	return s
}

func (s *ErrorDocument) SetKey(v string) *ErrorDocument {
	s.Key = &v
	return s
}

type ExtendWormConfiguration struct {
	RetentionPeriodInDays *int32 `json:"RetentionPeriodInDays,omitempty" xml:"RetentionPeriodInDays,omitempty"`
}

func (s ExtendWormConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ExtendWormConfiguration) GoString() string {
	return s.String()
}

func (s *ExtendWormConfiguration) SetRetentionPeriodInDays(v int32) *ExtendWormConfiguration {
	s.RetentionPeriodInDays = &v
	return s
}

type IndexDocument struct {
	Suffix        *string `json:"Suffix,omitempty" xml:"Suffix,omitempty"`
	SupportSubDir *bool   `json:"SupportSubDir,omitempty" xml:"SupportSubDir,omitempty"`
	Type          *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s IndexDocument) String() string {
	return tea.Prettify(s)
}

func (s IndexDocument) GoString() string {
	return s.String()
}

func (s *IndexDocument) SetSuffix(v string) *IndexDocument {
	s.Suffix = &v
	return s
}

func (s *IndexDocument) SetSupportSubDir(v bool) *IndexDocument {
	s.SupportSubDir = &v
	return s
}

func (s *IndexDocument) SetType(v string) *IndexDocument {
	s.Type = &v
	return s
}

type InitiateWormConfiguration struct {
	RetentionPeriodInDays *int32 `json:"RetentionPeriodInDays,omitempty" xml:"RetentionPeriodInDays,omitempty"`
}

func (s InitiateWormConfiguration) String() string {
	return tea.Prettify(s)
}

func (s InitiateWormConfiguration) GoString() string {
	return s.String()
}

func (s *InitiateWormConfiguration) SetRetentionPeriodInDays(v int32) *InitiateWormConfiguration {
	s.RetentionPeriodInDays = &v
	return s
}

type InputSerialization struct {
	Csv             *CSVInput  `json:"CSV,omitempty" xml:"CSV,omitempty"`
	CompressionType *string    `json:"CompressionType,omitempty" xml:"CompressionType,omitempty"`
	Json            *JSONInput `json:"JSON,omitempty" xml:"JSON,omitempty"`
}

func (s InputSerialization) String() string {
	return tea.Prettify(s)
}

func (s InputSerialization) GoString() string {
	return s.String()
}

func (s *InputSerialization) SetCsv(v *CSVInput) *InputSerialization {
	s.Csv = v
	return s
}

func (s *InputSerialization) SetCompressionType(v string) *InputSerialization {
	s.CompressionType = &v
	return s
}

func (s *InputSerialization) SetJson(v *JSONInput) *InputSerialization {
	s.Json = v
	return s
}

type InventoryConfiguration struct {
	Destination            *InventoryDestination                 `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Filter                 *InventoryFilter                      `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Id                     *string                               `json:"Id,omitempty" xml:"Id,omitempty"`
	IncludedObjectVersions *string                               `json:"IncludedObjectVersions,omitempty" xml:"IncludedObjectVersions,omitempty"`
	IsEnabled              *bool                                 `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	OptionalFields         *InventoryConfigurationOptionalFields `json:"OptionalFields,omitempty" xml:"OptionalFields,omitempty" type:"Struct"`
	Schedule               *InventorySchedule                    `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s InventoryConfiguration) String() string {
	return tea.Prettify(s)
}

func (s InventoryConfiguration) GoString() string {
	return s.String()
}

func (s *InventoryConfiguration) SetDestination(v *InventoryDestination) *InventoryConfiguration {
	s.Destination = v
	return s
}

func (s *InventoryConfiguration) SetFilter(v *InventoryFilter) *InventoryConfiguration {
	s.Filter = v
	return s
}

func (s *InventoryConfiguration) SetId(v string) *InventoryConfiguration {
	s.Id = &v
	return s
}

func (s *InventoryConfiguration) SetIncludedObjectVersions(v string) *InventoryConfiguration {
	s.IncludedObjectVersions = &v
	return s
}

func (s *InventoryConfiguration) SetIsEnabled(v bool) *InventoryConfiguration {
	s.IsEnabled = &v
	return s
}

func (s *InventoryConfiguration) SetOptionalFields(v *InventoryConfigurationOptionalFields) *InventoryConfiguration {
	s.OptionalFields = v
	return s
}

func (s *InventoryConfiguration) SetSchedule(v *InventorySchedule) *InventoryConfiguration {
	s.Schedule = v
	return s
}

type InventoryConfigurationOptionalFields struct {
	Fields []*string `json:"Field,omitempty" xml:"Field,omitempty" type:"Repeated"`
}

func (s InventoryConfigurationOptionalFields) String() string {
	return tea.Prettify(s)
}

func (s InventoryConfigurationOptionalFields) GoString() string {
	return s.String()
}

func (s *InventoryConfigurationOptionalFields) SetFields(v []*string) *InventoryConfigurationOptionalFields {
	s.Fields = v
	return s
}

type InventoryDestination struct {
	OSSBucketDestination *InventoryOSSBucketDestination `json:"OSSBucketDestination,omitempty" xml:"OSSBucketDestination,omitempty"`
}

func (s InventoryDestination) String() string {
	return tea.Prettify(s)
}

func (s InventoryDestination) GoString() string {
	return s.String()
}

func (s *InventoryDestination) SetOSSBucketDestination(v *InventoryOSSBucketDestination) *InventoryDestination {
	s.OSSBucketDestination = v
	return s
}

type InventoryEncryption struct {
	SSEKMS *SSEKMS `json:"SSE-KMS,omitempty" xml:"SSE-KMS,omitempty"`
	SSEOSS *SSEOSS `json:"SSE-OSS,omitempty" xml:"SSE-OSS,omitempty"`
}

func (s InventoryEncryption) String() string {
	return tea.Prettify(s)
}

func (s InventoryEncryption) GoString() string {
	return s.String()
}

func (s *InventoryEncryption) SetSSEKMS(v *SSEKMS) *InventoryEncryption {
	s.SSEKMS = v
	return s
}

func (s *InventoryEncryption) SetSSEOSS(v *SSEOSS) *InventoryEncryption {
	s.SSEOSS = v
	return s
}

type InventoryFilter struct {
	Prefix *string `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s InventoryFilter) String() string {
	return tea.Prettify(s)
}

func (s InventoryFilter) GoString() string {
	return s.String()
}

func (s *InventoryFilter) SetPrefix(v string) *InventoryFilter {
	s.Prefix = &v
	return s
}

type InventoryOSSBucketDestination struct {
	AccountId  *string              `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	Bucket     *string              `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Encryption *InventoryEncryption `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	Format     *string              `json:"Format,omitempty" xml:"Format,omitempty"`
	Prefix     *string              `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	RoleArn    *string              `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
}

func (s InventoryOSSBucketDestination) String() string {
	return tea.Prettify(s)
}

func (s InventoryOSSBucketDestination) GoString() string {
	return s.String()
}

func (s *InventoryOSSBucketDestination) SetAccountId(v string) *InventoryOSSBucketDestination {
	s.AccountId = &v
	return s
}

func (s *InventoryOSSBucketDestination) SetBucket(v string) *InventoryOSSBucketDestination {
	s.Bucket = &v
	return s
}

func (s *InventoryOSSBucketDestination) SetEncryption(v *InventoryEncryption) *InventoryOSSBucketDestination {
	s.Encryption = v
	return s
}

func (s *InventoryOSSBucketDestination) SetFormat(v string) *InventoryOSSBucketDestination {
	s.Format = &v
	return s
}

func (s *InventoryOSSBucketDestination) SetPrefix(v string) *InventoryOSSBucketDestination {
	s.Prefix = &v
	return s
}

func (s *InventoryOSSBucketDestination) SetRoleArn(v string) *InventoryOSSBucketDestination {
	s.RoleArn = &v
	return s
}

type InventorySchedule struct {
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
}

func (s InventorySchedule) String() string {
	return tea.Prettify(s)
}

func (s InventorySchedule) GoString() string {
	return s.String()
}

func (s *InventorySchedule) SetFrequency(v string) *InventorySchedule {
	s.Frequency = &v
	return s
}

type JSONInput struct {
	ParseJsonNumberAsString *bool   `json:"ParseJsonNumberAsString,omitempty" xml:"ParseJsonNumberAsString,omitempty"`
	Range                   *string `json:"Range,omitempty" xml:"Range,omitempty"`
	Type                    *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s JSONInput) String() string {
	return tea.Prettify(s)
}

func (s JSONInput) GoString() string {
	return s.String()
}

func (s *JSONInput) SetParseJsonNumberAsString(v bool) *JSONInput {
	s.ParseJsonNumberAsString = &v
	return s
}

func (s *JSONInput) SetRange(v string) *JSONInput {
	s.Range = &v
	return s
}

func (s *JSONInput) SetType(v string) *JSONInput {
	s.Type = &v
	return s
}

type JSONOutput struct {
	RecordDelimiter *string `json:"RecordDelimiter,omitempty" xml:"RecordDelimiter,omitempty"`
}

func (s JSONOutput) String() string {
	return tea.Prettify(s)
}

func (s JSONOutput) GoString() string {
	return s.String()
}

func (s *JSONOutput) SetRecordDelimiter(v string) *JSONOutput {
	s.RecordDelimiter = &v
	return s
}

type LifecycleConfiguration struct {
	Rule []*LifecycleRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s LifecycleConfiguration) String() string {
	return tea.Prettify(s)
}

func (s LifecycleConfiguration) GoString() string {
	return s.String()
}

func (s *LifecycleConfiguration) SetRule(v []*LifecycleRule) *LifecycleConfiguration {
	s.Rule = v
	return s
}

type LifecycleRule struct {
	LifecycleAbortMultipartUpload *LifecycleRuleLifecycleAbortMultipartUpload `json:"AbortMultipartUpload,omitempty" xml:"AbortMultipartUpload,omitempty" type:"Struct"`
	LifecycleExpiration           *LifecycleRuleLifecycleExpiration           `json:"Expiration,omitempty" xml:"Expiration,omitempty" type:"Struct"`
	ID                            *string                                     `json:"ID,omitempty" xml:"ID,omitempty"`
	NoncurrentVersionExpiration   *LifecycleRuleNoncurrentVersionExpiration   `json:"NoncurrentVersionExpiration,omitempty" xml:"NoncurrentVersionExpiration,omitempty" type:"Struct"`
	NoncurrentVersionTransition   []*LifecycleRuleNoncurrentVersionTransition `json:"NoncurrentVersionTransition,omitempty" xml:"NoncurrentVersionTransition,omitempty" type:"Repeated"`
	Prefix                        *string                                     `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Status                        *string                                     `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag                           []*LifecycleRuleTag                         `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	LifecycleTransition           []*LifecycleRuleLifecycleTransition         `json:"Transition,omitempty" xml:"Transition,omitempty" type:"Repeated"`
}

func (s LifecycleRule) String() string {
	return tea.Prettify(s)
}

func (s LifecycleRule) GoString() string {
	return s.String()
}

func (s *LifecycleRule) SetLifecycleAbortMultipartUpload(v *LifecycleRuleLifecycleAbortMultipartUpload) *LifecycleRule {
	s.LifecycleAbortMultipartUpload = v
	return s
}

func (s *LifecycleRule) SetLifecycleExpiration(v *LifecycleRuleLifecycleExpiration) *LifecycleRule {
	s.LifecycleExpiration = v
	return s
}

func (s *LifecycleRule) SetID(v string) *LifecycleRule {
	s.ID = &v
	return s
}

func (s *LifecycleRule) SetNoncurrentVersionExpiration(v *LifecycleRuleNoncurrentVersionExpiration) *LifecycleRule {
	s.NoncurrentVersionExpiration = v
	return s
}

func (s *LifecycleRule) SetNoncurrentVersionTransition(v []*LifecycleRuleNoncurrentVersionTransition) *LifecycleRule {
	s.NoncurrentVersionTransition = v
	return s
}

func (s *LifecycleRule) SetPrefix(v string) *LifecycleRule {
	s.Prefix = &v
	return s
}

func (s *LifecycleRule) SetStatus(v string) *LifecycleRule {
	s.Status = &v
	return s
}

func (s *LifecycleRule) SetTag(v []*LifecycleRuleTag) *LifecycleRule {
	s.Tag = v
	return s
}

func (s *LifecycleRule) SetLifecycleTransition(v []*LifecycleRuleLifecycleTransition) *LifecycleRule {
	s.LifecycleTransition = v
	return s
}

type LifecycleRuleLifecycleAbortMultipartUpload struct {
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	Days              *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
}

func (s LifecycleRuleLifecycleAbortMultipartUpload) String() string {
	return tea.Prettify(s)
}

func (s LifecycleRuleLifecycleAbortMultipartUpload) GoString() string {
	return s.String()
}

func (s *LifecycleRuleLifecycleAbortMultipartUpload) SetCreatedBeforeDate(v string) *LifecycleRuleLifecycleAbortMultipartUpload {
	s.CreatedBeforeDate = &v
	return s
}

func (s *LifecycleRuleLifecycleAbortMultipartUpload) SetDays(v int32) *LifecycleRuleLifecycleAbortMultipartUpload {
	s.Days = &v
	return s
}

type LifecycleRuleLifecycleExpiration struct {
	CreatedBeforeDate         *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	Days                      *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	ExpiredObjectDeleteMarker *bool   `json:"ExpiredObjectDeleteMarker,omitempty" xml:"ExpiredObjectDeleteMarker,omitempty"`
}

func (s LifecycleRuleLifecycleExpiration) String() string {
	return tea.Prettify(s)
}

func (s LifecycleRuleLifecycleExpiration) GoString() string {
	return s.String()
}

func (s *LifecycleRuleLifecycleExpiration) SetCreatedBeforeDate(v string) *LifecycleRuleLifecycleExpiration {
	s.CreatedBeforeDate = &v
	return s
}

func (s *LifecycleRuleLifecycleExpiration) SetDays(v int32) *LifecycleRuleLifecycleExpiration {
	s.Days = &v
	return s
}

func (s *LifecycleRuleLifecycleExpiration) SetExpiredObjectDeleteMarker(v bool) *LifecycleRuleLifecycleExpiration {
	s.ExpiredObjectDeleteMarker = &v
	return s
}

type LifecycleRuleNoncurrentVersionExpiration struct {
	NoncurrentDays *int32 `json:"NoncurrentDays,omitempty" xml:"NoncurrentDays,omitempty"`
}

func (s LifecycleRuleNoncurrentVersionExpiration) String() string {
	return tea.Prettify(s)
}

func (s LifecycleRuleNoncurrentVersionExpiration) GoString() string {
	return s.String()
}

func (s *LifecycleRuleNoncurrentVersionExpiration) SetNoncurrentDays(v int32) *LifecycleRuleNoncurrentVersionExpiration {
	s.NoncurrentDays = &v
	return s
}

type LifecycleRuleNoncurrentVersionTransition struct {
	NoncurrentDays *int32  `json:"NoncurrentDays,omitempty" xml:"NoncurrentDays,omitempty"`
	StorageClass   *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s LifecycleRuleNoncurrentVersionTransition) String() string {
	return tea.Prettify(s)
}

func (s LifecycleRuleNoncurrentVersionTransition) GoString() string {
	return s.String()
}

func (s *LifecycleRuleNoncurrentVersionTransition) SetNoncurrentDays(v int32) *LifecycleRuleNoncurrentVersionTransition {
	s.NoncurrentDays = &v
	return s
}

func (s *LifecycleRuleNoncurrentVersionTransition) SetStorageClass(v string) *LifecycleRuleNoncurrentVersionTransition {
	s.StorageClass = &v
	return s
}

type LifecycleRuleTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s LifecycleRuleTag) String() string {
	return tea.Prettify(s)
}

func (s LifecycleRuleTag) GoString() string {
	return s.String()
}

func (s *LifecycleRuleTag) SetKey(v string) *LifecycleRuleTag {
	s.Key = &v
	return s
}

func (s *LifecycleRuleTag) SetValue(v string) *LifecycleRuleTag {
	s.Value = &v
	return s
}

type LifecycleRuleLifecycleTransition struct {
	CreatedBeforeDate *string `json:"CreatedBeforeDate,omitempty" xml:"CreatedBeforeDate,omitempty"`
	Days              *int32  `json:"Days,omitempty" xml:"Days,omitempty"`
	StorageClass      *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
}

func (s LifecycleRuleLifecycleTransition) String() string {
	return tea.Prettify(s)
}

func (s LifecycleRuleLifecycleTransition) GoString() string {
	return s.String()
}

func (s *LifecycleRuleLifecycleTransition) SetCreatedBeforeDate(v string) *LifecycleRuleLifecycleTransition {
	s.CreatedBeforeDate = &v
	return s
}

func (s *LifecycleRuleLifecycleTransition) SetDays(v int32) *LifecycleRuleLifecycleTransition {
	s.Days = &v
	return s
}

func (s *LifecycleRuleLifecycleTransition) SetStorageClass(v string) *LifecycleRuleLifecycleTransition {
	s.StorageClass = &v
	return s
}

type LiveChannel struct {
	Description  *string                 `json:"Description,omitempty" xml:"Description,omitempty"`
	LastModified *string                 `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Name         *string                 `json:"Name,omitempty" xml:"Name,omitempty"`
	PlayUrls     *LiveChannelPlayUrls    `json:"PlayUrls,omitempty" xml:"PlayUrls,omitempty"`
	PublishUrls  *LiveChannelPublishUrls `json:"PublishUrls,omitempty" xml:"PublishUrls,omitempty"`
	Status       *string                 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s LiveChannel) String() string {
	return tea.Prettify(s)
}

func (s LiveChannel) GoString() string {
	return s.String()
}

func (s *LiveChannel) SetDescription(v string) *LiveChannel {
	s.Description = &v
	return s
}

func (s *LiveChannel) SetLastModified(v string) *LiveChannel {
	s.LastModified = &v
	return s
}

func (s *LiveChannel) SetName(v string) *LiveChannel {
	s.Name = &v
	return s
}

func (s *LiveChannel) SetPlayUrls(v *LiveChannelPlayUrls) *LiveChannel {
	s.PlayUrls = v
	return s
}

func (s *LiveChannel) SetPublishUrls(v *LiveChannelPublishUrls) *LiveChannel {
	s.PublishUrls = v
	return s
}

func (s *LiveChannel) SetStatus(v string) *LiveChannel {
	s.Status = &v
	return s
}

type LiveChannelAudio struct {
	Bandwidth  *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Codec      *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	SampleRate *int64  `json:"SampleRate,omitempty" xml:"SampleRate,omitempty"`
}

func (s LiveChannelAudio) String() string {
	return tea.Prettify(s)
}

func (s LiveChannelAudio) GoString() string {
	return s.String()
}

func (s *LiveChannelAudio) SetBandwidth(v int64) *LiveChannelAudio {
	s.Bandwidth = &v
	return s
}

func (s *LiveChannelAudio) SetCodec(v string) *LiveChannelAudio {
	s.Codec = &v
	return s
}

func (s *LiveChannelAudio) SetSampleRate(v int64) *LiveChannelAudio {
	s.SampleRate = &v
	return s
}

type LiveChannelConfiguration struct {
	Description *string              `json:"Description,omitempty" xml:"Description,omitempty"`
	Snapshot    *LiveChannelSnapshot `json:"Snapshot,omitempty" xml:"Snapshot,omitempty"`
	Status      *string              `json:"Status,omitempty" xml:"Status,omitempty"`
	Target      *LiveChannelTarget   `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s LiveChannelConfiguration) String() string {
	return tea.Prettify(s)
}

func (s LiveChannelConfiguration) GoString() string {
	return s.String()
}

func (s *LiveChannelConfiguration) SetDescription(v string) *LiveChannelConfiguration {
	s.Description = &v
	return s
}

func (s *LiveChannelConfiguration) SetSnapshot(v *LiveChannelSnapshot) *LiveChannelConfiguration {
	s.Snapshot = v
	return s
}

func (s *LiveChannelConfiguration) SetStatus(v string) *LiveChannelConfiguration {
	s.Status = &v
	return s
}

func (s *LiveChannelConfiguration) SetTarget(v *LiveChannelTarget) *LiveChannelConfiguration {
	s.Target = v
	return s
}

type LiveChannelPlayUrls struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s LiveChannelPlayUrls) String() string {
	return tea.Prettify(s)
}

func (s LiveChannelPlayUrls) GoString() string {
	return s.String()
}

func (s *LiveChannelPlayUrls) SetUrl(v string) *LiveChannelPlayUrls {
	s.Url = &v
	return s
}

type LiveChannelPublishUrls struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s LiveChannelPublishUrls) String() string {
	return tea.Prettify(s)
}

func (s LiveChannelPublishUrls) GoString() string {
	return s.String()
}

func (s *LiveChannelPublishUrls) SetUrl(v string) *LiveChannelPublishUrls {
	s.Url = &v
	return s
}

type LiveChannelSnapshot struct {
	DestBucket  *string `json:"DestBucket,omitempty" xml:"DestBucket,omitempty"`
	Interval    *int64  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	NotifyTopic *string `json:"NotifyTopic,omitempty" xml:"NotifyTopic,omitempty"`
	RoleName    *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s LiveChannelSnapshot) String() string {
	return tea.Prettify(s)
}

func (s LiveChannelSnapshot) GoString() string {
	return s.String()
}

func (s *LiveChannelSnapshot) SetDestBucket(v string) *LiveChannelSnapshot {
	s.DestBucket = &v
	return s
}

func (s *LiveChannelSnapshot) SetInterval(v int64) *LiveChannelSnapshot {
	s.Interval = &v
	return s
}

func (s *LiveChannelSnapshot) SetNotifyTopic(v string) *LiveChannelSnapshot {
	s.NotifyTopic = &v
	return s
}

func (s *LiveChannelSnapshot) SetRoleName(v string) *LiveChannelSnapshot {
	s.RoleName = &v
	return s
}

type LiveChannelTarget struct {
	FragCount    *int64  `json:"FragCount,omitempty" xml:"FragCount,omitempty"`
	FragDuration *int64  `json:"FragDuration,omitempty" xml:"FragDuration,omitempty"`
	PlaylistName *string `json:"PlaylistName,omitempty" xml:"PlaylistName,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LiveChannelTarget) String() string {
	return tea.Prettify(s)
}

func (s LiveChannelTarget) GoString() string {
	return s.String()
}

func (s *LiveChannelTarget) SetFragCount(v int64) *LiveChannelTarget {
	s.FragCount = &v
	return s
}

func (s *LiveChannelTarget) SetFragDuration(v int64) *LiveChannelTarget {
	s.FragDuration = &v
	return s
}

func (s *LiveChannelTarget) SetPlaylistName(v string) *LiveChannelTarget {
	s.PlaylistName = &v
	return s
}

func (s *LiveChannelTarget) SetType(v string) *LiveChannelTarget {
	s.Type = &v
	return s
}

type LiveChannelVideo struct {
	Bandwidth *int64  `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	Codec     *string `json:"Codec,omitempty" xml:"Codec,omitempty"`
	FrameRate *int64  `json:"FrameRate,omitempty" xml:"FrameRate,omitempty"`
	Height    *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	Width     *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s LiveChannelVideo) String() string {
	return tea.Prettify(s)
}

func (s LiveChannelVideo) GoString() string {
	return s.String()
}

func (s *LiveChannelVideo) SetBandwidth(v int64) *LiveChannelVideo {
	s.Bandwidth = &v
	return s
}

func (s *LiveChannelVideo) SetCodec(v string) *LiveChannelVideo {
	s.Codec = &v
	return s
}

func (s *LiveChannelVideo) SetFrameRate(v int64) *LiveChannelVideo {
	s.FrameRate = &v
	return s
}

func (s *LiveChannelVideo) SetHeight(v int64) *LiveChannelVideo {
	s.Height = &v
	return s
}

func (s *LiveChannelVideo) SetWidth(v int64) *LiveChannelVideo {
	s.Width = &v
	return s
}

type LiveRecord struct {
	EndTime    *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	RemoteAddr *string `json:"RemoteAddr,omitempty" xml:"RemoteAddr,omitempty"`
	StartTime  *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s LiveRecord) String() string {
	return tea.Prettify(s)
}

func (s LiveRecord) GoString() string {
	return s.String()
}

func (s *LiveRecord) SetEndTime(v string) *LiveRecord {
	s.EndTime = &v
	return s
}

func (s *LiveRecord) SetRemoteAddr(v string) *LiveRecord {
	s.RemoteAddr = &v
	return s
}

func (s *LiveRecord) SetStartTime(v string) *LiveRecord {
	s.StartTime = &v
	return s
}

type LocationTransferType struct {
	Location      *string                            `json:"Location,omitempty" xml:"Location,omitempty"`
	TransferTypes *LocationTransferTypeTransferTypes `json:"TransferTypes,omitempty" xml:"TransferTypes,omitempty" type:"Struct"`
}

func (s LocationTransferType) String() string {
	return tea.Prettify(s)
}

func (s LocationTransferType) GoString() string {
	return s.String()
}

func (s *LocationTransferType) SetLocation(v string) *LocationTransferType {
	s.Location = &v
	return s
}

func (s *LocationTransferType) SetTransferTypes(v *LocationTransferTypeTransferTypes) *LocationTransferType {
	s.TransferTypes = v
	return s
}

type LocationTransferTypeTransferTypes struct {
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LocationTransferTypeTransferTypes) String() string {
	return tea.Prettify(s)
}

func (s LocationTransferTypeTransferTypes) GoString() string {
	return s.String()
}

func (s *LocationTransferTypeTransferTypes) SetType(v string) *LocationTransferTypeTransferTypes {
	s.Type = &v
	return s
}

type LoggingEnabled struct {
	TargetBucket *string `json:"TargetBucket,omitempty" xml:"TargetBucket,omitempty"`
	TargetPrefix *string `json:"TargetPrefix,omitempty" xml:"TargetPrefix,omitempty"`
}

func (s LoggingEnabled) String() string {
	return tea.Prettify(s)
}

func (s LoggingEnabled) GoString() string {
	return s.String()
}

func (s *LoggingEnabled) SetTargetBucket(v string) *LoggingEnabled {
	s.TargetBucket = &v
	return s
}

func (s *LoggingEnabled) SetTargetPrefix(v string) *LoggingEnabled {
	s.TargetPrefix = &v
	return s
}

type ObjectIdentifier struct {
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ObjectIdentifier) String() string {
	return tea.Prettify(s)
}

func (s ObjectIdentifier) GoString() string {
	return s.String()
}

func (s *ObjectIdentifier) SetKey(v string) *ObjectIdentifier {
	s.Key = &v
	return s
}

func (s *ObjectIdentifier) SetVersionId(v string) *ObjectIdentifier {
	s.VersionId = &v
	return s
}

type ObjectSummary struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Owner        *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	Type         *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ObjectSummary) String() string {
	return tea.Prettify(s)
}

func (s ObjectSummary) GoString() string {
	return s.String()
}

func (s *ObjectSummary) SetETag(v string) *ObjectSummary {
	s.ETag = &v
	return s
}

func (s *ObjectSummary) SetKey(v string) *ObjectSummary {
	s.Key = &v
	return s
}

func (s *ObjectSummary) SetLastModified(v string) *ObjectSummary {
	s.LastModified = &v
	return s
}

func (s *ObjectSummary) SetOwner(v *Owner) *ObjectSummary {
	s.Owner = v
	return s
}

func (s *ObjectSummary) SetSize(v int64) *ObjectSummary {
	s.Size = &v
	return s
}

func (s *ObjectSummary) SetStorageClass(v string) *ObjectSummary {
	s.StorageClass = &v
	return s
}

func (s *ObjectSummary) SetType(v string) *ObjectSummary {
	s.Type = &v
	return s
}

type ObjectVersion struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	IsLatest     *bool   `json:"IsLatest,omitempty" xml:"IsLatest,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	Owner        *Owner  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	StorageClass *string `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	VersionId    *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s ObjectVersion) String() string {
	return tea.Prettify(s)
}

func (s ObjectVersion) GoString() string {
	return s.String()
}

func (s *ObjectVersion) SetETag(v string) *ObjectVersion {
	s.ETag = &v
	return s
}

func (s *ObjectVersion) SetIsLatest(v bool) *ObjectVersion {
	s.IsLatest = &v
	return s
}

func (s *ObjectVersion) SetKey(v string) *ObjectVersion {
	s.Key = &v
	return s
}

func (s *ObjectVersion) SetLastModified(v string) *ObjectVersion {
	s.LastModified = &v
	return s
}

func (s *ObjectVersion) SetOwner(v *Owner) *ObjectVersion {
	s.Owner = v
	return s
}

func (s *ObjectVersion) SetSize(v int64) *ObjectVersion {
	s.Size = &v
	return s
}

func (s *ObjectVersion) SetStorageClass(v string) *ObjectVersion {
	s.StorageClass = &v
	return s
}

func (s *ObjectVersion) SetVersionId(v string) *ObjectVersion {
	s.VersionId = &v
	return s
}

type OutputSerialization struct {
	Csv              *CSVOutput  `json:"CSV,omitempty" xml:"CSV,omitempty"`
	EnablePayloadCrc *bool       `json:"EnablePayloadCrc,omitempty" xml:"EnablePayloadCrc,omitempty"`
	Json             *JSONOutput `json:"JSON,omitempty" xml:"JSON,omitempty"`
	KeepAllColumns   *bool       `json:"KeepAllColumns,omitempty" xml:"KeepAllColumns,omitempty"`
	OutputHeader     *bool       `json:"OutputHeader,omitempty" xml:"OutputHeader,omitempty"`
	OutputRawData    *bool       `json:"OutputRawData,omitempty" xml:"OutputRawData,omitempty"`
}

func (s OutputSerialization) String() string {
	return tea.Prettify(s)
}

func (s OutputSerialization) GoString() string {
	return s.String()
}

func (s *OutputSerialization) SetCsv(v *CSVOutput) *OutputSerialization {
	s.Csv = v
	return s
}

func (s *OutputSerialization) SetEnablePayloadCrc(v bool) *OutputSerialization {
	s.EnablePayloadCrc = &v
	return s
}

func (s *OutputSerialization) SetJson(v *JSONOutput) *OutputSerialization {
	s.Json = v
	return s
}

func (s *OutputSerialization) SetKeepAllColumns(v bool) *OutputSerialization {
	s.KeepAllColumns = &v
	return s
}

func (s *OutputSerialization) SetOutputHeader(v bool) *OutputSerialization {
	s.OutputHeader = &v
	return s
}

func (s *OutputSerialization) SetOutputRawData(v bool) *OutputSerialization {
	s.OutputRawData = &v
	return s
}

type Owner struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ID          *string `json:"ID,omitempty" xml:"ID,omitempty"`
}

func (s Owner) String() string {
	return tea.Prettify(s)
}

func (s Owner) GoString() string {
	return s.String()
}

func (s *Owner) SetDisplayName(v string) *Owner {
	s.DisplayName = &v
	return s
}

func (s *Owner) SetID(v string) *Owner {
	s.ID = &v
	return s
}

type Part struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
	PartNumber   *int64  `json:"PartNumber,omitempty" xml:"PartNumber,omitempty"`
	Size         *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s Part) String() string {
	return tea.Prettify(s)
}

func (s Part) GoString() string {
	return s.String()
}

func (s *Part) SetETag(v string) *Part {
	s.ETag = &v
	return s
}

func (s *Part) SetLastModified(v string) *Part {
	s.LastModified = &v
	return s
}

func (s *Part) SetPartNumber(v int64) *Part {
	s.PartNumber = &v
	return s
}

func (s *Part) SetSize(v int64) *Part {
	s.Size = &v
	return s
}

type RefererConfiguration struct {
	AllowEmptyReferer        *bool                            `json:"AllowEmptyReferer,omitempty" xml:"AllowEmptyReferer,omitempty"`
	AllowTruncateQueryString *bool                            `json:"AllowTruncateQueryString,omitempty" xml:"AllowTruncateQueryString,omitempty"`
	RefererList              *RefererConfigurationRefererList `json:"RefererList,omitempty" xml:"RefererList,omitempty" type:"Struct"`
}

func (s RefererConfiguration) String() string {
	return tea.Prettify(s)
}

func (s RefererConfiguration) GoString() string {
	return s.String()
}

func (s *RefererConfiguration) SetAllowEmptyReferer(v bool) *RefererConfiguration {
	s.AllowEmptyReferer = &v
	return s
}

func (s *RefererConfiguration) SetAllowTruncateQueryString(v bool) *RefererConfiguration {
	s.AllowTruncateQueryString = &v
	return s
}

func (s *RefererConfiguration) SetRefererList(v *RefererConfigurationRefererList) *RefererConfiguration {
	s.RefererList = v
	return s
}

type RefererConfigurationRefererList struct {
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
}

func (s RefererConfigurationRefererList) String() string {
	return tea.Prettify(s)
}

func (s RefererConfigurationRefererList) GoString() string {
	return s.String()
}

func (s *RefererConfigurationRefererList) SetReferer(v []*string) *RefererConfigurationRefererList {
	s.Referer = v
	return s
}

type RegionInfo struct {
	AccelerateEndpoint *string `json:"AccelerateEndpoint,omitempty" xml:"AccelerateEndpoint,omitempty"`
	InternalEndpoint   *string `json:"InternalEndpoint,omitempty" xml:"InternalEndpoint,omitempty"`
	InternetEndpoint   *string `json:"InternetEndpoint,omitempty" xml:"InternetEndpoint,omitempty"`
	Region             *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s RegionInfo) String() string {
	return tea.Prettify(s)
}

func (s RegionInfo) GoString() string {
	return s.String()
}

func (s *RegionInfo) SetAccelerateEndpoint(v string) *RegionInfo {
	s.AccelerateEndpoint = &v
	return s
}

func (s *RegionInfo) SetInternalEndpoint(v string) *RegionInfo {
	s.InternalEndpoint = &v
	return s
}

func (s *RegionInfo) SetInternetEndpoint(v string) *RegionInfo {
	s.InternetEndpoint = &v
	return s
}

func (s *RegionInfo) SetRegion(v string) *RegionInfo {
	s.Region = &v
	return s
}

type ReplicationConfiguration struct {
	Rule *ReplicationRule `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s ReplicationConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ReplicationConfiguration) GoString() string {
	return s.String()
}

func (s *ReplicationConfiguration) SetRule(v *ReplicationRule) *ReplicationConfiguration {
	s.Rule = v
	return s
}

type ReplicationDestination struct {
	Bucket       *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
	TransferType *string `json:"TransferType,omitempty" xml:"TransferType,omitempty"`
}

func (s ReplicationDestination) String() string {
	return tea.Prettify(s)
}

func (s ReplicationDestination) GoString() string {
	return s.String()
}

func (s *ReplicationDestination) SetBucket(v string) *ReplicationDestination {
	s.Bucket = &v
	return s
}

func (s *ReplicationDestination) SetLocation(v string) *ReplicationDestination {
	s.Location = &v
	return s
}

func (s *ReplicationDestination) SetTransferType(v string) *ReplicationDestination {
	s.TransferType = &v
	return s
}

type ReplicationPrefixSet struct {
	Prefixs []*string `json:"Prefix,omitempty" xml:"Prefix,omitempty" type:"Repeated"`
}

func (s ReplicationPrefixSet) String() string {
	return tea.Prettify(s)
}

func (s ReplicationPrefixSet) GoString() string {
	return s.String()
}

func (s *ReplicationPrefixSet) SetPrefixs(v []*string) *ReplicationPrefixSet {
	s.Prefixs = v
	return s
}

type ReplicationProgressRule struct {
	Action                      *string                          `json:"Action,omitempty" xml:"Action,omitempty"`
	Destination                 *ReplicationDestination          `json:"Destination,omitempty" xml:"Destination,omitempty"`
	HistoricalObjectReplication *string                          `json:"HistoricalObjectReplication,omitempty" xml:"HistoricalObjectReplication,omitempty"`
	ID                          *string                          `json:"ID,omitempty" xml:"ID,omitempty"`
	PrefixSet                   *ReplicationPrefixSet            `json:"PrefixSet,omitempty" xml:"PrefixSet,omitempty"`
	Progress                    *ReplicationProgressRuleProgress `json:"Progress,omitempty" xml:"Progress,omitempty" type:"Struct"`
	Status                      *string                          `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ReplicationProgressRule) String() string {
	return tea.Prettify(s)
}

func (s ReplicationProgressRule) GoString() string {
	return s.String()
}

func (s *ReplicationProgressRule) SetAction(v string) *ReplicationProgressRule {
	s.Action = &v
	return s
}

func (s *ReplicationProgressRule) SetDestination(v *ReplicationDestination) *ReplicationProgressRule {
	s.Destination = v
	return s
}

func (s *ReplicationProgressRule) SetHistoricalObjectReplication(v string) *ReplicationProgressRule {
	s.HistoricalObjectReplication = &v
	return s
}

func (s *ReplicationProgressRule) SetID(v string) *ReplicationProgressRule {
	s.ID = &v
	return s
}

func (s *ReplicationProgressRule) SetPrefixSet(v *ReplicationPrefixSet) *ReplicationProgressRule {
	s.PrefixSet = v
	return s
}

func (s *ReplicationProgressRule) SetProgress(v *ReplicationProgressRuleProgress) *ReplicationProgressRule {
	s.Progress = v
	return s
}

func (s *ReplicationProgressRule) SetStatus(v string) *ReplicationProgressRule {
	s.Status = &v
	return s
}

type ReplicationProgressRuleProgress struct {
	HistoricalObject *string `json:"HistoricalObject,omitempty" xml:"HistoricalObject,omitempty"`
	NewObject        *string `json:"NewObject,omitempty" xml:"NewObject,omitempty"`
}

func (s ReplicationProgressRuleProgress) String() string {
	return tea.Prettify(s)
}

func (s ReplicationProgressRuleProgress) GoString() string {
	return s.String()
}

func (s *ReplicationProgressRuleProgress) SetHistoricalObject(v string) *ReplicationProgressRuleProgress {
	s.HistoricalObject = &v
	return s
}

func (s *ReplicationProgressRuleProgress) SetNewObject(v string) *ReplicationProgressRuleProgress {
	s.NewObject = &v
	return s
}

type ReplicationRule struct {
	Action                      *string                                 `json:"Action,omitempty" xml:"Action,omitempty"`
	Destination                 *ReplicationDestination                 `json:"Destination,omitempty" xml:"Destination,omitempty"`
	EncryptionConfiguration     *ReplicationRuleEncryptionConfiguration `json:"EncryptionConfiguration,omitempty" xml:"EncryptionConfiguration,omitempty" type:"Struct"`
	HistoricalObjectReplication *string                                 `json:"HistoricalObjectReplication,omitempty" xml:"HistoricalObjectReplication,omitempty"`
	ID                          *string                                 `json:"ID,omitempty" xml:"ID,omitempty"`
	PrefixSet                   *ReplicationPrefixSet                   `json:"PrefixSet,omitempty" xml:"PrefixSet,omitempty"`
	SourceSelectionCriteria     *ReplicationSourceSelectionCriteria     `json:"SourceSelectionCriteria,omitempty" xml:"SourceSelectionCriteria,omitempty"`
	Status                      *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	SyncRole                    *string                                 `json:"SyncRole,omitempty" xml:"SyncRole,omitempty"`
}

func (s ReplicationRule) String() string {
	return tea.Prettify(s)
}

func (s ReplicationRule) GoString() string {
	return s.String()
}

func (s *ReplicationRule) SetAction(v string) *ReplicationRule {
	s.Action = &v
	return s
}

func (s *ReplicationRule) SetDestination(v *ReplicationDestination) *ReplicationRule {
	s.Destination = v
	return s
}

func (s *ReplicationRule) SetEncryptionConfiguration(v *ReplicationRuleEncryptionConfiguration) *ReplicationRule {
	s.EncryptionConfiguration = v
	return s
}

func (s *ReplicationRule) SetHistoricalObjectReplication(v string) *ReplicationRule {
	s.HistoricalObjectReplication = &v
	return s
}

func (s *ReplicationRule) SetID(v string) *ReplicationRule {
	s.ID = &v
	return s
}

func (s *ReplicationRule) SetPrefixSet(v *ReplicationPrefixSet) *ReplicationRule {
	s.PrefixSet = v
	return s
}

func (s *ReplicationRule) SetSourceSelectionCriteria(v *ReplicationSourceSelectionCriteria) *ReplicationRule {
	s.SourceSelectionCriteria = v
	return s
}

func (s *ReplicationRule) SetStatus(v string) *ReplicationRule {
	s.Status = &v
	return s
}

func (s *ReplicationRule) SetSyncRole(v string) *ReplicationRule {
	s.SyncRole = &v
	return s
}

type ReplicationRuleEncryptionConfiguration struct {
	ReplicaKmsKeyID *string `json:"ReplicaKmsKeyID,omitempty" xml:"ReplicaKmsKeyID,omitempty"`
}

func (s ReplicationRuleEncryptionConfiguration) String() string {
	return tea.Prettify(s)
}

func (s ReplicationRuleEncryptionConfiguration) GoString() string {
	return s.String()
}

func (s *ReplicationRuleEncryptionConfiguration) SetReplicaKmsKeyID(v string) *ReplicationRuleEncryptionConfiguration {
	s.ReplicaKmsKeyID = &v
	return s
}

type ReplicationRuleProgress struct {
	Action    *string               `json:"Action,omitempty" xml:"Action,omitempty"`
	ID        *string               `json:"ID,omitempty" xml:"ID,omitempty"`
	PrefixSet *ReplicationPrefixSet `json:"PrefixSet,omitempty" xml:"PrefixSet,omitempty"`
}

func (s ReplicationRuleProgress) String() string {
	return tea.Prettify(s)
}

func (s ReplicationRuleProgress) GoString() string {
	return s.String()
}

func (s *ReplicationRuleProgress) SetAction(v string) *ReplicationRuleProgress {
	s.Action = &v
	return s
}

func (s *ReplicationRuleProgress) SetID(v string) *ReplicationRuleProgress {
	s.ID = &v
	return s
}

func (s *ReplicationRuleProgress) SetPrefixSet(v *ReplicationPrefixSet) *ReplicationRuleProgress {
	s.PrefixSet = v
	return s
}

type ReplicationRules struct {
	Ids []*string `json:"ID,omitempty" xml:"ID,omitempty" type:"Repeated"`
}

func (s ReplicationRules) String() string {
	return tea.Prettify(s)
}

func (s ReplicationRules) GoString() string {
	return s.String()
}

func (s *ReplicationRules) SetIds(v []*string) *ReplicationRules {
	s.Ids = v
	return s
}

type ReplicationSourceSelectionCriteria struct {
	SseKmsEncryptedObjects *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects `json:"SseKmsEncryptedObjects,omitempty" xml:"SseKmsEncryptedObjects,omitempty" type:"Struct"`
}

func (s ReplicationSourceSelectionCriteria) String() string {
	return tea.Prettify(s)
}

func (s ReplicationSourceSelectionCriteria) GoString() string {
	return s.String()
}

func (s *ReplicationSourceSelectionCriteria) SetSseKmsEncryptedObjects(v *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) *ReplicationSourceSelectionCriteria {
	s.SseKmsEncryptedObjects = v
	return s
}

type ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) String() string {
	return tea.Prettify(s)
}

func (s ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) GoString() string {
	return s.String()
}

func (s *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects) SetStatus(v string) *ReplicationSourceSelectionCriteriaSseKmsEncryptedObjects {
	s.Status = &v
	return s
}

type RequestPaymentConfiguration struct {
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
}

func (s RequestPaymentConfiguration) String() string {
	return tea.Prettify(s)
}

func (s RequestPaymentConfiguration) GoString() string {
	return s.String()
}

func (s *RequestPaymentConfiguration) SetPayer(v string) *RequestPaymentConfiguration {
	s.Payer = &v
	return s
}

type RestoreRequest struct {
	Days          *int64                       `json:"Days,omitempty" xml:"Days,omitempty"`
	JobParameters *RestoreRequestJobParameters `json:"JobParameters,omitempty" xml:"JobParameters,omitempty" type:"Struct"`
}

func (s RestoreRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreRequest) GoString() string {
	return s.String()
}

func (s *RestoreRequest) SetDays(v int64) *RestoreRequest {
	s.Days = &v
	return s
}

func (s *RestoreRequest) SetJobParameters(v *RestoreRequestJobParameters) *RestoreRequest {
	s.JobParameters = v
	return s
}

type RestoreRequestJobParameters struct {
	Tier *string `json:"Tier,omitempty" xml:"Tier,omitempty"`
}

func (s RestoreRequestJobParameters) String() string {
	return tea.Prettify(s)
}

func (s RestoreRequestJobParameters) GoString() string {
	return s.String()
}

func (s *RestoreRequestJobParameters) SetTier(v string) *RestoreRequestJobParameters {
	s.Tier = &v
	return s
}

type RoutingRule struct {
	Condition  *RoutingRuleCondition `json:"Condition,omitempty" xml:"Condition,omitempty"`
	Redirect   *RoutingRuleRedirect  `json:"Redirect,omitempty" xml:"Redirect,omitempty"`
	RuleNumber *int64                `json:"RuleNumber,omitempty" xml:"RuleNumber,omitempty"`
}

func (s RoutingRule) String() string {
	return tea.Prettify(s)
}

func (s RoutingRule) GoString() string {
	return s.String()
}

func (s *RoutingRule) SetCondition(v *RoutingRuleCondition) *RoutingRule {
	s.Condition = v
	return s
}

func (s *RoutingRule) SetRedirect(v *RoutingRuleRedirect) *RoutingRule {
	s.Redirect = v
	return s
}

func (s *RoutingRule) SetRuleNumber(v int64) *RoutingRule {
	s.RuleNumber = &v
	return s
}

type RoutingRuleCondition struct {
	HttpErrorCodeReturnedEquals *int64  `json:"HttpErrorCodeReturnedEquals,omitempty" xml:"HttpErrorCodeReturnedEquals,omitempty"`
	KeyPrefixEquals             *string `json:"KeyPrefixEquals,omitempty" xml:"KeyPrefixEquals,omitempty"`
}

func (s RoutingRuleCondition) String() string {
	return tea.Prettify(s)
}

func (s RoutingRuleCondition) GoString() string {
	return s.String()
}

func (s *RoutingRuleCondition) SetHttpErrorCodeReturnedEquals(v int64) *RoutingRuleCondition {
	s.HttpErrorCodeReturnedEquals = &v
	return s
}

func (s *RoutingRuleCondition) SetKeyPrefixEquals(v string) *RoutingRuleCondition {
	s.KeyPrefixEquals = &v
	return s
}

type RoutingRuleRedirect struct {
	EnableReplacePrefix            *bool                             `json:"EnableReplacePrefix,omitempty" xml:"EnableReplacePrefix,omitempty"`
	HostName                       *string                           `json:"HostName,omitempty" xml:"HostName,omitempty"`
	HttpRedirectCode               *int64                            `json:"HttpRedirectCode,omitempty" xml:"HttpRedirectCode,omitempty"`
	MirrorCheckMd5                 *bool                             `json:"MirrorCheckMd5,omitempty" xml:"MirrorCheckMd5,omitempty"`
	MirrorFollowRedirect           *bool                             `json:"MirrorFollowRedirect,omitempty" xml:"MirrorFollowRedirect,omitempty"`
	MirrorHeaders                  *RoutingRuleRedirectMirrorHeaders `json:"MirrorHeaders,omitempty" xml:"MirrorHeaders,omitempty" type:"Struct"`
	MirrorPassQueryString          *bool                             `json:"MirrorPassQueryString,omitempty" xml:"MirrorPassQueryString,omitempty"`
	MirrorURL                      *string                           `json:"MirrorURL,omitempty" xml:"MirrorURL,omitempty"`
	PassQueryString                *bool                             `json:"PassQueryString,omitempty" xml:"PassQueryString,omitempty"`
	Protocol                       *string                           `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	RedirectType                   *string                           `json:"RedirectType,omitempty" xml:"RedirectType,omitempty"`
	ReplaceKeyPrefixWith           *string                           `json:"ReplaceKeyPrefixWith,omitempty" xml:"ReplaceKeyPrefixWith,omitempty"`
	ReplaceKeyWith                 *string                           `json:"ReplaceKeyWith,omitempty" xml:"ReplaceKeyWith,omitempty"`
	TransparentMirrorResponseCodes *string                           `json:"TransparentMirrorResponseCodes,omitempty" xml:"TransparentMirrorResponseCodes,omitempty"`
}

func (s RoutingRuleRedirect) String() string {
	return tea.Prettify(s)
}

func (s RoutingRuleRedirect) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirect) SetEnableReplacePrefix(v bool) *RoutingRuleRedirect {
	s.EnableReplacePrefix = &v
	return s
}

func (s *RoutingRuleRedirect) SetHostName(v string) *RoutingRuleRedirect {
	s.HostName = &v
	return s
}

func (s *RoutingRuleRedirect) SetHttpRedirectCode(v int64) *RoutingRuleRedirect {
	s.HttpRedirectCode = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorCheckMd5(v bool) *RoutingRuleRedirect {
	s.MirrorCheckMd5 = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorFollowRedirect(v bool) *RoutingRuleRedirect {
	s.MirrorFollowRedirect = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorHeaders(v *RoutingRuleRedirectMirrorHeaders) *RoutingRuleRedirect {
	s.MirrorHeaders = v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorPassQueryString(v bool) *RoutingRuleRedirect {
	s.MirrorPassQueryString = &v
	return s
}

func (s *RoutingRuleRedirect) SetMirrorURL(v string) *RoutingRuleRedirect {
	s.MirrorURL = &v
	return s
}

func (s *RoutingRuleRedirect) SetPassQueryString(v bool) *RoutingRuleRedirect {
	s.PassQueryString = &v
	return s
}

func (s *RoutingRuleRedirect) SetProtocol(v string) *RoutingRuleRedirect {
	s.Protocol = &v
	return s
}

func (s *RoutingRuleRedirect) SetRedirectType(v string) *RoutingRuleRedirect {
	s.RedirectType = &v
	return s
}

func (s *RoutingRuleRedirect) SetReplaceKeyPrefixWith(v string) *RoutingRuleRedirect {
	s.ReplaceKeyPrefixWith = &v
	return s
}

func (s *RoutingRuleRedirect) SetReplaceKeyWith(v string) *RoutingRuleRedirect {
	s.ReplaceKeyWith = &v
	return s
}

func (s *RoutingRuleRedirect) SetTransparentMirrorResponseCodes(v string) *RoutingRuleRedirect {
	s.TransparentMirrorResponseCodes = &v
	return s
}

type RoutingRuleRedirectMirrorHeaders struct {
	Pass    []*string                              `json:"Pass,omitempty" xml:"Pass,omitempty" type:"Repeated"`
	PassAll *bool                                  `json:"PassAll,omitempty" xml:"PassAll,omitempty"`
	Remove  []*string                              `json:"Remove,omitempty" xml:"Remove,omitempty" type:"Repeated"`
	Set     []*RoutingRuleRedirectMirrorHeadersSet `json:"Set,omitempty" xml:"Set,omitempty" type:"Repeated"`
}

func (s RoutingRuleRedirectMirrorHeaders) String() string {
	return tea.Prettify(s)
}

func (s RoutingRuleRedirectMirrorHeaders) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorHeaders) SetPass(v []*string) *RoutingRuleRedirectMirrorHeaders {
	s.Pass = v
	return s
}

func (s *RoutingRuleRedirectMirrorHeaders) SetPassAll(v bool) *RoutingRuleRedirectMirrorHeaders {
	s.PassAll = &v
	return s
}

func (s *RoutingRuleRedirectMirrorHeaders) SetRemove(v []*string) *RoutingRuleRedirectMirrorHeaders {
	s.Remove = v
	return s
}

func (s *RoutingRuleRedirectMirrorHeaders) SetSet(v []*RoutingRuleRedirectMirrorHeadersSet) *RoutingRuleRedirectMirrorHeaders {
	s.Set = v
	return s
}

type RoutingRuleRedirectMirrorHeadersSet struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RoutingRuleRedirectMirrorHeadersSet) String() string {
	return tea.Prettify(s)
}

func (s RoutingRuleRedirectMirrorHeadersSet) GoString() string {
	return s.String()
}

func (s *RoutingRuleRedirectMirrorHeadersSet) SetKey(v string) *RoutingRuleRedirectMirrorHeadersSet {
	s.Key = &v
	return s
}

func (s *RoutingRuleRedirectMirrorHeadersSet) SetValue(v string) *RoutingRuleRedirectMirrorHeadersSet {
	s.Value = &v
	return s
}

type SSEKMS struct {
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
}

func (s SSEKMS) String() string {
	return tea.Prettify(s)
}

func (s SSEKMS) GoString() string {
	return s.String()
}

func (s *SSEKMS) SetKeyId(v string) *SSEKMS {
	s.KeyId = &v
	return s
}

type SSEOSS struct {
}

func (s SSEOSS) String() string {
	return tea.Prettify(s)
}

func (s SSEOSS) GoString() string {
	return s.String()
}

type SelectMetaRequest struct {
	InputSerialization *InputSerialization `json:"InputSerialization,omitempty" xml:"InputSerialization,omitempty"`
	OverwriteIfExists  *bool               `json:"OverwriteIfExists,omitempty" xml:"OverwriteIfExists,omitempty"`
}

func (s SelectMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s SelectMetaRequest) GoString() string {
	return s.String()
}

func (s *SelectMetaRequest) SetInputSerialization(v *InputSerialization) *SelectMetaRequest {
	s.InputSerialization = v
	return s
}

func (s *SelectMetaRequest) SetOverwriteIfExists(v bool) *SelectMetaRequest {
	s.OverwriteIfExists = &v
	return s
}

type SelectMetaStatus struct {
	ColsCount         *int64  `json:"ColsCount,omitempty" xml:"ColsCount,omitempty"`
	ErrorMessage      *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Offset            *int64  `json:"Offset,omitempty" xml:"Offset,omitempty"`
	RowsCount         *int64  `json:"RowsCount,omitempty" xml:"RowsCount,omitempty"`
	SplitsCount       *int64  `json:"SplitsCount,omitempty" xml:"SplitsCount,omitempty"`
	Status            *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TotalScannedBytes *int64  `json:"TotalScannedBytes,omitempty" xml:"TotalScannedBytes,omitempty"`
}

func (s SelectMetaStatus) String() string {
	return tea.Prettify(s)
}

func (s SelectMetaStatus) GoString() string {
	return s.String()
}

func (s *SelectMetaStatus) SetColsCount(v int64) *SelectMetaStatus {
	s.ColsCount = &v
	return s
}

func (s *SelectMetaStatus) SetErrorMessage(v string) *SelectMetaStatus {
	s.ErrorMessage = &v
	return s
}

func (s *SelectMetaStatus) SetOffset(v int64) *SelectMetaStatus {
	s.Offset = &v
	return s
}

func (s *SelectMetaStatus) SetRowsCount(v int64) *SelectMetaStatus {
	s.RowsCount = &v
	return s
}

func (s *SelectMetaStatus) SetSplitsCount(v int64) *SelectMetaStatus {
	s.SplitsCount = &v
	return s
}

func (s *SelectMetaStatus) SetStatus(v int64) *SelectMetaStatus {
	s.Status = &v
	return s
}

func (s *SelectMetaStatus) SetTotalScannedBytes(v int64) *SelectMetaStatus {
	s.TotalScannedBytes = &v
	return s
}

type SelectRequest struct {
	Expression          *string               `json:"Expression,omitempty" xml:"Expression,omitempty"`
	InputSerialization  *InputSerialization   `json:"InputSerialization,omitempty" xml:"InputSerialization,omitempty"`
	Options             *SelectRequestOptions `json:"Options,omitempty" xml:"Options,omitempty"`
	OutputSerialization *OutputSerialization  `json:"OutputSerialization,omitempty" xml:"OutputSerialization,omitempty"`
}

func (s SelectRequest) String() string {
	return tea.Prettify(s)
}

func (s SelectRequest) GoString() string {
	return s.String()
}

func (s *SelectRequest) SetExpression(v string) *SelectRequest {
	s.Expression = &v
	return s
}

func (s *SelectRequest) SetInputSerialization(v *InputSerialization) *SelectRequest {
	s.InputSerialization = v
	return s
}

func (s *SelectRequest) SetOptions(v *SelectRequestOptions) *SelectRequest {
	s.Options = v
	return s
}

func (s *SelectRequest) SetOutputSerialization(v *OutputSerialization) *SelectRequest {
	s.OutputSerialization = v
	return s
}

type SelectRequestOptions struct {
	MaxSkippedRecordsAllowed *int64 `json:"MaxSkippedRecordsAllowed,omitempty" xml:"MaxSkippedRecordsAllowed,omitempty"`
	SkipPartialDataRecord    *bool  `json:"SkipPartialDataRecord,omitempty" xml:"SkipPartialDataRecord,omitempty"`
}

func (s SelectRequestOptions) String() string {
	return tea.Prettify(s)
}

func (s SelectRequestOptions) GoString() string {
	return s.String()
}

func (s *SelectRequestOptions) SetMaxSkippedRecordsAllowed(v int64) *SelectRequestOptions {
	s.MaxSkippedRecordsAllowed = &v
	return s
}

func (s *SelectRequestOptions) SetSkipPartialDataRecord(v bool) *SelectRequestOptions {
	s.SkipPartialDataRecord = &v
	return s
}

type ServerSideEncryptionRule struct {
	ApplyServerSideEncryptionByDefault *ApplyServerSideEncryptionByDefault `json:"ApplyServerSideEncryptionByDefault,omitempty" xml:"ApplyServerSideEncryptionByDefault,omitempty"`
}

func (s ServerSideEncryptionRule) String() string {
	return tea.Prettify(s)
}

func (s ServerSideEncryptionRule) GoString() string {
	return s.String()
}

func (s *ServerSideEncryptionRule) SetApplyServerSideEncryptionByDefault(v *ApplyServerSideEncryptionByDefault) *ServerSideEncryptionRule {
	s.ApplyServerSideEncryptionByDefault = v
	return s
}

type Tag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Tag) String() string {
	return tea.Prettify(s)
}

func (s Tag) GoString() string {
	return s.String()
}

func (s *Tag) SetKey(v string) *Tag {
	s.Key = &v
	return s
}

func (s *Tag) SetValue(v string) *Tag {
	s.Value = &v
	return s
}

type TagSet struct {
	Tags []*Tag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagSet) String() string {
	return tea.Prettify(s)
}

func (s TagSet) GoString() string {
	return s.String()
}

func (s *TagSet) SetTags(v []*Tag) *TagSet {
	s.Tags = v
	return s
}

type Tagging struct {
	TagSet *TagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty"`
}

func (s Tagging) String() string {
	return tea.Prettify(s)
}

func (s Tagging) GoString() string {
	return s.String()
}

func (s *Tagging) SetTagSet(v *TagSet) *Tagging {
	s.TagSet = v
	return s
}

type TransferAccelerationConfiguration struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s TransferAccelerationConfiguration) String() string {
	return tea.Prettify(s)
}

func (s TransferAccelerationConfiguration) GoString() string {
	return s.String()
}

func (s *TransferAccelerationConfiguration) SetEnabled(v bool) *TransferAccelerationConfiguration {
	s.Enabled = &v
	return s
}

type Upload struct {
	Initiated *string `json:"Initiated,omitempty" xml:"Initiated,omitempty"`
	Key       *string `json:"Key,omitempty" xml:"Key,omitempty"`
	UploadId  *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s Upload) String() string {
	return tea.Prettify(s)
}

func (s Upload) GoString() string {
	return s.String()
}

func (s *Upload) SetInitiated(v string) *Upload {
	s.Initiated = &v
	return s
}

func (s *Upload) SetKey(v string) *Upload {
	s.Key = &v
	return s
}

func (s *Upload) SetUploadId(v string) *Upload {
	s.UploadId = &v
	return s
}

type VersioningConfiguration struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s VersioningConfiguration) String() string {
	return tea.Prettify(s)
}

func (s VersioningConfiguration) GoString() string {
	return s.String()
}

func (s *VersioningConfiguration) SetStatus(v string) *VersioningConfiguration {
	s.Status = &v
	return s
}

type WebsiteConfiguration struct {
	ErrorDocument *ErrorDocument                    `json:"ErrorDocument,omitempty" xml:"ErrorDocument,omitempty"`
	IndexDocument *IndexDocument                    `json:"IndexDocument,omitempty" xml:"IndexDocument,omitempty"`
	RoutingRules  *WebsiteConfigurationRoutingRules `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty" type:"Struct"`
}

func (s WebsiteConfiguration) String() string {
	return tea.Prettify(s)
}

func (s WebsiteConfiguration) GoString() string {
	return s.String()
}

func (s *WebsiteConfiguration) SetErrorDocument(v *ErrorDocument) *WebsiteConfiguration {
	s.ErrorDocument = v
	return s
}

func (s *WebsiteConfiguration) SetIndexDocument(v *IndexDocument) *WebsiteConfiguration {
	s.IndexDocument = v
	return s
}

func (s *WebsiteConfiguration) SetRoutingRules(v *WebsiteConfigurationRoutingRules) *WebsiteConfiguration {
	s.RoutingRules = v
	return s
}

type WebsiteConfigurationRoutingRules struct {
	RoutingRule []*RoutingRule `json:"RoutingRule,omitempty" xml:"RoutingRule,omitempty" type:"Repeated"`
}

func (s WebsiteConfigurationRoutingRules) String() string {
	return tea.Prettify(s)
}

func (s WebsiteConfigurationRoutingRules) GoString() string {
	return s.String()
}

func (s *WebsiteConfigurationRoutingRules) SetRoutingRule(v []*RoutingRule) *WebsiteConfigurationRoutingRules {
	s.RoutingRule = v
	return s
}

type AbortBucketWormResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s AbortBucketWormResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortBucketWormResponse) GoString() string {
	return s.String()
}

func (s *AbortBucketWormResponse) SetHeaders(v map[string]*string) *AbortBucketWormResponse {
	s.Headers = v
	return s
}

func (s *AbortBucketWormResponse) SetStatusCode(v int32) *AbortBucketWormResponse {
	s.StatusCode = &v
	return s
}

type AbortMultipartUploadRequest struct {
	UploadId *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s AbortMultipartUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s AbortMultipartUploadRequest) GoString() string {
	return s.String()
}

func (s *AbortMultipartUploadRequest) SetUploadId(v string) *AbortMultipartUploadRequest {
	s.UploadId = &v
	return s
}

type AbortMultipartUploadResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s AbortMultipartUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s AbortMultipartUploadResponse) GoString() string {
	return s.String()
}

func (s *AbortMultipartUploadResponse) SetHeaders(v map[string]*string) *AbortMultipartUploadResponse {
	s.Headers = v
	return s
}

func (s *AbortMultipartUploadResponse) SetStatusCode(v int32) *AbortMultipartUploadResponse {
	s.StatusCode = &v
	return s
}

type AppendObjectHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	CacheControl         *string            `json:"Cache-Control,omitempty" xml:"Cache-Control,omitempty"`
	ContentDisposition   *string            `json:"Content-Disposition,omitempty" xml:"Content-Disposition,omitempty"`
	ContentEncoding      *string            `json:"Content-Encoding,omitempty" xml:"Content-Encoding,omitempty"`
	ContentMD5           *string            `json:"Content-MD5,omitempty" xml:"Content-MD5,omitempty"`
	Expires              *string            `json:"Expires,omitempty" xml:"Expires,omitempty"`
	MetaData             map[string]*string `json:"x-oss-meta-*,omitempty" xml:"x-oss-meta-*,omitempty"`
	Acl                  *string            `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
	ServerSideEncryption *string            `json:"x-oss-server-side-encryption,omitempty" xml:"x-oss-server-side-encryption,omitempty"`
	StorageClass         *string            `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
}

func (s AppendObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s AppendObjectHeaders) GoString() string {
	return s.String()
}

func (s *AppendObjectHeaders) SetCommonHeaders(v map[string]*string) *AppendObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *AppendObjectHeaders) SetCacheControl(v string) *AppendObjectHeaders {
	s.CacheControl = &v
	return s
}

func (s *AppendObjectHeaders) SetContentDisposition(v string) *AppendObjectHeaders {
	s.ContentDisposition = &v
	return s
}

func (s *AppendObjectHeaders) SetContentEncoding(v string) *AppendObjectHeaders {
	s.ContentEncoding = &v
	return s
}

func (s *AppendObjectHeaders) SetContentMD5(v string) *AppendObjectHeaders {
	s.ContentMD5 = &v
	return s
}

func (s *AppendObjectHeaders) SetExpires(v string) *AppendObjectHeaders {
	s.Expires = &v
	return s
}

func (s *AppendObjectHeaders) SetMetaData(v map[string]*string) *AppendObjectHeaders {
	s.MetaData = v
	return s
}

func (s *AppendObjectHeaders) SetAcl(v string) *AppendObjectHeaders {
	s.Acl = &v
	return s
}

func (s *AppendObjectHeaders) SetServerSideEncryption(v string) *AppendObjectHeaders {
	s.ServerSideEncryption = &v
	return s
}

func (s *AppendObjectHeaders) SetStorageClass(v string) *AppendObjectHeaders {
	s.StorageClass = &v
	return s
}

type AppendObjectRequest struct {
	Body     io.Reader `json:"body,omitempty" xml:"body,omitempty"`
	Position *int64    `json:"position,omitempty" xml:"position,omitempty"`
}

func (s AppendObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s AppendObjectRequest) GoString() string {
	return s.String()
}

func (s *AppendObjectRequest) SetBody(v io.Reader) *AppendObjectRequest {
	s.Body = v
	return s
}

func (s *AppendObjectRequest) SetPosition(v int64) *AppendObjectRequest {
	s.Position = &v
	return s
}

type AppendObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s AppendObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s AppendObjectResponse) GoString() string {
	return s.String()
}

func (s *AppendObjectResponse) SetHeaders(v map[string]*string) *AppendObjectResponse {
	s.Headers = v
	return s
}

func (s *AppendObjectResponse) SetStatusCode(v int32) *AppendObjectResponse {
	s.StatusCode = &v
	return s
}

type CompleteBucketWormRequest struct {
	WormId *string `json:"wormId,omitempty" xml:"wormId,omitempty"`
}

func (s CompleteBucketWormRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteBucketWormRequest) GoString() string {
	return s.String()
}

func (s *CompleteBucketWormRequest) SetWormId(v string) *CompleteBucketWormRequest {
	s.WormId = &v
	return s
}

type CompleteBucketWormResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s CompleteBucketWormResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteBucketWormResponse) GoString() string {
	return s.String()
}

func (s *CompleteBucketWormResponse) SetHeaders(v map[string]*string) *CompleteBucketWormResponse {
	s.Headers = v
	return s
}

func (s *CompleteBucketWormResponse) SetStatusCode(v int32) *CompleteBucketWormResponse {
	s.StatusCode = &v
	return s
}

type CompleteMultipartUploadHeaders struct {
	CommonHeaders   map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	CompleteAll     *string            `json:"x-oss-complete-all,omitempty" xml:"x-oss-complete-all,omitempty"`
	ForbidOverwrite *string            `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
}

func (s CompleteMultipartUploadHeaders) String() string {
	return tea.Prettify(s)
}

func (s CompleteMultipartUploadHeaders) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadHeaders) SetCommonHeaders(v map[string]*string) *CompleteMultipartUploadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CompleteMultipartUploadHeaders) SetCompleteAll(v string) *CompleteMultipartUploadHeaders {
	s.CompleteAll = &v
	return s
}

func (s *CompleteMultipartUploadHeaders) SetForbidOverwrite(v string) *CompleteMultipartUploadHeaders {
	s.ForbidOverwrite = &v
	return s
}

type CompleteMultipartUploadRequest struct {
	CompleteMultipartUpload *CompleteMultipartUpload `json:"completeMultipartUpload,omitempty" xml:"completeMultipartUpload,omitempty"`
	EncodingType            *string                  `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	UploadId                *string                  `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s CompleteMultipartUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteMultipartUploadRequest) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadRequest) SetCompleteMultipartUpload(v *CompleteMultipartUpload) *CompleteMultipartUploadRequest {
	s.CompleteMultipartUpload = v
	return s
}

func (s *CompleteMultipartUploadRequest) SetEncodingType(v string) *CompleteMultipartUploadRequest {
	s.EncodingType = &v
	return s
}

func (s *CompleteMultipartUploadRequest) SetUploadId(v string) *CompleteMultipartUploadRequest {
	s.UploadId = &v
	return s
}

type CompleteMultipartUploadResponseBody struct {
	Bucket       *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Location     *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s CompleteMultipartUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CompleteMultipartUploadResponseBody) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadResponseBody) SetBucket(v string) *CompleteMultipartUploadResponseBody {
	s.Bucket = &v
	return s
}

func (s *CompleteMultipartUploadResponseBody) SetETag(v string) *CompleteMultipartUploadResponseBody {
	s.ETag = &v
	return s
}

func (s *CompleteMultipartUploadResponseBody) SetEncodingType(v string) *CompleteMultipartUploadResponseBody {
	s.EncodingType = &v
	return s
}

func (s *CompleteMultipartUploadResponseBody) SetKey(v string) *CompleteMultipartUploadResponseBody {
	s.Key = &v
	return s
}

func (s *CompleteMultipartUploadResponseBody) SetLocation(v string) *CompleteMultipartUploadResponseBody {
	s.Location = &v
	return s
}

type CompleteMultipartUploadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CompleteMultipartUploadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CompleteMultipartUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteMultipartUploadResponse) GoString() string {
	return s.String()
}

func (s *CompleteMultipartUploadResponse) SetHeaders(v map[string]*string) *CompleteMultipartUploadResponse {
	s.Headers = v
	return s
}

func (s *CompleteMultipartUploadResponse) SetStatusCode(v int32) *CompleteMultipartUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteMultipartUploadResponse) SetBody(v *CompleteMultipartUploadResponseBody) *CompleteMultipartUploadResponse {
	s.Body = v
	return s
}

type CopyObjectHeaders struct {
	CommonHeaders               map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	CopySource                  *string            `json:"x-oss-copy-source,omitempty" xml:"x-oss-copy-source,omitempty"`
	CopySourceIfMatch           *string            `json:"x-oss-copy-source-if-match,omitempty" xml:"x-oss-copy-source-if-match,omitempty"`
	CopySourceIfModifiedSince   *string            `json:"x-oss-copy-source-if-modified-since,omitempty" xml:"x-oss-copy-source-if-modified-since,omitempty"`
	CopySourceIfNoneMatch       *string            `json:"x-oss-copy-source-if-none-match,omitempty" xml:"x-oss-copy-source-if-none-match,omitempty"`
	CopySourceIfUnmodifiedSince *string            `json:"x-oss-copy-source-if-unmodified-since,omitempty" xml:"x-oss-copy-source-if-unmodified-since,omitempty"`
	ForbidOverwrite             *string            `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
	MetaData                    map[string]*string `json:"x-oss-meta-*,omitempty" xml:"x-oss-meta-*,omitempty"`
	MetadataDirective           *string            `json:"x-oss-metadata-directive,omitempty" xml:"x-oss-metadata-directive,omitempty"`
	Acl                         *string            `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
	ServerSideEncryption        *string            `json:"x-oss-server-side-encryption,omitempty" xml:"x-oss-server-side-encryption,omitempty"`
	SseKeyId                    *string            `json:"x-oss-server-side-encryption-key-id,omitempty" xml:"x-oss-server-side-encryption-key-id,omitempty"`
	StorageClass                *string            `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
	Tagging                     *string            `json:"x-oss-tagging,omitempty" xml:"x-oss-tagging,omitempty"`
	TaggingDirective            *string            `json:"x-oss-tagging-directive,omitempty" xml:"x-oss-tagging-directive,omitempty"`
}

func (s CopyObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s CopyObjectHeaders) GoString() string {
	return s.String()
}

func (s *CopyObjectHeaders) SetCommonHeaders(v map[string]*string) *CopyObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *CopyObjectHeaders) SetCopySource(v string) *CopyObjectHeaders {
	s.CopySource = &v
	return s
}

func (s *CopyObjectHeaders) SetCopySourceIfMatch(v string) *CopyObjectHeaders {
	s.CopySourceIfMatch = &v
	return s
}

func (s *CopyObjectHeaders) SetCopySourceIfModifiedSince(v string) *CopyObjectHeaders {
	s.CopySourceIfModifiedSince = &v
	return s
}

func (s *CopyObjectHeaders) SetCopySourceIfNoneMatch(v string) *CopyObjectHeaders {
	s.CopySourceIfNoneMatch = &v
	return s
}

func (s *CopyObjectHeaders) SetCopySourceIfUnmodifiedSince(v string) *CopyObjectHeaders {
	s.CopySourceIfUnmodifiedSince = &v
	return s
}

func (s *CopyObjectHeaders) SetForbidOverwrite(v string) *CopyObjectHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *CopyObjectHeaders) SetMetaData(v map[string]*string) *CopyObjectHeaders {
	s.MetaData = v
	return s
}

func (s *CopyObjectHeaders) SetMetadataDirective(v string) *CopyObjectHeaders {
	s.MetadataDirective = &v
	return s
}

func (s *CopyObjectHeaders) SetAcl(v string) *CopyObjectHeaders {
	s.Acl = &v
	return s
}

func (s *CopyObjectHeaders) SetServerSideEncryption(v string) *CopyObjectHeaders {
	s.ServerSideEncryption = &v
	return s
}

func (s *CopyObjectHeaders) SetSseKeyId(v string) *CopyObjectHeaders {
	s.SseKeyId = &v
	return s
}

func (s *CopyObjectHeaders) SetStorageClass(v string) *CopyObjectHeaders {
	s.StorageClass = &v
	return s
}

func (s *CopyObjectHeaders) SetTagging(v string) *CopyObjectHeaders {
	s.Tagging = &v
	return s
}

func (s *CopyObjectHeaders) SetTaggingDirective(v string) *CopyObjectHeaders {
	s.TaggingDirective = &v
	return s
}

type CopyObjectResponseBody struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s CopyObjectResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyObjectResponseBody) GoString() string {
	return s.String()
}

func (s *CopyObjectResponseBody) SetETag(v string) *CopyObjectResponseBody {
	s.ETag = &v
	return s
}

func (s *CopyObjectResponseBody) SetLastModified(v string) *CopyObjectResponseBody {
	s.LastModified = &v
	return s
}

type CopyObjectResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CopyObjectResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CopyObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyObjectResponse) GoString() string {
	return s.String()
}

func (s *CopyObjectResponse) SetHeaders(v map[string]*string) *CopyObjectResponse {
	s.Headers = v
	return s
}

func (s *CopyObjectResponse) SetStatusCode(v int32) *CopyObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyObjectResponse) SetBody(v *CopyObjectResponseBody) *CopyObjectResponse {
	s.Body = v
	return s
}

type CreateSelectObjectMetaRequest struct {
	SelectMetaRequest *SelectMetaRequest `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSelectObjectMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSelectObjectMetaRequest) GoString() string {
	return s.String()
}

func (s *CreateSelectObjectMetaRequest) SetSelectMetaRequest(v *SelectMetaRequest) *CreateSelectObjectMetaRequest {
	s.SelectMetaRequest = v
	return s
}

type CreateSelectObjectMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SelectMetaStatus  `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateSelectObjectMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSelectObjectMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateSelectObjectMetaResponse) SetHeaders(v map[string]*string) *CreateSelectObjectMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateSelectObjectMetaResponse) SetStatusCode(v int32) *CreateSelectObjectMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSelectObjectMetaResponse) SetBody(v *SelectMetaStatus) *CreateSelectObjectMetaResponse {
	s.Body = v
	return s
}

type DeleteBucketResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketResponse) SetHeaders(v map[string]*string) *DeleteBucketResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketResponse) SetStatusCode(v int32) *DeleteBucketResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketCorsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketCorsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketCorsResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketCorsResponse) SetHeaders(v map[string]*string) *DeleteBucketCorsResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketCorsResponse) SetStatusCode(v int32) *DeleteBucketCorsResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketEncryptionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketEncryptionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketEncryptionResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketEncryptionResponse) SetHeaders(v map[string]*string) *DeleteBucketEncryptionResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketEncryptionResponse) SetStatusCode(v int32) *DeleteBucketEncryptionResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketInventoryRequest struct {
	InventoryId *string `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
}

func (s DeleteBucketInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketInventoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketInventoryRequest) SetInventoryId(v string) *DeleteBucketInventoryRequest {
	s.InventoryId = &v
	return s
}

type DeleteBucketInventoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketInventoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketInventoryResponse) SetHeaders(v map[string]*string) *DeleteBucketInventoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketInventoryResponse) SetStatusCode(v int32) *DeleteBucketInventoryResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketLifecycleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketLifecycleResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketLifecycleResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketLifecycleResponse) SetHeaders(v map[string]*string) *DeleteBucketLifecycleResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketLifecycleResponse) SetStatusCode(v int32) *DeleteBucketLifecycleResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketLoggingResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketLoggingResponse) SetHeaders(v map[string]*string) *DeleteBucketLoggingResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketLoggingResponse) SetStatusCode(v int32) *DeleteBucketLoggingResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketPolicyResponse) SetHeaders(v map[string]*string) *DeleteBucketPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketPolicyResponse) SetStatusCode(v int32) *DeleteBucketPolicyResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketReplicationRequest struct {
	Body *ReplicationRules `json:"ReplicationRules,omitempty" xml:"ReplicationRules,omitempty"`
}

func (s DeleteBucketReplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketReplicationRequest) GoString() string {
	return s.String()
}

func (s *DeleteBucketReplicationRequest) SetBody(v *ReplicationRules) *DeleteBucketReplicationRequest {
	s.Body = v
	return s
}

type DeleteBucketReplicationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketReplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketReplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketReplicationResponse) SetHeaders(v map[string]*string) *DeleteBucketReplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketReplicationResponse) SetStatusCode(v int32) *DeleteBucketReplicationResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketTagsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketTagsResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketTagsResponse) SetHeaders(v map[string]*string) *DeleteBucketTagsResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketTagsResponse) SetStatusCode(v int32) *DeleteBucketTagsResponse {
	s.StatusCode = &v
	return s
}

type DeleteBucketWebsiteResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteBucketWebsiteResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteBucketWebsiteResponse) GoString() string {
	return s.String()
}

func (s *DeleteBucketWebsiteResponse) SetHeaders(v map[string]*string) *DeleteBucketWebsiteResponse {
	s.Headers = v
	return s
}

func (s *DeleteBucketWebsiteResponse) SetStatusCode(v int32) *DeleteBucketWebsiteResponse {
	s.StatusCode = &v
	return s
}

type DeleteLiveChannelResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteLiveChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveChannelResponse) SetHeaders(v map[string]*string) *DeleteLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveChannelResponse) SetStatusCode(v int32) *DeleteLiveChannelResponse {
	s.StatusCode = &v
	return s
}

type DeleteMultipleObjectsRequest struct {
	Delete       *Delete `json:"Delete,omitempty" xml:"Delete,omitempty"`
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
}

func (s DeleteMultipleObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultipleObjectsRequest) GoString() string {
	return s.String()
}

func (s *DeleteMultipleObjectsRequest) SetDelete(v *Delete) *DeleteMultipleObjectsRequest {
	s.Delete = v
	return s
}

func (s *DeleteMultipleObjectsRequest) SetEncodingType(v string) *DeleteMultipleObjectsRequest {
	s.EncodingType = &v
	return s
}

type DeleteMultipleObjectsResponseBody struct {
	Deleted      []*DeletedObject `json:"Deleted,omitempty" xml:"Deleted,omitempty" type:"Repeated"`
	EncodingType *string          `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
}

func (s DeleteMultipleObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultipleObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMultipleObjectsResponseBody) SetDeleted(v []*DeletedObject) *DeleteMultipleObjectsResponseBody {
	s.Deleted = v
	return s
}

func (s *DeleteMultipleObjectsResponseBody) SetEncodingType(v string) *DeleteMultipleObjectsResponseBody {
	s.EncodingType = &v
	return s
}

type DeleteMultipleObjectsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteMultipleObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteMultipleObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteMultipleObjectsResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultipleObjectsResponse) SetHeaders(v map[string]*string) *DeleteMultipleObjectsResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultipleObjectsResponse) SetStatusCode(v int32) *DeleteMultipleObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultipleObjectsResponse) SetBody(v *DeleteMultipleObjectsResponseBody) *DeleteMultipleObjectsResponse {
	s.Body = v
	return s
}

type DeleteObjectRequest struct {
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s DeleteObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectRequest) GoString() string {
	return s.String()
}

func (s *DeleteObjectRequest) SetVersionId(v string) *DeleteObjectRequest {
	s.VersionId = &v
	return s
}

type DeleteObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectResponse) GoString() string {
	return s.String()
}

func (s *DeleteObjectResponse) SetHeaders(v map[string]*string) *DeleteObjectResponse {
	s.Headers = v
	return s
}

func (s *DeleteObjectResponse) SetStatusCode(v int32) *DeleteObjectResponse {
	s.StatusCode = &v
	return s
}

type DeleteObjectTaggingRequest struct {
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s DeleteObjectTaggingRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectTaggingRequest) GoString() string {
	return s.String()
}

func (s *DeleteObjectTaggingRequest) SetVersionId(v string) *DeleteObjectTaggingRequest {
	s.VersionId = &v
	return s
}

type DeleteObjectTaggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s DeleteObjectTaggingResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteObjectTaggingResponse) GoString() string {
	return s.String()
}

func (s *DeleteObjectTaggingResponse) SetHeaders(v map[string]*string) *DeleteObjectTaggingResponse {
	s.Headers = v
	return s
}

func (s *DeleteObjectTaggingResponse) SetStatusCode(v int32) *DeleteObjectTaggingResponse {
	s.StatusCode = &v
	return s
}

type DescribeRegionsRequest struct {
	Regions *string `json:"regions,omitempty" xml:"regions,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetRegions(v string) *DescribeRegionsRequest {
	s.Regions = &v
	return s
}

type DescribeRegionsResponseBody struct {
	RegionInfos []*RegionInfo `json:"RegionInfo,omitempty" xml:"RegionInfo,omitempty" type:"Repeated"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegionInfos(v []*RegionInfo) *DescribeRegionsResponseBody {
	s.RegionInfos = v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type ExtendBucketWormRequest struct {
	ExtendWormConfiguration *ExtendWormConfiguration `json:"extendWormConfiguration,omitempty" xml:"extendWormConfiguration,omitempty"`
	WormId                  *string                  `json:"wormId,omitempty" xml:"wormId,omitempty"`
}

func (s ExtendBucketWormRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtendBucketWormRequest) GoString() string {
	return s.String()
}

func (s *ExtendBucketWormRequest) SetExtendWormConfiguration(v *ExtendWormConfiguration) *ExtendBucketWormRequest {
	s.ExtendWormConfiguration = v
	return s
}

func (s *ExtendBucketWormRequest) SetWormId(v string) *ExtendBucketWormRequest {
	s.WormId = &v
	return s
}

type ExtendBucketWormResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s ExtendBucketWormResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtendBucketWormResponse) GoString() string {
	return s.String()
}

func (s *ExtendBucketWormResponse) SetHeaders(v map[string]*string) *ExtendBucketWormResponse {
	s.Headers = v
	return s
}

func (s *ExtendBucketWormResponse) SetStatusCode(v int32) *ExtendBucketWormResponse {
	s.StatusCode = &v
	return s
}

type GetBucketAclResponseBody struct {
	AccessControlList *GetBucketAclResponseBodyAccessControlList `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty" type:"Struct"`
	Owner             *Owner                                     `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s GetBucketAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketAclResponseBody) SetAccessControlList(v *GetBucketAclResponseBodyAccessControlList) *GetBucketAclResponseBody {
	s.AccessControlList = v
	return s
}

func (s *GetBucketAclResponseBody) SetOwner(v *Owner) *GetBucketAclResponseBody {
	s.Owner = v
	return s
}

type GetBucketAclResponseBodyAccessControlList struct {
	Grant *string `json:"Grant,omitempty" xml:"Grant,omitempty"`
}

func (s GetBucketAclResponseBodyAccessControlList) String() string {
	return tea.Prettify(s)
}

func (s GetBucketAclResponseBodyAccessControlList) GoString() string {
	return s.String()
}

func (s *GetBucketAclResponseBodyAccessControlList) SetGrant(v string) *GetBucketAclResponseBodyAccessControlList {
	s.Grant = &v
	return s
}

type GetBucketAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketAclResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketAclResponse) GoString() string {
	return s.String()
}

func (s *GetBucketAclResponse) SetHeaders(v map[string]*string) *GetBucketAclResponse {
	s.Headers = v
	return s
}

func (s *GetBucketAclResponse) SetStatusCode(v int32) *GetBucketAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketAclResponse) SetBody(v *GetBucketAclResponseBody) *GetBucketAclResponse {
	s.Body = v
	return s
}

type GetBucketCorsResponseBody struct {
	CORSRule     []*CORSRule `json:"CORSRule,omitempty" xml:"CORSRule,omitempty" type:"Repeated"`
	ResponseVary *bool       `json:"ResponseVary,omitempty" xml:"ResponseVary,omitempty"`
}

func (s GetBucketCorsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketCorsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketCorsResponseBody) SetCORSRule(v []*CORSRule) *GetBucketCorsResponseBody {
	s.CORSRule = v
	return s
}

func (s *GetBucketCorsResponseBody) SetResponseVary(v bool) *GetBucketCorsResponseBody {
	s.ResponseVary = &v
	return s
}

type GetBucketCorsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketCorsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketCorsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketCorsResponse) GoString() string {
	return s.String()
}

func (s *GetBucketCorsResponse) SetHeaders(v map[string]*string) *GetBucketCorsResponse {
	s.Headers = v
	return s
}

func (s *GetBucketCorsResponse) SetStatusCode(v int32) *GetBucketCorsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketCorsResponse) SetBody(v *GetBucketCorsResponseBody) *GetBucketCorsResponse {
	s.Body = v
	return s
}

type GetBucketEncryptionResponseBody struct {
	ApplyServerSideEncryptionByDefault *ApplyServerSideEncryptionByDefault `json:"ApplyServerSideEncryptionByDefault,omitempty" xml:"ApplyServerSideEncryptionByDefault,omitempty"`
}

func (s GetBucketEncryptionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketEncryptionResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketEncryptionResponseBody) SetApplyServerSideEncryptionByDefault(v *ApplyServerSideEncryptionByDefault) *GetBucketEncryptionResponseBody {
	s.ApplyServerSideEncryptionByDefault = v
	return s
}

type GetBucketEncryptionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketEncryptionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketEncryptionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketEncryptionResponse) GoString() string {
	return s.String()
}

func (s *GetBucketEncryptionResponse) SetHeaders(v map[string]*string) *GetBucketEncryptionResponse {
	s.Headers = v
	return s
}

func (s *GetBucketEncryptionResponse) SetStatusCode(v int32) *GetBucketEncryptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketEncryptionResponse) SetBody(v *GetBucketEncryptionResponseBody) *GetBucketEncryptionResponse {
	s.Body = v
	return s
}

type GetBucketInfoResponseBody struct {
	BucketInfo *GetBucketInfoResponseBodyBucketInfo `json:"Bucket,omitempty" xml:"Bucket,omitempty" type:"Struct"`
}

func (s GetBucketInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponseBody) SetBucketInfo(v *GetBucketInfoResponseBodyBucketInfo) *GetBucketInfoResponseBody {
	s.BucketInfo = v
	return s
}

type GetBucketInfoResponseBodyBucketInfo struct {
	AccessControlList      *GetBucketInfoResponseBodyBucketInfoAccessControlList `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty" type:"Struct"`
	Comment                *string                                               `json:"Comment,omitempty" xml:"Comment,omitempty"`
	CreationDate           *string                                               `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	CrossRegionReplication *string                                               `json:"CrossRegionReplication,omitempty" xml:"CrossRegionReplication,omitempty"`
	ExtranetEndpoint       *string                                               `json:"ExtranetEndpoint,omitempty" xml:"ExtranetEndpoint,omitempty"`
	IntranetEndpoint       *string                                               `json:"IntranetEndpoint,omitempty" xml:"IntranetEndpoint,omitempty"`
	Location               *string                                               `json:"Location,omitempty" xml:"Location,omitempty"`
	Name                   *string                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Owner                  *Owner                                                `json:"Owner,omitempty" xml:"Owner,omitempty"`
	StorageClass           *string                                               `json:"StorageClass,omitempty" xml:"StorageClass,omitempty"`
	TransferAcceleration   *string                                               `json:"TransferAcceleration,omitempty" xml:"TransferAcceleration,omitempty"`
}

func (s GetBucketInfoResponseBodyBucketInfo) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInfoResponseBodyBucketInfo) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetAccessControlList(v *GetBucketInfoResponseBodyBucketInfoAccessControlList) *GetBucketInfoResponseBodyBucketInfo {
	s.AccessControlList = v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetComment(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.Comment = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetCreationDate(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.CreationDate = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetCrossRegionReplication(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.CrossRegionReplication = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetExtranetEndpoint(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.ExtranetEndpoint = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetIntranetEndpoint(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.IntranetEndpoint = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetLocation(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.Location = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetName(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.Name = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetOwner(v *Owner) *GetBucketInfoResponseBodyBucketInfo {
	s.Owner = v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetStorageClass(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.StorageClass = &v
	return s
}

func (s *GetBucketInfoResponseBodyBucketInfo) SetTransferAcceleration(v string) *GetBucketInfoResponseBodyBucketInfo {
	s.TransferAcceleration = &v
	return s
}

type GetBucketInfoResponseBodyBucketInfoAccessControlList struct {
	Grant *string `json:"Grant,omitempty" xml:"Grant,omitempty"`
}

func (s GetBucketInfoResponseBodyBucketInfoAccessControlList) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInfoResponseBodyBucketInfoAccessControlList) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponseBodyBucketInfoAccessControlList) SetGrant(v string) *GetBucketInfoResponseBodyBucketInfoAccessControlList {
	s.Grant = &v
	return s
}

type GetBucketInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInfoResponse) GoString() string {
	return s.String()
}

func (s *GetBucketInfoResponse) SetHeaders(v map[string]*string) *GetBucketInfoResponse {
	s.Headers = v
	return s
}

func (s *GetBucketInfoResponse) SetStatusCode(v int32) *GetBucketInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketInfoResponse) SetBody(v *GetBucketInfoResponseBody) *GetBucketInfoResponse {
	s.Body = v
	return s
}

type GetBucketInventoryRequest struct {
	InventoryId *string `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
}

func (s GetBucketInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInventoryRequest) GoString() string {
	return s.String()
}

func (s *GetBucketInventoryRequest) SetInventoryId(v string) *GetBucketInventoryRequest {
	s.InventoryId = &v
	return s
}

type GetBucketInventoryResponseBody struct {
	Destination            *InventoryDestination                         `json:"Destination,omitempty" xml:"Destination,omitempty"`
	Filter                 *InventoryFilter                              `json:"Filter,omitempty" xml:"Filter,omitempty"`
	Id                     *string                                       `json:"Id,omitempty" xml:"Id,omitempty"`
	IncludedObjectVersions *string                                       `json:"IncludedObjectVersions,omitempty" xml:"IncludedObjectVersions,omitempty"`
	IsEnabled              *bool                                         `json:"IsEnabled,omitempty" xml:"IsEnabled,omitempty"`
	OptionalFields         *GetBucketInventoryResponseBodyOptionalFields `json:"OptionalFields,omitempty" xml:"OptionalFields,omitempty" type:"Struct"`
	Schedule               *InventorySchedule                            `json:"Schedule,omitempty" xml:"Schedule,omitempty"`
}

func (s GetBucketInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketInventoryResponseBody) SetDestination(v *InventoryDestination) *GetBucketInventoryResponseBody {
	s.Destination = v
	return s
}

func (s *GetBucketInventoryResponseBody) SetFilter(v *InventoryFilter) *GetBucketInventoryResponseBody {
	s.Filter = v
	return s
}

func (s *GetBucketInventoryResponseBody) SetId(v string) *GetBucketInventoryResponseBody {
	s.Id = &v
	return s
}

func (s *GetBucketInventoryResponseBody) SetIncludedObjectVersions(v string) *GetBucketInventoryResponseBody {
	s.IncludedObjectVersions = &v
	return s
}

func (s *GetBucketInventoryResponseBody) SetIsEnabled(v bool) *GetBucketInventoryResponseBody {
	s.IsEnabled = &v
	return s
}

func (s *GetBucketInventoryResponseBody) SetOptionalFields(v *GetBucketInventoryResponseBodyOptionalFields) *GetBucketInventoryResponseBody {
	s.OptionalFields = v
	return s
}

func (s *GetBucketInventoryResponseBody) SetSchedule(v *InventorySchedule) *GetBucketInventoryResponseBody {
	s.Schedule = v
	return s
}

type GetBucketInventoryResponseBodyOptionalFields struct {
	Field []*string `json:"Field,omitempty" xml:"Field,omitempty" type:"Repeated"`
}

func (s GetBucketInventoryResponseBodyOptionalFields) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInventoryResponseBodyOptionalFields) GoString() string {
	return s.String()
}

func (s *GetBucketInventoryResponseBodyOptionalFields) SetField(v []*string) *GetBucketInventoryResponseBodyOptionalFields {
	s.Field = v
	return s
}

type GetBucketInventoryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketInventoryResponse) GoString() string {
	return s.String()
}

func (s *GetBucketInventoryResponse) SetHeaders(v map[string]*string) *GetBucketInventoryResponse {
	s.Headers = v
	return s
}

func (s *GetBucketInventoryResponse) SetStatusCode(v int32) *GetBucketInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketInventoryResponse) SetBody(v *GetBucketInventoryResponseBody) *GetBucketInventoryResponse {
	s.Body = v
	return s
}

type GetBucketLifecycleResponseBody struct {
	Rules []*LifecycleRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s GetBucketLifecycleResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketLifecycleResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketLifecycleResponseBody) SetRules(v []*LifecycleRule) *GetBucketLifecycleResponseBody {
	s.Rules = v
	return s
}

type GetBucketLifecycleResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketLifecycleResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketLifecycleResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketLifecycleResponse) GoString() string {
	return s.String()
}

func (s *GetBucketLifecycleResponse) SetHeaders(v map[string]*string) *GetBucketLifecycleResponse {
	s.Headers = v
	return s
}

func (s *GetBucketLifecycleResponse) SetStatusCode(v int32) *GetBucketLifecycleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketLifecycleResponse) SetBody(v *GetBucketLifecycleResponseBody) *GetBucketLifecycleResponse {
	s.Body = v
	return s
}

type GetBucketLocationResponseBody struct {
	LocationConstraint *string `json:"LocationConstraint,omitempty" xml:"LocationConstraint,omitempty"`
}

func (s GetBucketLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketLocationResponseBody) SetLocationConstraint(v string) *GetBucketLocationResponseBody {
	s.LocationConstraint = &v
	return s
}

type GetBucketLocationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketLocationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketLocationResponse) SetHeaders(v map[string]*string) *GetBucketLocationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketLocationResponse) SetStatusCode(v int32) *GetBucketLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketLocationResponse) SetBody(v *GetBucketLocationResponseBody) *GetBucketLocationResponse {
	s.Body = v
	return s
}

type GetBucketLoggingResponseBody struct {
	LoggingEnabled *LoggingEnabled `json:"LoggingEnabled,omitempty" xml:"LoggingEnabled,omitempty"`
}

func (s GetBucketLoggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketLoggingResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketLoggingResponseBody) SetLoggingEnabled(v *LoggingEnabled) *GetBucketLoggingResponseBody {
	s.LoggingEnabled = v
	return s
}

type GetBucketLoggingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketLoggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketLoggingResponse) GoString() string {
	return s.String()
}

func (s *GetBucketLoggingResponse) SetHeaders(v map[string]*string) *GetBucketLoggingResponse {
	s.Headers = v
	return s
}

func (s *GetBucketLoggingResponse) SetStatusCode(v int32) *GetBucketLoggingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketLoggingResponse) SetBody(v *GetBucketLoggingResponseBody) *GetBucketLoggingResponse {
	s.Body = v
	return s
}

type GetBucketPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketPolicyResponse) GoString() string {
	return s.String()
}

func (s *GetBucketPolicyResponse) SetHeaders(v map[string]*string) *GetBucketPolicyResponse {
	s.Headers = v
	return s
}

func (s *GetBucketPolicyResponse) SetStatusCode(v int32) *GetBucketPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketPolicyResponse) SetBody(v string) *GetBucketPolicyResponse {
	s.Body = &v
	return s
}

type GetBucketRefererResponseBody struct {
	AllowEmptyReferer *bool                                    `json:"AllowEmptyReferer,omitempty" xml:"AllowEmptyReferer,omitempty"`
	RefererList       *GetBucketRefererResponseBodyRefererList `json:"RefererList,omitempty" xml:"RefererList,omitempty" type:"Struct"`
}

func (s GetBucketRefererResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketRefererResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketRefererResponseBody) SetAllowEmptyReferer(v bool) *GetBucketRefererResponseBody {
	s.AllowEmptyReferer = &v
	return s
}

func (s *GetBucketRefererResponseBody) SetRefererList(v *GetBucketRefererResponseBodyRefererList) *GetBucketRefererResponseBody {
	s.RefererList = v
	return s
}

type GetBucketRefererResponseBodyRefererList struct {
	Referer []*string `json:"Referer,omitempty" xml:"Referer,omitempty" type:"Repeated"`
}

func (s GetBucketRefererResponseBodyRefererList) String() string {
	return tea.Prettify(s)
}

func (s GetBucketRefererResponseBodyRefererList) GoString() string {
	return s.String()
}

func (s *GetBucketRefererResponseBodyRefererList) SetReferer(v []*string) *GetBucketRefererResponseBodyRefererList {
	s.Referer = v
	return s
}

type GetBucketRefererResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketRefererResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketRefererResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketRefererResponse) GoString() string {
	return s.String()
}

func (s *GetBucketRefererResponse) SetHeaders(v map[string]*string) *GetBucketRefererResponse {
	s.Headers = v
	return s
}

func (s *GetBucketRefererResponse) SetStatusCode(v int32) *GetBucketRefererResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketRefererResponse) SetBody(v *GetBucketRefererResponseBody) *GetBucketRefererResponse {
	s.Body = v
	return s
}

type GetBucketReplicationResponseBody struct {
	Rules []*ReplicationRule `json:"Rule,omitempty" xml:"Rule,omitempty" type:"Repeated"`
}

func (s GetBucketReplicationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketReplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationResponseBody) SetRules(v []*ReplicationRule) *GetBucketReplicationResponseBody {
	s.Rules = v
	return s
}

type GetBucketReplicationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketReplicationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketReplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketReplicationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationResponse) SetHeaders(v map[string]*string) *GetBucketReplicationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketReplicationResponse) SetStatusCode(v int32) *GetBucketReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketReplicationResponse) SetBody(v *GetBucketReplicationResponseBody) *GetBucketReplicationResponse {
	s.Body = v
	return s
}

type GetBucketReplicationLocationResponseBody struct {
	Locations                      []*string                                                               `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
	LocationTransferTypeConstraint *GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint `json:"LocationTransferTypeConstraint,omitempty" xml:"LocationTransferTypeConstraint,omitempty" type:"Struct"`
}

func (s GetBucketReplicationLocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketReplicationLocationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationLocationResponseBody) SetLocations(v []*string) *GetBucketReplicationLocationResponseBody {
	s.Locations = v
	return s
}

func (s *GetBucketReplicationLocationResponseBody) SetLocationTransferTypeConstraint(v *GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint) *GetBucketReplicationLocationResponseBody {
	s.LocationTransferTypeConstraint = v
	return s
}

type GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint struct {
	LocationTransferTypes []*LocationTransferType `json:"LocationTransferType,omitempty" xml:"LocationTransferType,omitempty" type:"Repeated"`
}

func (s GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint) String() string {
	return tea.Prettify(s)
}

func (s GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint) SetLocationTransferTypes(v []*LocationTransferType) *GetBucketReplicationLocationResponseBodyLocationTransferTypeConstraint {
	s.LocationTransferTypes = v
	return s
}

type GetBucketReplicationLocationResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketReplicationLocationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketReplicationLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketReplicationLocationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationLocationResponse) SetHeaders(v map[string]*string) *GetBucketReplicationLocationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketReplicationLocationResponse) SetStatusCode(v int32) *GetBucketReplicationLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketReplicationLocationResponse) SetBody(v *GetBucketReplicationLocationResponseBody) *GetBucketReplicationLocationResponse {
	s.Body = v
	return s
}

type GetBucketReplicationProgressRequest struct {
	RuleId *string `json:"rule-id,omitempty" xml:"rule-id,omitempty"`
}

func (s GetBucketReplicationProgressRequest) String() string {
	return tea.Prettify(s)
}

func (s GetBucketReplicationProgressRequest) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationProgressRequest) SetRuleId(v string) *GetBucketReplicationProgressRequest {
	s.RuleId = &v
	return s
}

type GetBucketReplicationProgressResponseBody struct {
	Rule *ReplicationProgressRule `json:"Rule,omitempty" xml:"Rule,omitempty"`
}

func (s GetBucketReplicationProgressResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketReplicationProgressResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationProgressResponseBody) SetRule(v *ReplicationProgressRule) *GetBucketReplicationProgressResponseBody {
	s.Rule = v
	return s
}

type GetBucketReplicationProgressResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketReplicationProgressResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketReplicationProgressResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketReplicationProgressResponse) GoString() string {
	return s.String()
}

func (s *GetBucketReplicationProgressResponse) SetHeaders(v map[string]*string) *GetBucketReplicationProgressResponse {
	s.Headers = v
	return s
}

func (s *GetBucketReplicationProgressResponse) SetStatusCode(v int32) *GetBucketReplicationProgressResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketReplicationProgressResponse) SetBody(v *GetBucketReplicationProgressResponseBody) *GetBucketReplicationProgressResponse {
	s.Body = v
	return s
}

type GetBucketRequestPaymentResponseBody struct {
	Payer *string `json:"Payer,omitempty" xml:"Payer,omitempty"`
}

func (s GetBucketRequestPaymentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketRequestPaymentResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketRequestPaymentResponseBody) SetPayer(v string) *GetBucketRequestPaymentResponseBody {
	s.Payer = &v
	return s
}

type GetBucketRequestPaymentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketRequestPaymentResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketRequestPaymentResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketRequestPaymentResponse) GoString() string {
	return s.String()
}

func (s *GetBucketRequestPaymentResponse) SetHeaders(v map[string]*string) *GetBucketRequestPaymentResponse {
	s.Headers = v
	return s
}

func (s *GetBucketRequestPaymentResponse) SetStatusCode(v int32) *GetBucketRequestPaymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketRequestPaymentResponse) SetBody(v *GetBucketRequestPaymentResponseBody) *GetBucketRequestPaymentResponse {
	s.Body = v
	return s
}

type GetBucketTagsResponseBody struct {
	TagSet *TagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty"`
}

func (s GetBucketTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketTagsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketTagsResponseBody) SetTagSet(v *TagSet) *GetBucketTagsResponseBody {
	s.TagSet = v
	return s
}

type GetBucketTagsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketTagsResponse) GoString() string {
	return s.String()
}

func (s *GetBucketTagsResponse) SetHeaders(v map[string]*string) *GetBucketTagsResponse {
	s.Headers = v
	return s
}

func (s *GetBucketTagsResponse) SetStatusCode(v int32) *GetBucketTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketTagsResponse) SetBody(v *GetBucketTagsResponseBody) *GetBucketTagsResponse {
	s.Body = v
	return s
}

type GetBucketTransferAccelerationResponseBody struct {
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
}

func (s GetBucketTransferAccelerationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketTransferAccelerationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketTransferAccelerationResponseBody) SetEnabled(v bool) *GetBucketTransferAccelerationResponseBody {
	s.Enabled = &v
	return s
}

type GetBucketTransferAccelerationResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketTransferAccelerationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketTransferAccelerationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketTransferAccelerationResponse) GoString() string {
	return s.String()
}

func (s *GetBucketTransferAccelerationResponse) SetHeaders(v map[string]*string) *GetBucketTransferAccelerationResponse {
	s.Headers = v
	return s
}

func (s *GetBucketTransferAccelerationResponse) SetStatusCode(v int32) *GetBucketTransferAccelerationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketTransferAccelerationResponse) SetBody(v *GetBucketTransferAccelerationResponseBody) *GetBucketTransferAccelerationResponse {
	s.Body = v
	return s
}

type GetBucketVersioningResponseBody struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetBucketVersioningResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketVersioningResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketVersioningResponseBody) SetStatus(v string) *GetBucketVersioningResponseBody {
	s.Status = &v
	return s
}

type GetBucketVersioningResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketVersioningResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketVersioningResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketVersioningResponse) GoString() string {
	return s.String()
}

func (s *GetBucketVersioningResponse) SetHeaders(v map[string]*string) *GetBucketVersioningResponse {
	s.Headers = v
	return s
}

func (s *GetBucketVersioningResponse) SetStatusCode(v int32) *GetBucketVersioningResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketVersioningResponse) SetBody(v *GetBucketVersioningResponseBody) *GetBucketVersioningResponse {
	s.Body = v
	return s
}

type GetBucketWebsiteResponseBody struct {
	ErrorDocument *ErrorDocument                            `json:"ErrorDocument,omitempty" xml:"ErrorDocument,omitempty"`
	IndexDocument *IndexDocument                            `json:"IndexDocument,omitempty" xml:"IndexDocument,omitempty"`
	RoutingRules  *GetBucketWebsiteResponseBodyRoutingRules `json:"RoutingRules,omitempty" xml:"RoutingRules,omitempty" type:"Struct"`
}

func (s GetBucketWebsiteResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketWebsiteResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketWebsiteResponseBody) SetErrorDocument(v *ErrorDocument) *GetBucketWebsiteResponseBody {
	s.ErrorDocument = v
	return s
}

func (s *GetBucketWebsiteResponseBody) SetIndexDocument(v *IndexDocument) *GetBucketWebsiteResponseBody {
	s.IndexDocument = v
	return s
}

func (s *GetBucketWebsiteResponseBody) SetRoutingRules(v *GetBucketWebsiteResponseBodyRoutingRules) *GetBucketWebsiteResponseBody {
	s.RoutingRules = v
	return s
}

type GetBucketWebsiteResponseBodyRoutingRules struct {
	RoutingRules []*RoutingRule `json:"RoutingRule,omitempty" xml:"RoutingRule,omitempty" type:"Repeated"`
}

func (s GetBucketWebsiteResponseBodyRoutingRules) String() string {
	return tea.Prettify(s)
}

func (s GetBucketWebsiteResponseBodyRoutingRules) GoString() string {
	return s.String()
}

func (s *GetBucketWebsiteResponseBodyRoutingRules) SetRoutingRules(v []*RoutingRule) *GetBucketWebsiteResponseBodyRoutingRules {
	s.RoutingRules = v
	return s
}

type GetBucketWebsiteResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketWebsiteResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketWebsiteResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketWebsiteResponse) GoString() string {
	return s.String()
}

func (s *GetBucketWebsiteResponse) SetHeaders(v map[string]*string) *GetBucketWebsiteResponse {
	s.Headers = v
	return s
}

func (s *GetBucketWebsiteResponse) SetStatusCode(v int32) *GetBucketWebsiteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketWebsiteResponse) SetBody(v *GetBucketWebsiteResponseBody) *GetBucketWebsiteResponse {
	s.Body = v
	return s
}

type GetBucketWormResponseBody struct {
	CreationDate          *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	RetentionPeriodInDays *int32  `json:"RetentionPeriodInDays,omitempty" xml:"RetentionPeriodInDays,omitempty"`
	State                 *string `json:"State,omitempty" xml:"State,omitempty"`
	WormId                *string `json:"WormId,omitempty" xml:"WormId,omitempty"`
}

func (s GetBucketWormResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetBucketWormResponseBody) GoString() string {
	return s.String()
}

func (s *GetBucketWormResponseBody) SetCreationDate(v string) *GetBucketWormResponseBody {
	s.CreationDate = &v
	return s
}

func (s *GetBucketWormResponseBody) SetRetentionPeriodInDays(v int32) *GetBucketWormResponseBody {
	s.RetentionPeriodInDays = &v
	return s
}

func (s *GetBucketWormResponseBody) SetState(v string) *GetBucketWormResponseBody {
	s.State = &v
	return s
}

func (s *GetBucketWormResponseBody) SetWormId(v string) *GetBucketWormResponseBody {
	s.WormId = &v
	return s
}

type GetBucketWormResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetBucketWormResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetBucketWormResponse) String() string {
	return tea.Prettify(s)
}

func (s GetBucketWormResponse) GoString() string {
	return s.String()
}

func (s *GetBucketWormResponse) SetHeaders(v map[string]*string) *GetBucketWormResponse {
	s.Headers = v
	return s
}

func (s *GetBucketWormResponse) SetStatusCode(v int32) *GetBucketWormResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBucketWormResponse) SetBody(v *GetBucketWormResponseBody) *GetBucketWormResponse {
	s.Body = v
	return s
}

type GetLiveChannelHistoryResponseBody struct {
	LiveRecords []*LiveRecord `json:"LiveRecord,omitempty" xml:"LiveRecord,omitempty" type:"Repeated"`
}

func (s GetLiveChannelHistoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveChannelHistoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveChannelHistoryResponseBody) SetLiveRecords(v []*LiveRecord) *GetLiveChannelHistoryResponseBody {
	s.LiveRecords = v
	return s
}

type GetLiveChannelHistoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLiveChannelHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveChannelHistoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveChannelHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetLiveChannelHistoryResponse) SetHeaders(v map[string]*string) *GetLiveChannelHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetLiveChannelHistoryResponse) SetStatusCode(v int32) *GetLiveChannelHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveChannelHistoryResponse) SetBody(v *GetLiveChannelHistoryResponseBody) *GetLiveChannelHistoryResponse {
	s.Body = v
	return s
}

type GetLiveChannelInfoResponseBody struct {
	Description *string            `json:"Description,omitempty" xml:"Description,omitempty"`
	Status      *string            `json:"Status,omitempty" xml:"Status,omitempty"`
	Target      *LiveChannelTarget `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s GetLiveChannelInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveChannelInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveChannelInfoResponseBody) SetDescription(v string) *GetLiveChannelInfoResponseBody {
	s.Description = &v
	return s
}

func (s *GetLiveChannelInfoResponseBody) SetStatus(v string) *GetLiveChannelInfoResponseBody {
	s.Status = &v
	return s
}

func (s *GetLiveChannelInfoResponseBody) SetTarget(v *LiveChannelTarget) *GetLiveChannelInfoResponseBody {
	s.Target = v
	return s
}

type GetLiveChannelInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLiveChannelInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveChannelInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveChannelInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLiveChannelInfoResponse) SetHeaders(v map[string]*string) *GetLiveChannelInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLiveChannelInfoResponse) SetStatusCode(v int32) *GetLiveChannelInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveChannelInfoResponse) SetBody(v *GetLiveChannelInfoResponseBody) *GetLiveChannelInfoResponse {
	s.Body = v
	return s
}

type GetLiveChannelStatResponseBody struct {
	Audio         *LiveChannelAudio `json:"Audio,omitempty" xml:"Audio,omitempty"`
	ConnectedTime *string           `json:"ConnectedTime,omitempty" xml:"ConnectedTime,omitempty"`
	RemoteAddr    *string           `json:"RemoteAddr,omitempty" xml:"RemoteAddr,omitempty"`
	Status        *string           `json:"Status,omitempty" xml:"Status,omitempty"`
	Video         *LiveChannelVideo `json:"Video,omitempty" xml:"Video,omitempty"`
}

func (s GetLiveChannelStatResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLiveChannelStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetLiveChannelStatResponseBody) SetAudio(v *LiveChannelAudio) *GetLiveChannelStatResponseBody {
	s.Audio = v
	return s
}

func (s *GetLiveChannelStatResponseBody) SetConnectedTime(v string) *GetLiveChannelStatResponseBody {
	s.ConnectedTime = &v
	return s
}

func (s *GetLiveChannelStatResponseBody) SetRemoteAddr(v string) *GetLiveChannelStatResponseBody {
	s.RemoteAddr = &v
	return s
}

func (s *GetLiveChannelStatResponseBody) SetStatus(v string) *GetLiveChannelStatResponseBody {
	s.Status = &v
	return s
}

func (s *GetLiveChannelStatResponseBody) SetVideo(v *LiveChannelVideo) *GetLiveChannelStatResponseBody {
	s.Video = v
	return s
}

type GetLiveChannelStatResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLiveChannelStatResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLiveChannelStatResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLiveChannelStatResponse) GoString() string {
	return s.String()
}

func (s *GetLiveChannelStatResponse) SetHeaders(v map[string]*string) *GetLiveChannelStatResponse {
	s.Headers = v
	return s
}

func (s *GetLiveChannelStatResponse) SetStatusCode(v int32) *GetLiveChannelStatResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLiveChannelStatResponse) SetBody(v *GetLiveChannelStatResponseBody) *GetLiveChannelStatResponse {
	s.Body = v
	return s
}

type GetObjectHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AcceptEncoding    *string            `json:"Accept-Encoding,omitempty" xml:"Accept-Encoding,omitempty"`
	IfMatch           *string            `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	IfModifiedSince   *string            `json:"If-Modified-Since,omitempty" xml:"If-Modified-Since,omitempty"`
	IfNoneMatch       *string            `json:"If-None-Match,omitempty" xml:"If-None-Match,omitempty"`
	IfUnmodifiedSince *string            `json:"If-Unmodified-Since,omitempty" xml:"If-Unmodified-Since,omitempty"`
	Range             *string            `json:"Range,omitempty" xml:"Range,omitempty"`
}

func (s GetObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s GetObjectHeaders) GoString() string {
	return s.String()
}

func (s *GetObjectHeaders) SetCommonHeaders(v map[string]*string) *GetObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *GetObjectHeaders) SetAcceptEncoding(v string) *GetObjectHeaders {
	s.AcceptEncoding = &v
	return s
}

func (s *GetObjectHeaders) SetIfMatch(v string) *GetObjectHeaders {
	s.IfMatch = &v
	return s
}

func (s *GetObjectHeaders) SetIfModifiedSince(v string) *GetObjectHeaders {
	s.IfModifiedSince = &v
	return s
}

func (s *GetObjectHeaders) SetIfNoneMatch(v string) *GetObjectHeaders {
	s.IfNoneMatch = &v
	return s
}

func (s *GetObjectHeaders) SetIfUnmodifiedSince(v string) *GetObjectHeaders {
	s.IfUnmodifiedSince = &v
	return s
}

func (s *GetObjectHeaders) SetRange(v string) *GetObjectHeaders {
	s.Range = &v
	return s
}

type GetObjectRequest struct {
	ResponseCacheControl       *string `json:"response-cache-control,omitempty" xml:"response-cache-control,omitempty"`
	ResponseContentDisposition *string `json:"response-content-disposition,omitempty" xml:"response-content-disposition,omitempty"`
	ResponseContentEncoding    *string `json:"response-content-encoding,omitempty" xml:"response-content-encoding,omitempty"`
	ResponseContentLanguage    *string `json:"response-content-language,omitempty" xml:"response-content-language,omitempty"`
	ResponseContentType        *string `json:"response-content-type,omitempty" xml:"response-content-type,omitempty"`
	ResponseExpires            *string `json:"response-expires,omitempty" xml:"response-expires,omitempty"`
}

func (s GetObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s GetObjectRequest) GoString() string {
	return s.String()
}

func (s *GetObjectRequest) SetResponseCacheControl(v string) *GetObjectRequest {
	s.ResponseCacheControl = &v
	return s
}

func (s *GetObjectRequest) SetResponseContentDisposition(v string) *GetObjectRequest {
	s.ResponseContentDisposition = &v
	return s
}

func (s *GetObjectRequest) SetResponseContentEncoding(v string) *GetObjectRequest {
	s.ResponseContentEncoding = &v
	return s
}

func (s *GetObjectRequest) SetResponseContentLanguage(v string) *GetObjectRequest {
	s.ResponseContentLanguage = &v
	return s
}

func (s *GetObjectRequest) SetResponseContentType(v string) *GetObjectRequest {
	s.ResponseContentType = &v
	return s
}

func (s *GetObjectRequest) SetResponseExpires(v string) *GetObjectRequest {
	s.ResponseExpires = &v
	return s
}

type GetObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       io.Reader          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s GetObjectResponse) GoString() string {
	return s.String()
}

func (s *GetObjectResponse) SetHeaders(v map[string]*string) *GetObjectResponse {
	s.Headers = v
	return s
}

func (s *GetObjectResponse) SetStatusCode(v int32) *GetObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectResponse) SetBody(v io.Reader) *GetObjectResponse {
	s.Body = v
	return s
}

type GetObjectAclRequest struct {
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetObjectAclRequest) String() string {
	return tea.Prettify(s)
}

func (s GetObjectAclRequest) GoString() string {
	return s.String()
}

func (s *GetObjectAclRequest) SetVersionId(v string) *GetObjectAclRequest {
	s.VersionId = &v
	return s
}

type GetObjectAclResponseBody struct {
	AccessControlList *GetObjectAclResponseBodyAccessControlList `json:"AccessControlList,omitempty" xml:"AccessControlList,omitempty" type:"Struct"`
	Owner             *Owner                                     `json:"Owner,omitempty" xml:"Owner,omitempty"`
}

func (s GetObjectAclResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetObjectAclResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectAclResponseBody) SetAccessControlList(v *GetObjectAclResponseBodyAccessControlList) *GetObjectAclResponseBody {
	s.AccessControlList = v
	return s
}

func (s *GetObjectAclResponseBody) SetOwner(v *Owner) *GetObjectAclResponseBody {
	s.Owner = v
	return s
}

type GetObjectAclResponseBodyAccessControlList struct {
	ACL *string `json:"Grant,omitempty" xml:"Grant,omitempty"`
}

func (s GetObjectAclResponseBodyAccessControlList) String() string {
	return tea.Prettify(s)
}

func (s GetObjectAclResponseBodyAccessControlList) GoString() string {
	return s.String()
}

func (s *GetObjectAclResponseBodyAccessControlList) SetACL(v string) *GetObjectAclResponseBodyAccessControlList {
	s.ACL = &v
	return s
}

type GetObjectAclResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetObjectAclResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetObjectAclResponse) String() string {
	return tea.Prettify(s)
}

func (s GetObjectAclResponse) GoString() string {
	return s.String()
}

func (s *GetObjectAclResponse) SetHeaders(v map[string]*string) *GetObjectAclResponse {
	s.Headers = v
	return s
}

func (s *GetObjectAclResponse) SetStatusCode(v int32) *GetObjectAclResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectAclResponse) SetBody(v *GetObjectAclResponseBody) *GetObjectAclResponse {
	s.Body = v
	return s
}

type GetObjectMetaRequest struct {
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetObjectMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetObjectMetaRequest) GoString() string {
	return s.String()
}

func (s *GetObjectMetaRequest) SetVersionId(v string) *GetObjectMetaRequest {
	s.VersionId = &v
	return s
}

type GetObjectMetaResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s GetObjectMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetObjectMetaResponse) GoString() string {
	return s.String()
}

func (s *GetObjectMetaResponse) SetHeaders(v map[string]*string) *GetObjectMetaResponse {
	s.Headers = v
	return s
}

func (s *GetObjectMetaResponse) SetStatusCode(v int32) *GetObjectMetaResponse {
	s.StatusCode = &v
	return s
}

type GetObjectTaggingRequest struct {
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetObjectTaggingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetObjectTaggingRequest) GoString() string {
	return s.String()
}

func (s *GetObjectTaggingRequest) SetVersionId(v string) *GetObjectTaggingRequest {
	s.VersionId = &v
	return s
}

type GetObjectTaggingResponseBody struct {
	TagSet *TagSet `json:"TagSet,omitempty" xml:"TagSet,omitempty"`
}

func (s GetObjectTaggingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetObjectTaggingResponseBody) GoString() string {
	return s.String()
}

func (s *GetObjectTaggingResponseBody) SetTagSet(v *TagSet) *GetObjectTaggingResponseBody {
	s.TagSet = v
	return s
}

type GetObjectTaggingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetObjectTaggingResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetObjectTaggingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetObjectTaggingResponse) GoString() string {
	return s.String()
}

func (s *GetObjectTaggingResponse) SetHeaders(v map[string]*string) *GetObjectTaggingResponse {
	s.Headers = v
	return s
}

func (s *GetObjectTaggingResponse) SetStatusCode(v int32) *GetObjectTaggingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetObjectTaggingResponse) SetBody(v *GetObjectTaggingResponseBody) *GetObjectTaggingResponse {
	s.Body = v
	return s
}

type GetSymlinkRequest struct {
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s GetSymlinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSymlinkRequest) GoString() string {
	return s.String()
}

func (s *GetSymlinkRequest) SetVersionId(v string) *GetSymlinkRequest {
	s.VersionId = &v
	return s
}

type GetSymlinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s GetSymlinkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSymlinkResponse) GoString() string {
	return s.String()
}

func (s *GetSymlinkResponse) SetHeaders(v map[string]*string) *GetSymlinkResponse {
	s.Headers = v
	return s
}

func (s *GetSymlinkResponse) SetStatusCode(v int32) *GetSymlinkResponse {
	s.StatusCode = &v
	return s
}

type GetVodPlaylistRequest struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s GetVodPlaylistRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVodPlaylistRequest) GoString() string {
	return s.String()
}

func (s *GetVodPlaylistRequest) SetEndTime(v string) *GetVodPlaylistRequest {
	s.EndTime = &v
	return s
}

func (s *GetVodPlaylistRequest) SetStartTime(v string) *GetVodPlaylistRequest {
	s.StartTime = &v
	return s
}

type GetVodPlaylistResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       io.Reader          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVodPlaylistResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVodPlaylistResponse) GoString() string {
	return s.String()
}

func (s *GetVodPlaylistResponse) SetHeaders(v map[string]*string) *GetVodPlaylistResponse {
	s.Headers = v
	return s
}

func (s *GetVodPlaylistResponse) SetStatusCode(v int32) *GetVodPlaylistResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVodPlaylistResponse) SetBody(v io.Reader) *GetVodPlaylistResponse {
	s.Body = v
	return s
}

type HeadObjectHeaders struct {
	CommonHeaders     map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	IfMatch           *string            `json:"If-Match,omitempty" xml:"If-Match,omitempty"`
	IfModifiedSince   *string            `json:"If-Modified-Since,omitempty" xml:"If-Modified-Since,omitempty"`
	IfNoneMatch       *string            `json:"If-None-Match,omitempty" xml:"If-None-Match,omitempty"`
	IfUnmodifiedSince *string            `json:"If-Unmodified-Since,omitempty" xml:"If-Unmodified-Since,omitempty"`
}

func (s HeadObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s HeadObjectHeaders) GoString() string {
	return s.String()
}

func (s *HeadObjectHeaders) SetCommonHeaders(v map[string]*string) *HeadObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *HeadObjectHeaders) SetIfMatch(v string) *HeadObjectHeaders {
	s.IfMatch = &v
	return s
}

func (s *HeadObjectHeaders) SetIfModifiedSince(v string) *HeadObjectHeaders {
	s.IfModifiedSince = &v
	return s
}

func (s *HeadObjectHeaders) SetIfNoneMatch(v string) *HeadObjectHeaders {
	s.IfNoneMatch = &v
	return s
}

func (s *HeadObjectHeaders) SetIfUnmodifiedSince(v string) *HeadObjectHeaders {
	s.IfUnmodifiedSince = &v
	return s
}

type HeadObjectRequest struct {
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s HeadObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s HeadObjectRequest) GoString() string {
	return s.String()
}

func (s *HeadObjectRequest) SetVersionId(v string) *HeadObjectRequest {
	s.VersionId = &v
	return s
}

type HeadObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s HeadObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s HeadObjectResponse) GoString() string {
	return s.String()
}

func (s *HeadObjectResponse) SetHeaders(v map[string]*string) *HeadObjectResponse {
	s.Headers = v
	return s
}

func (s *HeadObjectResponse) SetStatusCode(v int32) *HeadObjectResponse {
	s.StatusCode = &v
	return s
}

type InitiateBucketWormRequest struct {
	InitiateWormConfiguration *InitiateWormConfiguration `json:"InitiateWormConfiguration,omitempty" xml:"InitiateWormConfiguration,omitempty"`
}

func (s InitiateBucketWormRequest) String() string {
	return tea.Prettify(s)
}

func (s InitiateBucketWormRequest) GoString() string {
	return s.String()
}

func (s *InitiateBucketWormRequest) SetInitiateWormConfiguration(v *InitiateWormConfiguration) *InitiateBucketWormRequest {
	s.InitiateWormConfiguration = v
	return s
}

type InitiateBucketWormResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s InitiateBucketWormResponse) String() string {
	return tea.Prettify(s)
}

func (s InitiateBucketWormResponse) GoString() string {
	return s.String()
}

func (s *InitiateBucketWormResponse) SetHeaders(v map[string]*string) *InitiateBucketWormResponse {
	s.Headers = v
	return s
}

func (s *InitiateBucketWormResponse) SetStatusCode(v int32) *InitiateBucketWormResponse {
	s.StatusCode = &v
	return s
}

type InitiateMultipartUploadHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	CacheControl         *string            `json:"Cache-Control,omitempty" xml:"Cache-Control,omitempty"`
	ContentDisposition   *string            `json:"Content-Disposition,omitempty" xml:"Content-Disposition,omitempty"`
	ContentEncoding      *string            `json:"Content-Encoding,omitempty" xml:"Content-Encoding,omitempty"`
	Expires              *string            `json:"Expires,omitempty" xml:"Expires,omitempty"`
	ForbidOverwrite      *string            `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
	SseDataEncryption    *string            `json:"x-oss-server-side-data-encryption,omitempty" xml:"x-oss-server-side-data-encryption,omitempty"`
	ServerSideEncryption *string            `json:"x-oss-server-side-encryption,omitempty" xml:"x-oss-server-side-encryption,omitempty"`
	SseKeyId             *string            `json:"x-oss-server-side-encryption-key-id,omitempty" xml:"x-oss-server-side-encryption-key-id,omitempty"`
	StorageClass         *string            `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
	Tagging              *string            `json:"x-oss-tagging,omitempty" xml:"x-oss-tagging,omitempty"`
}

func (s InitiateMultipartUploadHeaders) String() string {
	return tea.Prettify(s)
}

func (s InitiateMultipartUploadHeaders) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadHeaders) SetCommonHeaders(v map[string]*string) *InitiateMultipartUploadHeaders {
	s.CommonHeaders = v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetCacheControl(v string) *InitiateMultipartUploadHeaders {
	s.CacheControl = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetContentDisposition(v string) *InitiateMultipartUploadHeaders {
	s.ContentDisposition = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetContentEncoding(v string) *InitiateMultipartUploadHeaders {
	s.ContentEncoding = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetExpires(v string) *InitiateMultipartUploadHeaders {
	s.Expires = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetForbidOverwrite(v string) *InitiateMultipartUploadHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetSseDataEncryption(v string) *InitiateMultipartUploadHeaders {
	s.SseDataEncryption = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetServerSideEncryption(v string) *InitiateMultipartUploadHeaders {
	s.ServerSideEncryption = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetSseKeyId(v string) *InitiateMultipartUploadHeaders {
	s.SseKeyId = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetStorageClass(v string) *InitiateMultipartUploadHeaders {
	s.StorageClass = &v
	return s
}

func (s *InitiateMultipartUploadHeaders) SetTagging(v string) *InitiateMultipartUploadHeaders {
	s.Tagging = &v
	return s
}

type InitiateMultipartUploadRequest struct {
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
}

func (s InitiateMultipartUploadRequest) String() string {
	return tea.Prettify(s)
}

func (s InitiateMultipartUploadRequest) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadRequest) SetEncodingType(v string) *InitiateMultipartUploadRequest {
	s.EncodingType = &v
	return s
}

type InitiateMultipartUploadResponseBody struct {
	Bucket       *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	EncodingType *string `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	Key          *string `json:"Key,omitempty" xml:"Key,omitempty"`
	UploadId     *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s InitiateMultipartUploadResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InitiateMultipartUploadResponseBody) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadResponseBody) SetBucket(v string) *InitiateMultipartUploadResponseBody {
	s.Bucket = &v
	return s
}

func (s *InitiateMultipartUploadResponseBody) SetEncodingType(v string) *InitiateMultipartUploadResponseBody {
	s.EncodingType = &v
	return s
}

func (s *InitiateMultipartUploadResponseBody) SetKey(v string) *InitiateMultipartUploadResponseBody {
	s.Key = &v
	return s
}

func (s *InitiateMultipartUploadResponseBody) SetUploadId(v string) *InitiateMultipartUploadResponseBody {
	s.UploadId = &v
	return s
}

type InitiateMultipartUploadResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *InitiateMultipartUploadResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InitiateMultipartUploadResponse) String() string {
	return tea.Prettify(s)
}

func (s InitiateMultipartUploadResponse) GoString() string {
	return s.String()
}

func (s *InitiateMultipartUploadResponse) SetHeaders(v map[string]*string) *InitiateMultipartUploadResponse {
	s.Headers = v
	return s
}

func (s *InitiateMultipartUploadResponse) SetStatusCode(v int32) *InitiateMultipartUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *InitiateMultipartUploadResponse) SetBody(v *InitiateMultipartUploadResponseBody) *InitiateMultipartUploadResponse {
	s.Body = v
	return s
}

type ListBucketInventoryRequest struct {
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
}

func (s ListBucketInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBucketInventoryRequest) GoString() string {
	return s.String()
}

func (s *ListBucketInventoryRequest) SetContinuationToken(v string) *ListBucketInventoryRequest {
	s.ContinuationToken = &v
	return s
}

type ListBucketInventoryResponseBody struct {
	InventoryConfigurations []*InventoryConfiguration `json:"InventoryConfiguration,omitempty" xml:"InventoryConfiguration,omitempty" type:"Repeated"`
	IsTruncated             *bool                     `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	NextContinuationToken   *string                   `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
}

func (s ListBucketInventoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBucketInventoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketInventoryResponseBody) SetInventoryConfigurations(v []*InventoryConfiguration) *ListBucketInventoryResponseBody {
	s.InventoryConfigurations = v
	return s
}

func (s *ListBucketInventoryResponseBody) SetIsTruncated(v bool) *ListBucketInventoryResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListBucketInventoryResponseBody) SetNextContinuationToken(v string) *ListBucketInventoryResponseBody {
	s.NextContinuationToken = &v
	return s
}

type ListBucketInventoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBucketInventoryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBucketInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBucketInventoryResponse) GoString() string {
	return s.String()
}

func (s *ListBucketInventoryResponse) SetHeaders(v map[string]*string) *ListBucketInventoryResponse {
	s.Headers = v
	return s
}

func (s *ListBucketInventoryResponse) SetStatusCode(v int32) *ListBucketInventoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucketInventoryResponse) SetBody(v *ListBucketInventoryResponseBody) *ListBucketInventoryResponse {
	s.Body = v
	return s
}

type ListBucketsRequest struct {
	Marker  *string `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxKeys *int64  `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	Prefix  *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListBucketsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListBucketsRequest) GoString() string {
	return s.String()
}

func (s *ListBucketsRequest) SetMarker(v string) *ListBucketsRequest {
	s.Marker = &v
	return s
}

func (s *ListBucketsRequest) SetMaxKeys(v int64) *ListBucketsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListBucketsRequest) SetPrefix(v string) *ListBucketsRequest {
	s.Prefix = &v
	return s
}

type ListBucketsResponseBody struct {
	Buckets     *ListBucketsResponseBodyBuckets `json:"Buckets,omitempty" xml:"Buckets,omitempty" type:"Struct"`
	IsTruncated *bool                           `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker      *string                         `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys     *int64                          `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	NextMarker  *string                         `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Owner       *Owner                          `json:"Owner,omitempty" xml:"Owner,omitempty"`
	Prefix      *string                         `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListBucketsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListBucketsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBody) SetBuckets(v *ListBucketsResponseBodyBuckets) *ListBucketsResponseBody {
	s.Buckets = v
	return s
}

func (s *ListBucketsResponseBody) SetIsTruncated(v bool) *ListBucketsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListBucketsResponseBody) SetMarker(v string) *ListBucketsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListBucketsResponseBody) SetMaxKeys(v int64) *ListBucketsResponseBody {
	s.MaxKeys = &v
	return s
}

func (s *ListBucketsResponseBody) SetNextMarker(v string) *ListBucketsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListBucketsResponseBody) SetOwner(v *Owner) *ListBucketsResponseBody {
	s.Owner = v
	return s
}

func (s *ListBucketsResponseBody) SetPrefix(v string) *ListBucketsResponseBody {
	s.Prefix = &v
	return s
}

type ListBucketsResponseBodyBuckets struct {
	Buckets []*Bucket `json:"Bucket,omitempty" xml:"Bucket,omitempty" type:"Repeated"`
}

func (s ListBucketsResponseBodyBuckets) String() string {
	return tea.Prettify(s)
}

func (s ListBucketsResponseBodyBuckets) GoString() string {
	return s.String()
}

func (s *ListBucketsResponseBodyBuckets) SetBuckets(v []*Bucket) *ListBucketsResponseBodyBuckets {
	s.Buckets = v
	return s
}

type ListBucketsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListBucketsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListBucketsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListBucketsResponse) GoString() string {
	return s.String()
}

func (s *ListBucketsResponse) SetHeaders(v map[string]*string) *ListBucketsResponse {
	s.Headers = v
	return s
}

func (s *ListBucketsResponse) SetStatusCode(v int32) *ListBucketsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBucketsResponse) SetBody(v *ListBucketsResponseBody) *ListBucketsResponse {
	s.Body = v
	return s
}

type ListLiveChannelRequest struct {
	Marker  *string `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxKeys *int64  `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	Prefix  *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListLiveChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s ListLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *ListLiveChannelRequest) SetMarker(v string) *ListLiveChannelRequest {
	s.Marker = &v
	return s
}

func (s *ListLiveChannelRequest) SetMaxKeys(v int64) *ListLiveChannelRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListLiveChannelRequest) SetPrefix(v string) *ListLiveChannelRequest {
	s.Prefix = &v
	return s
}

type ListLiveChannelResponseBody struct {
	IsTruncated  *bool          `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	LiveChannels []*LiveChannel `json:"LiveChannel,omitempty" xml:"LiveChannel,omitempty" type:"Repeated"`
	Marker       *string        `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys      *int64         `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	NextMarker   *string        `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Prefix       *string        `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListLiveChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *ListLiveChannelResponseBody) SetIsTruncated(v bool) *ListLiveChannelResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListLiveChannelResponseBody) SetLiveChannels(v []*LiveChannel) *ListLiveChannelResponseBody {
	s.LiveChannels = v
	return s
}

func (s *ListLiveChannelResponseBody) SetMarker(v string) *ListLiveChannelResponseBody {
	s.Marker = &v
	return s
}

func (s *ListLiveChannelResponseBody) SetMaxKeys(v int64) *ListLiveChannelResponseBody {
	s.MaxKeys = &v
	return s
}

func (s *ListLiveChannelResponseBody) SetNextMarker(v string) *ListLiveChannelResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListLiveChannelResponseBody) SetPrefix(v string) *ListLiveChannelResponseBody {
	s.Prefix = &v
	return s
}

type ListLiveChannelResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListLiveChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s ListLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *ListLiveChannelResponse) SetHeaders(v map[string]*string) *ListLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *ListLiveChannelResponse) SetStatusCode(v int32) *ListLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLiveChannelResponse) SetBody(v *ListLiveChannelResponseBody) *ListLiveChannelResponse {
	s.Body = v
	return s
}

type ListMultipartUploadsRequest struct {
	Delimiter      *string `json:"delimiter,omitempty" xml:"delimiter,omitempty"`
	EncodingType   *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	KeyMarker      *string `json:"key-marker,omitempty" xml:"key-marker,omitempty"`
	MaxUploads     *int64  `json:"max-uploads,omitempty" xml:"max-uploads,omitempty"`
	Prefix         *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	UploadIdMarker *string `json:"upload-id-marker,omitempty" xml:"upload-id-marker,omitempty"`
}

func (s ListMultipartUploadsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMultipartUploadsRequest) GoString() string {
	return s.String()
}

func (s *ListMultipartUploadsRequest) SetDelimiter(v string) *ListMultipartUploadsRequest {
	s.Delimiter = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetEncodingType(v string) *ListMultipartUploadsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetKeyMarker(v string) *ListMultipartUploadsRequest {
	s.KeyMarker = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetMaxUploads(v int64) *ListMultipartUploadsRequest {
	s.MaxUploads = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetPrefix(v string) *ListMultipartUploadsRequest {
	s.Prefix = &v
	return s
}

func (s *ListMultipartUploadsRequest) SetUploadIdMarker(v string) *ListMultipartUploadsRequest {
	s.UploadIdMarker = &v
	return s
}

type ListMultipartUploadsResponseBody struct {
	Bucket             *string         `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	CommonPrefixes     []*CommonPrefix `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	Delimiter          *string         `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	EncodingType       *string         `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	IsTruncated        *bool           `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	KeyMarker          *string         `json:"KeyMarker,omitempty" xml:"KeyMarker,omitempty"`
	MaxUploads         *int64          `json:"MaxUploads,omitempty" xml:"MaxUploads,omitempty"`
	NextKeyMarker      *string         `json:"NextKeyMarker,omitempty" xml:"NextKeyMarker,omitempty"`
	NextUploadIdMarker *string         `json:"NextUploadIdMarker,omitempty" xml:"NextUploadIdMarker,omitempty"`
	Prefix             *string         `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Uploads            []*Upload       `json:"Upload,omitempty" xml:"Upload,omitempty" type:"Repeated"`
	UploadIdMarker     *string         `json:"UploadIdMarker,omitempty" xml:"UploadIdMarker,omitempty"`
}

func (s ListMultipartUploadsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMultipartUploadsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMultipartUploadsResponseBody) SetBucket(v string) *ListMultipartUploadsResponseBody {
	s.Bucket = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetCommonPrefixes(v []*CommonPrefix) *ListMultipartUploadsResponseBody {
	s.CommonPrefixes = v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetDelimiter(v string) *ListMultipartUploadsResponseBody {
	s.Delimiter = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetEncodingType(v string) *ListMultipartUploadsResponseBody {
	s.EncodingType = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetIsTruncated(v bool) *ListMultipartUploadsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetKeyMarker(v string) *ListMultipartUploadsResponseBody {
	s.KeyMarker = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetMaxUploads(v int64) *ListMultipartUploadsResponseBody {
	s.MaxUploads = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetNextKeyMarker(v string) *ListMultipartUploadsResponseBody {
	s.NextKeyMarker = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetNextUploadIdMarker(v string) *ListMultipartUploadsResponseBody {
	s.NextUploadIdMarker = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetPrefix(v string) *ListMultipartUploadsResponseBody {
	s.Prefix = &v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetUploads(v []*Upload) *ListMultipartUploadsResponseBody {
	s.Uploads = v
	return s
}

func (s *ListMultipartUploadsResponseBody) SetUploadIdMarker(v string) *ListMultipartUploadsResponseBody {
	s.UploadIdMarker = &v
	return s
}

type ListMultipartUploadsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMultipartUploadsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMultipartUploadsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMultipartUploadsResponse) GoString() string {
	return s.String()
}

func (s *ListMultipartUploadsResponse) SetHeaders(v map[string]*string) *ListMultipartUploadsResponse {
	s.Headers = v
	return s
}

func (s *ListMultipartUploadsResponse) SetStatusCode(v int32) *ListMultipartUploadsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultipartUploadsResponse) SetBody(v *ListMultipartUploadsResponseBody) *ListMultipartUploadsResponse {
	s.Body = v
	return s
}

type ListObjectVersionsRequest struct {
	Delimiter       *string `json:"delimiter,omitempty" xml:"delimiter,omitempty"`
	EncodingType    *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	KeyMarker       *string `json:"key-marker,omitempty" xml:"key-marker,omitempty"`
	MaxKeys         *int64  `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	Prefix          *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	VersionIdMarker *string `json:"version-id-marker,omitempty" xml:"version-id-marker,omitempty"`
}

func (s ListObjectVersionsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListObjectVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListObjectVersionsRequest) SetDelimiter(v string) *ListObjectVersionsRequest {
	s.Delimiter = &v
	return s
}

func (s *ListObjectVersionsRequest) SetEncodingType(v string) *ListObjectVersionsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListObjectVersionsRequest) SetKeyMarker(v string) *ListObjectVersionsRequest {
	s.KeyMarker = &v
	return s
}

func (s *ListObjectVersionsRequest) SetMaxKeys(v int64) *ListObjectVersionsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectVersionsRequest) SetPrefix(v string) *ListObjectVersionsRequest {
	s.Prefix = &v
	return s
}

func (s *ListObjectVersionsRequest) SetVersionIdMarker(v string) *ListObjectVersionsRequest {
	s.VersionIdMarker = &v
	return s
}

type ListObjectVersionsResponseBody struct {
	CommonPrefixes      []*CommonPrefix      `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	DeleteMarkers       []*DeleteMarkerEntry `json:"DeleteMarker,omitempty" xml:"DeleteMarker,omitempty" type:"Repeated"`
	Delimiter           *string              `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	EncodingType        *string              `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	IsTruncated         *bool                `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	KeyMarker           *string              `json:"KeyMarker,omitempty" xml:"KeyMarker,omitempty"`
	MaxKeys             *int64               `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	Name                *string              `json:"Name,omitempty" xml:"Name,omitempty"`
	NextKeyMarker       *string              `json:"NextKeyMarker,omitempty" xml:"NextKeyMarker,omitempty"`
	NextVersionIdMarker *string              `json:"NextVersionIdMarker,omitempty" xml:"NextVersionIdMarker,omitempty"`
	Prefix              *string              `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	Versions            []*ObjectVersion     `json:"Version,omitempty" xml:"Version,omitempty" type:"Repeated"`
	VersionIdMarker     *string              `json:"VersionIdMarker,omitempty" xml:"VersionIdMarker,omitempty"`
}

func (s ListObjectVersionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListObjectVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectVersionsResponseBody) SetCommonPrefixes(v []*CommonPrefix) *ListObjectVersionsResponseBody {
	s.CommonPrefixes = v
	return s
}

func (s *ListObjectVersionsResponseBody) SetDeleteMarkers(v []*DeleteMarkerEntry) *ListObjectVersionsResponseBody {
	s.DeleteMarkers = v
	return s
}

func (s *ListObjectVersionsResponseBody) SetDelimiter(v string) *ListObjectVersionsResponseBody {
	s.Delimiter = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetEncodingType(v string) *ListObjectVersionsResponseBody {
	s.EncodingType = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetIsTruncated(v bool) *ListObjectVersionsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetKeyMarker(v string) *ListObjectVersionsResponseBody {
	s.KeyMarker = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetMaxKeys(v int64) *ListObjectVersionsResponseBody {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetName(v string) *ListObjectVersionsResponseBody {
	s.Name = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetNextKeyMarker(v string) *ListObjectVersionsResponseBody {
	s.NextKeyMarker = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetNextVersionIdMarker(v string) *ListObjectVersionsResponseBody {
	s.NextVersionIdMarker = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetPrefix(v string) *ListObjectVersionsResponseBody {
	s.Prefix = &v
	return s
}

func (s *ListObjectVersionsResponseBody) SetVersions(v []*ObjectVersion) *ListObjectVersionsResponseBody {
	s.Versions = v
	return s
}

func (s *ListObjectVersionsResponseBody) SetVersionIdMarker(v string) *ListObjectVersionsResponseBody {
	s.VersionIdMarker = &v
	return s
}

type ListObjectVersionsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListObjectVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListObjectVersionsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListObjectVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListObjectVersionsResponse) SetHeaders(v map[string]*string) *ListObjectVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListObjectVersionsResponse) SetStatusCode(v int32) *ListObjectVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListObjectVersionsResponse) SetBody(v *ListObjectVersionsResponseBody) *ListObjectVersionsResponse {
	s.Body = v
	return s
}

type ListObjectsRequest struct {
	Delimiter    *string `json:"delimiter,omitempty" xml:"delimiter,omitempty"`
	EncodingType *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	Marker       *string `json:"marker,omitempty" xml:"marker,omitempty"`
	MaxKeys      *int64  `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	Prefix       *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
}

func (s ListObjectsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsRequest) GoString() string {
	return s.String()
}

func (s *ListObjectsRequest) SetDelimiter(v string) *ListObjectsRequest {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsRequest) SetEncodingType(v string) *ListObjectsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsRequest) SetMarker(v string) *ListObjectsRequest {
	s.Marker = &v
	return s
}

func (s *ListObjectsRequest) SetMaxKeys(v int64) *ListObjectsRequest {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsRequest) SetPrefix(v string) *ListObjectsRequest {
	s.Prefix = &v
	return s
}

type ListObjectsResponseBody struct {
	CommonPrefixes []*CommonPrefix  `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	Contents       []*ObjectSummary `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	Delimiter      *string          `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	EncodingType   *string          `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	IsTruncated    *bool            `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Marker         *string          `json:"Marker,omitempty" xml:"Marker,omitempty"`
	MaxKeys        *int32           `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	Name           *string          `json:"Name,omitempty" xml:"Name,omitempty"`
	NextMarker     *string          `json:"NextMarker,omitempty" xml:"NextMarker,omitempty"`
	Prefix         *string          `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
}

func (s ListObjectsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectsResponseBody) SetCommonPrefixes(v []*CommonPrefix) *ListObjectsResponseBody {
	s.CommonPrefixes = v
	return s
}

func (s *ListObjectsResponseBody) SetContents(v []*ObjectSummary) *ListObjectsResponseBody {
	s.Contents = v
	return s
}

func (s *ListObjectsResponseBody) SetDelimiter(v string) *ListObjectsResponseBody {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsResponseBody) SetEncodingType(v string) *ListObjectsResponseBody {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsResponseBody) SetIsTruncated(v bool) *ListObjectsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListObjectsResponseBody) SetMarker(v string) *ListObjectsResponseBody {
	s.Marker = &v
	return s
}

func (s *ListObjectsResponseBody) SetMaxKeys(v int32) *ListObjectsResponseBody {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsResponseBody) SetName(v string) *ListObjectsResponseBody {
	s.Name = &v
	return s
}

func (s *ListObjectsResponseBody) SetNextMarker(v string) *ListObjectsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListObjectsResponseBody) SetPrefix(v string) *ListObjectsResponseBody {
	s.Prefix = &v
	return s
}

type ListObjectsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListObjectsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsResponse) GoString() string {
	return s.String()
}

func (s *ListObjectsResponse) SetHeaders(v map[string]*string) *ListObjectsResponse {
	s.Headers = v
	return s
}

func (s *ListObjectsResponse) SetStatusCode(v int32) *ListObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListObjectsResponse) SetBody(v *ListObjectsResponseBody) *ListObjectsResponse {
	s.Body = v
	return s
}

type ListObjectsV2Request struct {
	ContinuationToken *string `json:"continuation-token,omitempty" xml:"continuation-token,omitempty"`
	Delimiter         *string `json:"delimiter,omitempty" xml:"delimiter,omitempty"`
	EncodingType      *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	FetchOwner        *bool   `json:"fetch-owner,omitempty" xml:"fetch-owner,omitempty"`
	MaxKeys           *int64  `json:"max-keys,omitempty" xml:"max-keys,omitempty"`
	Prefix            *string `json:"prefix,omitempty" xml:"prefix,omitempty"`
	StartAfter        *string `json:"start-after,omitempty" xml:"start-after,omitempty"`
}

func (s ListObjectsV2Request) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsV2Request) GoString() string {
	return s.String()
}

func (s *ListObjectsV2Request) SetContinuationToken(v string) *ListObjectsV2Request {
	s.ContinuationToken = &v
	return s
}

func (s *ListObjectsV2Request) SetDelimiter(v string) *ListObjectsV2Request {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsV2Request) SetEncodingType(v string) *ListObjectsV2Request {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsV2Request) SetFetchOwner(v bool) *ListObjectsV2Request {
	s.FetchOwner = &v
	return s
}

func (s *ListObjectsV2Request) SetMaxKeys(v int64) *ListObjectsV2Request {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsV2Request) SetPrefix(v string) *ListObjectsV2Request {
	s.Prefix = &v
	return s
}

func (s *ListObjectsV2Request) SetStartAfter(v string) *ListObjectsV2Request {
	s.StartAfter = &v
	return s
}

type ListObjectsV2ResponseBody struct {
	CommonPrefixes        []*CommonPrefix  `json:"CommonPrefixes,omitempty" xml:"CommonPrefixes,omitempty" type:"Repeated"`
	Contents              []*ObjectSummary `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
	ContinuationToken     *string          `json:"ContinuationToken,omitempty" xml:"ContinuationToken,omitempty"`
	Delimiter             *string          `json:"Delimiter,omitempty" xml:"Delimiter,omitempty"`
	EncodingType          *string          `json:"EncodingType,omitempty" xml:"EncodingType,omitempty"`
	IsTruncated           *bool            `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	KeyCount              *int32           `json:"KeyCount,omitempty" xml:"KeyCount,omitempty"`
	MaxKeys               *int32           `json:"MaxKeys,omitempty" xml:"MaxKeys,omitempty"`
	Name                  *string          `json:"Name,omitempty" xml:"Name,omitempty"`
	NextContinuationToken *string          `json:"NextContinuationToken,omitempty" xml:"NextContinuationToken,omitempty"`
	Prefix                *string          `json:"Prefix,omitempty" xml:"Prefix,omitempty"`
	StartAfter            *string          `json:"StartAfter,omitempty" xml:"StartAfter,omitempty"`
}

func (s ListObjectsV2ResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsV2ResponseBody) GoString() string {
	return s.String()
}

func (s *ListObjectsV2ResponseBody) SetCommonPrefixes(v []*CommonPrefix) *ListObjectsV2ResponseBody {
	s.CommonPrefixes = v
	return s
}

func (s *ListObjectsV2ResponseBody) SetContents(v []*ObjectSummary) *ListObjectsV2ResponseBody {
	s.Contents = v
	return s
}

func (s *ListObjectsV2ResponseBody) SetContinuationToken(v string) *ListObjectsV2ResponseBody {
	s.ContinuationToken = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetDelimiter(v string) *ListObjectsV2ResponseBody {
	s.Delimiter = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetEncodingType(v string) *ListObjectsV2ResponseBody {
	s.EncodingType = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetIsTruncated(v bool) *ListObjectsV2ResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetKeyCount(v int32) *ListObjectsV2ResponseBody {
	s.KeyCount = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetMaxKeys(v int32) *ListObjectsV2ResponseBody {
	s.MaxKeys = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetName(v string) *ListObjectsV2ResponseBody {
	s.Name = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetNextContinuationToken(v string) *ListObjectsV2ResponseBody {
	s.NextContinuationToken = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetPrefix(v string) *ListObjectsV2ResponseBody {
	s.Prefix = &v
	return s
}

func (s *ListObjectsV2ResponseBody) SetStartAfter(v string) *ListObjectsV2ResponseBody {
	s.StartAfter = &v
	return s
}

type ListObjectsV2Response struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListObjectsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListObjectsV2Response) String() string {
	return tea.Prettify(s)
}

func (s ListObjectsV2Response) GoString() string {
	return s.String()
}

func (s *ListObjectsV2Response) SetHeaders(v map[string]*string) *ListObjectsV2Response {
	s.Headers = v
	return s
}

func (s *ListObjectsV2Response) SetStatusCode(v int32) *ListObjectsV2Response {
	s.StatusCode = &v
	return s
}

func (s *ListObjectsV2Response) SetBody(v *ListObjectsV2ResponseBody) *ListObjectsV2Response {
	s.Body = v
	return s
}

type ListPartsRequest struct {
	EncodingType     *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	MaxParts         *int64  `json:"max-parts,omitempty" xml:"max-parts,omitempty"`
	PartNumberMarker *int64  `json:"part-number-marker,omitempty" xml:"part-number-marker,omitempty"`
	UploadId         *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s ListPartsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPartsRequest) GoString() string {
	return s.String()
}

func (s *ListPartsRequest) SetEncodingType(v string) *ListPartsRequest {
	s.EncodingType = &v
	return s
}

func (s *ListPartsRequest) SetMaxParts(v int64) *ListPartsRequest {
	s.MaxParts = &v
	return s
}

func (s *ListPartsRequest) SetPartNumberMarker(v int64) *ListPartsRequest {
	s.PartNumberMarker = &v
	return s
}

func (s *ListPartsRequest) SetUploadId(v string) *ListPartsRequest {
	s.UploadId = &v
	return s
}

type ListPartsShrinkRequest struct {
	EncodingTypeShrink *string `json:"encoding-type,omitempty" xml:"encoding-type,omitempty"`
	MaxParts           *int64  `json:"max-parts,omitempty" xml:"max-parts,omitempty"`
	PartNumberMarker   *int64  `json:"part-number-marker,omitempty" xml:"part-number-marker,omitempty"`
	UploadId           *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s ListPartsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPartsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListPartsShrinkRequest) SetEncodingTypeShrink(v string) *ListPartsShrinkRequest {
	s.EncodingTypeShrink = &v
	return s
}

func (s *ListPartsShrinkRequest) SetMaxParts(v int64) *ListPartsShrinkRequest {
	s.MaxParts = &v
	return s
}

func (s *ListPartsShrinkRequest) SetPartNumberMarker(v int64) *ListPartsShrinkRequest {
	s.PartNumberMarker = &v
	return s
}

func (s *ListPartsShrinkRequest) SetUploadId(v string) *ListPartsShrinkRequest {
	s.UploadId = &v
	return s
}

type ListPartsResponseBody struct {
	Bucket               *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	IsTruncated          *bool   `json:"IsTruncated,omitempty" xml:"IsTruncated,omitempty"`
	Key                  *string `json:"Key,omitempty" xml:"Key,omitempty"`
	MaxParts             *int64  `json:"MaxParts,omitempty" xml:"MaxParts,omitempty"`
	NextPartNumberMarker *int64  `json:"NextPartNumberMarker,omitempty" xml:"NextPartNumberMarker,omitempty"`
	Part                 []*Part `json:"Part,omitempty" xml:"Part,omitempty" type:"Repeated"`
	PartNumberMarker     *int64  `json:"PartNumberMarker,omitempty" xml:"PartNumberMarker,omitempty"`
	UploadId             *string `json:"UploadId,omitempty" xml:"UploadId,omitempty"`
}

func (s ListPartsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPartsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPartsResponseBody) SetBucket(v string) *ListPartsResponseBody {
	s.Bucket = &v
	return s
}

func (s *ListPartsResponseBody) SetIsTruncated(v bool) *ListPartsResponseBody {
	s.IsTruncated = &v
	return s
}

func (s *ListPartsResponseBody) SetKey(v string) *ListPartsResponseBody {
	s.Key = &v
	return s
}

func (s *ListPartsResponseBody) SetMaxParts(v int64) *ListPartsResponseBody {
	s.MaxParts = &v
	return s
}

func (s *ListPartsResponseBody) SetNextPartNumberMarker(v int64) *ListPartsResponseBody {
	s.NextPartNumberMarker = &v
	return s
}

func (s *ListPartsResponseBody) SetPart(v []*Part) *ListPartsResponseBody {
	s.Part = v
	return s
}

func (s *ListPartsResponseBody) SetPartNumberMarker(v int64) *ListPartsResponseBody {
	s.PartNumberMarker = &v
	return s
}

func (s *ListPartsResponseBody) SetUploadId(v string) *ListPartsResponseBody {
	s.UploadId = &v
	return s
}

type ListPartsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListPartsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPartsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPartsResponse) GoString() string {
	return s.String()
}

func (s *ListPartsResponse) SetHeaders(v map[string]*string) *ListPartsResponse {
	s.Headers = v
	return s
}

func (s *ListPartsResponse) SetStatusCode(v int32) *ListPartsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPartsResponse) SetBody(v *ListPartsResponseBody) *ListPartsResponse {
	s.Body = v
	return s
}

type OptionObjectHeaders struct {
	CommonHeaders               map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	AccessControlRequestHeaders *string            `json:"Access-Control-Request-Headers,omitempty" xml:"Access-Control-Request-Headers,omitempty"`
	AccessControlRequestMethod  *string            `json:"Access-Control-Request-Method,omitempty" xml:"Access-Control-Request-Method,omitempty"`
	Origin                      *string            `json:"Origin,omitempty" xml:"Origin,omitempty"`
}

func (s OptionObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s OptionObjectHeaders) GoString() string {
	return s.String()
}

func (s *OptionObjectHeaders) SetCommonHeaders(v map[string]*string) *OptionObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *OptionObjectHeaders) SetAccessControlRequestHeaders(v string) *OptionObjectHeaders {
	s.AccessControlRequestHeaders = &v
	return s
}

func (s *OptionObjectHeaders) SetAccessControlRequestMethod(v string) *OptionObjectHeaders {
	s.AccessControlRequestMethod = &v
	return s
}

func (s *OptionObjectHeaders) SetOrigin(v string) *OptionObjectHeaders {
	s.Origin = &v
	return s
}

type OptionObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s OptionObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s OptionObjectResponse) GoString() string {
	return s.String()
}

func (s *OptionObjectResponse) SetHeaders(v map[string]*string) *OptionObjectResponse {
	s.Headers = v
	return s
}

func (s *OptionObjectResponse) SetStatusCode(v int32) *OptionObjectResponse {
	s.StatusCode = &v
	return s
}

type PostObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PostObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PostObjectResponse) GoString() string {
	return s.String()
}

func (s *PostObjectResponse) SetHeaders(v map[string]*string) *PostObjectResponse {
	s.Headers = v
	return s
}

func (s *PostObjectResponse) SetStatusCode(v int32) *PostObjectResponse {
	s.StatusCode = &v
	return s
}

type PostVodPlaylistRequest struct {
	EndTime   *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s PostVodPlaylistRequest) String() string {
	return tea.Prettify(s)
}

func (s PostVodPlaylistRequest) GoString() string {
	return s.String()
}

func (s *PostVodPlaylistRequest) SetEndTime(v string) *PostVodPlaylistRequest {
	s.EndTime = &v
	return s
}

func (s *PostVodPlaylistRequest) SetStartTime(v string) *PostVodPlaylistRequest {
	s.StartTime = &v
	return s
}

type PostVodPlaylistResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PostVodPlaylistResponse) String() string {
	return tea.Prettify(s)
}

func (s PostVodPlaylistResponse) GoString() string {
	return s.String()
}

func (s *PostVodPlaylistResponse) SetHeaders(v map[string]*string) *PostVodPlaylistResponse {
	s.Headers = v
	return s
}

func (s *PostVodPlaylistResponse) SetStatusCode(v int32) *PostVodPlaylistResponse {
	s.StatusCode = &v
	return s
}

type PutBucketHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Acl           *string            `json:"x-oss-acl,omitempty" xml:"x-oss-acl,omitempty"`
}

func (s PutBucketHeaders) String() string {
	return tea.Prettify(s)
}

func (s PutBucketHeaders) GoString() string {
	return s.String()
}

func (s *PutBucketHeaders) SetCommonHeaders(v map[string]*string) *PutBucketHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutBucketHeaders) SetAcl(v string) *PutBucketHeaders {
	s.Acl = &v
	return s
}

type PutBucketRequest struct {
	CreateBucketConfiguration *CreateBucketConfiguration `json:"CreateBucketConfiguration,omitempty" xml:"CreateBucketConfiguration,omitempty"`
}

func (s PutBucketRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRequest) SetCreateBucketConfiguration(v *CreateBucketConfiguration) *PutBucketRequest {
	s.CreateBucketConfiguration = v
	return s
}

type PutBucketResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketResponse) GoString() string {
	return s.String()
}

func (s *PutBucketResponse) SetHeaders(v map[string]*string) *PutBucketResponse {
	s.Headers = v
	return s
}

func (s *PutBucketResponse) SetStatusCode(v int32) *PutBucketResponse {
	s.StatusCode = &v
	return s
}

type PutBucketAclHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Acl           *string            `json:"x-oss-acl,omitempty" xml:"x-oss-acl,omitempty"`
}

func (s PutBucketAclHeaders) String() string {
	return tea.Prettify(s)
}

func (s PutBucketAclHeaders) GoString() string {
	return s.String()
}

func (s *PutBucketAclHeaders) SetCommonHeaders(v map[string]*string) *PutBucketAclHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutBucketAclHeaders) SetAcl(v string) *PutBucketAclHeaders {
	s.Acl = &v
	return s
}

type PutBucketAclResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketAclResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketAclResponse) GoString() string {
	return s.String()
}

func (s *PutBucketAclResponse) SetHeaders(v map[string]*string) *PutBucketAclResponse {
	s.Headers = v
	return s
}

func (s *PutBucketAclResponse) SetStatusCode(v int32) *PutBucketAclResponse {
	s.StatusCode = &v
	return s
}

type PutBucketCorsRequest struct {
	CORSConfiguration *CORSConfiguration `json:"CORSConfiguration,omitempty" xml:"CORSConfiguration,omitempty"`
}

func (s PutBucketCorsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketCorsRequest) GoString() string {
	return s.String()
}

func (s *PutBucketCorsRequest) SetCORSConfiguration(v *CORSConfiguration) *PutBucketCorsRequest {
	s.CORSConfiguration = v
	return s
}

type PutBucketCorsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketCorsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketCorsResponse) GoString() string {
	return s.String()
}

func (s *PutBucketCorsResponse) SetHeaders(v map[string]*string) *PutBucketCorsResponse {
	s.Headers = v
	return s
}

func (s *PutBucketCorsResponse) SetStatusCode(v int32) *PutBucketCorsResponse {
	s.StatusCode = &v
	return s
}

type PutBucketEncryptionRequest struct {
	ServerSideEncryptionRule *ServerSideEncryptionRule `json:"ServerSideEncryptionRule,omitempty" xml:"ServerSideEncryptionRule,omitempty"`
}

func (s PutBucketEncryptionRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketEncryptionRequest) GoString() string {
	return s.String()
}

func (s *PutBucketEncryptionRequest) SetServerSideEncryptionRule(v *ServerSideEncryptionRule) *PutBucketEncryptionRequest {
	s.ServerSideEncryptionRule = v
	return s
}

type PutBucketEncryptionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketEncryptionResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketEncryptionResponse) GoString() string {
	return s.String()
}

func (s *PutBucketEncryptionResponse) SetHeaders(v map[string]*string) *PutBucketEncryptionResponse {
	s.Headers = v
	return s
}

func (s *PutBucketEncryptionResponse) SetStatusCode(v int32) *PutBucketEncryptionResponse {
	s.StatusCode = &v
	return s
}

type PutBucketInventoryRequest struct {
	InventoryConfiguration *InventoryConfiguration `json:"InventoryConfiguration,omitempty" xml:"InventoryConfiguration,omitempty"`
	InventoryId            *string                 `json:"inventoryId,omitempty" xml:"inventoryId,omitempty"`
}

func (s PutBucketInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketInventoryRequest) GoString() string {
	return s.String()
}

func (s *PutBucketInventoryRequest) SetInventoryConfiguration(v *InventoryConfiguration) *PutBucketInventoryRequest {
	s.InventoryConfiguration = v
	return s
}

func (s *PutBucketInventoryRequest) SetInventoryId(v string) *PutBucketInventoryRequest {
	s.InventoryId = &v
	return s
}

type PutBucketInventoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketInventoryResponse) GoString() string {
	return s.String()
}

func (s *PutBucketInventoryResponse) SetHeaders(v map[string]*string) *PutBucketInventoryResponse {
	s.Headers = v
	return s
}

func (s *PutBucketInventoryResponse) SetStatusCode(v int32) *PutBucketInventoryResponse {
	s.StatusCode = &v
	return s
}

type PutBucketLifecycleRequest struct {
	LifecycleConfiguration *LifecycleConfiguration `json:"LifecycleConfiguration,omitempty" xml:"LifecycleConfiguration,omitempty"`
}

func (s PutBucketLifecycleRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketLifecycleRequest) GoString() string {
	return s.String()
}

func (s *PutBucketLifecycleRequest) SetLifecycleConfiguration(v *LifecycleConfiguration) *PutBucketLifecycleRequest {
	s.LifecycleConfiguration = v
	return s
}

type PutBucketLifecycleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketLifecycleResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketLifecycleResponse) GoString() string {
	return s.String()
}

func (s *PutBucketLifecycleResponse) SetHeaders(v map[string]*string) *PutBucketLifecycleResponse {
	s.Headers = v
	return s
}

func (s *PutBucketLifecycleResponse) SetStatusCode(v int32) *PutBucketLifecycleResponse {
	s.StatusCode = &v
	return s
}

type PutBucketLoggingRequest struct {
	BucketLoggingStatus *BucketLoggingStatus `json:"BucketLoggingStatus,omitempty" xml:"BucketLoggingStatus,omitempty"`
}

func (s PutBucketLoggingRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketLoggingRequest) GoString() string {
	return s.String()
}

func (s *PutBucketLoggingRequest) SetBucketLoggingStatus(v *BucketLoggingStatus) *PutBucketLoggingRequest {
	s.BucketLoggingStatus = v
	return s
}

type PutBucketLoggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketLoggingResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketLoggingResponse) GoString() string {
	return s.String()
}

func (s *PutBucketLoggingResponse) SetHeaders(v map[string]*string) *PutBucketLoggingResponse {
	s.Headers = v
	return s
}

func (s *PutBucketLoggingResponse) SetStatusCode(v int32) *PutBucketLoggingResponse {
	s.StatusCode = &v
	return s
}

type PutBucketPolicyRequest struct {
	Policy *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutBucketPolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketPolicyRequest) GoString() string {
	return s.String()
}

func (s *PutBucketPolicyRequest) SetPolicy(v string) *PutBucketPolicyRequest {
	s.Policy = &v
	return s
}

type PutBucketPolicyResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketPolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketPolicyResponse) GoString() string {
	return s.String()
}

func (s *PutBucketPolicyResponse) SetHeaders(v map[string]*string) *PutBucketPolicyResponse {
	s.Headers = v
	return s
}

func (s *PutBucketPolicyResponse) SetStatusCode(v int32) *PutBucketPolicyResponse {
	s.StatusCode = &v
	return s
}

type PutBucketRefererRequest struct {
	RefererConfiguration *RefererConfiguration `json:"RefererConfiguration,omitempty" xml:"RefererConfiguration,omitempty"`
}

func (s PutBucketRefererRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketRefererRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRefererRequest) SetRefererConfiguration(v *RefererConfiguration) *PutBucketRefererRequest {
	s.RefererConfiguration = v
	return s
}

type PutBucketRefererResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketRefererResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketRefererResponse) GoString() string {
	return s.String()
}

func (s *PutBucketRefererResponse) SetHeaders(v map[string]*string) *PutBucketRefererResponse {
	s.Headers = v
	return s
}

func (s *PutBucketRefererResponse) SetStatusCode(v int32) *PutBucketRefererResponse {
	s.StatusCode = &v
	return s
}

type PutBucketReplicationRequest struct {
	ReplicationConfiguration *ReplicationConfiguration `json:"ReplicationConfiguration,omitempty" xml:"ReplicationConfiguration,omitempty"`
}

func (s PutBucketReplicationRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketReplicationRequest) GoString() string {
	return s.String()
}

func (s *PutBucketReplicationRequest) SetReplicationConfiguration(v *ReplicationConfiguration) *PutBucketReplicationRequest {
	s.ReplicationConfiguration = v
	return s
}

type PutBucketReplicationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketReplicationResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketReplicationResponse) GoString() string {
	return s.String()
}

func (s *PutBucketReplicationResponse) SetHeaders(v map[string]*string) *PutBucketReplicationResponse {
	s.Headers = v
	return s
}

func (s *PutBucketReplicationResponse) SetStatusCode(v int32) *PutBucketReplicationResponse {
	s.StatusCode = &v
	return s
}

type PutBucketRequestPaymentRequest struct {
	RequestPaymentConfiguration *RequestPaymentConfiguration `json:"RequestPaymentConfiguration,omitempty" xml:"RequestPaymentConfiguration,omitempty"`
}

func (s PutBucketRequestPaymentRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketRequestPaymentRequest) GoString() string {
	return s.String()
}

func (s *PutBucketRequestPaymentRequest) SetRequestPaymentConfiguration(v *RequestPaymentConfiguration) *PutBucketRequestPaymentRequest {
	s.RequestPaymentConfiguration = v
	return s
}

type PutBucketRequestPaymentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketRequestPaymentResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketRequestPaymentResponse) GoString() string {
	return s.String()
}

func (s *PutBucketRequestPaymentResponse) SetHeaders(v map[string]*string) *PutBucketRequestPaymentResponse {
	s.Headers = v
	return s
}

func (s *PutBucketRequestPaymentResponse) SetStatusCode(v int32) *PutBucketRequestPaymentResponse {
	s.StatusCode = &v
	return s
}

type PutBucketTagsRequest struct {
	Tagging *Tagging `json:"Tagging,omitempty" xml:"Tagging,omitempty"`
}

func (s PutBucketTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketTagsRequest) GoString() string {
	return s.String()
}

func (s *PutBucketTagsRequest) SetTagging(v *Tagging) *PutBucketTagsRequest {
	s.Tagging = v
	return s
}

type PutBucketTagsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketTagsResponse) GoString() string {
	return s.String()
}

func (s *PutBucketTagsResponse) SetHeaders(v map[string]*string) *PutBucketTagsResponse {
	s.Headers = v
	return s
}

func (s *PutBucketTagsResponse) SetStatusCode(v int32) *PutBucketTagsResponse {
	s.StatusCode = &v
	return s
}

type PutBucketTransferAccelerationRequest struct {
	TransferAccelerationConfiguration *TransferAccelerationConfiguration `json:"TransferAccelerationConfiguration,omitempty" xml:"TransferAccelerationConfiguration,omitempty"`
}

func (s PutBucketTransferAccelerationRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketTransferAccelerationRequest) GoString() string {
	return s.String()
}

func (s *PutBucketTransferAccelerationRequest) SetTransferAccelerationConfiguration(v *TransferAccelerationConfiguration) *PutBucketTransferAccelerationRequest {
	s.TransferAccelerationConfiguration = v
	return s
}

type PutBucketTransferAccelerationResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketTransferAccelerationResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketTransferAccelerationResponse) GoString() string {
	return s.String()
}

func (s *PutBucketTransferAccelerationResponse) SetHeaders(v map[string]*string) *PutBucketTransferAccelerationResponse {
	s.Headers = v
	return s
}

func (s *PutBucketTransferAccelerationResponse) SetStatusCode(v int32) *PutBucketTransferAccelerationResponse {
	s.StatusCode = &v
	return s
}

type PutBucketVersioningRequest struct {
	VersioningConfiguration *VersioningConfiguration `json:"VersioningConfiguration,omitempty" xml:"VersioningConfiguration,omitempty"`
}

func (s PutBucketVersioningRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketVersioningRequest) GoString() string {
	return s.String()
}

func (s *PutBucketVersioningRequest) SetVersioningConfiguration(v *VersioningConfiguration) *PutBucketVersioningRequest {
	s.VersioningConfiguration = v
	return s
}

type PutBucketVersioningResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketVersioningResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketVersioningResponse) GoString() string {
	return s.String()
}

func (s *PutBucketVersioningResponse) SetHeaders(v map[string]*string) *PutBucketVersioningResponse {
	s.Headers = v
	return s
}

func (s *PutBucketVersioningResponse) SetStatusCode(v int32) *PutBucketVersioningResponse {
	s.StatusCode = &v
	return s
}

type PutBucketWebsiteRequest struct {
	WebsiteConfiguration *WebsiteConfiguration `json:"WebsiteConfiguration,omitempty" xml:"WebsiteConfiguration,omitempty"`
}

func (s PutBucketWebsiteRequest) String() string {
	return tea.Prettify(s)
}

func (s PutBucketWebsiteRequest) GoString() string {
	return s.String()
}

func (s *PutBucketWebsiteRequest) SetWebsiteConfiguration(v *WebsiteConfiguration) *PutBucketWebsiteRequest {
	s.WebsiteConfiguration = v
	return s
}

type PutBucketWebsiteResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutBucketWebsiteResponse) String() string {
	return tea.Prettify(s)
}

func (s PutBucketWebsiteResponse) GoString() string {
	return s.String()
}

func (s *PutBucketWebsiteResponse) SetHeaders(v map[string]*string) *PutBucketWebsiteResponse {
	s.Headers = v
	return s
}

func (s *PutBucketWebsiteResponse) SetStatusCode(v int32) *PutBucketWebsiteResponse {
	s.StatusCode = &v
	return s
}

type PutLiveChannelRequest struct {
	LiveChannelConfiguration *LiveChannelConfiguration `json:"LiveChannelConfiguration,omitempty" xml:"LiveChannelConfiguration,omitempty"`
}

func (s PutLiveChannelRequest) String() string {
	return tea.Prettify(s)
}

func (s PutLiveChannelRequest) GoString() string {
	return s.String()
}

func (s *PutLiveChannelRequest) SetLiveChannelConfiguration(v *LiveChannelConfiguration) *PutLiveChannelRequest {
	s.LiveChannelConfiguration = v
	return s
}

type PutLiveChannelResponseBody struct {
	PlayUrls    *LiveChannelPlayUrls    `json:"PlayUrls,omitempty" xml:"PlayUrls,omitempty"`
	PublishUrls *LiveChannelPublishUrls `json:"PublishUrls,omitempty" xml:"PublishUrls,omitempty"`
}

func (s PutLiveChannelResponseBody) String() string {
	return tea.Prettify(s)
}

func (s PutLiveChannelResponseBody) GoString() string {
	return s.String()
}

func (s *PutLiveChannelResponseBody) SetPlayUrls(v *LiveChannelPlayUrls) *PutLiveChannelResponseBody {
	s.PlayUrls = v
	return s
}

func (s *PutLiveChannelResponseBody) SetPublishUrls(v *LiveChannelPublishUrls) *PutLiveChannelResponseBody {
	s.PublishUrls = v
	return s
}

type PutLiveChannelResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *PutLiveChannelResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s PutLiveChannelResponse) String() string {
	return tea.Prettify(s)
}

func (s PutLiveChannelResponse) GoString() string {
	return s.String()
}

func (s *PutLiveChannelResponse) SetHeaders(v map[string]*string) *PutLiveChannelResponse {
	s.Headers = v
	return s
}

func (s *PutLiveChannelResponse) SetStatusCode(v int32) *PutLiveChannelResponse {
	s.StatusCode = &v
	return s
}

func (s *PutLiveChannelResponse) SetBody(v *PutLiveChannelResponseBody) *PutLiveChannelResponse {
	s.Body = v
	return s
}

type PutLiveChannelStatusRequest struct {
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PutLiveChannelStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s PutLiveChannelStatusRequest) GoString() string {
	return s.String()
}

func (s *PutLiveChannelStatusRequest) SetStatus(v string) *PutLiveChannelStatusRequest {
	s.Status = &v
	return s
}

type PutLiveChannelStatusResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutLiveChannelStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s PutLiveChannelStatusResponse) GoString() string {
	return s.String()
}

func (s *PutLiveChannelStatusResponse) SetHeaders(v map[string]*string) *PutLiveChannelStatusResponse {
	s.Headers = v
	return s
}

func (s *PutLiveChannelStatusResponse) SetStatusCode(v int32) *PutLiveChannelStatusResponse {
	s.StatusCode = &v
	return s
}

type PutObjectHeaders struct {
	CommonHeaders        map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ForbidOverwrite      *bool              `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
	MetaData             map[string]*string `json:"x-oss-meta-*,omitempty" xml:"x-oss-meta-*,omitempty"`
	Acl                  *string            `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
	SseDataEncryption    *string            `json:"x-oss-server-side-data-encryption,omitempty" xml:"x-oss-server-side-data-encryption,omitempty"`
	ServerSideEncryption *string            `json:"x-oss-server-side-encryption,omitempty" xml:"x-oss-server-side-encryption,omitempty"`
	SseKeyId             *string            `json:"x-oss-server-side-encryption-key-id,omitempty" xml:"x-oss-server-side-encryption-key-id,omitempty"`
	StorageClass         *string            `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
	Tagging              *string            `json:"x-oss-tagging,omitempty" xml:"x-oss-tagging,omitempty"`
}

func (s PutObjectHeaders) String() string {
	return tea.Prettify(s)
}

func (s PutObjectHeaders) GoString() string {
	return s.String()
}

func (s *PutObjectHeaders) SetCommonHeaders(v map[string]*string) *PutObjectHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutObjectHeaders) SetForbidOverwrite(v bool) *PutObjectHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *PutObjectHeaders) SetMetaData(v map[string]*string) *PutObjectHeaders {
	s.MetaData = v
	return s
}

func (s *PutObjectHeaders) SetAcl(v string) *PutObjectHeaders {
	s.Acl = &v
	return s
}

func (s *PutObjectHeaders) SetSseDataEncryption(v string) *PutObjectHeaders {
	s.SseDataEncryption = &v
	return s
}

func (s *PutObjectHeaders) SetServerSideEncryption(v string) *PutObjectHeaders {
	s.ServerSideEncryption = &v
	return s
}

func (s *PutObjectHeaders) SetSseKeyId(v string) *PutObjectHeaders {
	s.SseKeyId = &v
	return s
}

func (s *PutObjectHeaders) SetStorageClass(v string) *PutObjectHeaders {
	s.StorageClass = &v
	return s
}

func (s *PutObjectHeaders) SetTagging(v string) *PutObjectHeaders {
	s.Tagging = &v
	return s
}

type PutObjectRequest struct {
	Body io.Reader `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s PutObjectRequest) GoString() string {
	return s.String()
}

func (s *PutObjectRequest) SetBody(v io.Reader) *PutObjectRequest {
	s.Body = v
	return s
}

type PutObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s PutObjectResponse) GoString() string {
	return s.String()
}

func (s *PutObjectResponse) SetHeaders(v map[string]*string) *PutObjectResponse {
	s.Headers = v
	return s
}

func (s *PutObjectResponse) SetStatusCode(v int32) *PutObjectResponse {
	s.StatusCode = &v
	return s
}

type PutObjectAclHeaders struct {
	CommonHeaders map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	Acl           *string            `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
}

func (s PutObjectAclHeaders) String() string {
	return tea.Prettify(s)
}

func (s PutObjectAclHeaders) GoString() string {
	return s.String()
}

func (s *PutObjectAclHeaders) SetCommonHeaders(v map[string]*string) *PutObjectAclHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutObjectAclHeaders) SetAcl(v string) *PutObjectAclHeaders {
	s.Acl = &v
	return s
}

type PutObjectAclRequest struct {
	VersionId *string `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s PutObjectAclRequest) String() string {
	return tea.Prettify(s)
}

func (s PutObjectAclRequest) GoString() string {
	return s.String()
}

func (s *PutObjectAclRequest) SetVersionId(v string) *PutObjectAclRequest {
	s.VersionId = &v
	return s
}

type PutObjectAclResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutObjectAclResponse) String() string {
	return tea.Prettify(s)
}

func (s PutObjectAclResponse) GoString() string {
	return s.String()
}

func (s *PutObjectAclResponse) SetHeaders(v map[string]*string) *PutObjectAclResponse {
	s.Headers = v
	return s
}

func (s *PutObjectAclResponse) SetStatusCode(v int32) *PutObjectAclResponse {
	s.StatusCode = &v
	return s
}

type PutObjectTaggingRequest struct {
	Tagging   *Tagging `json:"Tagging,omitempty" xml:"Tagging,omitempty"`
	VersionId *string  `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s PutObjectTaggingRequest) String() string {
	return tea.Prettify(s)
}

func (s PutObjectTaggingRequest) GoString() string {
	return s.String()
}

func (s *PutObjectTaggingRequest) SetTagging(v *Tagging) *PutObjectTaggingRequest {
	s.Tagging = v
	return s
}

func (s *PutObjectTaggingRequest) SetVersionId(v string) *PutObjectTaggingRequest {
	s.VersionId = &v
	return s
}

type PutObjectTaggingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutObjectTaggingResponse) String() string {
	return tea.Prettify(s)
}

func (s PutObjectTaggingResponse) GoString() string {
	return s.String()
}

func (s *PutObjectTaggingResponse) SetHeaders(v map[string]*string) *PutObjectTaggingResponse {
	s.Headers = v
	return s
}

func (s *PutObjectTaggingResponse) SetStatusCode(v int32) *PutObjectTaggingResponse {
	s.StatusCode = &v
	return s
}

type PutSymlinkHeaders struct {
	CommonHeaders    map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	ForbidOverwrite  *string            `json:"x-oss-forbid-overwrite,omitempty" xml:"x-oss-forbid-overwrite,omitempty"`
	Acl              *string            `json:"x-oss-object-acl,omitempty" xml:"x-oss-object-acl,omitempty"`
	StorageClass     *string            `json:"x-oss-storage-class,omitempty" xml:"x-oss-storage-class,omitempty"`
	SymlinkTargetKey *string            `json:"x-oss-symlink-target,omitempty" xml:"x-oss-symlink-target,omitempty"`
}

func (s PutSymlinkHeaders) String() string {
	return tea.Prettify(s)
}

func (s PutSymlinkHeaders) GoString() string {
	return s.String()
}

func (s *PutSymlinkHeaders) SetCommonHeaders(v map[string]*string) *PutSymlinkHeaders {
	s.CommonHeaders = v
	return s
}

func (s *PutSymlinkHeaders) SetForbidOverwrite(v string) *PutSymlinkHeaders {
	s.ForbidOverwrite = &v
	return s
}

func (s *PutSymlinkHeaders) SetAcl(v string) *PutSymlinkHeaders {
	s.Acl = &v
	return s
}

func (s *PutSymlinkHeaders) SetStorageClass(v string) *PutSymlinkHeaders {
	s.StorageClass = &v
	return s
}

func (s *PutSymlinkHeaders) SetSymlinkTargetKey(v string) *PutSymlinkHeaders {
	s.SymlinkTargetKey = &v
	return s
}

type PutSymlinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s PutSymlinkResponse) String() string {
	return tea.Prettify(s)
}

func (s PutSymlinkResponse) GoString() string {
	return s.String()
}

func (s *PutSymlinkResponse) SetHeaders(v map[string]*string) *PutSymlinkResponse {
	s.Headers = v
	return s
}

func (s *PutSymlinkResponse) SetStatusCode(v int32) *PutSymlinkResponse {
	s.StatusCode = &v
	return s
}

type RestoreObjectRequest struct {
	RestoreRequest *RestoreRequest `json:"RestoreRequest,omitempty" xml:"RestoreRequest,omitempty"`
	VersionId      *string         `json:"versionId,omitempty" xml:"versionId,omitempty"`
}

func (s RestoreObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreObjectRequest) GoString() string {
	return s.String()
}

func (s *RestoreObjectRequest) SetRestoreRequest(v *RestoreRequest) *RestoreObjectRequest {
	s.RestoreRequest = v
	return s
}

func (s *RestoreObjectRequest) SetVersionId(v string) *RestoreObjectRequest {
	s.VersionId = &v
	return s
}

type RestoreObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s RestoreObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreObjectResponse) GoString() string {
	return s.String()
}

func (s *RestoreObjectResponse) SetHeaders(v map[string]*string) *RestoreObjectResponse {
	s.Headers = v
	return s
}

func (s *RestoreObjectResponse) SetStatusCode(v int32) *RestoreObjectResponse {
	s.StatusCode = &v
	return s
}

type SelectObjectRequest struct {
	SelectRequest *SelectRequest `json:"SelectRequest,omitempty" xml:"SelectRequest,omitempty"`
}

func (s SelectObjectRequest) String() string {
	return tea.Prettify(s)
}

func (s SelectObjectRequest) GoString() string {
	return s.String()
}

func (s *SelectObjectRequest) SetSelectRequest(v *SelectRequest) *SelectObjectRequest {
	s.SelectRequest = v
	return s
}

type SelectObjectResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       io.Reader          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SelectObjectResponse) String() string {
	return tea.Prettify(s)
}

func (s SelectObjectResponse) GoString() string {
	return s.String()
}

func (s *SelectObjectResponse) SetHeaders(v map[string]*string) *SelectObjectResponse {
	s.Headers = v
	return s
}

func (s *SelectObjectResponse) SetStatusCode(v int32) *SelectObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *SelectObjectResponse) SetBody(v io.Reader) *SelectObjectResponse {
	s.Body = v
	return s
}

type UploadPartRequest struct {
	Body       io.Reader `json:"body,omitempty" xml:"body,omitempty"`
	PartNumber *int64    `json:"partNumber,omitempty" xml:"partNumber,omitempty"`
	UploadId   *string   `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s UploadPartRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadPartRequest) GoString() string {
	return s.String()
}

func (s *UploadPartRequest) SetBody(v io.Reader) *UploadPartRequest {
	s.Body = v
	return s
}

func (s *UploadPartRequest) SetPartNumber(v int64) *UploadPartRequest {
	s.PartNumber = &v
	return s
}

func (s *UploadPartRequest) SetUploadId(v string) *UploadPartRequest {
	s.UploadId = &v
	return s
}

type UploadPartResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
}

func (s UploadPartResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadPartResponse) GoString() string {
	return s.String()
}

func (s *UploadPartResponse) SetHeaders(v map[string]*string) *UploadPartResponse {
	s.Headers = v
	return s
}

func (s *UploadPartResponse) SetStatusCode(v int32) *UploadPartResponse {
	s.StatusCode = &v
	return s
}

type UploadPartCopyHeaders struct {
	CommonHeaders               map[string]*string `json:"commonHeaders,omitempty" xml:"commonHeaders,omitempty"`
	CopySource                  *string            `json:"x-oss-copy-source,omitempty" xml:"x-oss-copy-source,omitempty"`
	CopySourceIfMatch           *string            `json:"x-oss-copy-source-if-match,omitempty" xml:"x-oss-copy-source-if-match,omitempty"`
	CopySourceIfModifiedSince   *string            `json:"x-oss-copy-source-if-modified-since,omitempty" xml:"x-oss-copy-source-if-modified-since,omitempty"`
	CopySourceIfNoneMatch       *string            `json:"x-oss-copy-source-if-none-match,omitempty" xml:"x-oss-copy-source-if-none-match,omitempty"`
	CopySourceIfUnmodifiedSince *string            `json:"x-oss-copy-source-if-unmodified-since,omitempty" xml:"x-oss-copy-source-if-unmodified-since,omitempty"`
	CopySourceRange             *string            `json:"x-oss-copy-source-range,omitempty" xml:"x-oss-copy-source-range,omitempty"`
}

func (s UploadPartCopyHeaders) String() string {
	return tea.Prettify(s)
}

func (s UploadPartCopyHeaders) GoString() string {
	return s.String()
}

func (s *UploadPartCopyHeaders) SetCommonHeaders(v map[string]*string) *UploadPartCopyHeaders {
	s.CommonHeaders = v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySource(v string) *UploadPartCopyHeaders {
	s.CopySource = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceIfMatch(v string) *UploadPartCopyHeaders {
	s.CopySourceIfMatch = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceIfModifiedSince(v string) *UploadPartCopyHeaders {
	s.CopySourceIfModifiedSince = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceIfNoneMatch(v string) *UploadPartCopyHeaders {
	s.CopySourceIfNoneMatch = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceIfUnmodifiedSince(v string) *UploadPartCopyHeaders {
	s.CopySourceIfUnmodifiedSince = &v
	return s
}

func (s *UploadPartCopyHeaders) SetCopySourceRange(v string) *UploadPartCopyHeaders {
	s.CopySourceRange = &v
	return s
}

type UploadPartCopyRequest struct {
	PartNumber *int64  `json:"partNumber,omitempty" xml:"partNumber,omitempty"`
	UploadId   *string `json:"uploadId,omitempty" xml:"uploadId,omitempty"`
}

func (s UploadPartCopyRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadPartCopyRequest) GoString() string {
	return s.String()
}

func (s *UploadPartCopyRequest) SetPartNumber(v int64) *UploadPartCopyRequest {
	s.PartNumber = &v
	return s
}

func (s *UploadPartCopyRequest) SetUploadId(v string) *UploadPartCopyRequest {
	s.UploadId = &v
	return s
}

type UploadPartCopyResponseBody struct {
	ETag         *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	LastModified *string `json:"LastModified,omitempty" xml:"LastModified,omitempty"`
}

func (s UploadPartCopyResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UploadPartCopyResponseBody) GoString() string {
	return s.String()
}

func (s *UploadPartCopyResponseBody) SetETag(v string) *UploadPartCopyResponseBody {
	s.ETag = &v
	return s
}

func (s *UploadPartCopyResponseBody) SetLastModified(v string) *UploadPartCopyResponseBody {
	s.LastModified = &v
	return s
}

type UploadPartCopyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UploadPartCopyResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UploadPartCopyResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadPartCopyResponse) GoString() string {
	return s.String()
}

func (s *UploadPartCopyResponse) SetHeaders(v map[string]*string) *UploadPartCopyResponse {
	s.Headers = v
	return s
}

func (s *UploadPartCopyResponse) SetStatusCode(v int32) *UploadPartCopyResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadPartCopyResponse) SetBody(v *UploadPartCopyResponseBody) *UploadPartCopyResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.EndpointRule = tea.String("")
	return nil
}

func (client *Client) AbortBucketWorm(bucket *string) (_result *AbortBucketWormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AbortBucketWormResponse{}
	_body, _err := client.AbortBucketWormWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortBucketWormWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AbortBucketWormResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("AbortBucketWorm"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?worm"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &AbortBucketWormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AbortMultipartUpload(bucket *string, key *string, request *AbortMultipartUploadRequest) (_result *AbortMultipartUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AbortMultipartUploadResponse{}
	_body, _err := client.AbortMultipartUploadWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AbortMultipartUploadWithOptions(bucket *string, key *string, request *AbortMultipartUploadRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AbortMultipartUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UploadId)) {
		query["uploadId"] = request.UploadId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("AbortMultipartUpload"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &AbortMultipartUploadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AppendObject(bucket *string, key *string, request *AppendObjectRequest) (_result *AppendObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &AppendObjectHeaders{}
	_result = &AppendObjectResponse{}
	_body, _err := client.AppendObjectWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AppendObjectWithOptions(bucket *string, key *string, request *AppendObjectRequest, headers *AppendObjectHeaders, runtime *util.RuntimeOptions) (_result *AppendObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Position)) {
		query["position"] = request.Position
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.CacheControl)) {
		realHeaders["Cache-Control"] = util.ToJSONString(headers.CacheControl)
	}

	if !tea.BoolValue(util.IsUnset(headers.ContentDisposition)) {
		realHeaders["Content-Disposition"] = util.ToJSONString(headers.ContentDisposition)
	}

	if !tea.BoolValue(util.IsUnset(headers.ContentEncoding)) {
		realHeaders["Content-Encoding"] = util.ToJSONString(headers.ContentEncoding)
	}

	if !tea.BoolValue(util.IsUnset(headers.ContentMD5)) {
		realHeaders["Content-MD5"] = util.ToJSONString(headers.ContentMD5)
	}

	if !tea.BoolValue(util.IsUnset(headers.Expires)) {
		realHeaders["Expires"] = util.ToJSONString(headers.Expires)
	}

	if !tea.BoolValue(util.IsUnset(headers.MetaData)) {
		realHeaders["x-oss-meta-*"] = util.ToJSONString(headers.MetaData)
	}

	if !tea.BoolValue(util.IsUnset(headers.Acl)) {
		realHeaders["x-oss-object-acl"] = util.ToJSONString(headers.Acl)
	}

	if !tea.BoolValue(util.IsUnset(headers.ServerSideEncryption)) {
		realHeaders["x-oss-server-side-encryption"] = util.ToJSONString(headers.ServerSideEncryption)
	}

	if !tea.BoolValue(util.IsUnset(headers.StorageClass)) {
		realHeaders["x-oss-storage-class"] = util.ToJSONString(headers.StorageClass)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
		Stream:  request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("AppendObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?append"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("binary"),
		BodyType:    tea.String("xml"),
	}
	_result = &AppendObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteBucketWorm(bucket *string, request *CompleteBucketWormRequest) (_result *CompleteBucketWormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CompleteBucketWormResponse{}
	_body, _err := client.CompleteBucketWormWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteBucketWormWithOptions(bucket *string, request *CompleteBucketWormRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CompleteBucketWormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WormId)) {
		query["wormId"] = request.WormId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("CompleteBucketWorm"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &CompleteBucketWormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CompleteMultipartUpload(bucket *string, key *string, request *CompleteMultipartUploadRequest) (_result *CompleteMultipartUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CompleteMultipartUploadHeaders{}
	_result = &CompleteMultipartUploadResponse{}
	_body, _err := client.CompleteMultipartUploadWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CompleteMultipartUploadWithOptions(bucket *string, key *string, request *CompleteMultipartUploadRequest, headers *CompleteMultipartUploadHeaders, runtime *util.RuntimeOptions) (_result *CompleteMultipartUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodingType)) {
		query["encoding-type"] = request.EncodingType
	}

	if !tea.BoolValue(util.IsUnset(request.UploadId)) {
		query["uploadId"] = request.UploadId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CompleteMultipartUpload))) {
		body["completeMultipartUpload"] = request.CompleteMultipartUpload
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.CompleteAll)) {
		realHeaders["x-oss-complete-all"] = util.ToJSONString(headers.CompleteAll)
	}

	if !tea.BoolValue(util.IsUnset(headers.ForbidOverwrite)) {
		realHeaders["x-oss-forbid-overwrite"] = util.ToJSONString(headers.ForbidOverwrite)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompleteMultipartUpload"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &CompleteMultipartUploadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CopyObject(bucket *string, key *string) (_result *CopyObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &CopyObjectHeaders{}
	_result = &CopyObjectResponse{}
	_body, _err := client.CopyObjectWithOptions(bucket, key, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CopyObjectWithOptions(bucket *string, key *string, headers *CopyObjectHeaders, runtime *util.RuntimeOptions) (_result *CopyObjectResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySource)) {
		realHeaders["x-oss-copy-source"] = util.ToJSONString(headers.CopySource)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceIfMatch)) {
		realHeaders["x-oss-copy-source-if-match"] = util.ToJSONString(headers.CopySourceIfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceIfModifiedSince)) {
		realHeaders["x-oss-copy-source-if-modified-since"] = util.ToJSONString(headers.CopySourceIfModifiedSince)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceIfNoneMatch)) {
		realHeaders["x-oss-copy-source-if-none-match"] = util.ToJSONString(headers.CopySourceIfNoneMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceIfUnmodifiedSince)) {
		realHeaders["x-oss-copy-source-if-unmodified-since"] = util.ToJSONString(headers.CopySourceIfUnmodifiedSince)
	}

	if !tea.BoolValue(util.IsUnset(headers.ForbidOverwrite)) {
		realHeaders["x-oss-forbid-overwrite"] = util.ToJSONString(headers.ForbidOverwrite)
	}

	if !tea.BoolValue(util.IsUnset(headers.MetaData)) {
		realHeaders["x-oss-meta-*"] = util.ToJSONString(headers.MetaData)
	}

	if !tea.BoolValue(util.IsUnset(headers.MetadataDirective)) {
		realHeaders["x-oss-metadata-directive"] = util.ToJSONString(headers.MetadataDirective)
	}

	if !tea.BoolValue(util.IsUnset(headers.Acl)) {
		realHeaders["x-oss-object-acl"] = util.ToJSONString(headers.Acl)
	}

	if !tea.BoolValue(util.IsUnset(headers.ServerSideEncryption)) {
		realHeaders["x-oss-server-side-encryption"] = util.ToJSONString(headers.ServerSideEncryption)
	}

	if !tea.BoolValue(util.IsUnset(headers.SseKeyId)) {
		realHeaders["x-oss-server-side-encryption-key-id"] = util.ToJSONString(headers.SseKeyId)
	}

	if !tea.BoolValue(util.IsUnset(headers.StorageClass)) {
		realHeaders["x-oss-storage-class"] = util.ToJSONString(headers.StorageClass)
	}

	if !tea.BoolValue(util.IsUnset(headers.Tagging)) {
		realHeaders["x-oss-tagging"] = util.ToJSONString(headers.Tagging)
	}

	if !tea.BoolValue(util.IsUnset(headers.TaggingDirective)) {
		realHeaders["x-oss-tagging-directive"] = util.ToJSONString(headers.TaggingDirective)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("CopyObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &CopyObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateSelectObjectMeta(bucket *string, key *string, request *CreateSelectObjectMetaRequest) (_result *CreateSelectObjectMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSelectObjectMetaResponse{}
	_body, _err := client.CreateSelectObjectMetaWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateSelectObjectMetaWithOptions(bucket *string, key *string, request *CreateSelectObjectMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSelectObjectMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.SelectMetaRequest)),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSelectObjectMeta"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSelectObjectMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucket(bucket *string) (_result *DeleteBucketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketResponse{}
	_body, _err := client.DeleteBucketWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucket"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketCors(bucket *string) (_result *DeleteBucketCorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketCorsResponse{}
	_body, _err := client.DeleteBucketCorsWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketCorsWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketCorsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketCors"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?cors"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketCorsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketEncryption(bucket *string) (_result *DeleteBucketEncryptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketEncryptionResponse{}
	_body, _err := client.DeleteBucketEncryptionWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketEncryptionWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketEncryptionResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketEncryption"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?encryption"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketEncryptionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketInventory(bucket *string, request *DeleteBucketInventoryRequest) (_result *DeleteBucketInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketInventoryResponse{}
	_body, _err := client.DeleteBucketInventoryWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketInventoryWithOptions(bucket *string, request *DeleteBucketInventoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InventoryId)) {
		query["inventoryId"] = request.InventoryId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketInventory"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?inventory"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketLifecycle(bucket *string) (_result *DeleteBucketLifecycleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketLifecycleResponse{}
	_body, _err := client.DeleteBucketLifecycleWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketLifecycleWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketLifecycleResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketLifecycle"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?lifecycle"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketLifecycleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketLogging(bucket *string) (_result *DeleteBucketLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketLoggingResponse{}
	_body, _err := client.DeleteBucketLoggingWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketLoggingWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketLoggingResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketLogging"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?logging"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketPolicy(bucket *string) (_result *DeleteBucketPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketPolicyResponse{}
	_body, _err := client.DeleteBucketPolicyWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketPolicyWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketPolicyResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketPolicy"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?policy"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketReplication(bucket *string, request *DeleteBucketReplicationRequest) (_result *DeleteBucketReplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketReplicationResponse{}
	_body, _err := client.DeleteBucketReplicationWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketReplicationWithOptions(bucket *string, request *DeleteBucketReplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketReplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Body)),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketReplication"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?replication&comp=delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketReplicationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketTags(bucket *string) (_result *DeleteBucketTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketTagsResponse{}
	_body, _err := client.DeleteBucketTagsWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketTagsWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketTagsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketTags"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?tagging"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketTagsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteBucketWebsite(bucket *string) (_result *DeleteBucketWebsiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteBucketWebsiteResponse{}
	_body, _err := client.DeleteBucketWebsiteWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteBucketWebsiteWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteBucketWebsiteResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteBucketWebsite"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?website"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteBucketWebsiteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteLiveChannel(bucket *string, channel *string) (_result *DeleteLiveChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteLiveChannelResponse{}
	_body, _err := client.DeleteLiveChannelWithOptions(bucket, channel, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteLiveChannelWithOptions(bucket *string, channel *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteLiveChannelResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteLiveChannel"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(channel) + "?live"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteLiveChannelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteMultipleObjects(bucket *string, request *DeleteMultipleObjectsRequest) (_result *DeleteMultipleObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteMultipleObjectsResponse{}
	_body, _err := client.DeleteMultipleObjectsWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteMultipleObjectsWithOptions(bucket *string, request *DeleteMultipleObjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteMultipleObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodingType)) {
		query["encoding-type"] = request.EncodingType
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Delete)),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteMultipleObjects"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteMultipleObjectsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteObject(bucket *string, key *string, request *DeleteObjectRequest) (_result *DeleteObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteObjectResponse{}
	_body, _err := client.DeleteObjectWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteObjectWithOptions(bucket *string, key *string, request *DeleteObjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("none"),
	}
	_result = &DeleteObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteObjectTagging(bucket *string, key *string, request *DeleteObjectTaggingRequest) (_result *DeleteObjectTaggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteObjectTaggingResponse{}
	_body, _err := client.DeleteObjectTaggingWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteObjectTaggingWithOptions(bucket *string, key *string, request *DeleteObjectTaggingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteObjectTaggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteObjectTagging"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?tagging"),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DeleteObjectTaggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Regions)) {
		query["regions"] = request.Regions
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?regions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ExtendBucketWorm(bucket *string, request *ExtendBucketWormRequest) (_result *ExtendBucketWormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ExtendBucketWormResponse{}
	_body, _err := client.ExtendBucketWormWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ExtendBucketWormWithOptions(bucket *string, request *ExtendBucketWormRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ExtendBucketWormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WormId)) {
		query["wormId"] = request.WormId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ExtendWormConfiguration))) {
		body["extendWormConfiguration"] = request.ExtendWormConfiguration
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtendBucketWorm"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?wormExtend"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ExtendBucketWormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketAcl(bucket *string) (_result *GetBucketAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketAclResponse{}
	_body, _err := client.GetBucketAclWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketAclWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketAclResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketAcl"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?acl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketAclResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketCors(bucket *string) (_result *GetBucketCorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketCorsResponse{}
	_body, _err := client.GetBucketCorsWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketCorsWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketCorsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketCors"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?cors"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketCorsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketEncryption(bucket *string) (_result *GetBucketEncryptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketEncryptionResponse{}
	_body, _err := client.GetBucketEncryptionWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketEncryptionWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketEncryptionResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketEncryption"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?encryption"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketEncryptionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketInfo(bucket *string) (_result *GetBucketInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketInfoResponse{}
	_body, _err := client.GetBucketInfoWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketInfoWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketInfoResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketInfo"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?bucketInfo"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketInventory(bucket *string, request *GetBucketInventoryRequest) (_result *GetBucketInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketInventoryResponse{}
	_body, _err := client.GetBucketInventoryWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketInventoryWithOptions(bucket *string, request *GetBucketInventoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InventoryId)) {
		query["inventoryId"] = request.InventoryId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketInventory"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?inventory"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketLifecycle(bucket *string) (_result *GetBucketLifecycleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketLifecycleResponse{}
	_body, _err := client.GetBucketLifecycleWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketLifecycleWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketLifecycleResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketLifecycle"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?lifecycle"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketLifecycleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketLocation(bucket *string) (_result *GetBucketLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketLocationResponse{}
	_body, _err := client.GetBucketLocationWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketLocationWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketLocationResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketLocation"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?location"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketLocationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketLogging(bucket *string) (_result *GetBucketLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketLoggingResponse{}
	_body, _err := client.GetBucketLoggingWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketLoggingWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketLoggingResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketLogging"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?logging"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketPolicy(bucket *string) (_result *GetBucketPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketPolicyResponse{}
	_body, _err := client.GetBucketPolicyWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketPolicyWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketPolicyResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketPolicy"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?policy"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("string"),
	}
	_result = &GetBucketPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketReferer(bucket *string) (_result *GetBucketRefererResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketRefererResponse{}
	_body, _err := client.GetBucketRefererWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketRefererWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketRefererResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketReferer"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?referer"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketRefererResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketReplication(bucket *string) (_result *GetBucketReplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketReplicationResponse{}
	_body, _err := client.GetBucketReplicationWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketReplicationWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketReplicationResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketReplication"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?replication"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketReplicationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketReplicationLocation(bucket *string) (_result *GetBucketReplicationLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketReplicationLocationResponse{}
	_body, _err := client.GetBucketReplicationLocationWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketReplicationLocationWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketReplicationLocationResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketReplicationLocation"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?replicationLocation"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketReplicationLocationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketReplicationProgress(bucket *string, request *GetBucketReplicationProgressRequest) (_result *GetBucketReplicationProgressResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketReplicationProgressResponse{}
	_body, _err := client.GetBucketReplicationProgressWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketReplicationProgressWithOptions(bucket *string, request *GetBucketReplicationProgressRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketReplicationProgressResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RuleId)) {
		query["rule-id"] = request.RuleId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketReplicationProgress"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?replicationProgress"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketReplicationProgressResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketRequestPayment(bucket *string) (_result *GetBucketRequestPaymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketRequestPaymentResponse{}
	_body, _err := client.GetBucketRequestPaymentWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketRequestPaymentWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketRequestPaymentResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketRequestPayment"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?requestPayment"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketRequestPaymentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketTags(bucket *string) (_result *GetBucketTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketTagsResponse{}
	_body, _err := client.GetBucketTagsWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketTagsWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketTagsResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketTags"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?tagging"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketTagsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketTransferAcceleration(bucket *string) (_result *GetBucketTransferAccelerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketTransferAccelerationResponse{}
	_body, _err := client.GetBucketTransferAccelerationWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketTransferAccelerationWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketTransferAccelerationResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketTransferAcceleration"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?transferAcceleration"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketTransferAccelerationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketVersioning(bucket *string) (_result *GetBucketVersioningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketVersioningResponse{}
	_body, _err := client.GetBucketVersioningWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketVersioningWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketVersioningResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketVersioning"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?versioning"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketVersioningResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketWebsite(bucket *string) (_result *GetBucketWebsiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketWebsiteResponse{}
	_body, _err := client.GetBucketWebsiteWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketWebsiteWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketWebsiteResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketWebsite"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?website"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketWebsiteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetBucketWorm(bucket *string) (_result *GetBucketWormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetBucketWormResponse{}
	_body, _err := client.GetBucketWormWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetBucketWormWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetBucketWormResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetBucketWorm"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?worm"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetBucketWormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveChannelHistory(bucket *string, channel *string) (_result *GetLiveChannelHistoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLiveChannelHistoryResponse{}
	_body, _err := client.GetLiveChannelHistoryWithOptions(bucket, channel, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveChannelHistoryWithOptions(bucket *string, channel *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLiveChannelHistoryResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveChannelHistory"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(channel) + "?live&comp=history"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetLiveChannelHistoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveChannelInfo(bucket *string, channel *string) (_result *GetLiveChannelInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLiveChannelInfoResponse{}
	_body, _err := client.GetLiveChannelInfoWithOptions(bucket, channel, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveChannelInfoWithOptions(bucket *string, channel *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLiveChannelInfoResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveChannelInfo"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(channel) + "?live"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetLiveChannelInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLiveChannelStat(bucket *string, channel *string) (_result *GetLiveChannelStatResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLiveChannelStatResponse{}
	_body, _err := client.GetLiveChannelStatWithOptions(bucket, channel, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLiveChannelStatWithOptions(bucket *string, channel *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLiveChannelStatResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetLiveChannelStat"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(channel) + "?live&comp=stat"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetLiveChannelStatResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetObject(bucket *string, key *string, request *GetObjectRequest) (_result *GetObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &GetObjectHeaders{}
	_result = &GetObjectResponse{}
	_body, _err := client.GetObjectWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetObjectWithOptions(bucket *string, key *string, request *GetObjectRequest, headers *GetObjectHeaders, runtime *util.RuntimeOptions) (_result *GetObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResponseCacheControl)) {
		query["response-cache-control"] = request.ResponseCacheControl
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseContentDisposition)) {
		query["response-content-disposition"] = request.ResponseContentDisposition
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseContentEncoding)) {
		query["response-content-encoding"] = request.ResponseContentEncoding
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseContentLanguage)) {
		query["response-content-language"] = request.ResponseContentLanguage
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseContentType)) {
		query["response-content-type"] = request.ResponseContentType
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseExpires)) {
		query["response-expires"] = request.ResponseExpires
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AcceptEncoding)) {
		realHeaders["Accept-Encoding"] = util.ToJSONString(headers.AcceptEncoding)
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.IfModifiedSince)) {
		realHeaders["If-Modified-Since"] = util.ToJSONString(headers.IfModifiedSince)
	}

	if !tea.BoolValue(util.IsUnset(headers.IfNoneMatch)) {
		realHeaders["If-None-Match"] = util.ToJSONString(headers.IfNoneMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.IfUnmodifiedSince)) {
		realHeaders["If-Unmodified-Since"] = util.ToJSONString(headers.IfUnmodifiedSince)
	}

	if !tea.BoolValue(util.IsUnset(headers.Range)) {
		realHeaders["Range"] = util.ToJSONString(headers.Range)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("binary"),
	}
	_result = &GetObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetObjectAcl(bucket *string, key *string, request *GetObjectAclRequest) (_result *GetObjectAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetObjectAclResponse{}
	_body, _err := client.GetObjectAclWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetObjectAclWithOptions(bucket *string, key *string, request *GetObjectAclRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetObjectAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetObjectAcl"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?acl"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetObjectAclResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetObjectMeta(bucket *string, key *string, request *GetObjectMetaRequest) (_result *GetObjectMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetObjectMetaResponse{}
	_body, _err := client.GetObjectMetaWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetObjectMetaWithOptions(bucket *string, key *string, request *GetObjectMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetObjectMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetObjectMeta"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?objectMeta"),
		Method:      tea.String("HEAD"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("none"),
	}
	_result = &GetObjectMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetObjectTagging(bucket *string, key *string, request *GetObjectTaggingRequest) (_result *GetObjectTaggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetObjectTaggingResponse{}
	_body, _err := client.GetObjectTaggingWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetObjectTaggingWithOptions(bucket *string, key *string, request *GetObjectTaggingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetObjectTaggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetObjectTagging"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?tagging"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetObjectTaggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSymlink(bucket *string, key *string, request *GetSymlinkRequest) (_result *GetSymlinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSymlinkResponse{}
	_body, _err := client.GetSymlinkWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSymlinkWithOptions(bucket *string, key *string, request *GetSymlinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSymlinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetSymlink"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?symlink"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &GetSymlinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVodPlaylist(bucket *string, channel *string, request *GetVodPlaylistRequest) (_result *GetVodPlaylistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVodPlaylistResponse{}
	_body, _err := client.GetVodPlaylistWithOptions(bucket, channel, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVodPlaylistWithOptions(bucket *string, channel *string, request *GetVodPlaylistRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVodPlaylistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVodPlaylist"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(channel) + "?vod"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("binary"),
	}
	_result = &GetVodPlaylistResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) HeadObject(bucket *string, key *string, request *HeadObjectRequest) (_result *HeadObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &HeadObjectHeaders{}
	_result = &HeadObjectResponse{}
	_body, _err := client.HeadObjectWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) HeadObjectWithOptions(bucket *string, key *string, request *HeadObjectRequest, headers *HeadObjectHeaders, runtime *util.RuntimeOptions) (_result *HeadObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.IfMatch)) {
		realHeaders["If-Match"] = util.ToJSONString(headers.IfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.IfModifiedSince)) {
		realHeaders["If-Modified-Since"] = util.ToJSONString(headers.IfModifiedSince)
	}

	if !tea.BoolValue(util.IsUnset(headers.IfNoneMatch)) {
		realHeaders["If-None-Match"] = util.ToJSONString(headers.IfNoneMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.IfUnmodifiedSince)) {
		realHeaders["If-Unmodified-Since"] = util.ToJSONString(headers.IfUnmodifiedSince)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("HeadObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("HEAD"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("none"),
	}
	_result = &HeadObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitiateBucketWorm(bucket *string, request *InitiateBucketWormRequest) (_result *InitiateBucketWormResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InitiateBucketWormResponse{}
	_body, _err := client.InitiateBucketWormWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitiateBucketWormWithOptions(bucket *string, request *InitiateBucketWormRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InitiateBucketWormResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.InitiateWormConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("InitiateBucketWorm"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?worm"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &InitiateBucketWormResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InitiateMultipartUpload(bucket *string, key *string, request *InitiateMultipartUploadRequest) (_result *InitiateMultipartUploadResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &InitiateMultipartUploadHeaders{}
	_result = &InitiateMultipartUploadResponse{}
	_body, _err := client.InitiateMultipartUploadWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InitiateMultipartUploadWithOptions(bucket *string, key *string, request *InitiateMultipartUploadRequest, headers *InitiateMultipartUploadHeaders, runtime *util.RuntimeOptions) (_result *InitiateMultipartUploadResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodingType)) {
		query["encoding-type"] = request.EncodingType
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.CacheControl)) {
		realHeaders["Cache-Control"] = util.ToJSONString(headers.CacheControl)
	}

	if !tea.BoolValue(util.IsUnset(headers.ContentDisposition)) {
		realHeaders["Content-Disposition"] = util.ToJSONString(headers.ContentDisposition)
	}

	if !tea.BoolValue(util.IsUnset(headers.ContentEncoding)) {
		realHeaders["Content-Encoding"] = util.ToJSONString(headers.ContentEncoding)
	}

	if !tea.BoolValue(util.IsUnset(headers.Expires)) {
		realHeaders["Expires"] = util.ToJSONString(headers.Expires)
	}

	if !tea.BoolValue(util.IsUnset(headers.ForbidOverwrite)) {
		realHeaders["x-oss-forbid-overwrite"] = util.ToJSONString(headers.ForbidOverwrite)
	}

	if !tea.BoolValue(util.IsUnset(headers.SseDataEncryption)) {
		realHeaders["x-oss-server-side-data-encryption"] = util.ToJSONString(headers.SseDataEncryption)
	}

	if !tea.BoolValue(util.IsUnset(headers.ServerSideEncryption)) {
		realHeaders["x-oss-server-side-encryption"] = util.ToJSONString(headers.ServerSideEncryption)
	}

	if !tea.BoolValue(util.IsUnset(headers.SseKeyId)) {
		realHeaders["x-oss-server-side-encryption-key-id"] = util.ToJSONString(headers.SseKeyId)
	}

	if !tea.BoolValue(util.IsUnset(headers.StorageClass)) {
		realHeaders["x-oss-storage-class"] = util.ToJSONString(headers.StorageClass)
	}

	if !tea.BoolValue(util.IsUnset(headers.Tagging)) {
		realHeaders["x-oss-tagging"] = util.ToJSONString(headers.Tagging)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("InitiateMultipartUpload"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?uploads"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &InitiateMultipartUploadResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBucketInventory(bucket *string, request *ListBucketInventoryRequest) (_result *ListBucketInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBucketInventoryResponse{}
	_body, _err := client.ListBucketInventoryWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBucketInventoryWithOptions(bucket *string, request *ListBucketInventoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListBucketInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContinuationToken)) {
		query["continuation-token"] = request.ContinuationToken
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBucketInventory"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?inventory"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListBucketInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListBuckets(request *ListBucketsRequest) (_result *ListBucketsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListBucketsResponse{}
	_body, _err := client.ListBucketsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListBucketsWithOptions(request *ListBucketsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListBucketsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["max-keys"] = request.MaxKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListBuckets"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListBucketsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListLiveChannel(bucket *string, request *ListLiveChannelRequest) (_result *ListLiveChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListLiveChannelResponse{}
	_body, _err := client.ListLiveChannelWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListLiveChannelWithOptions(bucket *string, request *ListLiveChannelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListLiveChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["max-keys"] = request.MaxKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListLiveChannel"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?live"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListLiveChannelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMultipartUploads(bucket *string, request *ListMultipartUploadsRequest) (_result *ListMultipartUploadsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMultipartUploadsResponse{}
	_body, _err := client.ListMultipartUploadsWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMultipartUploadsWithOptions(bucket *string, request *ListMultipartUploadsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMultipartUploadsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Delimiter)) {
		query["delimiter"] = request.Delimiter
	}

	if !tea.BoolValue(util.IsUnset(request.EncodingType)) {
		query["encoding-type"] = request.EncodingType
	}

	if !tea.BoolValue(util.IsUnset(request.KeyMarker)) {
		query["key-marker"] = request.KeyMarker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxUploads)) {
		query["max-uploads"] = request.MaxUploads
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.UploadIdMarker)) {
		query["upload-id-marker"] = request.UploadIdMarker
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMultipartUploads"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?uploads"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListMultipartUploadsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListObjectVersions(bucket *string, request *ListObjectVersionsRequest) (_result *ListObjectVersionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListObjectVersionsResponse{}
	_body, _err := client.ListObjectVersionsWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListObjectVersionsWithOptions(bucket *string, request *ListObjectVersionsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListObjectVersionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Delimiter)) {
		query["delimiter"] = request.Delimiter
	}

	if !tea.BoolValue(util.IsUnset(request.EncodingType)) {
		query["encoding-type"] = request.EncodingType
	}

	if !tea.BoolValue(util.IsUnset(request.KeyMarker)) {
		query["key-marker"] = request.KeyMarker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["max-keys"] = request.MaxKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.VersionIdMarker)) {
		query["version-id-marker"] = request.VersionIdMarker
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListObjectVersions"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?versions"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListObjectVersionsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListObjects(bucket *string, request *ListObjectsRequest) (_result *ListObjectsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListObjectsResponse{}
	_body, _err := client.ListObjectsWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListObjectsWithOptions(bucket *string, request *ListObjectsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListObjectsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Delimiter)) {
		query["delimiter"] = request.Delimiter
	}

	if !tea.BoolValue(util.IsUnset(request.EncodingType)) {
		query["encoding-type"] = request.EncodingType
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		query["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["max-keys"] = request.MaxKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListObjects"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListObjectsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListObjectsV2(bucket *string, request *ListObjectsV2Request) (_result *ListObjectsV2Response, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListObjectsV2Response{}
	_body, _err := client.ListObjectsV2WithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListObjectsV2WithOptions(bucket *string, request *ListObjectsV2Request, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListObjectsV2Response, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContinuationToken)) {
		query["continuation-token"] = request.ContinuationToken
	}

	if !tea.BoolValue(util.IsUnset(request.Delimiter)) {
		query["delimiter"] = request.Delimiter
	}

	if !tea.BoolValue(util.IsUnset(request.EncodingType)) {
		query["encoding-type"] = request.EncodingType
	}

	if !tea.BoolValue(util.IsUnset(request.FetchOwner)) {
		query["fetch-owner"] = request.FetchOwner
	}

	if !tea.BoolValue(util.IsUnset(request.MaxKeys)) {
		query["max-keys"] = request.MaxKeys
	}

	if !tea.BoolValue(util.IsUnset(request.Prefix)) {
		query["prefix"] = request.Prefix
	}

	if !tea.BoolValue(util.IsUnset(request.StartAfter)) {
		query["start-after"] = request.StartAfter
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListObjectsV2"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?list-type=2"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListObjectsV2Response{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListParts(bucket *string, key *string, request *ListPartsRequest) (_result *ListPartsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListPartsResponse{}
	_body, _err := client.ListPartsWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPartsWithOptions(bucket *string, key *string, tmpReq *ListPartsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListPartsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	request := &ListPartsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.EncodingType)) {
		request.EncodingTypeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.EncodingType, tea.String("encoding-type"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EncodingTypeShrink)) {
		query["encoding-type"] = request.EncodingTypeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.MaxParts)) {
		query["max-parts"] = request.MaxParts
	}

	if !tea.BoolValue(util.IsUnset(request.PartNumberMarker)) {
		query["part-number-marker"] = request.PartNumberMarker
	}

	if !tea.BoolValue(util.IsUnset(request.UploadId)) {
		query["uploadId"] = request.UploadId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListParts"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &ListPartsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) OptionObject(bucket *string, key *string) (_result *OptionObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &OptionObjectHeaders{}
	_result = &OptionObjectResponse{}
	_body, _err := client.OptionObjectWithOptions(bucket, key, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) OptionObjectWithOptions(bucket *string, key *string, headers *OptionObjectHeaders, runtime *util.RuntimeOptions) (_result *OptionObjectResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.AccessControlRequestHeaders)) {
		realHeaders["Access-Control-Request-Headers"] = util.ToJSONString(headers.AccessControlRequestHeaders)
	}

	if !tea.BoolValue(util.IsUnset(headers.AccessControlRequestMethod)) {
		realHeaders["Access-Control-Request-Method"] = util.ToJSONString(headers.AccessControlRequestMethod)
	}

	if !tea.BoolValue(util.IsUnset(headers.Origin)) {
		realHeaders["Origin"] = util.ToJSONString(headers.Origin)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("OptionObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("OPTIONS"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &OptionObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostObject(bucket *string) (_result *PostObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PostObjectResponse{}
	_body, _err := client.PostObjectWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostObjectWithOptions(bucket *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PostObjectResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("PostObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("multiFormData"),
		BodyType:    tea.String("xml"),
	}
	_result = &PostObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PostVodPlaylist(bucket *string, channel *string, playlist *string, request *PostVodPlaylistRequest) (_result *PostVodPlaylistResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PostVodPlaylistResponse{}
	_body, _err := client.PostVodPlaylistWithOptions(bucket, channel, playlist, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PostVodPlaylistWithOptions(bucket *string, channel *string, playlist *string, request *PostVodPlaylistRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PostVodPlaylistResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["endTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["startTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PostVodPlaylist"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(channel) + "/" + tea.StringValue(playlist) + "?vod"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PostVodPlaylistResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucket(bucket *string, request *PutBucketRequest) (_result *PutBucketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PutBucketHeaders{}
	_result = &PutBucketResponse{}
	_body, _err := client.PutBucketWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketWithOptions(bucket *string, request *PutBucketRequest, headers *PutBucketHeaders, runtime *util.RuntimeOptions) (_result *PutBucketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Acl)) {
		realHeaders["x-oss-acl"] = util.ToJSONString(headers.Acl)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.CreateBucketConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucket"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketAcl(bucket *string) (_result *PutBucketAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PutBucketAclHeaders{}
	_result = &PutBucketAclResponse{}
	_body, _err := client.PutBucketAclWithOptions(bucket, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketAclWithOptions(bucket *string, headers *PutBucketAclHeaders, runtime *util.RuntimeOptions) (_result *PutBucketAclResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Acl)) {
		realHeaders["x-oss-acl"] = util.ToJSONString(headers.Acl)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketAcl"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?acl"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketAclResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketCors(bucket *string, request *PutBucketCorsRequest) (_result *PutBucketCorsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketCorsResponse{}
	_body, _err := client.PutBucketCorsWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketCorsWithOptions(bucket *string, request *PutBucketCorsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketCorsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.CORSConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketCors"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?cors"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketCorsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketEncryption(bucket *string, request *PutBucketEncryptionRequest) (_result *PutBucketEncryptionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketEncryptionResponse{}
	_body, _err := client.PutBucketEncryptionWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketEncryptionWithOptions(bucket *string, request *PutBucketEncryptionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketEncryptionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.ServerSideEncryptionRule)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketEncryption"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?encryption"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketEncryptionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketInventory(bucket *string, request *PutBucketInventoryRequest) (_result *PutBucketInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketInventoryResponse{}
	_body, _err := client.PutBucketInventoryWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketInventoryWithOptions(bucket *string, request *PutBucketInventoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InventoryId)) {
		query["inventoryId"] = request.InventoryId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(tea.ToMap(request.InventoryConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketInventory"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?inventory"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketInventoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketLifecycle(bucket *string, request *PutBucketLifecycleRequest) (_result *PutBucketLifecycleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketLifecycleResponse{}
	_body, _err := client.PutBucketLifecycleWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketLifecycleWithOptions(bucket *string, request *PutBucketLifecycleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketLifecycleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.LifecycleConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketLifecycle"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?lifecycle"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketLifecycleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketLogging(bucket *string, request *PutBucketLoggingRequest) (_result *PutBucketLoggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketLoggingResponse{}
	_body, _err := client.PutBucketLoggingWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketLoggingWithOptions(bucket *string, request *PutBucketLoggingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketLoggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.BucketLoggingStatus)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketLogging"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?logging"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketLoggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketPolicy(bucket *string, request *PutBucketPolicyRequest) (_result *PutBucketPolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketPolicyResponse{}
	_body, _err := client.PutBucketPolicyWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketPolicyWithOptions(bucket *string, request *PutBucketPolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketPolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    request.Policy,
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketPolicy"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?policy"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &PutBucketPolicyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketReferer(bucket *string, request *PutBucketRefererRequest) (_result *PutBucketRefererResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketRefererResponse{}
	_body, _err := client.PutBucketRefererWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketRefererWithOptions(bucket *string, request *PutBucketRefererRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketRefererResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.RefererConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketReferer"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?referer"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketRefererResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketReplication(bucket *string, request *PutBucketReplicationRequest) (_result *PutBucketReplicationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketReplicationResponse{}
	_body, _err := client.PutBucketReplicationWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketReplicationWithOptions(bucket *string, request *PutBucketReplicationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketReplicationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.ReplicationConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketReplication"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?replication&comp=add"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketReplicationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketRequestPayment(bucket *string, request *PutBucketRequestPaymentRequest) (_result *PutBucketRequestPaymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketRequestPaymentResponse{}
	_body, _err := client.PutBucketRequestPaymentWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketRequestPaymentWithOptions(bucket *string, request *PutBucketRequestPaymentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketRequestPaymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.RequestPaymentConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketRequestPayment"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?requestPayment"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketRequestPaymentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketTags(bucket *string, request *PutBucketTagsRequest) (_result *PutBucketTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketTagsResponse{}
	_body, _err := client.PutBucketTagsWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketTagsWithOptions(bucket *string, request *PutBucketTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Tagging)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketTags"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?tagging"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketTagsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketTransferAcceleration(bucket *string, request *PutBucketTransferAccelerationRequest) (_result *PutBucketTransferAccelerationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketTransferAccelerationResponse{}
	_body, _err := client.PutBucketTransferAccelerationWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketTransferAccelerationWithOptions(bucket *string, request *PutBucketTransferAccelerationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketTransferAccelerationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.TransferAccelerationConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketTransferAcceleration"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?transferAcceleration"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketTransferAccelerationResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketVersioning(bucket *string, request *PutBucketVersioningRequest) (_result *PutBucketVersioningResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketVersioningResponse{}
	_body, _err := client.PutBucketVersioningWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketVersioningWithOptions(bucket *string, request *PutBucketVersioningRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketVersioningResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.VersioningConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketVersioning"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?versioning"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketVersioningResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutBucketWebsite(bucket *string, request *PutBucketWebsiteRequest) (_result *PutBucketWebsiteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutBucketWebsiteResponse{}
	_body, _err := client.PutBucketWebsiteWithOptions(bucket, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutBucketWebsiteWithOptions(bucket *string, request *PutBucketWebsiteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutBucketWebsiteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.WebsiteConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutBucketWebsite"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/?website"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutBucketWebsiteResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutLiveChannel(bucket *string, channel *string, request *PutLiveChannelRequest) (_result *PutLiveChannelResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutLiveChannelResponse{}
	_body, _err := client.PutLiveChannelWithOptions(bucket, channel, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutLiveChannelWithOptions(bucket *string, channel *string, request *PutLiveChannelRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutLiveChannelResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.LiveChannelConfiguration)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutLiveChannel"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(channel) + "?live"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutLiveChannelResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutLiveChannelStatus(bucket *string, channel *string, request *PutLiveChannelStatusRequest) (_result *PutLiveChannelStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutLiveChannelStatusResponse{}
	_body, _err := client.PutLiveChannelStatusWithOptions(bucket, channel, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutLiveChannelStatusWithOptions(bucket *string, channel *string, request *PutLiveChannelStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutLiveChannelStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["status"] = request.Status
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutLiveChannelStatus"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(channel) + "?live"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutLiveChannelStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutObject(bucket *string, key *string, request *PutObjectRequest) (_result *PutObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PutObjectHeaders{}
	_result = &PutObjectResponse{}
	_body, _err := client.PutObjectWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutObjectWithOptions(bucket *string, key *string, request *PutObjectRequest, headers *PutObjectHeaders, runtime *util.RuntimeOptions) (_result *PutObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ForbidOverwrite)) {
		realHeaders["x-oss-forbid-overwrite"] = util.ToJSONString(headers.ForbidOverwrite)
	}

	if !tea.BoolValue(util.IsUnset(headers.MetaData)) {
		realHeaders["x-oss-meta-*"] = util.ToJSONString(headers.MetaData)
	}

	if !tea.BoolValue(util.IsUnset(headers.Acl)) {
		realHeaders["x-oss-object-acl"] = util.ToJSONString(headers.Acl)
	}

	if !tea.BoolValue(util.IsUnset(headers.SseDataEncryption)) {
		realHeaders["x-oss-server-side-data-encryption"] = util.ToJSONString(headers.SseDataEncryption)
	}

	if !tea.BoolValue(util.IsUnset(headers.ServerSideEncryption)) {
		realHeaders["x-oss-server-side-encryption"] = util.ToJSONString(headers.ServerSideEncryption)
	}

	if !tea.BoolValue(util.IsUnset(headers.SseKeyId)) {
		realHeaders["x-oss-server-side-encryption-key-id"] = util.ToJSONString(headers.SseKeyId)
	}

	if !tea.BoolValue(util.IsUnset(headers.StorageClass)) {
		realHeaders["x-oss-storage-class"] = util.ToJSONString(headers.StorageClass)
	}

	if !tea.BoolValue(util.IsUnset(headers.Tagging)) {
		realHeaders["x-oss-tagging"] = util.ToJSONString(headers.Tagging)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Body:    request.Body,
		Stream:  request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("PutObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("binary"),
		BodyType:    tea.String("none"),
	}
	_result = &PutObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutObjectAcl(bucket *string, key *string, request *PutObjectAclRequest) (_result *PutObjectAclResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PutObjectAclHeaders{}
	_result = &PutObjectAclResponse{}
	_body, _err := client.PutObjectAclWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutObjectAclWithOptions(bucket *string, key *string, request *PutObjectAclRequest, headers *PutObjectAclHeaders, runtime *util.RuntimeOptions) (_result *PutObjectAclResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.Acl)) {
		realHeaders["x-oss-object-acl"] = util.ToJSONString(headers.Acl)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("PutObjectAcl"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?acl"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutObjectAclResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutObjectTagging(bucket *string, key *string, request *PutObjectTaggingRequest) (_result *PutObjectTaggingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PutObjectTaggingResponse{}
	_body, _err := client.PutObjectTaggingWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutObjectTaggingWithOptions(bucket *string, key *string, request *PutObjectTaggingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PutObjectTaggingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(tea.ToMap(request.Tagging)),
	}
	params := &openapi.Params{
		Action:      tea.String("PutObjectTagging"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?tagging"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutObjectTaggingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) PutSymlink(bucket *string, key *string) (_result *PutSymlinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &PutSymlinkHeaders{}
	_result = &PutSymlinkResponse{}
	_body, _err := client.PutSymlinkWithOptions(bucket, key, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) PutSymlinkWithOptions(bucket *string, key *string, headers *PutSymlinkHeaders, runtime *util.RuntimeOptions) (_result *PutSymlinkResponse, _err error) {
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.ForbidOverwrite)) {
		realHeaders["x-oss-forbid-overwrite"] = util.ToJSONString(headers.ForbidOverwrite)
	}

	if !tea.BoolValue(util.IsUnset(headers.Acl)) {
		realHeaders["x-oss-object-acl"] = util.ToJSONString(headers.Acl)
	}

	if !tea.BoolValue(util.IsUnset(headers.StorageClass)) {
		realHeaders["x-oss-storage-class"] = util.ToJSONString(headers.StorageClass)
	}

	if !tea.BoolValue(util.IsUnset(headers.SymlinkTargetKey)) {
		realHeaders["x-oss-symlink-target"] = util.ToJSONString(headers.SymlinkTargetKey)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
	}
	params := &openapi.Params{
		Action:      tea.String("PutSymlink"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?symlink"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &PutSymlinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RestoreObject(bucket *string, key *string, request *RestoreObjectRequest) (_result *RestoreObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestoreObjectResponse{}
	_body, _err := client.RestoreObjectWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RestoreObjectWithOptions(bucket *string, key *string, request *RestoreObjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestoreObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VersionId)) {
		query["versionId"] = request.VersionId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(tea.ToMap(request.RestoreRequest)),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key) + "?restore"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &RestoreObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SelectObject(bucket *string, key *string, request *SelectObjectRequest) (_result *SelectObjectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SelectObjectResponse{}
	_body, _err := client.SelectObjectWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SelectObjectWithOptions(bucket *string, key *string, request *SelectObjectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SelectObjectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Body:    openapiutil.ParseToMap(tea.ToMap(request.SelectRequest)),
	}
	params := &openapi.Params{
		Action:      tea.String("SelectObject"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("binary"),
	}
	_result = &SelectObjectResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadPart(bucket *string, key *string, request *UploadPartRequest) (_result *UploadPartResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadPartResponse{}
	_body, _err := client.UploadPartWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadPartWithOptions(bucket *string, key *string, request *UploadPartRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadPartResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PartNumber)) {
		query["partNumber"] = request.PartNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UploadId)) {
		query["uploadId"] = request.UploadId
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    request.Body,
		Stream:  request.Body,
	}
	params := &openapi.Params{
		Action:      tea.String("UploadPart"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("binary"),
		BodyType:    tea.String("xml"),
	}
	_result = &UploadPartResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UploadPartCopy(bucket *string, key *string, request *UploadPartCopyRequest) (_result *UploadPartCopyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := &UploadPartCopyHeaders{}
	_result = &UploadPartCopyResponse{}
	_body, _err := client.UploadPartCopyWithOptions(bucket, key, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UploadPartCopyWithOptions(bucket *string, key *string, request *UploadPartCopyRequest, headers *UploadPartCopyHeaders, runtime *util.RuntimeOptions) (_result *UploadPartCopyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	hostMap := make(map[string]*string)
	hostMap["bucket"] = bucket
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PartNumber)) {
		query["partNumber"] = request.PartNumber
	}

	if !tea.BoolValue(util.IsUnset(request.UploadId)) {
		query["uploadId"] = request.UploadId
	}

	realHeaders := make(map[string]*string)
	if !tea.BoolValue(util.IsUnset(headers.CommonHeaders)) {
		realHeaders = headers.CommonHeaders
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySource)) {
		realHeaders["x-oss-copy-source"] = util.ToJSONString(headers.CopySource)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceIfMatch)) {
		realHeaders["x-oss-copy-source-if-match"] = util.ToJSONString(headers.CopySourceIfMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceIfModifiedSince)) {
		realHeaders["x-oss-copy-source-if-modified-since"] = util.ToJSONString(headers.CopySourceIfModifiedSince)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceIfNoneMatch)) {
		realHeaders["x-oss-copy-source-if-none-match"] = util.ToJSONString(headers.CopySourceIfNoneMatch)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceIfUnmodifiedSince)) {
		realHeaders["x-oss-copy-source-if-unmodified-since"] = util.ToJSONString(headers.CopySourceIfUnmodifiedSince)
	}

	if !tea.BoolValue(util.IsUnset(headers.CopySourceRange)) {
		realHeaders["x-oss-copy-source-range"] = util.ToJSONString(headers.CopySourceRange)
	}

	req := &openapi.OpenApiRequest{
		HostMap: hostMap,
		Headers: realHeaders,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UploadPartCopy"),
		Version:     tea.String("2019-05-17"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/" + tea.StringValue(key)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("xml"),
		BodyType:    tea.String("xml"),
	}
	_result = &UploadPartCopyResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
