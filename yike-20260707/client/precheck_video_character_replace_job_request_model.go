// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckVideoCharacterReplaceJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PrecheckVideoCharacterReplaceJobRequest
	GetAppId() *string
	SetAppParams(v string) *PrecheckVideoCharacterReplaceJobRequest
	GetAppParams() *string
}

type PrecheckVideoCharacterReplaceJobRequest struct {
	// The ID of the video character replacement application. Use a configured application.
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The application input parameters, which are provided as a JSON-serialized string. The field names and material requirements are subject to the input definition of the corresponding AppId.
	//
	// example:
	//
	// {"LoadImage.1.TargetImage":"1a7852b032c371f0b64fe6f6c74b****"}
	AppParams *string `json:"AppParams,omitempty" xml:"AppParams,omitempty"`
}

func (s PrecheckVideoCharacterReplaceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s PrecheckVideoCharacterReplaceJobRequest) GoString() string {
	return s.String()
}

func (s *PrecheckVideoCharacterReplaceJobRequest) GetAppId() *string {
	return s.AppId
}

func (s *PrecheckVideoCharacterReplaceJobRequest) GetAppParams() *string {
	return s.AppParams
}

func (s *PrecheckVideoCharacterReplaceJobRequest) SetAppId(v string) *PrecheckVideoCharacterReplaceJobRequest {
	s.AppId = &v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobRequest) SetAppParams(v string) *PrecheckVideoCharacterReplaceJobRequest {
	s.AppParams = &v
	return s
}

func (s *PrecheckVideoCharacterReplaceJobRequest) Validate() error {
	return dara.Validate(s)
}
