// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaEntityDefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *DeleteMetaEntityDefRequest
	GetEntityType() *string
	SetForce(v bool) *DeleteMetaEntityDefRequest
	GetForce() *bool
}

type DeleteMetaEntityDefRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// true
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
}

func (s DeleteMetaEntityDefRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaEntityDefRequest) GoString() string {
	return s.String()
}

func (s *DeleteMetaEntityDefRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *DeleteMetaEntityDefRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteMetaEntityDefRequest) SetEntityType(v string) *DeleteMetaEntityDefRequest {
	s.EntityType = &v
	return s
}

func (s *DeleteMetaEntityDefRequest) SetForce(v bool) *DeleteMetaEntityDefRequest {
	s.Force = &v
	return s
}

func (s *DeleteMetaEntityDefRequest) Validate() error {
	return dara.Validate(s)
}
