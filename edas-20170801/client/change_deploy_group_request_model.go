// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeDeployGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ChangeDeployGroupRequest
	GetAppId() *string
	SetEccInfo(v string) *ChangeDeployGroupRequest
	GetEccInfo() *string
	SetForceStatus(v bool) *ChangeDeployGroupRequest
	GetForceStatus() *bool
	SetGroupName(v string) *ChangeDeployGroupRequest
	GetGroupName() *string
}

type ChangeDeployGroupRequest struct {
	// The ID of the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-**********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the elastic compute component (ECC) that corresponds to the ECS instance for which you want to change the application instance group. You can call the ListApplicationEcc operation to query the ECC ID. For more information, see [ListApplicationEcc](https://help.aliyun.com/document_detail/199277.html).
	//
	// > You can change the application instance group for only one ECS instance at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0cf49a6c-95a8-4aa8******
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
	// Specifies whether to forcibly change the application instance group if the deployment package version of the ECC is different from that of the application instance group.
	//
	// example:
	//
	// true
	ForceStatus *bool `json:"ForceStatus,omitempty" xml:"ForceStatus,omitempty"`
	// The name of the application instance group. Examples: group_a and group_b. The parameter value for the default application instance group is `_DEFAULT_GROUP`. The name can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s ChangeDeployGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeDeployGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeDeployGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *ChangeDeployGroupRequest) GetEccInfo() *string {
	return s.EccInfo
}

func (s *ChangeDeployGroupRequest) GetForceStatus() *bool {
	return s.ForceStatus
}

func (s *ChangeDeployGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ChangeDeployGroupRequest) SetAppId(v string) *ChangeDeployGroupRequest {
	s.AppId = &v
	return s
}

func (s *ChangeDeployGroupRequest) SetEccInfo(v string) *ChangeDeployGroupRequest {
	s.EccInfo = &v
	return s
}

func (s *ChangeDeployGroupRequest) SetForceStatus(v bool) *ChangeDeployGroupRequest {
	s.ForceStatus = &v
	return s
}

func (s *ChangeDeployGroupRequest) SetGroupName(v string) *ChangeDeployGroupRequest {
	s.GroupName = &v
	return s
}

func (s *ChangeDeployGroupRequest) Validate() error {
	return dara.Validate(s)
}
