// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPoliciesResponseBody
	GetCode() *string
	SetData(v *ListPoliciesResponseBodyData) *ListPoliciesResponseBody
	GetData() *ListPoliciesResponseBodyData
	SetMessage(v string) *ListPoliciesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPoliciesResponseBody
	GetRequestId() *string
}

type ListPoliciesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                       `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListPoliciesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPoliciesResponseBody) GetData() *ListPoliciesResponseBodyData {
	return s.Data
}

func (s *ListPoliciesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPoliciesResponseBody) SetCode(v string) *ListPoliciesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPoliciesResponseBody) SetData(v *ListPoliciesResponseBodyData) *ListPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *ListPoliciesResponseBody) SetMessage(v string) *ListPoliciesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPoliciesResponseBody) SetRequestId(v string) *ListPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPoliciesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListPoliciesResponseBodyData struct {
	Items []*PolicyInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 18
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListPoliciesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPoliciesResponseBodyData) GetItems() []*PolicyInfo {
	return s.Items
}

func (s *ListPoliciesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPoliciesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPoliciesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListPoliciesResponseBodyData) SetItems(v []*PolicyInfo) *ListPoliciesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPoliciesResponseBodyData) SetPageNumber(v int32) *ListPoliciesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPoliciesResponseBodyData) SetPageSize(v int32) *ListPoliciesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPoliciesResponseBodyData) SetTotalSize(v int32) *ListPoliciesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListPoliciesResponseBodyData) Validate() error {
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
