// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateLayerRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateLayerRequest
	GetInstanceId() *string
	SetName(v string) *UpdateLayerRequest
	GetName() *string
}

type UpdateLayerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// This is a test.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pairec-test1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// layer1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLayerRequest) GoString() string {
	return s.String()
}

func (s *UpdateLayerRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateLayerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateLayerRequest) GetName() *string {
	return s.Name
}

func (s *UpdateLayerRequest) SetDescription(v string) *UpdateLayerRequest {
	s.Description = &v
	return s
}

func (s *UpdateLayerRequest) SetInstanceId(v string) *UpdateLayerRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateLayerRequest) SetName(v string) *UpdateLayerRequest {
	s.Name = &v
	return s
}

func (s *UpdateLayerRequest) Validate() error {
	return dara.Validate(s)
}
