// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
}

type PublishDatasetRequest struct {
}

func (s PublishDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s PublishDatasetRequest) GoString() string {
	return s.String()
}

func (s *PublishDatasetRequest) Validate() error {
	return dara.Validate(s)
}
