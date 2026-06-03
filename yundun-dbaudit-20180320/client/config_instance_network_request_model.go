// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceNetworkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConfigInstanceNetworkRequest
	GetInstanceId() *string
	SetPrivateWhiteList(v []*string) *ConfigInstanceNetworkRequest
	GetPrivateWhiteList() []*string
	SetPublicAccessControl(v int32) *ConfigInstanceNetworkRequest
	GetPublicAccessControl() *int32
	SetPublicWhiteList(v []*string) *ConfigInstanceNetworkRequest
	GetPublicWhiteList() []*string
	SetRegionId(v string) *ConfigInstanceNetworkRequest
	GetRegionId() *string
	SetSecurityGroupIds(v []*string) *ConfigInstanceNetworkRequest
	GetSecurityGroupIds() []*string
}

type ConfigInstanceNetworkRequest struct {
	// This parameter is required.
	InstanceId       *string   `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	PrivateWhiteList []*string `json:"PrivateWhiteList,omitempty" xml:"PrivateWhiteList,omitempty" type:"Repeated"`
	// This parameter is required.
	PublicAccessControl *int32    `json:"PublicAccessControl,omitempty" xml:"PublicAccessControl,omitempty"`
	PublicWhiteList     []*string `json:"PublicWhiteList,omitempty" xml:"PublicWhiteList,omitempty" type:"Repeated"`
	RegionId            *string   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// This parameter is required.
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
}

func (s ConfigInstanceNetworkRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceNetworkRequest) GoString() string {
	return s.String()
}

func (s *ConfigInstanceNetworkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConfigInstanceNetworkRequest) GetPrivateWhiteList() []*string {
	return s.PrivateWhiteList
}

func (s *ConfigInstanceNetworkRequest) GetPublicAccessControl() *int32 {
	return s.PublicAccessControl
}

func (s *ConfigInstanceNetworkRequest) GetPublicWhiteList() []*string {
	return s.PublicWhiteList
}

func (s *ConfigInstanceNetworkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConfigInstanceNetworkRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ConfigInstanceNetworkRequest) SetInstanceId(v string) *ConfigInstanceNetworkRequest {
	s.InstanceId = &v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetPrivateWhiteList(v []*string) *ConfigInstanceNetworkRequest {
	s.PrivateWhiteList = v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetPublicAccessControl(v int32) *ConfigInstanceNetworkRequest {
	s.PublicAccessControl = &v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetPublicWhiteList(v []*string) *ConfigInstanceNetworkRequest {
	s.PublicWhiteList = v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetRegionId(v string) *ConfigInstanceNetworkRequest {
	s.RegionId = &v
	return s
}

func (s *ConfigInstanceNetworkRequest) SetSecurityGroupIds(v []*string) *ConfigInstanceNetworkRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ConfigInstanceNetworkRequest) Validate() error {
	return dara.Validate(s)
}
