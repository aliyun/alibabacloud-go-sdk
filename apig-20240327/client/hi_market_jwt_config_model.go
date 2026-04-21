// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHiMarketJwtConfig interface {
	dara.Model
	String() string
	GoString() string
}

type HiMarketJwtConfig struct {
}

func (s HiMarketJwtConfig) String() string {
	return dara.Prettify(s)
}

func (s HiMarketJwtConfig) GoString() string {
	return s.String()
}

func (s *HiMarketJwtConfig) Validate() error {
	return dara.Validate(s)
}
