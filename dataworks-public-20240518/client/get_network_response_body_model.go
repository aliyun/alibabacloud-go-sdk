// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetwork(v *GetNetworkResponseBodyNetwork) *GetNetworkResponseBody
	GetNetwork() *GetNetworkResponseBodyNetwork
	SetRequestId(v string) *GetNetworkResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetNetworkResponseBody
	GetSuccess() *bool
}

type GetNetworkResponseBody struct {
	// The information about the network resource.
	Network *GetNetworkResponseBodyNetwork `json:"Network,omitempty" xml:"Network,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *GetNetworkResponseBody) GetNetwork() *GetNetworkResponseBodyNetwork {
	return s.Network
}

func (s *GetNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNetworkResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetNetworkResponseBody) SetNetwork(v *GetNetworkResponseBodyNetwork) *GetNetworkResponseBody {
	s.Network = v
	return s
}

func (s *GetNetworkResponseBody) SetRequestId(v string) *GetNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNetworkResponseBody) SetSuccess(v bool) *GetNetworkResponseBody {
	s.Success = &v
	return s
}

func (s *GetNetworkResponseBody) Validate() error {
	if s.Network != nil {
		if err := s.Network.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetNetworkResponseBodyNetwork struct {
	// The time when the network resource was created. The value is a 64-bit timestamp.
	//
	// example:
	//
	// 1727055811000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the user who creates the network resource.
	//
	// example:
	//
	// 11075500042XXXXX
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The network ID.
	//
	// example:
	//
	// 1000
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the serverless resource group.
	//
	// example:
	//
	// Serverless_res_group_524257424564736_6831777003XXXXX
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-2ze13vamugr7jenXXXXX
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The status of the network resource. Valid values:
	//
	// 	- Pending: The network resource is waiting to be created.
	//
	// 	- Creating: The network resource is being created.
	//
	// 	- Running: The network resource is running as expected.
	//
	// 	- Deleting: The network resource is being deleted.
	//
	// 	- Deleted: The network resource is deleted.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-m2et4f3oc8msfbccXXXXX
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The VSwitch ID.
	//
	// example:
	//
	// vsw-uf8usrhs7hjd9amsXXXXX
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s GetNetworkResponseBodyNetwork) String() string {
	return dara.Prettify(s)
}

func (s GetNetworkResponseBodyNetwork) GoString() string {
	return s.String()
}

func (s *GetNetworkResponseBodyNetwork) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetNetworkResponseBodyNetwork) GetCreateUser() *string {
	return s.CreateUser
}

func (s *GetNetworkResponseBodyNetwork) GetId() *int64 {
	return s.Id
}

func (s *GetNetworkResponseBodyNetwork) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetNetworkResponseBodyNetwork) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetNetworkResponseBodyNetwork) GetStatus() *string {
	return s.Status
}

func (s *GetNetworkResponseBodyNetwork) GetVpcId() *string {
	return s.VpcId
}

func (s *GetNetworkResponseBodyNetwork) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetNetworkResponseBodyNetwork) SetCreateTime(v int64) *GetNetworkResponseBodyNetwork {
	s.CreateTime = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetCreateUser(v string) *GetNetworkResponseBodyNetwork {
	s.CreateUser = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetId(v int64) *GetNetworkResponseBodyNetwork {
	s.Id = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetResourceGroupId(v string) *GetNetworkResponseBodyNetwork {
	s.ResourceGroupId = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetSecurityGroupId(v string) *GetNetworkResponseBodyNetwork {
	s.SecurityGroupId = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetStatus(v string) *GetNetworkResponseBodyNetwork {
	s.Status = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetVpcId(v string) *GetNetworkResponseBodyNetwork {
	s.VpcId = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) SetVswitchId(v string) *GetNetworkResponseBodyNetwork {
	s.VswitchId = &v
	return s
}

func (s *GetNetworkResponseBodyNetwork) Validate() error {
	return dara.Validate(s)
}
