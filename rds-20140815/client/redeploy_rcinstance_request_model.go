// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForceStop(v bool) *RedeployRCInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *RedeployRCInstanceRequest
	GetInstanceId() *string
}

type RedeployRCInstanceRequest struct {
	// Specifies whether to forcefully stop the instance that is in the Running state. Default value: false.
	//
	// >  A forced stop is equivalent to the shutdown operation for a physical database server and can result in loss of data that is not written to storage devices. We recommend that you redeploy instances when they are in the Stopped state.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-s8t4zcwr5fnmfycn****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s RedeployRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RedeployRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *RedeployRCInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *RedeployRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RedeployRCInstanceRequest) SetForceStop(v bool) *RedeployRCInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *RedeployRCInstanceRequest) SetInstanceId(v string) *RedeployRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RedeployRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
