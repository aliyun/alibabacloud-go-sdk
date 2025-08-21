// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBlackholeStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBlackholeStatus(v string) *ModifyBlackholeStatusRequest
	GetBlackholeStatus() *string
	SetInstanceId(v string) *ModifyBlackholeStatusRequest
	GetInstanceId() *string
}

type ModifyBlackholeStatusRequest struct {
	// The action that you want to perform on the instance. Set the value to **undo**, which indicates that you want to deactivate blackhole filtering.
	//
	// This parameter is required.
	//
	// example:
	//
	// undo
	BlackholeStatus *string `json:"BlackholeStatus,omitempty" xml:"BlackholeStatus,omitempty"`
	// The ID of the instance.
	//
	// > You can call the [DescribeInstanceIds](https://help.aliyun.com/document_detail/157459.html) operation to query the IDs of all instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyBlackholeStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyBlackholeStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyBlackholeStatusRequest) GetBlackholeStatus() *string {
	return s.BlackholeStatus
}

func (s *ModifyBlackholeStatusRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyBlackholeStatusRequest) SetBlackholeStatus(v string) *ModifyBlackholeStatusRequest {
	s.BlackholeStatus = &v
	return s
}

func (s *ModifyBlackholeStatusRequest) SetInstanceId(v string) *ModifyBlackholeStatusRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyBlackholeStatusRequest) Validate() error {
	return dara.Validate(s)
}
