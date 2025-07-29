// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteNamespaceResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNamespaceResponseBody
	GetSuccess() *bool
}

type DeleteNamespaceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The system namespace cannot be deleted
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4F68ABED-AC31-4412-9297-D9A8F0401107
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

func (s *DeleteNamespaceResponseBody) GetCode() *int32 {
	return s.Code
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

func (s *DeleteNamespaceResponseBody) SetCode(v int32) *DeleteNamespaceResponseBody {
	s.Code = &v
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
