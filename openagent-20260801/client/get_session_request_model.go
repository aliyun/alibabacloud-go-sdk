// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSessionRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetSessionRequest struct {
}

func (s GetSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSessionRequest) GoString() string {
	return s.String()
}

func (s *GetSessionRequest) Validate() error {
	return dara.Validate(s)
}
