// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExtCodeSignResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteExtCodeSignResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteExtCodeSignResponseBody
	GetCode() *string
	SetData(v bool) *DeleteExtCodeSignResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteExtCodeSignResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteExtCodeSignResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteExtCodeSignResponseBody
	GetSuccess() *bool
}

type DeleteExtCodeSignResponseBody struct {
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

func (s DeleteExtCodeSignResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteExtCodeSignResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteExtCodeSignResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteExtCodeSignResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteExtCodeSignResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteExtCodeSignResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteExtCodeSignResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteExtCodeSignResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteExtCodeSignResponseBody) SetAccessDeniedDetail(v string) *DeleteExtCodeSignResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetCode(v string) *DeleteExtCodeSignResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetData(v bool) *DeleteExtCodeSignResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetMessage(v string) *DeleteExtCodeSignResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetRequestId(v string) *DeleteExtCodeSignResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) SetSuccess(v bool) *DeleteExtCodeSignResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteExtCodeSignResponseBody) Validate() error {
	return dara.Validate(s)
}
