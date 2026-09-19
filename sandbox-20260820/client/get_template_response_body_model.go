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
	Code            *string                      `json:"code,omitempty" xml:"code,omitempty"`
	CreatedTime     *string                      `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	Message         *string                      `json:"message,omitempty" xml:"message,omitempty"`
	Name            *string                      `json:"name,omitempty" xml:"name,omitempty"`
	RequestId       *string                      `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResourceGroupID *string                      `json:"resourceGroupID,omitempty" xml:"resourceGroupID,omitempty"`
	RuntimeConfig   *PublicTemplateRuntimeConfig `json:"runtimeConfig,omitempty" xml:"runtimeConfig,omitempty"`
	Status          *PublicTemplateStatus        `json:"status,omitempty" xml:"status,omitempty"`
	TeamID          *string                      `json:"teamID,omitempty" xml:"teamID,omitempty"`
	TeamName        *string                      `json:"teamName,omitempty" xml:"teamName,omitempty"`
	TemplateID      *string                      `json:"templateID,omitempty" xml:"templateID,omitempty"`
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
