// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateLayerRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateLayerRequest
	GetInstanceId() *string
	SetLaboratoryId(v string) *CreateLayerRequest
	GetLaboratoryId() *string
	SetName(v string) *CreateLayerRequest
	GetName() *string
}

type CreateLayerRequest struct {
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
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3
	LaboratoryId *string `json:"LaboratoryId,omitempty" xml:"LaboratoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// layer1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLayerRequest) GoString() string {
	return s.String()
}

func (s *CreateLayerRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLayerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateLayerRequest) GetLaboratoryId() *string {
	return s.LaboratoryId
}

func (s *CreateLayerRequest) GetName() *string {
	return s.Name
}

func (s *CreateLayerRequest) SetDescription(v string) *CreateLayerRequest {
	s.Description = &v
	return s
}

func (s *CreateLayerRequest) SetInstanceId(v string) *CreateLayerRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateLayerRequest) SetLaboratoryId(v string) *CreateLayerRequest {
	s.LaboratoryId = &v
	return s
}

func (s *CreateLayerRequest) SetName(v string) *CreateLayerRequest {
	s.Name = &v
	return s
}

func (s *CreateLayerRequest) Validate() error {
	return dara.Validate(s)
}
