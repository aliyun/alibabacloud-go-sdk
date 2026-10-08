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
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// This parameter is required.
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
