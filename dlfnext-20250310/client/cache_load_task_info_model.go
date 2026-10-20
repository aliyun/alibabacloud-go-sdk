// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCacheLoadTaskInfo interface {
	dara.Model
	String() string
	GoString() string
	SetCatalogId(v string) *CacheLoadTaskInfo
	GetCatalogId() *string
	SetCompletedFiles(v int64) *CacheLoadTaskInfo
	GetCompletedFiles() *int64
	SetDatabase(v string) *CacheLoadTaskInfo
	GetDatabase() *string
	SetEnumeratedFiles(v int64) *CacheLoadTaskInfo
	GetEnumeratedFiles() *int64
	SetEnumerationDone(v bool) *CacheLoadTaskInfo
	GetEnumerationDone() *bool
	SetExtraOptions(v map[string]*string) *CacheLoadTaskInfo
	GetExtraOptions() map[string]*string
	SetFailedFiles(v int64) *CacheLoadTaskInfo
	GetFailedFiles() *int64
	SetLastError(v string) *CacheLoadTaskInfo
	GetLastError() *string
	SetLoadedBytes(v int64) *CacheLoadTaskInfo
	GetLoadedBytes() *int64
	SetPath(v string) *CacheLoadTaskInfo
	GetPath() *string
	SetStartedAt(v int64) *CacheLoadTaskInfo
	GetStartedAt() *int64
	SetStatus(v string) *CacheLoadTaskInfo
	GetStatus() *string
	SetTable(v string) *CacheLoadTaskInfo
	GetTable() *string
	SetTaskId(v string) *CacheLoadTaskInfo
	GetTaskId() *string
	SetTaskUpdatedAt(v int64) *CacheLoadTaskInfo
	GetTaskUpdatedAt() *int64
	SetTotalBytes(v int64) *CacheLoadTaskInfo
	GetTotalBytes() *int64
	SetTotalFiles(v int64) *CacheLoadTaskInfo
	GetTotalFiles() *int64
}

type CacheLoadTaskInfo struct {
	// The data catalog ID.
	//
	// example:
	//
	// clg-paimon-xxxx
	CatalogId *string `json:"catalogId,omitempty" xml:"catalogId,omitempty"`
	// The number of files with completed cache loading.
	//
	// example:
	//
	// 100
	CompletedFiles *int64 `json:"completedFiles,omitempty" xml:"completedFiles,omitempty"`
	// The database name.
	//
	// example:
	//
	// database_name
	Database *string `json:"database,omitempty" xml:"database,omitempty"`
	// The number of enumerated files.
	//
	// example:
	//
	// 100
	EnumeratedFiles *int64 `json:"enumeratedFiles,omitempty" xml:"enumeratedFiles,omitempty"`
	// The file enumeration completion status.
	//
	// example:
	//
	// true
	EnumerationDone *bool `json:"enumerationDone,omitempty" xml:"enumerationDone,omitempty"`
	// The additional parameters.
	ExtraOptions map[string]*string `json:"extraOptions,omitempty" xml:"extraOptions,omitempty"`
	// The number of files with failed cache loading.
	//
	// example:
	//
	// 0
	FailedFiles *int64 `json:"failedFiles,omitempty" xml:"failedFiles,omitempty"`
	// The most recent error message, which is empty when the status is normal.
	//
	// example:
	//
	// null
	LastError *string `json:"lastError,omitempty" xml:"lastError,omitempty"`
	// The number of bytes loaded into the cache.
	//
	// example:
	//
	// 1024
	LoadedBytes *int64 `json:"loadedBytes,omitempty" xml:"loadedBytes,omitempty"`
	// The storage location.
	//
	// example:
	//
	// oss://clg-xxxx/db-xxxx/tbl-xxxx
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The start time of the task execution, in milliseconds.
	//
	// example:
	//
	// 1790576716926
	StartedAt *int64 `json:"startedAt,omitempty" xml:"startedAt,omitempty"`
	// The task execution status.
	//
	// example:
	//
	// SUCCEEDED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The table name.
	//
	// example:
	//
	// table_name
	Table *string `json:"table,omitempty" xml:"table,omitempty"`
	// The task ID.
	//
	// example:
	//
	// lc-81afc0f7812ec737-1790576716926
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
	// The task update time, in milliseconds.
	//
	// example:
	//
	// 1790576746926
	TaskUpdatedAt *int64 `json:"taskUpdatedAt,omitempty" xml:"taskUpdatedAt,omitempty"`
	// The total number of bytes.
	//
	// example:
	//
	// 1024
	TotalBytes *int64 `json:"totalBytes,omitempty" xml:"totalBytes,omitempty"`
	// The total number of files.
	//
	// example:
	//
	// 100
	TotalFiles *int64 `json:"totalFiles,omitempty" xml:"totalFiles,omitempty"`
}

