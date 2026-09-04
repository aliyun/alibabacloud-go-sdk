// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListSkillsResponseBody
	GetCode() *string
	SetItems(v []*ListSkillsResponseBodyItems) *ListSkillsResponseBody
	GetItems() []*ListSkillsResponseBodyItems
	SetMessage(v string) *ListSkillsResponseBody
	GetMessage() *string
	SetPage(v int32) *ListSkillsResponseBody
	GetPage() *int32
	SetPageSize(v int32) *ListSkillsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSkillsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListSkillsResponseBody
	GetTotal() *int64
}

type ListSkillsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of skill cards.
	Items []*ListSkillsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The prompt message.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page. Default value: 20. Minimum value: 1. Maximum value: 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListSkillsResponseBody) GetItems() []*ListSkillsResponseBodyItems {
	return s.Items
}

func (s *ListSkillsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSkillsResponseBody) GetPage() *int32 {
	return s.Page
}

func (s *ListSkillsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListSkillsResponseBody) SetCode(v string) *ListSkillsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSkillsResponseBody) SetItems(v []*ListSkillsResponseBodyItems) *ListSkillsResponseBody {
	s.Items = v
	return s
}

func (s *ListSkillsResponseBody) SetMessage(v string) *ListSkillsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSkillsResponseBody) SetPage(v int32) *ListSkillsResponseBody {
	s.Page = &v
	return s
}

func (s *ListSkillsResponseBody) SetPageSize(v int32) *ListSkillsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSkillsResponseBody) SetRequestId(v string) *ListSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillsResponseBody) SetTotal(v int64) *ListSkillsResponseBody {
	s.Total = &v
	return s
}

func (s *ListSkillsResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillsResponseBodyItems struct {
	// The creation time.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	CreatedTime *string `json:"createdTime,omitempty" xml:"createdTime,omitempty"`
	// The description of the to-do card type.
	//
	// example:
	//
	// Sample description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution mode, such as CODE_AGENT or SYSTEM.
	//
	// example:
	//
	// string_value
	ExecuteMode *string `json:"executeMode,omitempty" xml:"executeMode,omitempty"`
	// Indicates whether the skill is globally accessible.
	//
	// example:
	//
	// true
	GlobalAccess *bool `json:"globalAccess,omitempty" xml:"globalAccess,omitempty"`
	// Indicates whether unpublished draft modifications exist.
	//
	// example:
	//
	// true
	HasDraftChanges *bool `json:"hasDraftChanges,omitempty" xml:"hasDraftChanges,omitempty"`
	// The file name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The actual skill code for execution.
	//
	// example:
	//
	// string_value
	SkillCode *string `json:"skillCode,omitempty" xml:"skillCode,omitempty"`
	// The skill definition ID.
	//
	// example:
	//
	// 1
	SkillHubDefinitionId *int64 `json:"skillHubDefinitionId,omitempty" xml:"skillHubDefinitionId,omitempty"`
	// The skill source type.
	//
	// example:
	//
	// BUILTIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The task status. Running is returned upon submission.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The key.
	//
	// example:
	//
	// string_value
	Tags []*string `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// The update timestamp, in milliseconds.
	//
	// example:
	//
	// 2023-10-01T12:00:00Z
	UpdatedTime *string `json:"updatedTime,omitempty" xml:"updatedTime,omitempty"`
	// The workflow definition version number.
	//
	// example:
	//
	// string_value
	VersionNumber *string `json:"versionNumber,omitempty" xml:"versionNumber,omitempty"`
}

func (s ListSkillsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyItems) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListSkillsResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListSkillsResponseBodyItems) GetExecuteMode() *string {
	return s.ExecuteMode
}

func (s *ListSkillsResponseBodyItems) GetGlobalAccess() *bool {
	return s.GlobalAccess
}

func (s *ListSkillsResponseBodyItems) GetHasDraftChanges() *bool {
	return s.HasDraftChanges
}

func (s *ListSkillsResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListSkillsResponseBodyItems) GetSkillCode() *string {
	return s.SkillCode
}

func (s *ListSkillsResponseBodyItems) GetSkillHubDefinitionId() *int64 {
	return s.SkillHubDefinitionId
}

func (s *ListSkillsResponseBodyItems) GetSourceType() *string {
	return s.SourceType
}

func (s *ListSkillsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListSkillsResponseBodyItems) GetTags() []*string {
	return s.Tags
}

func (s *ListSkillsResponseBodyItems) GetUpdatedTime() *string {
	return s.UpdatedTime
}

func (s *ListSkillsResponseBodyItems) GetVersionNumber() *string {
	return s.VersionNumber
}

func (s *ListSkillsResponseBodyItems) SetCreatedTime(v string) *ListSkillsResponseBodyItems {
	s.CreatedTime = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetDescription(v string) *ListSkillsResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetExecuteMode(v string) *ListSkillsResponseBodyItems {
	s.ExecuteMode = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetGlobalAccess(v bool) *ListSkillsResponseBodyItems {
	s.GlobalAccess = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetHasDraftChanges(v bool) *ListSkillsResponseBodyItems {
	s.HasDraftChanges = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetName(v string) *ListSkillsResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetSkillCode(v string) *ListSkillsResponseBodyItems {
	s.SkillCode = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetSkillHubDefinitionId(v int64) *ListSkillsResponseBodyItems {
	s.SkillHubDefinitionId = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetSourceType(v string) *ListSkillsResponseBodyItems {
	s.SourceType = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetStatus(v string) *ListSkillsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetTags(v []*string) *ListSkillsResponseBodyItems {
	s.Tags = v
	return s
}

func (s *ListSkillsResponseBodyItems) SetUpdatedTime(v string) *ListSkillsResponseBodyItems {
	s.UpdatedTime = &v
	return s
}

func (s *ListSkillsResponseBodyItems) SetVersionNumber(v string) *ListSkillsResponseBodyItems {
	s.VersionNumber = &v
	return s
}

func (s *ListSkillsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
