// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOrderDetails(v []*CreateFileSystemRequestOrderDetails) *CreateFileSystemRequest
	GetOrderDetails() []*CreateFileSystemRequestOrderDetails
}

type CreateFileSystemRequest struct {
	// The information about the orders.
	//
	// This parameter is required.
	OrderDetails []*CreateFileSystemRequestOrderDetails `json:"OrderDetails,omitempty" xml:"OrderDetails,omitempty" type:"Repeated"`
}

func (s CreateFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemRequest) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequest) GetOrderDetails() []*CreateFileSystemRequestOrderDetails {
	return s.OrderDetails
}

func (s *CreateFileSystemRequest) SetOrderDetails(v []*CreateFileSystemRequestOrderDetails) *CreateFileSystemRequest {
	s.OrderDetails = v
	return s
}

func (s *CreateFileSystemRequest) Validate() error {
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

type CreateFileSystemRequestOrderDetails struct {
	// The billing method of the NAS file system. Valid values:
	//
	// 	- PrePaid: subscription. This billing method is not supported.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// This parameter is required.
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The ID of the edge node.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-suzhou-telecom
	EnsRegionId *string `json:"EnsRegionId,omitempty" xml:"EnsRegionId,omitempty"`
	// The name of the file system. The name must be 1 to 80 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// nasFile
	FileSystemName *string `json:"FileSystemName,omitempty" xml:"FileSystemName,omitempty"`
	// The name of the mount target. The name must be 1 to 80 characters in length and can contain letters, digits, hyphens (-), and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// nasMount
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The ID of the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-xxx
	NetworkId *string `json:"NetworkId,omitempty" xml:"NetworkId,omitempty"`
	// The type of the order. Set the value to BUY.
	//
	// This parameter is required.
	//
	// example:
	//
	// BUY
	OrderType *string `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	// The storage protocol. Set the value to nfs.
	//
	// This parameter is required.
	//
	// example:
	//
	// nfs
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The storage type. Valid values:
	//
	// 	- Capacity.
	//
	// 	- Performance.
	//
	// This parameter is required.
	//
	// example:
	//
	// capacity
	StorgeType *string `json:"StorgeType,omitempty" xml:"StorgeType,omitempty"`
}

func (s CreateFileSystemRequestOrderDetails) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemRequestOrderDetails) GoString() string {
	return s.String()
}

func (s *CreateFileSystemRequestOrderDetails) GetChargeType() *string {
	return s.ChargeType
}

func (s *CreateFileSystemRequestOrderDetails) GetEnsRegionId() *string {
	return s.EnsRegionId
}

func (s *CreateFileSystemRequestOrderDetails) GetFileSystemName() *string {
	return s.FileSystemName
}

func (s *CreateFileSystemRequestOrderDetails) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *CreateFileSystemRequestOrderDetails) GetNetworkId() *string {
	return s.NetworkId
}

func (s *CreateFileSystemRequestOrderDetails) GetOrderType() *string {
	return s.OrderType
}

func (s *CreateFileSystemRequestOrderDetails) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *CreateFileSystemRequestOrderDetails) GetStorgeType() *string {
	return s.StorgeType
}

func (s *CreateFileSystemRequestOrderDetails) SetChargeType(v string) *CreateFileSystemRequestOrderDetails {
	s.ChargeType = &v
	return s
}

func (s *CreateFileSystemRequestOrderDetails) SetEnsRegionId(v string) *CreateFileSystemRequestOrderDetails {
	s.EnsRegionId = &v
	return s
}

func (s *CreateFileSystemRequestOrderDetails) SetFileSystemName(v string) *CreateFileSystemRequestOrderDetails {
	s.FileSystemName = &v
	return s
}

func (s *CreateFileSystemRequestOrderDetails) SetMountTargetDomain(v string) *CreateFileSystemRequestOrderDetails {
	s.MountTargetDomain = &v
	return s
}

func (s *CreateFileSystemRequestOrderDetails) SetNetworkId(v string) *CreateFileSystemRequestOrderDetails {
	s.NetworkId = &v
	return s
}

func (s *CreateFileSystemRequestOrderDetails) SetOrderType(v string) *CreateFileSystemRequestOrderDetails {
	s.OrderType = &v
	return s
}

func (s *CreateFileSystemRequestOrderDetails) SetProtocolType(v string) *CreateFileSystemRequestOrderDetails {
	s.ProtocolType = &v
	return s
}

func (s *CreateFileSystemRequestOrderDetails) SetStorgeType(v string) *CreateFileSystemRequestOrderDetails {
	s.StorgeType = &v
	return s
}

func (s *CreateFileSystemRequestOrderDetails) Validate() error {
	return dara.Validate(s)
}
