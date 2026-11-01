// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccessApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAccessApplicationRequest struct {
}

func (s GetAccessApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAccessApplicationRequest) GoString() string {
	return s.String()
}

func (s *GetAccessApplicationRequest) Validate() error {
	return dara.Validate(s)
}
