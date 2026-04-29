// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnknownThreatDetectStatisticRequest interface {
	dara.Model
	String() string
	GoString() string
}

type GetUnknownThreatDetectStatisticRequest struct {
}

func (s GetUnknownThreatDetectStatisticRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUnknownThreatDetectStatisticRequest) GoString() string {
	return s.String()
}

func (s *GetUnknownThreatDetectStatisticRequest) Validate() error {
	return dara.Validate(s)
}
