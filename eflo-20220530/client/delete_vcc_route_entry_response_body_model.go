// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVccRouteEntryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteVccRouteEntryResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *DeleteVccRouteEntryResponseBody
	GetCode() *int32
	SetContent(v interface{}) *DeleteVccRouteEntryResponseBody
	GetContent() interface{}
	SetMessage(v string) *DeleteVccRouteEntryResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteVccRouteEntryResponseBody
	GetRequestId() *string
}

type DeleteVccRouteEntryResponseBody struct {
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
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID of the current request
	//
	// example:
	//
	// 0901F411-28FA-5B9C-BAEE-7776463FF0DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVccRouteEntryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVccRouteEntryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVccRouteEntryResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteVccRouteEntryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteVccRouteEntryResponseBody) GetContent() interface{} {
	return s.Content
}

func (s *DeleteVccRouteEntryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteVccRouteEntryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVccRouteEntryResponseBody) SetAccessDeniedDetail(v string) *DeleteVccRouteEntryResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteVccRouteEntryResponseBody) SetCode(v int32) *DeleteVccRouteEntryResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteVccRouteEntryResponseBody) SetContent(v interface{}) *DeleteVccRouteEntryResponseBody {
	s.Content = v
	return s
}

func (s *DeleteVccRouteEntryResponseBody) SetMessage(v string) *DeleteVccRouteEntryResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteVccRouteEntryResponseBody) SetRequestId(v string) *DeleteVccRouteEntryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVccRouteEntryResponseBody) Validate() error {
	return dara.Validate(s)
}
