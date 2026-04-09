// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishImageRequest interface {
	dara.Model
	String() string
	GoString() string
}

type PublishImageRequest struct {
}

func (s PublishImageRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishImageRequest) GoString() string {
	return s.String()
}

func (s *PublishImageRequest) Validate() error {
	return dara.Validate(s)
}
