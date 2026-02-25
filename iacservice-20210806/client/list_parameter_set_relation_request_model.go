// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterSetRelationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceId(v string) *ListParameterSetRelationRequest
	GetResourceId() *string
	SetResourceType(v string) *ListParameterSetRelationRequest
	GetResourceType() *string
}

type ListParameterSetRelationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mod-edf123fr-v3
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Module
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListParameterSetRelationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListParameterSetRelationRequest) GoString() string {
	return s.String()
}

func (s *ListParameterSetRelationRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListParameterSetRelationRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListParameterSetRelationRequest) SetResourceId(v string) *ListParameterSetRelationRequest {
	s.ResourceId = &v
	return s
}

func (s *ListParameterSetRelationRequest) SetResourceType(v string) *ListParameterSetRelationRequest {
	s.ResourceType = &v
	return s
}

func (s *ListParameterSetRelationRequest) Validate() error {
	return dara.Validate(s)
}
