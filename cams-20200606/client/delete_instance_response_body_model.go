// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteInstanceResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteInstanceResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstanceResponseBody
	GetSuccess() *bool
}

type DeleteInstanceResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The request status code.
	//
	// - A value of `OK` means the request was successful.
	//
	// - For other error codes, see the [Error Code List](https://help.aliyun.com/document_detail/196974.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// NULL
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique ID for the request. Use it for troubleshooting.
	//
	// example:
	//
	// 3993939******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Specifies whether the request was successful. Valid values:
	//
	// - `true`: The request was successful.
	//
	// - `false`: The request failed.
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstanceResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstanceResponseBody) SetAccessDeniedDetail(v string) *DeleteInstanceResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetCode(v string) *DeleteInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetMessage(v string) *DeleteInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetRequestId(v string) *DeleteInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstanceResponseBody) SetSuccess(v bool) *DeleteInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
