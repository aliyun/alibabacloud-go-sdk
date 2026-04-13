// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDetectConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetDetectConfigRequest struct {
}

func (s GetDetectConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDetectConfigRequest) GoString() string {
	return s.String()
}

func (s *GetDetectConfigRequest) Validate() error {
	return dara.Validate(s)
}
