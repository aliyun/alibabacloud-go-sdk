// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateFlowVersionResponseBody
	GetCode() *string
	SetMessage(v string) *CreateFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateFlowVersionResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *CreateFlowVersionResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *CreateFlowVersionResponseBody
	GetSuccess() *bool
}

type CreateFlowVersionResponseBody struct {
	// Details of access denial; this field is only returned when RAM verification fails.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Request status code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// 无
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Content of the returned data.
	//
	// example:
	//
	// 无
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Indicates whether the operation was successful. Values: true for success, false for failure.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFlowVersionResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *CreateFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateFlowVersionResponseBody) SetAccessDeniedDetail(v string) *CreateFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateFlowVersionResponseBody) SetCode(v string) *CreateFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateFlowVersionResponseBody) SetMessage(v string) *CreateFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateFlowVersionResponseBody) SetRequestId(v string) *CreateFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFlowVersionResponseBody) SetResponse(v map[string]interface{}) *CreateFlowVersionResponseBody {
	s.Response = v
	return s
}

func (s *CreateFlowVersionResponseBody) SetSuccess(v bool) *CreateFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
