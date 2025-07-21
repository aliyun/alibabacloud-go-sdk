// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DeleteFlowResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *DeleteFlowResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFlowResponseBody
	GetRequestId() *string
}

type DeleteFlowResponseBody struct {
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

func (s DeleteFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFlowResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFlowResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DeleteFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFlowResponseBody) SetAccessDeniedDetail(v string) *DeleteFlowResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DeleteFlowResponseBody) SetCode(v string) *DeleteFlowResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFlowResponseBody) SetMessage(v string) *DeleteFlowResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFlowResponseBody) SetRequestId(v string) *DeleteFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFlowResponseBody) Validate() error {
	return dara.Validate(s)
}
