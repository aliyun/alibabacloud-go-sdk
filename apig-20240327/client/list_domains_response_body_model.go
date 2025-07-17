// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDomainsResponseBody
	GetCode() *string
	SetData(v *ListDomainsResponseBodyData) *ListDomainsResponseBody
	GetData() *ListDomainsResponseBodyData
	SetMessage(v string) *ListDomainsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDomainsResponseBody
	GetRequestId() *string
}

type ListDomainsResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *ListDomainsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to trace the API call link.
	//
	// example:
	//
	// C61E30D3-579A-5B43-994E-31E02EDC9129
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDomainsResponseBody) GetData() *ListDomainsResponseBodyData {
	return s.Data
}

func (s *ListDomainsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDomainsResponseBody) SetCode(v string) *ListDomainsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDomainsResponseBody) SetData(v *ListDomainsResponseBodyData) *ListDomainsResponseBody {
	s.Data = v
	return s
}

func (s *ListDomainsResponseBody) SetMessage(v string) *ListDomainsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDomainsResponseBody) SetRequestId(v string) *ListDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDomainsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDomainsResponseBodyData struct {
	// The information about the domain names.
	Items []*DomainInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The page number of the returned page.
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
	// The total number of entries returned.
	//
	// example:
	//
	// 9
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListDomainsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBodyData) GetItems() []*DomainInfo {
	return s.Items
}

func (s *ListDomainsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListDomainsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListDomainsResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListDomainsResponseBodyData) SetItems(v []*DomainInfo) *ListDomainsResponseBodyData {
	s.Items = v
	return s
}

func (s *ListDomainsResponseBodyData) SetPageNumber(v int32) *ListDomainsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListDomainsResponseBodyData) SetPageSize(v int32) *ListDomainsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListDomainsResponseBodyData) SetTotalSize(v int32) *ListDomainsResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListDomainsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
