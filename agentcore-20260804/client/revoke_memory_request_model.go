// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeMemoryRequest interface {
	dara.Model
	String() string
	GoString() string
}

type RevokeMemoryRequest struct {
}

func (s RevokeMemoryRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeMemoryRequest) GoString() string {
	return s.String()
}

func (s *RevokeMemoryRequest) Validate() error {
	return dara.Validate(s)
}
