// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemoryResponseBody
	GetCode() *string
	SetData(v *ListMemoryResponseBodyData) *ListMemoryResponseBody
	GetData() *ListMemoryResponseBodyData
	SetRequestId(v string) *ListMemoryResponseBody
	GetRequestId() *string
}

type ListMemoryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                     `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListMemoryResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 55D4BE40-2811-5CFB-8482-E0E98D575B1E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMemoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoryResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemoryResponseBody) GetData() *ListMemoryResponseBodyData {
	return s.Data
}

func (s *ListMemoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoryResponseBody) SetCode(v string) *ListMemoryResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemoryResponseBody) SetData(v *ListMemoryResponseBodyData) *ListMemoryResponseBody {
	s.Data = v
	return s
}

func (s *ListMemoryResponseBody) SetRequestId(v string) *ListMemoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMemoryResponseBodyData struct {
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 284
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMemoryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMemoryResponseBodyData) GetItems() []*string {
	return s.Items
}

func (s *ListMemoryResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemoryResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemoryResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListMemoryResponseBodyData) SetItems(v []*string) *ListMemoryResponseBodyData {
	s.Items = v
	return s
}

func (s *ListMemoryResponseBodyData) SetPageNumber(v int32) *ListMemoryResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMemoryResponseBodyData) SetPageSize(v int32) *ListMemoryResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMemoryResponseBodyData) SetTotal(v int64) *ListMemoryResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMemoryResponseBodyData) Validate() error {
	return dara.Validate(s)
}
