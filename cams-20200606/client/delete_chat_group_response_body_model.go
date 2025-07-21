// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChatGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteChatGroupResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteChatGroupResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteChatGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteChatGroupResponseBody
	GetRequestId() *string
	SetResult(v int64) *DeleteChatGroupResponseBody
	GetResult() *int64
	SetSuccess(v bool) *DeleteChatGroupResponseBody
	GetSuccess() *bool
}

type DeleteChatGroupResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E939E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteChatGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChatGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChatGroupResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteChatGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteChatGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteChatGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChatGroupResponseBody) GetResult() *int64 {
	return s.Result
}

func (s *DeleteChatGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteChatGroupResponseBody) SetAccessDeniedDetail(v string) *DeleteChatGroupResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetCode(v string) *DeleteChatGroupResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetMessage(v string) *DeleteChatGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetRequestId(v string) *DeleteChatGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetResult(v int64) *DeleteChatGroupResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteChatGroupResponseBody) SetSuccess(v bool) *DeleteChatGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteChatGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
