// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRayJobRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetRayJobRequest struct {
}

func (s GetRayJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRayJobRequest) GoString() string {
	return s.String()
}

func (s *GetRayJobRequest) Validate() error {
	return dara.Validate(s)
}
