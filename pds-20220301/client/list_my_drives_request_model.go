// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMyDrivesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListMyDrivesRequest
	GetLimit() *int32
	SetMarker(v string) *ListMyDrivesRequest
	GetMarker() *string
}

type ListMyDrivesRequest struct {
	// The maximum number of results to return. Default value: 100. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListMyDrivesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMyDrivesRequest) GoString() string {
	return s.String()
}

func (s *ListMyDrivesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListMyDrivesRequest) GetMarker() *string {
	return s.Marker
}

func (s *ListMyDrivesRequest) SetLimit(v int32) *ListMyDrivesRequest {
	s.Limit = &v
	return s
}

func (s *ListMyDrivesRequest) SetMarker(v string) *ListMyDrivesRequest {
	s.Marker = &v
	return s
}

func (s *ListMyDrivesRequest) Validate() error {
	return dara.Validate(s)
}
