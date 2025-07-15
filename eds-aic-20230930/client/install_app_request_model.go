// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdList(v []*string) *InstallAppRequest
	GetAppIdList() []*string
	SetInstanceGroupIdList(v []*string) *InstallAppRequest
	GetInstanceGroupIdList() []*string
	SetInstanceIdList(v []*string) *InstallAppRequest
	GetInstanceIdList() []*string
}

type InstallAppRequest struct {
	// The IDs of the apps that you want to install.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	// The IDs of the instance groups.
	InstanceGroupIdList []*string `json:"InstanceGroupIdList,omitempty" xml:"InstanceGroupIdList,omitempty" type:"Repeated"`
	// The IDs of the cloud phone instances.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s InstallAppRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallAppRequest) GoString() string {
	return s.String()
}

func (s *InstallAppRequest) GetAppIdList() []*string {
	return s.AppIdList
}

func (s *InstallAppRequest) GetInstanceGroupIdList() []*string {
	return s.InstanceGroupIdList
}

func (s *InstallAppRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *InstallAppRequest) SetAppIdList(v []*string) *InstallAppRequest {
	s.AppIdList = v
	return s
}

func (s *InstallAppRequest) SetInstanceGroupIdList(v []*string) *InstallAppRequest {
	s.InstanceGroupIdList = v
	return s
}

func (s *InstallAppRequest) SetInstanceIdList(v []*string) *InstallAppRequest {
	s.InstanceIdList = v
	return s
}

func (s *InstallAppRequest) Validate() error {
	return dara.Validate(s)
}
