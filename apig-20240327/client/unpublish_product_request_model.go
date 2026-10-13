// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnpublishProductRequest interface {
	dara.Model
	String() string
	GoString() string
}

type UnpublishProductRequest struct {
}

func (s UnpublishProductRequest) String() string {
	return dara.Prettify(s)
}

func (s UnpublishProductRequest) GoString() string {
	return s.String()
}

func (s *UnpublishProductRequest) Validate() error {
	return dara.Validate(s)
}
