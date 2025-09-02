// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataServiceApiTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v int64) *ListDataServiceApiTestRequest
	GetApiId() *int64
	SetPageSize(v int32) *ListDataServiceApiTestRequest
	GetPageSize() *int32
}

type ListDataServiceApiTestRequest struct {
	// The ID of the DataService Studio API on which tests are performed.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The number of entries to return on each page. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListDataServiceApiTestRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataServiceApiTestRequest) GoString() string {
	return s.String()
}

func (s *ListDataServiceApiTestRequest) GetApiId() *int64 {
	return s.ApiId
}

func (s *ListDataServiceApiTestRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDataServiceApiTestRequest) SetApiId(v int64) *ListDataServiceApiTestRequest {
	s.ApiId = &v
	return s
}

func (s *ListDataServiceApiTestRequest) SetPageSize(v int32) *ListDataServiceApiTestRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataServiceApiTestRequest) Validate() error {
	return dara.Validate(s)
}
