// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAllowedIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedListIp(v string) *UpdateAllowedIpRequest
	GetAllowedListIp() *string
	SetAllowedListType(v string) *UpdateAllowedIpRequest
	GetAllowedListType() *string
	SetDescription(v string) *UpdateAllowedIpRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateAllowedIpRequest
	GetInstanceId() *string
	SetPortRange(v string) *UpdateAllowedIpRequest
	GetPortRange() *string
	SetRegionId(v string) *UpdateAllowedIpRequest
	GetRegionId() *string
	SetUpdateType(v string) *UpdateAllowedIpRequest
	GetUpdateType() *string
}

type UpdateAllowedIpRequest struct {
	// The IP addresses that you want to manage. You can specify a CIDR block. Example: **192.168.0.0/16**.
	//
	// 	- If the **UpdateType*	- parameter is set to **add**, specify one or more IP addresses for this parameter. Separate multiple IP addresses with commas (,).
	//
	// 	- If the **UpdateType*	- parameter is set to **delete**, specify only one IP address.
	//
	// 	- Exercise caution when you delete IP addresses.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0.0.0.0/0
	AllowedListIp *string `json:"AllowedListIp,omitempty" xml:"AllowedListIp,omitempty"`
	// The type of the whitelist. Valid values:
	//
	// 	- **vpc**: a whitelist for access from a VPC.
	//
	// 	- **internet**: a whitelist for access from the Internet.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc
	AllowedListType *string `json:"AllowedListType,omitempty" xml:"AllowedListType,omitempty"`
	// The description of the whitelist.
	//
	// example:
	//
	// tf-testAccEcsImageConfigBasic3549descriptionChange
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-0pp1cng20***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The port range. Valid values:
	//
	// 	- **9092/9092**: Messages are transmitted in a virtual private cloud (VPC) by using the PLAINTEXT protocol.
	//
	// 	- **9093/9093**: Messages are transmitted over the Internet by using the SASL_SSL protocol.
	//
	// 	- **9094/9094**: Messages are transmitted in a VPC by using the SASL_PLAINTEXT protocol.
	//
	// 	- **9095/9095**: Messages are transmitted in a VPC by using the SASL_SSL protocol.
	//
	// This parameter must correspond to **AllowdedListType**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9092/9092
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the region where the instance resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of configuration change. Valid values:
	//
	// 	- **add**
	//
	// 	- **delete**
	//
	// This parameter is required.
	//
	// example:
	//
	// add
	UpdateType *string `json:"UpdateType,omitempty" xml:"UpdateType,omitempty"`
}

func (s UpdateAllowedIpRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAllowedIpRequest) GoString() string {
	return s.String()
}

func (s *UpdateAllowedIpRequest) GetAllowedListIp() *string {
	return s.AllowedListIp
}

func (s *UpdateAllowedIpRequest) GetAllowedListType() *string {
	return s.AllowedListType
}

func (s *UpdateAllowedIpRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateAllowedIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAllowedIpRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *UpdateAllowedIpRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateAllowedIpRequest) GetUpdateType() *string {
	return s.UpdateType
}

func (s *UpdateAllowedIpRequest) SetAllowedListIp(v string) *UpdateAllowedIpRequest {
	s.AllowedListIp = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetAllowedListType(v string) *UpdateAllowedIpRequest {
	s.AllowedListType = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetDescription(v string) *UpdateAllowedIpRequest {
	s.Description = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetInstanceId(v string) *UpdateAllowedIpRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetPortRange(v string) *UpdateAllowedIpRequest {
	s.PortRange = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetRegionId(v string) *UpdateAllowedIpRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateAllowedIpRequest) SetUpdateType(v string) *UpdateAllowedIpRequest {
	s.UpdateType = &v
	return s
}

func (s *UpdateAllowedIpRequest) Validate() error {
	return dara.Validate(s)
}
