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
	// example:
	//
	// []
	AttributeDefs []*MetaEntityAttributeDef `json:"AttributeDefs,omitempty" xml:"AttributeDefs,omitempty" type:"Repeated"`
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Business API
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// NONE
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
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
