// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPersonalSourcesByObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPersonalSourcesByObjectResponseBody
	GetCode() *string
	SetMessage(v string) *ListPersonalSourcesByObjectResponseBody
	GetMessage() *string
	SetPage(v int64) *ListPersonalSourcesByObjectResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListPersonalSourcesByObjectResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListPersonalSourcesByObjectResponseBody
	GetRequestId() *string
	SetSources(v []*ListPersonalSourcesByObjectResponseBodySources) *ListPersonalSourcesByObjectResponseBody
	GetSources() []*ListPersonalSourcesByObjectResponseBodySources
	SetTotalCount(v int64) *ListPersonalSourcesByObjectResponseBody
	GetTotalCount() *int64
}

type ListPersonalSourcesByObjectResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为 success
	//
	// example:
	//
	// 请求的资源不存在
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 当前页码
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 当前页资源列表；空数组表示当前页没有可见匹配资源
	Sources []*ListPersonalSourcesByObjectResponseBodySources `json:"sources,omitempty" xml:"sources,omitempty" type:"Repeated"`
	// 权限过滤并按 Source ID 去重后的匹配总数
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListPersonalSourcesByObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalSourcesByObjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListPersonalSourcesByObjectResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPersonalSourcesByObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPersonalSourcesByObjectResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListPersonalSourcesByObjectResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPersonalSourcesByObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPersonalSourcesByObjectResponseBody) GetSources() []*ListPersonalSourcesByObjectResponseBodySources {
	return s.Sources
}

func (s *ListPersonalSourcesByObjectResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPersonalSourcesByObjectResponseBody) SetCode(v string) *ListPersonalSourcesByObjectResponseBody {
	s.Code = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBody) SetMessage(v string) *ListPersonalSourcesByObjectResponseBody {
	s.Message = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBody) SetPage(v int64) *ListPersonalSourcesByObjectResponseBody {
	s.Page = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBody) SetPageSize(v int64) *ListPersonalSourcesByObjectResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBody) SetRequestId(v string) *ListPersonalSourcesByObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBody) SetSources(v []*ListPersonalSourcesByObjectResponseBodySources) *ListPersonalSourcesByObjectResponseBody {
	s.Sources = v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBody) SetTotalCount(v int64) *ListPersonalSourcesByObjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBody) Validate() error {
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

type ListPersonalSourcesByObjectResponseBodySources struct {
	// 资源描述，可为空
	//
	// example:
	//
	// user-data test user_data_test fa6c967-20260928141235
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// 权限过滤后选择的代表目录 ID；不暴露不可读目录
	//
	// example:
	//
	// root
	DirectoryId *string `json:"directoryId,omitempty" xml:"directoryId,omitempty"`
	// 资源名称
	//
	// example:
	//
	// openapi-personal-public-url-7vw69qzc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 资源范围，固定为 PERSONAL
	//
	// example:
	//
	// read:user,read:repo,write:repo,read:org,read:group
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// 资源 ID
	//
	// example:
	//
	// src-da170aem1hkj8qfor7ig
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
	// 资源类型，如 TEXT、FILE、ONLINE_DOC、VOICE_MEETING 等
	//
	// example:
	//
	// BUILTIN
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// 资源实际处理状态，如 READY、RUNNING、PENDING、FAILED 等
	//
	// example:
	//
	// Stopped
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListPersonalSourcesByObjectResponseBodySources) String() string {
	return dara.Prettify(s)
}

func (s ListPersonalSourcesByObjectResponseBodySources) GoString() string {
	return s.String()
}

func (s *ListPersonalSourcesByObjectResponseBodySources) GetDescription() *string {
	return s.Description
}

func (s *ListPersonalSourcesByObjectResponseBodySources) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListPersonalSourcesByObjectResponseBodySources) GetName() *string {
	return s.Name
}

func (s *ListPersonalSourcesByObjectResponseBodySources) GetScope() *string {
	return s.Scope
}

func (s *ListPersonalSourcesByObjectResponseBodySources) GetSourceId() *string {
	return s.SourceId
}

func (s *ListPersonalSourcesByObjectResponseBodySources) GetSourceType() *string {
	return s.SourceType
}

func (s *ListPersonalSourcesByObjectResponseBodySources) GetStatus() *string {
	return s.Status
}

func (s *ListPersonalSourcesByObjectResponseBodySources) SetDescription(v string) *ListPersonalSourcesByObjectResponseBodySources {
	s.Description = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBodySources) SetDirectoryId(v string) *ListPersonalSourcesByObjectResponseBodySources {
	s.DirectoryId = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBodySources) SetName(v string) *ListPersonalSourcesByObjectResponseBodySources {
	s.Name = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBodySources) SetScope(v string) *ListPersonalSourcesByObjectResponseBodySources {
	s.Scope = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBodySources) SetSourceId(v string) *ListPersonalSourcesByObjectResponseBodySources {
	s.SourceId = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBodySources) SetSourceType(v string) *ListPersonalSourcesByObjectResponseBodySources {
	s.SourceType = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBodySources) SetStatus(v string) *ListPersonalSourcesByObjectResponseBodySources {
	s.Status = &v
	return s
}

func (s *ListPersonalSourcesByObjectResponseBodySources) Validate() error {
	return dara.Validate(s)
}
