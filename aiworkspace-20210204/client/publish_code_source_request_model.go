// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishCodeSourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type PublishCodeSourceRequest struct {
}

func (s PublishCodeSourceRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishCodeSourceRequest) GoString() string {
	return s.String()
}

func (s *PublishCodeSourceRequest) Validate() error {
	return dara.Validate(s)
}
