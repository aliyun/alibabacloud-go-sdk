// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRCSSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateRCSSignatureResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateRCSSignatureResponseBody
	GetCode() *string
	SetData(v bool) *UpdateRCSSignatureResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateRCSSignatureResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRCSSignatureResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRCSSignatureResponseBody
	GetSuccess() *bool
}

type UpdateRCSSignatureResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRCSSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRCSSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRCSSignatureResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateRCSSignatureResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRCSSignatureResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateRCSSignatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRCSSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRCSSignatureResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRCSSignatureResponseBody) SetAccessDeniedDetail(v string) *UpdateRCSSignatureResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateRCSSignatureResponseBody) SetCode(v string) *UpdateRCSSignatureResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRCSSignatureResponseBody) SetData(v bool) *UpdateRCSSignatureResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRCSSignatureResponseBody) SetMessage(v string) *UpdateRCSSignatureResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRCSSignatureResponseBody) SetRequestId(v string) *UpdateRCSSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRCSSignatureResponseBody) SetSuccess(v bool) *UpdateRCSSignatureResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRCSSignatureResponseBody) Validate() error {
	return dara.Validate(s)
}
