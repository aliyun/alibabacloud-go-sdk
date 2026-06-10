// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListAgentsResponseBody
	GetRequestId() *string
	SetCode(v string) *ListAgentsResponseBody
	GetCode() *string
	SetData(v []*ListAgentsResponseBodyData) *ListAgentsResponseBody
	GetData() []*ListAgentsResponseBodyData
	SetMessage(v string) *ListAgentsResponseBody
	GetMessage() *string
	SetTotal(v int64) *ListAgentsResponseBody
	GetTotal() *int64
}

type ListAgentsResponseBody struct {
	// Request ID, which can be used for end-to-end Diagnosis
	//
	// example:
	//
	// 66EAED72-542B-583B-BCED-64433DC27AD7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Status code
	//
	// - `code == Success` indicates successful authorization;
	//
	// - Other status codes indicate authorization failure. When authorization fails, view the `message` field to obtain detailed error message;
	//
	// example:
	//
	// Success
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned Data
	Data []*ListAgentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// Error message
	//
	// - If `code == Success`, this field is empty.
	//
	// - Otherwise, this field contains the request error message.
	//
	// example:
	//
	// SysomOpenAPIAssumeRoleException: EntityNotExist.Role The role not exists: acs:ram::xxxxx:role/aliyunserviceroleforsysom
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentsResponseBody) GetData() []*ListAgentsResponseBodyData {
	return s.Data
}

func (s *ListAgentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListAgentsResponseBody) SetRequestId(v string) *ListAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentsResponseBody) SetCode(v string) *ListAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentsResponseBody) SetData(v []*ListAgentsResponseBodyData) *ListAgentsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentsResponseBody) SetMessage(v string) *ListAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentsResponseBody) SetTotal(v int64) *ListAgentsResponseBody {
	s.Total = &v
	return s
}

func (s *ListAgentsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentsResponseBodyData struct {
	// Widget creation time
	//
	// example:
	//
	// 2024-09-14T20:46:08
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// Widget description
	//
	// example:
	//
	// SysOM Agent
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Widget ID
	//
	// example:
	//
	// 74a86327-3170-412c-8e67-da3389ec56a9
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// Widget name
	//
	// example:
	//
	// SysOM Agent
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Supported architecture (multiple architectures separated by commas)
	//
	// example:
	//
	// x86
	SupportArch *string `json:"support_arch,omitempty" xml:"support_arch,omitempty"`
	// Widget type
	//
	// - Control: control-type widget
	//
	// - AI: AI widget
	//
	// example:
	//
	// Control
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Widget Update Time
	//
	// example:
	//
	// 2024-09-14T20:46:08
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Widget Version List
	Versions []*ListAgentsResponseBodyDataVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s ListAgentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListAgentsResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *ListAgentsResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListAgentsResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ListAgentsResponseBodyData) GetSupportArch() *string {
	return s.SupportArch
}

func (s *ListAgentsResponseBodyData) GetType() *string {
	return s.Type
}

func (s *ListAgentsResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListAgentsResponseBodyData) GetVersions() []*ListAgentsResponseBodyDataVersions {
	return s.Versions
}

func (s *ListAgentsResponseBodyData) SetCreatedAt(v string) *ListAgentsResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetDescription(v string) *ListAgentsResponseBodyData {
	s.Description = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetId(v string) *ListAgentsResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetName(v string) *ListAgentsResponseBodyData {
	s.Name = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetSupportArch(v string) *ListAgentsResponseBodyData {
	s.SupportArch = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetType(v string) *ListAgentsResponseBodyData {
	s.Type = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetUpdatedAt(v string) *ListAgentsResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetVersions(v []*ListAgentsResponseBodyDataVersions) *ListAgentsResponseBodyData {
	s.Versions = v
	return s
}

func (s *ListAgentsResponseBodyData) Validate() error {
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

type ListAgentsResponseBodyDataVersions struct {
	// Widget version creation time
	//
	// example:
	//
	// 2024-09-14T20:46:08
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The widget\\"s install script
	//
	// example:
	//
	// sysom.sh install
	InstallScript *string `json:"install_script,omitempty" xml:"install_script,omitempty"`
	// Widget uninstall script
	//
	// example:
	//
	// sysom.sh uninstall
	UninstallScript *string `json:"uninstall_script,omitempty" xml:"uninstall_script,omitempty"`
	// Widget version update time
	//
	// example:
	//
	// 2024-09-14T20:46:08
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// Widget upgrade script
	//
	// example:
	//
	// sysom.sh upgrade
	UpgradeScript *string `json:"upgrade_script,omitempty" xml:"upgrade_script,omitempty"`
	// Widget version number
	//
	// example:
	//
	// 3.4.0-1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListAgentsResponseBodyDataVersions) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyDataVersions) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *ListAgentsResponseBodyDataVersions) GetInstallScript() *string {
	return s.InstallScript
}

func (s *ListAgentsResponseBodyDataVersions) GetUninstallScript() *string {
	return s.UninstallScript
}

func (s *ListAgentsResponseBodyDataVersions) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *ListAgentsResponseBodyDataVersions) GetUpgradeScript() *string {
	return s.UpgradeScript
}

func (s *ListAgentsResponseBodyDataVersions) GetVersion() *string {
	return s.Version
}

func (s *ListAgentsResponseBodyDataVersions) SetCreatedAt(v string) *ListAgentsResponseBodyDataVersions {
	s.CreatedAt = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetInstallScript(v string) *ListAgentsResponseBodyDataVersions {
	s.InstallScript = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetUninstallScript(v string) *ListAgentsResponseBodyDataVersions {
	s.UninstallScript = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetUpdatedAt(v string) *ListAgentsResponseBodyDataVersions {
	s.UpdatedAt = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetUpgradeScript(v string) *ListAgentsResponseBodyDataVersions {
	s.UpgradeScript = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) SetVersion(v string) *ListAgentsResponseBodyDataVersions {
	s.Version = &v
	return s
}

func (s *ListAgentsResponseBodyDataVersions) Validate() error {
	return dara.Validate(s)
}
