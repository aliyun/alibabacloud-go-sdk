// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteCredentialsRequest struct {
}

func (s DeleteCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCredentialsRequest) GoString() string {
	return s.String()
}

func (s *DeleteCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
