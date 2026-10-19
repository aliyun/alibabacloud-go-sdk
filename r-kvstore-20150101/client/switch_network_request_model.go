// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClassicExpiredDays(v string) *SwitchNetworkRequest
	GetClassicExpiredDays() *string
	SetInstanceId(v string) *SwitchNetworkRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *SwitchNetworkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchNetworkRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SwitchNetworkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchNetworkRequest
	GetResourceOwnerId() *int64
	SetRetainClassic(v string) *SwitchNetworkRequest
	GetRetainClassic() *string
	SetSecurityToken(v string) *SwitchNetworkRequest
	GetSecurityToken() *string
	SetTargetNetworkType(v string) *SwitchNetworkRequest
	GetTargetNetworkType() *string
	SetVSwitchId(v string) *SwitchNetworkRequest
	GetVSwitchId() *string
	SetVpcId(v string) *SwitchNetworkRequest
	GetVpcId() *string
}

type SwitchNetworkRequest struct {
	// The retention period of the classic network connectivity endpoint. Valid values: **14**, **30**, **60**, and **120**. Unit: days.
	//
	// > 	- This parameter is active and required only when **RetainClassic*	- is set to **True**.
	//
	// > 	- After the switchover is complete, you can invoke [ModifyInstanceNetExpireTime](https://help.aliyun.com/document_detail/473793.html) to modify the retention period of the classic network connectivity endpoint.
	//
	// example:
	//
	// 30
	ClassicExpiredDays *string `json:"ClassicExpiredDays,omitempty" xml:"ClassicExpiredDays,omitempty"`
	// The instance ID. You can call [DescribeInstances](https://help.aliyun.com/document_detail/473778.html) to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to retain the classic network connectivity endpoint when you switch the instance from the classic network to a VPC. Valid values:
	//
	// 	- **True**: The classic network endpoint is retained.
	//
	// 	- **False**: The classic network endpoint is not retained. This is the default value.
	//
	// > This parameter is active only when the network type of the instance is classic network.
	//
	// example:
	//
	// True
	RetainClassic *string `json:"RetainClassic,omitempty" xml:"RetainClassic,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The network type to which you want to switch. Set the value to **VPC**.
	//
	// example:
	//
	// VPC
	TargetNetworkType *string `json:"TargetNetworkType,omitempty" xml:"TargetNetworkType,omitempty"`
	// The ID of the vSwitch in the destination VPC. You can call [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) to query vSwitch IDs.
	//
	// > The vSwitch must be in the same zone as the instance.
	//
	// example:
	//
	// vsw-bp1e7clcw529l773d****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the VPC to which you want to switch. You can call [DescribeVpcs](https://help.aliyun.com/document_detail/35739.html) to query VPC IDs.
	//
	// > 	- The VPC must be in the same region as the instance.
	//
	// > 	- After you specify this parameter, you must also specify the **VSwitchId*	- parameter.
	//
	// example:
	//
	// vpc-bp1nme44gek34slfc****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SwitchNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchNetworkRequest) GoString() string {
	return s.String()
}

func (s *SwitchNetworkRequest) GetClassicExpiredDays() *string {
	return s.ClassicExpiredDays
}

func (s *SwitchNetworkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchNetworkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchNetworkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchNetworkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchNetworkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchNetworkRequest) GetRetainClassic() *string {
	return s.RetainClassic
}

func (s *SwitchNetworkRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SwitchNetworkRequest) GetTargetNetworkType() *string {
	return s.TargetNetworkType
}

func (s *SwitchNetworkRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *SwitchNetworkRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *SwitchNetworkRequest) SetClassicExpiredDays(v string) *SwitchNetworkRequest {
	s.ClassicExpiredDays = &v
	return s
}

func (s *SwitchNetworkRequest) SetInstanceId(v string) *SwitchNetworkRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchNetworkRequest) SetOwnerAccount(v string) *SwitchNetworkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchNetworkRequest) SetOwnerId(v int64) *SwitchNetworkRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchNetworkRequest) SetResourceOwnerAccount(v string) *SwitchNetworkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchNetworkRequest) SetResourceOwnerId(v int64) *SwitchNetworkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchNetworkRequest) SetRetainClassic(v string) *SwitchNetworkRequest {
	s.RetainClassic = &v
	return s
}

func (s *SwitchNetworkRequest) SetSecurityToken(v string) *SwitchNetworkRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchNetworkRequest) SetTargetNetworkType(v string) *SwitchNetworkRequest {
	s.TargetNetworkType = &v
	return s
}

func (s *SwitchNetworkRequest) SetVSwitchId(v string) *SwitchNetworkRequest {
	s.VSwitchId = &v
	return s
}

func (s *SwitchNetworkRequest) SetVpcId(v string) *SwitchNetworkRequest {
	s.VpcId = &v
	return s
}

func (s *SwitchNetworkRequest) Validate() error {
	return dara.Validate(s)
}
