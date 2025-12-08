// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *DeleteFaceEntityRequest
	GetDbName() *string
	SetEntityId(v string) *DeleteFaceEntityRequest
	GetEntityId() *string
}

type DeleteFaceEntityRequest struct {
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
	// wood
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s DeleteFaceEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *DeleteFaceEntityRequest) GetDbName() *string {
	return s.DbName
}

func (s *DeleteFaceEntityRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *DeleteFaceEntityRequest) SetDbName(v string) *DeleteFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *DeleteFaceEntityRequest) SetEntityId(v string) *DeleteFaceEntityRequest {
	s.EntityId = &v
	return s
}

func (s *DeleteFaceEntityRequest) Validate() error {
	return dara.Validate(s)
}
