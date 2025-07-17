// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRumUploadFileUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *CreateRumUploadFileUrlRequest
	GetAppName() *string
	SetContentType(v string) *CreateRumUploadFileUrlRequest
	GetContentType() *string
	SetFileName(v string) *CreateRumUploadFileUrlRequest
	GetFileName() *string
	SetPid(v string) *CreateRumUploadFileUrlRequest
	GetPid() *string
	SetRegionId(v string) *CreateRumUploadFileUrlRequest
	GetRegionId() *string
	SetSourcemapType(v string) *CreateRumUploadFileUrlRequest
	GetSourcemapType() *string
	SetUuid(v string) *CreateRumUploadFileUrlRequest
	GetUuid() *string
	SetVersionId(v string) *CreateRumUploadFileUrlRequest
	GetVersionId() *string
}

type CreateRumUploadFileUrlRequest struct {
	// The application name.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// The type of the file. You can set this parameter to "application/zip", "text/plain", or an empty string.
	//
	// example:
	//
	// text/plain
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.js.map
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The process ID (PID) of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// iioe7jcnuk@582846f37******
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The file type. Valid values: source-map: SourceMap files. mapping: symbol table files for Android. dsym: dSYM files for iOS.
	//
	// example:
	//
	// source-map
	SourcemapType *string `json:"SourcemapType,omitempty" xml:"SourcemapType,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 125bdb39-a415-4503-bd96-e293925fc64c
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
	// The version number of the file.
	//
	// example:
	//
	// 1.0.0
	VersionId *string `json:"VersionId,omitempty" xml:"VersionId,omitempty"`
}

func (s CreateRumUploadFileUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRumUploadFileUrlRequest) GoString() string {
	return s.String()
}

func (s *CreateRumUploadFileUrlRequest) GetAppName() *string {
	return s.AppName
}

func (s *CreateRumUploadFileUrlRequest) GetContentType() *string {
	return s.ContentType
}

func (s *CreateRumUploadFileUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateRumUploadFileUrlRequest) GetPid() *string {
	return s.Pid
}

func (s *CreateRumUploadFileUrlRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateRumUploadFileUrlRequest) GetSourcemapType() *string {
	return s.SourcemapType
}

func (s *CreateRumUploadFileUrlRequest) GetUuid() *string {
	return s.Uuid
}

func (s *CreateRumUploadFileUrlRequest) GetVersionId() *string {
	return s.VersionId
}

func (s *CreateRumUploadFileUrlRequest) SetAppName(v string) *CreateRumUploadFileUrlRequest {
	s.AppName = &v
	return s
}

func (s *CreateRumUploadFileUrlRequest) SetContentType(v string) *CreateRumUploadFileUrlRequest {
	s.ContentType = &v
	return s
}

func (s *CreateRumUploadFileUrlRequest) SetFileName(v string) *CreateRumUploadFileUrlRequest {
	s.FileName = &v
	return s
}

func (s *CreateRumUploadFileUrlRequest) SetPid(v string) *CreateRumUploadFileUrlRequest {
	s.Pid = &v
	return s
}

func (s *CreateRumUploadFileUrlRequest) SetRegionId(v string) *CreateRumUploadFileUrlRequest {
	s.RegionId = &v
	return s
}

func (s *CreateRumUploadFileUrlRequest) SetSourcemapType(v string) *CreateRumUploadFileUrlRequest {
	s.SourcemapType = &v
	return s
}

func (s *CreateRumUploadFileUrlRequest) SetUuid(v string) *CreateRumUploadFileUrlRequest {
	s.Uuid = &v
	return s
}

func (s *CreateRumUploadFileUrlRequest) SetVersionId(v string) *CreateRumUploadFileUrlRequest {
	s.VersionId = &v
	return s
}

func (s *CreateRumUploadFileUrlRequest) Validate() error {
	return dara.Validate(s)
}
