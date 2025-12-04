// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateErRouteMapResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateErRouteMapResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *UpdateErRouteMapResponseBody
	GetCode() *int32
	SetContent(v map[string]interface{}) *UpdateErRouteMapResponseBody
	GetContent() map[string]interface{}
	SetMessage(v string) *UpdateErRouteMapResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateErRouteMapResponseBody
	GetRequestId() *string
}

type UpdateErRouteMapResponseBody struct {
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
	// The response parameters.
	//
	// example:
	//
	// {}
	Content map[string]interface{} `json:"Content,omitempty" xml:"Content,omitempty"`
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
	// BDBCC783-84CA-5733-8EEA-645C88B9009C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateErRouteMapResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateErRouteMapResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateErRouteMapResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateErRouteMapResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateErRouteMapResponseBody) GetContent() map[string]interface{} {
	return s.Content
}

func (s *UpdateErRouteMapResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateErRouteMapResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateErRouteMapResponseBody) SetAccessDeniedDetail(v string) *UpdateErRouteMapResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateErRouteMapResponseBody) SetCode(v int32) *UpdateErRouteMapResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateErRouteMapResponseBody) SetContent(v map[string]interface{}) *UpdateErRouteMapResponseBody {
	s.Content = v
	return s
}

func (s *UpdateErRouteMapResponseBody) SetMessage(v string) *UpdateErRouteMapResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateErRouteMapResponseBody) SetRequestId(v string) *UpdateErRouteMapResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateErRouteMapResponseBody) Validate() error {
	return dara.Validate(s)
}
