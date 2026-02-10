// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalSignOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CreateDigitalSignOrderResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CreateDigitalSignOrderResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *CreateDigitalSignOrderResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *CreateDigitalSignOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDigitalSignOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDigitalSignOrderResponseBody
	GetSuccess() *bool
}

type CreateDigitalSignOrderResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {\\"signId\\": \\"20067468291\\", \\"signCode\\": \\"SIGN_100000184736042_1744164758835_hpMd1\\", \\"signOrderId\\": 22469795330, \\"signName\\": u\\"\\u8d5b\\u745e\\u5a05\\u808c\\u80a4\\u7ba1\\u7406\\"}
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2D0133B9-6C0D-0BAE-8161-1EEF9E2D4069
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDigitalSignOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalSignOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDigitalSignOrderResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CreateDigitalSignOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDigitalSignOrderResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *CreateDigitalSignOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDigitalSignOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDigitalSignOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDigitalSignOrderResponseBody) SetAccessDeniedDetail(v string) *CreateDigitalSignOrderResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CreateDigitalSignOrderResponseBody) SetCode(v string) *CreateDigitalSignOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDigitalSignOrderResponseBody) SetData(v map[string]interface{}) *CreateDigitalSignOrderResponseBody {
	s.Data = v
	return s
}

func (s *CreateDigitalSignOrderResponseBody) SetMessage(v string) *CreateDigitalSignOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDigitalSignOrderResponseBody) SetRequestId(v string) *CreateDigitalSignOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDigitalSignOrderResponseBody) SetSuccess(v bool) *CreateDigitalSignOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDigitalSignOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
