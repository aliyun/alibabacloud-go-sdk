// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBatchTemplateVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetBatchTemplateVersionsResponseBody
	GetCode() *string
	SetData(v *GetBatchTemplateVersionsResponseBodyData) *GetBatchTemplateVersionsResponseBody
	GetData() *GetBatchTemplateVersionsResponseBodyData
	SetHttpStatusCode(v int32) *GetBatchTemplateVersionsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetBatchTemplateVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetBatchTemplateVersionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetBatchTemplateVersionsResponseBody
	GetSuccess() *bool
}

type GetBatchTemplateVersionsResponseBody struct {
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetBatchTemplateVersionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBatchTemplateVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetBatchTemplateVersionsResponseBody) GetData() *GetBatchTemplateVersionsResponseBodyData {
	return s.Data
}

func (s *GetBatchTemplateVersionsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetBatchTemplateVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetBatchTemplateVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBatchTemplateVersionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBatchTemplateVersionsResponseBody) SetCode(v string) *GetBatchTemplateVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetData(v *GetBatchTemplateVersionsResponseBodyData) *GetBatchTemplateVersionsResponseBody {
	s.Data = v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetHttpStatusCode(v int32) *GetBatchTemplateVersionsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetMessage(v string) *GetBatchTemplateVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetRequestId(v string) *GetBatchTemplateVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) SetSuccess(v bool) *GetBatchTemplateVersionsResponseBody {
	s.Success = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBatchTemplateVersionsResponseBodyData struct {
	TemplateVersionList []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionList `json:"TemplateVersionList,omitempty" xml:"TemplateVersionList,omitempty" type:"Repeated"`
}

func (s GetBatchTemplateVersionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBodyData) GetTemplateVersionList() []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	return s.TemplateVersionList
}

func (s *GetBatchTemplateVersionsResponseBodyData) SetTemplateVersionList(v []*GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) *GetBatchTemplateVersionsResponseBodyData {
	s.TemplateVersionList = v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyData) Validate() error {
	if s.TemplateVersionList != nil {
		for _, item := range s.TemplateVersionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetBatchTemplateVersionsResponseBodyDataTemplateVersionList struct {
	// example:
	//
	// 初始化提交
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// #!/bin/bash
	//
	// echo \\"hello world\\"
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 用于数据处理的Shell模板
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Python 3.7
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// example:
	//
	// 2026-01-01 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2026-05-28 15:30:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 李四
	Modifier *string `json:"Modifier,omitempty" xml:"Modifier,omitempty"`
	// example:
	//
	// 100002
	ModifierId *string `json:"ModifierId,omitempty" xml:"ModifierId,omitempty"`
	// example:
	//
	// 数据处理模板
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 10
	OperatorType *int32 `json:"OperatorType,omitempty" xml:"OperatorType,omitempty"`
	// example:
	//
	// 张三
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 100001
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 123456
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// 2
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 1
	Version *int64 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) String() string {
	return dara.Prettify(s)
}

func (s GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GoString() string {
	return s.String()
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetComment() *string {
	return s.Comment
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetContent() *string {
	return s.Content
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetDescription() *string {
	return s.Description
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetEngine() *string {
	return s.Engine
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetId() *int64 {
	return s.Id
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetModifier() *string {
	return s.Modifier
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetModifierId() *string {
	return s.ModifierId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetName() *string {
	return s.Name
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetOperatorType() *int32 {
	return s.OperatorType
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetOwner() *string {
	return s.Owner
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetStatus() *string {
	return s.Status
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) GetVersion() *int64 {
	return s.Version
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetComment(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Comment = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetContent(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Content = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetDescription(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Description = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetEngine(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Engine = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetGmtCreate(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.GmtCreate = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetGmtModified(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.GmtModified = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetId(v int64) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Id = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetModifier(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Modifier = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetModifierId(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.ModifierId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetName(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Name = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetOperatorType(v int32) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.OperatorType = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetOwner(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Owner = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetOwnerId(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.OwnerId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetProjectId(v int64) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.ProjectId = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetStatus(v string) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Status = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) SetVersion(v int64) *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList {
	s.Version = &v
	return s
}

func (s *GetBatchTemplateVersionsResponseBodyDataTemplateVersionList) Validate() error {
	return dara.Validate(s)
}
