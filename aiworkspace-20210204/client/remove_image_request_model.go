// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RemoveImageRequest struct {
}

func (s RemoveImageRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageRequest) Validate() error {
	return dara.Validate(s)
}
