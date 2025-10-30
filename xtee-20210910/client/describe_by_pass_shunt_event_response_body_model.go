// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeByPassShuntEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeByPassShuntEventResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeByPassShuntEventResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeByPassShuntEventResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeByPassShuntEventResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeByPassShuntEventResponseBodyResultObject) *DescribeByPassShuntEventResponseBody
	GetResultObject() *DescribeByPassShuntEventResponseBodyResultObject
	SetSuccess(v bool) *DescribeByPassShuntEventResponseBody
	GetSuccess() *bool
}

type DescribeByPassShuntEventResponseBody struct {
	// Error code.
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
	// Error details
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
	// Returned object
	ResultObject *DescribeByPassShuntEventResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Whether it was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeByPassShuntEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeByPassShuntEventResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeByPassShuntEventResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeByPassShuntEventResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeByPassShuntEventResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeByPassShuntEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeByPassShuntEventResponseBody) GetResultObject() *DescribeByPassShuntEventResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeByPassShuntEventResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeByPassShuntEventResponseBody) SetCode(v string) *DescribeByPassShuntEventResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeByPassShuntEventResponseBody) SetHttpStatusCode(v string) *DescribeByPassShuntEventResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeByPassShuntEventResponseBody) SetMessage(v string) *DescribeByPassShuntEventResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeByPassShuntEventResponseBody) SetRequestId(v string) *DescribeByPassShuntEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeByPassShuntEventResponseBody) SetResultObject(v *DescribeByPassShuntEventResponseBodyResultObject) *DescribeByPassShuntEventResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeByPassShuntEventResponseBody) SetSuccess(v bool) *DescribeByPassShuntEventResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeByPassShuntEventResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeByPassShuntEventResponseBodyResultObject struct {
	// Event name.
	//
	// example:
	//
	// 营销风险识别_增强版
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
}

func (s DescribeByPassShuntEventResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeByPassShuntEventResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeByPassShuntEventResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeByPassShuntEventResponseBodyResultObject) SetEventName(v string) *DescribeByPassShuntEventResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeByPassShuntEventResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
