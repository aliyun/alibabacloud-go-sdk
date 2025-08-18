// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKvNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteKvNamespaceResponseBody
	GetRequestId() *string
}

type DeleteKvNamespaceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// EEEBE525-F576-1196-8DAF-2D70CA3F4D2F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteKvNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteKvNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteKvNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteKvNamespaceResponseBody) SetRequestId(v string) *DeleteKvNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteKvNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
