// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvgExecuteCostReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAvgExecuteCostReportResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeAvgExecuteCostReportResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeAvgExecuteCostReportResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAvgExecuteCostReportResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeAvgExecuteCostReportResponseBodyResultObject) *DescribeAvgExecuteCostReportResponseBody
	GetResultObject() []*DescribeAvgExecuteCostReportResponseBodyResultObject
	SetSuccess(v bool) *DescribeAvgExecuteCostReportResponseBody
	GetSuccess() *bool
}

type DescribeAvgExecuteCostReportResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
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
	ResultObject []*DescribeAvgExecuteCostReportResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Whether the call was successful
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeAvgExecuteCostReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvgExecuteCostReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvgExecuteCostReportResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAvgExecuteCostReportResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeAvgExecuteCostReportResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAvgExecuteCostReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvgExecuteCostReportResponseBody) GetResultObject() []*DescribeAvgExecuteCostReportResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeAvgExecuteCostReportResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAvgExecuteCostReportResponseBody) SetCode(v string) *DescribeAvgExecuteCostReportResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAvgExecuteCostReportResponseBody) SetHttpStatusCode(v string) *DescribeAvgExecuteCostReportResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAvgExecuteCostReportResponseBody) SetMessage(v string) *DescribeAvgExecuteCostReportResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAvgExecuteCostReportResponseBody) SetRequestId(v string) *DescribeAvgExecuteCostReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvgExecuteCostReportResponseBody) SetResultObject(v []*DescribeAvgExecuteCostReportResponseBodyResultObject) *DescribeAvgExecuteCostReportResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeAvgExecuteCostReportResponseBody) SetSuccess(v bool) *DescribeAvgExecuteCostReportResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAvgExecuteCostReportResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeAvgExecuteCostReportResponseBodyResultObject struct {
	// Comparison with yesterday\\"s average execution time
	//
	// example:
	//
	// 0.2
	Ratio *string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// Today\\"s average execution time
	//
	// example:
	//
	// 0.1 毫秒
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeAvgExecuteCostReportResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvgExecuteCostReportResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeAvgExecuteCostReportResponseBodyResultObject) GetRatio() *string {
	return s.Ratio
}

func (s *DescribeAvgExecuteCostReportResponseBodyResultObject) GetValue() *string {
	return s.Value
}

func (s *DescribeAvgExecuteCostReportResponseBodyResultObject) SetRatio(v string) *DescribeAvgExecuteCostReportResponseBodyResultObject {
	s.Ratio = &v
	return s
}

func (s *DescribeAvgExecuteCostReportResponseBodyResultObject) SetValue(v string) *DescribeAvgExecuteCostReportResponseBodyResultObject {
	s.Value = &v
	return s
}

func (s *DescribeAvgExecuteCostReportResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
