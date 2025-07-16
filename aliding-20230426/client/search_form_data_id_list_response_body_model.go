// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchFormDataIdListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *SearchFormDataIdListResponseBody
	GetData() []*string
	SetPageNumber(v int64) *SearchFormDataIdListResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *SearchFormDataIdListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *SearchFormDataIdListResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *SearchFormDataIdListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *SearchFormDataIdListResponseBody
	GetVendorType() *string
}

type SearchFormDataIdListResponseBody struct {
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s SearchFormDataIdListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchFormDataIdListResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFormDataIdListResponseBody) GetData() []*string {
	return s.Data
}

func (s *SearchFormDataIdListResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *SearchFormDataIdListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchFormDataIdListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *SearchFormDataIdListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *SearchFormDataIdListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *SearchFormDataIdListResponseBody) SetData(v []*string) *SearchFormDataIdListResponseBody {
	s.Data = v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetPageNumber(v int64) *SearchFormDataIdListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetRequestId(v string) *SearchFormDataIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetTotalCount(v int64) *SearchFormDataIdListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetVendorRequestId(v string) *SearchFormDataIdListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *SearchFormDataIdListResponseBody) SetVendorType(v string) *SearchFormDataIdListResponseBody {
	s.VendorType = &v
	return s
}

func (s *SearchFormDataIdListResponseBody) Validate() error {
	return dara.Validate(s)
}
