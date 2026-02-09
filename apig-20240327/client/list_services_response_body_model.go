// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListServicesResponseBody
	GetCode() *string
	SetData(v *ListServicesResponseBodyData) *ListServicesResponseBody
	GetData() *ListServicesResponseBodyData
	SetMessage(v string) *ListServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListServicesResponseBody
	GetRequestId() *string
}

type ListServicesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response payload.
	Data *ListServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListServicesResponseBody) GetData() *ListServicesResponseBodyData {
	return s.Data
}

func (s *ListServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServicesResponseBody) SetCode(v string) *ListServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListServicesResponseBody) SetData(v *ListServicesResponseBodyData) *ListServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListServicesResponseBody) SetMessage(v string) *ListServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListServicesResponseBody) SetRequestId(v string) *ListServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServicesResponseBodyData struct {
	// The list of services.
	Items []*Service `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 18
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListServicesResponseBodyData) GetItems() []*Service {
	return s.Items
}

func (s *ListServicesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListServicesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListServicesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListServicesResponseBodyData) SetItems(v []*Service) *ListServicesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListServicesResponseBodyData) SetPageNumber(v int32) *ListServicesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListServicesResponseBodyData) SetPageSize(v int32) *ListServicesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListServicesResponseBodyData) SetTotalSize(v int32) *ListServicesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListServicesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
