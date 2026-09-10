// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaOssTempKeyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetMetaOssTempKeyRequest struct {
}

func (s GetMetaOssTempKeyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaOssTempKeyRequest) GoString() string {
	return s.String()
}

func (s *GetMetaOssTempKeyRequest) Validate() error {
	return dara.Validate(s)
}
