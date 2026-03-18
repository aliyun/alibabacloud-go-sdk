// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInnerIpWhitelistGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCidrIpList(v []*string) *UpdateInnerIpWhitelistGroupRequest
	GetCidrIpList() []*string
	SetInnerIpWhitelistGroupId(v string) *UpdateInnerIpWhitelistGroupRequest
	GetInnerIpWhitelistGroupId() *string
	SetInstanceId(v string) *UpdateInnerIpWhitelistGroupRequest
	GetInstanceId() *string
}

type UpdateInnerIpWhitelistGroupRequest struct {
	// This parameter is required.
	CidrIpList []*string `json:"CidrIpList,omitempty" xml:"CidrIpList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// sg-28sjsi12bs1
	InnerIpWhitelistGroupId *string `json:"InnerIpWhitelistGroupId,omitempty" xml:"InnerIpWhitelistGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c-0104730e9de40215
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateInnerIpWhitelistGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInnerIpWhitelistGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateInnerIpWhitelistGroupRequest) GetCidrIpList() []*string {
	return s.CidrIpList
}

func (s *UpdateInnerIpWhitelistGroupRequest) GetInnerIpWhitelistGroupId() *string {
	return s.InnerIpWhitelistGroupId
}

func (s *UpdateInnerIpWhitelistGroupRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInnerIpWhitelistGroupRequest) SetCidrIpList(v []*string) *UpdateInnerIpWhitelistGroupRequest {
	s.CidrIpList = v
	return s
}

func (s *UpdateInnerIpWhitelistGroupRequest) SetInnerIpWhitelistGroupId(v string) *UpdateInnerIpWhitelistGroupRequest {
	s.InnerIpWhitelistGroupId = &v
	return s
}

func (s *UpdateInnerIpWhitelistGroupRequest) SetInstanceId(v string) *UpdateInnerIpWhitelistGroupRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInnerIpWhitelistGroupRequest) Validate() error {
	return dara.Validate(s)
}
