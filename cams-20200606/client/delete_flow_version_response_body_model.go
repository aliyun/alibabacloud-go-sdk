// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteFlowVersionResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFlowVersionResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *DeleteFlowVersionResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *DeleteFlowVersionResponseBody
	GetSuccess() *bool
}

type DeleteFlowVersionResponseBody struct {
	// Detailed reason for access denial.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Status code.
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
	// Indicates whether the operation was successful. Values: true: success; false: failure.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowVersionResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *DeleteFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFlowVersionResponseBody) SetAccessDeniedDetail(v string) *DeleteFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteFlowVersionResponseBody) SetCode(v string) *DeleteFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlowVersionResponseBody) SetMessage(v string) *DeleteFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlowVersionResponseBody) SetRequestId(v string) *DeleteFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowVersionResponseBody) SetResponse(v map[string]interface{}) *DeleteFlowVersionResponseBody {
	s.Response = v
	return s
}

func (s *DeleteFlowVersionResponseBody) SetSuccess(v bool) *DeleteFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
