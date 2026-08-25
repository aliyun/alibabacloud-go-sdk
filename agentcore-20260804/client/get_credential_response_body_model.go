// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCredentialResponseBody
	GetCode() *string
	SetData(v *GetCredentialResponseBodyData) *GetCredentialResponseBody
	GetData() *GetCredentialResponseBodyData
	SetHttpStatusCode(v int32) *GetCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCredentialResponseBody
	GetSuccess() *bool
}

type GetCredentialResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                        `json:"code,omitempty" xml:"code,omitempty"`
	Data *GetCredentialResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCredentialResponseBody) GetData() *GetCredentialResponseBodyData {
	return s.Data
}

func (s *GetCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCredentialResponseBody) SetCode(v string) *GetCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetCredentialResponseBody) SetData(v *GetCredentialResponseBodyData) *GetCredentialResponseBody {
	s.Data = v
	return s
}

func (s *GetCredentialResponseBody) SetHttpStatusCode(v int32) *GetCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetCredentialResponseBody) SetMessage(v string) *GetCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetCredentialResponseBody) SetRequestId(v string) *GetCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCredentialResponseBody) SetSuccess(v bool) *GetCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *GetCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCredentialResponseBodyData struct {
	BoundAgents []*GetCredentialResponseBodyDataBoundAgents `json:"boundAgents,omitempty" xml:"boundAgents,omitempty" type:"Repeated"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// cred-123456
	CredentialId *string `json:"credentialId,omitempty" xml:"credentialId,omitempty"`
	// example:
	//
	// {"apiKey":"****************"}
	CredentialMetadata *string `json:"credentialMetadata,omitempty" xml:"credentialMetadata,omitempty"`
	// example:
	//
	// apiKey
	CredentialType *string `json:"credentialType,omitempty" xml:"credentialType,omitempty"`
	// example:
	//
	// 线上环境调用模型服务使用的 API Key
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// model-api-key
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 2026-08-12T03:04:05Z
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// example:
	//
	// ws-123456
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s GetCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyData) GetBoundAgents() []*GetCredentialResponseBodyDataBoundAgents {
	return s.BoundAgents
}

func (s *GetCredentialResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *GetCredentialResponseBodyData) GetCredentialId() *string {
	return s.CredentialId
}

func (s *GetCredentialResponseBodyData) GetCredentialMetadata() *string {
	return s.CredentialMetadata
}

func (s *GetCredentialResponseBodyData) GetCredentialType() *string {
	return s.CredentialType
}

func (s *GetCredentialResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetCredentialResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetCredentialResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCredentialResponseBodyData) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *GetCredentialResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetCredentialResponseBodyData) SetBoundAgents(v []*GetCredentialResponseBodyDataBoundAgents) *GetCredentialResponseBodyData {
	s.BoundAgents = v
	return s
}

func (s *GetCredentialResponseBodyData) SetCreatedAt(v string) *GetCredentialResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetCredentialId(v string) *GetCredentialResponseBodyData {
	s.CredentialId = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetCredentialMetadata(v string) *GetCredentialResponseBodyData {
	s.CredentialMetadata = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetCredentialType(v string) *GetCredentialResponseBodyData {
	s.CredentialType = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetDescription(v string) *GetCredentialResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetName(v string) *GetCredentialResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetRegionId(v string) *GetCredentialResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetUpdatedAt(v string) *GetCredentialResponseBodyData {
	s.UpdatedAt = &v
	return s
}

func (s *GetCredentialResponseBodyData) SetWorkspaceId(v string) *GetCredentialResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetCredentialResponseBodyData) Validate() error {
	if s.BoundAgents != nil {
		for _, item := range s.BoundAgents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCredentialResponseBodyDataBoundAgents struct {
	// example:
	//
	// agent-123456
	AgentId *string `json:"agentId,omitempty" xml:"agentId,omitempty"`
	// example:
	//
	// agent-01
	AgentName *string `json:"agentName,omitempty" xml:"agentName,omitempty"`
}

func (s GetCredentialResponseBodyDataBoundAgents) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialResponseBodyDataBoundAgents) GoString() string {
	return s.String()
}

func (s *GetCredentialResponseBodyDataBoundAgents) GetAgentId() *string {
	return s.AgentId
}

func (s *GetCredentialResponseBodyDataBoundAgents) GetAgentName() *string {
	return s.AgentName
}

func (s *GetCredentialResponseBodyDataBoundAgents) SetAgentId(v string) *GetCredentialResponseBodyDataBoundAgents {
	s.AgentId = &v
	return s
}

func (s *GetCredentialResponseBodyDataBoundAgents) SetAgentName(v string) *GetCredentialResponseBodyDataBoundAgents {
	s.AgentName = &v
	return s
}

func (s *GetCredentialResponseBodyDataBoundAgents) Validate() error {
	return dara.Validate(s)
}
