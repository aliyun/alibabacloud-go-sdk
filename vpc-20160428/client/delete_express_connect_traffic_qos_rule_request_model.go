// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteExpressConnectTrafficQosRuleRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteExpressConnectTrafficQosRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteExpressConnectTrafficQosRuleRequest
	GetOwnerId() *int64
	SetQosId(v string) *DeleteExpressConnectTrafficQosRuleRequest
	GetQosId() *string
	SetQueueId(v string) *DeleteExpressConnectTrafficQosRuleRequest
	GetQueueId() *string
	SetRegionId(v string) *DeleteExpressConnectTrafficQosRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteExpressConnectTrafficQosRuleRequest
	GetResourceOwnerAccount() *string
	SetRuleId(v string) *DeleteExpressConnectTrafficQosRuleRequest
	GetRuleId() *string
}

type DeleteExpressConnectTrafficQosRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-00****
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the QoS queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rcy4n5
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The region ID of the QoS policy.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// The ID of the QoS rule.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-rule-iugg0l9x27f2nocouj
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s DeleteExpressConnectTrafficQosRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) GetQosId() *string {
	return s.QosId
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) GetRuleId() *string {
	return s.RuleId
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) SetClientToken(v string) *DeleteExpressConnectTrafficQosRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) SetOwnerAccount(v string) *DeleteExpressConnectTrafficQosRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) SetOwnerId(v int64) *DeleteExpressConnectTrafficQosRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) SetQosId(v string) *DeleteExpressConnectTrafficQosRuleRequest {
	s.QosId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) SetQueueId(v string) *DeleteExpressConnectTrafficQosRuleRequest {
	s.QueueId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) SetRegionId(v string) *DeleteExpressConnectTrafficQosRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) SetResourceOwnerAccount(v string) *DeleteExpressConnectTrafficQosRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) SetRuleId(v string) *DeleteExpressConnectTrafficQosRuleRequest {
	s.RuleId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRuleRequest) Validate() error {
	return dara.Validate(s)
}
