// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateInstanceNameRequest
	GetInstanceId() *string
	SetInstanceName(v string) *UpdateInstanceNameRequest
	GetInstanceName() *string
}

type UpdateInstanceNameRequest struct {
	// The ID of the ApsaraMQ for RabbitMQ instance for which you want to update the name.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-zvp2ajsj****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the instance. No limits are imposed on the value. We recommend that you set this parameter to a maximum of 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// amqp-cn-ZVp2ajsj****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s UpdateInstanceNameRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNameRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNameRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceNameRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *UpdateInstanceNameRequest) SetInstanceId(v string) *UpdateInstanceNameRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceNameRequest) SetInstanceName(v string) *UpdateInstanceNameRequest {
	s.InstanceName = &v
	return s
}

func (s *UpdateInstanceNameRequest) Validate() error {
	return dara.Validate(s)
}
