// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUninstallAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppIdList(v []*string) *UninstallAppRequest
	GetAppIdList() []*string
	SetInstanceGroupIdList(v []*string) *UninstallAppRequest
	GetInstanceGroupIdList() []*string
	SetInstanceIdList(v []*string) *UninstallAppRequest
	GetInstanceIdList() []*string
}

type UninstallAppRequest struct {
	// The IDs of the apps.
	AppIdList []*string `json:"AppIdList,omitempty" xml:"AppIdList,omitempty" type:"Repeated"`
	// The ID of the instance groups. If you specify this parameter, you cannot specify InstanceIdList.
	InstanceGroupIdList []*string `json:"InstanceGroupIdList,omitempty" xml:"InstanceGroupIdList,omitempty" type:"Repeated"`
	// The IDs of the cloud phone instances. If you specify this parameter, you cannot specify InstanceGroupIdList.
	InstanceIdList []*string `json:"InstanceIdList,omitempty" xml:"InstanceIdList,omitempty" type:"Repeated"`
}

func (s UninstallAppRequest) String() string {
	return dara.Prettify(s)
}

func (s UninstallAppRequest) GoString() string {
	return s.String()
}

func (s *UninstallAppRequest) GetAppIdList() []*string {
	return s.AppIdList
}

func (s *UninstallAppRequest) GetInstanceGroupIdList() []*string {
	return s.InstanceGroupIdList
}

func (s *UninstallAppRequest) GetInstanceIdList() []*string {
	return s.InstanceIdList
}

func (s *UninstallAppRequest) SetAppIdList(v []*string) *UninstallAppRequest {
	s.AppIdList = v
	return s
}

func (s *UninstallAppRequest) SetInstanceGroupIdList(v []*string) *UninstallAppRequest {
	s.InstanceGroupIdList = v
	return s
}

func (s *UninstallAppRequest) SetInstanceIdList(v []*string) *UninstallAppRequest {
	s.InstanceIdList = v
	return s
}

func (s *UninstallAppRequest) Validate() error {
	return dara.Validate(s)
}
