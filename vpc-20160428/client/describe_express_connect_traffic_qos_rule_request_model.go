// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeExpressConnectTrafficQosRuleRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DescribeExpressConnectTrafficQosRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeExpressConnectTrafficQosRuleRequest
	GetOwnerId() *int64
	SetQosId(v string) *DescribeExpressConnectTrafficQosRuleRequest
	GetQosId() *string
	SetQueueId(v string) *DescribeExpressConnectTrafficQosRuleRequest
	GetQueueId() *string
	SetRegionId(v string) *DescribeExpressConnectTrafficQosRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeExpressConnectTrafficQosRuleRequest
	GetResourceOwnerAccount() *string
	SetRuleIdList(v []*string) *DescribeExpressConnectTrafficQosRuleRequest
	GetRuleIdList() []*string
	SetRuleNameList(v []*string) *DescribeExpressConnectTrafficQosRuleRequest
	GetRuleNameList() []*string
}

type DescribeExpressConnectTrafficQosRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The ID of the QoS queue.
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
	// The list of QoS rule IDs.
	RuleIdList []*string `json:"RuleIdList,omitempty" xml:"RuleIdList,omitempty" type:"Repeated"`
	// The list of QoS rule names.
	RuleNameList []*string `json:"RuleNameList,omitempty" xml:"RuleNameList,omitempty" type:"Repeated"`
}

func (s DescribeExpressConnectTrafficQosRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosRuleRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetQosId() *string {
	return s.QosId
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetRuleIdList() []*string {
	return s.RuleIdList
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) GetRuleNameList() []*string {
	return s.RuleNameList
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetClientToken(v string) *DescribeExpressConnectTrafficQosRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetOwnerAccount(v string) *DescribeExpressConnectTrafficQosRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetOwnerId(v int64) *DescribeExpressConnectTrafficQosRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetQosId(v string) *DescribeExpressConnectTrafficQosRuleRequest {
	s.QosId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetQueueId(v string) *DescribeExpressConnectTrafficQosRuleRequest {
	s.QueueId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetRegionId(v string) *DescribeExpressConnectTrafficQosRuleRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetResourceOwnerAccount(v string) *DescribeExpressConnectTrafficQosRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetRuleIdList(v []*string) *DescribeExpressConnectTrafficQosRuleRequest {
	s.RuleIdList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) SetRuleNameList(v []*string) *DescribeExpressConnectTrafficQosRuleRequest {
	s.RuleNameList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosRuleRequest) Validate() error {
	return dara.Validate(s)
}
