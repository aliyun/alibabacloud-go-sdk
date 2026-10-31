// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListUserProjectsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListUserProjectsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListUserProjectsResponseBody
	GetMessage() *string
	SetPageResult(v *ListUserProjectsResponseBodyPageResult) *ListUserProjectsResponseBody
	GetPageResult() *ListUserProjectsResponseBodyPageResult
	SetRequestId(v string) *ListUserProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListUserProjectsResponseBody
	GetSuccess() *bool
}

type ListUserProjectsResponseBody struct {
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
	// The details of the backend exception.
	//
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The paged query result.
	PageResult *ListUserProjectsResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListUserProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListUserProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserProjectsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListUserProjectsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListUserProjectsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListUserProjectsResponseBody) GetPageResult() *ListUserProjectsResponseBodyPageResult {
	return s.PageResult
}

func (s *ListUserProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListUserProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListUserProjectsResponseBody) SetCode(v string) *ListUserProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *ListUserProjectsResponseBody) SetHttpStatusCode(v int32) *ListUserProjectsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListUserProjectsResponseBody) SetMessage(v string) *ListUserProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *ListUserProjectsResponseBody) SetPageResult(v *ListUserProjectsResponseBodyPageResult) *ListUserProjectsResponseBody {
	s.PageResult = v
	return s
}

func (s *ListUserProjectsResponseBody) SetRequestId(v string) *ListUserProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListUserProjectsResponseBody) SetSuccess(v bool) *ListUserProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListUserProjectsResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListUserProjectsResponseBodyPageResult struct {
	// The list of projects that the user participates in.
	ProjectList []*ListUserProjectsResponseBodyPageResultProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// The total number of records.
	//
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListUserProjectsResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListUserProjectsResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListUserProjectsResponseBodyPageResult) GetProjectList() []*ListUserProjectsResponseBodyPageResultProjectList {
	return s.ProjectList
}

func (s *ListUserProjectsResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListUserProjectsResponseBodyPageResult) SetProjectList(v []*ListUserProjectsResponseBodyPageResultProjectList) *ListUserProjectsResponseBodyPageResult {
	s.ProjectList = v
	return s
}

func (s *ListUserProjectsResponseBodyPageResult) SetTotalCount(v int32) *ListUserProjectsResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResult) Validate() error {
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

type ListUserProjectsResponseBodyPageResultProjectList struct {
	// The ID of the business unit.
	//
	// example:
	//
	// 162112
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The ID of the compute source.
	//
	// example:
	//
	// 1121
	DataSourceId *int64 `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The name of the compute source.
	//
	// example:
	//
	// ds1
	DataSourceName *string `json:"DataSourceName,omitempty" xml:"DataSourceName,omitempty"`
	// The description.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the project.
	//
	// example:
	//
	// Test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment identifier.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
	// The creation time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-06-10 10:01:01
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2025-06-10 10:01:01
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the project.
	//
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The project mode.
	//
	// example:
	//
	// BASIC
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the project.
	//
	// example:
	//
	// demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the project.
	//
	// example:
	//
	// 30012112
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The owner of the project.
	//
	// example:
	//
	// Zhang San
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The ID of the real-time compute source.
	//
	// example:
	//
	// 1122
	RealtimeDataSourceId *int64 `json:"RealtimeDataSourceId,omitempty" xml:"RealtimeDataSourceId,omitempty"`
	// The name of the real-time compute source.
	//
	// example:
	//
	// ds2
	RealtimeDataSourceName *string `json:"RealtimeDataSourceName,omitempty" xml:"RealtimeDataSourceName,omitempty"`
	// The project type.
	//
	// example:
	//
	// GENERAL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListUserProjectsResponseBodyPageResultProjectList) String() string {
	return dara.Prettify(s)
}

func (s ListUserProjectsResponseBodyPageResultProjectList) GoString() string {
	return s.String()
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetDataSourceId() *int64 {
	return s.DataSourceId
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetDescription() *string {
	return s.Description
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetEnv() *string {
	return s.Env
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetId() *int64 {
	return s.Id
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetMode() *string {
	return s.Mode
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetName() *string {
	return s.Name
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetOwner() *string {
	return s.Owner
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetRealtimeDataSourceId() *int64 {
	return s.RealtimeDataSourceId
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetRealtimeDataSourceName() *string {
	return s.RealtimeDataSourceName
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) GetType() *string {
	return s.Type
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetBizUnitId(v int64) *ListUserProjectsResponseBodyPageResultProjectList {
	s.BizUnitId = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetDataSourceId(v int64) *ListUserProjectsResponseBodyPageResultProjectList {
	s.DataSourceId = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetDataSourceName(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.DataSourceName = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetDescription(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.Description = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetDisplayName(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.DisplayName = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetEnv(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.Env = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetGmtCreate(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.GmtCreate = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetGmtModified(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.GmtModified = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetId(v int64) *ListUserProjectsResponseBodyPageResultProjectList {
	s.Id = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetMode(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.Mode = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetName(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.Name = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetOwner(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.Owner = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetOwnerName(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.OwnerName = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetRealtimeDataSourceId(v int64) *ListUserProjectsResponseBodyPageResultProjectList {
	s.RealtimeDataSourceId = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetRealtimeDataSourceName(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.RealtimeDataSourceName = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) SetType(v string) *ListUserProjectsResponseBodyPageResultProjectList {
	s.Type = &v
	return s
}

func (s *ListUserProjectsResponseBodyPageResultProjectList) Validate() error {
	return dara.Validate(s)
}
