// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomerModuleBasicInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCustomerModuleBasicInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int64) *DescribeCustomerModuleBasicInfoResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *DescribeCustomerModuleBasicInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCustomerModuleBasicInfoResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeCustomerModuleBasicInfoResponseBodyResultObject) *DescribeCustomerModuleBasicInfoResponseBody
	GetResultObject() *DescribeCustomerModuleBasicInfoResponseBodyResultObject
	SetSuccess(v bool) *DescribeCustomerModuleBasicInfoResponseBody
	GetSuccess() *bool
}

type DescribeCustomerModuleBasicInfoResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID
	//
	// example:
	//
	// 055f1546-f465-4c92-a2da-bfb86abe6f56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result.
	ResultObject *DescribeCustomerModuleBasicInfoResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Indicates whether the application configuration information was successfully retrieved. Possible values: true: Success; false: Failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCustomerModuleBasicInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleBasicInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) GetResultObject() *DescribeCustomerModuleBasicInfoResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) SetCode(v string) *DescribeCustomerModuleBasicInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) SetHttpStatusCode(v int64) *DescribeCustomerModuleBasicInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) SetMessage(v string) *DescribeCustomerModuleBasicInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) SetRequestId(v string) *DescribeCustomerModuleBasicInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) SetResultObject(v *DescribeCustomerModuleBasicInfoResponseBodyResultObject) *DescribeCustomerModuleBasicInfoResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) SetSuccess(v bool) *DescribeCustomerModuleBasicInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomerModuleBasicInfoResponseBodyResultObject struct {
	// Customer model name.
	//
	// example:
	//
	// ModuleName
	CustomerModuleName *string `json:"CustomerModuleName,omitempty" xml:"CustomerModuleName,omitempty"`
	// Remarks.
	//
	// example:
	//
	// 备注。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Model identifier.
	//
	// example:
	//
	// ModuleName
	InnerModuleName *string `json:"InnerModuleName,omitempty" xml:"InnerModuleName,omitempty"`
	// Module type.
	//
	// example:
	//
	// PMML
	ModuleType *string `json:"ModuleType,omitempty" xml:"ModuleType,omitempty"`
	// Test address.
	//
	// example:
	//
	// model.pmml
	StorePath *string `json:"StorePath,omitempty" xml:"StorePath,omitempty"`
}

func (s DescribeCustomerModuleBasicInfoResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomerModuleBasicInfoResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) GetCustomerModuleName() *string {
	return s.CustomerModuleName
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) GetInnerModuleName() *string {
	return s.InnerModuleName
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) GetModuleType() *string {
	return s.ModuleType
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) GetStorePath() *string {
	return s.StorePath
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) SetCustomerModuleName(v string) *DescribeCustomerModuleBasicInfoResponseBodyResultObject {
	s.CustomerModuleName = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) SetDescription(v string) *DescribeCustomerModuleBasicInfoResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) SetInnerModuleName(v string) *DescribeCustomerModuleBasicInfoResponseBodyResultObject {
	s.InnerModuleName = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) SetModuleType(v string) *DescribeCustomerModuleBasicInfoResponseBodyResultObject {
	s.ModuleType = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) SetStorePath(v string) *DescribeCustomerModuleBasicInfoResponseBodyResultObject {
	s.StorePath = &v
	return s
}

func (s *DescribeCustomerModuleBasicInfoResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
