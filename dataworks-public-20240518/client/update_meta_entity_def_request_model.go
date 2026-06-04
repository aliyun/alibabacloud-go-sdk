// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaEntityDefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateMetaEntityDefRequest
	GetDescription() *string
	SetDisplayName(v string) *UpdateMetaEntityDefRequest
	GetDisplayName() *string
	SetEntityType(v string) *UpdateMetaEntityDefRequest
	GetEntityType() *string
	SetNewAttributeDefs(v []*MetaEntityAttributeDef) *UpdateMetaEntityDefRequest
	GetNewAttributeDefs() []*MetaEntityAttributeDef
	SetUpdateAttributeDefs(v []*MetaEntityAttributeDef) *UpdateMetaEntityDefRequest
	GetUpdateAttributeDefs() []*MetaEntityAttributeDef
}

type UpdateMetaEntityDefRequest struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// Business API
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// example:
	//
	// []
	NewAttributeDefs []*MetaEntityAttributeDef `json:"NewAttributeDefs,omitempty" xml:"NewAttributeDefs,omitempty" type:"Repeated"`
	// example:
	//
	// []
	UpdateAttributeDefs []*MetaEntityAttributeDef `json:"UpdateAttributeDefs,omitempty" xml:"UpdateAttributeDefs,omitempty" type:"Repeated"`
}

func (s UpdateMetaEntityDefRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityDefRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityDefRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateMetaEntityDefRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateMetaEntityDefRequest) GetEntityType() *string {
	return s.EntityType
}

func (s *UpdateMetaEntityDefRequest) GetNewAttributeDefs() []*MetaEntityAttributeDef {
	return s.NewAttributeDefs
}

func (s *UpdateMetaEntityDefRequest) GetUpdateAttributeDefs() []*MetaEntityAttributeDef {
	return s.UpdateAttributeDefs
}

func (s *UpdateMetaEntityDefRequest) SetDescription(v string) *UpdateMetaEntityDefRequest {
	s.Description = &v
	return s
}

func (s *UpdateMetaEntityDefRequest) SetDisplayName(v string) *UpdateMetaEntityDefRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateMetaEntityDefRequest) SetEntityType(v string) *UpdateMetaEntityDefRequest {
	s.EntityType = &v
	return s
}

func (s *UpdateMetaEntityDefRequest) SetNewAttributeDefs(v []*MetaEntityAttributeDef) *UpdateMetaEntityDefRequest {
	s.NewAttributeDefs = v
	return s
}

func (s *UpdateMetaEntityDefRequest) SetUpdateAttributeDefs(v []*MetaEntityAttributeDef) *UpdateMetaEntityDefRequest {
	s.UpdateAttributeDefs = v
	return s
}

func (s *UpdateMetaEntityDefRequest) Validate() error {
	if s.NewAttributeDefs != nil {
		for _, item := range s.NewAttributeDefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UpdateAttributeDefs != nil {
		for _, item := range s.UpdateAttributeDefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
