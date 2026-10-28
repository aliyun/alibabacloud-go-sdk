// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartQuotaUtilizationReportRequest interface {
	dara.Model
	String() string
	GoString() string
}

type StartQuotaUtilizationReportRequest struct {
}

func (s StartQuotaUtilizationReportRequest) String() string {
	return dara.Prettify(s)
}

func (s StartQuotaUtilizationReportRequest) GoString() string {
	return s.String()
}

func (s *StartQuotaUtilizationReportRequest) Validate() error {
	return dara.Validate(s)
}
