// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaceEntitiesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDbName(v string) *ListFaceEntitiesRequest
	GetDbName() *string
	SetEntityIdPrefix(v string) *ListFaceEntitiesRequest
	GetEntityIdPrefix() *string
	SetLabels(v string) *ListFaceEntitiesRequest
	GetLabels() *string
	SetLimit(v int32) *ListFaceEntitiesRequest
	GetLimit() *int32
	SetOffset(v int32) *ListFaceEntitiesRequest
	GetOffset() *int32
	SetOrder(v string) *ListFaceEntitiesRequest
	GetOrder() *string
	SetToken(v string) *ListFaceEntitiesRequest
	GetToken() *string
}

type ListFaceEntitiesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// example:
	//
	// U1
	EntityIdPrefix *string `json:"EntityIdPrefix,omitempty" xml:"EntityIdPrefix,omitempty"`
	Labels         *string `json:"Labels,omitempty" xml:"Labels,omitempty"`
	// example:
	//
	// 50
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// example:
	//
	// 1
	Offset *int32 `json:"Offset,omitempty" xml:"Offset,omitempty"`
	// example:
	//
	// asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// example:
	//
	// 2
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s ListFaceEntitiesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFaceEntitiesRequest) GoString() string {
	return s.String()
}

func (s *ListFaceEntitiesRequest) GetDbName() *string {
	return s.DbName
}

func (s *ListFaceEntitiesRequest) GetEntityIdPrefix() *string {
	return s.EntityIdPrefix
}

func (s *ListFaceEntitiesRequest) GetLabels() *string {
	return s.Labels
}

func (s *ListFaceEntitiesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFaceEntitiesRequest) GetOffset() *int32 {
	return s.Offset
}

func (s *ListFaceEntitiesRequest) GetOrder() *string {
	return s.Order
}

func (s *ListFaceEntitiesRequest) GetToken() *string {
	return s.Token
}

func (s *ListFaceEntitiesRequest) SetDbName(v string) *ListFaceEntitiesRequest {
	s.DbName = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetEntityIdPrefix(v string) *ListFaceEntitiesRequest {
	s.EntityIdPrefix = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetLabels(v string) *ListFaceEntitiesRequest {
	s.Labels = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetLimit(v int32) *ListFaceEntitiesRequest {
	s.Limit = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetOffset(v int32) *ListFaceEntitiesRequest {
	s.Offset = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetOrder(v string) *ListFaceEntitiesRequest {
	s.Order = &v
	return s
}

func (s *ListFaceEntitiesRequest) SetToken(v string) *ListFaceEntitiesRequest {
	s.Token = &v
	return s
}

func (s *ListFaceEntitiesRequest) Validate() error {
	return dara.Validate(s)
}
