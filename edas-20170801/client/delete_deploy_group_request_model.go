// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeployGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *DeleteDeployGroupRequest
	GetAppId() *string
	SetGroupName(v string) *DeleteDeployGroupRequest
	GetGroupName() *string
}

type DeleteDeployGroupRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-4413-b31*************
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the instance group.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteDeployGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeployGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeployGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *DeleteDeployGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteDeployGroupRequest) SetAppId(v string) *DeleteDeployGroupRequest {
	s.AppId = &v
	return s
}

func (s *DeleteDeployGroupRequest) SetGroupName(v string) *DeleteDeployGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteDeployGroupRequest) Validate() error {
	return dara.Validate(s)
}
