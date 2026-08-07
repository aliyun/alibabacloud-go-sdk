// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGuardLogStatsRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetGuardLogStatsRequest struct {
}

func (s GetGuardLogStatsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetGuardLogStatsRequest) GoString() string {
	return s.String()
}

func (s *GetGuardLogStatsRequest) Validate() error {
	return dara.Validate(s)
}
