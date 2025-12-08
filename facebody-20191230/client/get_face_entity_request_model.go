// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFaceEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *GetFaceEntityRequest
	GetDbName() *string
	SetEntityId(v string) *GetFaceEntityRequest
	GetEntityId() *string
}

type GetFaceEntityRequest struct {
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
	// 66
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
}

func (s GetFaceEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFaceEntityRequest) GoString() string {
	return s.String()
}

func (s *GetFaceEntityRequest) GetDbName() *string {
	return s.DbName
}

func (s *GetFaceEntityRequest) GetEntityId() *string {
	return s.EntityId
}

func (s *GetFaceEntityRequest) SetDbName(v string) *GetFaceEntityRequest {
	s.DbName = &v
	return s
}

func (s *GetFaceEntityRequest) SetEntityId(v string) *GetFaceEntityRequest {
	s.EntityId = &v
	return s
}

func (s *GetFaceEntityRequest) Validate() error {
	return dara.Validate(s)
}
