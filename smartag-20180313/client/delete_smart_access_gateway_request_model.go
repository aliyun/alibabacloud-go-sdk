// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartAccessGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSmartAccessGatewayRequest
	GetInstanceId() *string
	SetRegionId(v string) *DeleteSmartAccessGatewayRequest
	GetRegionId() *string
	SetResourceOwnerId(v int64) *DeleteSmartAccessGatewayRequest
	GetResourceOwnerId() *int64
}

type DeleteSmartAccessGatewayRequest struct {
	// The ID of the SAG instance that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-far8v6owtdxlua****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteSmartAccessGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartAccessGatewayRequest) GoString() string {
	return s.String()
}

func (s *DeleteSmartAccessGatewayRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSmartAccessGatewayRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteSmartAccessGatewayRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSmartAccessGatewayRequest) SetInstanceId(v string) *DeleteSmartAccessGatewayRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSmartAccessGatewayRequest) SetRegionId(v string) *DeleteSmartAccessGatewayRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteSmartAccessGatewayRequest) SetResourceOwnerId(v int64) *DeleteSmartAccessGatewayRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSmartAccessGatewayRequest) Validate() error {
	return dara.Validate(s)
}
