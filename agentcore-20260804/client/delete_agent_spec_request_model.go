// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgentSpecRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteAgentSpecRequest struct {
}

func (s DeleteAgentSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgentSpecRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgentSpecRequest) Validate() error {
	return dara.Validate(s)
}
