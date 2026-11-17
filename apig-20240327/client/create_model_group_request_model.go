// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateModelGroupRequest
	GetDescription() *string
	SetGatewayId(v string) *CreateModelGroupRequest
	GetGatewayId() *string
	SetModelCardIds(v []*string) *CreateModelGroupRequest
	GetModelCardIds() []*string
	SetName(v string) *CreateModelGroupRequest
	GetName() *string
}

type CreateModelGroupRequest struct {
	// The description of the model group. The description can be up to 1,024 UTF-16 code units in length and can be left empty.
	//
	// example:
	//
	// Model group for service integration
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the AI gateway instance to which the resource belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-d00000000000000000000
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The IDs of the actual or logical model cards within the same instance. The list can be empty and must not contain duplicate values. Model groups are not accepted as members.
	//
	// This parameter is required.
	ModelCardIds []*string `json:"modelCardIds,omitempty" xml:"modelCardIds,omitempty" type:"Repeated"`
	// The unique name within the instance. The name must be 1 to 128 UTF-8 bytes in length after trimming.
	//
	// This parameter is required.
	//
	// example:
	//
	// beta-model-group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s CreateModelGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateModelGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateModelGroupRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *CreateModelGroupRequest) GetModelCardIds() []*string {
	return s.ModelCardIds
}

func (s *CreateModelGroupRequest) GetName() *string {
	return s.Name
}

func (s *CreateModelGroupRequest) SetDescription(v string) *CreateModelGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateModelGroupRequest) SetGatewayId(v string) *CreateModelGroupRequest {
	s.GatewayId = &v
	return s
}

func (s *CreateModelGroupRequest) SetModelCardIds(v []*string) *CreateModelGroupRequest {
	s.ModelCardIds = v
	return s
}

func (s *CreateModelGroupRequest) SetName(v string) *CreateModelGroupRequest {
	s.Name = &v
	return s
}

func (s *CreateModelGroupRequest) Validate() error {
	return dara.Validate(s)
}
