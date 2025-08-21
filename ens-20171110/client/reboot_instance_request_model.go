// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceStop(v string) *RebootInstanceRequest
	GetForceStop() *string
	SetInstanceId(v string) *RebootInstanceRequest
	GetInstanceId() *string
}

type RebootInstanceRequest struct {
	// Specifies whether to forcefully stop the instance before you restart it.
	//
	// 	- **true**
	//
	// 	- **false*	- (default)
	//
	// example:
	//
	// true
	ForceStop *string `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The ID of the instance that you want to reboot. You can specify only one instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-instanceid****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RebootInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootInstanceRequest) GetForceStop() *string {
	return s.ForceStop
}

func (s *RebootInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootInstanceRequest) SetForceStop(v string) *RebootInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *RebootInstanceRequest) SetInstanceId(v string) *RebootInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootInstanceRequest) Validate() error {
	return dara.Validate(s)
}
