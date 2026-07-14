// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteGroupByIdResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteGroupByIdResponseBody
	GetCode() *string
	SetData(v string) *DeleteGroupByIdResponseBody
	GetData() *string
	SetMessage(v string) *DeleteGroupByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteGroupByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteGroupByIdResponseBody
	GetSuccess() *bool
}

type DeleteGroupByIdResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1223-1111
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteGroupByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupByIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupByIdResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteGroupByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteGroupByIdResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteGroupByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteGroupByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGroupByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteGroupByIdResponseBody) SetAccessDeniedDetail(v string) *DeleteGroupByIdResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteGroupByIdResponseBody) SetCode(v string) *DeleteGroupByIdResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteGroupByIdResponseBody) SetData(v string) *DeleteGroupByIdResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteGroupByIdResponseBody) SetMessage(v string) *DeleteGroupByIdResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteGroupByIdResponseBody) SetRequestId(v string) *DeleteGroupByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupByIdResponseBody) SetSuccess(v bool) *DeleteGroupByIdResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteGroupByIdResponseBody) Validate() error {
	return dara.Validate(s)
}
