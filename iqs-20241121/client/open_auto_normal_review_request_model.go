// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenAutoNormalReviewRequest interface {
	dara.Model
	String() string
	GoString() string
}

type OpenAutoNormalReviewRequest struct {
}

func (s OpenAutoNormalReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenAutoNormalReviewRequest) GoString() string {
	return s.String()
}

func (s *OpenAutoNormalReviewRequest) Validate() error {
	return dara.Validate(s)
}
