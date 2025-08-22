// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumUploadFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppType(v string) *GetRumUploadFilesRequest
	GetAppType() *string
	SetFileName(v string) *GetRumUploadFilesRequest
	GetFileName() *string
	SetNextToken(v string) *GetRumUploadFilesRequest
	GetNextToken() *string
	SetPageSize(v int32) *GetRumUploadFilesRequest
	GetPageSize() *int32
	SetPid(v string) *GetRumUploadFilesRequest
	GetPid() *string
	SetRegionId(v string) *GetRumUploadFilesRequest
	GetRegionId() *string
	SetVersionId(v string) *GetRumUploadFilesRequest
	GetVersionId() *string
}

type GetRumUploadFilesRequest struct {
	// The file type. Valid values: source-map: SourceMap files. mapping: symbol table files for Android. dsym: dSYM files for iOS.
	//
	// example:
	//
	// source-map
	AppType   *string `json:"AppType,omitempty" xml:"AppType,omitempty"`
	FileName  *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	PageSize  *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The process ID (PID) of the application.
	//
	// example:
	//
	// aoxxxxxly@741623b4e91****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The version number of the files. If you do not specify this parameter, all versions of the files are returned by default.
	//
	// example:
	//
	// 1.0.0
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s GetRumUploadFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumUploadFilesRequest) GoString() string {
	return s.String()
}

func (s *GetRumUploadFilesRequest) GetAppType() *string {
	return s.AppType
}

func (s *GetRumUploadFilesRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetRumUploadFilesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *GetRumUploadFilesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetRumUploadFilesRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRumUploadFilesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumUploadFilesRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *GetRumUploadFilesRequest) SetAppType(v string) *GetRumUploadFilesRequest {
	s.AppType = &v
	return s
}

func (s *GetRumUploadFilesRequest) SetFileName(v string) *GetRumUploadFilesRequest {
	s.FileName = &v
	return s
}

func (s *GetRumUploadFilesRequest) SetNextToken(v string) *GetRumUploadFilesRequest {
	s.NextToken = &v
	return s
}

func (s *GetRumUploadFilesRequest) SetPageSize(v int32) *GetRumUploadFilesRequest {
	s.PageSize = &v
	return s
}

func (s *GetRumUploadFilesRequest) SetPid(v string) *GetRumUploadFilesRequest {
	s.Pid = &v
	return s
}

func (s *GetRumUploadFilesRequest) SetRegionId(v string) *GetRumUploadFilesRequest {
	s.RegionId = &v
	return s
}

func (s *GetRumUploadFilesRequest) SetVersionId(v string) *GetRumUploadFilesRequest {
	s.VersionId = &v
	return s
}

func (s *GetRumUploadFilesRequest) Validate() error {
	return dara.Validate(s)
}
