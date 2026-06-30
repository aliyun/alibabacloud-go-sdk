// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallPolarClawSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *InstallPolarClawSkillResponseBody
	GetApplicationId() *string
	SetCode(v int32) *InstallPolarClawSkillResponseBody
	GetCode() *int32
	SetMessage(v string) *InstallPolarClawSkillResponseBody
	GetMessage() *string
	SetOk(v bool) *InstallPolarClawSkillResponseBody
	GetOk() *bool
	SetRequestId(v string) *InstallPolarClawSkillResponseBody
	GetRequestId() *string
	SetSlug(v string) *InstallPolarClawSkillResponseBody
	GetSlug() *string
	SetTargetDir(v string) *InstallPolarClawSkillResponseBody
	GetTargetDir() *string
	SetVersion(v string) *InstallPolarClawSkillResponseBody
	GetVersion() *string
}

type InstallPolarClawSkillResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Indicates whether the operation was successful.
	//
	// example:
	//
	// true
	Ok *bool `json:"Ok,omitempty" xml:"Ok,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F45FFACC-xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The identifier of the installed Skill.
	//
	// example:
	//
	// alibacloud-rds-copilot
	Slug *string `json:"Slug,omitempty" xml:"Slug,omitempty"`
	// The installation directory.
	//
	// example:
	//
	// /home/node/.openclaw/skills/alibacloud-rds-copilot
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The version number of the installed Skill.
	//
	// example:
	//
	// 1.2.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s InstallPolarClawSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InstallPolarClawSkillResponseBody) GoString() string {
	return s.String()
}

func (s *InstallPolarClawSkillResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *InstallPolarClawSkillResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InstallPolarClawSkillResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InstallPolarClawSkillResponseBody) GetOk() *bool {
	return s.Ok
}

func (s *InstallPolarClawSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InstallPolarClawSkillResponseBody) GetSlug() *string {
	return s.Slug
}

func (s *InstallPolarClawSkillResponseBody) GetTargetDir() *string {
	return s.TargetDir
}

func (s *InstallPolarClawSkillResponseBody) GetVersion() *string {
	return s.Version
}

func (s *InstallPolarClawSkillResponseBody) SetApplicationId(v string) *InstallPolarClawSkillResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *InstallPolarClawSkillResponseBody) SetCode(v int32) *InstallPolarClawSkillResponseBody {
	s.Code = &v
	return s
}

func (s *InstallPolarClawSkillResponseBody) SetMessage(v string) *InstallPolarClawSkillResponseBody {
	s.Message = &v
	return s
}

func (s *InstallPolarClawSkillResponseBody) SetOk(v bool) *InstallPolarClawSkillResponseBody {
	s.Ok = &v
	return s
}

func (s *InstallPolarClawSkillResponseBody) SetRequestId(v string) *InstallPolarClawSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *InstallPolarClawSkillResponseBody) SetSlug(v string) *InstallPolarClawSkillResponseBody {
	s.Slug = &v
	return s
}

func (s *InstallPolarClawSkillResponseBody) SetTargetDir(v string) *InstallPolarClawSkillResponseBody {
	s.TargetDir = &v
	return s
}

func (s *InstallPolarClawSkillResponseBody) SetVersion(v string) *InstallPolarClawSkillResponseBody {
	s.Version = &v
	return s
}

func (s *InstallPolarClawSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
