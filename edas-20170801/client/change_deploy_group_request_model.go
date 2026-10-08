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
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3616cdca-4f92-**********
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The Elastic Compute Container (ECC) ID of the ECS instance whose group you want to change. Call the ListApplicationEcc operation to query the ECC ID of an application. For more information, see [ListApplicationEcc](https://help.aliyun.com/document_detail/199277.html).
	//
	// > You can change the group for only one ECS instance at a time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0cf49a6c-95a8-4aa8******
	EccInfo *string `json:"EccInfo,omitempty" xml:"EccInfo,omitempty"`
	// Specifies whether to force the change when the deployment package version of the ECC is different from the deployment package version of the application group.
	//
	// example:
	//
	// true
	ForceStatus *bool `json:"ForceStatus,omitempty" xml:"ForceStatus,omitempty"`
	// The name of the application group, such as \\`group_a\\` and \\`group_b\\`. The GroupName for the default group is `_DEFAULT_GROUP`. The name can be up to 64 characters long.
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
