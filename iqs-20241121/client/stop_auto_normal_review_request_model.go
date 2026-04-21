// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAutoNormalReviewRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StopAutoNormalReviewRequest struct {
}

func (s StopAutoNormalReviewRequest) String() string {
	return dara.Prettify(s)
}

func (s StopAutoNormalReviewRequest) GoString() string {
	return s.String()
}

func (s *StopAutoNormalReviewRequest) Validate() error {
	return dara.Validate(s)
}
