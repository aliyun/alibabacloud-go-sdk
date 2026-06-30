// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelOperatorServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListModelOperatorServicesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *ListModelOperatorServicesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *ListModelOperatorServicesResponseBody
	GetRequestId() *string
	SetServices(v []*ListModelOperatorServicesResponseBodyServices) *ListModelOperatorServicesResponseBody
	GetServices() []*ListModelOperatorServicesResponseBodyServices
	SetTotalRecordCount(v int32) *ListModelOperatorServicesResponseBody
	GetTotalRecordCount() *int32
}

type ListModelOperatorServicesResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Services  []*ListModelOperatorServicesResponseBodyServices `json:"Services,omitempty" xml:"Services,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s ListModelOperatorServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelOperatorServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelOperatorServicesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListModelOperatorServicesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *ListModelOperatorServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelOperatorServicesResponseBody) GetServices() []*ListModelOperatorServicesResponseBodyServices {
	return s.Services
}

func (s *ListModelOperatorServicesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *ListModelOperatorServicesResponseBody) SetPageNumber(v int32) *ListModelOperatorServicesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListModelOperatorServicesResponseBody) SetPageRecordCount(v int32) *ListModelOperatorServicesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *ListModelOperatorServicesResponseBody) SetRequestId(v string) *ListModelOperatorServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelOperatorServicesResponseBody) SetServices(v []*ListModelOperatorServicesResponseBodyServices) *ListModelOperatorServicesResponseBody {
	s.Services = v
	return s
}

func (s *ListModelOperatorServicesResponseBody) SetTotalRecordCount(v int32) *ListModelOperatorServicesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *ListModelOperatorServicesResponseBody) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelOperatorServicesResponseBodyServices struct {
	// example:
	//
	// agdb-xxx
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s ListModelOperatorServicesResponseBodyServices) String() string {
	return dara.Prettify(s)
}

func (s ListModelOperatorServicesResponseBodyServices) GoString() string {
	return s.String()
}

func (s *ListModelOperatorServicesResponseBodyServices) GetServiceId() *string {
	return s.ServiceId
}

func (s *ListModelOperatorServicesResponseBodyServices) SetServiceId(v string) *ListModelOperatorServicesResponseBodyServices {
	s.ServiceId = &v
	return s
}

func (s *ListModelOperatorServicesResponseBodyServices) Validate() error {
	return dara.Validate(s)
}
