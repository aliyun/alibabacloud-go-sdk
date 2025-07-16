// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFormRemarksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFormRemarkVoMap(v map[string]interface{}) *ListFormRemarksResponseBody
	GetFormRemarkVoMap() map[string]interface{}
	SetRequestId(v string) *ListFormRemarksResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *ListFormRemarksResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListFormRemarksResponseBody
	GetVendorType() *string
}

type ListFormRemarksResponseBody struct {
	// example:
	//
	// {}
	FormRemarkVoMap map[string]interface{} `json:"formRemarkVoMap,omitempty" xml:"formRemarkVoMap,omitempty"`
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

func (s ListFormRemarksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFormRemarksResponseBody) GoString() string {
	return s.String()
}

func (s *ListFormRemarksResponseBody) GetFormRemarkVoMap() map[string]interface{} {
	return s.FormRemarkVoMap
}

func (s *ListFormRemarksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFormRemarksResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListFormRemarksResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListFormRemarksResponseBody) SetFormRemarkVoMap(v map[string]interface{}) *ListFormRemarksResponseBody {
	s.FormRemarkVoMap = v
	return s
}

func (s *ListFormRemarksResponseBody) SetRequestId(v string) *ListFormRemarksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFormRemarksResponseBody) SetVendorRequestId(v string) *ListFormRemarksResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListFormRemarksResponseBody) SetVendorType(v string) *ListFormRemarksResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListFormRemarksResponseBody) Validate() error {
	return dara.Validate(s)
}
