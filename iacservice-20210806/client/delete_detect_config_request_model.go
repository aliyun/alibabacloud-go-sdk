// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDetectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteDetectConfigRequest struct {
}

func (s DeleteDetectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDetectConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteDetectConfigRequest) Validate() error {
	return dara.Validate(s)
}
