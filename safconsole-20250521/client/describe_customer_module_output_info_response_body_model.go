// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleOutputInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *DescribeCustomerModuleOutputInfoResponseBody
	GetCode() *int64
	SetHttpStatusCode(v int64) *DescribeCustomerModuleOutputInfoResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *DescribeCustomerModuleOutputInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomerModuleOutputInfoResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeCustomerModuleOutputInfoResponseBodyResultObject) *DescribeCustomerModuleOutputInfoResponseBody
	GetResultObject() *DescribeCustomerModuleOutputInfoResponseBodyResultObject
	SetSuccess(v bool) *DescribeCustomerModuleOutputInfoResponseBody
	GetSuccess() *bool
}

type DescribeCustomerModuleOutputInfoResponseBody struct {
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 055f1546-f465-4c92-a2da-bfb86abe6f56
	RequestId    *string                                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeCustomerModuleOutputInfoResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomerModuleOutputInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleOutputInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) GetResultObject() *DescribeCustomerModuleOutputInfoResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) SetCode(v int64) *DescribeCustomerModuleOutputInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) SetHttpStatusCode(v int64) *DescribeCustomerModuleOutputInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) SetMessage(v string) *DescribeCustomerModuleOutputInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) SetRequestId(v string) *DescribeCustomerModuleOutputInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) SetResultObject(v *DescribeCustomerModuleOutputInfoResponseBodyResultObject) *DescribeCustomerModuleOutputInfoResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) SetSuccess(v bool) *DescribeCustomerModuleOutputInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomerModuleOutputInfoResponseBodyResultObject struct {
	// example:
	//
	// 2
	FinalScoreFormat *string `json:"FinalScoreFormat,omitempty" xml:"FinalScoreFormat,omitempty"`
	// example:
	//
	// score
	ProcessExpression *string `json:"ProcessExpression,omitempty" xml:"ProcessExpression,omitempty"`
	// example:
	//
	// customer/xxxxxxxxx/xxxxxxxx.pmml
	TestFile *string `json:"TestFile,omitempty" xml:"TestFile,omitempty"`
}

func (s DescribeCustomerModuleOutputInfoResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleOutputInfoResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleOutputInfoResponseBodyResultObject) GetFinalScoreFormat() *string {
	return s.FinalScoreFormat
}

func (s *DescribeCustomerModuleOutputInfoResponseBodyResultObject) GetProcessExpression() *string {
	return s.ProcessExpression
}

func (s *DescribeCustomerModuleOutputInfoResponseBodyResultObject) GetTestFile() *string {
	return s.TestFile
}

func (s *DescribeCustomerModuleOutputInfoResponseBodyResultObject) SetFinalScoreFormat(v string) *DescribeCustomerModuleOutputInfoResponseBodyResultObject {
	s.FinalScoreFormat = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBodyResultObject) SetProcessExpression(v string) *DescribeCustomerModuleOutputInfoResponseBodyResultObject {
	s.ProcessExpression = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBodyResultObject) SetTestFile(v string) *DescribeCustomerModuleOutputInfoResponseBodyResultObject {
	s.TestFile = &v
	return s
}

func (s *DescribeCustomerModuleOutputInfoResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
