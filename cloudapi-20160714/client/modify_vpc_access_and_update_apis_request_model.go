// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVpcAccessAndUpdateApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyVpcAccessAndUpdateApisRequest
	GetInstanceId() *string
	SetName(v string) *ModifyVpcAccessAndUpdateApisRequest
	GetName() *string
	SetNeedBatchWork(v bool) *ModifyVpcAccessAndUpdateApisRequest
	GetNeedBatchWork() *bool
	SetPort(v int32) *ModifyVpcAccessAndUpdateApisRequest
	GetPort() *int32
	SetRefresh(v bool) *ModifyVpcAccessAndUpdateApisRequest
	GetRefresh() *bool
	SetSecurityToken(v string) *ModifyVpcAccessAndUpdateApisRequest
	GetSecurityToken() *string
	SetToken(v string) *ModifyVpcAccessAndUpdateApisRequest
	GetToken() *string
	SetVpcId(v string) *ModifyVpcAccessAndUpdateApisRequest
	GetVpcId() *string
	SetVpcTargetHostName(v string) *ModifyVpcAccessAndUpdateApisRequest
	GetVpcTargetHostName() *string
}

type ModifyVpcAccessAndUpdateApisRequest struct {
	// The ID of the new instance.
	//
	// example:
	//
	// i-uf6bzcg1pr4oh5jjmxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the VPC authorization.
	//
	// >
	//
	// 	- The name of a VPC authorization cannot be changed. You cannot use this parameter to change the name of a VPC authorization.
	//
	// 	- You must set this parameter to the name of the current VPC authorization.
	//
	// This parameter is required.
	//
	// example:
	//
	// VpcName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies whether to update the associated API.
	//
	// **
	//
	// **Warning:*	- If you want to update the VPC authorization of a published API, you must set this parameter to true. Otherwise, the update will not be synchronized to the backend service of the API.
	//
	// example:
	//
	// true
	NeedBatchWork *bool `json:"NeedBatchWork,omitempty" xml:"NeedBatchWork,omitempty"`
	// The new port number.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// Specifies whether to update the VPC authorization.
	//
	// >
	//
	// 	- If the ID of the instance in your VPC is changed but the IP address of the instance remains unchanged, you can set this parameter to true to update the VPC authorization.
	//
	// example:
	//
	// false
	Refresh       *bool   `json:"Refresh,omitempty" xml:"Refresh,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The token of the request.
	//
	// example:
	//
	// c20d86c4-1eb3-4d0b-afe9-c586df1e2136
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
	// The ID of the new VPC.
	//
	// example:
	//
	// vpc-m5e7jqfppv5wbvmdw5pg2
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The hostname of the backend service.
	//
	// example:
	//
	// iot.hu***ng.com
	VpcTargetHostName *string `json:"VpcTargetHostName,omitempty" xml:"VpcTargetHostName,omitempty"`
}

func (s ModifyVpcAccessAndUpdateApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVpcAccessAndUpdateApisRequest) GoString() string {
	return s.String()
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetName() *string {
	return s.Name
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetNeedBatchWork() *bool {
	return s.NeedBatchWork
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetPort() *int32 {
	return s.Port
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetRefresh() *bool {
	return s.Refresh
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetToken() *string {
	return s.Token
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *ModifyVpcAccessAndUpdateApisRequest) GetVpcTargetHostName() *string {
	return s.VpcTargetHostName
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetInstanceId(v string) *ModifyVpcAccessAndUpdateApisRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetName(v string) *ModifyVpcAccessAndUpdateApisRequest {
	s.Name = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetNeedBatchWork(v bool) *ModifyVpcAccessAndUpdateApisRequest {
	s.NeedBatchWork = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetPort(v int32) *ModifyVpcAccessAndUpdateApisRequest {
	s.Port = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetRefresh(v bool) *ModifyVpcAccessAndUpdateApisRequest {
	s.Refresh = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetSecurityToken(v string) *ModifyVpcAccessAndUpdateApisRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetToken(v string) *ModifyVpcAccessAndUpdateApisRequest {
	s.Token = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetVpcId(v string) *ModifyVpcAccessAndUpdateApisRequest {
	s.VpcId = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) SetVpcTargetHostName(v string) *ModifyVpcAccessAndUpdateApisRequest {
	s.VpcTargetHostName = &v
	return s
}

func (s *ModifyVpcAccessAndUpdateApisRequest) Validate() error {
	return dara.Validate(s)
}
