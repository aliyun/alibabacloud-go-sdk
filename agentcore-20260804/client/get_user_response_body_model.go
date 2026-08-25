// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserResponseBody
	GetCode() *string
	SetData(v *GetUserResponseBodyData) *GetUserResponseBody
	GetData() *GetUserResponseBodyData
	SetHttpStatusCode(v int32) *GetUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserResponseBody
	GetSuccess() *bool
}

type GetUserResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                  `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserResponseBody) GetData() *GetUserResponseBodyData {
	return s.Data
}

func (s *GetUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserResponseBody) SetCode(v string) *GetUserResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserResponseBody) SetData(v *GetUserResponseBodyData) *GetUserResponseBody {
	s.Data = v
	return s
}

func (s *GetUserResponseBody) SetHttpStatusCode(v int32) *GetUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserResponseBody) SetMessage(v string) *GetUserResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserResponseBody) SetRequestId(v string) *GetUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserResponseBody) SetSuccess(v bool) *GetUserResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserResponseBodyData struct {
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

func (s GetUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserResponseBodyData) GetAgentCoreUserId() *string {
	return s.AgentCoreUserId
}

func (s *GetUserResponseBodyData) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *GetUserResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetUserResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *GetUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetUserResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *GetUserResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUserResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetUserResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetUserResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetUserResponseBodyData) SetAgentCoreUserId(v string) *GetUserResponseBodyData {
	s.AgentCoreUserId = &v
	return s
}

func (s *GetUserResponseBodyData) SetAuthMethod(v string) *GetUserResponseBodyData {
	s.AuthMethod = &v
	return s
}

func (s *GetUserResponseBodyData) SetCreatedAt(v string) *GetUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetUserResponseBodyData) SetDisplayName(v string) *GetUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *GetUserResponseBodyData) SetEmail(v string) *GetUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *GetUserResponseBodyData) SetName(v string) *GetUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetUserResponseBodyData) SetNote(v string) *GetUserResponseBodyData {
	s.Note = &v
	return s
}

func (s *GetUserResponseBodyData) SetRegionId(v string) *GetUserResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetUserResponseBodyData) SetStatus(v string) *GetUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetUserResponseBodyData) SetUpdatedAt(v string) *GetUserResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetUserResponseBodyData) SetWorkspaceId(v string) *GetUserResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
