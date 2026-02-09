// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApiOperationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHttpApiOperationsResponseBody
	GetCode() *string
	SetData(v *ListHttpApiOperationsResponseBodyData) *ListHttpApiOperationsResponseBody
	GetData() *ListHttpApiOperationsResponseBodyData
	SetMessage(v string) *ListHttpApiOperationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHttpApiOperationsResponseBody
	GetRequestId() *string
}

type ListHttpApiOperationsResponseBody struct {
	// Response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The operations.
	Data *ListHttpApiOperationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 585657D2-1C20-5B8A-AF17-D727C6490BE4
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListHttpApiOperationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApiOperationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHttpApiOperationsResponseBody) GetData() *ListHttpApiOperationsResponseBodyData {
	return s.Data
}

func (s *ListHttpApiOperationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHttpApiOperationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpApiOperationsResponseBody) SetCode(v string) *ListHttpApiOperationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListHttpApiOperationsResponseBody) SetData(v *ListHttpApiOperationsResponseBodyData) *ListHttpApiOperationsResponseBody {
	s.Data = v
	return s
}

func (s *ListHttpApiOperationsResponseBody) SetMessage(v string) *ListHttpApiOperationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListHttpApiOperationsResponseBody) SetRequestId(v string) *ListHttpApiOperationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpApiOperationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHttpApiOperationsResponseBodyData struct {
	// The operations.
	Items []*HttpApiOperationInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total count.
	//
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListHttpApiOperationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApiOperationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHttpApiOperationsResponseBodyData) GetItems() []*HttpApiOperationInfo {
	return s.Items
}

func (s *ListHttpApiOperationsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpApiOperationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpApiOperationsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListHttpApiOperationsResponseBodyData) SetItems(v []*HttpApiOperationInfo) *ListHttpApiOperationsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListHttpApiOperationsResponseBodyData) SetPageNumber(v int32) *ListHttpApiOperationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApiOperationsResponseBodyData) SetPageSize(v int32) *ListHttpApiOperationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHttpApiOperationsResponseBodyData) SetTotalSize(v int32) *ListHttpApiOperationsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListHttpApiOperationsResponseBodyData) Validate() error {
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
