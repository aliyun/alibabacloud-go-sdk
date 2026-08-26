// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteManagedAgentRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteManagedAgentRequest struct {
}

func (s DeleteManagedAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteManagedAgentRequest) GoString() string {
	return s.String()
}

func (s *DeleteManagedAgentRequest) Validate() error {
	return dara.Validate(s)
}
