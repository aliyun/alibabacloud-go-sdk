// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListResourceServicesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceServicesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourceServicesResponseBody
	GetRequestId() *string
	SetServices(v []*Service) *ListResourceServicesResponseBody
	GetServices() []*Service
	SetTotalCount(v int32) *ListResourceServicesResponseBody
	GetTotalCount() *int32
}

type ListResourceServicesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The services.
	Services []*Service `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListResourceServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceServicesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceServicesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceServicesResponseBody) GetServices() []*Service {
	return s.Services
}

func (s *ListResourceServicesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceServicesResponseBody) SetPageNumber(v int32) *ListResourceServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetPageSize(v int32) *ListResourceServicesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetRequestId(v string) *ListResourceServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceServicesResponseBody) SetServices(v []*Service) *ListResourceServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListResourceServicesResponseBody) SetTotalCount(v int32) *ListResourceServicesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceServicesResponseBody) Validate() error {
	return dara.Validate(s)
}
