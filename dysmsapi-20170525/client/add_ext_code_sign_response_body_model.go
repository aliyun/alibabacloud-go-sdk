// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddExtCodeSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AddExtCodeSignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AddExtCodeSignResponseBody
	GetCode() *string
	SetData(v bool) *AddExtCodeSignResponseBody
	GetData() *bool
	SetMessage(v string) *AddExtCodeSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddExtCodeSignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddExtCodeSignResponseBody
	GetSuccess() *bool
}

type AddExtCodeSignResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
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

func (s AddExtCodeSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddExtCodeSignResponseBody) GoString() string {
	return s.String()
}

func (s *AddExtCodeSignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AddExtCodeSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddExtCodeSignResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddExtCodeSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddExtCodeSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddExtCodeSignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddExtCodeSignResponseBody) SetAccessDeniedDetail(v string) *AddExtCodeSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetCode(v string) *AddExtCodeSignResponseBody {
	s.Code = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetData(v bool) *AddExtCodeSignResponseBody {
	s.Data = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetMessage(v string) *AddExtCodeSignResponseBody {
	s.Message = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetRequestId(v string) *AddExtCodeSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddExtCodeSignResponseBody) SetSuccess(v bool) *AddExtCodeSignResponseBody {
	s.Success = &v
	return s
}

func (s *AddExtCodeSignResponseBody) Validate() error {
	return dara.Validate(s)
}
