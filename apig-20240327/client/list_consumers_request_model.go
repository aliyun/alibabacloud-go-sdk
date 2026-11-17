// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExpireStatus(v string) *ListConsumersRequest
	GetExpireStatus() *string
	SetGatewayType(v string) *ListConsumersRequest
	GetGatewayType() *string
	SetIncludeExpirationSummary(v bool) *ListConsumersRequest
	GetIncludeExpirationSummary() *bool
	SetNameLike(v string) *ListConsumersRequest
	GetNameLike() *string
	SetPageNumber(v int32) *ListConsumersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListConsumersRequest
	GetPageSize() *int32
	SetStatus(v string) *ListConsumersRequest
	GetStatus() *string
}

type ListConsumersRequest struct {
	// The validity status filter. Valid values: Permanent, InEffect, AboutToExpire, and Expired. You can specify multiple values separated by commas to form a union. If this parameter is not specified, all statuses are included and intersected with the name and display status filters.
	//
	// example:
	//
	// Expired,AboutToExpire
	ExpireStatus *string `json:"expireStatus,omitempty" xml:"expireStatus,omitempty"`
	// The gateway type used for filtering. Currently, **AI*	- and **API*	- gateway types are supported.
	//
	// example:
	//
	// AI
	GatewayType *string `json:"gatewayType,omitempty" xml:"gatewayType,omitempty"`
	// Specifies whether to return the full expiration reminder count. Default value: false. The statistical scope is limited only by the current account and gatewayType.
	//
	// example:
	//
	// true
	IncludeExpirationSummary *bool `json:"includeExpirationSummary,omitempty" xml:"includeExpirationSummary,omitempty"`
	// The keyword used for fuzzy search by API name.
	//
	// example:
	//
	// UI-test
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The page number. Pages start from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The display status filter. Valid values: Enabled (enabled and not expired), Disabled (disabled regardless of expiration), and Expired (enabled and expired). If this parameter is not specified, all statuses are included and intersected with other filter conditions.
	//
	// example:
	//
	// Enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListConsumersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumersRequest) GoString() string {
	return s.String()
}

func (s *ListConsumersRequest) GetExpireStatus() *string {
	return s.ExpireStatus
}

func (s *ListConsumersRequest) GetGatewayType() *string {
	return s.GatewayType
}

func (s *ListConsumersRequest) GetIncludeExpirationSummary() *bool {
	return s.IncludeExpirationSummary
}

func (s *ListConsumersRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListConsumersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumersRequest) GetStatus() *string {
	return s.Status
}

func (s *ListConsumersRequest) SetExpireStatus(v string) *ListConsumersRequest {
	s.ExpireStatus = &v
	return s
}

func (s *ListConsumersRequest) SetGatewayType(v string) *ListConsumersRequest {
	s.GatewayType = &v
	return s
}

func (s *ListConsumersRequest) SetIncludeExpirationSummary(v bool) *ListConsumersRequest {
	s.IncludeExpirationSummary = &v
	return s
}

func (s *ListConsumersRequest) SetNameLike(v string) *ListConsumersRequest {
	s.NameLike = &v
	return s
}

func (s *ListConsumersRequest) SetPageNumber(v int32) *ListConsumersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListConsumersRequest) SetPageSize(v int32) *ListConsumersRequest {
	s.PageSize = &v
	return s
}

func (s *ListConsumersRequest) SetStatus(v string) *ListConsumersRequest {
	s.Status = &v
	return s
}

func (s *ListConsumersRequest) Validate() error {
	return dara.Validate(s)
}
