// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitYikeProjectExportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportType(v string) *SubmitYikeProjectExportJobRequest
	GetExportType() *string
	SetProjectId(v string) *SubmitYikeProjectExportJobRequest
	GetProjectId() *string
	SetUserData(v string) *SubmitYikeProjectExportJobRequest
	GetUserData() *string
}

type SubmitYikeProjectExportJobRequest struct {
	// The export type of the editing project. Valid values:
	//
	// - PureAudio: pure audio (the export result returns a pure audio file and a subtitle file).
	//
	// example:
	//
	// PureAudio
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// The ID of the online editing project.
	//
	// example:
	//
	// 01a6adbbd181437eb5030d3d93e0182d
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The custom parameter in JSON string format. The callback result carries this parameter as-is (for example, newsKey).
	//
	// The system reserved field NotifyAddress specifies the callback URL. After the task is completed, a callback is sent. Example: {"NotifyAddress": "http://xxx.callback.url"}
	//
	// example:
	//
	// {\\"newsKey\\":\\"NEWS_20260420_001\\",\\"key1\\":\\"value1\\", \\"NotifyAddress\\":\\"https://cms.example.com/callback/video-task\\"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitYikeProjectExportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitYikeProjectExportJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitYikeProjectExportJobRequest) GetExportType() *string {
	return s.ExportType
}

func (s *SubmitYikeProjectExportJobRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SubmitYikeProjectExportJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitYikeProjectExportJobRequest) SetExportType(v string) *SubmitYikeProjectExportJobRequest {
	s.ExportType = &v
	return s
}

func (s *SubmitYikeProjectExportJobRequest) SetProjectId(v string) *SubmitYikeProjectExportJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitYikeProjectExportJobRequest) SetUserData(v string) *SubmitYikeProjectExportJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitYikeProjectExportJobRequest) Validate() error {
	return dara.Validate(s)
}
