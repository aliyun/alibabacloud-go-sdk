// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateProjectFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateProjectFieldResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateProjectFieldResponseBody
	GetErrorMessage() *string
	SetProject(v *UpdateProjectFieldResponseBodyProject) *UpdateProjectFieldResponseBody
	GetProject() *UpdateProjectFieldResponseBodyProject
	SetRequestId(v string) *UpdateProjectFieldResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateProjectFieldResponseBody
	GetSuccess() *bool
}

type UpdateProjectFieldResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// errormessage
	ErrorMessage *string                                `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	Project      *UpdateProjectFieldResponseBodyProject `json:"project,omitempty" xml:"project,omitempty" type:"Struct"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// true/false
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s UpdateProjectFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectFieldResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateProjectFieldResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateProjectFieldResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateProjectFieldResponseBody) GetProject() *UpdateProjectFieldResponseBodyProject {
	return s.Project
}

func (s *UpdateProjectFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateProjectFieldResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateProjectFieldResponseBody) SetErrorCode(v string) *UpdateProjectFieldResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateProjectFieldResponseBody) SetErrorMessage(v string) *UpdateProjectFieldResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateProjectFieldResponseBody) SetProject(v *UpdateProjectFieldResponseBodyProject) *UpdateProjectFieldResponseBody {
	s.Project = v
	return s
}

func (s *UpdateProjectFieldResponseBody) SetRequestId(v string) *UpdateProjectFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateProjectFieldResponseBody) SetSuccess(v bool) *UpdateProjectFieldResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateProjectFieldResponseBody) Validate() error {
	if s.Project != nil {
		if err := s.Project.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateProjectFieldResponseBodyProject struct {
	// example:
	//
	// Project
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
	CustomCode  *string                                           `json:"customCode,omitempty" xml:"customCode,omitempty"`
	Description *string                                           `json:"description,omitempty" xml:"description,omitempty"`
	FieldList   []*UpdateProjectFieldResponseBodyProjectFieldList `json:"fieldList,omitempty" xml:"fieldList,omitempty" type:"Repeated"`
	// example:
	//
	// 1623916393000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// example:
	//
	// 1684134690000
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
	// 66893
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 1e2ea73d24652d5874d001b351
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
	StatusName       *string `json:"statusName,omitempty" xml:"statusName,omitempty"`
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

func (s UpdateProjectFieldResponseBodyProject) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectFieldResponseBodyProject) GoString() string {
	return s.String()
}

func (s *UpdateProjectFieldResponseBodyProject) GetCategory() *string {
	return s.Category
}

func (s *UpdateProjectFieldResponseBodyProject) GetCategoryIdentifier() *string {
	return s.CategoryIdentifier
}

func (s *UpdateProjectFieldResponseBodyProject) GetCreator() *string {
	return s.Creator
}

func (s *UpdateProjectFieldResponseBodyProject) GetCustomCode() *string {
	return s.CustomCode
}

func (s *UpdateProjectFieldResponseBodyProject) GetDescription() *string {
	return s.Description
}

func (s *UpdateProjectFieldResponseBodyProject) GetFieldList() []*UpdateProjectFieldResponseBodyProjectFieldList {
	return s.FieldList
}

func (s *UpdateProjectFieldResponseBodyProject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *UpdateProjectFieldResponseBodyProject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *UpdateProjectFieldResponseBodyProject) GetIcon() *string {
	return s.Icon
}

func (s *UpdateProjectFieldResponseBodyProject) GetIconBig() *string {
	return s.IconBig
}

func (s *UpdateProjectFieldResponseBodyProject) GetIconGroup() *string {
	return s.IconGroup
}

func (s *UpdateProjectFieldResponseBodyProject) GetIconSmall() *string {
	return s.IconSmall
}

func (s *UpdateProjectFieldResponseBodyProject) GetId() *string {
	return s.Id
}

func (s *UpdateProjectFieldResponseBodyProject) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateProjectFieldResponseBodyProject) GetIdentifierPath() *string {
	return s.IdentifierPath
}

func (s *UpdateProjectFieldResponseBodyProject) GetLogicalStatus() *string {
	return s.LogicalStatus
}

func (s *UpdateProjectFieldResponseBodyProject) GetModifier() *string {
	return s.Modifier
}

func (s *UpdateProjectFieldResponseBodyProject) GetName() *string {
	return s.Name
}

func (s *UpdateProjectFieldResponseBodyProject) GetOrganizationIdentifier() *string {
	return s.OrganizationIdentifier
}

func (s *UpdateProjectFieldResponseBodyProject) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *UpdateProjectFieldResponseBodyProject) GetScope() *string {
	return s.Scope
}

