// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegistryNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetRegistryNamespaceRequest struct {
}

func (s GetRegistryNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRegistryNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetRegistryNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
