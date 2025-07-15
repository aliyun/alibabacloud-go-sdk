// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyInstanceNameRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyInstanceNameRequest
	GetInstanceName() *string
	SetRegionId(v string) *ModifyInstanceNameRequest
	GetRegionId() *string
}

type ModifyInstanceNameRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_post-cn-v0h1fgs2****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The instance name. Valid values:
	//
	// 	- The name can contain only letters, digits, hyphens (-), and underscores (_).
	//
	// 	- The name must be 3 to 64 characters in length. A name that contains more than 64 characters is automatically truncated.
	//
	// This parameter is required.
	//
	// example:
	//
	// dev-test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceNameRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceNameRequest) SetInstanceId(v string) *ModifyInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetInstanceName(v string) *ModifyInstanceNameRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceNameRequest) SetRegionId(v string) *ModifyInstanceNameRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
