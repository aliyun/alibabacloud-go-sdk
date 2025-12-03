// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateProjectResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *CreateProjectResponseBody
	GetErrorMsg() *string
	SetProject(v *CreateProjectResponseBodyProject) *CreateProjectResponseBody
	GetProject() *CreateProjectResponseBodyProject
	SetRequestId(v string) *CreateProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateProjectResponseBody
	GetSuccess() *bool
}

type CreateProjectResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string                           `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	Project  *CreateProjectResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBody) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateProjectResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *CreateProjectResponseBody) GetProject() *CreateProjectResponseBodyProject {
	return s.Project
}

func (s *CreateProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateProjectResponseBody) SetErrorCode(v string) *CreateProjectResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateProjectResponseBody) SetErrorMsg(v string) *CreateProjectResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *CreateProjectResponseBody) SetProject(v *CreateProjectResponseBodyProject) *CreateProjectResponseBody {
	s.Project = v
	return s
}

func (s *CreateProjectResponseBody) SetRequestId(v string) *CreateProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateProjectResponseBody) SetSuccess(v bool) *CreateProjectResponseBody {
	s.Success = &v
	return s
}

func (s *CreateProjectResponseBody) Validate() error {
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateProjectResponseBodyProject struct {
	// example:
	//
	// Project
	CategoryIdentifier *string `json:"categoryIdentifier,omitempty" xml:"categoryIdentifier,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// OJAY
	CustomCode *string `json:"customCode,omitempty" xml:"customCode,omitempty"`
	// example:
	//
	// xxx
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 1623916393000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1623916393000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// example:
	//
	// https://xxxxxx.png
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// NORMAL
	LogicalStatus *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// demo示例项目
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 5e70xxxxxxcd000xxxxe96
	OrganizationIdentifier *string `json:"organizationIdentifier,omitempty" xml:"organizationIdentifier,omitempty"`
	// example:
	//
	// public
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// example:
	//
	// 8a40xxxxxxxxxxxxx64
	StatusIdentifier *string `json:"statusIdentifier,omitempty" xml:"statusIdentifier,omitempty"`
	// example:
	//
	// 1
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	// example:
	//
	// CustomProject
	TypeIdentifier *string `json:"typeIdentifier,omitempty" xml:"typeIdentifier,omitempty"`
}

func (s CreateProjectResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s CreateProjectResponseBodyProject) GoString() string {
	return s.String()
}

func (s *CreateProjectResponseBodyProject) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *CreateProjectResponseBodyProject) GetCreator() *string {
	return s.Creator
}

func (s *CreateProjectResponseBodyProject) GetCustomCode() *string {
	return s.CustomCode
}

func (s *CreateProjectResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *CreateProjectResponseBodyProject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateProjectResponseBodyProject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *CreateProjectResponseBodyProject) GetIcon() *string {
	return s.Icon
}

func (s *CreateProjectResponseBodyProject) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateProjectResponseBodyProject) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *CreateProjectResponseBodyProject) GetModifier() *string {
	return s.Modifier
}

func (s *CreateProjectResponseBodyProject) GetName() *string {
	return s.Name
}

func (s *CreateProjectResponseBodyProject) GetOrganizationIdentifier() *string {
	return s.OrganizationIdentifier
}

func (s *CreateProjectResponseBodyProject) GetScope() *string {
	return s.Scope
}

func (s *CreateProjectResponseBodyProject) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *CreateProjectResponseBodyProject) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *CreateProjectResponseBodyProject) GetTypeIdentifier() *string {
	return s.TypeIdentifier
}

func (s *CreateProjectResponseBodyProject) SetCategoryIdentifier(v string) *CreateProjectResponseBodyProject {
	s.CategoryIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetCreator(v string) *CreateProjectResponseBodyProject {
	s.Creator = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetCustomCode(v string) *CreateProjectResponseBodyProject {
	s.CustomCode = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetDescription(v string) *CreateProjectResponseBodyProject {
	s.Description = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetGmtCreate(v int64) *CreateProjectResponseBodyProject {
	s.GmtCreate = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetGmtModified(v int64) *CreateProjectResponseBodyProject {
	s.GmtModified = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetIcon(v string) *CreateProjectResponseBodyProject {
	s.Icon = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetIdentifier(v string) *CreateProjectResponseBodyProject {
	s.Identifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetLogicalStatus(v string) *CreateProjectResponseBodyProject {
	s.LogicalStatus = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetModifier(v string) *CreateProjectResponseBodyProject {
	s.Modifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetName(v string) *CreateProjectResponseBodyProject {
	s.Name = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetOrganizationIdentifier(v string) *CreateProjectResponseBodyProject {
	s.OrganizationIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetScope(v string) *CreateProjectResponseBodyProject {
	s.Scope = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetStatusIdentifier(v string) *CreateProjectResponseBodyProject {
	s.StatusIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetStatusStageIdentifier(v string) *CreateProjectResponseBodyProject {
	s.StatusStageIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) SetTypeIdentifier(v string) *CreateProjectResponseBodyProject {
	s.TypeIdentifier = &v
	return s
}

func (s *CreateProjectResponseBodyProject) Validate() error {
	return dara.Validate(s)
}
