// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogicalModelRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetLogicalModelRequest struct {
}

func (s GetLogicalModelRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLogicalModelRequest) GoString() string {
	return s.String()
}

func (s *GetLogicalModelRequest) Validate() error {
	return dara.Validate(s)
}
