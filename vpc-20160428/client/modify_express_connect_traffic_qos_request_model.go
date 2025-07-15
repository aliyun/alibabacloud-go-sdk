// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExpressConnectTrafficQosRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddInstanceList(v []*ModifyExpressConnectTrafficQosRequestAddInstanceList) *ModifyExpressConnectTrafficQosRequest
	GetAddInstanceList() []*ModifyExpressConnectTrafficQosRequestAddInstanceList
	SetClientToken(v string) *ModifyExpressConnectTrafficQosRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *ModifyExpressConnectTrafficQosRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyExpressConnectTrafficQosRequest
	GetOwnerId() *int64
	SetQosDescription(v string) *ModifyExpressConnectTrafficQosRequest
	GetQosDescription() *string
	SetQosId(v string) *ModifyExpressConnectTrafficQosRequest
	GetQosId() *string
	SetQosName(v string) *ModifyExpressConnectTrafficQosRequest
	GetQosName() *string
	SetRegionId(v string) *ModifyExpressConnectTrafficQosRequest
	GetRegionId() *string
	SetRemoveInstanceList(v []*ModifyExpressConnectTrafficQosRequestRemoveInstanceList) *ModifyExpressConnectTrafficQosRequest
	GetRemoveInstanceList() []*ModifyExpressConnectTrafficQosRequestRemoveInstanceList
	SetResourceOwnerAccount(v string) *ModifyExpressConnectTrafficQosRequest
	GetResourceOwnerAccount() *string
}

type ModifyExpressConnectTrafficQosRequest struct {
	// The instances to be added. Ignore this parameter if no instances are to be added.
	//
	// if can be null:
	// false
	AddInstanceList []*ModifyExpressConnectTrafficQosRequestAddInstanceList `json:"AddInstanceList,omitempty" xml:"AddInstanceList,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- may be different for each request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken  *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the QoS policy.
	//
	// example:
	//
	// qos-test
	QosDescription *string `json:"QosDescription,omitempty" xml:"QosDescription,omitempty"`
	// The ID of the QoS policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// qos-2giu0a6vd5x0mv4700
	QosId *string `json:"QosId,omitempty" xml:"QosId,omitempty"`
	// The name of the QoS policy.
	//
	// example:
	//
	// qos-test
	QosName *string `json:"QosName,omitempty" xml:"QosName,omitempty"`
	// The region ID of the resource.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instances to be removed. Ignore this parameter if no instances are to be removed.
	RemoveInstanceList   []*ModifyExpressConnectTrafficQosRequestRemoveInstanceList `json:"RemoveInstanceList,omitempty" xml:"RemoveInstanceList,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                                                    `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
}

func (s ModifyExpressConnectTrafficQosRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosRequest) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosRequest) GetAddInstanceList() []*ModifyExpressConnectTrafficQosRequestAddInstanceList {
	return s.AddInstanceList
}

func (s *ModifyExpressConnectTrafficQosRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyExpressConnectTrafficQosRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyExpressConnectTrafficQosRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyExpressConnectTrafficQosRequest) GetQosDescription() *string {
	return s.QosDescription
}

func (s *ModifyExpressConnectTrafficQosRequest) GetQosId() *string {
	return s.QosId
}

func (s *ModifyExpressConnectTrafficQosRequest) GetQosName() *string {
	return s.QosName
}

func (s *ModifyExpressConnectTrafficQosRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyExpressConnectTrafficQosRequest) GetRemoveInstanceList() []*ModifyExpressConnectTrafficQosRequestRemoveInstanceList {
	return s.RemoveInstanceList
}

func (s *ModifyExpressConnectTrafficQosRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyExpressConnectTrafficQosRequest) SetAddInstanceList(v []*ModifyExpressConnectTrafficQosRequestAddInstanceList) *ModifyExpressConnectTrafficQosRequest {
	s.AddInstanceList = v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetClientToken(v string) *ModifyExpressConnectTrafficQosRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetOwnerAccount(v string) *ModifyExpressConnectTrafficQosRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetOwnerId(v int64) *ModifyExpressConnectTrafficQosRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetQosDescription(v string) *ModifyExpressConnectTrafficQosRequest {
	s.QosDescription = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetQosId(v string) *ModifyExpressConnectTrafficQosRequest {
	s.QosId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetQosName(v string) *ModifyExpressConnectTrafficQosRequest {
	s.QosName = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetRegionId(v string) *ModifyExpressConnectTrafficQosRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetRemoveInstanceList(v []*ModifyExpressConnectTrafficQosRequestRemoveInstanceList) *ModifyExpressConnectTrafficQosRequest {
	s.RemoveInstanceList = v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) SetResourceOwnerAccount(v string) *ModifyExpressConnectTrafficQosRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequest) Validate() error {
	return dara.Validate(s)
}

type ModifyExpressConnectTrafficQosRequestAddInstanceList struct {
	// The ID of the instance to be associated.
	//
	// example:
	//
	// pc-bp159zj8zujwy3p07****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of instance to be associated. Set the value to **PHYSICALCONNECTION**.
	//
	// example:
	//
	// PHYSICALCONNECTION
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ModifyExpressConnectTrafficQosRequestAddInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosRequestAddInstanceList) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosRequestAddInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyExpressConnectTrafficQosRequestAddInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyExpressConnectTrafficQosRequestAddInstanceList) SetInstanceId(v string) *ModifyExpressConnectTrafficQosRequestAddInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequestAddInstanceList) SetInstanceType(v string) *ModifyExpressConnectTrafficQosRequestAddInstanceList {
	s.InstanceType = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequestAddInstanceList) Validate() error {
	return dara.Validate(s)
}

type ModifyExpressConnectTrafficQosRequestRemoveInstanceList struct {
	// The ID of the associated instance.
	//
	// example:
	//
	// pc-bp1j37am632492qzw****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The type of the associated instance. Set the value to **PHYSICALCONNECTION**.
	//
	// example:
	//
	// PHYSICALCONNECTION
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
}

func (s ModifyExpressConnectTrafficQosRequestRemoveInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ModifyExpressConnectTrafficQosRequestRemoveInstanceList) GoString() string {
	return s.String()
}

func (s *ModifyExpressConnectTrafficQosRequestRemoveInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyExpressConnectTrafficQosRequestRemoveInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyExpressConnectTrafficQosRequestRemoveInstanceList) SetInstanceId(v string) *ModifyExpressConnectTrafficQosRequestRemoveInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequestRemoveInstanceList) SetInstanceType(v string) *ModifyExpressConnectTrafficQosRequestRemoveInstanceList {
	s.InstanceType = &v
	return s
}

func (s *ModifyExpressConnectTrafficQosRequestRemoveInstanceList) Validate() error {
	return dara.Validate(s)
}
