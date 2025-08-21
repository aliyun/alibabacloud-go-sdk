// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceStop(v string) *StopInstanceRequest
	GetForceStop() *string
	SetInstanceId(v string) *StopInstanceRequest
	GetInstanceId() *string
}

type StopInstanceRequest struct {
	// Specifies whether to forcibly stop the instance.
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	ForceStop *string `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The ID of the instance that you want to stop. You can specify only one instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-instanceid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) GetForceStop() *string {
	return s.ForceStop
}

func (s *StopInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopInstanceRequest) SetForceStop(v string) *StopInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) Validate() error {
	return dara.Validate(s)
}
