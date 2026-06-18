// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkspaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateWorkspaceResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *CreateWorkspaceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateWorkspaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateWorkspaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateWorkspaceResponseBody
	GetSuccess() *bool
	SetWorkspace(v *CreateWorkspaceResponseBodyWorkspace) *CreateWorkspaceResponseBody
	GetWorkspace() *CreateWorkspaceResponseBodyWorkspace
}

type CreateWorkspaceResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BB521414-5D38-5E66-AA66-963B2B4200E2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The business workspace information.
	Workspace *CreateWorkspaceResponseBodyWorkspace `json:"workspace,omitempty" xml:"workspace,omitempty" type:"Struct"`
}

func (s CreateWorkspaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateWorkspaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateWorkspaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateWorkspaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWorkspaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateWorkspaceResponseBody) GetWorkspace() *CreateWorkspaceResponseBodyWorkspace {
	return s.Workspace
}

func (s *CreateWorkspaceResponseBody) SetCode(v string) *CreateWorkspaceResponseBody {
	s.Code = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetHttpStatusCode(v int32) *CreateWorkspaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetMessage(v string) *CreateWorkspaceResponseBody {
	s.Message = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetRequestId(v string) *CreateWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetSuccess(v bool) *CreateWorkspaceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWorkspaceResponseBody) SetWorkspace(v *CreateWorkspaceResponseBodyWorkspace) *CreateWorkspaceResponseBody {
	s.Workspace = v
	return s
}

func (s *CreateWorkspaceResponseBody) Validate() error {
	if s.Workspace != nil {
		if err := s.Workspace.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateWorkspaceResponseBodyWorkspace struct {
	// API Host。
	//
	// example:
	//
	// llm-34o9ts1dai60z5sf.cn-beijing.maas.aliyuncs.com
	ApiHost *string `json:"apiHost,omitempty" xml:"apiHost,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1742785623772
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The service deployment scope. For more information, see [documentation](https://www.alibabacloud.com/help/zh/model-studio/regions/).
	//
	// example:
	//
	// global
	ServiceSite *string `json:"serviceSite,omitempty" xml:"serviceSite,omitempty"`
	// The business workspace ID.
	//
	// example:
	//
	// ws-ac3ef438bec22dc5
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// The name of the business workspace.
	//
	// example:
	//
	// default
	WorkspaceName *string `json:"workspaceName,omitempty" xml:"workspaceName,omitempty"`
}

func (s CreateWorkspaceResponseBodyWorkspace) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkspaceResponseBodyWorkspace) GoString() string {
	return s.String()
}

func (s *CreateWorkspaceResponseBodyWorkspace) GetApiHost() *string {
	return s.ApiHost
}

func (s *CreateWorkspaceResponseBodyWorkspace) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *CreateWorkspaceResponseBodyWorkspace) GetRegion() *string {
	return s.Region
}

func (s *CreateWorkspaceResponseBodyWorkspace) GetServiceSite() *string {
	return s.ServiceSite
}

func (s *CreateWorkspaceResponseBodyWorkspace) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateWorkspaceResponseBodyWorkspace) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetApiHost(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.ApiHost = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetGmtCreate(v int64) *CreateWorkspaceResponseBodyWorkspace {
	s.GmtCreate = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetRegion(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.Region = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetServiceSite(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.ServiceSite = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetWorkspaceId(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.WorkspaceId = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) SetWorkspaceName(v string) *CreateWorkspaceResponseBodyWorkspace {
	s.WorkspaceName = &v
	return s
}

func (s *CreateWorkspaceResponseBodyWorkspace) Validate() error {
	return dara.Validate(s)
}
