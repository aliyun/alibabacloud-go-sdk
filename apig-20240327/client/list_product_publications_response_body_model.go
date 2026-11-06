// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductPublicationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListProductPublicationsResponseBody
	GetCode() *string
	SetData(v *ListProductPublicationsResponseBodyData) *ListProductPublicationsResponseBody
	GetData() *ListProductPublicationsResponseBodyData
	SetMessage(v string) *ListProductPublicationsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListProductPublicationsResponseBody
	GetRequestId() *string
}

type ListProductPublicationsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListProductPublicationsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListProductPublicationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProductPublicationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListProductPublicationsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListProductPublicationsResponseBody) GetData() *ListProductPublicationsResponseBodyData {
	return s.Data
}

func (s *ListProductPublicationsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListProductPublicationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProductPublicationsResponseBody) SetCode(v string) *ListProductPublicationsResponseBody {
	s.Code = &v
	return s
}

func (s *ListProductPublicationsResponseBody) SetData(v *ListProductPublicationsResponseBodyData) *ListProductPublicationsResponseBody {
	s.Data = v
	return s
}

func (s *ListProductPublicationsResponseBody) SetMessage(v string) *ListProductPublicationsResponseBody {
	s.Message = &v
	return s
}

func (s *ListProductPublicationsResponseBody) SetRequestId(v string) *ListProductPublicationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProductPublicationsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListProductPublicationsResponseBodyData struct {
	// The list of publication records for the API product. The list element is named productPublication.
	Content []*HiMarketProductPublicationDetail `json:"content,omitempty" xml:"content,omitempty" type:"Repeated"`
	// The page number.
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
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListProductPublicationsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProductPublicationsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProductPublicationsResponseBodyData) GetContent() []*HiMarketProductPublicationDetail {
	return s.Content
}

func (s *ListProductPublicationsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListProductPublicationsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProductPublicationsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListProductPublicationsResponseBodyData) SetContent(v []*HiMarketProductPublicationDetail) *ListProductPublicationsResponseBodyData {
	s.Content = v
	return s
}

func (s *ListProductPublicationsResponseBodyData) SetPageNumber(v int32) *ListProductPublicationsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListProductPublicationsResponseBodyData) SetPageSize(v int32) *ListProductPublicationsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListProductPublicationsResponseBodyData) SetTotalCount(v int32) *ListProductPublicationsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListProductPublicationsResponseBodyData) Validate() error {
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
