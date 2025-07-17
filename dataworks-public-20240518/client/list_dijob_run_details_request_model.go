// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobRunDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *ListDIJobRunDetailsRequest
	GetDIJobId() *int64
	SetInstanceId(v int64) *ListDIJobRunDetailsRequest
	GetInstanceId() *int64
	SetPageNumber(v int64) *ListDIJobRunDetailsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDIJobRunDetailsRequest
	GetPageSize() *int64
	SetSourceDataSourceName(v string) *ListDIJobRunDetailsRequest
	GetSourceDataSourceName() *string
	SetSourceDatabaseName(v string) *ListDIJobRunDetailsRequest
	GetSourceDatabaseName() *string
	SetSourceSchemaName(v string) *ListDIJobRunDetailsRequest
	GetSourceSchemaName() *string
	SetSourceTableName(v string) *ListDIJobRunDetailsRequest
	GetSourceTableName() *string
}

type ListDIJobRunDetailsRequest struct {
	// The ID of the synchronization task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 11265
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// 1234
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The name of the source.
	//
	// example:
	//
	// ds_name
	SourceDataSourceName *string `json:"SourceDataSourceName,omitempty" xml:"SourceDataSourceName,omitempty"`
	// The name of the database in the source.
	//
	// example:
	//
	// db_name
	SourceDatabaseName *string `json:"SourceDatabaseName,omitempty" xml:"SourceDatabaseName,omitempty"`
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
}

func (s ListDIJobRunDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobRunDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobRunDetailsRequest) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIJobRunDetailsRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListDIJobRunDetailsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDIJobRunDetailsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDIJobRunDetailsRequest) GetSourceDataSourceName() *string {
	return s.SourceDataSourceName
}

func (s *ListDIJobRunDetailsRequest) GetSourceDatabaseName() *string {
	return s.SourceDatabaseName
}

func (s *ListDIJobRunDetailsRequest) GetSourceSchemaName() *string {
	return s.SourceSchemaName
}

func (s *ListDIJobRunDetailsRequest) GetSourceTableName() *string {
	return s.SourceTableName
}

func (s *ListDIJobRunDetailsRequest) SetDIJobId(v int64) *ListDIJobRunDetailsRequest {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetInstanceId(v int64) *ListDIJobRunDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetPageNumber(v int64) *ListDIJobRunDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetPageSize(v int64) *ListDIJobRunDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetSourceDataSourceName(v string) *ListDIJobRunDetailsRequest {
	s.SourceDataSourceName = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetSourceDatabaseName(v string) *ListDIJobRunDetailsRequest {
	s.SourceDatabaseName = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetSourceSchemaName(v string) *ListDIJobRunDetailsRequest {
	s.SourceSchemaName = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) SetSourceTableName(v string) *ListDIJobRunDetailsRequest {
	s.SourceTableName = &v
	return s
}

func (s *ListDIJobRunDetailsRequest) Validate() error {
	return dara.Validate(s)
}
