// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkMeetingMetricDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *GetDingtalkMeetingMetricDataResponseBody
	GetData() interface{}
	SetRequestId(v string) *GetDingtalkMeetingMetricDataResponseBody
	GetRequestId() *string
	SetVendorRequestId(v string) *GetDingtalkMeetingMetricDataResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetDingtalkMeetingMetricDataResponseBody
	GetVendorType() *string
}

type GetDingtalkMeetingMetricDataResponseBody struct {
	// example:
	//
	// {}
	Data interface{} `json:"data,omitempty" xml:"data,omitempty"`
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

func (s GetDingtalkMeetingMetricDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkMeetingMetricDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetDingtalkMeetingMetricDataResponseBody) GetData() interface{} {
	return s.Data
}

func (s *GetDingtalkMeetingMetricDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDingtalkMeetingMetricDataResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetDingtalkMeetingMetricDataResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetDingtalkMeetingMetricDataResponseBody) SetData(v interface{}) *GetDingtalkMeetingMetricDataResponseBody {
	s.Data = v
	return s
}

func (s *GetDingtalkMeetingMetricDataResponseBody) SetRequestId(v string) *GetDingtalkMeetingMetricDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataResponseBody) SetVendorRequestId(v string) *GetDingtalkMeetingMetricDataResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataResponseBody) SetVendorType(v string) *GetDingtalkMeetingMetricDataResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetDingtalkMeetingMetricDataResponseBody) Validate() error {
	return dara.Validate(s)
}
