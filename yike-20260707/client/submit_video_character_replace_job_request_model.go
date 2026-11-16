// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitVideoCharacterReplaceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SubmitVideoCharacterReplaceJobRequest
	GetAppId() *string
	SetAppParams(v string) *SubmitVideoCharacterReplaceJobRequest
	GetAppParams() *string
	SetFolderId(v string) *SubmitVideoCharacterReplaceJobRequest
	GetFolderId() *string
	SetProjectId(v string) *SubmitVideoCharacterReplaceJobRequest
	GetProjectId() *string
}

type SubmitVideoCharacterReplaceJobRequest struct {
	// The ID of the video character replacement application. Use a configured application.
	//
	// This parameter is required.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application input parameters as a JSON-serialized string. The field names and material requirements are subject to the input definition of the corresponding AppId.
	//
	// example:
	//
	// {"LoadImage.1.TargetImage":"1a7852b032c371f0b64fe6f6c74b****"}
	AppParams *string `json:"AppParams,omitempty" xml:"AppParams,omitempty"`
	// The folder ID.
	//
	// example:
	//
	// folder-u3ilwhoc36ux9a****
	FolderId *string `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// pd_061716****
	ProjectId *string `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s SubmitVideoCharacterReplaceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitVideoCharacterReplaceJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitVideoCharacterReplaceJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *SubmitVideoCharacterReplaceJobRequest) GetAppParams() *string {
	return s.AppParams
}

func (s *SubmitVideoCharacterReplaceJobRequest) GetFolderId() *string {
	return s.FolderId
}

func (s *SubmitVideoCharacterReplaceJobRequest) GetProjectId() *string {
	return s.ProjectId
}

func (s *SubmitVideoCharacterReplaceJobRequest) SetAppId(v string) *SubmitVideoCharacterReplaceJobRequest {
	s.AppId = &v
	return s
}

func (s *SubmitVideoCharacterReplaceJobRequest) SetAppParams(v string) *SubmitVideoCharacterReplaceJobRequest {
	s.AppParams = &v
	return s
}

func (s *SubmitVideoCharacterReplaceJobRequest) SetFolderId(v string) *SubmitVideoCharacterReplaceJobRequest {
	s.FolderId = &v
	return s
}

func (s *SubmitVideoCharacterReplaceJobRequest) SetProjectId(v string) *SubmitVideoCharacterReplaceJobRequest {
	s.ProjectId = &v
	return s
}

func (s *SubmitVideoCharacterReplaceJobRequest) Validate() error {
	return dara.Validate(s)
}
