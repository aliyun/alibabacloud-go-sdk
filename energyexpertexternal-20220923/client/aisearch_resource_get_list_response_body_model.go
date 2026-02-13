// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAISearchResourceGetListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AISearchResourceGetListResponseBodyData) *AISearchResourceGetListResponseBody
	GetData() *AISearchResourceGetListResponseBodyData
	SetRequestId(v string) *AISearchResourceGetListResponseBody
	GetRequestId() *string
}

type AISearchResourceGetListResponseBody struct {
	Data *AISearchResourceGetListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// A8AEC6D9-A359-5169-BD1A-BD848BA60D65
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AISearchResourceGetListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceGetListResponseBody) GoString() string {
	return s.String()
}

func (s *AISearchResourceGetListResponseBody) GetData() *AISearchResourceGetListResponseBodyData {
	return s.Data
}

func (s *AISearchResourceGetListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AISearchResourceGetListResponseBody) SetData(v *AISearchResourceGetListResponseBodyData) *AISearchResourceGetListResponseBody {
	s.Data = v
	return s
}

func (s *AISearchResourceGetListResponseBody) SetRequestId(v string) *AISearchResourceGetListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AISearchResourceGetListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AISearchResourceGetListResponseBodyData struct {
	// example:
	//
	// 1
	CurrentPage *string       `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	Items       []interface{} `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 68
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s AISearchResourceGetListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AISearchResourceGetListResponseBodyData) GoString() string {
	return s.String()
}

func (s *AISearchResourceGetListResponseBodyData) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *AISearchResourceGetListResponseBodyData) GetItems() []interface{} {
	return s.Items
}

func (s *AISearchResourceGetListResponseBodyData) GetPageSize() *string {
	return s.PageSize
}

func (s *AISearchResourceGetListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *AISearchResourceGetListResponseBodyData) SetCurrentPage(v string) *AISearchResourceGetListResponseBodyData {
	s.CurrentPage = &v
	return s
}

func (s *AISearchResourceGetListResponseBodyData) SetItems(v []interface{}) *AISearchResourceGetListResponseBodyData {
	s.Items = v
	return s
}

func (s *AISearchResourceGetListResponseBodyData) SetPageSize(v string) *AISearchResourceGetListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *AISearchResourceGetListResponseBodyData) SetTotal(v int64) *AISearchResourceGetListResponseBodyData {
	s.Total = &v
	return s
}

func (s *AISearchResourceGetListResponseBodyData) Validate() error {
	return dara.Validate(s)
}
