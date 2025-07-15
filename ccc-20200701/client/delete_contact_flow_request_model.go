// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactFlowId(v string) *DeleteContactFlowRequest
	GetContactFlowId() *string
	SetForce(v bool) *DeleteContactFlowRequest
	GetForce() *bool
	SetInstanceId(v string) *DeleteContactFlowRequest
	GetInstanceId() *string
}

type DeleteContactFlowRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0f87c997-b0c1-41d4-9e9e-1b791de6ad1f
	ContactFlowId *string `json:"ContactFlowId,omitempty" xml:"ContactFlowId,omitempty"`
	Force         *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteContactFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactFlowRequest) GoString() string {
	return s.String()
}

func (s *DeleteContactFlowRequest) GetContactFlowId() *string {
	return s.ContactFlowId
}

func (s *DeleteContactFlowRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteContactFlowRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteContactFlowRequest) SetContactFlowId(v string) *DeleteContactFlowRequest {
	s.ContactFlowId = &v
	return s
}

func (s *DeleteContactFlowRequest) SetForce(v bool) *DeleteContactFlowRequest {
	s.Force = &v
	return s
}

func (s *DeleteContactFlowRequest) SetInstanceId(v string) *DeleteContactFlowRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteContactFlowRequest) Validate() error {
	return dara.Validate(s)
}
