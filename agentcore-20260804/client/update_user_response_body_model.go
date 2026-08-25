// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUserResponseBody
	GetCode() *string
	SetData(v *UpdateUserResponseBodyData) *UpdateUserResponseBody
	GetData() *UpdateUserResponseBodyData
	SetHttpStatusCode(v int32) *UpdateUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUserResponseBody
	GetSuccess() *bool
}

type UpdateUserResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *UpdateUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s UpdateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUserResponseBody) GetData() *UpdateUserResponseBodyData {
	return s.Data
}

func (s *UpdateUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetData(v *UpdateUserResponseBodyData) *UpdateUserResponseBody {
	s.Data = v
	return s
}

func (s *UpdateUserResponseBody) SetHttpStatusCode(v int32) *UpdateUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserResponseBodyData struct {
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

func (s UpdateUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBodyData) GetAgentCoreUserId() *string {
	return s.AgentCoreUserId
}

func (s *UpdateUserResponseBodyData) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *UpdateUserResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateUserResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateUserResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *UpdateUserResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateUserResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UpdateUserResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *UpdateUserResponseBodyData) SetAgentCoreUserId(v string) *UpdateUserResponseBodyData {
	s.AgentCoreUserId = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetAuthMethod(v string) *UpdateUserResponseBodyData {
	s.AuthMethod = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetCreatedAt(v string) *UpdateUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetDisplayName(v string) *UpdateUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetEmail(v string) *UpdateUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetName(v string) *UpdateUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetNote(v string) *UpdateUserResponseBodyData {
	s.Note = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetRegionId(v string) *UpdateUserResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetStatus(v string) *UpdateUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetUpdatedAt(v string) *UpdateUserResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetWorkspaceId(v string) *UpdateUserResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
