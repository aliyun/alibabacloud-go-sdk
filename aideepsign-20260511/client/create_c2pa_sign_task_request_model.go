// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateC2paSignTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileUrl(v string) *CreateC2paSignTaskRequest
	GetFileUrl() *string
	SetSoftwareAgent(v string) *CreateC2paSignTaskRequest
	GetSoftwareAgent() *string
	SetSourceType(v string) *CreateC2paSignTaskRequest
	GetSourceType() *string
}

type CreateC2paSignTaskRequest struct {
	// The absolute HTTP or HTTPS URL of the file to be signed. The URL can be up to 4,096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://example.com/images/photo.png?token=xxx
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// Not supported in the current version.
	//
	// example:
	//
	// Alibaba Cloud Qwen ImageGen
	SoftwareAgent *string `json:"SoftwareAgent,omitempty" xml:"SoftwareAgent,omitempty"`
	// Not supported in the current version.
	//
	// example:
	//
	// trainedAlgorithmicMedia
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s CreateC2paSignTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateC2paSignTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateC2paSignTaskRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *CreateC2paSignTaskRequest) GetSoftwareAgent() *string {
	return s.SoftwareAgent
}

func (s *CreateC2paSignTaskRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateC2paSignTaskRequest) SetFileUrl(v string) *CreateC2paSignTaskRequest {
	s.FileUrl = &v
	return s
}

func (s *CreateC2paSignTaskRequest) SetSoftwareAgent(v string) *CreateC2paSignTaskRequest {
	s.SoftwareAgent = &v
	return s
}

func (s *CreateC2paSignTaskRequest) SetSourceType(v string) *CreateC2paSignTaskRequest {
	s.SourceType = &v
	return s
}

func (s *CreateC2paSignTaskRequest) Validate() error {
	return dara.Validate(s)
}
