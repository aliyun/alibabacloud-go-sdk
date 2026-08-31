// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTemplateResponseBody
	GetCode() *string
	SetCreatedTime(v string) *GetTemplateResponseBody
	GetCreatedTime() *string
	SetMessage(v string) *GetTemplateResponseBody
	GetMessage() *string
	SetName(v string) *GetTemplateResponseBody
	GetName() *string
	SetRequestId(v string) *GetTemplateResponseBody
	GetRequestId() *string
	SetResourceGroupID(v string) *GetTemplateResponseBody
	GetResourceGroupID() *string
	SetRuntimeConfig(v *PublicTemplateRuntimeConfig) *GetTemplateResponseBody
	GetRuntimeConfig() *PublicTemplateRuntimeConfig
	SetStatus(v *PublicTemplateStatus) *GetTemplateResponseBody
	GetStatus() *PublicTemplateStatus
	SetTeamID(v string) *GetTemplateResponseBody
	GetTeamID() *string
	SetTeamName(v string) *GetTemplateResponseBody
	GetTeamName() *string
	SetTemplateID(v string) *GetTemplateResponseBody
	GetTemplateID() *string
}

type GetTemplateResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The time when the template was created.
	//
	// example:
	//
	// 2026-08-31T12:00:00Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The template name.
	//
	// example:
	//
	// my-template
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Id of the request
	//
	// example:
	//
	// B5AD8B54-4358-5F5B-ACAA-52F2016459C6
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-acfmwxqyrgwabcd
	ResourceGroupID *string `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	// The runtime configuration of the template.
	RuntimeConfig *PublicTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
	// The template status.
	Status *PublicTemplateStatus `json:"status,omitempty" xml:"status,omitempty"`
	// The unique identifier of the team.
	//
	// example:
	//
	// 88a4c762-b0ce-4661-9413-578b2309e60f
	TeamID *string `json:"teamID,omitempty" xml:"teamID,omitempty"`
	// The team name.
	//
	// example:
	//
	// codeclaw-localenv
	TeamName *string `json:"teamName,omitempty" xml:"teamName,omitempty"`
	// The unique identifier of the template.
	//
	// example:
	//
	// tpl-9f3a2b7c8d1e4f5a6b0c7d8e9f1a2b3c
	TemplateID *string `json:"templateID,omitempty" xml:"templateID,omitempty"`
}

func (s GetTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTemplateResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *GetTemplateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTemplateResponseBody) GetName() *string {
	return s.Name
}

func (s *GetTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateResponseBody) GetResourceGroupID() *string {
	return s.ResourceGroupID
}

func (s *GetTemplateResponseBody) GetRuntimeConfig() *PublicTemplateRuntimeConfig {
	return s.RuntimeConfig
}

func (s *GetTemplateResponseBody) GetStatus() *PublicTemplateStatus {
	return s.Status
}

func (s *GetTemplateResponseBody) GetTeamID() *string {
	return s.TeamID
}

func (s *GetTemplateResponseBody) GetTeamName() *string {
	return s.TeamName
}

func (s *GetTemplateResponseBody) GetTemplateID() *string {
	return s.TemplateID
}

func (s *GetTemplateResponseBody) SetCode(v string) *GetTemplateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTemplateResponseBody) SetCreatedTime(v string) *GetTemplateResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetTemplateResponseBody) SetMessage(v string) *GetTemplateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTemplateResponseBody) SetName(v string) *GetTemplateResponseBody {
	s.Name = &v
	return s
}

func (s *GetTemplateResponseBody) SetRequestId(v string) *GetTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateResponseBody) SetResourceGroupID(v string) *GetTemplateResponseBody {
	s.ResourceGroupID = &v
	return s
}

func (s *GetTemplateResponseBody) SetRuntimeConfig(v *PublicTemplateRuntimeConfig) *GetTemplateResponseBody {
	s.RuntimeConfig = v
	return s
}

func (s *GetTemplateResponseBody) SetStatus(v *PublicTemplateStatus) *GetTemplateResponseBody {
	s.Status = v
	return s
}

func (s *GetTemplateResponseBody) SetTeamID(v string) *GetTemplateResponseBody {
	s.TeamID = &v
	return s
}

func (s *GetTemplateResponseBody) SetTeamName(v string) *GetTemplateResponseBody {
	s.TeamName = &v
	return s
}

func (s *GetTemplateResponseBody) SetTemplateID(v string) *GetTemplateResponseBody {
	s.TemplateID = &v
	return s
}

func (s *GetTemplateResponseBody) Validate() error {
	if s.RuntimeConfig != nil {
		if err := s.RuntimeConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Status != nil {
		if err := s.Status.Validate(); err != nil {
			return err
		}
	}
	return nil
}
