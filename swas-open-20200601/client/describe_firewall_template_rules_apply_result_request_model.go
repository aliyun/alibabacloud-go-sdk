// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFirewallTemplateRulesApplyResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeFirewallTemplateRulesApplyResultRequest
	GetClientToken() *string
	SetFirewallTemplateId(v string) *DescribeFirewallTemplateRulesApplyResultRequest
	GetFirewallTemplateId() *string
	SetInstanceId(v string) *DescribeFirewallTemplateRulesApplyResultRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeFirewallTemplateRulesApplyResultRequest
	GetRegionId() *string
	SetTaskId(v string) *DescribeFirewallTemplateRulesApplyResultRequest
	GetTaskId() *string
}

type DescribeFirewallTemplateRulesApplyResultRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the firewall template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ft-bcf1a7hrdq717****
	FirewallTemplateId *string `json:"FirewallTemplateId,omitempty" xml:"FirewallTemplateId,omitempty"`
	// The ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID of the simple application server. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the execution to apply the template rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// aft-ikgt0bynitvs3****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeFirewallTemplateRulesApplyResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFirewallTemplateRulesApplyResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) GetFirewallTemplateId() *string {
	return s.FirewallTemplateId
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetClientToken(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetFirewallTemplateId(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.FirewallTemplateId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetInstanceId(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetRegionId(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) SetTaskId(v string) *DescribeFirewallTemplateRulesApplyResultRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeFirewallTemplateRulesApplyResultRequest) Validate() error {
	return dara.Validate(s)
}
