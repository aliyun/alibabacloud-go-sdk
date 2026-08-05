// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetCredentialsRequest struct {
}

func (s GetCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCredentialsRequest) GoString() string {
	return s.String()
}

func (s *GetCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
