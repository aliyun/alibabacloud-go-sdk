// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceUpgradeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppVersion(v string) *GetDeviceUpgradeStatusRequest
	GetAppVersion() *string
	SetClientUid(v string) *GetDeviceUpgradeStatusRequest
	GetClientUid() *string
	SetProject(v string) *GetDeviceUpgradeStatusRequest
	GetProject() *string
	SetTaskUid(v string) *GetDeviceUpgradeStatusRequest
	GetTaskUid() *string
}

type GetDeviceUpgradeStatusRequest struct {
	AppVersion *string `json:"AppVersion,omitempty" xml:"AppVersion,omitempty"`
	ClientUid  *string `json:"ClientUid,omitempty" xml:"ClientUid,omitempty"`
	Project    *string `json:"Project,omitempty" xml:"Project,omitempty"`
	TaskUid    *string `json:"TaskUid,omitempty" xml:"TaskUid,omitempty"`
}

func (s GetDeviceUpgradeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceUpgradeStatusRequest) GoString() string {
	return s.String()
}

func (s *GetDeviceUpgradeStatusRequest) GetAppVersion() *string {
	return s.AppVersion
}

func (s *GetDeviceUpgradeStatusRequest) GetClientUid() *string {
	return s.ClientUid
}

func (s *GetDeviceUpgradeStatusRequest) GetProject() *string {
	return s.Project
}

func (s *GetDeviceUpgradeStatusRequest) GetTaskUid() *string {
	return s.TaskUid
}

func (s *GetDeviceUpgradeStatusRequest) SetAppVersion(v string) *GetDeviceUpgradeStatusRequest {
	s.AppVersion = &v
	return s
}

func (s *GetDeviceUpgradeStatusRequest) SetClientUid(v string) *GetDeviceUpgradeStatusRequest {
	s.ClientUid = &v
	return s
}

func (s *GetDeviceUpgradeStatusRequest) SetProject(v string) *GetDeviceUpgradeStatusRequest {
	s.Project = &v
	return s
}

func (s *GetDeviceUpgradeStatusRequest) SetTaskUid(v string) *GetDeviceUpgradeStatusRequest {
	s.TaskUid = &v
	return s
}

func (s *GetDeviceUpgradeStatusRequest) Validate() error {
	return dara.Validate(s)
}
