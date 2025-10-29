// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEnvironmentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListEnvironmentsResponseBody
	GetCode() *string
	SetData(v *ListEnvironmentsResponseBodyData) *ListEnvironmentsResponseBody
	GetData() *ListEnvironmentsResponseBodyData
	SetMessage(v string) *ListEnvironmentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListEnvironmentsResponseBody
	GetRequestId() *string
}

type ListEnvironmentsResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Paged query environment list response.
	Data *ListEnvironmentsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the call chain.
	//
	// example:
	//
	// CE857A85-251D-5018-8103-A38957D71E20
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListEnvironmentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListEnvironmentsResponseBody) GetData() *ListEnvironmentsResponseBodyData {
	return s.Data
}

func (s *ListEnvironmentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListEnvironmentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEnvironmentsResponseBody) SetCode(v string) *ListEnvironmentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetData(v *ListEnvironmentsResponseBodyData) *ListEnvironmentsResponseBody {
	s.Data = v
	return s
}

func (s *ListEnvironmentsResponseBody) SetMessage(v string) *ListEnvironmentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListEnvironmentsResponseBody) SetRequestId(v string) *ListEnvironmentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEnvironmentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEnvironmentsResponseBodyData struct {
	// List of environment information.
	Items []*EnvironmentInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Number of items per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 25
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListEnvironmentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListEnvironmentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListEnvironmentsResponseBodyData) GetItems() []*EnvironmentInfo {
	return s.Items
}

func (s *ListEnvironmentsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListEnvironmentsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListEnvironmentsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListEnvironmentsResponseBodyData) SetItems(v []*EnvironmentInfo) *ListEnvironmentsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetPageNumber(v int32) *ListEnvironmentsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetPageSize(v int32) *ListEnvironmentsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) SetTotalSize(v int32) *ListEnvironmentsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListEnvironmentsResponseBodyData) Validate() error {
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
