// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaEntityDefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeDefs(v []*MetaEntityAttributeDef) *CreateMetaEntityDefRequest
	GetAttributeDefs() []*MetaEntityAttributeDef
	SetDescription(v string) *CreateMetaEntityDefRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateMetaEntityDefRequest
	GetDisplayName() *string
	SetExtend(v string) *CreateMetaEntityDefRequest
	GetExtend() *string
	SetName(v string) *CreateMetaEntityDefRequest
	GetName() *string
}

type CreateMetaEntityDefRequest struct {
	// A list of attribute definitions for the pure custom type. Do not specify this parameter if the `Extend` parameter is set to `TABLE`.
	//
	// example:
	//
	// []
	AttributeDefs []*MetaEntityAttributeDef `json:"AttributeDefs,omitempty" xml:"AttributeDefs,omitempty" type:"Repeated"`
	// A description of the entity definition.
	//
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name. The maximum length is 32 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business API
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The extension mode. Valid values:
	//
	// - `NONE`: The default value. Specifies a pure custom type with user-defined attributes.
	//
	// - `TABLE`: Specifies an extended table type that references an existing table type in Data Map. Attribute definitions are not required for this type. You can create corresponding `Database` and `Table` objects for it.
	//
	// example:
	//
	// NONE
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The type definition name. For a pure custom type, the name must match `^[a-z0-9][a-z0-9_]*$`. For an extended table type, the name must match `^[a-z0-9][a-z0-9_]*-table$`.
	//
	// This parameter is required.
	//
	// example:
	//
	// biz_api
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateMetaEntityDefRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaEntityDefRequest) GoString() string {
	return s.String()
}

func (s *CreateMetaEntityDefRequest) GetAttributeDefs() []*MetaEntityAttributeDef {
	return s.AttributeDefs
}

func (s *CreateMetaEntityDefRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateMetaEntityDefRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateMetaEntityDefRequest) GetExtend() *string {
	return s.Extend
}

func (s *CreateMetaEntityDefRequest) GetName() *string {
	return s.Name
}

func (s *CreateMetaEntityDefRequest) SetAttributeDefs(v []*MetaEntityAttributeDef) *CreateMetaEntityDefRequest {
	s.AttributeDefs = v
	return s
}

func (s *CreateMetaEntityDefRequest) SetDescription(v string) *CreateMetaEntityDefRequest {
	s.Description = &v
	return s
}

func (s *CreateMetaEntityDefRequest) SetDisplayName(v string) *CreateMetaEntityDefRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateMetaEntityDefRequest) SetExtend(v string) *CreateMetaEntityDefRequest {
	s.Extend = &v
	return s
}

func (s *CreateMetaEntityDefRequest) SetName(v string) *CreateMetaEntityDefRequest {
	s.Name = &v
	return s
}

func (s *CreateMetaEntityDefRequest) Validate() error {
	if s.AttributeDefs != nil {
		for _, item := range s.AttributeDefs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
