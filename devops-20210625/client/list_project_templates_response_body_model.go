// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListProjectTemplatesResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListProjectTemplatesResponseBody
	GetErrorMsg() *string
	SetRequestId(v string) *ListProjectTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectTemplatesResponseBody
	GetSuccess() *bool
	SetTemplates(v []*ListProjectTemplatesResponseBodyTemplates) *ListProjectTemplatesResponseBody
	GetTemplates() []*ListProjectTemplatesResponseBodyTemplates
}

type ListProjectTemplatesResponseBody struct {
	// example:
	//
	// Openapi.RequestError
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// error
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success   *bool                                        `json:"success,omitempty" xml:"success,omitempty"`
	Templates []*ListProjectTemplatesResponseBodyTemplates `json:"templates,omitempty" xml:"templates,omitempty" type:"Repeated"`
}

func (s ListProjectTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectTemplatesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListProjectTemplatesResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListProjectTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectTemplatesResponseBody) GetTemplates() []*ListProjectTemplatesResponseBodyTemplates {
	return s.Templates
}

func (s *ListProjectTemplatesResponseBody) SetErrorCode(v string) *ListProjectTemplatesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProjectTemplatesResponseBody) SetErrorMsg(v string) *ListProjectTemplatesResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListProjectTemplatesResponseBody) SetRequestId(v string) *ListProjectTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectTemplatesResponseBody) SetSuccess(v bool) *ListProjectTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectTemplatesResponseBody) SetTemplates(v []*ListProjectTemplatesResponseBodyTemplates) *ListProjectTemplatesResponseBody {
	s.Templates = v
	return s
}

func (s *ListProjectTemplatesResponseBody) Validate() error {
	if s.Templates != nil {
		for _, item := range s.Templates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectTemplatesResponseBodyTemplates struct {
	// example:
	//
	// null
	CopyFrom *string `json:"copyFrom,omitempty" xml:"copyFrom,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 迭代的具体信息
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
	// https://img.xxxx.png
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	// example:
	//
	// e8bxxxxxxxxxxxxxxxx23
	Identifier *string `json:"identifier,omitempty" xml:"identifier,omitempty"`
	// example:
	//
	// 19xx7043xxxxxxx914
	Modifier *string `json:"modifier,omitempty" xml:"modifier,omitempty"`
	// example:
	//
	// 专用模板
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Test Template
	NameEn *string `json:"nameEn,omitempty" xml:"nameEn,omitempty"`
	// example:
	//
	// Project
	ResourceCategory *string `json:"resourceCategory,omitempty" xml:"resourceCategory,omitempty"`
	// example:
	//
	// null
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// null
	SpaceIdentifier *string `json:"spaceIdentifier,omitempty" xml:"spaceIdentifier,omitempty"`
	// example:
	//
	// null
	SpaceType *string `json:"spaceType,omitempty" xml:"spaceType,omitempty"`
	// example:
	//
	// 4
	Type *int64 `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListProjectTemplatesResponseBodyTemplates) String() string {
	return dara.Prettify(s)
}

func (s ListProjectTemplatesResponseBodyTemplates) GoString() string {
	return s.String()
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetCopyFrom() *string {
	return s.CopyFrom
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetCreator() *string {
	return s.Creator
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetDescription() *string {
	return s.Description
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetIcon() *string {
	return s.Icon
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetModifier() *string {
	return s.Modifier
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetName() *string {
	return s.Name
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetNameEn() *string {
	return s.NameEn
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetResourceCategory() *string {
	return s.ResourceCategory
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetSpaceIdentifier() *string {
	return s.SpaceIdentifier
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetSpaceType() *string {
	return s.SpaceType
}

func (s *ListProjectTemplatesResponseBodyTemplates) GetType() *int64 {
	return s.Type
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetCopyFrom(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.CopyFrom = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetCreator(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Creator = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetDescription(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Description = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetGmtCreate(v int64) *ListProjectTemplatesResponseBodyTemplates {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetGmtModified(v int64) *ListProjectTemplatesResponseBodyTemplates {
	s.GmtModified = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetIcon(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Icon = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetIdentifier(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Identifier = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetModifier(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Modifier = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetName(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.Name = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetNameEn(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.NameEn = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetResourceCategory(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.ResourceCategory = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetResourceType(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.ResourceType = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetSpaceIdentifier(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.SpaceIdentifier = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetSpaceType(v string) *ListProjectTemplatesResponseBodyTemplates {
	s.SpaceType = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) SetType(v int64) *ListProjectTemplatesResponseBodyTemplates {
	s.Type = &v
	return s
}

func (s *ListProjectTemplatesResponseBodyTemplates) Validate() error {
	return dara.Validate(s)
}
