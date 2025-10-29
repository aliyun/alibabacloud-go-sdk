// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHttpApisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListHttpApisResponseBody
	GetCode() *string
	SetData(v *ListHttpApisResponseBodyData) *ListHttpApisResponseBody
	GetData() *ListHttpApisResponseBodyData
	SetMessage(v string) *ListHttpApisResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHttpApisResponseBody
	GetRequestId() *string
}

type ListHttpApisResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The APIs.
	Data *ListHttpApisResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ListHttpApisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApisResponseBody) GoString() string {
	return s.String()
}

func (s *ListHttpApisResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListHttpApisResponseBody) GetData() *ListHttpApisResponseBodyData {
	return s.Data
}

func (s *ListHttpApisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHttpApisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHttpApisResponseBody) SetCode(v string) *ListHttpApisResponseBody {
	s.Code = &v
	return s
}

func (s *ListHttpApisResponseBody) SetData(v *ListHttpApisResponseBodyData) *ListHttpApisResponseBody {
	s.Data = v
	return s
}

func (s *ListHttpApisResponseBody) SetMessage(v string) *ListHttpApisResponseBody {
	s.Message = &v
	return s
}

func (s *ListHttpApisResponseBody) SetRequestId(v string) *ListHttpApisResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHttpApisResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListHttpApisResponseBodyData struct {
	// The API information.
	Items []*HttpApiInfoByName `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListHttpApisResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListHttpApisResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListHttpApisResponseBodyData) GetItems() []*HttpApiInfoByName {
	return s.Items
}

func (s *ListHttpApisResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHttpApisResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHttpApisResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListHttpApisResponseBodyData) SetItems(v []*HttpApiInfoByName) *ListHttpApisResponseBodyData {
	s.Items = v
	return s
}

func (s *ListHttpApisResponseBodyData) SetPageNumber(v int32) *ListHttpApisResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListHttpApisResponseBodyData) SetPageSize(v int32) *ListHttpApisResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListHttpApisResponseBodyData) SetTotalSize(v int32) *ListHttpApisResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListHttpApisResponseBodyData) Validate() error {
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
