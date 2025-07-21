// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PublishFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *PublishFlowResponseBody
	GetCode() *string
	SetMessage(v string) *PublishFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *PublishFlowResponseBody
	GetRequestId() *string
}

type PublishFlowResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// If OK is returned, the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// None
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 90E63D28-E31D-1EB2-8939-A9486641****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PublishFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PublishFlowResponseBody) GoString() string {
	return s.String()
}

func (s *PublishFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PublishFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *PublishFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PublishFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PublishFlowResponseBody) SetAccessDeniedDetail(v string) *PublishFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PublishFlowResponseBody) SetCode(v string) *PublishFlowResponseBody {
	s.Code = &v
	return s
}

func (s *PublishFlowResponseBody) SetMessage(v string) *PublishFlowResponseBody {
	s.Message = &v
	return s
}

func (s *PublishFlowResponseBody) SetRequestId(v string) *PublishFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *PublishFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
