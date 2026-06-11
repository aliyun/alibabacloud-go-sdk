// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteNamespaceResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteNamespaceResponseBody
	GetRequestId() *string
}

type DeleteNamespaceResponseBody struct {
	Data      *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNamespaceResponseBody) SetData(v bool) *DeleteNamespaceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNamespaceResponseBody) SetRequestId(v string) *DeleteNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
