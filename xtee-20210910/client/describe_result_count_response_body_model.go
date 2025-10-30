// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResultCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeResultCountResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeResultCountResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeResultCountResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeResultCountResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeResultCountResponseBodyResultObject) *DescribeResultCountResponseBody
	GetResultObject() []*DescribeResultCountResponseBodyResultObject
	SetSuccess(v bool) *DescribeResultCountResponseBody
	GetSuccess() *bool
}

type DescribeResultCountResponseBody struct {
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
	ResultObject []*DescribeResultCountResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeResultCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResultCountResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeResultCountResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeResultCountResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeResultCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResultCountResponseBody) GetResultObject() []*DescribeResultCountResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeResultCountResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeResultCountResponseBody) SetCode(v string) *DescribeResultCountResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeResultCountResponseBody) SetHttpStatusCode(v string) *DescribeResultCountResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeResultCountResponseBody) SetMessage(v string) *DescribeResultCountResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeResultCountResponseBody) SetRequestId(v string) *DescribeResultCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResultCountResponseBody) SetResultObject(v []*DescribeResultCountResponseBodyResultObject) *DescribeResultCountResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeResultCountResponseBody) SetSuccess(v bool) *DescribeResultCountResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeResultCountResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeResultCountResponseBodyResultObject struct {
	// Execution result
	//
	// example:
	//
	// PASS
	Result *string `json:"result,omitempty" xml:"result,omitempty"`
	// Quantity
	//
	// example:
	//
	// 200
	Total *string `json:"total,omitempty" xml:"total,omitempty"`
}

func (s DescribeResultCountResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeResultCountResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeResultCountResponseBodyResultObject) GetResult() *string {
	return s.Result
}

func (s *DescribeResultCountResponseBodyResultObject) GetTotal() *string {
	return s.Total
}

func (s *DescribeResultCountResponseBodyResultObject) SetResult(v string) *DescribeResultCountResponseBodyResultObject {
	s.Result = &v
	return s
}

func (s *DescribeResultCountResponseBodyResultObject) SetTotal(v string) *DescribeResultCountResponseBodyResultObject {
	s.Total = &v
	return s
}

func (s *DescribeResultCountResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
