// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretValueRequest interface {
	dara.Model
	String() string
	GoString() string
}

type ListSecretValueRequest struct {
}

func (s ListSecretValueRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretValueRequest) GoString() string {
	return s.String()
}

func (s *ListSecretValueRequest) Validate() error {
	return dara.Validate(s)
}
