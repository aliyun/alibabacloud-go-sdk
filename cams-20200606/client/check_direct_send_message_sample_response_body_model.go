// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDirectSendMessageSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CheckDirectSendMessageSampleResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CheckDirectSendMessageSampleResponseBody
	GetCode() *string
	SetData(v *CheckDirectSendMessageSampleResponseBodyData) *CheckDirectSendMessageSampleResponseBody
	GetData() *CheckDirectSendMessageSampleResponseBodyData
	SetMessage(v string) *CheckDirectSendMessageSampleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CheckDirectSendMessageSampleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CheckDirectSendMessageSampleResponseBody
	GetSuccess() *bool
}

type CheckDirectSendMessageSampleResponseBody struct {
	// Details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code of the request.
	//
	// - `OK` indicates that the request was successful.
	//
	// - For other error codes, see [Error codes](https://www.alibabacloud.com/help/zh/cams/latest/api-error-codes).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *CheckDirectSendMessageSampleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// example
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. Use this ID to trace logs and troubleshoot issues.
	//
	// example:
	//
	// example
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the request was successful. Valid values:
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

func (s CheckDirectSendMessageSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDirectSendMessageSampleResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDirectSendMessageSampleResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CheckDirectSendMessageSampleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CheckDirectSendMessageSampleResponseBody) GetData() *CheckDirectSendMessageSampleResponseBodyData {
	return s.Data
}

func (s *CheckDirectSendMessageSampleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CheckDirectSendMessageSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDirectSendMessageSampleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CheckDirectSendMessageSampleResponseBody) SetAccessDeniedDetail(v string) *CheckDirectSendMessageSampleResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CheckDirectSendMessageSampleResponseBody) SetCode(v string) *CheckDirectSendMessageSampleResponseBody {
	s.Code = &v
	return s
}

func (s *CheckDirectSendMessageSampleResponseBody) SetData(v *CheckDirectSendMessageSampleResponseBodyData) *CheckDirectSendMessageSampleResponseBody {
	s.Data = v
	return s
}

func (s *CheckDirectSendMessageSampleResponseBody) SetMessage(v string) *CheckDirectSendMessageSampleResponseBody {
	s.Message = &v
	return s
}

func (s *CheckDirectSendMessageSampleResponseBody) SetRequestId(v string) *CheckDirectSendMessageSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDirectSendMessageSampleResponseBody) SetSuccess(v bool) *CheckDirectSendMessageSampleResponseBody {
	s.Success = &v
	return s
}

func (s *CheckDirectSendMessageSampleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckDirectSendMessageSampleResponseBodyData struct {
	// The category of the WhatsApp template. Valid values:
	//
	// - **UTILITY**: A template related to a specific, agreed-upon transaction.
	//
	// - **MARKETING**: A template for promotional or marketing purposes.
	//
	// example:
	//
	// example
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Whether the validation was successful. Valid values:
	//
	// - true: The validation succeeded.
	//
	// - false: The validation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CheckDirectSendMessageSampleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CheckDirectSendMessageSampleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CheckDirectSendMessageSampleResponseBodyData) GetCategory() *string {
	return s.Category
}

func (s *CheckDirectSendMessageSampleResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CheckDirectSendMessageSampleResponseBodyData) SetCategory(v string) *CheckDirectSendMessageSampleResponseBodyData {
	s.Category = &v
	return s
}

func (s *CheckDirectSendMessageSampleResponseBodyData) SetSuccess(v bool) *CheckDirectSendMessageSampleResponseBodyData {
	s.Success = &v
	return s
}

func (s *CheckDirectSendMessageSampleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
