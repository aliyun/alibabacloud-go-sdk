// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessengerPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteMessengerPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteMessengerPageResponseBody
	GetCode() *string
	SetData(v string) *DeleteMessengerPageResponseBody
	GetData() *string
	SetMessage(v string) *DeleteMessengerPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMessengerPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMessengerPageResponseBody
	GetSuccess() *bool
}

type DeleteMessengerPageResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// JHGJG-65HFGF**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMessengerPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessengerPageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMessengerPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteMessengerPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMessengerPageResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteMessengerPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMessengerPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMessengerPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMessengerPageResponseBody) SetAccessDeniedDetail(v string) *DeleteMessengerPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteMessengerPageResponseBody) SetCode(v string) *DeleteMessengerPageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMessengerPageResponseBody) SetData(v string) *DeleteMessengerPageResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteMessengerPageResponseBody) SetMessage(v string) *DeleteMessengerPageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMessengerPageResponseBody) SetRequestId(v string) *DeleteMessengerPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMessengerPageResponseBody) SetSuccess(v bool) *DeleteMessengerPageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMessengerPageResponseBody) Validate() error {
	return dara.Validate(s)
}
