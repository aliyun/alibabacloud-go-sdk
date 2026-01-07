// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayType(v string) *ListSecretsRequest
	GetGatewayType() *string
	SetNameLike(v string) *ListSecretsRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListSecretsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSecretsRequest
	GetPageSize() *int32
}

type ListSecretsRequest struct {
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	NameLike    *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	PageNumber  *int32  `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretsRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListSecretsRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListSecretsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSecretsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSecretsRequest) SetGatewayType(v string) *ListSecretsRequest {
	s.GatewayType = &v
	return s
}

func (s *ListSecretsRequest) SetNameLike(v string) *ListSecretsRequest {
	s.NameLike = &v
	return s
}

func (s *ListSecretsRequest) SetPageNumber(v int32) *ListSecretsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecretsRequest) SetPageSize(v int32) *ListSecretsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSecretsRequest) Validate() error {
	return dara.Validate(s)
}
