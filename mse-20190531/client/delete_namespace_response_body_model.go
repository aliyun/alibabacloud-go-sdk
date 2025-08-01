// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteNamespaceResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *DeleteNamespaceResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNamespaceResponseBody
	GetSuccess() *bool
}

type DeleteNamespaceResponseBody struct {
	// The error code.
	//
	// example:
	//
	// NO_PERMISSION
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// abcde-efg
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteNamespaceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNamespaceResponseBody) SetErrorCode(v string) *DeleteNamespaceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetHttpStatusCode(v int32) *DeleteNamespaceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetMessage(v string) *DeleteNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetSuccess(v bool) *DeleteNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
