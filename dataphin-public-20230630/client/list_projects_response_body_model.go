// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProjectsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListProjectsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListProjectsResponseBody
	GetMessage() *string
	SetPageResult(v *ListProjectsResponseBodyPageResult) *ListProjectsResponseBody
	GetPageResult() *ListProjectsResponseBodyPageResult
	SetRequestId(v string) *ListProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProjectsResponseBody
	GetSuccess() *bool
}

type ListProjectsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListProjectsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProjectsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListProjectsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProjectsResponseBody) GetPageResult() *ListProjectsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProjectsResponseBody) SetCode(v string) *ListProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProjectsResponseBody) SetHttpStatusCode(v int32) *ListProjectsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListProjectsResponseBody) SetMessage(v string) *ListProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProjectsResponseBody) SetPageResult(v *ListProjectsResponseBodyPageResult) *ListProjectsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListProjectsResponseBody) SetRequestId(v string) *ListProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProjectsResponseBody) SetSuccess(v bool) *ListProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListProjectsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProjectsResponseBodyPageResult struct {
	ProjectList []*ListProjectsResponseBodyPageResultProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListProjectsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPageResult) GetProjectList() []*ListProjectsResponseBodyPageResultProjectList {
	return s.ProjectList
}

func (s *ListProjectsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProjectsResponseBodyPageResult) SetProjectList(v []*ListProjectsResponseBodyPageResultProjectList) *ListProjectsResponseBodyPageResult {
	s.ProjectList = v
	return s
}

func (s *ListProjectsResponseBodyPageResult) SetTotalCount(v int32) *ListProjectsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListProjectsResponseBodyPageResult) Validate() error {
	if s.ProjectList != nil {
		for _, item := range s.ProjectList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListProjectsResponseBodyPageResultProjectList struct {
	// example:
	//
	// 162112
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// 1121
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// ds1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// example:
	//
	// 2025-06-10 10:01:01
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-06-10 10:01:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BASIC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// example:
	//
	// demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012112
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 1122
	RealtimeDataSourceId *int64 `json:"RealtimeDataSourceId,omitempty" xml:"RealtimeDataSourceId,omitempty"`
	// example:
	//
	// ds2
	RealtimeDataSourceName *string `json:"RealtimeDataSourceName,omitempty" xml:"RealtimeDataSourceName,omitempty"`
	// example:
	//
	// GENERAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProjectsResponseBodyPageResultProjectList) String() string {
	return dara.Prettify(s)
}

func (s ListProjectsResponseBodyPageResultProjectList) GoString() string {
	return s.String()
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetDescription() *string {
	return s.Description
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetEnv() *string {
	return s.Env
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetId() *int64 {
	return s.Id
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetMode() *string {
	return s.Mode
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetName() *string {
	return s.Name
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetOwner() *string {
	return s.Owner
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetRealtimeDataSourceId() *int64 {
	return s.RealtimeDataSourceId
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetRealtimeDataSourceName() *string {
	return s.RealtimeDataSourceName
}

func (s *ListProjectsResponseBodyPageResultProjectList) GetType() *string {
	return s.Type
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetBizUnitId(v int64) *ListProjectsResponseBodyPageResultProjectList {
	s.BizUnitId = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetDataSourceId(v int64) *ListProjectsResponseBodyPageResultProjectList {
	s.DataSourceId = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetDataSourceName(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.DataSourceName = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetDescription(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.Description = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetDisplayName(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.DisplayName = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetEnv(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.Env = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetGmtCreate(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.GmtCreate = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetGmtModified(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.GmtModified = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetId(v int64) *ListProjectsResponseBodyPageResultProjectList {
	s.Id = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetMode(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.Mode = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetName(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.Name = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetOwner(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.Owner = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetOwnerName(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.OwnerName = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetRealtimeDataSourceId(v int64) *ListProjectsResponseBodyPageResultProjectList {
	s.RealtimeDataSourceId = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetRealtimeDataSourceName(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.RealtimeDataSourceName = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) SetType(v string) *ListProjectsResponseBodyPageResultProjectList {
	s.Type = &v
	return s
}

func (s *ListProjectsResponseBodyPageResultProjectList) Validate() error {
	return dara.Validate(s)
}
