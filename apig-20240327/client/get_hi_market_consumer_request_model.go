// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHiMarketConsumerRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetHiMarketConsumerRequest struct {
}

func (s GetHiMarketConsumerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHiMarketConsumerRequest) GoString() string {
	return s.String()
}

func (s *GetHiMarketConsumerRequest) Validate() error {
	return dara.Validate(s)
}
