// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobRunDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDIJobRunDetailsResponseBodyPagingInfo) *ListDIJobRunDetailsResponseBody
	GetPagingInfo() *ListDIJobRunDetailsResponseBodyPagingInfo
	SetRequestId(v string) *ListDIJobRunDetailsResponseBody
	GetRequestId() *string
}

type ListDIJobRunDetailsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDIJobRunDetailsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobRunDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobRunDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsResponseBody) GetPagingInfo() *ListDIJobRunDetailsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDIJobRunDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDIJobRunDetailsResponseBody) SetPagingInfo(v *ListDIJobRunDetailsResponseBodyPagingInfo) *ListDIJobRunDetailsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobRunDetailsResponseBody) SetRequestId(v string) *ListDIJobRunDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDIJobRunDetailsResponseBodyPagingInfo struct {
	// The running information about the synchronization task.
	JobRunInfos []*ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos `json:"JobRunInfos,omitempty" xml:"JobRunInfos,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 131
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobRunDetailsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobRunDetailsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) GetJobRunInfos() []*ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	return s.JobRunInfos
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) GetTotalCount() *string {
	return s.TotalCount
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) SetJobRunInfos(v []*ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) *ListDIJobRunDetailsResponseBodyPagingInfo {
	s.JobRunInfos = v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) SetPageNumber(v string) *ListDIJobRunDetailsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) SetPageSize(v string) *ListDIJobRunDetailsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) SetTotalCount(v string) *ListDIJobRunDetailsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos struct {
	// The name of the database in the destination.
	//
	// example:
	//
	// dst_db
	DestinationDatabaseName *string `json:"DestinationDatabaseName,omitempty" xml:"DestinationDatabaseName,omitempty"`
	// The name of the destination.
	//
	// example:
	//
	// dst_name
	DestinationDatasourceName *string `json:"DestinationDatasourceName,omitempty" xml:"DestinationDatasourceName,omitempty"`
	// The name of the schema of the destination.
	//
	// example:
	//
	// dst_schema
	DestinationSchemaName *string `json:"DestinationSchemaName,omitempty" xml:"DestinationSchemaName,omitempty"`
	// The name of the table in the destination.
	//
	// example:
	//
	// dst_name
	DestinationTableName *string `json:"DestinationTableName,omitempty" xml:"DestinationTableName,omitempty"`
	// The error message that is returned if an error occurs during full batch synchronization. If no error occurs, no value is returned for this parameter.
	//
	// example:
	//
	// sync table t1 fail.
	FullMigrationErrorMessage *string `json:"FullMigrationErrorMessage,omitempty" xml:"FullMigrationErrorMessage,omitempty"`
	// The status of full batch synchronization.
	//
	// example:
	//
	// Finished
	FullMigrationStatus *string `json:"FullMigrationStatus,omitempty" xml:"FullMigrationStatus,omitempty"`
	// The total number of errors that occur during full synchronization.
	//
	// example:
	//
	// 0
	OfflineErrorRecords *int64 `json:"OfflineErrorRecords,omitempty" xml:"OfflineErrorRecords,omitempty"`
	// The total number of bytes that are synchronized during full synchronization.
	//
	// example:
	//
	// 100
	OfflineTotalBytes *int64 `json:"OfflineTotalBytes,omitempty" xml:"OfflineTotalBytes,omitempty"`
	// The total number of data records that are synchronized during full synchronization.
	//
	// example:
	//
	// 10
	OfflineTotalRecords *int64 `json:"OfflineTotalRecords,omitempty" xml:"OfflineTotalRecords,omitempty"`
	// The error message that is returned if an error occurs during real-time synchronization. If no error occurs, no value is returned for this parameter.
	//
	// example:
	//
	// sync table t1 fail.
	RealtimeMigrationErrorMessage *string `json:"RealtimeMigrationErrorMessage,omitempty" xml:"RealtimeMigrationErrorMessage,omitempty"`
	// The status of real-time synchronization.
	//
	// example:
	//
	// Running
	RealtimeMigrationStatus *string `json:"RealtimeMigrationStatus,omitempty" xml:"RealtimeMigrationStatus,omitempty"`
	// The name of the database in the source.
	//
	// example:
	//
	// db_name
	SourceDatabaseName *string `json:"SourceDatabaseName,omitempty" xml:"SourceDatabaseName,omitempty"`
	// The name of the source.
	//
	// example:
	//
	// ds_name
	SourceDatasourceName *string `json:"SourceDatasourceName,omitempty" xml:"SourceDatasourceName,omitempty"`
	// The name of the schema of the source.
	//
	// example:
	//
	// schema_name
	SourceSchemaName *string `json:"SourceSchemaName,omitempty" xml:"SourceSchemaName,omitempty"`
	// The name of the table in the source.
	//
	// example:
	//
	// table_name
	SourceTableName *string `json:"SourceTableName,omitempty" xml:"SourceTableName,omitempty"`
	// The error message that is returned if an error occurs during schema synchronization. If no error occurs, no value is returned for this parameter.
	//
	// example:
	//
	// create table t1 fail.
	StructureMigrationErrorMessage *string `json:"StructureMigrationErrorMessage,omitempty" xml:"StructureMigrationErrorMessage,omitempty"`
	// The synchronization status of the schema.
	//
	// example:
	//
	// Finished
	StructureMigrationStatus *string `json:"StructureMigrationStatus,omitempty" xml:"StructureMigrationStatus,omitempty"`
}

func (s ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetDestinationDatabaseName() *string {
	return s.DestinationDatabaseName
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetDestinationDatasourceName() *string {
	return s.DestinationDatasourceName
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetDestinationSchemaName() *string {
	return s.DestinationSchemaName
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetDestinationTableName() *string {
	return s.DestinationTableName
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetFullMigrationErrorMessage() *string {
	return s.FullMigrationErrorMessage
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetFullMigrationStatus() *string {
	return s.FullMigrationStatus
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetOfflineErrorRecords() *int64 {
	return s.OfflineErrorRecords
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetOfflineTotalBytes() *int64 {
	return s.OfflineTotalBytes
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetOfflineTotalRecords() *int64 {
	return s.OfflineTotalRecords
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetRealtimeMigrationErrorMessage() *string {
	return s.RealtimeMigrationErrorMessage
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetRealtimeMigrationStatus() *string {
	return s.RealtimeMigrationStatus
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetSourceDatabaseName() *string {
	return s.SourceDatabaseName
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetSourceDatasourceName() *string {
	return s.SourceDatasourceName
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetSourceSchemaName() *string {
	return s.SourceSchemaName
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetStructureMigrationErrorMessage() *string {
	return s.StructureMigrationErrorMessage
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) GetStructureMigrationStatus() *string {
	return s.StructureMigrationStatus
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetDestinationDatabaseName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.DestinationDatabaseName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetDestinationDatasourceName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.DestinationDatasourceName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetDestinationSchemaName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.DestinationSchemaName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetDestinationTableName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.DestinationTableName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetFullMigrationErrorMessage(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.FullMigrationErrorMessage = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetFullMigrationStatus(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.FullMigrationStatus = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetOfflineErrorRecords(v int64) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.OfflineErrorRecords = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetOfflineTotalBytes(v int64) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.OfflineTotalBytes = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetOfflineTotalRecords(v int64) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.OfflineTotalRecords = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetRealtimeMigrationErrorMessage(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.RealtimeMigrationErrorMessage = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetRealtimeMigrationStatus(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.RealtimeMigrationStatus = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetSourceDatabaseName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.SourceDatabaseName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetSourceDatasourceName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.SourceDatasourceName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetSourceSchemaName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.SourceSchemaName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetSourceTableName(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.SourceTableName = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetStructureMigrationErrorMessage(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.StructureMigrationErrorMessage = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) SetStructureMigrationStatus(v string) *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos {
	s.StructureMigrationStatus = &v
	return s
}

func (s *ListDIJobRunDetailsResponseBodyPagingInfoJobRunInfos) Validate() error {
	return dara.Validate(s)
}
