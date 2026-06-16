// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySecurityCheckReportRequest interface {
	dara.Model
	String() string
	GoString() string
}

type QuerySecurityCheckReportRequest struct {
}

func (s QuerySecurityCheckReportRequest) String() string {
	return dara.Prettify(s)
}

func (s QuerySecurityCheckReportRequest) GoString() string {
	return s.String()
}

func (s *QuerySecurityCheckReportRequest) Validate() error {
	return dara.Validate(s)
}
