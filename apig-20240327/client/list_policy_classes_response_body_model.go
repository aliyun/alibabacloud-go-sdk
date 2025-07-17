// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolicyClassesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPolicyClassesResponseBody
	GetCode() *string
	SetData(v *ListPolicyClassesResponseBodyData) *ListPolicyClassesResponseBody
	GetData() *ListPolicyClassesResponseBodyData
	SetMessage(v string) *ListPolicyClassesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPolicyClassesResponseBody
	GetRequestId() *string
}

type ListPolicyClassesResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Policy template information.
	Data *ListPolicyClassesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// ResponseMessage
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 23B45FA9-7208-5E55-B5CE-B6B2567DD822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPolicyClassesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyClassesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPolicyClassesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPolicyClassesResponseBody) GetData() *ListPolicyClassesResponseBodyData {
	return s.Data
}

func (s *ListPolicyClassesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPolicyClassesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPolicyClassesResponseBody) SetCode(v string) *ListPolicyClassesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPolicyClassesResponseBody) SetData(v *ListPolicyClassesResponseBodyData) *ListPolicyClassesResponseBody {
	s.Data = v
	return s
}

func (s *ListPolicyClassesResponseBody) SetMessage(v string) *ListPolicyClassesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPolicyClassesResponseBody) SetRequestId(v string) *ListPolicyClassesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPolicyClassesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListPolicyClassesResponseBodyData struct {
	// List of policy templates
	Items []*PolicyClassInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// Page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// Page size
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Total number of items.
	//
	// example:
	//
	// 10
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListPolicyClassesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPolicyClassesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPolicyClassesResponseBodyData) GetItems() []*PolicyClassInfo {
	return s.Items
}

func (s *ListPolicyClassesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListPolicyClassesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListPolicyClassesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListPolicyClassesResponseBodyData) SetItems(v []*PolicyClassInfo) *ListPolicyClassesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListPolicyClassesResponseBodyData) SetPageNumber(v int32) *ListPolicyClassesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListPolicyClassesResponseBodyData) SetPageSize(v int32) *ListPolicyClassesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListPolicyClassesResponseBodyData) SetTotalSize(v int32) *ListPolicyClassesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListPolicyClassesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
