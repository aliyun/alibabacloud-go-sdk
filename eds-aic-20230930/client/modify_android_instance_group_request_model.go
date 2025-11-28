// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAndroidInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceGroupId(v string) *ModifyAndroidInstanceGroupRequest
	GetInstanceGroupId() *string
	SetNewInstanceGroupName(v string) *ModifyAndroidInstanceGroupRequest
	GetNewInstanceGroupName() *string
	SetPolicyGroupId(v string) *ModifyAndroidInstanceGroupRequest
	GetPolicyGroupId() *string
	SetStreamMode(v int32) *ModifyAndroidInstanceGroupRequest
	GetStreamMode() *int32
}

type ModifyAndroidInstanceGroupRequest struct {
	// The ID of the instance group.
	//
	// example:
	//
	// ag-cuv4scs4obxhs****
	InstanceGroupId *string `json:"InstanceGroupId,omitempty" xml:"InstanceGroupId,omitempty"`
	// The new name of the instance group.
	//
	// >
	//
	// 	- The name can be up to 30 characters in length. It can contain letters, digits, colons (:), underscores (_), periods (.), or hyphens (-). It must start with letters but cannot start with http:// or https://.
	//
	// example:
	//
	// newName
	NewInstanceGroupName *string `json:"NewInstanceGroupName,omitempty" xml:"NewInstanceGroupName,omitempty"`
	// The ID of the policy.
	//
	// example:
	//
	// pg-2w97kp89gnsif****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	StreamMode    *int32  `json:"StreamMode,omitempty" xml:"StreamMode,omitempty"`
}

func (s ModifyAndroidInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAndroidInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyAndroidInstanceGroupRequest) GetInstanceGroupId() *string {
	return s.InstanceGroupId
}

func (s *ModifyAndroidInstanceGroupRequest) GetNewInstanceGroupName() *string {
	return s.NewInstanceGroupName
}

func (s *ModifyAndroidInstanceGroupRequest) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *ModifyAndroidInstanceGroupRequest) GetStreamMode() *int32 {
	return s.StreamMode
}

func (s *ModifyAndroidInstanceGroupRequest) SetInstanceGroupId(v string) *ModifyAndroidInstanceGroupRequest {
	s.InstanceGroupId = &v
	return s
}

func (s *ModifyAndroidInstanceGroupRequest) SetNewInstanceGroupName(v string) *ModifyAndroidInstanceGroupRequest {
	s.NewInstanceGroupName = &v
	return s
}

func (s *ModifyAndroidInstanceGroupRequest) SetPolicyGroupId(v string) *ModifyAndroidInstanceGroupRequest {
	s.PolicyGroupId = &v
	return s
}

func (s *ModifyAndroidInstanceGroupRequest) SetStreamMode(v int32) *ModifyAndroidInstanceGroupRequest {
	s.StreamMode = &v
	return s
}

func (s *ModifyAndroidInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}
