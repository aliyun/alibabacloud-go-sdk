// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstallableGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayType(v string) *ListInstallableGatewaysRequest
	GetGatewayType() *string
	SetPageNumber(v int32) *ListInstallableGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListInstallableGatewaysRequest
	GetPageSize() *int32
}

type ListInstallableGatewaysRequest struct {
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
}

func (s ListInstallableGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstallableGatewaysRequest) GoString() string {
	return s.String()
}

func (s *ListInstallableGatewaysRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListInstallableGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListInstallableGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInstallableGatewaysRequest) SetGatewayType(v string) *ListInstallableGatewaysRequest {
	s.GatewayType = &v
	return s
}

func (s *ListInstallableGatewaysRequest) SetPageNumber(v int32) *ListInstallableGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListInstallableGatewaysRequest) SetPageSize(v int32) *ListInstallableGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *ListInstallableGatewaysRequest) Validate() error {
	return dara.Validate(s)
}
