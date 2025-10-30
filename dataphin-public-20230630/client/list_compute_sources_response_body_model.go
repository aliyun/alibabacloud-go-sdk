// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListComputeSourcesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListComputeSourcesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListComputeSourcesResponseBody
	GetMessage() *string
	SetPageResult(v *ListComputeSourcesResponseBodyPageResult) *ListComputeSourcesResponseBody
	GetPageResult() *ListComputeSourcesResponseBodyPageResult
	SetRequestId(v string) *ListComputeSourcesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListComputeSourcesResponseBody
	GetSuccess() *bool
}

type ListComputeSourcesResponseBody struct {
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
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	PageResult *ListComputeSourcesResponseBodyPageResult `json:"PageResult,omitempty" xml:"PageResult,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListComputeSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeSourcesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListComputeSourcesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListComputeSourcesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListComputeSourcesResponseBody) GetPageResult() *ListComputeSourcesResponseBodyPageResult {
	return s.PageResult
}

func (s *ListComputeSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeSourcesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListComputeSourcesResponseBody) SetCode(v string) *ListComputeSourcesResponseBody {
	s.Code = &v
	return s
}

func (s *ListComputeSourcesResponseBody) SetHttpStatusCode(v int32) *ListComputeSourcesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListComputeSourcesResponseBody) SetMessage(v string) *ListComputeSourcesResponseBody {
	s.Message = &v
	return s
}

func (s *ListComputeSourcesResponseBody) SetPageResult(v *ListComputeSourcesResponseBodyPageResult) *ListComputeSourcesResponseBody {
	s.PageResult = v
	return s
}

func (s *ListComputeSourcesResponseBody) SetRequestId(v string) *ListComputeSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeSourcesResponseBody) SetSuccess(v bool) *ListComputeSourcesResponseBody {
	s.Success = &v
	return s
}

func (s *ListComputeSourcesResponseBody) Validate() error {
	if s.PageResult != nil {
		if err := s.PageResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeSourcesResponseBodyPageResult struct {
	ComputeSourceList []*ListComputeSourcesResponseBodyPageResultComputeSourceList `json:"ComputeSourceList,omitempty" xml:"ComputeSourceList,omitempty" type:"Repeated"`
	// example:
	//
	// 68
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListComputeSourcesResponseBodyPageResult) String() string {
	return dara.Prettify(s)
}

func (s ListComputeSourcesResponseBodyPageResult) GoString() string {
	return s.String()
}

func (s *ListComputeSourcesResponseBodyPageResult) GetComputeSourceList() []*ListComputeSourcesResponseBodyPageResultComputeSourceList {
	return s.ComputeSourceList
}

func (s *ListComputeSourcesResponseBodyPageResult) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListComputeSourcesResponseBodyPageResult) SetComputeSourceList(v []*ListComputeSourcesResponseBodyPageResultComputeSourceList) *ListComputeSourcesResponseBodyPageResult {
	s.ComputeSourceList = v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResult) SetTotalCount(v int32) *ListComputeSourcesResponseBodyPageResult {
	s.TotalCount = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResult) Validate() error {
	if s.ComputeSourceList != nil {
		for _, item := range s.ComputeSourceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeSourcesResponseBodyPageResultComputeSourceList struct {
	BindProject *bool `json:"BindProject,omitempty" xml:"BindProject,omitempty"`
	// example:
	//
	// 10132131111
	BindProjectId *int64 `json:"BindProjectId,omitempty" xml:"BindProjectId,omitempty"`
	// example:
	//
	// testPrj
	BindProjectName *string `json:"BindProjectName,omitempty" xml:"BindProjectName,omitempty"`
	// example:
	//
	// 30012211
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// example:
	//
	// 张三
	CreatorName *string `json:"CreatorName,omitempty" xml:"CreatorName,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test1011
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2025-06-30 08:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2025-06-30 08:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 102111
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// test1011
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 30012211
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListComputeSourcesResponseBodyPageResultComputeSourceList) String() string {
	return dara.Prettify(s)
}

func (s ListComputeSourcesResponseBodyPageResultComputeSourceList) GoString() string {
	return s.String()
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetBindProject() *bool {
	return s.BindProject
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetBindProjectId() *int64 {
	return s.BindProjectId
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetBindProjectName() *string {
	return s.BindProjectName
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetCreator() *string {
	return s.Creator
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetCreatorName() *string {
	return s.CreatorName
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetDescription() *string {
	return s.Description
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetId() *int64 {
	return s.Id
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetName() *string {
	return s.Name
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetOwner() *string {
	return s.Owner
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) GetType() *string {
	return s.Type
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetBindProject(v bool) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.BindProject = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetBindProjectId(v int64) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.BindProjectId = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetBindProjectName(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.BindProjectName = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetCreator(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.Creator = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetCreatorName(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.CreatorName = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetDescription(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.Description = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetDisplayName(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.DisplayName = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetGmtCreate(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.GmtCreate = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetGmtModified(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.GmtModified = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetId(v int64) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.Id = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetName(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.Name = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetOwner(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.Owner = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetOwnerName(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.OwnerName = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) SetType(v string) *ListComputeSourcesResponseBodyPageResultComputeSourceList {
	s.Type = &v
	return s
}

func (s *ListComputeSourcesResponseBodyPageResultComputeSourceList) Validate() error {
	return dara.Validate(s)
}
