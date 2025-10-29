// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApiRoutesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHttpApiRoutesResponseBody
	GetCode() *string
	SetData(v *ListHttpApiRoutesResponseBodyData) *ListHttpApiRoutesResponseBody
	GetData() *ListHttpApiRoutesResponseBodyData
	SetMessage(v string) *ListHttpApiRoutesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHttpApiRoutesResponseBody
	GetRequestId() *string
}

type ListHttpApiRoutesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response parameters.
	Data *ListHttpApiRoutesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// CBEEB8C1-108E-50F0-9BEA-DED79553C309
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListHttpApiRoutesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApiRoutesResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpApiRoutesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHttpApiRoutesResponseBody) GetData() *ListHttpApiRoutesResponseBodyData {
	return s.Data
}

func (s *ListHttpApiRoutesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHttpApiRoutesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpApiRoutesResponseBody) SetCode(v string) *ListHttpApiRoutesResponseBody {
	s.Code = &v
	return s
}

func (s *ListHttpApiRoutesResponseBody) SetData(v *ListHttpApiRoutesResponseBodyData) *ListHttpApiRoutesResponseBody {
	s.Data = v
	return s
}

func (s *ListHttpApiRoutesResponseBody) SetMessage(v string) *ListHttpApiRoutesResponseBody {
	s.Message = &v
	return s
}

func (s *ListHttpApiRoutesResponseBody) SetRequestId(v string) *ListHttpApiRoutesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpApiRoutesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHttpApiRoutesResponseBodyData struct {
	// The routes.
	Items []*HttpRoute `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 9
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListHttpApiRoutesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApiRoutesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHttpApiRoutesResponseBodyData) GetItems() []*HttpRoute {
	return s.Items
}

func (s *ListHttpApiRoutesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpApiRoutesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpApiRoutesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListHttpApiRoutesResponseBodyData) SetItems(v []*HttpRoute) *ListHttpApiRoutesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListHttpApiRoutesResponseBodyData) SetPageNumber(v int32) *ListHttpApiRoutesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApiRoutesResponseBodyData) SetPageSize(v int32) *ListHttpApiRoutesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHttpApiRoutesResponseBodyData) SetTotalSize(v int32) *ListHttpApiRoutesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListHttpApiRoutesResponseBodyData) Validate() error {
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
