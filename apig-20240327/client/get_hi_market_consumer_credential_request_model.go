// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHiMarketConsumerCredentialRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetHiMarketConsumerCredentialRequest struct {
}

func (s GetHiMarketConsumerCredentialRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHiMarketConsumerCredentialRequest) GoString() string {
	return s.String()
}

func (s *GetHiMarketConsumerCredentialRequest) Validate() error {
	return dara.Validate(s)
}
