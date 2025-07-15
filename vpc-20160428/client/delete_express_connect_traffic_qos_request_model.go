// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExpressConnectTrafficQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteExpressConnectTrafficQosRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteExpressConnectTrafficQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteExpressConnectTrafficQosRequest
	GetOwnerId() *int64
	SetQosId(v string) *DeleteExpressConnectTrafficQosRequest
	GetQosId() *string
	SetRegionId(v string) *DeleteExpressConnectTrafficQosRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteExpressConnectTrafficQosRequest
	GetResourceOwnerAccount() *string
}

type DeleteExpressConnectTrafficQosRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The client token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The instance ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
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

func (s DeleteExpressConnectTrafficQosRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteExpressConnectTrafficQosRequest) GoString() string {
	return s.String()
}

func (s *DeleteExpressConnectTrafficQosRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteExpressConnectTrafficQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteExpressConnectTrafficQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteExpressConnectTrafficQosRequest) GetQosId() *string {
	return s.QosId
}

func (s *DeleteExpressConnectTrafficQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteExpressConnectTrafficQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteExpressConnectTrafficQosRequest) SetClientToken(v string) *DeleteExpressConnectTrafficQosRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRequest) SetOwnerAccount(v string) *DeleteExpressConnectTrafficQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRequest) SetOwnerId(v int64) *DeleteExpressConnectTrafficQosRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRequest) SetQosId(v string) *DeleteExpressConnectTrafficQosRequest {
	s.QosId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRequest) SetRegionId(v string) *DeleteExpressConnectTrafficQosRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRequest) SetResourceOwnerAccount(v string) *DeleteExpressConnectTrafficQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteExpressConnectTrafficQosRequest) Validate() error {
	return dara.Validate(s)
}
