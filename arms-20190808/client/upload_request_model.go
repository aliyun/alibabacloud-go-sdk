// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEdition(v string) *UploadRequest
	GetEdition() *string
	SetFile(v string) *UploadRequest
	GetFile() *string
	SetFileName(v string) *UploadRequest
	GetFileName() *string
	SetPid(v string) *UploadRequest
	GetPid() *string
	SetRegionId(v string) *UploadRequest
	GetRegionId() *string
	SetVersion(v string) *UploadRequest
	GetVersion() *string
}

type UploadRequest struct {
	// The version of the SourceMap file.
	//
	// example:
	//
	// 0.0.0
	Edition *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	// The string of the SourceMap file.
	//
	// example:
	//
	// test file content
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// The name of the SourceMap file.
	//
	// This parameter is required.
	//
	// example:
	//
	// test.js.map
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The application ID.
	//
	// Log on to the **ARMS console**. In the left-side navigation pane, choose **Browser Monitoring*	- > **Browser Monitoring**. On the Browser Monitoring page, click the name of an application. The URL in the address bar contains the process ID (PID) of the application. The PID is indicated in the pid=xxx format. The PID is usually percent encoded as xxx%40xxx. You must modify this value to remove the percent encoding. For example, if the PID in the URL is eb4zdose6v%409781be0f44d\\*\\*\\*\\*, you must replace %40 with @ to obtain eb4zdose6v@9781be0f44d\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// b590lhguqs@8cc3f6354******
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The ID of the region to which the SourceMap file is uploaded.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// We recommend that you do not specify this parameter.
	//
	// example:
	//
	// null
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s UploadRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadRequest) GoString() string {
	return s.String()
}

func (s *UploadRequest) GetEdition() *string {
	return s.Edition
}

func (s *UploadRequest) GetFile() *string {
	return s.File
}

func (s *UploadRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadRequest) GetPid() *string {
	return s.Pid
}

func (s *UploadRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadRequest) GetVersion() *string {
	return s.Version
}

func (s *UploadRequest) SetEdition(v string) *UploadRequest {
	s.Edition = &v
	return s
}

func (s *UploadRequest) SetFile(v string) *UploadRequest {
	s.File = &v
	return s
}

func (s *UploadRequest) SetFileName(v string) *UploadRequest {
	s.FileName = &v
	return s
}

func (s *UploadRequest) SetPid(v string) *UploadRequest {
	s.Pid = &v
	return s
}

func (s *UploadRequest) SetRegionId(v string) *UploadRequest {
	s.RegionId = &v
	return s
}

func (s *UploadRequest) SetVersion(v string) *UploadRequest {
	s.Version = &v
	return s
}

func (s *UploadRequest) Validate() error {
	return dara.Validate(s)
}
