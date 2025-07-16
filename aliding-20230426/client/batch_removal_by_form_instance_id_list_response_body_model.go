// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchRemovalByFormInstanceIdListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchRemovalByFormInstanceIdListResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *BatchRemovalByFormInstanceIdListResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *BatchRemovalByFormInstanceIdListResponseBody
	GetVendorType() *string
}

type BatchRemovalByFormInstanceIdListResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s BatchRemovalByFormInstanceIdListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchRemovalByFormInstanceIdListResponseBody) GoString() string {
	return s.String()
}

func (s *BatchRemovalByFormInstanceIdListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchRemovalByFormInstanceIdListResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *BatchRemovalByFormInstanceIdListResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *BatchRemovalByFormInstanceIdListResponseBody) SetRequestId(v string) *BatchRemovalByFormInstanceIdListResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListResponseBody) SetVendorRequestId(v string) *BatchRemovalByFormInstanceIdListResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListResponseBody) SetVendorType(v string) *BatchRemovalByFormInstanceIdListResponseBody {
	s.VendorType = &v
	return s
}

func (s *BatchRemovalByFormInstanceIdListResponseBody) Validate() error {
	return dara.Validate(s)
}