func (s CacheLoadTaskInfo) String() string {
	return dara.Prettify(s)
}

func (s CacheLoadTaskInfo) GoString() string {
	return s.String()
}

func (s *CacheLoadTaskInfo) GetCatalogId() *string {
	return s.CatalogId
}

func (s *CacheLoadTaskInfo) GetCompletedFiles() *int64 {
	return s.CompletedFiles
}

func (s *CacheLoadTaskInfo) GetDatabase() *string {
	return s.Database
}

func (s *CacheLoadTaskInfo) GetEnumeratedFiles() *int64 {
	return s.EnumeratedFiles
}

func (s *CacheLoadTaskInfo) GetEnumerationDone() *bool {
	return s.EnumerationDone
}

func (s *CacheLoadTaskInfo) GetExtraOptions() map[string]*string {
	return s.ExtraOptions
}

func (s *CacheLoadTaskInfo) GetFailedFiles() *int64 {
	return s.FailedFiles
}

func (s *CacheLoadTaskInfo) GetLastError() *string {
	return s.LastError
}

func (s *CacheLoadTaskInfo) GetLoadedBytes() *int64 {
	return s.LoadedBytes
}

func (s *CacheLoadTaskInfo) GetPath() *string {
	return s.Path
}

func (s *CacheLoadTaskInfo) GetStartedAt() *int64 {
	return s.StartedAt
}

func (s *CacheLoadTaskInfo) GetStatus() *string {
	return s.Status
}

func (s *CacheLoadTaskInfo) GetTable() *string {
	return s.Table
}

func (s *CacheLoadTaskInfo) GetTaskId() *string {
	return s.TaskId
}

func (s *CacheLoadTaskInfo) GetTaskUpdatedAt() *int64 {
	return s.TaskUpdatedAt
}

func (s *CacheLoadTaskInfo) GetTotalBytes() *int64 {
	return s.TotalBytes
}

func (s *CacheLoadTaskInfo) GetTotalFiles() *int64 {
	return s.TotalFiles
}

func (s *CacheLoadTaskInfo) SetCatalogId(v string) *CacheLoadTaskInfo {
	s.CatalogId = &v
	return s
}

func (s *CacheLoadTaskInfo) SetCompletedFiles(v int64) *CacheLoadTaskInfo {
	s.CompletedFiles = &v
	return s
}

func (s *CacheLoadTaskInfo) SetDatabase(v string) *CacheLoadTaskInfo {
	s.Database = &v
	return s
}

func (s *CacheLoadTaskInfo) SetEnumeratedFiles(v int64) *CacheLoadTaskInfo {
	s.EnumeratedFiles = &v
	return s
}

func (s *CacheLoadTaskInfo) SetEnumerationDone(v bool) *CacheLoadTaskInfo {
	s.EnumerationDone = &v
	return s
}

func (s *CacheLoadTaskInfo) SetExtraOptions(v map[string]*string) *CacheLoadTaskInfo {
	s.ExtraOptions = v
	return s
}

func (s *CacheLoadTaskInfo) SetFailedFiles(v int64) *CacheLoadTaskInfo {
	s.FailedFiles = &v
	return s
}

func (s *CacheLoadTaskInfo) SetLastError(v string) *CacheLoadTaskInfo {
	s.LastError = &v
	return s
}

func (s *CacheLoadTaskInfo) SetLoadedBytes(v int64) *CacheLoadTaskInfo {
	s.LoadedBytes = &v
	return s
}

func (s *CacheLoadTaskInfo) SetPath(v string) *CacheLoadTaskInfo {
	s.Path = &v
	return s
}

func (s *CacheLoadTaskInfo) SetStartedAt(v int64) *CacheLoadTaskInfo {
	s.StartedAt = &v
	return s
}

func (s *CacheLoadTaskInfo) SetStatus(v string) *CacheLoadTaskInfo {
	s.Status = &v
	return s
}

func (s *CacheLoadTaskInfo) SetTable(v string) *CacheLoadTaskInfo {
	s.Table = &v
	return s
}

func (s *CacheLoadTaskInfo) SetTaskId(v string) *CacheLoadTaskInfo {
	s.TaskId = &v
	return s
}

func (s *CacheLoadTaskInfo) SetTaskUpdatedAt(v int64) *CacheLoadTaskInfo {
	s.TaskUpdatedAt = &v
	return s
}

func (s *CacheLoadTaskInfo) SetTotalBytes(v int64) *CacheLoadTaskInfo {
	s.TotalBytes = &v
	return s
}

func (s *CacheLoadTaskInfo) SetTotalFiles(v int64) *CacheLoadTaskInfo {
	s.TotalFiles = &v
	return s
}

func (s *CacheLoadTaskInfo) Validate() error {
	return dara.Validate(s)
}
