// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateNamespaceResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateNamespaceResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateNamespaceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateNamespaceResponseBody
	GetSuccess() *bool
}

type UpdateNamespaceResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// The system namespace cannot be modified
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 39090022-1F3B-4797-8518-6B61095F1AF0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateNamespaceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNamespaceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateNamespaceResponseBody) SetCode(v int32) *UpdateNamespaceResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetMessage(v string) *UpdateNamespaceResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetRequestId(v string) *UpdateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetSuccess(v bool) *UpdateNamespaceResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
