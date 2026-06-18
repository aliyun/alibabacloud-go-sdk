// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentBySkillGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentBySkillGroupIdResponseBody
	GetCode() *string
	SetData(v []*ListAgentBySkillGroupIdResponseBodyData) *ListAgentBySkillGroupIdResponseBody
	GetData() []*ListAgentBySkillGroupIdResponseBodyData
	SetMessage(v string) *ListAgentBySkillGroupIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAgentBySkillGroupIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentBySkillGroupIdResponseBody
	GetSuccess() *bool
}

type ListAgentBySkillGroupIdResponseBody struct {
	// Status code. A return value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Agent information.
	Data []*ListAgentBySkillGroupIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// Description of the status code.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API invocation succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgentBySkillGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentBySkillGroupIdResponseBody) GetData() []*ListAgentBySkillGroupIdResponseBodyData {
	return s.Data
}

func (s *ListAgentBySkillGroupIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentBySkillGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentBySkillGroupIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentBySkillGroupIdResponseBody) SetCode(v string) *ListAgentBySkillGroupIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetData(v []*ListAgentBySkillGroupIdResponseBodyData) *ListAgentBySkillGroupIdResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetMessage(v string) *ListAgentBySkillGroupIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetRequestId(v string) *ListAgentBySkillGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) SetSuccess(v bool) *ListAgentBySkillGroupIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentBySkillGroupIdResponseBodyData struct {
	// Account name of the agent, which is the phone number or mailbox entered during account registration. It is unique within the instance.
	//
	// example:
	//
	// username@example.com
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// Agent ID.
	//
	// example:
	//
	// 666666
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// Display name of the agent.
	//
	// example:
	//
	// 刘测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Agent status. Only agents with a Normal status can perform Business Activities. Valid values:
	//
	// - **0**: Normal
	//
	// - **1**: Frozen
	//
	// - **2**: Deleted
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// Tenant ID to which the agent belongs. This corresponds to the instance ID provided in the input parameter.
	//
	// example:
	//
	// 0
	TenantId *int64 `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s ListAgentBySkillGroupIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentBySkillGroupIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentBySkillGroupIdResponseBodyData) GetAccountName() *string {
	return s.AccountName
}

func (s *ListAgentBySkillGroupIdResponseBodyData) GetAgentId() *int64 {
	return s.AgentId
}

func (s *ListAgentBySkillGroupIdResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListAgentBySkillGroupIdResponseBodyData) GetStatus() *int32 {
	return s.Status
}

func (s *ListAgentBySkillGroupIdResponseBodyData) GetTenantId() *int64 {
	return s.TenantId
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetAccountName(v string) *ListAgentBySkillGroupIdResponseBodyData {
	s.AccountName = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetAgentId(v int64) *ListAgentBySkillGroupIdResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetDisplayName(v string) *ListAgentBySkillGroupIdResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetStatus(v int32) *ListAgentBySkillGroupIdResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) SetTenantId(v int64) *ListAgentBySkillGroupIdResponseBodyData {
	s.TenantId = &v
	return s
}

func (s *ListAgentBySkillGroupIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
