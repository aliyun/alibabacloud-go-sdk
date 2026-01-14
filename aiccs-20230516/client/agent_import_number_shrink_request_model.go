// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentImportNumberShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *AgentImportNumberShrinkRequest
	GetAgentId() *int64
	SetAgentTag(v string) *AgentImportNumberShrinkRequest
	GetAgentTag() *string
	SetCallType(v int64) *AgentImportNumberShrinkRequest
	GetCallType() *int64
	SetCustomersShrink(v string) *AgentImportNumberShrinkRequest
	GetCustomersShrink() *string
	SetGatewayId(v int64) *AgentImportNumberShrinkRequest
	GetGatewayId() *int64
	SetOutId(v string) *AgentImportNumberShrinkRequest
	GetOutId() *string
	SetOwnerId(v int64) *AgentImportNumberShrinkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AgentImportNumberShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentImportNumberShrinkRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int64) *AgentImportNumberShrinkRequest
	GetTemplateId() *int64
}

type AgentImportNumberShrinkRequest struct {
	// 坐席ID，可以为空，但agentId与agentTag必须其中1个有值。用于查询对应的坐席信息
	//
	// example:
	//
	// 1
	AgentId *int64 `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// 坐席唯一标识（创建坐席时的用户唯一标识），可以为空，但agentId与agentTag必须其中1个有值。用于查询对应的坐席信息
	//
	// example:
	//
	// 1001
	AgentTag *string `json:"AgentTag,omitempty" xml:"AgentTag,omitempty"`
	// 外呼类型 可选项：1001：坐席-人工外呼，1002：坐席-AI外呼-不转人工，1003：坐席-AI外呼-接通转人工，1004：坐席-AI外呼-智能转人工"
	//
	// This parameter is required.
	//
	// example:
	//
	// 79
	CallType *int64 `json:"CallType,omitempty" xml:"CallType,omitempty"`
	// 具体用户信息，如手机号、姓名等，需根据具体任务参数要求传输。注：当callType为1001时，只会导入第1条数据
	//
	// This parameter is required.
	CustomersShrink *string `json:"Customers,omitempty" xml:"Customers,omitempty"`
	// 坐席-人工外呼选择的外呼线路，只针对callType=1001生效，其他callType不生效。如callType=1001，且gatewayId不传值，默认按系统的线路配置外呼
	//
	// example:
	//
	// 1
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// 请求id，具有唯一性，用来做请求幂等处理，一个请求id有效期10分钟。不传则不做幂等处理
	//
	// example:
	//
	// 1asgfh
	OutId                *string `json:"OutId,omitempty" xml:"OutId,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// AI话术ID 客户已制作并已发布/平台授权的AI话术ID，如callType=1001可不填；如callType=1002,1003或1004 ，必填。
	//
	// example:
	//
	// 1
	TemplateId *int64 `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s AgentImportNumberShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentImportNumberShrinkRequest) GoString() string {
	return s.String()
}

func (s *AgentImportNumberShrinkRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentImportNumberShrinkRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentImportNumberShrinkRequest) GetCallType() *int64 {
	return s.CallType
}

func (s *AgentImportNumberShrinkRequest) GetCustomersShrink() *string {
	return s.CustomersShrink
}

func (s *AgentImportNumberShrinkRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *AgentImportNumberShrinkRequest) GetOutId() *string {
	return s.OutId
}

func (s *AgentImportNumberShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentImportNumberShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentImportNumberShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentImportNumberShrinkRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *AgentImportNumberShrinkRequest) SetAgentId(v int64) *AgentImportNumberShrinkRequest {
	s.AgentId = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetAgentTag(v string) *AgentImportNumberShrinkRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetCallType(v int64) *AgentImportNumberShrinkRequest {
	s.CallType = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetCustomersShrink(v string) *AgentImportNumberShrinkRequest {
	s.CustomersShrink = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetGatewayId(v int64) *AgentImportNumberShrinkRequest {
	s.GatewayId = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetOutId(v string) *AgentImportNumberShrinkRequest {
	s.OutId = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetOwnerId(v int64) *AgentImportNumberShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetResourceOwnerAccount(v string) *AgentImportNumberShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetResourceOwnerId(v int64) *AgentImportNumberShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) SetTemplateId(v int64) *AgentImportNumberShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *AgentImportNumberShrinkRequest) Validate() error {
	return dara.Validate(s)
}
