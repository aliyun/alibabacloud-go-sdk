// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMemoryCollectionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMemoryCollectionRequest struct {
}

func (s DeleteMemoryCollectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMemoryCollectionRequest) GoString() string {
	return s.String()
}

func (s *DeleteMemoryCollectionRequest) Validate() error {
	return dara.Validate(s)
}
