// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLargeModelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateLargeModelResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateLargeModelResponseBody
	GetCode() *string
	SetData(v bool) *UpdateLargeModelResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateLargeModelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLargeModelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLargeModelResponseBody
	GetSuccess() *bool
}

type UpdateLargeModelResponseBody struct {
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLargeModelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLargeModelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLargeModelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateLargeModelResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLargeModelResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateLargeModelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLargeModelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLargeModelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLargeModelResponseBody) SetAccessDeniedDetail(v string) *UpdateLargeModelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLargeModelResponseBody) SetCode(v string) *UpdateLargeModelResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLargeModelResponseBody) SetData(v bool) *UpdateLargeModelResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateLargeModelResponseBody) SetMessage(v string) *UpdateLargeModelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLargeModelResponseBody) SetRequestId(v string) *UpdateLargeModelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLargeModelResponseBody) SetSuccess(v bool) *UpdateLargeModelResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLargeModelResponseBody) Validate() error {
	return dara.Validate(s)
}
