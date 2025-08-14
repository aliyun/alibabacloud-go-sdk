// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestPeakReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeRequestPeakReportResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeRequestPeakReportResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeRequestPeakReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeRequestPeakReportResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeRequestPeakReportResponseBodyResultObject) *DescribeRequestPeakReportResponseBody
	GetResultObject() []*DescribeRequestPeakReportResponseBodyResultObject
	SetSuccess(v bool) *DescribeRequestPeakReportResponseBody
	GetSuccess() *bool
}

type DescribeRequestPeakReportResponseBody struct {
	// Status code
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Return object
	ResultObject []*DescribeRequestPeakReportResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Whether the request was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeRequestPeakReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestPeakReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRequestPeakReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeRequestPeakReportResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeRequestPeakReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeRequestPeakReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRequestPeakReportResponseBody) GetResultObject() []*DescribeRequestPeakReportResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeRequestPeakReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRequestPeakReportResponseBody) SetCode(v string) *DescribeRequestPeakReportResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeRequestPeakReportResponseBody) SetHttpStatusCode(v string) *DescribeRequestPeakReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeRequestPeakReportResponseBody) SetMessage(v string) *DescribeRequestPeakReportResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeRequestPeakReportResponseBody) SetRequestId(v string) *DescribeRequestPeakReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRequestPeakReportResponseBody) SetResultObject(v []*DescribeRequestPeakReportResponseBodyResultObject) *DescribeRequestPeakReportResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeRequestPeakReportResponseBody) SetSuccess(v bool) *DescribeRequestPeakReportResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRequestPeakReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRequestPeakReportResponseBodyResultObject struct {
	// Return value
	//
	// example:
	//
	// 1.0
	Ratio *string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Return text
	//
	// example:
	//
	// 1.0 次/秒
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeRequestPeakReportResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestPeakReportResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeRequestPeakReportResponseBodyResultObject) GetRatio() *string {
	return s.Ratio
}

func (s *DescribeRequestPeakReportResponseBodyResultObject) GetValue() *string {
	return s.Value
}

func (s *DescribeRequestPeakReportResponseBodyResultObject) SetRatio(v string) *DescribeRequestPeakReportResponseBodyResultObject {
	s.Ratio = &v
	return s
}

func (s *DescribeRequestPeakReportResponseBodyResultObject) SetValue(v string) *DescribeRequestPeakReportResponseBodyResultObject {
	s.Value = &v
	return s
}

func (s *DescribeRequestPeakReportResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
