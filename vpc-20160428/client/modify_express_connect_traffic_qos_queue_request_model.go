// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPercent(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetBandwidthPercent() *string
	SetClientToken(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyExpressConnectTrafficQosQueueRequest
	GetOwnerId() *int64
	SetQosId(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetQosId() *string
	SetQueueDescription(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetQueueDescription() *string
	SetQueueId(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetQueueId() *string
	SetQueueName(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetQueueName() *string
	SetRegionId(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyExpressConnectTrafficQosQueueRequest
	GetResourceOwnerAccount() *string
}

type ModifyExpressConnectTrafficQosQueueRequest struct {
	// The bandwidth percentage of the QoS queue.
	//
	// 	- If the QoS queue type is **Medium**, this parameter is required. Valid values: 1 to 100.
	//
	// 	- If the QoS queue type is **Default**, this parameter is set to "-".
	//
	// example:
	//
	// 100
	BandwidthPercent *string `json:"BandwidthPercent,omitempty" xml:"BandwidthPercent,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the **RequestId*	- of the API request as the **ClientToken**. The **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv****
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The description of the QoS queue.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-queue-test
	QueueDescription *string `json:"QueueDescription,omitempty" xml:"QueueDescription,omitempty"`
	// The ID of the QoS queue.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-queue-9nyx2u7n71s2rc****
	QueueId *string `json:"QueueId,omitempty" xml:"QueueId,omitempty"`
	// The name of the QoS queue.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-queue-test
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The region ID of the QoS policy.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s ModifyExpressConnectTrafficQosQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosQueueRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetBandwidthPercent() *string {
	return s.BandwidthPercent
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetQosId() *string {
	return s.QosId
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetQueueDescription() *string {
	return s.QueueDescription
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetBandwidthPercent(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.BandwidthPercent = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetClientToken(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetOwnerAccount(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetOwnerId(v int64) *ModifyExpressConnectTrafficQosQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetQosId(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.QosId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetQueueDescription(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.QueueDescription = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetQueueId(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.QueueId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetQueueName(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.QueueName = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetRegionId(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) SetResourceOwnerAccount(v string) *ModifyExpressConnectTrafficQosQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosQueueRequest) Validate() error {
	return dara.Validate(s)
}
