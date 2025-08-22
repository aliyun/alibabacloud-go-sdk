// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateJobScriptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *UpdateJobScriptRequest
	GetAppName() *string
	SetClusterId(v string) *UpdateJobScriptRequest
	GetClusterId() *string
	SetJobId(v int64) *UpdateJobScriptRequest
	GetJobId() *int64
	SetScriptContent(v string) *UpdateJobScriptRequest
	GetScriptContent() *string
	SetVersionDescription(v string) *UpdateJobScriptRequest
	GetVersionDescription() *string
}

type UpdateJobScriptRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId         *int64  `json:"JobId,omitempty" xml:"JobId,omitempty"`
	ScriptContent *string `json:"ScriptContent,omitempty" xml:"ScriptContent,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// init version
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
}

func (s UpdateJobScriptRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateJobScriptRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobScriptRequest) GetAppName() *string {
	return s.AppName
}

func (s *UpdateJobScriptRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateJobScriptRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *UpdateJobScriptRequest) GetScriptContent() *string {
	return s.ScriptContent
}

func (s *UpdateJobScriptRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *UpdateJobScriptRequest) SetAppName(v string) *UpdateJobScriptRequest {
	s.AppName = &v
	return s
}

func (s *UpdateJobScriptRequest) SetClusterId(v string) *UpdateJobScriptRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateJobScriptRequest) SetJobId(v int64) *UpdateJobScriptRequest {
	s.JobId = &v
	return s
}

func (s *UpdateJobScriptRequest) SetScriptContent(v string) *UpdateJobScriptRequest {
	s.ScriptContent = &v
	return s
}

func (s *UpdateJobScriptRequest) SetVersionDescription(v string) *UpdateJobScriptRequest {
	s.VersionDescription = &v
	return s
}

func (s *UpdateJobScriptRequest) Validate() error {
	return dara.Validate(s)
}
