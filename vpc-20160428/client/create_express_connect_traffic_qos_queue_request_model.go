// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExpressConnectTrafficQosQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPercent(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetBandwidthPercent() *string
	SetClientToken(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateExpressConnectTrafficQosQueueRequest
	GetOwnerId() *int64
	SetQosId(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetQosId() *string
	SetQueueDescription(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetQueueDescription() *string
	SetQueueName(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetQueueName() *string
	SetQueueType(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetQueueType() *string
	SetRegionId(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateExpressConnectTrafficQosQueueRequest
	GetResourceOwnerAccount() *string
}

type CreateExpressConnectTrafficQosQueueRequest struct {
	// The percentage of bandwidth allocated to the QoS queue.
	//
	// 	- If QueueType is set to **Medium**, this parameter is required. Valid values: 1 to 100.
	//
	// 	- If QueueType is set to **Default**, a value of - is returned.
	//
	// example:
	//
	// 100
	BandwidthPercent *string `json:"BandwidthPercent,omitempty" xml:"BandwidthPercent,omitempty"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among all requests. ClientToken can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, **ClientToken*	- is set to the value of **RequestId**. The value of **RequestId*	- for each API request may be different.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
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
	// The description of the QoS queue.
	//
	// It must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-queue-test
	QueueDescription *string `json:"QueueDescription,omitempty" xml:"QueueDescription,omitempty"`
	// The name of the QoS queue.
	//
	// It must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// qos-queue-test
	QueueName *string `json:"QueueName,omitempty" xml:"QueueName,omitempty"`
	// The priority of the QoS queue. Valid values:
	//
	// 	- **High**
	//
	// 	- **Medium**
	//
	// 	- **Default**: default queue.
	//
	// > You cannot create a QoS queue of the default priority.
	//
	// This parameter is required.
	//
	// example:
	//
	// High
	QueueType *string `json:"QueueType,omitempty" xml:"QueueType,omitempty"`
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
}

func (s CreateExpressConnectTrafficQosQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateExpressConnectTrafficQosQueueRequest) GoString() string {
	return s.String()
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetBandwidthPercent() *string {
	return s.BandwidthPercent
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetQosId() *string {
	return s.QosId
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetQueueDescription() *string {
	return s.QueueDescription
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetQueueName() *string {
	return s.QueueName
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetQueueType() *string {
	return s.QueueType
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateExpressConnectTrafficQosQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetBandwidthPercent(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.BandwidthPercent = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetClientToken(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetOwnerAccount(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetOwnerId(v int64) *CreateExpressConnectTrafficQosQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetQosId(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.QosId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetQueueDescription(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.QueueDescription = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetQueueName(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.QueueName = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetQueueType(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.QueueType = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetRegionId(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.RegionId = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) SetResourceOwnerAccount(v string) *CreateExpressConnectTrafficQosQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateExpressConnectTrafficQosQueueRequest) Validate() error {
	return dara.Validate(s)
}
