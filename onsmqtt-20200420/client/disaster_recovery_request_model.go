// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisasterRecoveryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisasterRecoveryRequest
	GetInstanceId() *string
	SetRecoveryInstanceId(v string) *DisasterRecoveryRequest
	GetRecoveryInstanceId() *string
}

type DisasterRecoveryRequest struct {
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the instance to recover.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-xxx
	RecoveryInstanceId *string `json:"RecoveryInstanceId,omitempty" xml:"RecoveryInstanceId,omitempty"`
}

func (s DisasterRecoveryRequest) String() string {
	return dara.Prettify(s)
}

func (s DisasterRecoveryRequest) GoString() string {
	return s.String()
}

func (s *DisasterRecoveryRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisasterRecoveryRequest) GetRecoveryInstanceId() *string {
	return s.RecoveryInstanceId
}

func (s *DisasterRecoveryRequest) SetInstanceId(v string) *DisasterRecoveryRequest {
	s.InstanceId = &v
	return s
}

func (s *DisasterRecoveryRequest) SetRecoveryInstanceId(v string) *DisasterRecoveryRequest {
	s.RecoveryInstanceId = &v
	return s
}

func (s *DisasterRecoveryRequest) Validate() error {
	return dara.Validate(s)
}
