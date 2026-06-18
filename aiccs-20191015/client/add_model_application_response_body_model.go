// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddModelApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddModelApplicationResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddModelApplicationResponseBody
	GetCode() *string
	SetData(v string) *AddModelApplicationResponseBody
	GetData() *string
	SetMessage(v string) *AddModelApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddModelApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddModelApplicationResponseBody
	GetSuccess() *bool
}

type AddModelApplicationResponseBody struct {
	// The detailed reason for the access denied error.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The status code. A value of OK indicates a successful request.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned application code.
	//
	// example:
	//
	// 325****2D2
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message that explains the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C93B345-F702-5449-BA7E-7D110D4BF798
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddModelApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddModelApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *AddModelApplicationResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddModelApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddModelApplicationResponseBody) GetData() *string {
	return s.Data
}

func (s *AddModelApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddModelApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddModelApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddModelApplicationResponseBody) SetAccessDeniedDetail(v string) *AddModelApplicationResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetCode(v string) *AddModelApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetData(v string) *AddModelApplicationResponseBody {
	s.Data = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetMessage(v string) *AddModelApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetRequestId(v string) *AddModelApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddModelApplicationResponseBody) SetSuccess(v bool) *AddModelApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *AddModelApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}
