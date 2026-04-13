// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRegistryNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteRegistryNamespaceRequest struct {
}

func (s DeleteRegistryNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRegistryNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteRegistryNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
