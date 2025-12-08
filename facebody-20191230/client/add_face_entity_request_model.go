// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *AddFaceEntityRequest
	GetDbName() *string
	SetEntityId(v string) *AddFaceEntityRequest
	GetEntityId() *string
	SetLabels(v string) *AddFaceEntityRequest
	GetLabels() *string
}

type AddFaceEntityRequest struct {
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

func (s AddFaceEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *AddFaceEntityRequest) GetDbName() *string {
	return s.DbName
}

func (s *AddFaceEntityRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *AddFaceEntityRequest) GetLabels() *string {
	return s.Labels
}

func (s *AddFaceEntityRequest) SetDbName(v string) *AddFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *AddFaceEntityRequest) SetEntityId(v string) *AddFaceEntityRequest {
	s.EntityId = &v
	return s
}

func (s *AddFaceEntityRequest) SetLabels(v string) *AddFaceEntityRequest {
	s.Labels = &v
	return s
}

func (s *AddFaceEntityRequest) Validate() error {
	return dara.Validate(s)
}
