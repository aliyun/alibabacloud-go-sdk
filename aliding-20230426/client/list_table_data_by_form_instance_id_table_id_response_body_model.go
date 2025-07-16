// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTableDataByFormInstanceIdTableIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *ListTableDataByFormInstanceIdTableIdResponseBody
	GetData() []*string
	SetPageNumber(v int64) *ListTableDataByFormInstanceIdTableIdResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *ListTableDataByFormInstanceIdTableIdResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListTableDataByFormInstanceIdTableIdResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *ListTableDataByFormInstanceIdTableIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListTableDataByFormInstanceIdTableIdResponseBody
	GetVendorType() *string
}

type ListTableDataByFormInstanceIdTableIdResponseBody struct {
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

func (s ListTableDataByFormInstanceIdTableIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTableDataByFormInstanceIdTableIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) GetData() []*string {
	return s.Data
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetData(v []*string) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.Data = v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetPageNumber(v int64) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetRequestId(v string) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetTotalCount(v int64) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetVendorRequestId(v string) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) SetVendorType(v string) *ListTableDataByFormInstanceIdTableIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListTableDataByFormInstanceIdTableIdResponseBody) Validate() error {
	return dara.Validate(s)
}
