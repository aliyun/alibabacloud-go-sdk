// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributes(v map[string]*string) *UpdateMetaEntityRequest
	GetAttributes() map[string]*string
	SetComment(v string) *UpdateMetaEntityRequest
	GetComment() *string
	SetCustomAttributes(v map[string][]*string) *UpdateMetaEntityRequest
	GetCustomAttributes() map[string][]*string
	SetId(v string) *UpdateMetaEntityRequest
	GetId() *string
}

type UpdateMetaEntityRequest struct {
	// example:
	//
	// []
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// []
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api:api_001
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateMetaEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaEntityRequest) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *UpdateMetaEntityRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateMetaEntityRequest) GetCustomAttributes() map[string][]*string {
	return s.CustomAttributes
}

func (s *UpdateMetaEntityRequest) GetId() *string {
	return s.Id
}

func (s *UpdateMetaEntityRequest) SetAttributes(v map[string]*string) *UpdateMetaEntityRequest {
	s.Attributes = v
	return s
}

func (s *UpdateMetaEntityRequest) SetComment(v string) *UpdateMetaEntityRequest {
	s.Comment = &v
	return s
}

func (s *UpdateMetaEntityRequest) SetCustomAttributes(v map[string][]*string) *UpdateMetaEntityRequest {
	s.CustomAttributes = v
	return s
}

func (s *UpdateMetaEntityRequest) SetId(v string) *UpdateMetaEntityRequest {
	s.Id = &v
	return s
}

func (s *UpdateMetaEntityRequest) Validate() error {
	return dara.Validate(s)
}
