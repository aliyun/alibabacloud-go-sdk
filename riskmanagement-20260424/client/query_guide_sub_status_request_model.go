// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGuideSubStatusRequest interface {
	dara.Model
	String() string
	GoString() string
}

type QueryGuideSubStatusRequest struct {
}

func (s QueryGuideSubStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryGuideSubStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryGuideSubStatusRequest) Validate() error {
	return dara.Validate(s)
}
