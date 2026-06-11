// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelRayJobRequest interface {
	dara.Model
	String() string
	GoString() string
}

type CancelRayJobRequest struct {
}

func (s CancelRayJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelRayJobRequest) GoString() string {
	return s.String()
}

func (s *CancelRayJobRequest) Validate() error {
	return dara.Validate(s)
}
