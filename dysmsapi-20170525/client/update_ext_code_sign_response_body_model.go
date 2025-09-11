// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExtCodeSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateExtCodeSignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateExtCodeSignResponseBody
	GetCode() *string
	SetData(v bool) *UpdateExtCodeSignResponseBody
	GetData() *bool
	SetMessage(v string) *UpdateExtCodeSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateExtCodeSignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateExtCodeSignResponseBody
	GetSuccess() *bool
}

type UpdateExtCodeSignResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// false
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateExtCodeSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateExtCodeSignResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateExtCodeSignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateExtCodeSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateExtCodeSignResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateExtCodeSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateExtCodeSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateExtCodeSignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateExtCodeSignResponseBody) SetAccessDeniedDetail(v string) *UpdateExtCodeSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetCode(v string) *UpdateExtCodeSignResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetData(v bool) *UpdateExtCodeSignResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetMessage(v string) *UpdateExtCodeSignResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetRequestId(v string) *UpdateExtCodeSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) SetSuccess(v bool) *UpdateExtCodeSignResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateExtCodeSignResponseBody) Validate() error {
	return dara.Validate(s)
}
