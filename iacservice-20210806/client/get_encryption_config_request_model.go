// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEncryptionConfigRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetEncryptionConfigRequest struct {
}

func (s GetEncryptionConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEncryptionConfigRequest) GoString() string {
	return s.String()
}

func (s *GetEncryptionConfigRequest) Validate() error {
	return dara.Validate(s)
}
