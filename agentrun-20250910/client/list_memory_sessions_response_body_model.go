// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemorySessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemorySessionsResponseBody
	GetCode() *string
	SetData(v *ListMemorySessionsResponseBodyData) *ListMemorySessionsResponseBody
	GetData() *ListMemorySessionsResponseBodyData
	SetRequestId(v string) *ListMemorySessionsResponseBody
	GetRequestId() *string
}

type ListMemorySessionsResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListMemorySessionsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// E2C43519-6095-5487-9526-051AB8F50B4A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMemorySessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemorySessionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemorySessionsResponseBody) GetData() *ListMemorySessionsResponseBodyData {
	return s.Data
}

func (s *ListMemorySessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemorySessionsResponseBody) SetCode(v string) *ListMemorySessionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemorySessionsResponseBody) SetData(v *ListMemorySessionsResponseBodyData) *ListMemorySessionsResponseBody {
	s.Data = v
	return s
}

func (s *ListMemorySessionsResponseBody) SetRequestId(v string) *ListMemorySessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemorySessionsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMemorySessionsResponseBodyData struct {
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
	// 211
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMemorySessionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMemorySessionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMemorySessionsResponseBodyData) GetItems() []*string {
	return s.Items
}

func (s *ListMemorySessionsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemorySessionsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemorySessionsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListMemorySessionsResponseBodyData) SetItems(v []*string) *ListMemorySessionsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListMemorySessionsResponseBodyData) SetPageNumber(v int32) *ListMemorySessionsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMemorySessionsResponseBodyData) SetPageSize(v int32) *ListMemorySessionsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMemorySessionsResponseBodyData) SetTotal(v int64) *ListMemorySessionsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMemorySessionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
