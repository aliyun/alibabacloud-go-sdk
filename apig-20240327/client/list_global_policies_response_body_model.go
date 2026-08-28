// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGlobalPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGlobalPoliciesResponseBody
	GetCode() *string
	SetData(v *ListGlobalPoliciesResponseBodyData) *ListGlobalPoliciesResponseBody
	GetData() *ListGlobalPoliciesResponseBodyData
	SetMessage(v string) *ListGlobalPoliciesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGlobalPoliciesResponseBody
	GetRequestId() *string
}

type ListGlobalPoliciesResponseBody struct {
	// example:
	//
	// 200
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListGlobalPoliciesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 350E9393-B90C-5540-B2BE-6F4CF5965CDA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGlobalPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGlobalPoliciesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGlobalPoliciesResponseBody) GetData() *ListGlobalPoliciesResponseBodyData {
	return s.Data
}

func (s *ListGlobalPoliciesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGlobalPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGlobalPoliciesResponseBody) SetCode(v string) *ListGlobalPoliciesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGlobalPoliciesResponseBody) SetData(v *ListGlobalPoliciesResponseBodyData) *ListGlobalPoliciesResponseBody {
	s.Data = v
	return s
}

func (s *ListGlobalPoliciesResponseBody) SetMessage(v string) *ListGlobalPoliciesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGlobalPoliciesResponseBody) SetRequestId(v string) *ListGlobalPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGlobalPoliciesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGlobalPoliciesResponseBodyData struct {
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
	// 25
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListGlobalPoliciesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGlobalPoliciesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGlobalPoliciesResponseBodyData) GetItems() []*PolicyInfo {
	return s.Items
}

func (s *ListGlobalPoliciesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGlobalPoliciesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGlobalPoliciesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListGlobalPoliciesResponseBodyData) SetItems(v []*PolicyInfo) *ListGlobalPoliciesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGlobalPoliciesResponseBodyData) SetPageNumber(v int32) *ListGlobalPoliciesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGlobalPoliciesResponseBodyData) SetPageSize(v int32) *ListGlobalPoliciesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGlobalPoliciesResponseBodyData) SetTotalSize(v int32) *ListGlobalPoliciesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGlobalPoliciesResponseBodyData) Validate() error {
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
