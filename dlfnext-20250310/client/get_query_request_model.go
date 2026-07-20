// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQueryRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetQueryRequest struct {
}

func (s GetQueryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetQueryRequest) GoString() string {
	return s.String()
}

func (s *GetQueryRequest) Validate() error {
	return dara.Validate(s)
}
