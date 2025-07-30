// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCuPreCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyCuPreCheckRequest
	GetInstanceId() *string
	SetNodeGroupId(v string) *ModifyCuPreCheckRequest
	GetNodeGroupId() *string
	SetTarget(v int32) *ModifyCuPreCheckRequest
	GetTarget() *int32
}

type ModifyCuPreCheckRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The warehouse ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ng-d332aa8bca48****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The number of CUs to which you want to change.
	//
	// Valid values:
	//
	// 	- 2
	//
	// 	- 4
	//
	// 	- 8
	//
	// 	- 16
	//
	// 	- 32
	//
	// 	- 64
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Target *int32 `json:"Target,omitempty" xml:"Target,omitempty"`
}

func (s ModifyCuPreCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCuPreCheckRequest) GoString() string {
	return s.String()
}

func (s *ModifyCuPreCheckRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyCuPreCheckRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ModifyCuPreCheckRequest) GetTarget() *int32 {
	return s.Target
}

func (s *ModifyCuPreCheckRequest) SetInstanceId(v string) *ModifyCuPreCheckRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyCuPreCheckRequest) SetNodeGroupId(v string) *ModifyCuPreCheckRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ModifyCuPreCheckRequest) SetTarget(v int32) *ModifyCuPreCheckRequest {
	s.Target = &v
	return s
}

func (s *ModifyCuPreCheckRequest) Validate() error {
	return dara.Validate(s)
}
