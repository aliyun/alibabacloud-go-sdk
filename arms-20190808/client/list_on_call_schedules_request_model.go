// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOnCallSchedulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListOnCallSchedulesRequest
	GetName() *string
	SetPage(v int64) *ListOnCallSchedulesRequest
	GetPage() *int64
	SetSize(v int64) *ListOnCallSchedulesRequest
	GetSize() *int64
}

type ListOnCallSchedulesRequest struct {
	// The name of the scheduling policy.
	//
	// example:
	//
	// OnCallSchedule_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListOnCallSchedulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOnCallSchedulesRequest) GoString() string {
	return s.String()
}

func (s *ListOnCallSchedulesRequest) GetName() *string {
	return s.Name
}

func (s *ListOnCallSchedulesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListOnCallSchedulesRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListOnCallSchedulesRequest) SetName(v string) *ListOnCallSchedulesRequest {
	s.Name = &v
	return s
}

func (s *ListOnCallSchedulesRequest) SetPage(v int64) *ListOnCallSchedulesRequest {
	s.Page = &v
	return s
}

func (s *ListOnCallSchedulesRequest) SetSize(v int64) *ListOnCallSchedulesRequest {
	s.Size = &v
	return s
}

func (s *ListOnCallSchedulesRequest) Validate() error {
	return dara.Validate(s)
}
