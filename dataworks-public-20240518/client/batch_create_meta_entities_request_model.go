// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchCreateMetaEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntities(v []*BatchCreateMetaEntitiesRequestEntities) *BatchCreateMetaEntitiesRequest
	GetEntities() []*BatchCreateMetaEntitiesRequestEntities
}

type BatchCreateMetaEntitiesRequest struct {
	// The list of entities. A maximum of five entities are supported. All entities in the same batch must have the same entityType.
	//
	// This parameter is required.
	//
	// example:
	//
	// []
	Entities []*BatchCreateMetaEntitiesRequestEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Repeated"`
}

func (s BatchCreateMetaEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMetaEntitiesRequest) GoString() string {
	return s.String()
}

func (s *BatchCreateMetaEntitiesRequest) GetEntities() []*BatchCreateMetaEntitiesRequestEntities {
	return s.Entities
}

func (s *BatchCreateMetaEntitiesRequest) SetEntities(v []*BatchCreateMetaEntitiesRequestEntities) *BatchCreateMetaEntitiesRequest {
	s.Entities = v
	return s
}

func (s *BatchCreateMetaEntitiesRequest) Validate() error {
	if s.Entities != nil {
		for _, item := range s.Entities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchCreateMetaEntitiesRequestEntities struct {
	// The entity attributes. Complex values must be serialized as JSON strings.
	Attributes map[string]*string `json:"Attributes,omitempty" xml:"Attributes,omitempty"`
	// The comment.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The custom attribute values. The key is the custom attribute identifier, and the value currently supports only a single value.
	//
	// <notice>The custom attributes used here must be created in advance by calling the CreateCustomAttribute operation. For example, after you call the API to create a custom attribute with the ID `custom-attribute:owner_name`, you can configure {\\"owner_name\\": [\\"Bob\\"]} here to complete the custom attribute configuration.</notice>
	CustomAttributes map[string][]*string `json:"CustomAttributes,omitempty" xml:"CustomAttributes,omitempty"`
	// The entity type. All entities in the same batch must have the same type. The following types are supported:
	//
	// - Custom entity types, such as custom_entity-biz_api.
	//
	// - Extension table types. If the metadata entity type custom_dw-table is registered, you can create objects of the corresponding database type custom_dw-database and table type custom_dw-table.
	//
	// This parameter is required.
	//
	// example:
	//
	// custom_entity-customer_api
	EntityType *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	// The entity name. The name can contain uppercase letters, lowercase letters, digits, and underscores (_). It must start with a letter and can be up to 64 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// api_001
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s BatchCreateMetaEntitiesRequestEntities) String() string {
	return dara.Prettify(s)
}

func (s BatchCreateMetaEntitiesRequestEntities) GoString() string {
	return s.String()
}

func (s *BatchCreateMetaEntitiesRequestEntities) GetAttributes() map[string]*string {
	return s.Attributes
}

func (s *BatchCreateMetaEntitiesRequestEntities) GetComment() *string {
	return s.Comment
}

func (s *BatchCreateMetaEntitiesRequestEntities) GetCustomAttributes() map[string][]*string {
	return s.CustomAttributes
}

func (s *BatchCreateMetaEntitiesRequestEntities) GetEntityType() *string {
	return s.EntityType
}

func (s *BatchCreateMetaEntitiesRequestEntities) GetName() *string {
	return s.Name
}

func (s *BatchCreateMetaEntitiesRequestEntities) SetAttributes(v map[string]*string) *BatchCreateMetaEntitiesRequestEntities {
	s.Attributes = v
	return s
}

func (s *BatchCreateMetaEntitiesRequestEntities) SetComment(v string) *BatchCreateMetaEntitiesRequestEntities {
	s.Comment = &v
	return s
}

func (s *BatchCreateMetaEntitiesRequestEntities) SetCustomAttributes(v map[string][]*string) *BatchCreateMetaEntitiesRequestEntities {
	s.CustomAttributes = v
	return s
}

func (s *BatchCreateMetaEntitiesRequestEntities) SetEntityType(v string) *BatchCreateMetaEntitiesRequestEntities {
	s.EntityType = &v
	return s
}

func (s *BatchCreateMetaEntitiesRequestEntities) SetName(v string) *BatchCreateMetaEntitiesRequestEntities {
	s.Name = &v
	return s
}

func (s *BatchCreateMetaEntitiesRequestEntities) Validate() error {
	return dara.Validate(s)
}
