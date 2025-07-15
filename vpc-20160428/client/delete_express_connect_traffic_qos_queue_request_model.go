// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteExpressConnectTrafficQosQueueRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteExpressConnectTrafficQosQueueRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteExpressConnectTrafficQosQueueRequest
	GetOwnerId() *int64
	SetQosId(v string) *DeleteExpressConnectTrafficQosQueueRequest
	GetQosId() *string
	SetQueueId(v string) *DeleteExpressConnectTrafficQosQueueRequest
	GetQueueId() *string
	SetRegionId(v string) *DeleteExpressConnectTrafficQosQueueRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteExpressConnectTrafficQosQueueRequest
	GetResourceOwnerAccount() *string
}

type DeleteExpressConnectTrafficQosQueueRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
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
}

func (s DeleteExpressConnectTrafficQosQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosQueueRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) GetQosId() *string {
	return s.QosId
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) GetQueueId() *string {
	return s.QueueId
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) SetClientToken(v string) *DeleteExpressConnectTrafficQosQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) SetOwnerAccount(v string) *DeleteExpressConnectTrafficQosQueueRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) SetOwnerId(v int64) *DeleteExpressConnectTrafficQosQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) SetQosId(v string) *DeleteExpressConnectTrafficQosQueueRequest {
	s.QosId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) SetQueueId(v string) *DeleteExpressConnectTrafficQosQueueRequest {
	s.QueueId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) SetRegionId(v string) *DeleteExpressConnectTrafficQosQueueRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) SetResourceOwnerAccount(v string) *DeleteExpressConnectTrafficQosQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosQueueRequest) Validate() error {
	return dara.Validate(s)
}
