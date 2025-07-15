// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartTranscodeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *RestartTranscodeTaskRequest
	GetApp() *string
	SetPushDomain(v string) *RestartTranscodeTaskRequest
	GetPushDomain() *string
	SetSecurityToken(v string) *RestartTranscodeTaskRequest
	GetSecurityToken() *string
	SetStreamName(v string) *RestartTranscodeTaskRequest
	GetStreamName() *string
	SetTranscodingTemplate(v string) *RestartTranscodeTaskRequest
	GetTranscodingTemplate() *string
}

type RestartTranscodeTaskRequest struct {
	App                 *string `json:"App,omitempty" xml:"App,omitempty"`
	PushDomain          *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	SecurityToken       *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StreamName          *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	TranscodingTemplate *string `json:"TranscodingTemplate,omitempty" xml:"TranscodingTemplate,omitempty"`
}

func (s RestartTranscodeTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartTranscodeTaskRequest) GoString() string {
	return s.String()
}

func (s *RestartTranscodeTaskRequest) GetApp() *string {
	return s.App
}

func (s *RestartTranscodeTaskRequest) GetPushDomain() *string {
	return s.PushDomain
}

func (s *RestartTranscodeTaskRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *RestartTranscodeTaskRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *RestartTranscodeTaskRequest) GetTranscodingTemplate() *string {
	return s.TranscodingTemplate
}

func (s *RestartTranscodeTaskRequest) SetApp(v string) *RestartTranscodeTaskRequest {
	s.App = &v
	return s
}

func (s *RestartTranscodeTaskRequest) SetPushDomain(v string) *RestartTranscodeTaskRequest {
	s.PushDomain = &v
	return s
}

func (s *RestartTranscodeTaskRequest) SetSecurityToken(v string) *RestartTranscodeTaskRequest {
	s.SecurityToken = &v
	return s
}

func (s *RestartTranscodeTaskRequest) SetStreamName(v string) *RestartTranscodeTaskRequest {
	s.StreamName = &v
	return s
}

func (s *RestartTranscodeTaskRequest) SetTranscodingTemplate(v string) *RestartTranscodeTaskRequest {
	s.TranscodingTemplate = &v
	return s
}

func (s *RestartTranscodeTaskRequest) Validate() error {
	return dara.Validate(s)
}
