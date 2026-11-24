// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecretValueRequest interface {
	dara.Model
	String() string
	GoString() string
}

type DeleteSecretValueRequest struct {
}

func (s DeleteSecretValueRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecretValueRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecretValueRequest) Validate() error {
	return dara.Validate(s)
}
