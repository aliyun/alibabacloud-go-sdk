// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineFlowVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *OnlineFlowVersionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *OnlineFlowVersionResponseBody
	GetCode() *string
	SetMessage(v string) *OnlineFlowVersionResponseBody
	GetMessage() *string
	SetRequestId(v string) *OnlineFlowVersionResponseBody
	GetRequestId() *string
	SetResponse(v map[string]interface{}) *OnlineFlowVersionResponseBody
	GetResponse() map[string]interface{}
	SetSuccess(v bool) *OnlineFlowVersionResponseBody
	GetSuccess() *bool
}

type OnlineFlowVersionResponseBody struct {
	// Details of access denial.
	//
	// example:
	//
	// 无
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// Status code.
	//
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Error message.
	//
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Content of the returned data.
	//
	// example:
	//
	// 无
	Response map[string]interface{} `json:"Response,omitempty" xml:"Response,omitempty"`
	// Indicates whether the operation was successful. true means success, false means failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OnlineFlowVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineFlowVersionResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineFlowVersionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *OnlineFlowVersionResponseBody) GetCode() *string {
	return s.Code
}

func (s *OnlineFlowVersionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OnlineFlowVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineFlowVersionResponseBody) GetResponse() map[string]interface{} {
	return s.Response
}

func (s *OnlineFlowVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OnlineFlowVersionResponseBody) SetAccessDeniedDetail(v string) *OnlineFlowVersionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *OnlineFlowVersionResponseBody) SetCode(v string) *OnlineFlowVersionResponseBody {
	s.Code = &v
	return s
}

func (s *OnlineFlowVersionResponseBody) SetMessage(v string) *OnlineFlowVersionResponseBody {
	s.Message = &v
	return s
}

func (s *OnlineFlowVersionResponseBody) SetRequestId(v string) *OnlineFlowVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineFlowVersionResponseBody) SetResponse(v map[string]interface{}) *OnlineFlowVersionResponseBody {
	s.Response = v
	return s
}

func (s *OnlineFlowVersionResponseBody) SetSuccess(v bool) *OnlineFlowVersionResponseBody {
	s.Success = &v
	return s
}

func (s *OnlineFlowVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
