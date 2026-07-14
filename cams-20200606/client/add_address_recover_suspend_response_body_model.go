// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressRecoverSuspendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddAddressRecoverSuspendResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddAddressRecoverSuspendResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *AddAddressRecoverSuspendResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *AddAddressRecoverSuspendResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddAddressRecoverSuspendResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddAddressRecoverSuspendResponseBody
	GetSuccess() *bool
}

type AddAddressRecoverSuspendResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code.
	//
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// 示例值
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message. A success message is returned if the request succeeds. A failure reason is returned if the request fails.
	//
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 示例值示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: Successful.
	//
	// - false: Failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddAddressRecoverSuspendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAddressRecoverSuspendResponseBody) GoString() string {
	return s.String()
}

func (s *AddAddressRecoverSuspendResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddAddressRecoverSuspendResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddAddressRecoverSuspendResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *AddAddressRecoverSuspendResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddAddressRecoverSuspendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAddressRecoverSuspendResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddAddressRecoverSuspendResponseBody) SetAccessDeniedDetail(v string) *AddAddressRecoverSuspendResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddAddressRecoverSuspendResponseBody) SetCode(v string) *AddAddressRecoverSuspendResponseBody {
	s.Code = &v
	return s
}

func (s *AddAddressRecoverSuspendResponseBody) SetData(v map[string]interface{}) *AddAddressRecoverSuspendResponseBody {
	s.Data = v
	return s
}

func (s *AddAddressRecoverSuspendResponseBody) SetMessage(v string) *AddAddressRecoverSuspendResponseBody {
	s.Message = &v
	return s
}

func (s *AddAddressRecoverSuspendResponseBody) SetRequestId(v string) *AddAddressRecoverSuspendResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAddressRecoverSuspendResponseBody) SetSuccess(v bool) *AddAddressRecoverSuspendResponseBody {
	s.Success = &v
	return s
}

func (s *AddAddressRecoverSuspendResponseBody) Validate() error {
	return dara.Validate(s)
}
