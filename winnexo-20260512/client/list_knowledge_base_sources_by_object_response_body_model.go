// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKnowledgeBaseSourcesByObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListKnowledgeBaseSourcesByObjectResponseBody
	GetCode() *string
	SetMessage(v string) *ListKnowledgeBaseSourcesByObjectResponseBody
	GetMessage() *string
	SetPage(v int64) *ListKnowledgeBaseSourcesByObjectResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListKnowledgeBaseSourcesByObjectResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListKnowledgeBaseSourcesByObjectResponseBody
	GetRequestId() *string
	SetSources(v []*ListKnowledgeBaseSourcesByObjectResponseBodySources) *ListKnowledgeBaseSourcesByObjectResponseBody
	GetSources() []*ListKnowledgeBaseSourcesByObjectResponseBodySources
	SetTotalCount(v int64) *ListKnowledgeBaseSourcesByObjectResponseBody
	GetTotalCount() *int64
}

type ListKnowledgeBaseSourcesByObjectResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The subject type filter.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The referenced sources. The key is the project name, and the value is the list of selected Logstores under the project, which is consistent with the structure submitted during creation. `["*"]` indicates all Logstores under the project. Otherwise, it is a list of exact Logstore names. If no sources are referenced, an empty object `{}` is returned.
	Sources []*ListKnowledgeBaseSourcesByObjectResponseBodySources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
	// The total number of results.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListKnowledgeBaseSourcesByObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseSourcesByObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) GetSources() []*ListKnowledgeBaseSourcesByObjectResponseBodySources {
	return s.Sources
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) SetCode(v string) *ListKnowledgeBaseSourcesByObjectResponseBody {
	s.Code = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) SetMessage(v string) *ListKnowledgeBaseSourcesByObjectResponseBody {
	s.Message = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) SetPage(v int64) *ListKnowledgeBaseSourcesByObjectResponseBody {
	s.Page = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) SetPageSize(v int64) *ListKnowledgeBaseSourcesByObjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) SetRequestId(v string) *ListKnowledgeBaseSourcesByObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) SetSources(v []*ListKnowledgeBaseSourcesByObjectResponseBodySources) *ListKnowledgeBaseSourcesByObjectResponseBody {
	s.Sources = v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) SetTotalCount(v int64) *ListKnowledgeBaseSourcesByObjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBody) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKnowledgeBaseSourcesByObjectResponseBodySources struct {
	// The description of the to-do card type.
	//
	// example:
	//
	// user-data test user_data_test fa6c967-20260928141235
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The directory ID.
	//
	// example:
	//
	// oort
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// The file name.
	//
	// example:
	//
	// openapi-feishu-personal-2ec49p76
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The permission scope.
	//
	// example:
	//
	// read:user,read:repo,write:repo,read:org,read:group
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The ID of the enterprise knowledge base FILE data source to be replaced.
	//
	// example:
	//
	// 2001188
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// The data source type.
	//
	// example:
	//
	// BUILTIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// The data source status after re-parsing.
	//
	// example:
	//
	// Stopped
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListKnowledgeBaseSourcesByObjectResponseBodySources) String() string {
	return dara.Prettify(s)
}

func (s ListKnowledgeBaseSourcesByObjectResponseBodySources) GoString() string {
	return s.String()
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) GetDescription() *string {
	return s.Description
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) GetName() *string {
	return s.Name
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) GetScope() *string {
	return s.Scope
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) GetSourceId() *string {
	return s.SourceId
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) GetSourceType() *string {
	return s.SourceType
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) GetStatus() *string {
	return s.Status
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) SetDescription(v string) *ListKnowledgeBaseSourcesByObjectResponseBodySources {
	s.Description = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) SetDirectoryId(v string) *ListKnowledgeBaseSourcesByObjectResponseBodySources {
	s.DirectoryId = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) SetName(v string) *ListKnowledgeBaseSourcesByObjectResponseBodySources {
	s.Name = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) SetScope(v string) *ListKnowledgeBaseSourcesByObjectResponseBodySources {
	s.Scope = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) SetSourceId(v string) *ListKnowledgeBaseSourcesByObjectResponseBodySources {
	s.SourceId = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) SetSourceType(v string) *ListKnowledgeBaseSourcesByObjectResponseBodySources {
	s.SourceType = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) SetStatus(v string) *ListKnowledgeBaseSourcesByObjectResponseBodySources {
	s.Status = &v
	return s
}

func (s *ListKnowledgeBaseSourcesByObjectResponseBodySources) Validate() error {
	return dara.Validate(s)
}
