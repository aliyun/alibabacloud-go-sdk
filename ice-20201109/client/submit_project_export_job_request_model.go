// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitProjectExportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExportType(v string) *SubmitProjectExportJobRequest
	GetExportType() *string
	SetOutputMediaConfig(v string) *SubmitProjectExportJobRequest
	GetOutputMediaConfig() *string
	SetProjectId(v string) *SubmitProjectExportJobRequest
	GetProjectId() *string
	SetTimeline(v string) *SubmitProjectExportJobRequest
	GetTimeline() *string
	SetUserData(v string) *SubmitProjectExportJobRequest
	GetUserData() *string
}

type SubmitProjectExportJobRequest struct {
	// example:
	//
	// BaseTimeline
	ExportType *string `json:"ExportType,omitempty" xml:"ExportType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	// 	"Bucket": "example-bucket",
	//
	//     "Prefix": "example_prefix"
	//
	// }
	OutputMediaConfig *string `json:"OutputMediaConfig,omitempty" xml:"OutputMediaConfig,omitempty"`
	// example:
	//
	// *****67ae06542b9b93e0d1c387*****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// {"VideoTracks":[{"VideoTrackClips":[{"MediaId":"****4d7cf14dc7b83b0e801c****"},{"MediaId":"****4d7cf14dc7b83b0e801c****"}]}]}
	Timeline *string `json:"Timeline,omitempty" xml:"Timeline,omitempty"`
	// example:
	//
	// {"NotifyAddress":"http://xx.xx.xxx","Key":"Valuexxx"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s SubmitProjectExportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitProjectExportJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitProjectExportJobRequest) GetExportType() *string {
	return s.ExportType
}

func (s *SubmitProjectExportJobRequest) GetOutputMediaConfig() *string {
	return s.OutputMediaConfig
}

func (s *SubmitProjectExportJobRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SubmitProjectExportJobRequest) GetTimeline() *string {
	return s.Timeline
}

func (s *SubmitProjectExportJobRequest) GetUserData() *string {
	return s.UserData
}

func (s *SubmitProjectExportJobRequest) SetExportType(v string) *SubmitProjectExportJobRequest {
	s.ExportType = &v
	return s
}

func (s *SubmitProjectExportJobRequest) SetOutputMediaConfig(v string) *SubmitProjectExportJobRequest {
	s.OutputMediaConfig = &v
	return s
}

func (s *SubmitProjectExportJobRequest) SetProjectId(v string) *SubmitProjectExportJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitProjectExportJobRequest) SetTimeline(v string) *SubmitProjectExportJobRequest {
	s.Timeline = &v
	return s
}

func (s *SubmitProjectExportJobRequest) SetUserData(v string) *SubmitProjectExportJobRequest {
	s.UserData = &v
	return s
}

func (s *SubmitProjectExportJobRequest) Validate() error {
	return dara.Validate(s)
}
