// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResourceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetInstanceResourceRequest struct {
}

func (s GetInstanceResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResourceRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceRequest) Validate() error {
	return dara.Validate(s)
}
