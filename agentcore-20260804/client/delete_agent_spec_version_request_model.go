// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpecVersionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAgentSpecVersionRequest struct {
}

func (s DeleteAgentSpecVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpecVersionRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpecVersionRequest) Validate() error {
	return dara.Validate(s)
}
