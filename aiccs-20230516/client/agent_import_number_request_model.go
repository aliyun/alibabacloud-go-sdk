// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentImportNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentId(v int64) *AgentImportNumberRequest
	GetAgentId() *int64
	SetAgentTag(v string) *AgentImportNumberRequest
	GetAgentTag() *string
	SetCallType(v int64) *AgentImportNumberRequest
	GetCallType() *int64
	SetCustomers(v []*AgentImportNumberRequestCustomers) *AgentImportNumberRequest
	GetCustomers() []*AgentImportNumberRequestCustomers
	SetGatewayId(v int64) *AgentImportNumberRequest
	GetGatewayId() *int64
	SetOutId(v string) *AgentImportNumberRequest
	GetOutId() *string
	SetOwnerId(v int64) *AgentImportNumberRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *AgentImportNumberRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AgentImportNumberRequest
	GetResourceOwnerId() *int64
	SetTemplateId(v int64) *AgentImportNumberRequest
	GetTemplateId() *int64
}

type AgentImportNumberRequest struct {
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
	Customers []*AgentImportNumberRequestCustomers `json:"Customers,omitempty" xml:"Customers,omitempty" type:"Repeated"`
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

func (s AgentImportNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s AgentImportNumberRequest) GoString() string {
	return s.String()
}

func (s *AgentImportNumberRequest) GetAgentId() *int64 {
	return s.AgentId
}

func (s *AgentImportNumberRequest) GetAgentTag() *string {
	return s.AgentTag
}

func (s *AgentImportNumberRequest) GetCallType() *int64 {
	return s.CallType
}

func (s *AgentImportNumberRequest) GetCustomers() []*AgentImportNumberRequestCustomers {
	return s.Customers
}

func (s *AgentImportNumberRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *AgentImportNumberRequest) GetOutId() *string {
	return s.OutId
}

func (s *AgentImportNumberRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AgentImportNumberRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AgentImportNumberRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AgentImportNumberRequest) GetTemplateId() *int64 {
	return s.TemplateId
}

func (s *AgentImportNumberRequest) SetAgentId(v int64) *AgentImportNumberRequest {
	s.AgentId = &v
	return s
}

func (s *AgentImportNumberRequest) SetAgentTag(v string) *AgentImportNumberRequest {
	s.AgentTag = &v
	return s
}

func (s *AgentImportNumberRequest) SetCallType(v int64) *AgentImportNumberRequest {
	s.CallType = &v
	return s
}

func (s *AgentImportNumberRequest) SetCustomers(v []*AgentImportNumberRequestCustomers) *AgentImportNumberRequest {
	s.Customers = v
	return s
}

func (s *AgentImportNumberRequest) SetGatewayId(v int64) *AgentImportNumberRequest {
	s.GatewayId = &v
	return s
}

func (s *AgentImportNumberRequest) SetOutId(v string) *AgentImportNumberRequest {
	s.OutId = &v
	return s
}

func (s *AgentImportNumberRequest) SetOwnerId(v int64) *AgentImportNumberRequest {
	s.OwnerId = &v
	return s
}

func (s *AgentImportNumberRequest) SetResourceOwnerAccount(v string) *AgentImportNumberRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AgentImportNumberRequest) SetResourceOwnerId(v int64) *AgentImportNumberRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AgentImportNumberRequest) SetTemplateId(v int64) *AgentImportNumberRequest {
	s.TemplateId = &v
	return s
}

func (s *AgentImportNumberRequest) Validate() error {
	if s.Customers != nil {
		for _, item := range s.Customers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AgentImportNumberRequestCustomers struct {
	// 客户详情URL,用于展示客户用户在客户业务系统的详细信息，做多80个字符
	//
	// example:
	//
	// a
	ClientUrl *string `json:"ClientUrl,omitempty" xml:"ClientUrl,omitempty"`
	// 号码，如手机号等
	//
	// example:
	//
	// a
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// 用户电话号码MD5，和number两者必传一个
	//
	// example:
	//
	// a
	NumberMD5 *string `json:"NumberMD5,omitempty" xml:"NumberMD5,omitempty"`
	// 用户具体信息。如AI话术模板没变量要求或为人工外呼，可为空(总长度500个字符，多余的会被剔除)
	//
	// example:
	//
	// {"test":"234"}
	Properties map[string]interface{} `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// 用户自定义标签,最多50个字符
	//
	// example:
	//
	// a
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s AgentImportNumberRequestCustomers) String() string {
	return dara.Prettify(s)
}

func (s AgentImportNumberRequestCustomers) GoString() string {
	return s.String()
}

func (s *AgentImportNumberRequestCustomers) GetClientUrl() *string {
	return s.ClientUrl
}

func (s *AgentImportNumberRequestCustomers) GetNumber() *string {
	return s.Number
}

func (s *AgentImportNumberRequestCustomers) GetNumberMD5() *string {
	return s.NumberMD5
}

func (s *AgentImportNumberRequestCustomers) GetProperties() map[string]interface{} {
	return s.Properties
}

func (s *AgentImportNumberRequestCustomers) GetTag() *string {
	return s.Tag
}

func (s *AgentImportNumberRequestCustomers) SetClientUrl(v string) *AgentImportNumberRequestCustomers {
	s.ClientUrl = &v
	return s
}

func (s *AgentImportNumberRequestCustomers) SetNumber(v string) *AgentImportNumberRequestCustomers {
	s.Number = &v
	return s
}

func (s *AgentImportNumberRequestCustomers) SetNumberMD5(v string) *AgentImportNumberRequestCustomers {
	s.NumberMD5 = &v
	return s
}

func (s *AgentImportNumberRequestCustomers) SetProperties(v map[string]interface{}) *AgentImportNumberRequestCustomers {
	s.Properties = v
	return s
}

func (s *AgentImportNumberRequestCustomers) SetTag(v string) *AgentImportNumberRequestCustomers {
	s.Tag = &v
	return s
}

func (s *AgentImportNumberRequestCustomers) Validate() error {
	return dara.Validate(s)
}
