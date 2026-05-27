// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpServiceRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMcpServiceRequest struct {
}

func (s GetMcpServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcpServiceRequest) GoString() string {
	return s.String()
}

func (s *GetMcpServiceRequest) Validate() error {
	return dara.Validate(s)
}
