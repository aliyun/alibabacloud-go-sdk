// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSymbolicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *SubmitSymbolicRequest
	GetAppKey() *int64
	SetAppVersion(v string) *SubmitSymbolicRequest
	GetAppVersion() *string
	SetBuildId(v string) *SubmitSymbolicRequest
	GetBuildId() *string
	SetFileName(v string) *SubmitSymbolicRequest
	GetFileName() *string
	SetFilePath(v string) *SubmitSymbolicRequest
	GetFilePath() *string
	SetFileType(v string) *SubmitSymbolicRequest
	GetFileType() *string
	SetOs(v string) *SubmitSymbolicRequest
	GetOs() *string
	SetUuid(v string) *SubmitSymbolicRequest
	GetUuid() *string
}

type SubmitSymbolicRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 24781204
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// 1.0
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	// example:
	//
	// ab6b81d800968f2
	BuildId *string `json:"BuildId,omitempty" xml:"BuildId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// app_so.zip
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 24781204@android/1743506690915-app_so.zip
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
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
	// uuid
	//
	// example:
	//
	// abcf4a4b-158c-4a0b-b81c-262785d84c4f
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s SubmitSymbolicRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitSymbolicRequest) GoString() string {
	return s.String()
}

func (s *SubmitSymbolicRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *SubmitSymbolicRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *SubmitSymbolicRequest) GetBuildId() *string {
	return s.BuildId
}

func (s *SubmitSymbolicRequest) GetFileName() *string {
	return s.FileName
}

func (s *SubmitSymbolicRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *SubmitSymbolicRequest) GetFileType() *string {
	return s.FileType
}

func (s *SubmitSymbolicRequest) GetOs() *string {
	return s.Os
}

func (s *SubmitSymbolicRequest) GetUuid() *string {
	return s.Uuid
}

func (s *SubmitSymbolicRequest) SetAppKey(v int64) *SubmitSymbolicRequest {
	s.AppKey = &v
	return s
}

func (s *SubmitSymbolicRequest) SetAppVersion(v string) *SubmitSymbolicRequest {
	s.AppVersion = &v
	return s
}

func (s *SubmitSymbolicRequest) SetBuildId(v string) *SubmitSymbolicRequest {
	s.BuildId = &v
	return s
}

func (s *SubmitSymbolicRequest) SetFileName(v string) *SubmitSymbolicRequest {
	s.FileName = &v
	return s
}

func (s *SubmitSymbolicRequest) SetFilePath(v string) *SubmitSymbolicRequest {
	s.FilePath = &v
	return s
}

func (s *SubmitSymbolicRequest) SetFileType(v string) *SubmitSymbolicRequest {
	s.FileType = &v
	return s
}

func (s *SubmitSymbolicRequest) SetOs(v string) *SubmitSymbolicRequest {
	s.Os = &v
	return s
}

func (s *SubmitSymbolicRequest) SetUuid(v string) *SubmitSymbolicRequest {
	s.Uuid = &v
	return s
}

func (s *SubmitSymbolicRequest) Validate() error {
	return dara.Validate(s)
}
