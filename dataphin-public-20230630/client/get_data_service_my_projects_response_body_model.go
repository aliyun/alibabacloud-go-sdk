// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceMyProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceMyProjectsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetDataServiceMyProjectsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceMyProjectsResponseBody
	GetMessage() *string
	SetProjectList(v []*GetDataServiceMyProjectsResponseBodyProjectList) *GetDataServiceMyProjectsResponseBody
	GetProjectList() []*GetDataServiceMyProjectsResponseBodyProjectList
	SetRequestId(v string) *GetDataServiceMyProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceMyProjectsResponseBody
	GetSuccess() *bool
}

type GetDataServiceMyProjectsResponseBody struct {
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
	Message     *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	ProjectList []*GetDataServiceMyProjectsResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceMyProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceMyProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceMyProjectsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceMyProjectsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceMyProjectsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceMyProjectsResponseBody) GetProjectList() []*GetDataServiceMyProjectsResponseBodyProjectList {
	return s.ProjectList
}

func (s *GetDataServiceMyProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceMyProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceMyProjectsResponseBody) SetCode(v string) *GetDataServiceMyProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceMyProjectsResponseBody) SetHttpStatusCode(v int32) *GetDataServiceMyProjectsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceMyProjectsResponseBody) SetMessage(v string) *GetDataServiceMyProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceMyProjectsResponseBody) SetProjectList(v []*GetDataServiceMyProjectsResponseBodyProjectList) *GetDataServiceMyProjectsResponseBody {
	s.ProjectList = v
	return s
}

func (s *GetDataServiceMyProjectsResponseBody) SetRequestId(v string) *GetDataServiceMyProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceMyProjectsResponseBody) SetSuccess(v bool) *GetDataServiceMyProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceMyProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceMyProjectsResponseBodyProjectList struct {
	// example:
	//
	// 102011
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 1
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s GetDataServiceMyProjectsResponseBodyProjectList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceMyProjectsResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *GetDataServiceMyProjectsResponseBodyProjectList) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceMyProjectsResponseBodyProjectList) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDataServiceMyProjectsResponseBodyProjectList) GetRole() *int32 {
	return s.Role
}

func (s *GetDataServiceMyProjectsResponseBodyProjectList) SetProjectId(v int32) *GetDataServiceMyProjectsResponseBodyProjectList {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceMyProjectsResponseBodyProjectList) SetProjectName(v string) *GetDataServiceMyProjectsResponseBodyProjectList {
	s.ProjectName = &v
	return s
}

func (s *GetDataServiceMyProjectsResponseBodyProjectList) SetRole(v int32) *GetDataServiceMyProjectsResponseBodyProjectList {
	s.Role = &v
	return s
}

func (s *GetDataServiceMyProjectsResponseBodyProjectList) Validate() error {
	return dara.Validate(s)
}
