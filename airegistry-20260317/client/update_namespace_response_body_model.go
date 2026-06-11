// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateNamespaceResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateNamespaceResponseBody
	GetRequestId() *string
}

type UpdateNamespaceResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNamespaceResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNamespaceResponseBody) SetData(v bool) *UpdateNamespaceResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateNamespaceResponseBody) SetRequestId(v string) *UpdateNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
