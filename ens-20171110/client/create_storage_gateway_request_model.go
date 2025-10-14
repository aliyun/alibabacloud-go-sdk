// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageGatewayRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderDetails(v []*CreateStorageGatewayRequestOrderDetails) *CreateStorageGatewayRequest
	GetOrderDetails() []*CreateStorageGatewayRequestOrderDetails
}

type CreateStorageGatewayRequest struct {
	// The array of orders.
	//
	// This parameter is required.
	OrderDetails []*CreateStorageGatewayRequestOrderDetails `json:"OrderDetails,omitempty" xml:"OrderDetails,omitempty" type:"Repeated"`
}

func (s CreateStorageGatewayRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageGatewayRequest) GoString() string {
	return s.String()
}

func (s *CreateStorageGatewayRequest) GetOrderDetails() []*CreateStorageGatewayRequestOrderDetails {
	return s.OrderDetails
}

func (s *CreateStorageGatewayRequest) SetOrderDetails(v []*CreateStorageGatewayRequestOrderDetails) *CreateStorageGatewayRequest {
	s.OrderDetails = v
	return s
}

func (s *CreateStorageGatewayRequest) Validate() error {
	if s.OrderDetails != nil {
		for _, item := range s.OrderDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateStorageGatewayRequestOrderDetails struct {
	// The description of the gateway. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen-3
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the gateway. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (.), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testGatewayName
	GatewayName *string `json:"GatewayName,omitempty" xml:"GatewayName,omitempty"`
	// The type of the gateway. Set this parameter to **1**. **1*	- indicates iSCSI.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	GatewayType *string `json:"GatewayType,omitempty" xml:"GatewayType,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// n-123
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateStorageGatewayRequestOrderDetails) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageGatewayRequestOrderDetails) GoString() string {
	return s.String()
}

func (s *CreateStorageGatewayRequestOrderDetails) GetDescription() *string {
	return s.Description
}

func (s *CreateStorageGatewayRequestOrderDetails) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateStorageGatewayRequestOrderDetails) GetGatewayName() *string {
	return s.GatewayName
}

func (s *CreateStorageGatewayRequestOrderDetails) GetGatewayType() *string {
	return s.GatewayType
}

func (s *CreateStorageGatewayRequestOrderDetails) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateStorageGatewayRequestOrderDetails) SetDescription(v string) *CreateStorageGatewayRequestOrderDetails {
	s.Description = &v
	return s
}

func (s *CreateStorageGatewayRequestOrderDetails) SetEnsRegionId(v string) *CreateStorageGatewayRequestOrderDetails {
	s.EnsRegionId = &v
	return s
}

func (s *CreateStorageGatewayRequestOrderDetails) SetGatewayName(v string) *CreateStorageGatewayRequestOrderDetails {
	s.GatewayName = &v
	return s
}

func (s *CreateStorageGatewayRequestOrderDetails) SetGatewayType(v string) *CreateStorageGatewayRequestOrderDetails {
	s.GatewayType = &v
	return s
}

func (s *CreateStorageGatewayRequestOrderDetails) SetVpcId(v string) *CreateStorageGatewayRequestOrderDetails {
	s.VpcId = &v
	return s
}

func (s *CreateStorageGatewayRequestOrderDetails) Validate() error {
	return dara.Validate(s)
}
