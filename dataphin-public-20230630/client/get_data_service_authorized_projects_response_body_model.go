// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataServiceAuthorizedProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataServiceAuthorizedProjectsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetDataServiceAuthorizedProjectsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataServiceAuthorizedProjectsResponseBody
	GetMessage() *string
	SetProjectList(v []*GetDataServiceAuthorizedProjectsResponseBodyProjectList) *GetDataServiceAuthorizedProjectsResponseBody
	GetProjectList() []*GetDataServiceAuthorizedProjectsResponseBodyProjectList
	SetRequestId(v string) *GetDataServiceAuthorizedProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataServiceAuthorizedProjectsResponseBody
	GetSuccess() *bool
}

type GetDataServiceAuthorizedProjectsResponseBody struct {
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
	Message     *string                                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	ProjectList []*GetDataServiceAuthorizedProjectsResponseBodyProjectList `json:"ProjectList,omitempty" xml:"ProjectList,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataServiceAuthorizedProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAuthorizedProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) GetProjectList() []*GetDataServiceAuthorizedProjectsResponseBodyProjectList {
	return s.ProjectList
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) SetCode(v string) *GetDataServiceAuthorizedProjectsResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) SetHttpStatusCode(v int32) *GetDataServiceAuthorizedProjectsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) SetMessage(v string) *GetDataServiceAuthorizedProjectsResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) SetProjectList(v []*GetDataServiceAuthorizedProjectsResponseBodyProjectList) *GetDataServiceAuthorizedProjectsResponseBody {
	s.ProjectList = v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) SetRequestId(v string) *GetDataServiceAuthorizedProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) SetSuccess(v bool) *GetDataServiceAuthorizedProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataServiceAuthorizedProjectsResponseBodyProjectList struct {
	// example:
	//
	// 1011
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// test
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// 2
	Role *int32 `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s GetDataServiceAuthorizedProjectsResponseBodyProjectList) String() string {
	return dara.Prettify(s)
}

func (s GetDataServiceAuthorizedProjectsResponseBodyProjectList) GoString() string {
	return s.String()
}

func (s *GetDataServiceAuthorizedProjectsResponseBodyProjectList) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GetDataServiceAuthorizedProjectsResponseBodyProjectList) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetDataServiceAuthorizedProjectsResponseBodyProjectList) GetRole() *int32 {
	return s.Role
}

func (s *GetDataServiceAuthorizedProjectsResponseBodyProjectList) SetProjectId(v int32) *GetDataServiceAuthorizedProjectsResponseBodyProjectList {
	s.ProjectId = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBodyProjectList) SetProjectName(v string) *GetDataServiceAuthorizedProjectsResponseBodyProjectList {
	s.ProjectName = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBodyProjectList) SetRole(v int32) *GetDataServiceAuthorizedProjectsResponseBodyProjectList {
	s.Role = &v
	return s
}

func (s *GetDataServiceAuthorizedProjectsResponseBodyProjectList) Validate() error {
	return dara.Validate(s)
}
