// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListConfigTemplatesResponseBody
	GetCode() *int32
	SetData(v *ListConfigTemplatesResponseBodyData) *ListConfigTemplatesResponseBody
	GetData() *ListConfigTemplatesResponseBodyData
	SetMessage(v string) *ListConfigTemplatesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConfigTemplatesResponseBody
	GetRequestId() *string
}

type ListConfigTemplatesResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of the request.
	Data *ListConfigTemplatesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The additional information that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4D9F-DR94-FD****************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListConfigTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConfigTemplatesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListConfigTemplatesResponseBody) GetData() *ListConfigTemplatesResponseBodyData {
	return s.Data
}

func (s *ListConfigTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConfigTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConfigTemplatesResponseBody) SetCode(v int32) *ListConfigTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListConfigTemplatesResponseBody) SetData(v *ListConfigTemplatesResponseBodyData) *ListConfigTemplatesResponseBody {
	s.Data = v
	return s
}

func (s *ListConfigTemplatesResponseBody) SetMessage(v string) *ListConfigTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConfigTemplatesResponseBody) SetRequestId(v string) *ListConfigTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConfigTemplatesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConfigTemplatesResponseBodyData struct {
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The configuration templates.
	Result []*ListConfigTemplatesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// The total number of configuration templates.
	//
	// example:
	//
	// 100
	TotalSize *int64 `json:"TotalSize,omitempty" xml:"TotalSize,omitempty"`
}

func (s ListConfigTemplatesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTemplatesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConfigTemplatesResponseBodyData) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListConfigTemplatesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConfigTemplatesResponseBodyData) GetResult() []*ListConfigTemplatesResponseBodyDataResult {
	return s.Result
}

func (s *ListConfigTemplatesResponseBodyData) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *ListConfigTemplatesResponseBodyData) SetCurrentPage(v int32) *ListConfigTemplatesResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyData) SetPageSize(v int32) *ListConfigTemplatesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyData) SetResult(v []*ListConfigTemplatesResponseBodyDataResult) *ListConfigTemplatesResponseBodyData {
	s.Result = v
	return s
}

func (s *ListConfigTemplatesResponseBodyData) SetTotalSize(v int64) *ListConfigTemplatesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConfigTemplatesResponseBodyDataResult struct {
	// The content of the configuration template.
	//
	// example:
	//
	// {"name":"William"}
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description of the configuration template.
	//
	// example:
	//
	// Test configuration template
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The data format of the configuration template.
	//
	// example:
	//
	// JSON
	Format *string `json:"Format,omitempty" xml:"Format,omitempty"`
	// The time when the configuration template was created.
	//
	// example:
	//
	// 1638171689626
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The time when the configuration template was updated.
	//
	// example:
	//
	// 1638171689626
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the configuration template.
	//
	// example:
	//
	// 3d84efaf-37d9-49fb-a3a8-b38d5c2b460c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the configuration template.
	//
	// example:
	//
	// config-tmpl-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListConfigTemplatesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ListConfigTemplatesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ListConfigTemplatesResponseBodyDataResult) GetContent() *string {
	return s.Content
}

func (s *ListConfigTemplatesResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *ListConfigTemplatesResponseBodyDataResult) GetFormat() *string {
	return s.Format
}

func (s *ListConfigTemplatesResponseBodyDataResult) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *ListConfigTemplatesResponseBodyDataResult) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *ListConfigTemplatesResponseBodyDataResult) GetId() *string {
	return s.Id
}

func (s *ListConfigTemplatesResponseBodyDataResult) GetName() *string {
	return s.Name
}

func (s *ListConfigTemplatesResponseBodyDataResult) SetContent(v string) *ListConfigTemplatesResponseBodyDataResult {
	s.Content = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyDataResult) SetDescription(v string) *ListConfigTemplatesResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyDataResult) SetFormat(v string) *ListConfigTemplatesResponseBodyDataResult {
	s.Format = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyDataResult) SetGmtCreate(v int64) *ListConfigTemplatesResponseBodyDataResult {
	s.GmtCreate = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyDataResult) SetGmtModified(v int64) *ListConfigTemplatesResponseBodyDataResult {
	s.GmtModified = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyDataResult) SetId(v string) *ListConfigTemplatesResponseBodyDataResult {
	s.Id = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyDataResult) SetName(v string) *ListConfigTemplatesResponseBodyDataResult {
	s.Name = &v
	return s
}

func (s *ListConfigTemplatesResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
