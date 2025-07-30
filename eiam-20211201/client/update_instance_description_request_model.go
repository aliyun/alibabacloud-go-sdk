// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateInstanceDescriptionRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateInstanceDescriptionRequest
	GetInstanceId() *string
}

type UpdateInstanceDescriptionRequest struct {
	// The new description of the instance.
	//
	// example:
	//
	// 测试实例
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the instance whose description you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceDescriptionRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceDescriptionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceDescriptionRequest) SetDescription(v string) *UpdateInstanceDescriptionRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceDescriptionRequest) SetInstanceId(v string) *UpdateInstanceDescriptionRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}
