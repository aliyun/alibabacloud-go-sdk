// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateInstanceRequest
	GetDescription() *string
	SetMaxSlot(v string) *UpdateInstanceRequest
	GetMaxSlot() *string
	SetName(v string) *UpdateInstanceRequest
	GetName() *string
}

type UpdateInstanceRequest struct {
	// example:
	//
	// xgboost数据集加速实例
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 20
	MaxSlot *string `json:"MaxSlot,omitempty" xml:"MaxSlot,omitempty"`
	// example:
	//
	// acc_instance_1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateInstanceRequest) GetMaxSlot() *string {
	return s.MaxSlot
}

func (s *UpdateInstanceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateInstanceRequest) SetDescription(v string) *UpdateInstanceRequest {
	s.Description = &v
	return s
}

func (s *UpdateInstanceRequest) SetMaxSlot(v string) *UpdateInstanceRequest {
	s.MaxSlot = &v
	return s
}

func (s *UpdateInstanceRequest) SetName(v string) *UpdateInstanceRequest {
	s.Name = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
