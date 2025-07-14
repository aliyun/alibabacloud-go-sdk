// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataServiceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *QueryDataServiceListRequest
	GetName() *string
	SetPageNo(v int32) *QueryDataServiceListRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryDataServiceListRequest
	GetPageSize() *int32
	SetUserId(v string) *QueryDataServiceListRequest
	GetUserId() *string
}

type QueryDataServiceListRequest struct {
	// Data service name.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Number of items per page in a paginated query:
	//
	// - Default value: 10
	//
	// - Maximum value: 1000
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// User ID.
	//
	// example:
	//
	// dasdfdsa-csddf-dsadsa
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryDataServiceListRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListRequest) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListRequest) GetName() *string {
	return s.Name
}

func (s *QueryDataServiceListRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryDataServiceListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryDataServiceListRequest) GetUserId() *string {
	return s.UserId
}

func (s *QueryDataServiceListRequest) SetName(v string) *QueryDataServiceListRequest {
	s.Name = &v
	return s
}

func (s *QueryDataServiceListRequest) SetPageNo(v int32) *QueryDataServiceListRequest {
	s.PageNo = &v
	return s
}

func (s *QueryDataServiceListRequest) SetPageSize(v int32) *QueryDataServiceListRequest {
	s.PageSize = &v
	return s
}

func (s *QueryDataServiceListRequest) SetUserId(v string) *QueryDataServiceListRequest {
	s.UserId = &v
	return s
}

func (s *QueryDataServiceListRequest) Validate() error {
	return dara.Validate(s)
}
