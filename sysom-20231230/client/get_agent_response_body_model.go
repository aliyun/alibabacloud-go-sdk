// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAgentResponseBody
	GetRequestId() *string
	SetCode(v string) *GetAgentResponseBody
	GetCode() *string
	SetData(v *GetAgentResponseBodyData) *GetAgentResponseBody
	GetData() *GetAgentResponseBodyData
	SetMessage(v string) *GetAgentResponseBody
	GetMessage() *string
}

type GetAgentResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// Success
	Code *string                   `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetAgentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// SysomOpenAPIException: SysomOpenAPI.InvalidParameter Invalid params, should be json string or dict
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s GetAgentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentResponseBody) GetData() *GetAgentResponseBodyData {
	return s.Data
}

func (s *GetAgentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentResponseBody) SetRequestId(v string) *GetAgentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentResponseBody) SetCode(v string) *GetAgentResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentResponseBody) SetData(v *GetAgentResponseBodyData) *GetAgentResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentResponseBody) SetMessage(v string) *GetAgentResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentResponseBodyData struct {
	CreatedAt   *string                             `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description *string                             `json:"description,omitempty" xml:"description,omitempty"`
	Id          *string                             `json:"id,omitempty" xml:"id,omitempty"`
	Name        *string                             `json:"name,omitempty" xml:"name,omitempty"`
	SupportArch *string                             `json:"support_arch,omitempty" xml:"support_arch,omitempty"`
	Type        *string                             `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt   *string                             `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Versions    []*GetAgentResponseBodyDataVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetAgentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAgentResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgentResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetAgentResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAgentResponseBodyData) GetSupportArch() *string {
	return s.SupportArch
}

func (s *GetAgentResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetAgentResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetAgentResponseBodyData) GetVersions() []*GetAgentResponseBodyDataVersions {
	return s.Versions
}

func (s *GetAgentResponseBodyData) SetCreatedAt(v string) *GetAgentResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentResponseBodyData) SetDescription(v string) *GetAgentResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentResponseBodyData) SetId(v string) *GetAgentResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetAgentResponseBodyData) SetName(v string) *GetAgentResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAgentResponseBodyData) SetSupportArch(v string) *GetAgentResponseBodyData {
	s.SupportArch = &v
	return s
}

func (s *GetAgentResponseBodyData) SetType(v string) *GetAgentResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetAgentResponseBodyData) SetUpdatedAt(v string) *GetAgentResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetAgentResponseBodyData) SetVersions(v []*GetAgentResponseBodyDataVersions) *GetAgentResponseBodyData {
	s.Versions = v
	return s
}

func (s *GetAgentResponseBodyData) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentResponseBodyDataVersions struct {
	CreatedAt       *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	InstallScript   *string `json:"install_script,omitempty" xml:"install_script,omitempty"`
	UninstallScript *string `json:"uninstall_script,omitempty" xml:"uninstall_script,omitempty"`
	UpdatedAt       *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UpgradeScript   *string `json:"upgrade_script,omitempty" xml:"upgrade_script,omitempty"`
	Version         *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAgentResponseBodyDataVersions) String() string {
	return dara.Prettify(s)
}

func (s GetAgentResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *GetAgentResponseBodyDataVersions) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetAgentResponseBodyDataVersions) GetInstallScript() *string {
	return s.InstallScript
}

func (s *GetAgentResponseBodyDataVersions) GetUninstallScript() *string {
	return s.UninstallScript
}

func (s *GetAgentResponseBodyDataVersions) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetAgentResponseBodyDataVersions) GetUpgradeScript() *string {
	return s.UpgradeScript
}

func (s *GetAgentResponseBodyDataVersions) GetVersion() *string {
	return s.Version
}

func (s *GetAgentResponseBodyDataVersions) SetCreatedAt(v string) *GetAgentResponseBodyDataVersions {
	s.CreatedAt = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetInstallScript(v string) *GetAgentResponseBodyDataVersions {
	s.InstallScript = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetUninstallScript(v string) *GetAgentResponseBodyDataVersions {
	s.UninstallScript = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetUpdatedAt(v string) *GetAgentResponseBodyDataVersions {
	s.UpdatedAt = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetUpgradeScript(v string) *GetAgentResponseBodyDataVersions {
	s.UpgradeScript = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) SetVersion(v string) *GetAgentResponseBodyDataVersions {
	s.Version = &v
	return s
}

func (s *GetAgentResponseBodyDataVersions) Validate() error {
	return dara.Validate(s)
}
