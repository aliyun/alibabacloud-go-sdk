// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPage(v int32) *ListConfigsRequest
	GetPage() *int32
	SetPageSize(v int32) *ListConfigsRequest
	GetPageSize() *int32
}

type ListConfigsRequest struct {
	Page     *int32 `json:"Page,omitempty" xml:"Page,omitempty"`
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListConfigsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConfigsRequest) GoString() string {
	return s.String()
}

func (s *ListConfigsRequest) GetPage() *int32 {
	return s.Page
}

func (s *ListConfigsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConfigsRequest) SetPage(v int32) *ListConfigsRequest {
	s.Page = &v
	return s
}

func (s *ListConfigsRequest) SetPageSize(v int32) *ListConfigsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConfigsRequest) Validate() error {
	return dara.Validate(s)
}
