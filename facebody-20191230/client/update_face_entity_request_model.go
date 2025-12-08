// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFaceEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *UpdateFaceEntityRequest
	GetDbName() *string
	SetEntityId(v string) *UpdateFaceEntityRequest
	GetEntityId() *string
	SetLabels(v string) *UpdateFaceEntityRequest
	GetLabels() *string
}

type UpdateFaceEntityRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mm2
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	Labels   *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
}

func (s UpdateFaceEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *UpdateFaceEntityRequest) GetDbName() *string {
	return s.DbName
}

func (s *UpdateFaceEntityRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *UpdateFaceEntityRequest) GetLabels() *string {
	return s.Labels
}

func (s *UpdateFaceEntityRequest) SetDbName(v string) *UpdateFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *UpdateFaceEntityRequest) SetEntityId(v string) *UpdateFaceEntityRequest {
	s.EntityId = &v
	return s
}

func (s *UpdateFaceEntityRequest) SetLabels(v string) *UpdateFaceEntityRequest {
	s.Labels = &v
	return s
}

func (s *UpdateFaceEntityRequest) Validate() error {
	return dara.Validate(s)
}
