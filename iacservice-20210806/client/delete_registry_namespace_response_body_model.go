// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryNamespaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRegistryNamespaceResponseBody
	GetRequestId() *string
}

type DeleteRegistryNamespaceResponseBody struct {
	// example:
	//
	// 1D0CD708-E433-5F13-8A42-823C95FC756C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteRegistryNamespaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryNamespaceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRegistryNamespaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRegistryNamespaceResponseBody) SetRequestId(v string) *DeleteRegistryNamespaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRegistryNamespaceResponseBody) Validate() error {
	return dara.Validate(s)
}
