// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetManagedAgentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetManagedAgentRequest struct {
}

func (s GetManagedAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetManagedAgentRequest) GoString() string {
	return s.String()
}

func (s *GetManagedAgentRequest) Validate() error {
	return dara.Validate(s)
}
