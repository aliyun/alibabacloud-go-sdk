// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImportFieldResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *ImportFieldResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *ImportFieldResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportFieldResponseBody
	GetRequestId() *string
	SetResultObject(v *ImportFieldResponseBodyResultObject) *ImportFieldResponseBody
	GetResultObject() *ImportFieldResponseBodyResultObject
	SetSuccess(v bool) *ImportFieldResponseBody
	GetSuccess() *bool
}

type ImportFieldResponseBody struct {
	// API status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned result
	ResultObject *ImportFieldResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// Indicator of whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImportFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportFieldResponseBody) GoString() string {
	return s.String()
}

func (s *ImportFieldResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImportFieldResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ImportFieldResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportFieldResponseBody) GetResultObject() *ImportFieldResponseBodyResultObject {
	return s.ResultObject
}

func (s *ImportFieldResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportFieldResponseBody) SetCode(v string) *ImportFieldResponseBody {
	s.Code = &v
	return s
}

func (s *ImportFieldResponseBody) SetHttpStatusCode(v string) *ImportFieldResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportFieldResponseBody) SetMessage(v string) *ImportFieldResponseBody {
	s.Message = &v
	return s
}

func (s *ImportFieldResponseBody) SetRequestId(v string) *ImportFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportFieldResponseBody) SetResultObject(v *ImportFieldResponseBodyResultObject) *ImportFieldResponseBody {
	s.ResultObject = v
	return s
}

func (s *ImportFieldResponseBody) SetSuccess(v bool) *ImportFieldResponseBody {
	s.Success = &v
	return s
}

func (s *ImportFieldResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImportFieldResponseBodyResultObject struct {
	// Names of fields that failed to upload
	//
	// example:
	//
	// test
	FailFieldNames *string `json:"FailFieldNames,omitempty" xml:"FailFieldNames,omitempty"`
	// Number of successful executions.
	//
	// example:
	//
	// 7
	SuccessNum *int32 `json:"SuccessNum,omitempty" xml:"SuccessNum,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 7
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ImportFieldResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s ImportFieldResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *ImportFieldResponseBodyResultObject) GetFailFieldNames() *string {
	return s.FailFieldNames
}

func (s *ImportFieldResponseBodyResultObject) GetSuccessNum() *int32 {
	return s.SuccessNum
}

func (s *ImportFieldResponseBodyResultObject) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *ImportFieldResponseBodyResultObject) SetFailFieldNames(v string) *ImportFieldResponseBodyResultObject {
	s.FailFieldNames = &v
	return s
}

func (s *ImportFieldResponseBodyResultObject) SetSuccessNum(v int32) *ImportFieldResponseBodyResultObject {
	s.SuccessNum = &v
	return s
}

func (s *ImportFieldResponseBodyResultObject) SetTotalNum(v int32) *ImportFieldResponseBodyResultObject {
	s.TotalNum = &v
	return s
}

func (s *ImportFieldResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
