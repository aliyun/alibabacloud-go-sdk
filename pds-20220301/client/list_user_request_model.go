// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListUserRequest
	GetLimit() *int32
	SetMarker(v string) *ListUserRequest
	GetMarker() *string
}

type ListUserRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListUserRequest) String() string {
	return dara.Prettify(s)
}

func (s ListUserRequest) GoString() string {
	return s.String()
}

func (s *ListUserRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListUserRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListUserRequest) SetLimit(v int32) *ListUserRequest {
	s.Limit = &v
	return s
}

func (s *ListUserRequest) SetMarker(v string) *ListUserRequest {
	s.Marker = &v
	return s
}

func (s *ListUserRequest) Validate() error {
	return dara.Validate(s)
}