func (s *UpdateProjectFieldResponseBodyProject) GetStatusIdentifier() *string {
	return s.StatusIdentifier
}

func (s *UpdateProjectFieldResponseBodyProject) GetStatusName() *string {
	return s.StatusName
}

func (s *UpdateProjectFieldResponseBodyProject) GetStatusStageIdentifier() *string {
	return s.StatusStageIdentifier
}

func (s *UpdateProjectFieldResponseBodyProject) GetSubType() *string {
	return s.SubType
}

func (s *UpdateProjectFieldResponseBodyProject) GetTypeIdentifier() *string {
	return s.TypeIdentifier
}

func (s *UpdateProjectFieldResponseBodyProject) SetCategory(v string) *UpdateProjectFieldResponseBodyProject {
	s.Category = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetCategoryIdentifier(v string) *UpdateProjectFieldResponseBodyProject {
	s.CategoryIdentifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetCreator(v string) *UpdateProjectFieldResponseBodyProject {
	s.Creator = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetCustomCode(v string) *UpdateProjectFieldResponseBodyProject {
	s.CustomCode = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetDescription(v string) *UpdateProjectFieldResponseBodyProject {
	s.Description = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetFieldList(v []*UpdateProjectFieldResponseBodyProjectFieldList) *UpdateProjectFieldResponseBodyProject {
	s.FieldList = v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetGmtCreate(v int64) *UpdateProjectFieldResponseBodyProject {
	s.GmtCreate = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetGmtModified(v int64) *UpdateProjectFieldResponseBodyProject {
	s.GmtModified = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetIcon(v string) *UpdateProjectFieldResponseBodyProject {
	s.Icon = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetIconBig(v string) *UpdateProjectFieldResponseBodyProject {
	s.IconBig = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetIconGroup(v string) *UpdateProjectFieldResponseBodyProject {
	s.IconGroup = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetIconSmall(v string) *UpdateProjectFieldResponseBodyProject {
	s.IconSmall = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetId(v string) *UpdateProjectFieldResponseBodyProject {
	s.Id = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetIdentifier(v string) *UpdateProjectFieldResponseBodyProject {
	s.Identifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetIdentifierPath(v string) *UpdateProjectFieldResponseBodyProject {
	s.IdentifierPath = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetLogicalStatus(v string) *UpdateProjectFieldResponseBodyProject {
	s.LogicalStatus = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetModifier(v string) *UpdateProjectFieldResponseBodyProject {
	s.Modifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetName(v string) *UpdateProjectFieldResponseBodyProject {
	s.Name = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetOrganizationIdentifier(v string) *UpdateProjectFieldResponseBodyProject {
	s.OrganizationIdentifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetParentIdentifier(v string) *UpdateProjectFieldResponseBodyProject {
	s.ParentIdentifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetScope(v string) *UpdateProjectFieldResponseBodyProject {
	s.Scope = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetStatusIdentifier(v string) *UpdateProjectFieldResponseBodyProject {
	s.StatusIdentifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetStatusName(v string) *UpdateProjectFieldResponseBodyProject {
	s.StatusName = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetStatusStageIdentifier(v string) *UpdateProjectFieldResponseBodyProject {
	s.StatusStageIdentifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetSubType(v string) *UpdateProjectFieldResponseBodyProject {
	s.SubType = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) SetTypeIdentifier(v string) *UpdateProjectFieldResponseBodyProject {
	s.TypeIdentifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProject) Validate() error {
	if s.FieldList != nil {
		for _, item := range s.FieldList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateProjectFieldResponseBodyProjectFieldList struct {
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// b787ff504c0b36f8b71d15a4fd
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// ace3wxxxxxxx7rff3ee
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// example:
	//
	// user
	ValueType *string `json:"valueType,omitempty" xml:"valueType,omitempty"`
}

func (s UpdateProjectFieldResponseBodyProjectFieldList) String() string {
	return dara.Prettify(s)
}

func (s UpdateProjectFieldResponseBodyProjectFieldList) GoString() string {
	return s.String()
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) GetIdentifier() *string {
	return s.Identifier
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) GetValue() *string {
	return s.Value
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) GetValueType() *string {
	return s.ValueType
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) SetDisplayName(v string) *UpdateProjectFieldResponseBodyProjectFieldList {
	s.DisplayName = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) SetIdentifier(v string) *UpdateProjectFieldResponseBodyProjectFieldList {
	s.Identifier = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) SetValue(v string) *UpdateProjectFieldResponseBodyProjectFieldList {
	s.Value = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) SetValueType(v string) *UpdateProjectFieldResponseBodyProjectFieldList {
	s.ValueType = &v
	return s
}

func (s *UpdateProjectFieldResponseBodyProjectFieldList) Validate() error {
	return dara.Validate(s)
}
