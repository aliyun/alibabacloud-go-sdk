// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubtaskItemsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListSubtaskItemsResponseBody
	GetCode() *int32
	SetDetails(v string) *ListSubtaskItemsResponseBody
	GetDetails() *string
	SetErrorCode(v string) *ListSubtaskItemsResponseBody
	GetErrorCode() *string
	SetItems(v []*SubtaskItemDetail) *ListSubtaskItemsResponseBody
	GetItems() []*SubtaskItemDetail
	SetMessage(v string) *ListSubtaskItemsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListSubtaskItemsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSubtaskItemsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSubtaskItemsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListSubtaskItemsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListSubtaskItemsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListSubtaskItemsResponseBody
	GetTotalPage() *int32
}

type ListSubtaskItemsResponseBody struct {
	// Return encoding. The default value is 0, indicating normal execution.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Details.
	//
	// example:
	//
	// ""
	Details *string `json:"Details,omitempty" xml:"Details,omitempty"`
	// Error code.
	//
	// example:
	//
	// ""
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// List of data items.
	Items []*SubtaskItemDetail `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// Return message of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Page number of the queried task package annotation data returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Quantity of annotated task package data entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90ABA848-AD74-1F6E-84BC-4182A7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request succeeded. Possible values:
	//
	// - true: The request succeeded.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total quantity of annotated data in the task package.
	//
	// example:
	//
	// 22
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSubtaskItemsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSubtaskItemsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSubtaskItemsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListSubtaskItemsResponseBody) GetDetails() *string {
	return s.Details
}

func (s *ListSubtaskItemsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListSubtaskItemsResponseBody) GetItems() []*SubtaskItemDetail {
	return s.Items
}

func (s *ListSubtaskItemsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListSubtaskItemsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSubtaskItemsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSubtaskItemsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSubtaskItemsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListSubtaskItemsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSubtaskItemsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListSubtaskItemsResponseBody) SetCode(v int32) *ListSubtaskItemsResponseBody {
	s.Code = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetDetails(v string) *ListSubtaskItemsResponseBody {
	s.Details = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetErrorCode(v string) *ListSubtaskItemsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetItems(v []*SubtaskItemDetail) *ListSubtaskItemsResponseBody {
	s.Items = v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetMessage(v string) *ListSubtaskItemsResponseBody {
	s.Message = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetPageNumber(v int32) *ListSubtaskItemsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetPageSize(v int32) *ListSubtaskItemsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetRequestId(v string) *ListSubtaskItemsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetSuccess(v bool) *ListSubtaskItemsResponseBody {
	s.Success = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetTotalCount(v int32) *ListSubtaskItemsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) SetTotalPage(v int32) *ListSubtaskItemsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSubtaskItemsResponseBody) Validate() error {
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
