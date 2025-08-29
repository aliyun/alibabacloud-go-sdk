// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrowdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateCrowdRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateCrowdRequest
	GetInstanceId() *string
	SetName(v string) *UpdateCrowdRequest
	GetName() *string
}

type UpdateCrowdRequest struct {
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
	// xx人群
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateCrowdRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrowdRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrowdRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateCrowdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateCrowdRequest) GetName() *string {
	return s.Name
}

func (s *UpdateCrowdRequest) SetDescription(v string) *UpdateCrowdRequest {
	s.Description = &v
	return s
}

func (s *UpdateCrowdRequest) SetInstanceId(v string) *UpdateCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateCrowdRequest) SetName(v string) *UpdateCrowdRequest {
	s.Name = &v
	return s
}

func (s *UpdateCrowdRequest) Validate() error {
	return dara.Validate(s)
}
