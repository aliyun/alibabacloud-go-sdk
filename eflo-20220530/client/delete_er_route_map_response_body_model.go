// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteErRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteErRouteMapResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteErRouteMapResponseBody
	GetCode() *int32
	SetContent(v interface{}) *DeleteErRouteMapResponseBody
	GetContent() interface{}
	SetMessage(v string) *DeleteErRouteMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteErRouteMapResponseBody
	GetRequestId() *string
}

type DeleteErRouteMapResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The response status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Response body
	//
	// example:
	//
	// {}
	Content interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
	// The error message. (If the instance is in the Exception state, the exception cause is prompted.)
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// AC8C713A-A9F4-5984-A5E1-76496DF35153
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteErRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteErRouteMapResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteErRouteMapResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteErRouteMapResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *DeleteErRouteMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteErRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteErRouteMapResponseBody) SetAccessDeniedDetail(v string) *DeleteErRouteMapResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteErRouteMapResponseBody) SetCode(v int32) *DeleteErRouteMapResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteErRouteMapResponseBody) SetContent(v interface{}) *DeleteErRouteMapResponseBody {
	s.Content = v
	return s
}

func (s *DeleteErRouteMapResponseBody) SetMessage(v string) *DeleteErRouteMapResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteErRouteMapResponseBody) SetRequestId(v string) *DeleteErRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteErRouteMapResponseBody) Validate() error {
	return dara.Validate(s)
}
