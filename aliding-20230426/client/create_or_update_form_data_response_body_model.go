// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateFormDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOrUpdateFormDataResponseBody
	GetRequestId() *string
	SetResult(v []*string) *CreateOrUpdateFormDataResponseBody
	GetResult() []*string
	SetVendorRequestId(v string) *CreateOrUpdateFormDataResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *CreateOrUpdateFormDataResponseBody
	GetVendorType() *string
}

type CreateOrUpdateFormDataResponseBody struct {
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

func (s CreateOrUpdateFormDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateFormDataResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateFormDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateFormDataResponseBody) GetResult() []*string {
	return s.Result
}

func (s *CreateOrUpdateFormDataResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *CreateOrUpdateFormDataResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *CreateOrUpdateFormDataResponseBody) SetRequestId(v string) *CreateOrUpdateFormDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateFormDataResponseBody) SetResult(v []*string) *CreateOrUpdateFormDataResponseBody {
	s.Result = v
	return s
}

func (s *CreateOrUpdateFormDataResponseBody) SetVendorRequestId(v string) *CreateOrUpdateFormDataResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *CreateOrUpdateFormDataResponseBody) SetVendorType(v string) *CreateOrUpdateFormDataResponseBody {
	s.VendorType = &v
	return s
}

func (s *CreateOrUpdateFormDataResponseBody) Validate() error {
	return dara.Validate(s)
}
