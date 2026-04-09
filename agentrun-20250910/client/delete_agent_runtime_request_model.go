// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentRuntimeRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAgentRuntimeRequest struct {
}

func (s DeleteAgentRuntimeRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentRuntimeRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentRuntimeRequest) Validate() error {
	return dara.Validate(s)
}
