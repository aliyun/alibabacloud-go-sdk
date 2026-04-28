// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreator(v string) *ListShareLinkRequest
	GetCreator() *string
	SetIncludeCancelled(v bool) *ListShareLinkRequest
	GetIncludeCancelled() *bool
	SetLimit(v int32) *ListShareLinkRequest
	GetLimit() *int32
	SetMarker(v string) *ListShareLinkRequest
	GetMarker() *string
	SetOrderBy(v string) *ListShareLinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *ListShareLinkRequest
	GetOrderDirection() *string
}

type ListShareLinkRequest struct {
	// The creator of the share.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Specifies whether to return the shares that are canceled.
	//
	// example:
	//
	// true
	IncludeCancelled *bool `json:"include_cancelled,omitempty" xml:"include_cancelled,omitempty"`
	// The maximum number of results to return. Valid values: 0 to 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The field by which to sort the returned results. Default value: created_at. Valid values:
	//
	// 	- share_name: sorts the results by the name of the share.
	//
	// 	- updated_at: sorts the results by the time when the share was modified.
	//
	// 	- description: sorts the results by the description of the share.
	//
	// 	- created_at: sorts the results by the time when the share was created.
	//
	// example:
	//
	// share_name
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// The order in which you want to sort the returned results. By default, order_direction is set to DESC if order_by is set to created_at or updated_at, and is set to ASC if order_by is set to other values. Valid values:
	//
	// 	- ASC: sorts the results in ascending order.
	//
	// 	- DESC: sorts the results in descending order.
	//
	// example:
	//
	// ASC
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
}

func (s ListShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListShareLinkRequest) GoString() string {
	return s.String()
}

func (s *ListShareLinkRequest) GetCreator() *string {
	return s.Creator
}

func (s *ListShareLinkRequest) GetIncludeCancelled() *bool {
	return s.IncludeCancelled
}

func (s *ListShareLinkRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListShareLinkRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListShareLinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *ListShareLinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *ListShareLinkRequest) SetCreator(v string) *ListShareLinkRequest {
	s.Creator = &v
	return s
}

func (s *ListShareLinkRequest) SetIncludeCancelled(v bool) *ListShareLinkRequest {
	s.IncludeCancelled = &v
	return s
}

func (s *ListShareLinkRequest) SetLimit(v int32) *ListShareLinkRequest {
	s.Limit = &v
	return s
}

func (s *ListShareLinkRequest) SetMarker(v string) *ListShareLinkRequest {
	s.Marker = &v
	return s
}

func (s *ListShareLinkRequest) SetOrderBy(v string) *ListShareLinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListShareLinkRequest) SetOrderDirection(v string) *ListShareLinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
