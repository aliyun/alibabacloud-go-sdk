// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPackageTypeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPackageTypeInfoResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *QueryPackageTypeInfoResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *QueryPackageTypeInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryPackageTypeInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryPackageTypeInfoResponseBody
	GetSuccess() *bool
}

type QueryPackageTypeInfoResponseBody struct {
	// The status code.
	//
	// - OK indicates that the request is successful.
	//
	// - For other error codes, see the [error code list](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// Example
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	//
	// example:
	//
	// {"efactorVerification":"平台风控-二要素核验查询套餐包"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The response message.
	//
	// example:
	//
	// Example
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// Example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - true: The request was successful.
	//
	// - false: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryPackageTypeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPackageTypeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPackageTypeInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPackageTypeInfoResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *QueryPackageTypeInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPackageTypeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPackageTypeInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryPackageTypeInfoResponseBody) SetCode(v string) *QueryPackageTypeInfoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) SetData(v map[string]interface{}) *QueryPackageTypeInfoResponseBody {
	s.Data = v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) SetMessage(v string) *QueryPackageTypeInfoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) SetRequestId(v string) *QueryPackageTypeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) SetSuccess(v bool) *QueryPackageTypeInfoResponseBody {
	s.Success = &v
	return s
}

func (s *QueryPackageTypeInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
