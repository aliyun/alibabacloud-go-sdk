// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTranscodeTaskStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApp(v string) *GetTranscodeTaskStatusRequest
	GetApp() *string
	SetPushDomain(v string) *GetTranscodeTaskStatusRequest
	GetPushDomain() *string
	SetSecurityToken(v string) *GetTranscodeTaskStatusRequest
	GetSecurityToken() *string
	SetStreamName(v string) *GetTranscodeTaskStatusRequest
	GetStreamName() *string
	SetTranscodingTemplate(v string) *GetTranscodeTaskStatusRequest
	GetTranscodingTemplate() *string
}

type GetTranscodeTaskStatusRequest struct {
	App                 *string `json:"App,omitempty" xml:"App,omitempty"`
	PushDomain          *string `json:"PushDomain,omitempty" xml:"PushDomain,omitempty"`
	SecurityToken       *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	StreamName          *string `json:"StreamName,omitempty" xml:"StreamName,omitempty"`
	TranscodingTemplate *string `json:"TranscodingTemplate,omitempty" xml:"TranscodingTemplate,omitempty"`
}

func (s GetTranscodeTaskStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTranscodeTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTranscodeTaskStatusRequest) GetApp() *string {
	return s.App
}

func (s *GetTranscodeTaskStatusRequest) GetPushDomain() *string {
	return s.PushDomain
}

func (s *GetTranscodeTaskStatusRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetTranscodeTaskStatusRequest) GetStreamName() *string {
	return s.StreamName
}

func (s *GetTranscodeTaskStatusRequest) GetTranscodingTemplate() *string {
	return s.TranscodingTemplate
}

func (s *GetTranscodeTaskStatusRequest) SetApp(v string) *GetTranscodeTaskStatusRequest {
	s.App = &v
	return s
}

func (s *GetTranscodeTaskStatusRequest) SetPushDomain(v string) *GetTranscodeTaskStatusRequest {
	s.PushDomain = &v
	return s
}

func (s *GetTranscodeTaskStatusRequest) SetSecurityToken(v string) *GetTranscodeTaskStatusRequest {
	s.SecurityToken = &v
	return s
}

func (s *GetTranscodeTaskStatusRequest) SetStreamName(v string) *GetTranscodeTaskStatusRequest {
	s.StreamName = &v
	return s
}

func (s *GetTranscodeTaskStatusRequest) SetTranscodingTemplate(v string) *GetTranscodeTaskStatusRequest {
	s.TranscodingTemplate = &v
	return s
}

func (s *GetTranscodeTaskStatusRequest) Validate() error {
	return dara.Validate(s)
}
