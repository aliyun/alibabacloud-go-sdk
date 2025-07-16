// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIdListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *GetInstanceIdListResponseBody
	GetData() []*string
	SetPageNumber(v int64) *GetInstanceIdListResponseBody
	GetPageNumber() *int64
	SetRequestId(v string) *GetInstanceIdListResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *GetInstanceIdListResponseBody
	GetTotalCount() *int64
	SetVendorRequestId(v string) *GetInstanceIdListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetInstanceIdListResponseBody
	GetVendorType() *string
}

type GetInstanceIdListResponseBody struct {
	// example:
	//
	// [ "FINST-BOOxxx" ]
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
	// 10
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

func (s GetInstanceIdListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIdListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIdListResponseBody) GetData() []*string {
	return s.Data
}

func (s *GetInstanceIdListResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetInstanceIdListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceIdListResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *GetInstanceIdListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetInstanceIdListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetInstanceIdListResponseBody) SetData(v []*string) *GetInstanceIdListResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceIdListResponseBody) SetPageNumber(v int64) *GetInstanceIdListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetInstanceIdListResponseBody) SetRequestId(v string) *GetInstanceIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceIdListResponseBody) SetTotalCount(v int64) *GetInstanceIdListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceIdListResponseBody) SetVendorRequestId(v string) *GetInstanceIdListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetInstanceIdListResponseBody) SetVendorType(v string) *GetInstanceIdListResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetInstanceIdListResponseBody) Validate() error {
	return dara.Validate(s)
}
