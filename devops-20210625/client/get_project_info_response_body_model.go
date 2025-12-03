// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProjectInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetProjectInfoResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetProjectInfoResponseBody
	GetErrorMessage() *string
	SetProject(v *GetProjectInfoResponseBodyProject) *GetProjectInfoResponseBody
	GetProject() *GetProjectInfoResponseBodyProject
	SetRequestId(v string) *GetProjectInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProjectInfoResponseBody
	GetSuccess() *bool
}

type GetProjectInfoResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMessage *string                            `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Project      *GetProjectInfoResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetProjectInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProjectInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetProjectInfoResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetProjectInfoResponseBody) GetProject() *GetProjectInfoResponseBodyProject {
	return s.Project
}

func (s *GetProjectInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProjectInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProjectInfoResponseBody) SetErrorCode(v string) *GetProjectInfoResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetErrorMessage(v string) *GetProjectInfoResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetProject(v *GetProjectInfoResponseBodyProject) *GetProjectInfoResponseBody {
	s.Project = v
	return s
}

func (s *GetProjectInfoResponseBody) SetRequestId(v string) *GetProjectInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProjectInfoResponseBody) SetSuccess(v bool) *GetProjectInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetProjectInfoResponseBody) Validate() error {
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProjectInfoResponseBodyProject struct {
	// example:
	//
	// null
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
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
	// https://xxxxxx.png
	IconBig *string `json:"iconBig,omitempty" xml:"iconBig,omitempty"`
	// example:
	//
	// {"small":"https://xxxxxx.png","big":"https://img.yyyyyy.png"}
	IconGroup *string `json:"iconGroup,omitempty" xml:"iconGroup,omitempty"`
	// example:
	//
	// https://img.yyyyyy.png
	IconSmall *string `json:"iconSmall,omitempty" xml:"iconSmall,omitempty"`
	// example:
	//
	// null
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// e8b2xxxxxx2abdxxxxxxxx23
	IdentifierPath *string `json:"identifierPath,omitempty" xml:"identifierPath,omitempty"`
	// example:
	//
	// NORMAL
	LogicalStatus *string `json:"logicalStatus,omitempty" xml:"logicalStatus,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	Name     *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 5e70xxxxxxcd000xxxxe96
	OrganizationIdentifier *string `json:"organizationIdentifier,omitempty" xml:"organizationIdentifier,omitempty"`
	// example:
	//
	// null
	ParentIdentifier *string `json:"parentIdentifier,omitempty" xml:"parentIdentifier,omitempty"`
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
	// 8a4058a71159b68254......
	StatusStageIdentifier *string `json:"statusStageIdentifier,omitempty" xml:"statusStageIdentifier,omitempty"`
	// example:
	//
	// null
	SubType *string `json:"subType,omitempty" xml:"subType,omitempty"`
	// example:
	//
	// CustomProject
	TypeIdentifier *string `json:"typeIdentifier,omitempty" xml:"typeIdentifier,omitempty"`
}

func (s GetProjectInfoResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s GetProjectInfoResponseBodyProject) GoString() string {
	return s.String()
}

func (s *GetProjectInfoResponseBodyProject) GetCategory() *string {
	return s.Category
}

func (s *GetProjectInfoResponseBodyProject) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *GetProjectInfoResponseBodyProject) GetCreator() *string {
	return s.Creator
}

func (s *GetProjectInfoResponseBodyProject) GetCustomCode() *string {
	return s.CustomCode
}

func (s *GetProjectInfoResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *GetProjectInfoResponseBodyProject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetProjectInfoResponseBodyProject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetProjectInfoResponseBodyProject) GetIcon() *string {
	return s.Icon
}

func (s *GetProjectInfoResponseBodyProject) GetIconBig() *string {
	return s.IconBig
}

func (s *GetProjectInfoResponseBodyProject) GetIconGroup() *string {
	return s.IconGroup
}

func (s *GetProjectInfoResponseBodyProject) GetIconSmall() *string {
	return s.IconSmall
}

func (s *GetProjectInfoResponseBodyProject) GetId() *string {
	return s.Id
}

func (s *GetProjectInfoResponseBodyProject) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetProjectInfoResponseBodyProject) GetIdentifierPath() *string {
	return s.IdentifierPath
}

func (s *GetProjectInfoResponseBodyProject) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *GetProjectInfoResponseBodyProject) GetModifier() *string {
	return s.Modifier
}

func (s *GetProjectInfoResponseBodyProject) GetName() *string {
	return s.Name
}

func (s *GetProjectInfoResponseBodyProject) GetOrganizationIdentifier() *string {
	return s.OrganizationIdentifier
}

func (s *GetProjectInfoResponseBodyProject) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *GetProjectInfoResponseBodyProject) GetScope() *string {
	return s.Scope
}

func (s *GetProjectInfoResponseBodyProject) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *GetProjectInfoResponseBodyProject) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *GetProjectInfoResponseBodyProject) GetSubType() *string {
	return s.SubType
}

func (s *GetProjectInfoResponseBodyProject) GetTypeIdentifier() *string {
	return s.TypeIdentifier
}

func (s *GetProjectInfoResponseBodyProject) SetCategory(v string) *GetProjectInfoResponseBodyProject {
	s.Category = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetCategoryIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.CategoryIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetCreator(v string) *GetProjectInfoResponseBodyProject {
	s.Creator = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetCustomCode(v string) *GetProjectInfoResponseBodyProject {
	s.CustomCode = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetDescription(v string) *GetProjectInfoResponseBodyProject {
	s.Description = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetGmtCreate(v int64) *GetProjectInfoResponseBodyProject {
	s.GmtCreate = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetGmtModified(v int64) *GetProjectInfoResponseBodyProject {
	s.GmtModified = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIcon(v string) *GetProjectInfoResponseBodyProject {
	s.Icon = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIconBig(v string) *GetProjectInfoResponseBodyProject {
	s.IconBig = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIconGroup(v string) *GetProjectInfoResponseBodyProject {
	s.IconGroup = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIconSmall(v string) *GetProjectInfoResponseBodyProject {
	s.IconSmall = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetId(v string) *GetProjectInfoResponseBodyProject {
	s.Id = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.Identifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetIdentifierPath(v string) *GetProjectInfoResponseBodyProject {
	s.IdentifierPath = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetLogicalStatus(v string) *GetProjectInfoResponseBodyProject {
	s.LogicalStatus = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetModifier(v string) *GetProjectInfoResponseBodyProject {
	s.Modifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetName(v string) *GetProjectInfoResponseBodyProject {
	s.Name = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetOrganizationIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.OrganizationIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetParentIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.ParentIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetScope(v string) *GetProjectInfoResponseBodyProject {
	s.Scope = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetStatusIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.StatusIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetStatusStageIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.StatusStageIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetSubType(v string) *GetProjectInfoResponseBodyProject {
	s.SubType = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) SetTypeIdentifier(v string) *GetProjectInfoResponseBodyProject {
	s.TypeIdentifier = &v
	return s
}

func (s *GetProjectInfoResponseBodyProject) Validate() error {
	return dara.Validate(s)
}
