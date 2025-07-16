// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUpdateFormDataByInstanceIdResponseBody
	GetRequestId() *string
	SetResult(v []*string) *BatchUpdateFormDataByInstanceIdResponseBody
	GetResult() []*string
	SetVendorRequestId(v string) *BatchUpdateFormDataByInstanceIdResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *BatchUpdateFormDataByInstanceIdResponseBody
	GetVendorType() *string
}

type BatchUpdateFormDataByInstanceIdResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// [ "FINST-SASNOO39NSIFF780" ]
	Result []*string `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s BatchUpdateFormDataByInstanceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) GetResult() []*string {
	return s.Result
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) SetRequestId(v string) *BatchUpdateFormDataByInstanceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) SetResult(v []*string) *BatchUpdateFormDataByInstanceIdResponseBody {
	s.Result = v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) SetVendorRequestId(v string) *BatchUpdateFormDataByInstanceIdResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) SetVendorType(v string) *BatchUpdateFormDataByInstanceIdResponseBody {
	s.VendorType = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceIdResponseBody) Validate() error {
	return dara.Validate(s)
}
