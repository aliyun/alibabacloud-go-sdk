// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceResourceTableRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetInstanceResourceTableRequest struct {
}

func (s GetInstanceResourceTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceResourceTableRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceResourceTableRequest) Validate() error {
	return dara.Validate(s)
}
