// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisasterDowngradeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDowngradeInstanceId(v string) *DisasterDowngradeRequest
	GetDowngradeInstanceId() *string
	SetInstanceId(v string) *DisasterDowngradeRequest
	GetInstanceId() *string
}

type DisasterDowngradeRequest struct {
	// The instance ID to be downgraded.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-xxxx
	DowngradeInstanceId *string `json:"DowngradeInstanceId,omitempty" xml:"DowngradeInstanceId,omitempty"`
	// The ID of the ApsaraMQ for MQTT (MSMQ) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisasterDowngradeRequest) String() string {
	return dara.Prettify(s)
}

func (s DisasterDowngradeRequest) GoString() string {
	return s.String()
}

func (s *DisasterDowngradeRequest) GetDowngradeInstanceId() *string {
	return s.DowngradeInstanceId
}

func (s *DisasterDowngradeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisasterDowngradeRequest) SetDowngradeInstanceId(v string) *DisasterDowngradeRequest {
	s.DowngradeInstanceId = &v
	return s
}

func (s *DisasterDowngradeRequest) SetInstanceId(v string) *DisasterDowngradeRequest {
	s.InstanceId = &v
	return s
}

func (s *DisasterDowngradeRequest) Validate() error {
	return dara.Validate(s)
}
