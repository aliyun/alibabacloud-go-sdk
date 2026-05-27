// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMcpServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteMcpServiceRequest struct {
}

func (s DeleteMcpServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteMcpServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteMcpServiceRequest) Validate() error {
	return dara.Validate(s)
}
