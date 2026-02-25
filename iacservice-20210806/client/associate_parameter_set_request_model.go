// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateParameterSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetParameterSetIds(v []*string) *AssociateParameterSetRequest
	GetParameterSetIds() []*string
	SetResourceId(v string) *AssociateParameterSetRequest
	GetResourceId() *string
	SetResourceType(v string) *AssociateParameterSetRequest
	GetResourceType() *string
}

type AssociateParameterSetRequest struct {
	// This parameter is required.
	ParameterSetIds []*string `json:"parameterSetIds,omitempty" xml:"parameterSetIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// task-433aead756057ffdf5326bf1e12ed
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Task
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s AssociateParameterSetRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateParameterSetRequest) GoString() string {
	return s.String()
}

func (s *AssociateParameterSetRequest) GetParameterSetIds() []*string {
	return s.ParameterSetIds
}

func (s *AssociateParameterSetRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AssociateParameterSetRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AssociateParameterSetRequest) SetParameterSetIds(v []*string) *AssociateParameterSetRequest {
	s.ParameterSetIds = v
	return s
}

func (s *AssociateParameterSetRequest) SetResourceId(v string) *AssociateParameterSetRequest {
	s.ResourceId = &v
	return s
}

func (s *AssociateParameterSetRequest) SetResourceType(v string) *AssociateParameterSetRequest {
	s.ResourceType = &v
	return s
}

func (s *AssociateParameterSetRequest) Validate() error {
	return dara.Validate(s)
}
