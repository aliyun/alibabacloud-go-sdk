// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstagramPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteInstagramPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteInstagramPageResponseBody
	GetCode() *string
	SetData(v string) *DeleteInstagramPageResponseBody
	GetData() *string
	SetMessage(v string) *DeleteInstagramPageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstagramPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstagramPageResponseBody
	GetSuccess() *bool
}

type DeleteInstagramPageResponseBody struct {
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
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// gfdg435t-hf544**
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteInstagramPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstagramPageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstagramPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteInstagramPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstagramPageResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteInstagramPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstagramPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstagramPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstagramPageResponseBody) SetAccessDeniedDetail(v string) *DeleteInstagramPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetCode(v string) *DeleteInstagramPageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetData(v string) *DeleteInstagramPageResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetMessage(v string) *DeleteInstagramPageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetRequestId(v string) *DeleteInstagramPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) SetSuccess(v bool) *DeleteInstagramPageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstagramPageResponseBody) Validate() error {
	return dara.Validate(s)
}
