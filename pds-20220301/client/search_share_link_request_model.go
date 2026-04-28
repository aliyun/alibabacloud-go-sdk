// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchShareLinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreators(v []*string) *SearchShareLinkRequest
	GetCreators() []*string
	SetLimit(v int32) *SearchShareLinkRequest
	GetLimit() *int32
	SetMarker(v string) *SearchShareLinkRequest
	GetMarker() *string
	SetOrderBy(v string) *SearchShareLinkRequest
	GetOrderBy() *string
	SetOrderDirection(v string) *SearchShareLinkRequest
	GetOrderDirection() *string
	SetQuery(v string) *SearchShareLinkRequest
	GetQuery() *string
	SetReturnTotalCount(v bool) *SearchShareLinkRequest
	GetReturnTotalCount() *bool
}

type SearchShareLinkRequest struct {
	// The creators of shares. Set this parameter to a user ID. If you do not specify this parameter, Drive and Photo Service automatically queries shares based on your permissions. If you are a drive administrator or the super administrator, the shares created by all members are queried. If you are a team administrator, the shares created by all team members are queried. If you are a common user, only the shares created by yourself are queried.
	//
	// If you specify this parameter, this parameter must be set to the ID of the super administrator, a drive administrator, or a team administrator.
	Creators []*string `json:"creators,omitempty" xml:"creators,omitempty" type:"Repeated"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
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
	// created_at
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
	// The query condition that is used to search for share URLs. You can use the following fields to specify query conditions: created_at: queries a share URL based on the time when the share URL was created. updated_at: queries a share URL based on the time when the share URL was modified. share_name_for_fuzzy: queries a share URL based on the name of the share. A fuzzy match is supported. status: queries a share URL based on the status of the share. Valid values: enabled and disabled. A value of enabled indicates that the share is valid. A value of disabled indicates that the share is canceled. expired_time: queries a share URL based on the expiration time of the share. If the share never expires, set this field to 1970-01-01T00:00:00. Otherwise, set this field to 1970-01-02T00:00:00.
	//
	// example:
	//
	// created_at>=\\"2022-01-18T02:50:00\\" and created_at<\\"2022-01-19T09:45:28\\" and share_name_for_fuzzy match \\"HD.mp4\\" and status in [\\"enabled\\", \\"disabled\\"] and expired_time=\\"1970-01-01T00:00:00\\"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Specifies whether to return the total number of returned results.
	//
	// example:
	//
	// false
	ReturnTotalCount *bool `json:"return_total_count,omitempty" xml:"return_total_count,omitempty"`
}

func (s SearchShareLinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchShareLinkRequest) GoString() string {
	return s.String()
}

func (s *SearchShareLinkRequest) GetCreators() []*string {
	return s.Creators
}

func (s *SearchShareLinkRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *SearchShareLinkRequest) GetMarker() *string {
	return s.Marker
}

func (s *SearchShareLinkRequest) GetOrderBy() *string {
	return s.OrderBy
}

func (s *SearchShareLinkRequest) GetOrderDirection() *string {
	return s.OrderDirection
}

func (s *SearchShareLinkRequest) GetQuery() *string {
	return s.Query
}

func (s *SearchShareLinkRequest) GetReturnTotalCount() *bool {
	return s.ReturnTotalCount
}

func (s *SearchShareLinkRequest) SetCreators(v []*string) *SearchShareLinkRequest {
	s.Creators = v
	return s
}

func (s *SearchShareLinkRequest) SetLimit(v int32) *SearchShareLinkRequest {
	s.Limit = &v
	return s
}

func (s *SearchShareLinkRequest) SetMarker(v string) *SearchShareLinkRequest {
	s.Marker = &v
	return s
}

func (s *SearchShareLinkRequest) SetOrderBy(v string) *SearchShareLinkRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchShareLinkRequest) SetOrderDirection(v string) *SearchShareLinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *SearchShareLinkRequest) SetQuery(v string) *SearchShareLinkRequest {
	s.Query = &v
	return s
}

func (s *SearchShareLinkRequest) SetReturnTotalCount(v bool) *SearchShareLinkRequest {
	s.ReturnTotalCount = &v
	return s
}

func (s *SearchShareLinkRequest) Validate() error {
	return dara.Validate(s)
}
