// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveFormDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *BatchSaveFormDataResponseBody
	GetRequestId() *string
	SetResult(v []*string) *BatchSaveFormDataResponseBody
	GetResult() []*string
	SetVendorRequestId(v string) *BatchSaveFormDataResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *BatchSaveFormDataResponseBody
	GetVendorType() *string
}

type BatchSaveFormDataResponseBody struct {
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

func (s BatchSaveFormDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *BatchSaveFormDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchSaveFormDataResponseBody) GetResult() []*string {
	return s.Result
}

func (s *BatchSaveFormDataResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *BatchSaveFormDataResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *BatchSaveFormDataResponseBody) SetRequestId(v string) *BatchSaveFormDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchSaveFormDataResponseBody) SetResult(v []*string) *BatchSaveFormDataResponseBody {
	s.Result = v
	return s
}

func (s *BatchSaveFormDataResponseBody) SetVendorRequestId(v string) *BatchSaveFormDataResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *BatchSaveFormDataResponseBody) SetVendorType(v string) *BatchSaveFormDataResponseBody {
	s.VendorType = &v
	return s
}

func (s *BatchSaveFormDataResponseBody) Validate() error {
	return dara.Validate(s)
}
