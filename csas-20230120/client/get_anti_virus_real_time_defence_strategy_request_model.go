// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAntiVirusRealTimeDefenceStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetAntiVirusRealTimeDefenceStrategyRequest struct {
}

func (s GetAntiVirusRealTimeDefenceStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAntiVirusRealTimeDefenceStrategyRequest) GoString() string {
	return s.String()
}

func (s *GetAntiVirusRealTimeDefenceStrategyRequest) Validate() error {
	return dara.Validate(s)
}
