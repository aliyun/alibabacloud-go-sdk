// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListMemoryEventResponseBody
	GetCode() *string
	SetData(v *ListMemoryEventResponseBodyData) *ListMemoryEventResponseBody
	GetData() *ListMemoryEventResponseBodyData
	SetRequestId(v string) *ListMemoryEventResponseBody
	GetRequestId() *string
}

type ListMemoryEventResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                          `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListMemoryEventResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 0FB1162C-D50B-5DA7-AD04-3417ABBF133A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListMemoryEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryEventResponseBody) GoString() string {
	return s.String()
}

func (s *ListMemoryEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListMemoryEventResponseBody) GetData() *ListMemoryEventResponseBodyData {
	return s.Data
}

func (s *ListMemoryEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMemoryEventResponseBody) SetCode(v string) *ListMemoryEventResponseBody {
	s.Code = &v
	return s
}

func (s *ListMemoryEventResponseBody) SetData(v *ListMemoryEventResponseBodyData) *ListMemoryEventResponseBody {
	s.Data = v
	return s
}

func (s *ListMemoryEventResponseBody) SetRequestId(v string) *ListMemoryEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMemoryEventResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMemoryEventResponseBodyData struct {
	Items []*string `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 2
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListMemoryEventResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryEventResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListMemoryEventResponseBodyData) GetItems() []*string {
	return s.Items
}

func (s *ListMemoryEventResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListMemoryEventResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListMemoryEventResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListMemoryEventResponseBodyData) SetItems(v []*string) *ListMemoryEventResponseBodyData {
	s.Items = v
	return s
}

func (s *ListMemoryEventResponseBodyData) SetPageNumber(v int32) *ListMemoryEventResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListMemoryEventResponseBodyData) SetPageSize(v int32) *ListMemoryEventResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListMemoryEventResponseBodyData) SetTotal(v int64) *ListMemoryEventResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListMemoryEventResponseBodyData) Validate() error {
	return dara.Validate(s)
}
