// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSymbolicFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetSymbolicFilesRequest
	GetAppKey() *int64
	SetAppVersion(v string) *GetSymbolicFilesRequest
	GetAppVersion() *string
	SetBuildId(v string) *GetSymbolicFilesRequest
	GetBuildId() *string
	SetEndTime(v int64) *GetSymbolicFilesRequest
	GetEndTime() *int64
	SetExportStatus(v string) *GetSymbolicFilesRequest
	GetExportStatus() *string
	SetFileName(v string) *GetSymbolicFilesRequest
	GetFileName() *string
	SetFileType(v string) *GetSymbolicFilesRequest
	GetFileType() *string
	SetOs(v string) *GetSymbolicFilesRequest
	GetOs() *string
	SetPageIndex(v int32) *GetSymbolicFilesRequest
	GetPageIndex() *int32
	SetPageSize(v int32) *GetSymbolicFilesRequest
	GetPageSize() *int32
	SetStartTime(v int64) *GetSymbolicFilesRequest
	GetStartTime() *int64
	SetUuid(v string) *GetSymbolicFilesRequest
	GetUuid() *string
}

type GetSymbolicFilesRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 24780725
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	BuildId    *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// example:
	//
	// 1743523199999
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// NORMAL
	ExportStatus *string `json:"ExportStatus,omitempty" xml:"ExportStatus,omitempty"`
	// example:
	//
	// app_so.zip
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// APP_SO
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageIndex *int32 `json:"PageIndex,omitempty" xml:"PageIndex,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1742918400000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// uuid
	//
	// example:
	//
	// abcf4a4b-158c-4a0b-b81c-262785d84c4f
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetSymbolicFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSymbolicFilesRequest) GoString() string {
	return s.String()
}

func (s *GetSymbolicFilesRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetSymbolicFilesRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetSymbolicFilesRequest) GetBuildId() *string {
	return s.BuildId
}

func (s *GetSymbolicFilesRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetSymbolicFilesRequest) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *GetSymbolicFilesRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetSymbolicFilesRequest) GetFileType() *string {
	return s.FileType
}

func (s *GetSymbolicFilesRequest) GetOs() *string {
	return s.Os
}

func (s *GetSymbolicFilesRequest) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *GetSymbolicFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetSymbolicFilesRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *GetSymbolicFilesRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetSymbolicFilesRequest) SetAppKey(v int64) *GetSymbolicFilesRequest {
	s.AppKey = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetAppVersion(v string) *GetSymbolicFilesRequest {
	s.AppVersion = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetBuildId(v string) *GetSymbolicFilesRequest {
	s.BuildId = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetEndTime(v int64) *GetSymbolicFilesRequest {
	s.EndTime = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetExportStatus(v string) *GetSymbolicFilesRequest {
	s.ExportStatus = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetFileName(v string) *GetSymbolicFilesRequest {
	s.FileName = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetFileType(v string) *GetSymbolicFilesRequest {
	s.FileType = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetOs(v string) *GetSymbolicFilesRequest {
	s.Os = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetPageIndex(v int32) *GetSymbolicFilesRequest {
	s.PageIndex = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetPageSize(v int32) *GetSymbolicFilesRequest {
	s.PageSize = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetStartTime(v int64) *GetSymbolicFilesRequest {
	s.StartTime = &v
	return s
}

func (s *GetSymbolicFilesRequest) SetUuid(v string) *GetSymbolicFilesRequest {
	s.Uuid = &v
	return s
}

func (s *GetSymbolicFilesRequest) Validate() error {
	return dara.Validate(s)
}
