// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFastMode(v bool) *RestartInstanceRequest
	GetFastMode() *bool
	SetInstanceId(v string) *RestartInstanceRequest
	GetInstanceId() *string
}

type RestartInstanceRequest struct {
	// Specifies whether to restart compute nodes in quick restart mode. Default value: false. Valid values:
	//
	// 	- true: Compute nodes are restarted in quick restart mode in multiple batches. The batches are executed in parallel, and the nodes in each batch are restarted at the same time.
	//
	// 	- false: Compute nodes are restarted in rolling restart mode.
	//
	// example:
	//
	// true
	FastMode *bool `json:"FastMode,omitempty" xml:"FastMode,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RestartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RestartInstanceRequest) GoString() string {
	return s.String()
}

func (s *RestartInstanceRequest) GetFastMode() *bool {
	return s.FastMode
}

func (s *RestartInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RestartInstanceRequest) SetFastMode(v bool) *RestartInstanceRequest {
	s.FastMode = &v
	return s
}

func (s *RestartInstanceRequest) SetInstanceId(v string) *RestartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RestartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
