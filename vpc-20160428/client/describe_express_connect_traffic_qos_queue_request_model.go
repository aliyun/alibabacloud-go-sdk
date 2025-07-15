// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressConnectTrafficQosQueueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeExpressConnectTrafficQosQueueRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DescribeExpressConnectTrafficQosQueueRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeExpressConnectTrafficQosQueueRequest
	GetOwnerId() *int64
	SetQosId(v string) *DescribeExpressConnectTrafficQosQueueRequest
	GetQosId() *string
	SetQueueIdList(v []*string) *DescribeExpressConnectTrafficQosQueueRequest
	GetQueueIdList() []*string
	SetQueueNameList(v []*string) *DescribeExpressConnectTrafficQosQueueRequest
	GetQueueNameList() []*string
	SetRegionId(v string) *DescribeExpressConnectTrafficQosQueueRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeExpressConnectTrafficQosQueueRequest
	GetResourceOwnerAccount() *string
}

type DescribeExpressConnectTrafficQosQueueRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the QoS policy.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The IDs of the QoS queues.
	QueueIdList []*string `json:"QueueIdList,omitempty" xml:"QueueIdList,omitempty" type:"Repeated"`
	// The names of the QoS queues.
	QueueNameList []*string `json:"QueueNameList,omitempty" xml:"QueueNameList,omitempty" type:"Repeated"`
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

func (s DescribeExpressConnectTrafficQosQueueRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressConnectTrafficQosQueueRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) GetQosId() *string {
	return s.QosId
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) GetQueueIdList() []*string {
	return s.QueueIdList
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) GetQueueNameList() []*string {
	return s.QueueNameList
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) SetClientToken(v string) *DescribeExpressConnectTrafficQosQueueRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) SetOwnerAccount(v string) *DescribeExpressConnectTrafficQosQueueRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) SetOwnerId(v int64) *DescribeExpressConnectTrafficQosQueueRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) SetQosId(v string) *DescribeExpressConnectTrafficQosQueueRequest {
	s.QosId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) SetQueueIdList(v []*string) *DescribeExpressConnectTrafficQosQueueRequest {
	s.QueueIdList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) SetQueueNameList(v []*string) *DescribeExpressConnectTrafficQosQueueRequest {
	s.QueueNameList = v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) SetRegionId(v string) *DescribeExpressConnectTrafficQosQueueRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) SetResourceOwnerAccount(v string) *DescribeExpressConnectTrafficQosQueueRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeExpressConnectTrafficQosQueueRequest) Validate() error {
	return dara.Validate(s)
}
