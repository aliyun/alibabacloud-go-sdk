// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUpdateFormDataByInstanceMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchUpdateFormDataByInstanceMapResponseBody
	GetRequestId() *string
	SetResult(v []*string) *BatchUpdateFormDataByInstanceMapResponseBody
	GetResult() []*string
	SetVendorRequestId(v string) *BatchUpdateFormDataByInstanceMapResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *BatchUpdateFormDataByInstanceMapResponseBody
	GetVendorType() *string
}

type BatchUpdateFormDataByInstanceMapResponseBody struct {
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

func (s BatchUpdateFormDataByInstanceMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUpdateFormDataByInstanceMapResponseBody) GoString() string {
	return s.String()
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) GetResult() []*string {
	return s.Result
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) SetRequestId(v string) *BatchUpdateFormDataByInstanceMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) SetResult(v []*string) *BatchUpdateFormDataByInstanceMapResponseBody {
	s.Result = v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) SetVendorRequestId(v string) *BatchUpdateFormDataByInstanceMapResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) SetVendorType(v string) *BatchUpdateFormDataByInstanceMapResponseBody {
	s.VendorType = &v
	return s
}

func (s *BatchUpdateFormDataByInstanceMapResponseBody) Validate() error {
	return dara.Validate(s)
}
