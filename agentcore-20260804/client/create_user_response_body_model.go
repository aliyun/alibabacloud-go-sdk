// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateUserResponseBody
	GetCode() *string
	SetData(v *CreateUserResponseBodyData) *CreateUserResponseBody
	GetData() *CreateUserResponseBodyData
	SetHttpStatusCode(v int32) *CreateUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUserResponseBody
	GetSuccess() *bool
}

type CreateUserResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *CreateUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Success   *bool   `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateUserResponseBody) GetData() *CreateUserResponseBodyData {
	return s.Data
}

func (s *CreateUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUserResponseBody) SetCode(v string) *CreateUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserResponseBody) SetData(v *CreateUserResponseBodyData) *CreateUserResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserResponseBody) SetHttpStatusCode(v int32) *CreateUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetSuccess(v bool) *CreateUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserResponseBodyData struct {
	// example:
	//
	// usr-123456
	AgentCoreUserId *string `json:"agentCoreUserId,omitempty" xml:"agentCoreUserId,omitempty"`
	// example:
	//
	// password
	AuthMethod *string `json:"authMethod,omitempty" xml:"authMethod,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// 张三
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// example:
	//
	// user-01@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// Example@2026
	InitialPassword *string `json:"initialPassword,omitempty" xml:"initialPassword,omitempty"`
	// example:
	//
	// user-01
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 智能体运营组成员
	Note *string `json:"note,omitempty" xml:"note,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// Active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s CreateUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyData) GetAgentCoreUserId() *string {
	return s.AgentCoreUserId
}

func (s *CreateUserResponseBodyData) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *CreateUserResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateUserResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *CreateUserResponseBodyData) GetInitialPassword() *string {
	return s.InitialPassword
}

func (s *CreateUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateUserResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *CreateUserResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateUserResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *CreateUserResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *CreateUserResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateUserResponseBodyData) SetAgentCoreUserId(v string) *CreateUserResponseBodyData {
	s.AgentCoreUserId = &v
	return s
}

func (s *CreateUserResponseBodyData) SetAuthMethod(v string) *CreateUserResponseBodyData {
	s.AuthMethod = &v
	return s
}

func (s *CreateUserResponseBodyData) SetCreatedAt(v string) *CreateUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *CreateUserResponseBodyData) SetDisplayName(v string) *CreateUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *CreateUserResponseBodyData) SetEmail(v string) *CreateUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *CreateUserResponseBodyData) SetInitialPassword(v string) *CreateUserResponseBodyData {
	s.InitialPassword = &v
	return s
}

func (s *CreateUserResponseBodyData) SetName(v string) *CreateUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateUserResponseBodyData) SetNote(v string) *CreateUserResponseBodyData {
	s.Note = &v
	return s
}

func (s *CreateUserResponseBodyData) SetRegionId(v string) *CreateUserResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *CreateUserResponseBodyData) SetStatus(v string) *CreateUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *CreateUserResponseBodyData) SetUpdatedAt(v string) *CreateUserResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *CreateUserResponseBodyData) SetWorkspaceId(v string) *CreateUserResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *CreateUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
