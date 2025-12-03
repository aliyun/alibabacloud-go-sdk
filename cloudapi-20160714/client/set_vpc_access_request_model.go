// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetVpcAccessRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *SetVpcAccessRequest
	GetDescription() *string
	SetInstanceId(v string) *SetVpcAccessRequest
	GetInstanceId() *string
	SetName(v string) *SetVpcAccessRequest
	GetName() *string
	SetPort(v int32) *SetVpcAccessRequest
	GetPort() *int32
	SetSecurityToken(v string) *SetVpcAccessRequest
	GetSecurityToken() *string
	SetTag(v []*SetVpcAccessRequestTag) *SetVpcAccessRequest
	GetTag() []*SetVpcAccessRequestTag
	SetVpcId(v string) *SetVpcAccessRequest
	GetVpcId() *string
	SetVpcTargetHostName(v string) *SetVpcAccessRequest
	GetVpcTargetHostName() *string
}

type SetVpcAccessRequest struct {
	// The description of the VPC.
	//
	// example:
	//
	// description of the VPC
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of an ECS or SLB instance in the VPC.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf6bzcg1pr4oh5jjmxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the authorization. The name must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The port number that corresponds to the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	Port          *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The tag of objects that match the rule. You can specify multiple tags.
	Tag []*SetVpcAccessRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the VPC. The VPC must be an available one that belongs to the same account as the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-uf657qec7lx42paw3qxxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The host of the backend service.
	//
	// example:
	//
	// iot.hu***ng.com
	VpcTargetHostName *string `json:"VpcTargetHostName,omitempty" xml:"VpcTargetHostName,omitempty"`
}

func (s SetVpcAccessRequest) String() string {
	return dara.Prettify(s)
}

func (s SetVpcAccessRequest) GoString() string {
	return s.String()
}

func (s *SetVpcAccessRequest) GetDescription() *string {
	return s.Description
}

func (s *SetVpcAccessRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetVpcAccessRequest) GetName() *string {
	return s.Name
}

func (s *SetVpcAccessRequest) GetPort() *int32 {
	return s.Port
}

func (s *SetVpcAccessRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetVpcAccessRequest) GetTag() []*SetVpcAccessRequestTag {
	return s.Tag
}

func (s *SetVpcAccessRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *SetVpcAccessRequest) GetVpcTargetHostName() *string {
	return s.VpcTargetHostName
}

func (s *SetVpcAccessRequest) SetDescription(v string) *SetVpcAccessRequest {
	s.Description = &v
	return s
}

func (s *SetVpcAccessRequest) SetInstanceId(v string) *SetVpcAccessRequest {
	s.InstanceId = &v
	return s
}

func (s *SetVpcAccessRequest) SetName(v string) *SetVpcAccessRequest {
	s.Name = &v
	return s
}

func (s *SetVpcAccessRequest) SetPort(v int32) *SetVpcAccessRequest {
	s.Port = &v
	return s
}

func (s *SetVpcAccessRequest) SetSecurityToken(v string) *SetVpcAccessRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetVpcAccessRequest) SetTag(v []*SetVpcAccessRequestTag) *SetVpcAccessRequest {
	s.Tag = v
	return s
}

func (s *SetVpcAccessRequest) SetVpcId(v string) *SetVpcAccessRequest {
	s.VpcId = &v
	return s
}

func (s *SetVpcAccessRequest) SetVpcTargetHostName(v string) *SetVpcAccessRequest {
	s.VpcTargetHostName = &v
	return s
}

func (s *SetVpcAccessRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SetVpcAccessRequestTag struct {
	// The key of the tag.
	//
	// example:
	//
	// key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// 123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SetVpcAccessRequestTag) String() string {
	return dara.Prettify(s)
}

func (s SetVpcAccessRequestTag) GoString() string {
	return s.String()
}

func (s *SetVpcAccessRequestTag) GetKey() *string {
	return s.Key
}

func (s *SetVpcAccessRequestTag) GetValue() *string {
	return s.Value
}

func (s *SetVpcAccessRequestTag) SetKey(v string) *SetVpcAccessRequestTag {
	s.Key = &v
	return s
}

func (s *SetVpcAccessRequestTag) SetValue(v string) *SetVpcAccessRequestTag {
	s.Value = &v
	return s
}

func (s *SetVpcAccessRequestTag) Validate() error {
	return dara.Validate(s)
}
