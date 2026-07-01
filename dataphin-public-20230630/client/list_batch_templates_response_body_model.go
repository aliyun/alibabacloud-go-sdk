// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBatchTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBatchTemplatesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListBatchTemplatesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBatchTemplatesResponseBody
	GetMessage() *string
	SetPageResult(v *ListBatchTemplatesResponseBodyPageResult) *ListBatchTemplatesResponseBody
	GetPageResult() *ListBatchTemplatesResponseBodyPageResult
	SetRequestId(v string) *ListBatchTemplatesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBatchTemplatesResponseBody
	GetSuccess() *bool
}

type ListBatchTemplatesResponseBody struct {
	// The backend response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The backend exception details.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result.
	PageResult *ListBatchTemplatesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListBatchTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBatchTemplatesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBatchTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBatchTemplatesResponseBody) GetPageResult() *ListBatchTemplatesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListBatchTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBatchTemplatesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBatchTemplatesResponseBody) SetCode(v string) *ListBatchTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetHttpStatusCode(v int32) *ListBatchTemplatesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetMessage(v string) *ListBatchTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetPageResult(v *ListBatchTemplatesResponseBodyPageResult) *ListBatchTemplatesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetRequestId(v string) *ListBatchTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) SetSuccess(v bool) *ListBatchTemplatesResponseBody {
	s.Success = &v
	return s
}

func (s *ListBatchTemplatesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBatchTemplatesResponseBodyPageResult struct {
	// The list of template records.
	TemplateList []*ListBatchTemplatesResponseBodyPageResultTemplateList `json:"TemplateList,omitempty" xml:"TemplateList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListBatchTemplatesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBodyPageResult) GetTemplateList() []*ListBatchTemplatesResponseBodyPageResultTemplateList {
	return s.TemplateList
}

func (s *ListBatchTemplatesResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListBatchTemplatesResponseBodyPageResult) SetTemplateList(v []*ListBatchTemplatesResponseBodyPageResultTemplateList) *ListBatchTemplatesResponseBodyPageResult {
	s.TemplateList = v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResult) SetTotalCount(v int32) *ListBatchTemplatesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResult) Validate() error {
	if s.TemplateList != nil {
		for _, item := range s.TemplateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBatchTemplatesResponseBodyPageResultTemplateList struct {
	// The template submission comment.
	//
	// example:
	//
	// 初始化提交
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The template content.
	//
	// example:
	//
	// #!/bin/bash
	//
	// echo \\"hello world\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The current latest version.
	//
	// example:
	//
	// 1
	CurrentVersion *int64 `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	// The template description.
	//
	// example:
	//
	// 用于数据处理的Shell模板
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The compute engine version. Currently supported Python versions: Python 2.7 and Python 3.7.
	//
	// example:
	//
	// Python 3.7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The template creation time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-01-01 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The template update time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-05-28 15:30:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The template ID, which is the same as the menu tree node ID.
	//
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The last modifier of the template.
	//
	// example:
	//
	// 李四
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// The ID of the last modifier of the template.
	//
	// example:
	//
	// 100002
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// The template name.
	//
	// example:
	//
	// 数据处理模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The node type. Valid values: 10 (Shell) and 21 (Python).
	//
	// example:
	//
	// 10
	OperatorType *int32 `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// The template owner.
	//
	// example:
	//
	// 张三
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The template owner ID.
	//
	// example:
	//
	// 100001
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The project ID.
	//
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The template status. Valid values: 0 (draft), 2 (submitted), and 100 (in development).
	//
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListBatchTemplatesResponseBodyPageResultTemplateList) String() string {
	return dara.Prettify(s)
}

func (s ListBatchTemplatesResponseBodyPageResultTemplateList) GoString() string {
	return s.String()
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetComment() *string {
	return s.Comment
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetContent() *string {
	return s.Content
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetCurrentVersion() *int64 {
	return s.CurrentVersion
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetDescription() *string {
	return s.Description
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetEngine() *string {
	return s.Engine
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetId() *int64 {
	return s.Id
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetModifier() *string {
	return s.Modifier
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetModifierId() *string {
	return s.ModifierId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetName() *string {
	return s.Name
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetOperatorType() *int32 {
	return s.OperatorType
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetOwner() *string {
	return s.Owner
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) GetStatus() *string {
	return s.Status
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetComment(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Comment = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetContent(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Content = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetCurrentVersion(v int64) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.CurrentVersion = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetDescription(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Description = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetEngine(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Engine = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetGmtCreate(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.GmtCreate = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetGmtModified(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.GmtModified = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetId(v int64) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Id = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetModifier(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Modifier = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetModifierId(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.ModifierId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetName(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Name = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetOperatorType(v int32) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.OperatorType = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetOwner(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Owner = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetOwnerId(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.OwnerId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetProjectId(v int64) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.ProjectId = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) SetStatus(v string) *ListBatchTemplatesResponseBodyPageResultTemplateList {
	s.Status = &v
	return s
}

func (s *ListBatchTemplatesResponseBodyPageResultTemplateList) Validate() error {
	return dara.Validate(s)
}
