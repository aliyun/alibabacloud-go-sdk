// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceJobRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetInstanceJobRequest struct {
}

func (s GetInstanceJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceJobRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceJobRequest) Validate() error {
	return dara.Validate(s)
}
